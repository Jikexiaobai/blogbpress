import { axios } from '@/utils/request'
import api from './index'


// 获取列表
/* eslint-disable */
export function List (parameter) {
    return axios({
        url: api.getCarouselList,
        method: 'get',
        params: parameter
    })
}


/* eslint-disable */
export function Create (formData) {
    return axios({
        url: api.postCarouselCreate,
        method: 'post',
        data: formData
    })
}

// 删除
/* eslint-disable */
export function Remove (formData) {
    return axios({
        url: api.postRemoveCarousel,
        method: 'post',
        data: formData
    })
}

