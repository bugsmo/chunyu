import type { AxiosRequestConfig } from 'axios'
import $api from 'Plugins/api'

/** 登录 POST */
interface LoginDto {
  username: string
  password: string
}

// 这里约定所有的接口方法名前加个“$”前缀，跟普通方法名区分开
export async function $authLogin(data?: LoginDto, options?: AxiosRequestConfig) {
  return $api({
    url: '/user/login',
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data,
    ...(options || {}),
  })
}
