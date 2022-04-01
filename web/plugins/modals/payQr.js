import Vue from 'vue'
import QrCode from '@/components/modals/payQr/QrCode'


export default ({store,app: { $axios,$cookies}}) => {
    let instance
    if (!instance) {
        const QrCodeConstructor = Vue.extend(QrCode)
        instance = new QrCodeConstructor({store,$axios,$cookies})
        instance.$mount(document.createElement('div'))
        document.body.appendChild(instance.$el)
    }
    Vue.prototype.$PayQr = instance.confirm
    return instance
}
