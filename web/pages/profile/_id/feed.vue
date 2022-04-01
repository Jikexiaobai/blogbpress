<template>
    <div class="feed">
        <div class="feed-box">
            <ul v-if="list.length > 0">
                <li v-for="(item,index) in list" :key="index">
                    <FeedList :info="item"/>
                </li>
                <li v-if="isShow" class="feed-bottom">
                    已经到底了
                </li>
            </ul>
            <div v-if="list.length < 1"  class="empty">
                <a-config-provider :locale="locale">
                    <a-empty />
                </a-config-provider>
            </div>
        </div>
    </div>
</template>
<style lang="less" scoped>
.feed{
    .feed-box{
        background: white;
        padding: 20px;
        .empty{
            display: flex;
            justify-content: center;
            align-items: center;
            height: 300px;
        }
        .feed-bottom{
            display: flex;
            align-items: center;
            justify-content: center;
            padding-top: 20px;
        }
    }
}
</style>


<script>
import FeedList from "@/components/list/feedItem"
import zhCN from 'ant-design-vue/lib/locale-provider/zh_CN'
import api from "@/api/index"

export default {
    components:{
        FeedList,
    },
    data(){
        return{
            locale: zhCN,
            queryParam : {
                page: 1,
                limit: 8,
                module: "topic",
                userId:0,
                mode: 2,
                type: 0,
            },
            list:[],
            isShow:false
        }
    },
    mounted(){
        this.queryParam.userId = parseInt(this.$route.params.id)
        this.getTopicList()
        window.addEventListener('scroll', this.scrollList)
    },
    destroyed () {
        // 离开页面取消监听
        window.removeEventListener('scroll', this.scrollList, false)
    },
    methods:{
        scrollList(){
            //变量scrollTop是滚动条滚动时，距离顶部的距离
            var scrollTop = document.documentElement.scrollTop||document.body.scrollTop;
            //变量windowHeight是可视区的高度
            var windowHeight = document.documentElement.clientHeight || document.body.clientHeight;
            //变量scrollHeight是滚动条的总高度
            var scrollHeight = document.documentElement.scrollHeight||document.body.scrollHeight;

            //滚动条到底部的条件
            if (scrollTop+windowHeight > scrollHeight-50 && !this.isShow) {
                this.queryParam.page += 1
                this.getTopicList()
                return
            }
        },
        async getTopicList(){
            const res = await this.$axios.get(api.getUserPosts,{params: this.queryParam}) 
            if (res.code != 1) {
                this.$router.push(`/404`)
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
            res.data.list = res.data.list == null ? [] : res.data.list.map((item)=>{
                if (item.type == 1 && item.files != "") {
                    item.files = JSON.parse(item.files)
                }
                return item
            })
            this.isShow = res.data.list.length > 0 ? false : true
            this.list = [...this.list,...res.data.list]
            this.total = res.data.total
        },
    }
}
</script>