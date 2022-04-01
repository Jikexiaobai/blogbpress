import Vue from 'vue'
import Pay from '@/components/modals/pay/Pay'


export default ({store,app: { $axios,$cookies}}) => {
    let instance
    if (!instance) {
        const PayConstructor = Vue.extend(Pay)
        instance = new PayConstructor({store,$axios,$cookies})
        instance.$mount(document.createElement('div'))
        document.body.appendChild(instance.$el)
    }
    Vue.prototype.$Pay = instance.confirm
    return instance
}
