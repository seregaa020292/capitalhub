import axios, { AxiosError, AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios'
import { MessageService } from '@/services/message/MessageService'
import {
  AuthPresenterContainer,
  AuthServiceContainer,
  StorageServiceContainer,
} from '@/infrastructure/di/containers'
import { baseURL } from '@/infrastructure/network/urls'
import { responseReject } from '@/utils/server'
import { config as configApp} from '@/data/config/app'

export interface HttpResponse {
  status?: number
  message?: any
}

export interface HttErrorResponse {
  status: number
  error: unknown
}

interface extendAxiosRequestConfig extends AxiosRequestConfig {
  _isRetry?: boolean
}

/**
 ******************************
 * Create instance Axios
 ******************************
 */
const createAxiosInstance = () => axios.create({
  baseURL,
  timeout: 10000,
  withCredentials: true,
  // headers: { 'Access-Control-Allow-Origin': '*' },
})

const http: AxiosInstance = createAxiosInstance()
const httpEasy: AxiosInstance = createAxiosInstance()

/**
 ******************************
 * Request interceptors
 ******************************
 */
const requestOnFulfilled = (config: AxiosRequestConfig) => {
  const accessTokenWithPrefix = StorageServiceContainer().getAccessTokenWithPrefix()
  if (accessTokenWithPrefix !== '') {
    config.headers.Authorization = accessTokenWithPrefix
  }
  const csrf = AuthPresenterContainer().getCsrf()
  if (csrf !== '') {
    config.headers[configApp.xsrfHeaderName] = csrf
  }
  return config
}

const requestOnRejected = (error: AxiosError) => {
  console.error('Request error: ', error)
  Promise.reject(error)
}

/**
 ******************************
 * Response interceptors
 ******************************
 */
const responseOnFulfilled = (res: AxiosResponse) => {
  const data = res.data

  if (res.status.toString().startsWith('4')) {
    MessageService.instance().error('Ошибка на сервере')
    return Promise.reject(new Error(data.message || 'Error'))
  }

  return data
}

const responseOnRejected = async (error: AxiosError) => {
  if (![401, 403].includes(Number(error?.response?.status))) {
    return responseReject(error)
  }

  const originalRequest: extendAxiosRequestConfig = error.config
  if (!originalRequest._isRetry) {
    originalRequest._isRetry = true

    try {
      await AuthServiceContainer().refreshToken()
      originalRequest.headers.Authorization = StorageServiceContainer().getAccessTokenWithPrefix()
      originalRequest.headers[configApp.xsrfHeaderName] = AuthPresenterContainer().getCsrf()
      return await httpEasy.request(originalRequest)
    } catch (error) {
      if ([401, 403].includes(Number(error?.data?.status))) {
        await AuthServiceContainer().autoLogout()
        return Promise.reject(new Error('401'))
      }
      throw error
    }
  }
}

const responseOnRejectedEasy = async (error: AxiosError) => {
  return responseReject(error)
}

/**
 ******************************
 * Use interceptors
 ******************************
 */
http.interceptors.request.use(requestOnFulfilled, requestOnRejected)
http.interceptors.response.use(responseOnFulfilled, responseOnRejected)

httpEasy.interceptors.request.use(requestOnFulfilled, requestOnRejected)
httpEasy.interceptors.response.use(responseOnFulfilled, responseOnRejectedEasy)

export { http, httpEasy }
