import { axios } from '@/utils/request'
import api from './index'


// 获取列表
/* eslint-disable */
export function List (parameter) {
    return axios({
        url: api.getOrderList,
        method: 'get',
        params: parameter
    })
}

// 
/* eslint-disable */
export function Info (parameter) {
    return axios({
        url: api.getOrderInfo,
        method: 'get',
        params: parameter
    })
}


// 删除
/* eslint-disable */
export function Remove (formData) {
    return axios({
        url: api.postRemoveOrder,
        method: 'post',
        data: formData
    })
}


/* eslint-disable */
export function Review (formData) {
    return axios({
        url: api.postReviewOrder,
        method: 'post',
        data: formData
    })
}
