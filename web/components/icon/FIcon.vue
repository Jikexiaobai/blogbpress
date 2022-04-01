<template>
  <FIcon 
  :style="{fontSize:size+'px',color:color}" 
  :type="type" />
</template>

<script>
import { Icon } from 'ant-design-vue';

const FIcon = Icon.createFromIconfontCN({
  scriptUrl: '//at.alicdn.com/t/font_3028671_f2rm6va5c37.js',
});
export default {
  props:{ 
    type:{
      type: String, //指定传入的类型
      default: "" //这样可以指定默认的值
    },
    size:{
      type: Number, //指定传入的类型
      default: 20 //这样可以指定默认的值
    },
    color:{
      type: String, //指定传入的类型
      default: "rgb(255, 255, 255)" //这样可以指定默认的值
    },
  },
  components:{
    FIcon
  }
}
</script>