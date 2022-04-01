<template>
    <div class="answer-item">
        <div class="answer-user">
            <nuxt-link v-if="info.userInfo != null" :to="{path:'/profile/' + info.userInfo.id }" class="item-link">   
                <Avatar 
                    class="user-avatar"
                    :verifyRight="-5"
                    :verifyBottom="5"
                    :isVerify="info.userInfo.isVerify"
                    shape="circle" 
                    :src="info.userInfo.avatar+'@w60_h60'" 
                    :size="45"
                />
            </nuxt-link>
            <div class="user-info">
                <div class="user-name-box" v-if="info.userInfo != null">
                    <nuxt-link :to="{path:'/profile/' + info.userInfo.id }" class="item-link">   
                        <h2 class="user-name">{{info.userInfo.nickName}}</h2>
                    </nuxt-link>
                    <span>{{info.createTime | resetData}}</span>
                </div>
                <div class="user-role">
                    <a-space>
                        <img :src="info.userInfo.grade.icon" :alt="info.userInfo.grade.title">
                        <img v-if="info.userInfo.vip" :src="info.userInfo.vip.icon" :alt="info.userInfo.vip.title">
                    </a-space>
                </div>
            </div>
        </div>
        <div class="entry-content" v-html="info.content" v-if="info.isView" ref='imgTooles'>       
        </div>
        <div class="answer-yes" v-if="info.isView">
            <a-space >
                <a-button @click="postLike(info.id)" :type="info.isLike ? 'primary' : ''">
                    <a-icon :theme="info.isLike ? 'filled' : 'outlined'" :type="info.isLike ? 'caret-down' : 'caret-up'" />
                    {{info.isLike ? '已赞同' : '赞同'}} {{info.likes == 0 ? "" : info.likes}}
                </a-button>
                <a-button @click="openComment(info.id)" type="link" icon="message">
                    {{info.id == openCommentkey ? '收起评论' : '评论'}}
                </a-button>
                <!-- <a-button type="link" icon="star">
                    收藏
                </a-button> -->
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
         <div class="answer-show" v-if="!info.isView">
            <div class="answer-show-desc">
                <a-icon type="api" />
                付费可见
            </div>
            <div  class="answer-show-ac">
                <p>
                    支付费用阅读隐藏答案
                </p>
                <a-button @click="postPay" type="primary">
                    付费
                </a-button>
            </div>
        </div>
        <div v-if="info.id == openCommentkey" class="answer-comment">
             <Comment  module="answer" :relatedId="info.id"/>
        </div>
    </div>
</template>

<style lang="less" scoped>
.answer-item{
    padding-bottom: 10px;
    border-bottom: 1px #f0f2f7 solid;
    .answer-user{
        display: flex;
        margin: 10px 0;
        .user-info{
            flex: 1;
            .user-name-box{
                display: flex;
                justify-content: space-between;
                h2{
                    font-size: 14px;
                }
            }
            .user-role{
                img{
                    max-height: 20px;
                    max-width: 20px;
                }
            }
        }
        
    }
    .answer-yes{
        display: flex;
        align-items: center;
        margin-top: 10px;
        /deep/ .ant-btn-link{
            color: #8590a6;
            padding: 0 5px;
        }
    }
    .answer-show{
        position: relative;
        background: #f5f5f5;
        max-width: 390px;
        width: 100%;
        margin-top: 10px;
        padding: 16px;
        .answer-show-desc{
            padding-bottom: 10px;
            border-bottom: 1px solid #e5e5e5;
            margin-bottom: 10px;
            font-size: 13px;
            line-height: 1;
        }
        .answer-show-ac{
            p{
                margin-bottom: 16px;
                font-size: 12px;
                color: #878787;
            }
        }
    }
}
</style>

<script>
import { mapState } from "vuex"
import api from "@/api/index"
import Avatar from "@/components/avatar/avatar"
import Comment from "@/components/comment/List"
import {MODULE} from "@/shared/module"
import {ORDERTYPE} from "@/shared/order"
export default {
    props: {
        info:{
            type: Object,
            default: null
        }
    },
    components:{
        Avatar,
        Comment
    },
    computed:{
        ...mapState("user",["token","userInfo"]),
    },
    data(){
        return{
            openCommentkey:null,
        }
    },
    methods:{
        async upadteView(){
            const res = await this.$axios.get(api.getAnswer,{params:{id:this.info.id}})
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
            this.info = res.data.info
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
            const res = await this.$axios.post(api.postAnswerLike,query)
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
        openComment(id){
            if (this.openCommentkey == id) {
                this.openCommentkey = null
            }else{
                this.openCommentkey = id
            }
        },
        postPay(){
            if (this.token != null) {
                const product = {
                    detailId:this.info.id,
                    authorId:this.info.userInfo.id,
                    detailModule:MODULE.ANSWER,
                    orderMoney:this.info.price,
                    orderType: ORDERTYPE.VIEWANSWER,
                }
                this.$Pay("查看付费答案",product).then(async (res)=>{
                    if (res != false) {
                       
                        this.$emit("upadteView")
                    }
                }).catch((err)=>{
                    console.log(err)
                    // this.createForm.cover = undefined
                })
                
                
            } else {
                this.$Auth("login","登录","快速登录")
            }
        },
        report(id){
            this.$Report(id,"answer")
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
                        const res = await this.$axios.post(api.postAnswerRemove,formData)
                        if (res.code != 1) {
                            this.$message.error(
                                res.message,
                                3
                            )
                            return
                        }
                        this.$emit("remove",e)
                    } catch (error) {
                        console.log(error)
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