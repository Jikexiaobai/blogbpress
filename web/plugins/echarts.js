// 引入 ECharts 主模块
var echarts = require("echarts/lib/echarts");
import Vue from "vue";
// 引入折线图等组件
require('echarts/lib/chart/line')
require('echarts/lib/chart/bar')
require('echarts/lib/chart/Pie')

// 引入提示框和title组件，图例
require('echarts/lib/component/tooltip')
require('echarts/lib/component/title')
require('echarts/lib/component/legend')
require('echarts/lib/component/legendScroll')//图例翻译滚动
Vue.prototype.$echarts = echarts;