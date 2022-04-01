<template>
  <a-card :bordered="false">
       <a-tabs type="card">
            <a-tab-pane tab="七牛云" key="1">
              <QiNiu/>
            </a-tab-pane>
            <a-tab-pane tab="阿里云" key="2">
              <Aly />
            </a-tab-pane>
        </a-tabs>
  </a-card>
</template>
<script>
import QiNiu from "./components/Oss/QiNiu"
import Aly from "./components/Oss/Aly"
export default {
  name:"OssOption",
  components:{
    QiNiu,
    Aly
  }
};
</script>
