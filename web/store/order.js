export const state = () => ({
    orderNum: null,
})
  
export const mutations = {
    M_UPDATE_ORDER_NUM: (state, payload) => {
        state.orderNum = payload
    },
}

export const actions = {
    A_UPDATE_ORDER_NUM({commit},orderNum){
        commit('M_UPDATE_ORDER_NUM', orderNum)
    },
}