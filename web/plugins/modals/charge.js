import Vue from 'vue'
import Charge from '@/components/modals/charge/Charge'


export default ({store,app: { $axios,$cookies}},) => {
    let instance
    if (!instance) {
        const ChargeConstructor = Vue.extend(Charge)
        instance = new ChargeConstructor({store,$axios,$cookies})
        instance.$mount(document.createElement('div'))
        document.body.appendChild(instance.$el)
    }

    Vue.prototype.$Charge = instance.confirm
    return instance
}
