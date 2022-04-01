import Vue from 'vue'
import Cash from '@/components/modals/cash/Cash'
export default ({store,app: { $axios,$cookies}}) => {
    let instance
    if (!instance) {
        const CashConstructor = Vue.extend(Cash)
        instance = new CashConstructor({store,$axios,$cookies})
        instance.$mount(document.createElement('div'))
        document.body.appendChild(instance.$el)
    }
    Vue.prototype.$Cash = instance.confirm
    return instance
}