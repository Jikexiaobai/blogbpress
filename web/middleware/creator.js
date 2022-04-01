export default function ({store,redirect}) {
    if (store.state.user.token == null) {
        redirect("/404")
    }
}