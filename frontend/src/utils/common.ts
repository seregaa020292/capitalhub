export const byteToSize = (bytes: number, decimals = 2): string => {
  if (bytes === 0) return '0 Bytes'

  const k = 1024
  const dm = decimals < 0 ? 0 : decimals
  const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']

  const i = Math.floor(Math.log(bytes) / Math.log(k))

  return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i]
}

export const sleep = (timeMs = 0) => new Promise(resolve => setTimeout(resolve, timeMs))

export const isDigit = (value: string): boolean => !(/^[0-9]*$/.exec(value) == null || value === '')

export const isUUID = (uuid: string): boolean =>
  new RegExp('\b[0-9a-f]{8}\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\b[0-9a-f]{12}\b').test(uuid)

export const numberToArrayNumber = (number: number) => Array.from(String(number), Number)

export const uniqueId = (): number => {
  const a = Math.random
  const b: any = parseInt
  return Number(Number(new Date()).toString() + b(10 * a()) + b(10 * a()) + b(10 * a()))
}

export const getBytesLength = (s: string): number =>
  // eslint-disable-next-line no-control-regex
  s.replace(/[^\x00-\xff]/gi, '--').length

export const randomElement = (arr = []): any => arr[Math.floor(Math.random() * arr.length)]

export const kebab = (str: string): string =>
  (str || '').replace(/([a-z])([A-Z])/g, '$1-$2').toLowerCase()

export const firstTwoLetters = (name = ''): string =>
  name
    .replace(/\s+/, ' ')
    .split(' ')
    .slice(0, 2)
    .map((v) => v && v[0].toUpperCase())
    .join('')

export const htmlEncode = (text: string): string =>
  text
    .replace(/&/g, '&')
    .replace(/"/g, '"')
    .replace(/</g, '<')
    .replace(/>/g, '>')

/**
 * https://stackoverflow.com/questions/35228052/debounce-function-implemented-with-promises
 * @param inner
 * @param ms
 * @returns {function(...[*]): Promise<unknown>}
 * @private
 */
export const debounce = (inner: () => unknown, ms = 0): (() => Promise<unknown>) => {
  let timer: NodeJS.Timeout | null = null
  let resolves = [] as Array<(value: unknown) => void>

  return function () {
    if (timer) {
      clearTimeout(timer)
    }
    timer = setTimeout(() => {
      const result = inner()
      resolves.forEach((resolve) => resolve(result))
      resolves = []
    }, ms)

    return new Promise((resolve) => resolves.push(resolve))
  }
}
