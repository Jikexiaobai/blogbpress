import Vue from 'vue'
import store from './../../store/'
import BigFileUpload from '@/components/modals/bigFileUpload'


export default function bigFileUpload () {
    let instance
    if (!instance) {
        const BigFileUploadConstructor = Vue.extend(BigFileUpload)
        instance = new BigFileUploadConstructor({store})
        instance.$mount(document.createElement('div'))
        document.body.appendChild(instance.$el)
    }
    Vue.prototype.$BigFileUpload = instance.confirm
    return instance
}