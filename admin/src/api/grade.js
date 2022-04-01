import { axios } from '@/utils/request'
import api from './index'


// 获取列表
/* eslint-disable */
export function List (parameter) {
    return axios({
        url: api.getGradeList,
        method: 'get',
        params: parameter
    })
}



/* eslint-disable */
export function Create (formData) {
    return axios({
        url: api.postGradeCreate,
        method: 'post',
        data: formData
    })
}

// 
/* eslint-disable */
export function EditInfo (parameter) {
    return axios({
        url: api.getGradeEditInfo,
        method: 'get',
        params: parameter
    })
}


/* eslint-disable */
export function Edit (formData) {
    return axios({
        url: api.postGradeEdit,
        method: 'post',
        data: formData
    })
}



// 删除
/* eslint-disable */
export function Remove (formData) {
    return axios({
        url: api.postRemoveGrade,
        method: 'post',
        data: formData
    })
}