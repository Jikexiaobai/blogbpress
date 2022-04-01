import { axios } from '@/utils/request'
import api from './index'
export function postUploadFile (formData) {
  return axios({
    url: api.postuploadFile,
    method: 'post',
    headers: {
      'Content-Type': 'application/json'
    },
    data:formData
  })
}

export function getvalidChunk (parameter) {
  return axios({
    url: api.getvalidChunk,
    method: 'get',
    params: parameter
  })
}

export function postmergeChunk (formData) {
  return axios({
    url: api.postmergeChunk,
    method: 'post',
    data: formData,
  })
}

export function postUploadChunk (formData) {
  return axios({
    url: api.postuploadChunk,
    method: 'post',
    headers: {
      'Content-Type': 'application/json'
    },
    cancelToken:formData.cancelToken,
    data: formData.formData,
    onUploadProgress: formData.onUploadProgress
  })
}


