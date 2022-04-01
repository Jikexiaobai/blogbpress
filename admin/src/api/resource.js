import { axios } from '@/utils/request'
import api from './index'


/* eslint-disable */
export function ListMeta () {
    return axios({
        url: api.getResourceListCate,
        method: 'get',
    })
}

// 获取列表
/* eslint-disable */
export function List (parameter) {
    return axios({
        url: api.getResourceList,
        method: 'get',
        params: parameter
    })
}

// 创建
/* eslint-disable */
export function CreateMeta () {
    return axios({
        url: api.getResourceMeta,
        method: 'get',
    })
}




/* eslint-disable */
export function Create (formData) {
    return axios({
        url: api.postResourceCreate,
        method: 'post',
        data: formData
    })
}

// 
/* eslint-disable */
export function EditInfo (parameter) {
    return axios({
        url: api.getResourceEditInfo,
        method: 'get',
        params: parameter
    })
}


/* eslint-disable */
export function Edit (formData) {
    return axios({
        url: api.postResourceEdit,
        method: 'post',
        data: formData
    })
}


// 还原
/* eslint-disable */
export function Recover (formData) {
    return axios({
        url: api.postRecoverResource,
        method: 'post',
        data: formData
    })
}


// 还原
/* eslint-disable */
export function Reduction (formData) {
    return axios({
        url: api.postReductionResource,
        method: 'post',
        data: formData
    })
}


// 删除
/* eslint-disable */
export function Remove (formData) {
    return axios({
        url: api.postRemoveResource,
        method: 'post',
        data: formData
    })
}


/* eslint-disable */
export function Review (formData) {
    return axios({
        url: api.postReviewResource,
        method: 'post',
        data: formData
    })
}