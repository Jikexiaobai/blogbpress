<template>
    <div class="content-box">
        <div class="warper" :style="{ width: design.width + 'px' }">
            <div class="left">
                <div class="meun-list">
                    <ul>
                        <li>
                            <button :class="queryParam.mode == MODE.NEW ? 'active': ''" @click="changeMeun(MODE.NEW)">最新问题</button>
                        </li>
                        <li>
                            <button :class="queryParam.mode == MODE.HOT ? 'active': ''" @click="changeMeun(MODE.HOT)">热门问题</button>
                        </li>
                        <li>
                            <button :class="queryParam.mode == MODE.FAVORITE ? 'active': ''" @click="changeMeun(MODE.FAVORITE)">收藏最多</button>
                        </li>
                        <li>
                            <button :class="queryParam.mode == MODE.VIEW ? 'active': ''" @click="changeMeun(MODE.VIEW)">浏览最多</button>
                        </li>
                    </ul>
                    <a-button @click="goCreate" type="primary">
                        发布问题
                    </a-button>
                </div>
                <div  class="ask-list-box">
                    <ul v-if="list.length > 0 && !loading">
                        <li v-for="(item,index) in list" :key="index" >
                            <div class="ask-list-item">
                                <div class="ask-rank">
                                    <div class="answere-count">
                                    {{item.answers | resetNum}} <small>回答</small>  
                                    </div>
                                    <div class="ask-view">
                                        {{item.views | resetNum}} <small>查看</small>  
                                    </div>
                                </div>
                                <div class="ask-info">
                                    <div class="ask-author-date">
                                        <nuxt-link :to="{path:'/profile/' + item.userInfo.id }" class="item-link">   
                                            <h2 class="user-name">{{item.userInfo.nickName}}</h2>
                                        </nuxt-link>
                                        <span class="date">{{item.createTime | resetData}}</span>
                                    </div>
                                    <nuxt-link :to="{path:'/question/' + item.id }" class="item-link">   
                                        <h2 class="ask-title">{{item.title}}</h2>
                                    </nuxt-link>
                                </div>
                            </div>
                        </li>
                    </ul>
                    <ul v-if="loading">
                        <li class="loading">
                            <a-skeleton :paragraph="{ rows: 2 }" />
                        </li>
                        <li class="loading">
                            <a-skeleton :paragraph="{ rows: 2}" />
                        </li>
                        <li class="loading">
                            <a-skeleton :paragraph="{ rows: 2}" />
                        </li>
                    </ul>
                    <div v-if="list.length < 1 && !loading" class="list-emty">
                        <a-config-provider :locale="locale">
                            <a-empty />
                        </a-config-provider>
                    </div>
                </div>
                <div v-if="list.length > 0" class="ask-page">
                    <a-config-provider :locale="locale">
                        <a-pagination
                            @change="changePage"
                            :pageSize="queryParam.limit"
                            :total="total"
                            show-quick-jumper
                        >
                        </a-pagination>
                    </a-config-provider>
                </div>
            </div>
            <div class="right">
                
                <SidebarAnswerUserList/>
                <SidbarAdv />
                <SidebarQuestionList/>
            </div>
        </div>
    </div>
</template>


<script>
import SidebarAnswerUserList from "@/components/sidebar/sidebarAnswerUserList"
import SidebarQuestionList from "@/components/sidebar/sidebarQuestionList"
import SidbarAdv from "@/components/sidebar/sidbarAdv"
import { mapState } from "vuex"
import api from "@/api/index"
import zhCN from 'ant-design-vue/lib/locale-provider/zh_CN'
import {MODE} from "@/shared/mode"
export default {
    components:{
        SidebarAnswerUserList,
        SidbarAdv,
        SidebarQuestionList,
    },
    computed:{
        ...mapState(["design","base"]),
    },
    head(){
        return this.$seo(`问答中心-${this.base.title}`,`${this.base.childTitle}`,[{
            hid:"fiber-desc",
            name:"description",
            content:`${this.base.description}`
        }])
    },
    data () {
        return {
            MODE,
            locale: zhCN,
            queryParam:{
                page: 1,
                limit: 20,
                mode: MODE.NEW,
            },
            list: [],
            total: 0,
            loading:false,
        }
    },
    mounted(){
        this.getData()
    },
    methods: {
        async getData(){
            this.loading = true
            const res = await this.$axios.get(api.getQuestionList,{params: this.queryParam})
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
            this.total = res.data.total
            this.list = res.data.list == null ? [] : res.data.list
            this.loading = false
        },
        goCreate(){
            this.$router.push({ path: `/question/create`})
        },
        // 修改菜单
        changeMeun(e){
            this.queryParam.mode = e
            this.getData()
        },
        changePage(page,limit){
            this.queryParam.limit = limit
            this.queryParam.page = page
            this.getData()
        },
    },
}
</script>

<style lang="less" scoped>
.content-box{
    margin: 80px 0;
    display: flex;
    justify-content: center;
    // min-height: 550px;
    .warper{
        display: flex;
        .left{
            flex: 1;
            
            .meun-list{
                background: white;
                display: flex;
                align-items: center;
                justify-content: space-between;
                margin: 0;
                // position: relative;
                padding: 20px;
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
                            padding: 10px 10px;
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
            .ask-list-box{
                background: white;
                padding: 0px 20px 0 20px;
                .ask-list-item{
                    padding: .75rem 0;
                    border-bottom: 1px solid rgba(0,0,0,.0625);
                    display: flex;
                    justify-content: space-between;
                    align-items: flex-start;
                    .ask-rank{
                        display: flex;
                        .answere-count{
                            margin-right: 10px;
                            text-align: center;
                            line-height: 1.2;
                            display: inline-block;
                            padding-top: 3px;
                            width: 40px;
                            height: 40px;
                            -moz-border-radius: 1px;
                            -webkit-border-radius: 1px;
                            border-radius: 3px;
                            background: #0084ff;
                            border: none;
                            color: #fff;
                            small{
                                display: block;
                                font-size: 12px;
                            }
                        }
                        .ask-view{
                            margin-right: 10px;
                            text-align: center;
                            line-height: 1.2;
                            display: inline-block;
                            padding-top: 3px;
                            width: 40px;
                            height: 40px;
                            -moz-border-radius: 1px;
                            -webkit-border-radius: 1px;
                            border-radius: 3px;
                            background-color: #f3f3f3;
                            border-color: #f3f3f3;
                            color: #7b7b7b;
                            small{
                                display: block;
                                font-size: 12px;
                            }
                        }
                    }
                    .ask-info{
                        flex: 1;
                        display: flex;
                        height: 40px;
                        flex-direction: column;
                        justify-content: space-between;
                        .ask-author-date{
                            display: flex;
                            align-items: center;
                            
                            .user-name{
                                font-weight: 600;
                                font-size: 14px;
                                color: #7b7b7b;
                            }
                            .date{
                                margin-left: 10px;
                                font-size: 12px;
                                color: #7b7b7b;
                            }
                        }
                        .ask-title{
                            font-size: 14px;
                        }
                    }
                }
                .list-emty{
                    background: white;
                    padding: 20px;
                }
            }
            .ask-page{
                background: white;
                padding: 20px 20px 20px 20px;
                display: flex;
                align-items: center;
                justify-content: flex-end;
            }
        }
        .right{
            margin-left: 20px;
            width: 280px;
        }
    }

}
</style>
