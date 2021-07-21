export const getScrollTop = (): number =>
  Math.max(
    //chrome
    document.body.scrollTop,
    //firefox/IE
    document.documentElement.scrollTop
  )

export const getPageHeight = (): number => {
  const g = document
  const a = g.body
  const f = g.documentElement
  const d = g.compatMode === 'BackCompat' ? a : g.documentElement
  return Math.max(f.scrollHeight, a.scrollHeight, d.clientHeight)
}

export const getPageViewHeight = (): number => {
  const d = document
  const a = d.compatMode === 'BackCompat' ? d.body : d.documentElement
  return a.clientHeight
}

export const getPageViewWidth = (): number => {
  const d = document
  const a = d.compatMode === 'BackCompat' ? d.body : d.documentElement
  return a.clientWidth
}

export const getPageWidth = (): number => {
  const g = document
  const a = g.body
  const f = g.documentElement
  const d = g.compatMode === 'BackCompat' ? a : g.documentElement
  return Math.max(f.scrollWidth, a.scrollWidth, d.clientWidth)
}

export const getViewSize = (): Array<number> => {
  const de = document.documentElement
  const db = document.body
  const viewW = de.clientWidth === 0 ? db.clientWidth : de.clientWidth
  const viewH = de.clientHeight === 0 ? db.clientHeight : de.clientHeight
  return [viewW, viewH]
}

export const copyText = (text: string): void => {
  if (document !== null) {
    const tag = document.createElement('textarea')
    tag.setAttribute('id', 'cp_input')
    tag.value = text
    tag.style.opacity = '0'
    tag.style.position = 'absolute'
    document.getElementsByTagName('body')[0].appendChild(tag)
    const target: any = document.getElementById('cp_input')
    if (target) {
      target.select()
      document.execCommand('copy')
      target.remove()
    }
  }
}

export const toggleFullScreen = (): void => {
  const doc = window.document as Document & {
    mozCancelFullScreen(): Promise<void>
    webkitExitFullscreen(): Promise<void>
    msExitFullscreen(): Promise<void>
    mozFullScreenElement(): Promise<void>
    webkitFullscreenElement(): Promise<void>
    msFullscreenElement(): Promise<void>
  }
  const docEl = doc.documentElement as HTMLElement & {
    mozRequestFullScreen(): Promise<void>
    webkitRequestFullScreen(): Promise<void>
    msRequestFullscreen(): Promise<void>
  }

  const requestFullScreen =
    docEl.requestFullscreen ||
    docEl.mozRequestFullScreen ||
    docEl.webkitRequestFullScreen ||
    docEl.msRequestFullscreen

  const cancelFullScreen =
    doc.exitFullscreen ||
    doc.mozCancelFullScreen ||
    doc.webkitExitFullscreen ||
    doc.msExitFullscreen

  if (
    !doc.fullscreenElement &&
    !doc.mozFullScreenElement &&
    !doc.webkitFullscreenElement &&
    !doc.msFullscreenElement
  ) {
    requestFullScreen.call(docEl)
  } else {
    cancelFullScreen.call(doc)
  }
}
