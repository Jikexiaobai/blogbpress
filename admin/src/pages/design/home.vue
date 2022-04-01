<template>
  <a-card :bordered="false">
    <a-tabs @change="callback" type="card">
        <a-tab-pane tab="网站首页设计" key="1">
            <Web />
        </a-tab-pane>
        <a-tab-pane tab="H5首页设计" key="2">
            <HFive />
        </a-tab-pane>
        <a-tab-pane tab="微信小程序首页设计" key="3">
            <WechatMiniProgram />
        </a-tab-pane>
    </a-tabs>
  </a-card>
</template>
<script>
import Web from "./components/home/web"
import HFive from "./components/home/hfive"
import WechatMiniProgram from "./components/home/wechatMiniProgram"
export default {
    components:{
        Web,
        HFive,
        WechatMiniProgram
    },
    data() {
        return {

        }
    },
    methods: {
        callback(key) {
        // console.log(key);
        },
    },
};
</script>
