import axios from '@/utils/request'
import api from './index'


// 获取列表
/* eslint-disable */
export function List (parameter) {
  return axios({
      url: api.getUserList,
      method: 'get',
      params: parameter
  })
}

// 创建
/* eslint-disable */
export function CreateMeta () {
  return axios({
      url: api.getUserMeta,
      method: 'get',
  })
}




/* eslint-disable */
export function Create (formData) {
  return axios({
      url: api.postUserCreate,
      method: 'post',
      data: formData
  })
}

// 
/* eslint-disable */
export function EditInfo (parameter) {
  return axios({
      url: api.getUserEditInfo,
      method: 'get',
      params: parameter
  })
}


/* eslint-disable */
export function Edit (formData) {
  return axios({
      url: api.postUserEdit,
      method: 'post',
      data: formData
  })
}



// 删除
/* eslint-disable */
export function Remove (formData) {
  return axios({
      url: api.postRemoveUser,
      method: 'post',
      data: formData
  })
}


/* eslint-disable */
export function Review (formData) {
  return axios({
      url: api.postReviewUser,
      method: 'post',
      data: formData
  })
}