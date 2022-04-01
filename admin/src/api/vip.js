import { axios } from '@/utils/request'
import api from './index'


// 获取列表
/* eslint-disable */
export function List (parameter) {
    return axios({
        url: api.getVipList,
        method: 'get',
        params: parameter
    })
}



/* eslint-disable */
export function Create (formData) {
    return axios({
        url: api.postVipCreate,
        method: 'post',
        data: formData
    })
}

// 
/* eslint-disable */
export function EditInfo (parameter) {
    return axios({
        url: api.getVipEditInfo,
        method: 'get',
        params: parameter
    })
}


/* eslint-disable */
export function Edit (formData) {
    return axios({
        url: api.postVipEdit,
        method: 'post',
        data: formData
    })
}



// 删除
/* eslint-disable */
export function Remove (formData) {
    return axios({
        url: api.postRemoveVip,
        method: 'post',
        data: formData
    })
}