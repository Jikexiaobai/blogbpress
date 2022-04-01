import Vue from 'vue'
import ImgPreview from '@/components/modals/image/ImgPreview'


export default ({store,app: { $axios,$cookies}}) => {
    let instance
    if (!instance) {
        const ImgPreviewConstructor = Vue.extend(ImgPreview)
        instance = new ImgPreviewConstructor({store,$axios,$cookies})
        instance.$mount(document.createElement('div'))
        document.body.appendChild(instance.$el)
    }
    Vue.prototype.$ImgPreview = instance.confirm
    return instance
}
