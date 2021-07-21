export default {
  encode: (str: string): string => btoa(encodeURIComponent(str)),
  decode: (str: string): string => decodeURIComponent(atob(str)),
}
