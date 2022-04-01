import { axios } from '@/utils/request'
import api from './index'


// 获取列表
/* eslint-disable */
export function List (parameter) {
    return axios({
        url: api.getQuestionList,
        method: 'get',
        params: parameter
    })
}




/* eslint-disable */
export function Create (formData) {
    return axios({
        url: api.postQuestionCreate,
        method: 'post',
        data: formData
    })
}

// 
/* eslint-disable */
export function EditInfo (parameter) {
    return axios({
        url: api.getQuestionEditInfo,
        method: 'get',
        params: parameter
    })
}


/* eslint-disable */
export function Edit (formData) {
    return axios({
        url: api.postQuestionEdit,
        method: 'post',
        data: formData
    })
}

// 还原
/* eslint-disable */
export function Recover (formData) {
    return axios({
        url: api.postRecoverQuestion,
        method: 'post',
        data: formData
    })
}


// 还原
/* eslint-disable */
export function Reduction (formData) {
    return axios({
        url: api.postReductionQuestion,
        method: 'post',
        data: formData
    })
}


// 删除
/* eslint-disable */
export function Remove (formData) {
    return axios({
        url: api.postRemoveQuestion,
        method: 'post',
        data: formData
    })
}


/* eslint-disable */
export function Review (formData) {
    return axios({
        url: api.postReviewQuestion,
        method: 'post',
        data: formData
    })
}