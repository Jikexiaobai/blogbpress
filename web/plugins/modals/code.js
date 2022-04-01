import Vue from 'vue'
import Code from '@/components/modals/code/Code'
export default ({store,app: { $axios,$cookies}}) => {
    let instance
    if (!instance) {
        const CodeConstructor = Vue.extend(Code)
        instance = new CodeConstructor({store,$axios,$cookies})
        instance.$mount(document.createElement('div'))
        document.body.appendChild(instance.$el)
    }
    Vue.prototype.$Code = instance.confirm
    return instance
}