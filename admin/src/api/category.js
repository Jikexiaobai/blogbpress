import { axios } from '@/utils/request'
import api from './index'


// 删除
/* eslint-disable */
export function ByModule (formData) {
    return axios({
        url: api.getCateByModule,
        method: 'get',
        params: formData
    })
}


// 获取列表
/* eslint-disable */
export function List (parameter) {
    return axios({
        url: api.getCateList,
        method: 'get',
        params: parameter
    })
}



/* eslint-disable */
export function Create (formData) {
    return axios({
        url: api.postCateCreate,
        method: 'post',
        data: formData
    })
}

// 
/* eslint-disable */
export function EditInfo (parameter) {
    return axios({
        url: api.getCateEditInfo,
        method: 'get',
        params: parameter
    })
}


/* eslint-disable */
export function Edit (formData) {
    return axios({
        url: api.postCateEdit,
        method: 'post',
        data: formData
    })
}


// 删除
/* eslint-disable */
export function Remove (formData) {
    return axios({
        url: api.postCateRemove,
        method: 'post',
        data: formData
    })
}