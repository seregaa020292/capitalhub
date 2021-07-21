import EventEmitter from 'events'
import { dayjs } from '@/utils/dayjs'
import { config } from '@/data/config/app'
import { sleep } from '@/utils/common'

interface IOptions {
  protocols?: string | string[]
  debug?: boolean
  maxAttempt?: number
}

export type Unsubscribe = {
  remove: () => void
}

/**
 * @link https://github.com/Luka967/websocket-close-codes
 */
enum CloseCodes {
  CLOSE_NORMAL = 1000, // Successful operation / regular socket shutdown
  CLOSE_GOING_AWAY = 1001, // Client is leaving (browser tab closing)
  CLOSED_NO_STATUS = 1005, // Expected close status, received none
}

export class Socket {
  private ws: WebSocket | null = null
  private ee = new EventEmitter()
  private success = new Map()
  private failure = new Map()
  private reconnected = true
  private numAttempt = 0
  private maxAttempt!: number

  public constructor(private url: string, private options: IOptions = {}) {}

  public init(): void {
    this.ws = new WebSocket(this.url, this.options.protocols)
    this.ws.binaryType = 'arraybuffer'
    this.ws.onopen = this.onopen.bind(this)
    this.ws.onclose = this.onclose.bind(this)
    this.ws.onerror = this.onerror.bind(this)
    this.ws.onmessage = this.onmessage.bind(this)
    this.maxAttempt = this.options.maxAttempt || 100
    this.debug('>> WebSocket init:', this.url)
  }

  public subscribe(
    name: string,
    message: unknown,
    callback: (...args: any[]) => void,
  ): Unsubscribe {
    let ee: EventEmitter
    if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
      ee = this.failurePush(name, message, callback)
    } else {
      ee = this.successPush(name, message, callback)
    }
    return {
      remove: () => {
        this.unsubscribe(name)
        ee.removeAllListeners(name)
      },
    }
  }

  public unsubscribe(name: string): void {
    if (this.failure.has(name)) {
      this.failure.delete(name)
    }
    if (!this.success.has(name)) {
      return
    }
    if (!this.ws) {
      this.success.delete(name)
      return
    }
    const unsub = JSON.parse(this.success.get(name))
    unsub.cmd = 'unsub'
    const message = JSON.stringify(unsub)
    this.ws.send(message)
    this.ee.removeAllListeners(name)
    this.debug('>> WebSocket send:', message)
    this.success.delete(name)
  }

  public terminate(code?: number, message?: string): void {
    this.success.clear()
    this.failure.clear()
    this.ee.removeAllListeners()
    this.downReconnected()
    this.ws?.close(code || 1000, message)
  }

  private onopen() {
    this.debug('>> WebSocket open...')

    if (this.isIncAttempt) {
      // Повторная подписка после отключения
      this.success.forEach((msg) => {
        this.ws?.send(msg)
        this.debug('>> WebSocket send:', msg)
      })
    }
    // Не удалось повторно подписаться
    this.failure.forEach((msg, key) => {
      if (this.success.has(key)) {
        return
      }
      this.ws?.send(msg)
      this.success.set(key, msg)
      this.debug('>> WebSocket send:', msg)
    })

    this.zeroAttempt()
    this.upReconnected()
    this.failure.clear()
  }

  private onmessage(event: MessageEvent<string>) {
    if (!event.data) {
      return
    }
    // Декомпрессия данных с pako js
    // const text = pako.inflate(event.data, {
    //   to: 'string',
    // })
    const data = JSON.parse(event.data)
    // console.log('---inflate----', data)
    if (data && data.ping) {
      this.ws?.send(
        JSON.stringify({
          pong: Date.now(),
        }),
      )
      return
    }
    this.onBroadcast(data)
  }

  /**
   * Уведомление о трансляции
   */
  private onBroadcast(msg: any) {
    // if (!this.success[msg.ch]) {
    //   return
    // }
    this.ee.emit(msg.ch || '', msg)
  }

  private successPush(name: string, message: unknown, callback: (...args: unknown[]) => void) {
    this.success.set(name, JSON.stringify(message))
    this.ws?.send(this.success.get(name))
    this.debug('>> WebSocket send:', this.success.get(name))
    return this.ee.on(name, callback)
  }

  private failurePush(name: string, message: unknown, callback: (...args: unknown[]) => void) {
    this.failure.set(name, JSON.stringify(message))
    this.debug('>> WebSocket ready to subscribe:', this.failure.get(name))
    return this.ee.on(name, callback)
  }

  private onReconnection() {
    if (!this.url) {
      return
    }

    this.incAttempt()
    if (this.isMaxAttempt) {
      this.terminate()
      return
    }
    sleep(500).then(() => {
      this.init()
      this.debug('>> WebSocket reconnect:', dayjs().format('MM-DD HH:mm:ss'))
    })
  }

  private onclose(event: CloseEvent) {
    this.ws = null

    if (Socket.closeCodeList.includes(event.code)) {
      this.debug('>> Websocket normal close:', CloseCodes[event.code])
      return
    }

    this.debug('>> Websocket close...')

    if (this.reconnected) {
      this.onReconnection()
    }
  }

  private onerror(event: Event) {
    this.debug('>> Websocket error:', event)
  }

  private isConnectOpen(): boolean {
    return this.ws!.readyState === WebSocket.OPEN
  }

  private upReconnected(): void {
    this.reconnected = true
  }

  private downReconnected(): void {
    this.reconnected = false
  }

  private static get closeCodeList(): Array<number> {
    return Object.values(CloseCodes).filter(
      (code: string | number): code is number => typeof code === 'number',
    )
  }

  private get isMaxAttempt(): boolean {
    this.debug(`>> Num Attempt [${this.numAttempt}], max attempt [${this.maxAttempt}]`)
    return this.numAttempt >= this.maxAttempt
  }

  private get isIncAttempt(): boolean {
    return this.numAttempt > 0
  }

  private incAttempt(): void {
    this.numAttempt++
  }

  private zeroAttempt(): void {
    this.numAttempt = 0
  }

  private debug(msg: string | number, ...args: unknown[]) {
    if (this.options.debug || config.isDev) {
      console.log.apply(console, [`%c ${msg} `, 'background: #222; color: lime', ...args])
    }
  }
}
