import { axios } from '@/utils/request'
import api from './index'


/* eslint-disable */
export function ListMeta (formData) {
    return axios({
        url: api.getAudioListCate,
        method: 'get',
    })
}

// 获取列表
/* eslint-disable */
export function List (parameter) {
    return axios({
        url: api.getAudioList,
        method: 'get',
        params: parameter
    })
}

// 创建
/* eslint-disable */
export function CreateMeta () {
    return axios({
        url: api.getAudioMeta,
        method: 'get',
    })
}




/* eslint-disable */
export function Create (formData) {
    return axios({
        url: api.postAudioCreate,
        method: 'post',
        data: formData
    })
}

// 
/* eslint-disable */
export function EditInfo (parameter) {
    return axios({
        url: api.getAudioEditInfo,
        method: 'get',
        params: parameter
    })
}


/* eslint-disable */
export function Edit (formData) {
    return axios({
        url: api.postAudioEdit,
        method: 'post',
        data: formData
    })
}



// 删除
/* eslint-disable */
export function Remove (formData) {
    return axios({
        url: api.postRemoveAudio,
        method: 'post',
        data: formData
    })
}

// 还原
/* eslint-disable */
export function Recover (formData) {
    return axios({
        url: api.postRecoverAudio,
        method: 'post',
        data: formData
    })
}


// 还原
/* eslint-disable */
export function Reduction (formData) {
    return axios({
        url: api.postReductionAudio,
        method: 'post',
        data: formData
    })
}


/* eslint-disable */
export function Review (formData) {
    return axios({
        url: api.postReviewAudio,
        method: 'post',
        data: formData
    })
}