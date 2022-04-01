import storage from 'store'
import { ACCESS_TOKEN } from '@/store/mutation-types'
import { getAccount } from '@/api/manger'
const user = {
  namespaced: true,
  state: {
    token: null,
    roles: [],
    userInfo: {},
    permissions: []
  },
  mutations: {
    M_SET_TOKEN: (state, token) => {
      state.token = token
    },
    M_SET_ROLES: (state, roles) => {
      state.roles = roles
    },
    M_SET_PERMISSIONS: (state, permissions) => {
      state.permissions = permissions
    },
    M_SET_USERINFO: (state, info) => {
      state.info = info
    }
  },
  actions: {
    A_UPDATE_TOKEN({commit},token){
      commit('M_SET_TOKEN', token)
      storage.set(ACCESS_TOKEN, token, 7 * 24 * 60 * 60 * 1000)
    },
    
    A_UPDATE_USERINFO({commit},userInfo){
        commit('M_SET_USERINFO', userInfo)
    },

    // 获取用户信息
    GetAccount ({ commit }) {
      return new Promise((resolve, reject) => {
        getAccount().then((res)=>{
          const result = res.data
          if (result.roles && result.roles.length > 0) {
            
            commit('M_SET_ROLES', result.roles)
            commit('M_SET_PERMISSIONS', result.authority)
          } else {
            commit('M_SET_ROLES', ["ROLE_DEFAULT"])
            commit('M_SET_PERMISSIONS', result.authority)
          }
          // Vue.ls.set(ACCESS_TOKEN, result.data.token)
          commit('M_SET_USERINFO', result)
          resolve(result.authority)
        }).catch(error => {
          reject(error)
        })
      })
    },
  }
}

export default user
