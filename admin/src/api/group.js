import { axios } from '@/utils/request'
import api from './index'


/* eslint-disable */
export function ListMeta (formData) {
    return axios({
        url: api.getGroupListCate,
        method: 'get',
    })
}

// 获取列表
/* eslint-disable */
export function List (parameter) {
    return axios({
        url: api.getGroupList,
        method: 'get',
        params: parameter
    })
}

// 创建
/* eslint-disable */
export function CreateMeta () {
    return axios({
        url: api.getGroupMeta,
        method: 'get',
    })
}




/* eslint-disable */
export function Create (formData) {
    return axios({
        url: api.postGroupCreate,
        method: 'post',
        data: formData
    })
}

// 
/* eslint-disable */
export function EditInfo (parameter) {
    return axios({
        url: api.getGroupEditInfo,
        method: 'get',
        params: parameter
    })
}


/* eslint-disable */
export function Edit (formData) {
    return axios({
        url: api.postGroupEdit,
        method: 'post',
        data: formData
    })
}

// 还原
/* eslint-disable */
export function Recover (formData) {
    return axios({
        url: api.postRecoverGroup,
        method: 'post',
        data: formData
    })
}


// 还原
/* eslint-disable */
export function Reduction (formData) {
    return axios({
        url: api.postReductionGroup,
        method: 'post',
        data: formData
    })
}


// 删除
/* eslint-disable */
export function Remove (formData) {
    return axios({
        url: api.postRemoveGroup,
        method: 'post',
        data: formData
    })
}


/* eslint-disable */
export function Review (formData) {
    return axios({
        url: api.postReviewGroup,
        method: 'post',
        data: formData
    })
}