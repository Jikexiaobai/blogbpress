<template>
    <div class="question">
        <div class="question-box">
            <ul v-if="list.length > 0">
                <li v-for="(item,index) in list" :key="index" >
                    <div class="ask-list-item">
                        <!-- <div class="ask-rank">
                            <div class="answere-count">
                                {{item.answers | resetNum}} <small>回答</small>  
                            </div>
                            <div class="ask-view">
                                {{item.views | resetNum}} <small>查看</small>  
                            </div>
                        </div> -->
                        <div class="ask-info">
                            <div class="ask-author-date">
                                <nuxt-link :to="{path:'/profile/' + item.userInfo.id }" class="item-link">   
                                    <h2 class="user-name">{{item.userInfo.nickName}}</h2>
                                </nuxt-link>
                                <span class="date">{{item.createTime | resetData}}</span>
                            </div>
                            <nuxt-link :to="{path:'/question/' + item.questionId }" class="item-link">   
                                <h2 class="ask-title">{{item.title}}</h2>
                            </nuxt-link>
                        </div>
                    </div>
                    <a-divider/>
                </li>
            </ul>
            <div  v-if="list.length > 0" class="content-pagination" >
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
            <div v-if="list.length < 1"  class="empty">
                <a-config-provider :locale="locale">
                    <a-empty />
                </a-config-provider>
            </div>
        </div>
    </div>
</template>
<style lang="less" scoped>
.question{
    .question-box{
        .empty{
            display: flex;
            justify-content: center;
            align-items: center;
            padding-bottom: 20px;
            height: 300px;
        }
        .content-pagination{
            padding: 0 0 20px 0;
        }
        background: white;
        padding: 20px 20px 0 20px;
        .ask-list-item{
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
                .ask-author-date{
                    display: flex;
                    align-items: center;
                    
                    .user-name{
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
}
</style>


<script>
import zhCN from 'ant-design-vue/lib/locale-provider/zh_CN'
import api from "@/api/index"
export default {
    data(){
        return{
            locale: zhCN,
            queryParam : {
                page: 1,
                limit: 8,
                module: "question",
                userId:0,
                mode: 1,
            },
            list:[],
            total:0,
        }
    },
    mounted () {
        this.queryParam.userId = parseInt(this.$route.params.id)
        this.getData()
    },
    methods:{
        async getData(){
            const res = await this.$axios.get(api.getUserPosts,{params:this.queryParam})
            if (res.code != 1) {
                this.$router.push(`/404`)
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
            this.list = res.data.list || []
            this.total = res.data.total || 0
        },
        changePage(page,limit){
            this.queryParam.limit = limit
            this.queryParam.page = page
            this.getData()
        },
    }
}
</script>