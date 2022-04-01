import { axios } from '@/utils/request'
import api from './index'


// 获取列表
/* eslint-disable */
export function List (parameter) {
    return axios({
        url: api.getAuthorityList,
        method: 'get',
        params: parameter
    })
}

// 创建
/* eslint-disable */
export function Create (formData) {
    return axios({
        url: api.postaAuthorityCreate,
        method: 'post',
        data: formData
    })
}


// 
/* eslint-disable */
export function EditInfo (parameter) {
    return axios({
        url: api.getAuthorityEditInfo,
        method: 'get',
        params: parameter
    })
}


/* eslint-disable */
export function Edit (formData) {
    return axios({
        url: api.postAuthorityEdit,
        method: 'post',
        data: formData
    })
}



// 删除
/* eslint-disable */
export function Remove (formData) {
    return axios({
        url: api.postAuthorityRemove,
        method: 'post',
        data: formData
    })
}