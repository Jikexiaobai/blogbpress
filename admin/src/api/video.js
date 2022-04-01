import { axios } from '@/utils/request'
import api from './index'


/* eslint-disable */
export function ListMeta () {
    return axios({
        url: api.getVideoListCate,
        method: 'get',
    })
}

// 获取列表
/* eslint-disable */
export function List (parameter) {
    return axios({
        url: api.getVideoList,
        method: 'get',
        params: parameter
    })
}

// 创建
/* eslint-disable */
export function CreateMeta () {
    return axios({
        url: api.getVideoMeta,
        method: 'get',
    })
}




/* eslint-disable */
export function Create (formData) {
    return axios({
        url: api.postVideoCreate,
        method: 'post',
        data: formData
    })
}

// 
/* eslint-disable */
export function EditInfo (parameter) {
    return axios({
        url: api.getVideoEditInfo,
        method: 'get',
        params: parameter
    })
}


/* eslint-disable */
export function Edit (formData) {
    return axios({
        url: api.postVideoEdit,
        method: 'post',
        data: formData
    })
}

// 还原
/* eslint-disable */
export function Recover (formData) {
    return axios({
        url: api.postRecoverVideo,
        method: 'post',
        data: formData
    })
}


// 还原
/* eslint-disable */
export function Reduction (formData) {
    return axios({
        url: api.postReductionVideo,
        method: 'post',
        data: formData
    })
}



// 删除
/* eslint-disable */
export function Remove (formData) {
    return axios({
        url: api.postRemoveVideo,
        method: 'post',
        data: formData
    })
}


/* eslint-disable */
export function Review (formData) {
    return axios({
        url: api.postReviewVideo,
        method: 'post',
        data: formData
    })
}