import { UserLayout, BasicLayout,PageView } from '@/layouts'

const RouteView = {
  name: 'RouteView',
  render: h => h('router-view')
}

export const asyncRouterMap = [
  {
    path: '/',
    name: 'index',
    component: BasicLayout,
    meta: { title: 'menu.home' },
    redirect: '/system/setting',
    children: [

      //  文章
      {
        path: '/article',
        name: 'menu.article.manage',
        redirect: '/article/list',
        component: PageView,
        meta: { 
          title: 'menu.article.manage',  
          keepAlive: true, 
          icon: "bars", 
          permission: ['/api/v1/admin/article/list'] 
        },
        children: [
          {
            path: '/article/list',
            name: 'menu.article.list',
            component: () => import('@/pages/article/list'),
            meta: { 
              title: 'menu.article.list', 
              keepAlive: true, 
              permission: ['/api/v1/admin/article/list'] 
            }
          },
          {
            path: '/article/create',
            name: 'menu.article.create',
            component: () => import('@/pages/article/create'),
            meta: { 
              title: 'menu.article.create', 
              keepAlive: true, 
              permission: ['/api/v1/admin/article/create'] 
            }
          },
          {
            path: '/article/edit',
            hidden: true,
            name: 'menu.article.edit',
            component: () => import('@/pages/article/edit'),
            meta: { 
              title: 'menu.article.edit', 
              keepAlive: true, 
              permission: ['/api/v1/admin/article/edit'] 
            }
          }
        ]
      },

      //  音频
      {
        path: '/audio',
        name: 'menu.audio.manage',
        redirect: '/audio/list',
        component: PageView,
        meta: { 
          title: 'menu.audio.manage',  
          keepAlive: true, 
          icon: "bars", 
          permission: ['/api/v1/admin/audio/list'] 
        },
        children: [
          {
            path: '/audio/list',
            name: 'menu.audio.list',
            component: () => import('@/pages/audio/list'),
            meta: { 
              title: 'menu.audio.list', 
              keepAlive: true, 
              permission: ['/api/v1/admin/audio/list'] 
            }
          },
          {
            path: '/audio/create',
            name: 'menu.audio.create',
            component: () => import('@/pages/audio/create'),
            meta: { 
              title: 'menu.audio.create', 
              keepAlive: true, 
              permission: ['/api/v1/admin/audio/create'] 
            }
          },
          {
            path: '/audio/edit',
            hidden: true,
            name: 'menu.audio.edit',
            component: () => import('@/pages/audio/edit'),
            meta: { 
              title: 'menu.audio.edit', 
              keepAlive: true, 
              permission: ['/api/v1/admin/audio/edit'] 
            }
          }
        ]
      },

      //  视频
      {
        path: '/video',
        name: 'menu.video.manage',
        redirect: '/video/list',
        component: PageView,
        meta: { 
          title: 'menu.video.manage',  
          keepAlive: true, 
          icon: "bars", 
          permission: ['/api/v1/admin/video/list'] 
        },
        children: [
          {
            path: '/video/list',
            name: 'menu.video.list',
            component: () => import('@/pages/video/list'),
            meta: { 
              title: 'menu.video.list', 
              keepAlive: true, 
              permission: ['/api/v1/admin/video/list'] 
            }
          },
          {
            path: '/video/create',
            name: 'menu.video.create',
            component: () => import('@/pages/video/create'),
            meta: { 
              title: 'menu.video.create', 
              keepAlive: true, 
              permission: ['/api/v1/admin/video/create'] 
            }
          },
          {
            path: '/video/edit',
            hidden: true,
            name: 'menu.audio.edit',
            component: () => import('@/pages/video/edit'),
            meta: { 
              title: 'menu.video.edit', 
              keepAlive: true, 
              permission: ['/api/v1/admin/video/edit'] 
            }
          }
        ]
      },

      //  资源
      {
        path: '/resource',
        name: 'menu.resource.manage',
        redirect: '/resource/list',
        component: PageView,
        meta: { 
          title: 'menu.resource.manage',  
          keepAlive: true, 
          icon: "bars", 
          permission: ['/api/v1/admin/resource/list'] 
        },
        children: [
          {
            path: '/resource/list',
            name: 'menu.resource.list',
            component: () => import('@/pages/resource/list'),
            meta: { 
              title: 'menu.resource.list', 
              keepAlive: true, 
              permission: ['/api/v1/admin/resource/list'] 
            }
          },
          {
            path: '/resource/create',
            name: 'menu.resource.create',
            component: () => import('@/pages/resource/create'),
            meta: { 
              title: 'menu.resource.create', 
              keepAlive: true, 
              permission: ['/api/v1/admin/resource/create'] 
            }
          },
          {
            path: '/resource/edit',
            hidden: true,
            name: 'menu.audio.edit',
            component: () => import('@/pages/resource/edit'),
            meta: { 
              title: 'menu.resource.edit', 
              keepAlive: true, 
              permission: ['/api/v1/admin/resource/edit'] 
            }
          }
        ]
      },

      //  课程
      {
        path: '/edu',
        name: 'menu.edu.manage',
        redirect: '/edu/list',
        component: PageView,
        meta: { 
          title: 'menu.edu.manage', 
          keepAlive: true, 
          icon: "bars", 
          permission: ['/api/v1/admin/edu/list'] 
        },
        children: [
          {
            path: '/edu/list',
            name: 'menu.edu.list',
            component: () => import('@/pages/edu/list'),
            meta: { 
              title: 'menu.edu.list', 
              keepAlive: true,
              permission: ['/api/v1/admin/edu/list'] 
            }
          },
          {
            path: '/edu/create',
            name: 'menu.edu.create',
            component: () => import('@/pages/edu/create'),
            meta: { 
              title: 'menu.edu.create', 
              keepAlive: true, 
              permission: ['/api/v1/admin/edu/create'] 
            }
          },
          {
            path: '/edu/edit',
            hidden: true,
            name: 'menu.edu.edit',
            component: () => import('@/pages/edu/edit'),
            meta: { 
              title: 'menu.edu.edit', 
              keepAlive: true, 
              permission: ['/api/v1/admin/edu/edit'] 
            }
          },
          {
            path: '/edu/join',
            hidden: true,
            name: 'menu.edu.join',
            component: () => import('@/pages/edu/join'),
            meta: { 
              title: 'menu.edu.join', 
              keepAlive: true, 
              permission: ['/api/v1/admin/edu/edit'] 
            }
          }
        ]
      },
      
      //  圈子
      // {
      //   path: '/group',
      //   name: 'menu.group.manage',
      //   redirect: '/group/list',
      //   component: PageView,
      //   meta: { 
      //     title: 'menu.group.manage',  
      //     keepAlive: true, icon: "bars", 
      //     permission: ['/api/v1/admin/group/list'] 
      //   },
      //   children: [
      //     {
      //       path: '/group/list',
      //       name: 'menu.group.list',
      //       component: () => import('@/pages/group/list'),
      //       meta: { 
      //         title: 'menu.group.list', 
      //         keepAlive: true, 
      //         permission: ['/api/v1/admin/group/list'] 
      //       }
      //     },
      //     {
      //       path: '/group/create',
      //       name: 'menu.group.create',
      //       component: () => import('@/pages/group/create'),
      //       meta: { 
      //         title: 'menu.group.create', 
      //         keepAlive: true, 
      //         permission: ['/api/v1/admin/group/create'] 
      //       }
      //     },
      //     {
      //       path: '/group/edit',
      //       hidden: true,
      //       name: 'menu.group.edit',
      //       component: () => import('@/pages/group/edit'),
      //       meta: { 
      //         title: 'menu.group.edit', 
      //         keepAlive: true, 
      //         permission: ['/api/v1/admin/group/edit'] 
      //       }
      //     }
      //   ]
      // },
      
      // //  问题
      // {
      //   path: '/question',
      //   name: 'menu.question.manage',
      //   redirect: '/question/list',
      //   component: PageView,
      //   meta: { 
      //     title: 'menu.question.manage',
      //     keepAlive: true, 
      //     icon: "bars", 
      //     permission: ['/api/v1/admin/question/list'] 
      //   },
      //   children: [
      //     {
      //       path: '/question/list',
      //       name: 'menu.question.list',
      //       component: () => import('@/pages/question/list'),
      //       meta: { title: 'menu.question.list', keepAlive: true, permission: ['/api/v1/admin/question/list'] }
      //     },
      //     {
      //       path: '/question/create',
      //       name: 'menu.question.create',
      //       component: () => import('@/pages/question/create'),
      //       meta: { title: 'menu.question.create', keepAlive: true, permission: ['/api/v1/admin/question/create'] }
      //     },
      //     {
      //       path: '/question/edit',
      //       hidden: true,
      //       name: 'menu.question.edit',
      //       component: () => import('@/pages/question/edit'),
      //       meta: { title: 'menu.question.edit', keepAlive: true, permission: ['/api/v1/admin/question/edit'] }
      //     }
      //   ]
      // },

      //  互动
      // {
      //   path: '/interactive',
      //   name: 'menu.interactive.manage',
      //   redirect: '/topic/list',
      //   component: PageView,
      //   meta: { title: 'menu.interactive.manage',  keepAlive: true, icon: "bars", permission: ['/api/v1/admin/topic/list'] },
      //   children: [
      //     {
      //       path: '/topic/list',
      //       name: 'menu.topic.list',
      //       component: () => import('@/pages/interactive/topic/list'),
      //       meta: { title: 'menu.topic.list', keepAlive: true, permission: ['/api/v1/admin/topic/list'] }
      //     },
      //     {
      //       path: '/comment/list',
      //       name: 'menu.comment.list',
      //       component: () => import('@/pages/interactive/comment/list'),
      //       meta: { title: 'menu.comment.list', keepAlive: true, permission: ['/api/v1/admin/comment/list'] }
      //     },
      //     {
      //       path: '/answer/list',
      //       name: 'menu.answer.list',
      //       component: () => import('@/pages/interactive/answer/list'),
      //       meta: { title: 'menu.answer.list', keepAlive: true, permission: ['/api/v1/admin/answer/list'] }
      //     },
      //     {
      //       path: '/report/list',
      //       name: 'menu.report.list',
      //       component: () => import('@/pages/interactive/report/list'),
      //       meta: { title: 'menu.report.list', keepAlive: true, permission: ['/api/v1/admin/report/list'] }
      //     },
      //     {
      //       path: '/notice/list',
      //       name: 'menu.notice.list',
      //       component: () => import('@/pages/interactive/report/list'),
      //       meta: { title: 'menu.notice.list', keepAlive: true, permission: ['/api/v1/admin/notice/list'] }
      //     }
      //   ]
      // },

      //  聚合管理
      {
        path: '/meta',
        name: 'menu.meta.manage',
        redirect: '/meta/category',
        component: PageView,
        meta: { title: 'menu.meta.manage',  keepAlive: true, icon: "bars", permission: ['/api/v1/admin/category/list'] },
        children: [
          {
            path: '/meta/category',
            name: 'menu.meta.category',
            component: () => import('@/pages/category/list'),
            meta: { title: 'menu.meta.category', keepAlive: true, permission: ['/api/v1/admin/category/list'] }
          },
          {
            path: '/meta/tag',
            name: 'menu.meta.tag',
            component: () => import('@/pages/tag/list'),
            meta: { title: 'menu.meta.tag', keepAlive: true, permission: ['/api/v1/admin/tag/list'] }
          },
          {
            path: '/meta/carousel',
            name: 'menu.meta.carousel',
            component: () => import('@/pages/carousel/list'),
            meta: { title: 'menu.meta.carousel', keepAlive: true, permission: ['/api/v1/admin/carousel/list'] }
          },
          {
            path: '/meta/media',
            name: 'menu.meta.media',
            component: () => import('@/pages/media/list'),
            meta: { title: 'menu.meta.media', keepAlive: true, permission: ['/api/v1/admin/media/list'] }
          },
        ]
      },

      //  用户
      {
        path: '/user',
        name: 'menu.user.manage',
        redirect: '/user/list',
        component: PageView,
        meta: { 
          title: 'menu.user.manage',  
          keepAlive: true, 
          icon: "bars", 
          permission: [
            '/api/v1/admin/user/list',
            '/api/v1/admin/verify/list',
            '/api/v1/admin/role/list',
            '/api/v1/admin/user/create',
            '/api/v1/admin/user/edit'
          ] 
        },
        children: [
          {
            path: '/user/list',
            name: 'menu.user.list',
            component: () => import('@/pages/user/list'),
            meta: { title: 'menu.user.list', keepAlive: true, permission: ['/api/v1/admin/user/list'] }
          },
          {
            path: '/user/create',
            name: 'menu.user.create',
            component: () => import('@/pages/user/create'),
            meta: { title: 'menu.user.create', keepAlive: true, permission: ['/api/v1/admin/user/create'] }
          },
          {
            path: '/user/verify',
            name: 'menu.user.verify',
            component: () => import('@/pages/verify/list'),
            meta: { title: 'menu.user.verify', keepAlive: true, permission: ['/api/v1/admin/verify/list'] }
          },
          {
            path: '/role/list',
            name: 'menu.role.list',
            component: () => import('@/pages/role/list'),
            meta: { title: 'menu.role.list', keepAlive: true, permission: ['/api/v1/admin/role/list'] }
          },
          {
            path: '/grade/list',
            name: 'menu.grade.list',
            component: () => import('@/pages/grade/list'),
            meta: { title: 'menu.grade.list', keepAlive: true, permission: ['/api/v1/admin/grade/list'] }
          },
          {
            path: '/vip/list',
            name: 'menu.vip.list',
            component: () => import('@/pages/vip/list'),
            meta: { title: 'menu.vip.list', keepAlive: true, permission: ['/api/v1/admin/vip/list'] }
          },
          {
            path: '/user/edit',
            hidden: true,
            name: 'menu.user.edit',
            component: () => import('@/pages/user/edit'),
            meta: { title: 'menu.user.edit', keepAlive: true, permission: ['/api/v1/admin/user/edit'] }
          }
        ]
      },

      //  财务
      {
        path: '/finance',
        name: 'menu.finance.manage',
        redirect: '/order/list',
        component: PageView,
        meta: { title: 'menu.finance.manage',  keepAlive: true, icon: "bars", permission: ['/api/v1/admin/order/list'] },
        children: [
          {
            path: '/order/list',
            name: 'menu.order.list',
            component: () => import('@/pages/order/list'),
            meta: { title: 'menu.order.list', keepAlive: true, permission: ['/api/v1/admin/order/list'] }
          },
          {
            path: '/order/info',
            hidden: true,
            name: 'menu.order.info',
            component: () => import('@/pages/order/info'),
            meta: { title: 'menu.order.info', keepAlive: true, permission: ['/api/v1/admin/order/info'] }
          },
          // {
          //   path: '/order/edit',
          //   hidden: true,
          //   name: 'menu.order.edit',
          //   component: () => import('@/pages/order/edit'),
          //   meta: { title: 'menu.order.edit', keepAlive: true, permission: ['/api/v1/admin/order/edit'] }
          // },
          {
            path: '/cash/list',
            name: 'menu.cash.list',
            component: () => import('@/pages/cash/list'),
            meta: { title: 'menu.cash.list', keepAlive: true, permission: ['/api/v1/admin/cash/list'] }
          },
        ]
      },

      //  设置
      {
        path: '/system',
        name: 'menu.setting.manage',
        redirect: '/system/setting',
        component: PageView,
        meta: { title: 'menu.setting.manage', keepAlive: true,  icon: "setting", permission: ['/api/v1/admin/system/setting/info'] },
        children: [
          {
            path: '/system/setting',
            name: 'menu.setting.base',
            component: () => import('@/pages/system/setting'),
            meta: { title: 'menu.setting.base', keepAlive: true, permission: ['/api/v1/admin/system/setting/info'] }
          },
          {
            path: '/pay/option',
            name: 'menu.option.pay',
            component: () => import('@/pages/option/pay'),
            meta: { title: 'menu.option.pay', keepAlive: true, permission: ['/api/v1/admin/system/setting/info'] }
          },
          {
            path: '/oss/option',
            name: 'menu.option.oss',
            component: () => import('@/pages/option/oss'),
            meta: { title: 'menu.option.oss', keepAlive: true, permission: ['/api/v1/admin/system/setting/info'] }
          },
          {
            path: '/email/option',
            name: 'menu.option.email',
            component: () => import('@/pages/option/email'),
            meta: { title: 'menu.option.email', keepAlive: true, permission: ['/api/v1/admin/system/setting/info'] }
          },
        ]
      },

      //  外观
      // {
      //   path: '/design',
      //   name: 'menu.design.manage',
      //   redirect: '/design/web',
      //   component: PageView,
      //   meta: { title: 'menu.design.manage',  keepAlive: true, icon: "setting", permission: ['/api/v1/admin/system/setting/info'] },
      //   children: [
      //     {
      //       path: 'web',
      //       name: 'menu.design.web',
      //       component: () => import('@/pages/design/Web'),
      //       meta: { title: 'menu.design.web', keepAlive: true, permission: ['/api/v1/admin/system/setting/info'] }
      //     },
      //   ]
      // },

      //  权限管理
      {
        path: '/authority',
        name: '权限管理',
        redirect: '/authority/list',
        component: PageView,
        meta: { title: '权限管理',  keepAlive: true, icon: "api", permission: ['/api/v1/admin/system/setting/info'] },
        children: [
          {
            path: '/authority/list',
            name: 'menu.authority.list',
            component: () => import('@/pages/authority/list'),
            meta: { 
              title: 'menu.authority.list',  
              keepAlive: true, 
              permission: ['/api/v1/admin/authority/list'] 
            },
          },
          {
            path: '/authority/create',
            name: '权限创建',
            hidden: true,
            component: () => import('@/pages/authority/create'),
            meta: { 
              title: '权限创建',  
              keepAlive: true, 
              permission: ['/api/v1/admin/authority/list'] 
            },
          },
        ]
      },
      
    ]
  },
  {
    path: '*',
    redirect: '/404',
    hidden: true
  }
]

/**
 * 基础路由
 * @type { *[] }
 */
export const constantRouterMap = [
  {
    path: '/auth',
    component: UserLayout,
    redirect: '/auth/login',
    hidden: true,
    children: [
      {
        path: 'login',
        meta: { title: 'menu.login' },
        name: 'login',
        component: () => import(/* webpackChunkName: "user" */ '@/pages/auth/Login')
      },
      {
        path: 'register',
        name: 'register',
        component: () => import(/* webpackChunkName: "user" */ '@/pages/auth/Register')
      },
      {
        path: 'register-result',
        name: 'registerResult',
        component: () => import(/* webpackChunkName: "user" */ '@/pages/auth/RegisterResult')
      },
      {
        path: 'recover',
        name: 'recover',
        component: undefined
      }
    ]
  },

  {
    path: '/404',
    component: () => import(/* webpackChunkName: "fail" */ '@/views/exception/404')
  }
]