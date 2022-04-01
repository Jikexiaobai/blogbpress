<template>
  <div class="container">
        <div class="warper" :style="{ width: this.design.width + 'px' }">
            <div class="account-router">
                <div class="account-router-l">
                    <a-menu
                        mode="inline"
                        :selectedKeys="selectedKeys"
                        :open-keys="openKeys"
                        type="inner"
                        @openChange="onOpenChange"
                        >
                        <!-- <a-menu-item key="/member/dashboard">
                            <nuxt-link :to="{ name: 'member-dashboard' }">
                                <a-icon type="user" />
                                仪表盘
                            </nuxt-link>
                        </a-menu-item> -->

                        <a-menu-item key="/member/article/list">
                                <nuxt-link :to="{ name: 'member-article-list'}">
                                    <a-icon type="menu" />
                                    文章管理
                                </nuxt-link>
                        </a-menu-item>

                        <a-menu-item key="/member/resource/list">
                                <nuxt-link :to="{ name: 'member-resource-list'}">
                                    <a-icon type="menu" />
                                    资源管理
                                </nuxt-link>
                        </a-menu-item>

                        <a-menu-item key="/member/audio/list">
                                <nuxt-link :to="{ name: 'member-audio-list'}">
                                    <a-icon type="menu" />
                                    音频管理
                                </nuxt-link>
                        </a-menu-item>


                        <a-menu-item key="/member/video/list">
                                <nuxt-link :to="{ name: 'member-video-list'}">
                                    <a-icon type="menu" />
                                    视频管理
                                </nuxt-link>
                        </a-menu-item>

                        <a-menu-item key="/member/edu/list">
                                <nuxt-link :to="{ name: 'member-edu-list'}">
                                    <a-icon type="menu" />
                                    课程管理
                                </nuxt-link>
                        </a-menu-item>
<!-- 
                        <a-menu-item key="/member/group/list">
                            <nuxt-link :to="{ name: 'member-group-list'}">
                                <a-icon type="menu" />
                                圈子管理
                            </nuxt-link>
                        </a-menu-item> -->
<!-- 
                        <a-menu-item key="/member/comment">
                            <nuxt-link :to="{ name: 'member-comment' ,query:{type:'article'}}">
                                <a-icon type="dollar" />
                                评论中心
                            </nuxt-link>
                        </a-menu-item>

                        <a-menu-item key="/member/allowance">
                            <nuxt-link :to="{ name: 'member-allowance' }">
                                <a-icon type="rocket" />
                                收益管理
                            </nuxt-link>
                        </a-menu-item> -->
                    </a-menu>
                </div>
                <div class="account-router-r">
                    <Nuxt />
                </div>
            </div>
        </div>
  </div>
</template>

<style lang="less" scoped>
.container{
    margin: 20px 0;
    display: flex;
    justify-content: center;
    .warper{
        .account-router{
            display: flex;
            .account-router-l{
                background: white;
                width: 300px;
                // min-height: 100vh;
                margin-right: 20px;
            }
            .account-router-r{
                width: 100%;
                // min-height: 100vh;
            }
        }
    }
}

// 响应式处理
@media only screen and (max-width: 768px) {
  .container{
    .account-content{
      .warper {
        width: 100% !important;
      }
    }
  }
}
</style>


<script>
import { mapState } from "vuex"
export default {
    layout:"member",
    asyncData({ redirect,route }) {
        if (route.path == "/member") {
            redirect("/member/dashboard")
        }
    },
    computed:{
        ...mapState(["design"]),
        ...mapState("user",["token"]),
    },
    data () {
        return {
        // horizontal  inline
        //   mode: 'inline',
        openKeys: [],
        selectedKeys: [],
        }
    },
    mounted () {
        this.updateMenu()
    },
    methods: {
        onOpenChange (openKeys) {
        this.openKeys = openKeys
        },
        updateMenu () {
            const routes = this.$route.matched.concat()
           
            this.selectedKeys = [ routes.pop().path ]
             console.log(this.selectedKeys)
        }
    },
    watch: {
        '$route' (val) {
            this.updateMenu()
        }
    }
}
</script>