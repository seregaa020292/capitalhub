import info from '../../../setup/data/project.json'

export interface IApp {
  projectName: string
  baseUrl: string
  redirectUri: string
  xsrfHeaderName: string
  isMock: boolean
  isProd: boolean
  isDev: boolean
}

const config: IApp = {
  projectName: info.name,
  baseUrl: import.meta.env.BASE_URL,
  redirectUri: '/login',
  xsrfHeaderName: 'x-csrf-token',
  isMock: false,
  isProd: import.meta.env.PROD,
  isDev: import.meta.env.DEV,
}

if (config.isDev) {
  // config.redirectUri = '/login'
}

export { config }
