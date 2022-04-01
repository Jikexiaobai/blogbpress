export const state = () => ({
    system: 0,
    finance: 0,
    comment: 0,
    like: 0,
    follow: 0,
    answer:0
})
  
export const mutations = {
    M_UPDATE_HAVE_NOTICE: (state, payload) => {
        state.system = payload.system
        state.finance = payload.finance
        state.comment = payload.comment
        state.answer = payload.answer
        state.like = payload.like
        state.follow = payload.follow
        // console.log(state)
    },
}