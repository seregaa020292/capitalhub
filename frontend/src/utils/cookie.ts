export const getCookie = (name: string): string | null => {
  const arr = document.cookie.match(new RegExp('(^| )' + name + '=([^;]*)(;|$)'))
  if (arr != null) return unescape(arr[2])
  return null
}

export const setCookie = (name: string, value: string, hours: number): void => {
  const d = new Date()
  const offset = 8
  const utc = d.getTime() + d.getTimezoneOffset() * 60000
  const nd = utc + 3600000 * offset
  const exp = new Date(nd)
  const domain = window.location.hostname.split('.').reverse().slice(0, 2).reverse().join('.')
  exp.setTime(exp.getTime() + hours * 60 * 60 * 1000)
  document.cookie =
    name + '=' + escape(value) + ';path=/;expires=' + exp.toUTCString() + ';domain=' + domain + ';'
}
