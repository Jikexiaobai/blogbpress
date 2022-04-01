import Vue from 'vue'
import Auth from '@/components/modals/auth/Auth'
export default ({store,app: { $axios,$cookies}}) => {
    let instance
    if (!instance) {
        const AuthConstructor = Vue.extend(Auth)
        instance = new AuthConstructor({store,$axios,$cookies})
        instance.$mount(document.createElement('div'))
        document.body.appendChild(instance.$el)
    }
    Vue.prototype.$Auth = instance.confirm
    return instance
}