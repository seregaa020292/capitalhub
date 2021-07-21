import { AxiosError } from 'axios'
import { ResponseException, ServerException } from '@/infrastructure/exceptions'
import { config } from '@/data/config/app'

export const responseReject = (error: AxiosError) => {
  if (error?.response?.status.toString().startsWith('5')) {
    return Promise.reject(new ServerException(error.response.status))
  }
  if (error?.response?.data) {
    return Promise.reject(new ResponseException(error.response.data))
  }
  return Promise.reject(error)
}

export const responseWithCSRF = () => ({
  transformResponse: (data: string, headers: any) => {
    let response
    try {
      response = JSON.parse(data)
    } catch (error) {
      response = {}
      console.error(error.message)
    }
    response.csrf = headers[config.xsrfHeaderName]
    return response
  },
})

export const serverResponse = ({
  message = '',
  status = 0,
}: {
  message: unknown
  status?: number
}): unknown => ({ message, status })

export const serverErrorResponse = ({
  error = '',
  status = 0,
}: {
  error: unknown
  status?: number
}): unknown => ({ error, status })

export const getServerTime = (): number => {
  let xmlHttp: XMLHttpRequest
  const serverTime = () => {
    try {
      // FF, Opera, Safari, Chrome
      xmlHttp = new XMLHttpRequest()
    } catch (err) {
      //IE
      try {
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-ignore
        xmlHttp = new ActiveXObject('Msxml2.XMLHTTP')
      } catch (err) {
        try {
          // eslint-disable-next-line @typescript-eslint/ban-ts-comment
          // @ts-ignore
          xmlHttp = new ActiveXObject('Microsoft.XMLHTTP')
        } catch (err) {
          //AJAX not supported, use CPU time.
          console.error('AJAX not supported')
        }
      }
    }
    xmlHttp.open('HEAD', window.location.href.toString(), false)
    xmlHttp.setRequestHeader('Content-Type', 'text/html')
    xmlHttp.send('')
    return xmlHttp.getResponseHeader('Date')
  }

  return new Date(serverTime() as string).getTime()
}
