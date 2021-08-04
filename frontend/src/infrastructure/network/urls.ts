/**
 ******************************
 * Список внешних маршрутов
 ******************************
 */
const baseURL = '/api/'

const mapUrls = {
  api_v1: {
    LOGIN: 'auth/login',
    LOGOUT: 'auth/logout',
    REGISTER: 'auth/register',
    REFRESH_TOKEN: 'auth/refresh',
    CHECK_LOGGED: 'auth/check',
    CONFIRMED: 'auth/confirmed',
    USER: 'user',
    ASSETS: 'asset/all',
    TOTAL_ASSETS: 'asset/total-all',
    ASSET_ADD: 'asset/add',
    MARKETS: 'market/all',
    SEARCH_MARKETS: 'market/search',
    PORTFOLIO_ACTIVE_TOTAL: 'portfolio/active-total',
    PORTFOLIO_ALL_STATS: 'portfolio/all-stats',
    PORTFOLIO_ADD: 'portfolio/add',
    PORTFOLIO_CHOOSE: 'portfolio/{portfolioId}/choose',
    APP_DASHBOARD: 'application/dashboard',
  },
  ws: {
    BITFINEX: 'wss://api.bitfinex.com/ws',
    QUOTES: 'ws://localhost:8081/ws/market/quotes',
  },
}

/**
 ******************************
 * Версии для маршрутов
 ******************************
 */
const versions = {
  api_v1: 'v1/',
}

/**
 ******************************
 * Версионирование маршрутов
 *
 * @param {versions}
 ******************************
 */
type Urls = typeof mapUrls
type Versions = typeof versions
type ProxyPaths = { [key: string]: unknown }

const hasVersion = (type: string): type is keyof Versions => type in versions

const proxyUrls = (proxyPaths: ProxyPaths, [type, paths]: [string, any]) => {
  const version = hasVersion(type) ? versions[type] : ''
  proxyPaths[type] = new Proxy(paths, {
    get: (target, prop) => version + target[prop],
  })
  return proxyPaths
}

const urls = Object.entries(mapUrls).reduce(proxyUrls, {}) as Urls

/**
 ******************************
 * Склеить маршрут по шаблону
 ******************************
 */
const parsePatternUrl = (url: string, ...params: string[]): string =>
  url.replace(/({\S+?})/gi, () => params.shift() as string)

export { urls, baseURL, parsePatternUrl }
