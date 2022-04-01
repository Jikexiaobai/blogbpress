<template>
    <div class="content-box">
        <div class="ask-detail">
            <div class="ask-info-box" :style="{ width: design.width + 'px' }">
                <a-row :gutter="[{md:14}]">
                    <a-col class="ask-l" :span="18">
                        <div>
                            <a-tag color="#f50">
                                {{info.groupInfo.title}}
                            </a-tag>
                        </div>
                        <h2 class="ask-title">
                            {{info.title}}
                        </h2>
                        <p class="date">
                            发布于<span class="ask-date">{{info.createTime | resetData}}</span>
                        </p>
                        <div class="ask-desc" v-html="info.content"></div>
                        <div class="ask-btn">
                            <a-space >
                                <a-button @click="postLike(info.id)" type="link">
                                    <a-icon :theme="info.isLike ? 'filled' : 'outlined'" type="like" />
                                    {{info.isLike ? '已赞' : '赞'}} {{info.likes == 0 ? "" : info.likes | resetNum}}
                                </a-button>
                                <a-button @click="postFavorite(info.id)" type="link">
                                    <a-icon :theme="info.isFavorite ? 'filled' : 'outlined'" type="star" />
                                    {{info.isFavorite ? '已收藏' : '收藏'}} {{info.favorites == 0 ? "" : info.favorites | resetNum}}
                                </a-button>
                                <a-button type="link" icon="share-alt">
                                    分享
                                </a-button>
                                <a-button @click="report(info.id)" type="link" icon="warning">
                                    举报
                                </a-button>
                                <a-button v-if="userInfo.userId == info.userInfo.id && info.userInfo != null" @click="remove(info.id)" type="link" icon="delete">
                                    删除
                                </a-button>
                            </a-space>
                        </div>
                    </a-col>
                    <a-col class="ask-r" :span="6">
                        <div class="ask-answer">
                            <h2>关注者</h2>
                            <p>{{info.favorites | resetNum}}</p>
                        </div>
                        <div class="ask-view">
                            <h2>被浏览</h2>
                            <p>{{info.views | resetNum}}</p>
                        </div>
                    </a-col>
                </a-row>
            </div>
        </div>
        <div class="answer-detail" >
            <div class="answer-info-box" :style="{ width: design.width + 'px' }">
                <a-row :gutter="[{md:14}]">
                    <a-col class="answer-l" :span="18">
                        <div  class="answer-write">
                            <div class="answer-list-title">
                                <h2>撰写答案</h2>
                            </div>
                            <div class="answer-editor">
                                <a-textarea
                                    v-model="content"
                                    placeholder="请填写答案内容"
                                    :auto-size="{ minRows: 6, maxRows: 5 }"
                                />
                            </div>
                            <p class="answer-ds">如果权限设置为付费查看答案，请设置付费金额,默认0为免费</p>
                            <div class="answer-post">
                                <a-space>
                                    <a-button @click="createAnswer(info.id)" type="primary">提交答案</a-button>
                                    <a-input-number
                                        v-model="price"
                                        placeholder="如果权限设置为付费查看，请设置付费金额"
                                        :style="{ width: '100%' }"
                         
                                        :precision="2"
                                        :min="0"
                                    />
                                </a-space>
                            </div>
                        </div>

                        <div class="answer-list" v-if="list.length > 0">
                            <div class="answer-list-title">
                                <h2><span class="answer-list-count">{{total}}</span>个回答</h2>
                            </div>
                            
                            <ul>
                                <!-- <li v-for="(item,index) in list" :key="index">
                                    <AnswerList 
                                    @upadteView="upadteView"
                                    @remove="removeAnswer" 
                                    :info="item"/>
                                </li> -->
                                <li v-if="isShow" class="bottom">
                                    已经到底了
                                </li>
                            </ul>
                        </div>
                    </a-col>
                    <a-col class="answer-r" :span="6">
                        <SidebarUserInfo :isFollow.sync="info.isFollow"  :info="info.userInfo"/>
                        <!-- <SidebarHotEdu/> -->
                        <SidebarQuestionList />
                    </a-col>
                </a-row>
            </div>
        </div>
    </div>
</template>


<script>
import { mapState } from "vuex"
import api from "@/api/index"

import Avatar from "@/components/avatar/avatar"
// import AnswerList from "@/components/answer/List"
import SidebarHotEdu from "@/components/sidebar/sidebarHotEdu"
import SidebarUserInfo from "@/components/sidebar/sidebarUserInfo"
import SidebarQuestionList from "@/components/sidebar/sidebarQuestionList"
export default {
    components:{
        Avatar,
        // AnswerList,
        SidebarUserInfo,
        SidebarHotEdu,
        SidebarQuestionList
    },
    computed:{
        ...mapState(["design"]),
        ...mapState("user",["token","userInfo"]),
    },
    head(){
        return this.$seo(`${this.info.title}-${this.base.title}`,`${this.info.title}`,[{
            hid:"fiber-desc",
            name:"description",
            content:`${this.info.description}`
        }])
    },
    validate({ params }) {
        if (params.id != null && params.id != undefined && params.id != NaN) {
            return true // 如果参数有效
        }
        return false // 参数无效，Nuxt.js 停止渲染当前页面并显示错误页面
    },
    async asyncData({params,$axios,redirect,store}){
    
        const id = parseInt(params.id)
        const res = await $axios.get(api.getQuestion,{params:{id:id}})
        if (res.code != 1) {
            redirect("/404")
        }
        const queryParam = {
            page: 1,
            limit: 20,
            relatedId:id,
        }
        const answerRes = await $axios.get(api.getAnswerList,{params: queryParam})
        if (answerRes.code != 1) {
            redirect("/404")
        }
       
        // this.list = [...this.list,...res.data.list]
        return {
            base:store.state.base,
            id:id,
            info:res.data.info,
            queryParam: queryParam,
            list:answerRes.data.list ? answerRes.data.list : [],
            total:answerRes.data.total,
            isShow:answerRes.data.list != null ? false : true,
        }
    },
    data(){
        return{
            price:0,
            content:"",
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
        async postFavorite(id){
            if (this.token == null) {
                this.$Auth("login","登录","快速登录")
                return
            }
            this.info.isFavorite = !this.info.isFavorite
            if (this.info.isFavorite) {
                this.info.favorites =  this.info.favorites + 1
            } else {
                this.info.favorites =  this.info.favorites - 1
            }
            const query = {
                id:id
            }
            const res = await this.$axios.post(api.postQuestionFavorite,query)
            if (res.code != 1) {
                 this.$message.error(
                    res.message,
                    3
                )
                if (this.info.id == id) {
                this.info.isFavorite = !this.info.isFavorite
                    if (this.info.isFavorite) {
                        this.info.favorites =  this.info.favorites + 1
                    } else {
                        this.info.favorites =  this.info.favorites - 1
                    }
                }
                return
            }
        },
        async postLike(id){
            if (this.token == null) {
                this.$Auth("login","登录","快速登录")
                return
            }
            this.info.isLike = !this.info.isLike
            if (this.info.isLike) {
                this.info.likes =  this.info.likes + 1
            } else {
                this.info.likes =  this.info.likes - 1
            }
            const query = {
                id:id
            }
            const res = await this.$axios.post(api.postQuestionLike,query)
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
                if (this.info.id == id) {
                    this.info.isLike = !this.info.isLike
                    if (this.info.isLike) {
                        this.info.likes =  this.info.likes + 1
                    } else {
                        this.info.likes =  this.info.likes - 1
                    }
                }
                return
            }
        },
        report(id){
            this.$Report(id,"question")
        },
        async createAnswer(id){
            if (this.token == null) {
                this.$Auth("login","登录","快速登录")
                return
            }
            if (id == null || id == undefined ) {
                this.$message.error(
                    "请设置问题id",
                    3
                )
                return
            }

            if (this.content == "" ||this.content == null || this.content == undefined ) {
                this.$message.error(
                    "请编写你的答案",
                    3
                )
                return
            }

            let formData = {
                content:this.content,
                price:this.price,
                questionId:id,
            }
            const res = await this.$axios.post(api.postAnswerCreate,formData)
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
            this.$message.success(
                "创建成功",
                3
            )
            this.content = ""
            this.price = 0
            this.total += 1
            this.list = [res.data.info,...this.list]
        },
        async upadteView(){
            const res = await this.$axios.get(api.getAnswerList,{params: this.queryParam})
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
            }
            this.list = res.data.list
        },
        async getAnswerList(){
            try {
                const res = await this.$axios.get(api.getAnswerList,{params: this.queryParam})
                if (res.code != 1) {
                    this.$message.error(
                        res.message,
                        3
                    )
                }
                if (res.data.list == null) {
                    this.isShow = true
                    return
                }
                this.queryParam.page += 1
                this.list = [...this.list,...res.data.list]
            } catch (error) {
            
                setTimeout(() => {
                    this.$notification.error({
                        message: '网络错误',
                        description: "请稍后再试"
                    })
                }, 1000)
            }
        },
        scrollList(){
            const scrollTop = document.documentElement.scrollTop || document.body.scrollTop
            const clientHeight = document.documentElement.clientHeight
            const scrollHeight = document.documentElement.scrollHeight
            if (scrollTop + clientHeight >= scrollHeight) {
                if (!this.isShow) {
                    this.queryParam.page += 1
                    this.getAnswerList()
                }
                return
            }
           
        },
        removeAnswer(e){
            this.list.forEach((item,index) => {
                if (item.id == e) {
                    this.list.splice(index,1)
                    this.total -= 1
                }
            })
        },
        remove(e){
            if (this.token == null) {
                this.$Auth("login","登录","快速登录")
                return
            }
            this.$confirm({
                okText:"确定",
                cancelText:"取消",
                title: '正在删除',
                content: '请注意，您现在正在删除',
                onOk: async () => {e
                    const formData = {
                        id:parseInt(e),
                    }
                    try {
                        const res = await this.$axios.post(api.postQuestionRemove,formData)
                        if (res.code != 1) {
                            this.$message.error(
                                res.message,
                                3
                            )
                            return
                        }
                        this.$router.push({ name: "question"})
                    } catch (error) {
                  
                        setTimeout(() => {
                            this.$notification.error({
                                message: '网络错误',
                                description: "请稍后再试"
                            })
                        }, 1000)
                    }
                    return false;
                },
                onCancel() {},
            });
        }
    }
}
</script>

<style lang="less" scoped>
.content-box{
    margin-top: 70px;
    min-height: 550px;
    .ask-detail{
        background: white;
        // height: 300px;
        display: flex;
        justify-content: center;
        .ask-info-box{
            margin-top: 20px;
            .ask-l{
                .ask-title{
                    margin-top: 10px;
                    font-size: 22px;
                    font-weight: 600;
                    line-height: 32px;
                    color: #1a1a1a;
                }
                .date{
                    margin-top: 10px;
                    font-size: 12px;
                    color: #777;
                    .ask-date{
                        margin-left: 10px;
                    }
                }
                .ask-desc{
                    margin-top: 10px;
                    font-size: 14px;
                    color: #777;
                }
                .ask-btn{
                    margin: 20px 0;
                    /deep/ .ant-btn-link{
                        color: #8590a6;
                        padding: 0 5px;
                    }
                }
            }
            .ask-r{
                display: flex;
                justify-content: space-around;
                .ask-answer{
                    h2{
                        font-weight: 600;
                        font-size: 16px;
                    }
                    display: flex;
                    flex-direction: column;
                    align-items: center;
                    justify-content: center;
                }
                .ask-view{
                    h2{
                        font-weight: 600;
                        font-size: 16px;
                    }
                    display: flex;
                    flex-direction: column;
                    align-items: center;
                    justify-content: center;
                }
            }
        }
    }
    .answer-detail{
        margin-top: 10px;
        margin-bottom: 20px;
        // height: 300px;
        display: flex;
        justify-content: center;
        .answer-l{
            .answer-write{
                background: white;
                padding: 10px 20px 10px 20px;
                .answer-editor{
                    margin-top: 10px;
                    // min-height: 200px;
                }
                .answer-ds{
                    margin: 10px 0;
                    font-size: 12px;
                    color: #8590a6;
                }
            }
            .answer-list-title{
                padding-bottom: 10px;
                border-bottom: 1px #f0f2f7 solid;
                h2{
                    font-size: 14px;
                    .answer-list-count{
                        margin-right: 10px;
                        font-size: 14px;
                    }
                }
                
            }
            .answer-list{
                margin: 10px 0;
                background: white;
                padding: 10px 20px 10px 20px;
                ul{
                    margin-top: 10px;
                    .bottom{
                        margin-top: 10px;
                        text-align: center;
                    }
                }
            }
        }
    }
}
</style>
