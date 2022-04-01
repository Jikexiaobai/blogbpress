import Vue from 'vue'
import Report from '@/components/modals/report/Report'


export default ({store,app: { $axios,$cookies}}) => {
    let instance
    if (!instance) {
        const ReportConstructor = Vue.extend(Report)
        instance = new ReportConstructor({store,$axios,$cookies})
        instance.$mount(document.createElement('div'))
        document.body.appendChild(instance.$el)
    }
    Vue.prototype.$Report = instance.confirm
    return instance
}
