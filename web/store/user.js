export const state = () => ({
    token: null,
    userInfo:{
        avatar:"",
        sign:null,
    },
    
})
  
export const mutations = {
    M_UPDATE_TOKEN: (state, payload) => {
        state.token = payload
    },
    M_UPDATE_USER: (state, payload) => {
        state.userInfo = payload
    },
    M_UPDATE_NICKNAME: (state, payload) => {
        state.userInfo.nickName = payload
    },
    M_UPDATE_AVATAR: (state, payload) => {
        state.userInfo.avatar = payload
    },
    M_UPDATE_SIGN: (state, payload) => {
        state.userInfo.sign = payload
    },
}

export const actions = {
    A_UPDATE_TOKEN({commit},token){
        commit('M_UPDATE_TOKEN', token)
    },
    A_UPDATE_USER({commit},userInfo){
        commit('M_UPDATE_USER', userInfo)
    },
    A_UPDATE_SIGN({commit},sign){
        commit('M_UPDATE_SIGN', sign)
    },
}