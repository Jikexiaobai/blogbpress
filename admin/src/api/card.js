import { axios } from '@/utils/request'
import api from './index'


// 获取列表
/* eslint-disable */
export function List (parameter) {
    return axios({
        url: api.getCardList,
        method: 'get',
        params: parameter
    })
}


// 删除
/* eslint-disable */
export function Remove (formData) {
    return axios({
        url: api.postRemoveCard,
        method: 'post',
        data: formData
    })
}


/* eslint-disable */
export function Create (formData) {
    return axios({
        url: api.postCreateCard,
        method: 'post',
        data: formData
    })
}
