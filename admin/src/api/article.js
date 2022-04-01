import { axios } from '@/utils/request'
import api from './index'




// 获取列表
/* eslint-disable */
export function List (parameter) {
    return axios({
        url: api.getArticleList,
        method: 'get',
        params: parameter
    })
}

// 创建
/* eslint-disable */
export function CreateMeta () {
    return axios({
        url: api.getArticleMeta,
        method: 'get',
    })
}




/* eslint-disable */
export function Create (formData) {
    return axios({
        url: api.postArticleCreate,
        method: 'post',
        data: formData
    })
}

// 
/* eslint-disable */
export function EditInfo (parameter) {
    return axios({
        url: api.getArticleEditInfo,
        method: 'get',
        params: parameter
    })
}


/* eslint-disable */
export function Edit (formData) {
    return axios({
        url: api.postArticleEdit,
        method: 'post',
        data: formData
    })
}



// 删除
/* eslint-disable */
export function Remove (formData) {
    return axios({
        url: api.postRemoveArticle,
        method: 'post',
        data: formData
    })
}

// 还原
/* eslint-disable */
export function Recover (formData) {
    return axios({
        url: api.postRecoverArticle,
        method: 'post',
        data: formData
    })
}


// 还原
/* eslint-disable */
export function Reduction (formData) {
    return axios({
        url: api.postReductionArticle,
        method: 'post',
        data: formData
    })
}


/* eslint-disable */
export function Review (formData) {
    return axios({
        url: api.postReviewArticle,
        method: 'post',
        data: formData
    })
}