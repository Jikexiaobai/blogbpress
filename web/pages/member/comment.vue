<template>
    <div class="bg-box">
        <a-tabs :default-active-key="contentType" @change="callback">
            <a-tab-pane key="article" tab="文章评论">
                <!-- <CArticle v-if="contentType == 'article'"/> -->
                asdasd
            </a-tab-pane>
            <a-tab-pane key="works" tab="作品评论" force-render>
                <!-- <UWork v-if="contentType == 'works'"/> -->
            </a-tab-pane>
            <a-tab-pane key="resources" tab="资源评论">
                <!-- <UResources v-if="contentType == 'resources'"/> -->
            </a-tab-pane>
            <a-button slot="tabBarExtraContent" @click="goUpload">
                投稿
            </a-button>
        </a-tabs>
    </div>
</template>

<style lang="less" scoped>
.bg-box{
    background-color: #fff;
    padding: 20px;

    width: 1152px;
}
</style>

<script>

export default {
    middleware: ['auth'],
    name:"MemberConter",
    components:{
        // CArticle,
        // UResources,
        // UWork
    },
    data(){
        return{
           contentType:"" 
        }
    },
    created(){
        if (this.$route.query && this.$route.query.type ) {
            this.contentType = this.$route.query.type
        }
    },
    validate({ query }) {
        if (query.type == "article" || query.type == "works" || query.type == "resources") {
             return true // 如果参数有效
        }
        return false // 参数无效，Nuxt.js 停止渲染当前页面并显示错误页面
    },
    methods: {
        callback(key) {
            this.$router.push({ name: "member-comment",
                query:{
                    type:key,
                }
            })
            this.contentType = key
        },
        goUpload(){
            this.$router.push({ name: "member-comment",
                query:{
                    type:"article",
                }
            })
        }
    },
}
</script>