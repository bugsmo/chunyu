import axios from './axios'
import { API_DEFAULT_CONFIG } from 'Config/index'
import type { AxiosRequestConfig } from 'axios'

// 根据当前环境设置baseUrl
const mockBaseUrl = import.meta.env.DEV ? API_DEFAULT_CONFIG.mockBaseUrl : ''

// 封装API请求
const API = (option: AxiosRequestConfig) => {
  option['url'] = mockBaseUrl + option.url
  if (option.method?.toLowerCase() === 'get') {
    option['params'] = option.data
  }

  return axios(option)
}

export default API
