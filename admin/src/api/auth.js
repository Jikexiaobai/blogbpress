import axios from '@/utils/request'
import api from './index'

/**
 * login func
 * parameter: {
 *     username: '',
 *     password: '',
 *     remember_me: true,
 *     captcha: '12345'
 * }
 * @param parameter
 * @returns {*}
 */
 export function login (parameter) {
    return axios({
      url: api.Login,
      method: 'post',
      data: parameter
    })
  }

  // export function logout () {
  //   return axios({
  //     url: api.Logout,
  //     method: 'post',
  //   })
  // }
  
  /**
   * get user 2step code open?
   * @param parameter {*}
   */
  // 获取验证码
  export function getCaptcha () {
    return axios({
      url: api.getCaptcha,
      method: 'get',
    })
  }
  
  