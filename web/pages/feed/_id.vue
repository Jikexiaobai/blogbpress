<template>
    <div class="detail">
        <div class="info" :style="{ width: design.width + 'px' }">
            <a-row :gutter="[{md:12}]">
                <a-col :span="18">
                    <div class="feed-info">
                        <div class="user-info">
                            <div class="user-info-l">
                                <nuxt-link :to="{path:'/profile/' + info.userInfo.id }" class="item-link">   

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
                                <div class="user-name-box">
                                    <nuxt-link :to="{path:'/profile/' + info.userInfo.id }" class="item-link">   
                                        <h2 class="user-name">{{info.userInfo.nickName}}</h2>
                                    </nuxt-link>
                                    <div  class="user-role">
                                        <a-space>
                                            <img :src="info.userInfo.grade.icon" :alt="info.userInfo.grade.title">
                                            <img v-if="info.userInfo.vip" :src="info.userInfo.vip.icon" :alt="info.userInfo.vip.title">
                                        </a-space>
                                    </div>
                                </div>
                            </div>
                            <div class="user-info-r">
                                <a-button @click="remove" v-if="userInfo.userId == info.userInfo.id" type="link" icon="delete">
                                    删除
                                </a-button>
                                <a-button @click="report" type="link" icon="info-circle">
                                    举报
                                </a-button>
                            </div>
                        </div>
                        <div  class="feed-content">
                            <div>
                                <p class="feed-text">{{info.title}}</p>
                                <div v-if="info.type == 1" class="feed-img-box">
                                    <div v-if="info.files.length == 1" class="feed-img-one" :style="{ width: this.width + 'px'}">
                                        <img  :src="info.files[0]">
                                    </div>
                                    <ul v-if="info.files.length > 1">
                                        <li v-for="(item,index) in info.files.slice(0, 3)" :key="index">
                                            <div>
                                                <img :src="item" alt="">
                                                <span v-if="index === 2 && info.files.length > 3" class="image-number">
                                                    +<b v-text="info.files.length - 3"></b>
                                                </span>
                                            </div>
                                        </li>
                                    </ul>
                                </div>
                                <div @click="goRelated" v-if="info.type == 2" class="feed-share-box">
                                    <img class="cover" :src="info.relatedInfo.cover" :alt="info.relatedInfo.title">
                                    <div class="title-desc">
                                        <h2>{{info.relatedInfo.title}}</h2>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="feed-meta">
                            <div class="feed-meta-l">
                                <button @click="postLike" class="text">
                                    <a-icon :theme="info.isLike ? 'filled' : 'outlined'" type="like" />
                                    <span>{{info.isLike ? '已赞' : '赞'}} </span>
                                    <b>{{info.likes == 0 ? "" : info.likes}}</b>
                                </button>
                                <button 
                                    v-clipboard:copy="'aasdasdasd'"
                                    v-clipboard:success="onCopy" class="text">
                                    <a-icon type="share-alt" />
                                    <span>复制链接</span>
                                </button>
                                <nuxt-link  :to="`/group/${info.groupInfo.id}`" class="item-link"> 
                                    <a-tag color="#f50">
                                        #{{info.groupInfo.title}}
                                    </a-tag>
                                </nuxt-link>
                            </div>
                            <div class="feed-meta-r">
                                <span class="feed-meta-date">
                                    {{info.createTime | resetData}}
                                </span>
                            </div>
                        </div>
                    </div>
                    <div class="comment-box">
                        <Comment  module="topic" :relatedId="id"/>
                    </div>
                </a-col>
                <a-col :span="6">
                    <SidebarUserInfo :isFollow.sync="info.isFollow"  :info="info.userInfo"/>
                </a-col>
            </a-row>
        </div>
    </div>
</template>



<script>
import Avatar from "@/components/avatar/avatar"
import Comment from "@/components/comment/List"
import SidebarUserInfo from "@/components/sidebar/sidebarUserInfo"

import { mapState } from "vuex"
import api from "@/api/index"
export default {
    computed:{
        ...mapState(["design","title"]),
        ...mapState("user",["token","userInfo"]),
    },
    components:{
        Avatar,
        Comment,
        SidebarUserInfo,
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
        const res = await $axios.get(api.getTopic,{params:{id:id}})
    
        if (res.code != 1) {
            redirect("/404")
        }
        if (res.data.info.type == 1 && res.data.info.files != "") {
            res.data.info.files = JSON.parse(res.data.info.files)
        }

        return {
            base:store.state.base,
            id:id,
            info:res.data.info
        }
    },
    methods: { 
        goRelated(){
            this.$router.push(`/${this.info.relatedInfo.module}/${this.info.relatedInfo.id}`)
        },
        goGroup(e){
            this.$router.push(`/group/${e}`)
        },
        report(){
            this.$Report(this.id,"topic")
        },
        async postLike(){
            if (this.token == null) {
                this.$Auth("login","登录","快速登录")
                return
            }
            this.info.isLike = !this.info.isLike
            if (this.info.isLike) {
                this.info.likes = this.info.likes + 1
            } else {
                 this.info.likes = this.info.likes - 1
            }
            const query = {
                id:this.id
            }
            const res = await this.$axios.post(api.postTopicLike,query)
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
                this.info.isLike = !this.info.isLike
                if (this.info.isLike) {
                    this.info.likes = this.info.likes + 1
                } else {
                    this.info.likes = this.info.likes - 1
                }
                return
            }
        },
        onCopy(e){
            this.$message.success(
                "复制成功",
                3
            )
        },
        async remove(){
            this.$confirm({
                okText:"确定",
                cancelText:"取消",
                title: '正在删除',
                content: '请注意，您现在正在删除',
                onOk:() => {
                    const formData = {
                        id:this.id,
                    }
                    this.postDelete(formData)
                    return false;
                },
                onCancel() {},
            });
        },
        async postDelete(formData){
            try {
                const res = await this.$axios.post(api.postTopicRemove,formData)
                if (res.code != 1) {
                    this.$message.error(
                        res.message,
                        3
                    )
                    return
                }
                 this.$router.push({ name: "feed"})
            } catch (error) {
                console.log(error)
                setTimeout(() => {
                    this.$notification.error({
                        message: '网络错误',
                        description: "请稍后再试"
                    })
                }, 1000)
            }
        },

    }
}
</script>

<style lang="less" scoped>
.detail{
    margin-top: 80px;
    display: flex;
    justify-content: center;
    min-height: 550px;
    .info{
        min-height: 550px;
        .feed-info{
            background: white;
            padding: 20px;
            .user-info{
                display: flex;
                justify-content: space-between;
                .user-info-l{
                    display: flex;
                    align-items: center;
                    margin-right: 10px;
                    .user-name-box{
                        .user-name{
                            font-size: 14px;
                        }
                        .user-meta{
                            font-size: 12px;
                        }
                        .user-role{
                            img{
                                max-height: 18px;
                                max-width: 18px;
                            }
                        }
                    }
                }
                .user-info-r{
                    color: #d0d4dc;
                    font-size: 14px;
                    /deep/ .ant-btn-link{
                        color: #d0d4dc;
                    }
                }
            }
            .feed-content{
                margin: 10px 0;
                .feed-text{
                    font-size: 14px!important;
                }
                .feed-img-box{
                    margin: 10px 0;
                    .feed-img-one{
                        img{
                            width: 100%;
                            height: 100%;
                        }
                    }
                    ul{
                        display: flex;
                        margin: -5px;
                        li{
                            width: 33.333333%;
                            padding: 5px;
                            div{
                                height: 0;
                                padding-top: 100%;
                                cursor: pointer;
                                overflow: hidden;
                                position: relative;
                                transition: padding-top .2s;
                                max-width: 100%;
                                img{
                                    position: absolute;
                                    left: 0;
                                    top: 0;
                                    background-color: #f5f5f5;
                                    width: 100%;
                                    height: 100%;
                                    display: block;
                                }
                                .image-number{
                                    position: absolute;
                                    right: 10px;
                                    top: 10px;
                                    height: 24px;
                                    line-height: 24px;
                                    border-radius: 15px;
                                    -webkit-backdrop-filter: blur(5px);
                                    backdrop-filter: blur(5px);
                                    padding: 0 8px;
                                    font-size: 13px;
                                    font-weight: 500;
                                    color: #fff;
                                    white-space: nowrap;
                                    background-color: rgba(26,26,26,.3);
                                }
                                
                            }
                        }
                    }
                }
                .feed-share-box{
                    cursor: pointer;
                    margin: 10px 0;
                    margin-top: 10px;
                    background-color: #F7F8FA;
                    display: flex;
                    cursor: pointer;
                    .cover{
                        height: 80px;
                        width: 80px;
                    }
                    .title-desc{
                        flex:1;
                        margin-left: 10px;
                        margin-top: 5px;
                        h2{
                            font-size: 18px;
                            font-weight: bold;
                            display: -webkit-box;
                            -webkit-box-orient: vertical;
                            -webkit-line-clamp: 1;
                            overflow: hidden;
                            text-justify: inter-ideograph;
                            word-break: break-all;
                        }
                        p{
                            font-size: 14px;
                            font-weight: bold;
                            display: -webkit-box;
                            -webkit-box-orient: vertical;
                            -webkit-line-clamp: 2;
                            overflow: hidden;
                            text-justify: inter-ideograph;
                            word-break: break-all;
                        }
                    }
                }
                .feed-show{
                    position: relative;
                    background: #f5f5f5;
                    max-width: 390px;
                    width: 100%;
                    margin-top: 10px;
                    padding: 16px;
                    .feed-show-desc{
                        padding-bottom: 10px;
                        border-bottom: 1px solid #e5e5e5;
                        margin-bottom: 10px;
                        font-size: 13px;
                        line-height: 1;
                    }
                    .feed-show-ac{
                        p{
                            margin-bottom: 16px;
                            font-size: 12px;
                            color: #878787;
                        }
                    }
                }
            }
            .feed-meta{
                display: flex;
                justify-content: space-between;
                align-items: center;
                .feed-meta-l{
                    display: flex;
                    position: relative;
                    .text{
                        outline: none;
                        -webkit-tap-highlight-color: rgba(0,0,0,0);
                        font-family: font-regular,'Helvetica Neue',sans-serif;
                        border: 1px solid #ccc;
                        box-sizing: border-box;
                        margin-right: 10px;
                        font-size: 12px;
                        border-radius: 2px;
                        border: 0;
                        padding: 0 6px;
                        background: #f5f5f5;
                        color: #b9b9b9;
                        cursor: pointer;
                        font-weight: 400;
                        display: flex;
                        align-items: center;
                        span{
                            padding: 0 3px;
                            margin: 0;
                        }
                    }
                    
                }
                .feed-meta-r{
                    .feed-meta-date{
                        height: 20px;
                        line-height: 20px;
                        font-size: 12px;
                        color: #b9b9b9;
                    }
                }
            }
        }
        .comment-box{
            margin: 20px 0;
            padding: 20px;
            background: white; 
            ul{
                li{
                    margin-bottom: 10px;
                }
            }
        }
    }  
}
</style>