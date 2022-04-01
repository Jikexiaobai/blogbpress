import { axios } from '@/utils/request'
import api from './index'


// 获取列表
/* eslint-disable */
export function List (parameter) {
    return axios({
        url: api.getTopicList,
        method: 'get',
        params: parameter
    })
}




/* eslint-disable */
export function Create (formData) {
    return axios({
        url: api.postTopicCreate,
        method: 'post',
        data: formData
    })
}

// 还原
/* eslint-disable */
export function Recover (formData) {
    return axios({
        url: api.postRecoverTopic,
        method: 'post',
        data: formData
    })
}


// 还原
/* eslint-disable */
export function Reduction (formData) {
    return axios({
        url: api.postReductionTopic,
        method: 'post',
        data: formData
    })
}


// 删除
/* eslint-disable */
export function Remove (formData) {
    return axios({
        url: api.postRemoveTopic,
        method: 'post',
        data: formData
    })
}


/* eslint-disable */
export function Review (formData) {
    return axios({
        url: api.postReviewTopic,
        method: 'post',
        data: formData
    })
}

/* eslint-disable */
export function Top (formData) {
    return axios({
        url: api.postTopTopic,
        method: 'post',
        data: formData
    })
}