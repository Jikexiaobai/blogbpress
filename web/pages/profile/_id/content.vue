<template>
    <div class="post-box">
        <a-row class="account-router" :gutter="[{md:12}]">
            <a-col :span="24" :md="4">
                <!-- <Author /> -->
                <a-menu
                    mode="inline"
                    :selectedKeys="selectedKeys"
                    type="inner"
                    @openChange="onOpenChange"
                    >
                    <a-menu-item :key="`/profile/:id?/content/article`">
                        
                        <nuxt-link :to="{ path: `/profile/${id}/content/article` }">
                            <a-icon type="user" />
                            文章
                        </nuxt-link>
                    </a-menu-item>

                    <a-menu-item :key="`/profile/:id?/content/video`">
                        
                        <nuxt-link :to="{ path: `/profile/${id}/content/video` }">
                            <a-icon type="user" />
                            视频
                        </nuxt-link>
                    </a-menu-item>

                    <a-menu-item :key="`/profile/:id?/content/audio`">
                        <nuxt-link :to="{ path: `/profile/${id}/content/audio` }">
                            <a-icon type="safety-certificate" />
                            音频
                        </nuxt-link>
                    </a-menu-item>

                    <a-menu-item :key="`/profile/:id?/content/resource`">
                        <nuxt-link :to="{ path: `/profile/${id}/content/resource` }">
                            <a-icon type="dollar" />
                            资源
                        </nuxt-link>
                    </a-menu-item>

                    <a-menu-item :key="`/profile/:id?/content/course`">
                        <nuxt-link :to="{ path: `/profile/${id}/content/course` }">
                            <a-icon type="rocket" />
                            课程
                        </nuxt-link>
                    </a-menu-item>
                </a-menu>
            </a-col>
            <a-col :span="24" :md="20">
                <Nuxt />
            </a-col>
        </a-row>
    </div>
</template>

<style lang="less" scoped>
.post-box{
    margin-top: 10px;
    background-color: white;
   
}
</style>


<script>
import { mapState } from "vuex"
export default {
    asyncData({ params,redirect,route }) {
        if (route.path == `/profile/${params.id}/content`) {
            redirect(`/profile/${params.id}/content/article`)
        }
    },
    data () {
        return {
        // horizontal  inline
        //   mode: 'inline',
            openKeys: [],
            selectedKeys: [],
            id:0
        }
    },
    computed:{
        ...mapState(["allWidth"])
    },
    mounted () {
        this.id = this.$route.params.id
        this.updateMenu()
    },
    methods: {
        onOpenChange (openKeys) {
            this.openKeys = openKeys
        },
        updateMenu () {
            const routes = this.$route.matched.concat()
            
            this.selectedKeys = [ routes.pop().path ]
        }
    },
    watch: {
        '$route' (val) {
            this.updateMenu()
        }
    }
}
</script>