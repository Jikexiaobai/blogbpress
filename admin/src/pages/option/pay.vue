<template>
  <a-card :bordered="false">
       <a-tabs @change="callback" type="card">
            <a-tab-pane tab="支付宝" key="1">
              <AlyPay/>
            </a-tab-pane>
            <a-tab-pane tab="微信支付" key="2">
              <WechatPay />
            </a-tab-pane>
        </a-tabs>
  </a-card>
</template>
<script>
import AlyPay from "./components/pay/AlyPay"
import WechatPay from "./components/pay/WechatPay"
export default {
  name:"OssOption",
  components:{
    AlyPay,
    WechatPay
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
