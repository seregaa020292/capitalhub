import AxiosMockAdapter from 'axios-mock-adapter'
import urls from '@/infrastructure/network/urls'
import { serverResponse } from '@/utils/server'
import { setCookie } from '@/utils/cookie'

export default (mock: AxiosMockAdapter): void => {
  mock.onGet(urls.api_v1.USER).reply(
    200,
    serverResponse({
      message: {
        id: 'number',
        name: 'string',
        email: 'string',
        avatar: 'string',
      },
    })
  )

  mock.onPost(urls.api_v1.REGISTER).reply((config) => {
    setCookie('registerData', JSON.stringify(config.data), 1)
    setCookie('refreshToken', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9', 24 * 7)
    return [
      201,
      serverResponse({
        message: {
          prefixToken: 'Bearer',
          accessToken:
            'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c',
        },
      }),
    ]
  })
}
