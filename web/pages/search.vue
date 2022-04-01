<template>
    <div class="container">
        <div class="content-box" :style="{ width: design.width + 'px' }">
            <a-row  :gutter="[{md:12}]">
                <a-col :span="18">
                    <div class="content-box-menu">
                        <ul>
                            <li>
                                <button :class="queryParam.module == MODULE.ALL ? 'active': ''" 
                                 @click="changeType(MODULE.ALL)">综合</button>
                            </li>
                            <li>
                                <button :class="queryParam.module == MODULE.GROUP ? 'active': ''" 
                                 @click="changeType(MODULE.GROUP)">圈子</button>
                            </li>
                            <li>
                                <button :class="queryParam.module == MODULE.USER ? 'active': ''" 
                                 @click="changeType(MODULE.USER)">用户</button>
                            </li>
                            <li>
                                <button :class="queryParam.module == MODULE.ARTICLE ? 'active': ''" 
                                 @click="changeType(MODULE.ARTICLE)">文章</button>
                            </li>
                            <li>
                                <button :class="queryParam.module == MODULE.RESOURCE ? 'active': ''" 
                                 @click="changeType(MODULE.RESOURCE)">资源</button>
                            </li>
                            <li>
                                <button :class="queryParam.module == MODULE.VIDEO ? 'active': ''" 
                                 @click="changeType(MODULE.VIDEO)">视频</button>
                            </li>
                            <li>
                                <button :class="queryParam.module == MODULE.EDU ? 'active': ''" 
                                 @click="changeType(MODULE.EDU)">课程</button>
                            </li>
                            <li>
                                <button :class="queryParam.module == MODULE.AUDIO ? 'active': ''" 
                                 @click="changeType(MODULE.AUDIO)">音频</button>
                            </li>
                        </ul>
                        <div>
                            <a-input-search 
                                placeholder="输入搜索内容" 
                                style="width: 200px" 
                                @search="onSearch" 
                                v-model="queryParam.title"
                            />
                        </div>
                    </div>
                    <div class="hot-search">
                        <span>热搜: </span>
                        <ul>
                            <li  v-for="(item, index) in hotKeyWords" @click="onSearch(item)" :key="index">
                                {{item}}
                            </li>
                        </ul>
                    </div>
                    <div class="search">
                        <!-- 用户 -->
                        <div class="search-list" 
                            v-if="(user.list.length > 0 && user.list != null ) || queryParam.module == MODULE.USER"
                            >
                            <div class="more">
                                <h2>活跃用户</h2>
                                <a-space v-if="queryParam.module == MODULE.ALL">
                                    <span @click="changeType(MODULE.USER)">更多</span>
                                    <a-icon @click="changeType(MODULE.USER)" type="right" />
                                </a-space>
                            </div>
                            <a-row v-if="user.list.length > 0 && user.list != null" :gutter="[{md:12}]">
                                <a-col v-for="(item,index) in user.list" :key="index" :span="6">
                                    <ListSix :info="item"/>
                                </a-col>
                            </a-row>
                            <div class="pagination" v-if="queryParam.module != MODULE.ALL && (user.list.length > 0 && user.list != null)">
                                <a-config-provider :locale="locale">
                                    <a-pagination
                                        @change="changePage"
                                        :pageSize="queryParam.limit"
                                        :total="user.total"
                                        show-quick-jumper
                                    >
                                    </a-pagination>
                                </a-config-provider>
                            </div>
                            <div v-if="user.list.length < 1 || user.list == null" class="group-content-list-emty">
                                <a-config-provider :locale="locale">
                                    <a-empty />
                                </a-config-provider>
                            </div>
                        </div>
                        <!-- 圈子 -->
                        <div class="search-list"
                            v-if="(group.list.length > 0 && group.list != null) || queryParam.module == MODULE.GROUP"
                            >
                            <div class="more">
                                <h2>热门圈子</h2>
                                <a-space v-if="queryParam.module == MODULE.ALL">
                                    <span  @click="changeType(MODULE.GROUP)">更多</span>
                                    <a-icon  @click="changeType(MODULE.GROUP)" type="right" />
                                </a-space>
                            </div>
                            <a-row v-if="group.list.length > 0 && group.list != null" :gutter="[{md:12}]">
                                <a-col :span="8" v-for="(item,index) in group.list" :key="index" >
                                    <ListSeven :info="item"/>
                                </a-col>
                            </a-row>
                            <div class="pagination" v-if="queryParam.module != MODULE.ALL && (group.list.length > 0 && group.list != null)">
                                <a-config-provider :locale="locale">
                                    <a-pagination
                                        @change="changePage"
                                        :pageSize="queryParam.limit"
                                        :total="group.total"
                                        show-quick-jumper
                                    >
                                    </a-pagination>
                                </a-config-provider>
                            </div>
                            <div v-if="group.list.length < 1 || group.list == null" class="group-content-list-emty">
                                <a-config-provider :locale="locale">
                                    <a-empty />
                                </a-config-provider>
                            </div>
                        </div>
                        <!-- 文章 -->
                        <div class="search-list"
                            v-if="(article.list.length > 0 && article.list != null) || queryParam.module == MODULE.ARTICLE"
                            >
                            <div class="more">
                                <h2>热门文章</h2>
                                <a-space v-if="queryParam.module == MODULE.ALL">
                                    <span  @click="changeType(MODULE.ARTICLE)">更多</span>
                                    <a-icon  @click="changeType(MODULE.ARTICLE)" type="right" />
                                </a-space>
                            </div>
                            <ul v-if="article.list.length > 0 && article.list != null">
                                <li v-for="(item,index) in article.list" :key="index">
                                    <list-one :info="item"/>           
                                </li>
                            </ul>
                            <div class="pagination" v-if="queryParam.module != MODULE.ALL && (article.list.length > 0 && article.list != null)" >
                                <a-config-provider :locale="locale">
                                    <a-pagination
                                        @change="changePage"
                                        :pageSize="queryParam.limit"
                                        :total="article.total"
                                        show-quick-jumper
                                    >
                                    </a-pagination>
                                </a-config-provider>
                            </div>
                            <div v-if="article.list.length < 1 || article.list == null" class="group-content-list-emty">
                                <a-config-provider :locale="locale">
                                    <a-empty />
                                </a-config-provider>
                            </div>
                        </div>
                        <!-- 音频 -->
                        <div class="search-list"
                            v-if="(audio.list.length > 0 && audio.list != null) || queryParam.module == MODULE.AUDIO"
                            >
                            <div class="more">
                                <h2>热门音频</h2>
                                <a-space v-if="queryParam.module == MODULE.ALL">
                                    <span  @click="changeType(MODULE.AUDIO)">更多</span>
                                    <a-icon  @click="changeType(MODULE.AUDIO)" type="right" />
                                </a-space>
                            </div>
                            <a-row v-if="audio.list.length > 0 && audio.list != null" :gutter="[{md:12}]">
                                <a-col v-for="(item,index) in audio.list" :key="index" :span="6">
                                   <list-eight :info="item"/>     
                                </a-col>
                            </a-row>
                            <div class="pagination" v-if="queryParam.module != MODULE.ALL && (audio.list.length > 0 && audio.list != null)" >
                                <a-config-provider :locale="locale">
                                    <a-pagination
                                        @change="changePage"
                                        :pageSize="queryParam.limit"
                                        :total="audio.total"
                                        show-quick-jumper
                                    >
                                    </a-pagination>
                                </a-config-provider>
                            </div>
                            <div v-if="audio.list.length < 1 || audio.list == null" class="group-content-list-emty">
                                <a-config-provider :locale="locale">
                                    <a-empty />
                                </a-config-provider>
                            </div>
                        </div>
                        <!-- 视频 -->
                        <div class="search-list"
                            v-if="(video.list.length > 0 && video.list != null) || queryParam.module == MODULE.VIDEO"
                            >
                            <div class="more">
                                <h2>热门视频</h2>
                                <a-space v-if="queryParam.module == MODULE.ALL">
                                    <span @click="changeType(MODULE.VIDEO)">更多</span>
                                    <a-icon @click="changeType(MODULE.VIDEO)" type="right" />
                                </a-space>
                            </div>
                            <a-row v-if="video.list.length > 0 && video.list != null" :gutter="[{md:12}]">
                                <a-col v-for="(item,index) in video.list" :key="index" :span="6">
                                   <list-eight :info="item"/>     
                                </a-col>
                            </a-row>
                            <div class="pagination" v-if="queryParam.module != MODULE.ALL && (video.list.length > 0 && video.list != null)" >
                                <a-config-provider :locale="locale">
                                    <a-pagination
                                        @change="changePage"
                                        :pageSize="queryParam.limit"
                                        :total="video.total"
                                        show-quick-jumper
                                    >
                                    </a-pagination>
                                </a-config-provider>
                            </div>
                            <div v-if="video.list.length < 1 || video.list == null" class="group-content-list-emty">
                                <a-config-provider :locale="locale">
                                    <a-empty />
                                </a-config-provider>
                            </div>
                        </div>
                        <!-- 资源 -->
                        <div class="search-list"
                            v-if="(resource.list.length > 0 && resource.list != null) || queryParam.module == MODULE.RESOURCE"
                            >
                            <div class="more">
                                <h2>热门资源</h2>
                                <a-space v-if="queryParam.module == MODULE.ALL">
                                    <span @click="changeType(MODULE.RESOURCE)">更多</span>
                                    <a-icon @click="changeType(MODULE.RESOURCE)" type="right" />
                                </a-space>
                            </div>
                            <a-row v-if="resource.list.length > 0 && resource.list != null" :gutter="[{md:12}]">
                                <a-col v-for="(item,index) in resource.list" :key="index" :span="6">
                                   <list-eight :info="item"/>     
                                </a-col>
                            </a-row>
                            <div class="pagination" v-if="queryParam.module != MODULE.ALL && (resource.list.length > 0 && resource.list != null)" >
                                <a-config-provider :locale="locale">
                                    <a-pagination
                                        @change="changePage"
                                        :pageSize="queryParam.limit"
                                        :total="resource.total"
                                        show-quick-jumper
                                    >
                                    </a-pagination>
                                </a-config-provider>
                            </div>
                            <div v-if="resource.list.length < 1 || resource.list == null" class="group-content-list-emty">
                                <a-config-provider :locale="locale">
                                    <a-empty />
                                </a-config-provider>
                            </div>
                        </div>
                        <!-- 课程 -->
                        <div class="search-list"
                            v-if="(edu.list.length > 0 && edu.list != null) || queryParam.module == 'course'"
                            >
                            <div class="more">
                                <h2>热门课程</h2>
                                <a-space v-if="queryParam.module == MODULE.ALL">
                                    <span  @click="changeType(MODULE.EDU)">更多</span>
                                    <a-icon @click="changeType(MODULE.EDU)" type="right" />
                                </a-space>
                            </div>
                            <a-row v-if="edu.list.length > 0 && edu.list != null" :gutter="[{md:12}]">
                                <a-col v-for="(item,index) in edu.list" :key="index" :span="6">
                                   <list-eight :info="item"/>     
                                </a-col>
                            </a-row>
                            <div class="pagination" v-if="queryParam.module != MODULE.ALL && (edu.list.length > 0 && edu.list != null)" >
                                <a-config-provider :locale="locale">
                                    <a-pagination
                                        @change="changePage"
                                        :pageSize="queryParam.limit"
                                        :total="edu.total"
                                        show-quick-jumper
                                    >
                                    </a-pagination>
                                </a-config-provider>
                            </div>
                            <div v-if="edu.list.length < 1 || edu.list == null" class="group-content-list-emty">
                                <a-config-provider :locale="locale">
                                    <a-empty />
                                </a-config-provider>
                            </div>
                        </div>
                        
                    </div>
                </a-col>
                <a-col :span="6">
                    sdfasd
                </a-col>
            </a-row>
        </div>
    </div>
</template>


<script>
import { mapState } from "vuex"
// import ChartsSwipe from "@/components/carousel/chartsSwipe"

import ListOne from "@/components/group/ListOne"
import ListTwo from "@/components/group/ListTwo"
import ListThree from "@/components/group/ListThree"
import ListFour from "@/components/group/ListFour"
import ListFive from "@/components/group/ListFive"
import ListSix from "@/components/group/ListSix"
import ListSeven from "@/components/group/ListSeven"
import ListEight from "@/components/group/ListEight"
import ListNine from "@/components/group/ListNine"

import zh_CN from 'ant-design-vue/lib/locale-provider/zh_CN';
import {MODULE} from "@/shared/module"
import api from "@/api/index"
export default {
    components:{
        // ChartsSwipe,
        ListOne,
        ListTwo,
        ListThree,
        ListFour,
        ListFive,
        ListSix,
        ListSeven,
        ListEight,
        ListNine,
    },
    async asyncData({query,$axios,store}){
        const queryParam = {
            page: 1,
            limit: 6,
            title:"",
            module:"",
        }
        const hotList = await $axios.get(api.getSystemHotSearch)

        if (query.keyword != "" && typeof(query.keyword) != "undefined" && query.keyword != 0) {
            queryParam.title = query.keyword
            const res = await $axios.get(api.getSystemSearch,{params:queryParam})
            return {
                base:store.state.base,
                hotKeyWords:hotList.data.list == null ? [] : hotList.data.list,
                queryParam,
                article:{
                    list: res.data.list == null ? [] : res.data.list.filter((item)=>{
                        return item.module == MODULE.ARTICLE
                    }),
                    total: res.data.total == 0 ? 0 :res.data.total,
                },
                audio:{
                    list: res.data.list == null ? [] : res.data.list.filter((item)=>{
                        return item.module == MODULE.AUDIO
                    }),
                    total: res.data.total == 0 ? 0 :res.data.total,
                },
                resource:{
                    list: res.data.list == null ? [] : res.data.list.filter((item)=>{
                        return item.module == MODULE.RESOURCE
                    }),
                    total: res.data.total == 0 ? 0 :res.data.total,
                },
                video:{
                    list: res.data.list == null ? [] : res.data.list.filter((item)=>{
                        return item.module == MODULE.VIDEO
                    }),
                    total: res.data.total == 0 ? 0 :res.data.total,
                },
                edu:{
                    list: res.data.list == null ? [] : res.data.list.filter((item)=>{
                        return item.module == 'course'
                    }),
                    total: res.data.total == 0 ? 0 :res.data.total,
                },
                group:{
                    list: res.data.list == null ? [] : res.data.list.filter((item)=>{
                        return item.module == MODULE.GROUP
                    }),
                    total: res.data.total == 0 ? 0 :res.data.total,
                },
                user:{
                    list: res.data.list == null ? [] : res.data.list.filter((item)=>{
                        return item.module == MODULE.USER
                    }),
                    total: res.data.total == 0 ? 0 :res.data.total,
                }
                // isShow: list != [] ? false :true,
            }
        }
        

        return {
            base:store.state.base,
            queryParam,
            hotKeyWords:hotList.data.list == null ? [] : hotList.data.list,
            article:{
                list: [],
                total: 0,
            },
            audio:{
                list: [],
                total: 0,
            },
            resource:{
                list: [],
                total: 0,
            },
            video:{
                list: [],
                total: 0,
            },
            edu:{
                list: [],
                total: 0,
            },
            group:{
                list: [],
                total: 0,
            },
            user:{
                list: [],
                total: 0,
            }
            // isShow: list != [] ? false :true,
        }
    },
    data(){
        return{
            locale: zh_CN,
            MODULE,
        }
    },
    computed:{
        ...mapState(["design"])
    },
    mounted(){
        // console.log(this.hotKeyWords)
    },
    methods: {
        
        changePage(page,limit){
            this.queryParam.limit = limit
            this.queryParam.page = page
            this.getSearch()
        },
        onSearch(e){
            this.queryParam.title = e
            if (this.queryParam.title == "") {
                this.$message.error(
                    "请输入搜素内容",
                    3
                )
                return
            }
            this.article.list = []
            this.article.total = 0

            this.article.list = []
            this.article.total = 0
            
            this.audio.list = []
            this.audio.total = 0

            this.resource.list = []
            this.resource.total = 0

            this.video.list = []
            this.video.total = 0

            this.edu.list = []
            this.edu.total = 0

            this.group.list = []
            this.group.total = 0

            this.user.list = []
            this.user.total = 0
            if (this.queryParam.module == MODULE.ALL) {
                this.queryParam.page = 1
                this.queryParam.limit = 2
            }
            this.queryParam.page = 1
            this.queryParam.limit = 10
            this.getSearch()
        },
        changeType(t){
            this.article.list = []
            this.article.total = 0

            this.article.list = []
            this.article.total = 0
            
            this.audio.list = []
            this.audio.total = 0

            this.resource.list = []
            this.resource.total = 0

            this.video.list = []
            this.video.total = 0

            this.edu.list = []
            this.edu.total = 0

            this.group.list = []
            this.group.total = 0

            this.user.list = []
            this.user.total = 0

            this.queryParam.module = t
            if (this.queryParam.module == MODULE.ALL) {
                this.queryParam.page = 1
                this.queryParam.limit = 2
            }
            this.queryParam.page = 1
            this.queryParam.limit = 10
            this.getSearch()
        },
        async getSearch(){
            if (this.queryParam.title == "") {
                this.article.list = []
                this.article.total = 0

                this.article.list = []
                this.article.total = 0
                
                this.audio.list = []
                this.audio.total = 0

                this.resource.list = []
                this.resource.total = 0

                this.video.list = []
                this.video.total = 0

                this.edu.list = []
                this.edu.total = 0

                this.group.list = []
                this.group.total = 0

                this.user.list = []
                this.user.total = 0
                return
            }
            const res = await this.$axios.get(api.getSystemSearch,{params: this.queryParam}) 
            this.article.list = res.data.list == null ? [] : res.data.list.filter((item)=>{
                    return item.module == MODULE.ARTICLE
                })
            this.article.total = res.data.total == 0 ? 0 :res.data.total

            this.audio.list = res.data.list == null ? [] : res.data.list.filter((item)=>{
                    return item.module == MODULE.AUDIO
                })
            this.audio.total = res.data.total == 0 ? 0 :res.data.total

            this.resource.list = res.data.list == null ? [] : res.data.list.filter((item)=>{
                    return item.module == MODULE.RESOURCE
                })
            this.resource.total = res.data.total == 0 ? 0 :res.data.total
            
            this.edu.list = res.data.list == null ? [] : res.data.list.filter((item)=>{
                    return item.module == 'course'
                })
            this.edu.total = res.data.total == 0 ? 0 :res.data.total

            this.group.list = res.data.list == null ? [] : res.data.list.filter((item)=>{
                    return item.module == MODULE.GROUP
                })
            this.group.total = res.data.total == 0 ? 0 :res.data.total

            this.user.list = res.data.list == null ? [] : res.data.list.filter((item)=>{
                    return item.module == MODULE.USER
                })
            this.user.total = res.data.total == 0 ? 0 :res.data.total

        },
    },
}
</script>

<style lang="less" scoped>
.container{
    margin: 60px 0;
    margin-top: 70px;
    display: flex;
    justify-content: center;
    min-height: 550px;
    .content-box{
        .content-box-menu{
            background: white;
            padding: 20px;
            display: flex;
            justify-content: space-between;
            align-items: center;
            max-width: 100%;
            ul{
                display: flex;
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
        .hot-search{
            background: white;
            margin: 10px 0;
            padding: 10px;
            display: flex;
            align-items: center;
            span{
                margin-right: 10px;
            }
            ul{
                li{
                    cursor: pointer;
                    border: 1px solid #9E9E9E;
                    padding: 5px;
                    margin-right: 20px;
                    display: inline-block;
                    color: #9E9E9E;
                    border-radius: 2px;
                }
            }
        }
        .search{
            max-width: 100%;
            
            .search-list{
                background: white;
                margin: 10px 0;
                padding: 20px;
                .more{
                    display: flex;
                    justify-content: space-between;
                    align-items: center;
                    margin-bottom: 10px;
                    h2{
                        font-size: 18px;
                        font-weight: 700;
                        color: #0b0b37;
                    }
                    /deep/ .ant-space-item{
                        cursor: pointer;
                    }
                }
                .pagination{
                    margin: 10px 0;
                }
            }
        }
    }
}
</style>

