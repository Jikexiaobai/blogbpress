import antd from 'ant-design-vue/es/locale-provider/zh_CN'
import momentCN from 'moment/locale/zh-cn'
import global from './zh-CN/global'

import menu from './zh-CN/menu'
import setting from './zh-CN/setting'
import user from './zh-CN/user'
import dashboard from './zh-CN/dashboard'
import form from './zh-CN/form'
import result from './zh-CN/result'
import account from './zh-CN/account'
import authority from './zh-CN/authority'
import modal from './zh-CN/modal'
import role from './zh-CN/role'
import grade from './zh-CN/grade'
import vip from './zh-CN/vip'
import article from './zh-CN/article'
import audio from './zh-CN/audio'
import video from './zh-CN/video'
import resource from './zh-CN/resource'
import question from './zh-CN/question'
import topic from './zh-CN/topic'
import comment from './zh-CN/comment'
import answer from './zh-CN/answer'
import group from './zh-CN/group'
import edu from './zh-CN/edu'
import report from './zh-CN/report'
import order from './zh-CN/order'
import cash from './zh-CN/cash'
import verify from './zh-CN/verify'
import option from './zh-CN/option'


import category from './zh-CN/category'
import tag from './zh-CN/tag'
import carousel from './zh-CN/carousel'
import media from './zh-CN/media'

const components = {
  antLocale: antd,
  momentName: 'zh-cn',
  momentLocale: momentCN
}

export default {
  message: '-',

  'layouts.usermenu.dialog.title': '信息',
  'layouts.usermenu.dialog.content': '您确定要注销吗？',
  'layouts.userLayout.title': 'Ant Design 是西湖区最具影响力的 Web 设计规范',
  ...components,
  ...global,
  ...menu,
  ...setting,
  ...user,
  ...dashboard,
  ...form,
  ...result,
  ...account,
  ...authority,
  ...modal,
  ...role,
  ...grade,
  ...vip,
  ...article,
  ...audio,
  ...video,
  ...resource,
  ...question,
  ...topic,
  ...comment,
  ...answer,
  ...group,
  ...edu,
  ...category,
  ...tag,
  ...report,
  ...carousel,
  ...media,
  ...order,
  ...cash,
  ...verify,
  ...option
}
