import Vue from 'vue'
import Recharge from '@/components/modals/recharge/Recharge'


export default ({store,app: { $axios,$cookies}},) => {
    let instance
    if (!instance) {
        const RechargeConstructor = Vue.extend(Recharge)
        instance = new RechargeConstructor({store,$axios,$cookies})
        instance.$mount(document.createElement('div'))
        document.body.appendChild(instance.$el)
    }

    Vue.prototype.$Recharge = instance.confirm
    // inject('payModal', instance.confirm)
    return instance
}
