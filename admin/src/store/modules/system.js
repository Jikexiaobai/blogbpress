import { SystemInfo } from '@/api/system'
const system = {
  namespaced: true,
  state: {
    base: {
      title:"",
      logo:"",
      icon:"",
    },
    file: null,
  },
  mutations: {
    M_SET_BASE: (state, payload) => {
      state.base = payload
    },
    M_SET_FILE: (state, payload) => {
      state.file = payload
    },
  },
  actions: {
    // 获取系统信息
    A_UPDATE_SYSTEM ({ commit }) {
      return new Promise((resolve, reject) => {
        SystemInfo().then((res)=>{
          commit("M_SET_BASE", res.data.base)
          commit("M_SET_FILE", res.data.file)
        }).catch(error => {
          reject(error)
        })
      })
    },
  }
}

export default system
