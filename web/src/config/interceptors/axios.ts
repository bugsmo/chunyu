import type { InternalAxiosRequestConfig, AxiosRequestConfig, AxiosResponse } from 'axios'

export function requestSuccessFunc(requestObj: InternalAxiosRequestConfig) {
  // 自定义请求拦截逻辑，可以处理权限，请求发送监控等
  // ...

  return requestObj
}

export function requestFailFunc(requestError: AxiosRequestConfig) {
  // 自定义发送请求失败逻辑，断网，请求发送监控等
  // ...

  return Promise.reject(requestError)
}

export function responseSuccessFunc(responseObj: AxiosResponse) {
  // 自定义响应成功逻辑，全局拦截接口，根据不同业务做不同处理，响应成功监控等
  // ...

  return responseObj
}

export function responseFailFunc(responseError: AxiosResponse) {
  // 响应失败，可根据 responseError.status 来做监控处理
  // ...

  return Promise.reject(responseError)
}
