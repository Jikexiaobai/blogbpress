<template>
  <a-card :bordered="false">
       <a-tabs @change="callback" type="card">
            <a-tab-pane tab="阿里云" key="1">
              <AliSms/>
            </a-tab-pane>
            <a-tab-pane tab="微信支付" key="2">
              <WechatPay />
            </a-tab-pane>
        </a-tabs>
  </a-card>
</template>
<script>
import AliSms from "./components/sms/AliSms"
import WechatPay from "./components/pay/WechatPay"
export default {
  name:"SmsOption",
  components:{
    AliSms,
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
