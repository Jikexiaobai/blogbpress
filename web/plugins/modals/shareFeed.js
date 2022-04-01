import Vue from 'vue'
import ShareFeed from '@/components/modals/shareFeed/ShareFeed'


export default ({store,app: { $axios,$cookies}}) => {
    let instance
    if (!instance) {
        const ShareFeedConstructor = Vue.extend(ShareFeed)
        instance = new ShareFeedConstructor({store,$axios,$cookies})
        instance.$mount(document.createElement('div'))
        document.body.appendChild(instance.$el)
    }
    Vue.prototype.$ShareFeed = instance.confirm
    return instance
}
