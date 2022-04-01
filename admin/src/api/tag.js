import { axios } from '@/utils/request'
import api from './index'

// 获取列表
/* eslint-disable */
export function Hots (parameter) {
    return axios({
        url: api.getTagHots,
        method: 'get',
    })
}


// 获取列表
/* eslint-disable */
export function List (parameter) {
    return axios({
        url: api.getTagList,
        method: 'get',
        params: parameter
    })
}

// 删除
/* eslint-disable */
export function Top (formData) {
    return axios({
        url: api.postTagTop,
        method: 'post',
        data: formData
    })
}

// 删除
/* eslint-disable */
export function Remove (formData) {
    return axios({
        url: api.postTagRemove,
        method: 'post',
        data: formData
    })
}