<template>
  <a-card :bordered="false">
       <a-tabs @change="callback" type="card" default-active-key="1">
            <a-tab-pane tab="基础设置" key="1">
              <base-setting />
            </a-tab-pane>
            <a-tab-pane tab="登录设置" key="2">
              <AuthSetting />
            </a-tab-pane>
            <a-tab-pane tab="用户设置" key="3">
              <UserSetting />
            </a-tab-pane>
            <a-tab-pane tab="文件设置" key="4">
              <file-setting />
            </a-tab-pane>
            <a-tab-pane tab="支付设置" key="5">
              <PaySetting />
            </a-tab-pane>
            <a-tab-pane tab="积分设置" key="6">
              <IntegralSetting />
            </a-tab-pane>
            <!-- <a-tab-pane :tab="$t('system.setting.content')" key="8">
              <ContentSetting />
            </a-tab-pane> -->
        </a-tabs>
  </a-card>
</template>
<script>
import BaseSetting from "./components/BaseSetting"
import FileSetting from "./components/FileSetting"
import AuthSetting from "./components/AuthSetting"
import UserSetting from "./components/UserSetting"
import PaySetting from "./components/PaySetting"
import IntegralSetting from "./components/IntegralSetting"

export default {
  components:{
    BaseSetting,
    FileSetting,
    AuthSetting,
    UserSetting,
    PaySetting,
    IntegralSetting,
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
