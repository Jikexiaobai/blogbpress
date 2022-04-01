import { axios } from '@/utils/request'
import api from './index'


/* eslint-disable */
export function ListMeta (formData) {
    return axios({
        url: api.getEduListCate,
        method: 'get',
    })
}

// 获取列表
/* eslint-disable */
export function List (parameter) {
    return axios({
        url: api.getEduList,
        method: 'get',
        params: parameter
    })
}

// 获取列表
/* eslint-disable */
export function UserJoinList (parameter) {
    return axios({
        url: api.getEduUserJoinList,
        method: 'get',
        params: parameter
    })
}

// 创建
/* eslint-disable */
export function CreateMeta () {
    return axios({
        url: api.getEduMeta,
        method: 'get',
    })
}




/* eslint-disable */
export function Create (formData) {
    return axios({
        url: api.postEduCreate,
        method: 'post',
        data: formData
    })
}

// 
/* eslint-disable */
export function EditInfo (parameter) {
    return axios({
        url: api.getEduEditInfo,
        method: 'get',
        params: parameter
    })
}


/* eslint-disable */
export function Edit (formData) {
    return axios({
        url: api.postEduEdit,
        method: 'post',
        data: formData
    })
}

// 还原
/* eslint-disable */
export function Recover (formData) {
    return axios({
        url: api.postRecoverEdu,
        method: 'post',
        data: formData
    })
}


// 还原
/* eslint-disable */
export function Reduction (formData) {
    return axios({
        url: api.postReductionEdu,
        method: 'post',
        data: formData
    })
}


// 删除
/* eslint-disable */
export function Remove (formData) {
    return axios({
        url: api.postRemoveEdu,
        method: 'post',
        data: formData
    })
}


/* eslint-disable */
export function Review (formData) {
    return axios({
        url: api.postReviewEdu,
        method: 'post',
        data: formData
    })
}