import Vue from 'vue'
import store from './../../store/'
import Upload from '@/components/modals/upload'


export default function upload () {
    let instance
    if (!instance) {
        const UploadConstructor = Vue.extend(Upload)
        instance = new UploadConstructor({store})
        instance.$mount(document.createElement('div'))
        document.body.appendChild(instance.$el)
    }
    Vue.prototype.$Upload = instance.confirm
    return instance
}