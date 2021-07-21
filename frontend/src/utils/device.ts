import FingerprintJS from '@fingerprintjs/fingerprintjs'

export const isMobile = (): boolean =>
  /iphone|ipod|android.*mobile|windows.*phone|blackberry.*mobile/i.test(
    window.navigator.userAgent.toLowerCase()
  )

export const isAndroidMobileDevice = (): boolean =>
  /android/i.test(navigator.userAgent.toLowerCase())

export const isAppleMobileDevice = (): boolean =>
  /iphone|ipod|ipad|Macintosh/i.test(navigator.userAgent.toLowerCase())

export const isViewportOpen = (): boolean => !!document.getElementById('wixMobileViewport')

export const initFingerprintJS = async () => {
  const fpPromise = FingerprintJS.load()
  const fp = await fpPromise
  return await fp.get()
}
