// eslint-disable-next-line
import {MangerAuthority} from '@/api/manger'
// eslint-disable-next-line
import { BasicLayout, BlankLayout, PageView, RouteView } from '@/layouts'
import { loopMenu } from '@/utils/tree/menu'
// 前端路由表
const constantRouterComponents = {
  // 基础页面 layout 必须引入
  BasicLayout: BasicLayout,
  BlankLayout: BlankLayout,
  RouteView: RouteView,
  PageView: PageView,

  // --------------------- 权限组件
  'AuthorityList': () => import('@/pages/authority/list'),
  'AuthorityCreate': () => import('@/pages/authority/create'),
  'AuthorityEdit': () => import('@/pages/authority/edit'),
  // --------------------- 角色组件
  'RoleList': () => import('@/pages/role/list'),
  'RoleCreate': () => import('@/pages/role/create'),
  'RoleEdit': () => import('@/pages/role/edit'),
  // --------------------- 管理组件
  'MangerList': () => import('@/pages/manger/list'),
  'MangerCreate': () => import('@/pages/manger/create'),
  'MangerEdit': () => import('@/pages/manger/edit'),
  // --------------------- 系统组件
  'SystemSetting': () => import('@/pages/system/setting'),
  'PayOption': () => import('@/pages/option/pay'),
  'OssOption': () => import('@/pages/option/oss'),
  'EmailOption': () => import('@/pages/option/email'),
  'SmsOption': () => import('@/pages/option/sms'),
  // --------------------- 外观
  // 'DesignWeb': () => import('@/pages/design/web'),
  'DesignHome': () => import('@/pages/design/home'),
  // --------------------- 系统工具
  'MediaList': () => import('@/pages/media/list'),
  // --------------------- 运营中心
  'CategoryList': () => import('@/pages/category/list'),
  'CategoryCreate': () => import('@/pages/category/create'),
  'CategoryEdit': () => import('@/pages/category/edit'),
  'TagList': () => import('@/pages/tag/list'),
  // --------------------- 文章管理
  'ArticleList': () => import('@/pages/article/list'),
  'ArticleCreate': () => import('@/pages/article/create'),
  'ArticleEdit': () => import('@/pages/article/edit'),
  // --------------------- 音频管理
  'AudioList': () => import('@/pages/audio/list'),
  'AudioCreate': () => import('@/pages/audio/create'),
  'AudioEdit': () => import('@/pages/audio/edit'),
  // --------------------- 资源管理
  'ResourceList': () => import('@/pages/resource/list'),
  'ResourceCreate': () => import('@/pages/resource/create'),
  'ResourceEdit': () => import('@/pages/resource/edit'),
  // --------------------- 视频管理
  'VideoList': () => import('@/pages/video/list'),
  'VideoCreate': () => import('@/pages/video/create'),
  'VideoEdit': () => import('@/pages/video/edit'),
  // --------------------- 课程管理
  'EduList': () => import('@/pages/edu/list'),
  'EduJoinList': () => import('@/pages/edu/join'),
  'EduCreate': () => import('@/pages/edu/create'),
  'EduEdit': () => import('@/pages/edu/edit'),
  // --------------------- 社区管理
  'TopicList': () => import('@/pages/community/topic/list'),
  'QuestionList': () => import('@/pages/community/question/list'),
  'GroupList': () => import('@/pages/community/group/list'),
  'CommentList': () => import('@/pages/community/comment/list'),
  'AnswerList': () => import('@/pages/community/answer/list'),
  'ReportList': () => import('@/pages/community/report/list'),
  // --------------------- 用户管理
  'UserList': () => import('@/pages/user/list'),
  'UserCreate': () => import('@/pages/user/create'),
  'UserEdit': () => import('@/pages/user/edit'),
  'GradeList': () => import('@/pages/grade/list'),
  'VipList': () => import('@/pages/vip/list'),
  'VerifyList': () => import('@/pages/verify/list'),
  // --------------------- 财务管理
  'OrderList': () => import('@/pages/order/list'),
  'OrderInfo': () => import('@/pages/order/info'),
  'CashList': () => import('@/pages/cash/list'),
  'RechargeList': () => import('@/pages/recharge/list'),
  'CardList': () => import('@/pages/card/list'),

  // --------------------- 财务管理
  'Dashboard': () => import('@/pages/dashboard/Analysis'),

}

// 前端未找到页面路由（固定不用改）
const notFoundRouter = {
  path: '*', redirect: '/404', hidden: true
}

// 根级菜单
const rootRouter = {
  key: '',
  name: '首页',
  path: '/',
  component: 'BasicLayout',
  redirect: '/authority/list',
  meta: {
    title: '首页'
  },
  children: []
}

/**
 * 动态生成菜单
 * @param token
 * @returns {Promise<Router>}
 */
export const generatorDynamicRouter = async () => {
  try {
    const menus = await MangerAuthority()
    menus.data = menus.data == null ? [] : menus.data.filter((item)=>{
      return item.type != 3
    })
    let menuLoot = loopMenu(menus.data || [])
 
    const menuNav = []
    rootRouter.children = menuLoot
    menuNav.push(rootRouter)

    
    const routers = generator(menuNav)
    routers.push(notFoundRouter)
   
    return routers
  } catch (error) {
    throw error
  }
}

/**
 * 格式化树形结构数据 生成 vue-router 层级路由表
 *
 * @param routerMap
 * @param parent
 * @returns {*}
 */
 export const generator = (routerMap, parent) => {
  return routerMap.map((item) => {
    // if (item.type == 1 || item.type == 1) {
      
    // }
    // const { title, show, hideChildren, hiddenHeaderContent, target, icon } = item.meta || {}
    const currentRouter = {
      // 如果路由设置了 path，则作为默认 path，否则 路由地址 动态拼接生成如 /dashboard/workplace
      path: item.path || `${parent && parent.path || ''}/${item.key}`,
      // 路由名称，建议唯一
      name: item.title || item.key || '',
      // 该路由对应页面的 组件 :方案1
      component: constantRouterComponents[item.component || item.key],
      // 该路由对应页面的 组件 :方案2 (动态加载)
      // component: (constantRouterComponents[item.component || item.key]) || (() => import(`@/pages/${item.component}`)),
      key: item.title || item.key || '',
      // meta: 页面标题, 菜单图标, 页面权限(供指令权限用，可去掉)
      meta: {
        title: item.title || item.name,
        icon: item.icon || undefined,
        target: item.target == 2 ? "_blank" : undefined,
        // keepAlive: true,
        permission: item.perms
      }
    }

    //是否设置了隐藏菜单
    if (item.hidden == 1) {
      currentRouter.hidden = true
    }else{
      currentRouter.hidden = false
    }
    // 是否设置了隐藏子菜单
    // if (item.visible != "0") {
    //   currentRouter.hideChildrenInMenu = true
    // }
    // 为了防止出现后端返回结果不规范，处理有可能出现拼接出两个 反斜杠
    if (!currentRouter.path.startsWith('http')) {
      currentRouter.path = currentRouter.path.replace('//', '/')
    }
    // 重定向
    item.redirect && (currentRouter.redirect = item.redirect)
    // 是否有子菜单，并递归处理
    if (item.children && item.children.length > 0) {
      // Recursion
      currentRouter.children = generator(item.children, currentRouter)
    } else{
      delete currentRouter.children
    }
    return currentRouter
  })
}
