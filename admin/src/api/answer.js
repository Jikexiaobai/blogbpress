import { axios } from '@/utils/request'
import api from './index'


// 获取列表
/* eslint-disable */
export function List (parameter) {
    return axios({
        url: api.getAnswerList,
        method: 'get',
        params: parameter
    })
}


// 还原
/* eslint-disable */
export function Recover (formData) {
    return axios({
        url: api.postRecoverAnswer,
        method: 'post',
        data: formData
    })
}


// 还原
/* eslint-disable */
export function Reduction (formData) {
    return axios({
        url: api.postReductionAnswer,
        method: 'post',
        data: formData
    })
}


// 删除
/* eslint-disable */
export function Remove (formData) {
    return axios({
        url: api.postRemoveAnswer,
        method: 'post',
        data: formData
    })
}


/* eslint-disable */
export function Review (formData) {
    return axios({
        url: api.postReviewAnswer,
        method: 'post',
        data: formData
    })
}
