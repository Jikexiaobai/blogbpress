import Vue from 'vue'
import JoinEdu from '@/components/modals/joinEdu/JoinEdu'


export default ({store,app: { $axios,$cookies}}) => {
    let instance
    if (!instance) {
        const JoinEduConstructor = Vue.extend(JoinEdu)
        instance = new JoinEduConstructor({store,$axios,$cookies})
        instance.$mount(document.createElement('div'))
        document.body.appendChild(instance.$el)
    }
    Vue.prototype.$JoinEdu = instance.confirm
    return instance
}
