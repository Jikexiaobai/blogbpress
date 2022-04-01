<template>
    <div class="container">
        <div class="warper" :style="{ width: design.width + 'px' }">
            <div class="news-screen">
                <div class="news-screen-cate">
                    <ul>
                        <li>
                            <button :class="queryParam.cateId == 0 ? 'active': ''"  @click="changeCate(0)">不限</button>
                        </li>
                        <li v-for="item in cateList" :key="item.cateId">
                            <button :class="queryParam.cateId == item.cateId ? 'active': ''"  @click="changeCate(item.cateId)">{{item.title}}</button>
                        </li>
                    </ul>
                </div>
            </div>
            <!-- <WidgetOne class="swiper"/> -->
            <a-row class="news-content" :gutter="[{md:12}]">
                <a-col :span="18">
                    <div class="news-list"  v-if="list.length > 0">
                        <ul>
                            <li  v-for="(item,index) in list" :key="index">
                                <list-one :info="item"/>
                            </li>
                            <li v-if="isShow" class="more">
                                已经到底了
                            </li>
                            <li class="more" v-if="!isShow">
                                <a-button type="primary" @click="changeMore" :loading="loading">
                                    加载更多
                                </a-button>
                            </li>
                        </ul>
                    </div>
                    <div class="news-empty" v-if="list.length < 1">
                        <a-config-provider :locale="locale">
                            <a-empty />
                        </a-config-provider>
                    </div>
                </a-col>
                <a-col :span="6">
                    <SidebarUploadArticle/>
                    
                    <SidebarHotArticle />
                    <SidbarAdv />
                    <SidebarComment module="article"/>
                </a-col>
            </a-row>
        </div>
    </div>
</template>



<script>
import api from "@/api/index"
import { mapState } from "vuex"
import {MODULE} from "@/shared/module"
import {MODE} from "@/shared/mode"
import WidgetOne from "@/components/widget/article/widgetOne"
import zhCN from 'ant-design-vue/lib/locale-provider/zh_CN'
import ListOne from '@/components/list/listOne.vue'
import SidebarHotArticle from "@/components/sidebar/sidebarHotArticle"
import SidebarComment from "@/components/sidebar/sidebarComment"
import SidebarUploadArticle from "@/components/sidebar/sidebarUploadArticle"
import SidbarAdv from "@/components/sidebar/sidbarAdv"
export default {
    components:{
        WidgetOne,
        ListOne,
        SidebarHotArticle,
        SidebarComment,
        SidebarUploadArticle,
        SidbarAdv
    },
    head(){
        return this.$seo(`专栏-${this.base.title}`,`专栏`,[{
            hid:"fiber",
            name:"description",
            content:`专栏`
        }])
    },
    computed:{
        ...mapState(["design"]),
        ...mapState("user",["token"]),
    },
    async asyncData({$axios,store}){
   
        const queryParam = {
            page:1,
            limit: 12,
            cateId: 0,
            tagId: 0,
            mode: 0,
            module:"article"
        }
        
       
        const res = await $axios.get(api.getSystemFilter,{params:queryParam})
        
        const cateList = await $axios.get(api.getSystemCate,{params:{module:queryParam.module}})
      
        
        return {
            base:store.state.base,
            queryParam,
            cateList:cateList.data.list != null ?cateList.data.list : [],
            list: res.data.list != null ?res.data.list : [],
            total:res.data.total != 0 ?res.data.total : 0,
            isShow:res.data.list != null ? false : true,
        }
    },
    data(){
        return{
            MODULE,
            MODE,
            locale: zhCN,
            loading: false,
        }
    },
    mounted(){
        window.addEventListener('scroll', this.scrollList)
    },
    destroyed () {
        // 离开页面取消监听
        window.removeEventListener('scroll', this.scrollList, false)
    },
    methods:{
        async scrollList(){
            const scrollTop = document.documentElement.scrollTop || document.body.scrollTop
            const clientHeight = document.documentElement.clientHeight
            const scrollHeight = document.documentElement.scrollHeight
            if (scrollTop + clientHeight >= scrollHeight) {
                if (!this.isShow) {
                    this.queryParam.page += 1
                    const res = await this.$axios.get(api.getSystemFilter,{params:this.queryParam})
                    if (res.code != 1) {
                        this.$router.push(`/404`)
                        this.$message.error(
                            res.message,
                            3
                        )
                        return
                    }
                    if (res.data.list == null) {
                        this.isShow = true
                        return
                    }
                    this.list = [...this.list,...res.data.list]
                    this.total = res.data.total || 0
                }
                return
            }
           
        },
        async getData(){
            const res = await this.$axios.get(api.getSystemFilter,{params:this.queryParam})
            if (res.code != 1) {
                this.$router.push(`/404`)
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
            this.list = res.data.list != null ? res.data.list : []
            this.total = res.data.total != 0 ? res.data.total : 0
        },
        changeCate(e){
            this.queryParam.page = 1
            this.queryParam.cateId = e
            this.getData()

        },
        changeMode(e){
            this.queryParam.mode = e
            this.getData()
        },
        async changeMore(){
            if (!this.isShow) {
                this.queryParam.page += 1
                this.loading = true
                const res = await this.$axios.get(api.getSystemFilter,{params:this.queryParam})
                if (res.code != 1) {
                    this.$router.push(`/404`)
                    this.$message.error(
                        res.message,
                        3
                    )
                    this.loading = false
                    return
                }
                
                if (res.data.list == null) {
                    this.isShow = true
                    this.loading = false
                    return
                }
                this.list = [...this.list,...res.data.list]
                this.total = res.data.total || 0
                this.loading = false
            }
        },
    }
}
</script>
<style lang="less" scoped>
.container{
    margin: 80px 0;
    min-height: 550px;
    display: flex;
    justify-content: center;
    .warper{
        .news-screen{
            top: 120px;
            position: fixed;
            display: flex;
            transition: .2s;
            -webkit-transition: .2s;
            z-index: 9;
            .news-screen-mode{
                ul{
                    display: flex;
                    flex-direction: column;
                    li{
                        font-size: 12px;
                        margin-right: 8px;
                        padding: 5;
                        button{
                            color: #787878;
                            cursor: pointer;
                            background: 0 0;
                            border: 0;
                            color: initial;
                            padding: 5px 10px;
                            border-radius: 2px;
                            -webkit-appearance: none;
                            outline: none;
                            -webkit-tap-highlight-color: rgba(0,0,0,0);
                            font-family: font-regular,'Helvetica Neue',sans-serif;
                            // border: 1px solid #ccc;
                            box-sizing: border-box;
                            user-select: none;
                        }
                        .active{
                            // background-color: #4560c9;
                            color: #4560c9;
                        }
                    }
                }
            }
            .news-screen-cate{
                ul{
                    display: flex;
                    flex-direction: column;
                    li{
                        font-size: 14px;
                        margin-right: 8px;
                        button{
                            cursor: pointer;
                            background: 0 0;
                            border: 0;
                            color: initial;
                            padding: 5px 10px;
                            border-radius: 2px;
                            -webkit-appearance: none;
                            outline: none;
                            -webkit-tap-highlight-color: rgba(0,0,0,0);
                            font-family: font-regular,'Helvetica Neue',sans-serif;
                            // border: 1px solid #ccc;
                            box-sizing: border-box;
                            user-select: none;
                        }
                        .active{
                            background-color: #4560c9;
                            color: #fff;
                        }
                    }
                }
            }
        }
        .swiper{
            padding-left: 120px;
            margin-bottom: 10px;
        }
        .news-content{
            padding-left: 120px;
            .more{
                display: flex;
                justify-content: center;
                align-items: center;
                margin-top: 10px;
            }
        }
    }
}
</style>