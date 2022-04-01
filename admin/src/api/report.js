import { axios } from '@/utils/request'
import api from './index'


// 获取列表
/* eslint-disable */
export function List (parameter) {
    return axios({
        url: api.getReportList,
        method: 'get',
        params: parameter
    })
}




// 删除
/* eslint-disable */
export function Remove (formData) {
    return axios({
        url: api.postRemoveReport,
        method: 'post',
        data: formData
    })
}


/* eslint-disable */
export function Review (formData) {
    return axios({
        url: api.postReviewReport,
        method: 'post',
        data: formData
    })
}
