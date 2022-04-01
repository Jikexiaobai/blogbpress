import axios from '@/utils/request'
import api from './index'
// 创建
/* eslint-disable */
export function Save (formData) {
    return axios({
        url: api.postSystemSettingSave,
        method: 'post',
        data: formData
    })
}

// 获取信息
/* eslint-disable */
export function Info (parameter) {
    return axios({
        url: api.getSystemSettingInfo,
        method: 'get',
        params: {
            key:parameter
        }
    })
}

// 获取信息
/* eslint-disable */
export function SystemInfo () {
    return axios({
        url: api.getSystemInfo,
        method: 'get',
    })
}
