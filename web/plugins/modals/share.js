import Vue from 'vue'
import Share from '@/components/modals/share/Share'


export default ({store,app: { $axios,$cookies}}) => {
    let instance
    if (!instance) {
        const ShareConstructor = Vue.extend(Share)
        instance = new ShareConstructor({store,$axios,$cookies})
        instance.$mount(document.createElement('div'))
        document.body.appendChild(instance.$el)
    }
    Vue.prototype.$Share = instance.confirm
    return instance
}
