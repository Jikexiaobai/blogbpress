import axios from '@/utils/request'
import api from './index'

/**
 * get Manger 2step code open?
 * @param parameter {*}
 */
export function getAccount () {
  return axios({
    url: api.getAccount,
    method: 'get',
  })
}


// 获取列表
/* eslint-disable */
export function MangerAuthority (parameter) {
  return axios({
      url: api.getMangerAuthority,
      method: 'get',
      params: parameter
  })
}



// 获取列表
/* eslint-disable */
export function List (parameter) {
  return axios({
      url: api.getMangerList,
      method: 'get',
      params: parameter
  })
}

// 创建
/* eslint-disable */
export function CreateMeta () {
  return axios({
      url: api.getMangerMeta,
      method: 'get',
  })
}




/* eslint-disable */
export function Create (formData) {
  return axios({
      url: api.postMangerCreate,
      method: 'post',
      data: formData
  })
}

// 
/* eslint-disable */
export function EditInfo (parameter) {
  return axios({
      url: api.getMangerEditInfo,
      method: 'get',
      params: parameter
  })
}


/* eslint-disable */
export function Edit (formData) {
  return axios({
      url: api.postMangerEdit,
      method: 'post',
      data: formData
  })
}



// 删除
/* eslint-disable */
export function Remove (formData) {
  return axios({
      url: api.postRemoveManger,
      method: 'post',
      data: formData
  })
}


/* eslint-disable */
export function Review (formData) {
  return axios({
      url: api.postReviewManger,
      method: 'post',
      data: formData
  })
}