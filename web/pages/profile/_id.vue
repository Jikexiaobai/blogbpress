<template>
    <div class="profile-box">
        <div class="profile-info" :style="{ width: design.width + 'px' }">
            <div class="profile-info-cover" :style="{ backgroundImage: `url(${info.cover})` }">
                <div class="profile-info-user">
                    <div class="profile-info-username-avatar">
                        <Avatar 
                            isVerify 
                            :src="info.avatar+'@w60_h60'" 
                            shape="circle" 
                            :size="60"
                        />
                        <div>
                            <a-space>
                                <h2 class="username">{{info.nickName}}</h2>
                                <div class="user-role">
                                    <span class="grade-title">
                                        {{info.grade.title}}
                                    </span>
                                    <span class="vip-title" v-if="info.vip">
                                        {{info.vip.title}}
                                    </span>
                                    <!-- <img :src="info.grade.icon" :alt="info.grade.title">
                                    <img v-if="info.vip != null" :src="info.vip.icon" :alt="info.vip.title"> -->
                                </div>
                            </a-space>
                            <p class="desc">{{info.description}}</p>
                        </div>
                    </div>
                    <div v-if="token == null || userInfo.userId != info.id || info.isFollow" class="profile-info-follow">
                        <a-button @click="follow" type="primary">
                            {{info.isFollow ? "取消关注" : "关注"}}
                        </a-button>
                        <!-- <a-button>私信</a-button> -->
                    </div>
                </div>
            </div>
            <div class="profile-info-nav">
                <ul class="profile-info-nav-l">
                    <li>
                        <nuxt-link :to="{path:'/profile/' + id}" class="item-link">
                            <span class="nav-title">主页</span>
                        </nuxt-link>
                    </li>
                    <li>
                        <nuxt-link :to="{path:'/profile/' + id + '/question'}" class="item-link">
                            <span class="nav-title">提问</span>
                        </nuxt-link>
                    </li>
                    <li>
                        <nuxt-link :to="{path:'/profile/' + id + '/feed'}" class="item-link">
                            <span class="nav-title">动态</span>
                        </nuxt-link>
                    </li>
                    <li>
                        <nuxt-link :to="{path:'/profile/' + id + '/content/article'}"  class="item-link">
                            <span class="nav-title">投稿</span>
                        </nuxt-link>
                    </li>
                </ul>
                <ul class="profile-info-nav-r">
                    <li>
                        <nuxt-link  :to="{path:'/profile/' + id + '/follow'}"  class="item-link">
                            <span class="nav-title">关注</span>
                            <span class="nav-num">{{info.follows | resetNum}}</span>
                        </nuxt-link>
                    </li>  
                    <li>
                        <nuxt-link  :to="{path:'/profile/' + id + '/fans'}"  class="item-link">
                            <span class="nav-title">粉丝</span>
                            <span class="nav-num">{{info.fans | resetNum}}</span>
                        </nuxt-link>
                    </li>
                    <li>
                        <span class="nav-title">获赞</span>
                        <span class="nav-num">{{info.likes | resetNum}}</span>
                    </li>
                </ul>
            </div>
            <div class="profile-info-body">
                <a-row :gutter="{ md: '20'}">
                    <a-col :md="18">
                        <Nuxt/>
                    </a-col>
                    <a-col :md="6">
                        <SidebarReward />
                        <!-- <SidebarUserGroup />
                        <SidebarHotUserList /> -->
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
import SidebarReward from "@/components/sidebar/sidebarReward"
import SidebarUserGroup from "@/components/sidebar/sidebarUserGroup"
import SidebarHotUserList from "@/components/sidebar/sidebarHotUserList"
import SidebarAccess from "@/components/sidebar/sidebarAccess"

export default {
    components:{
        Avatar,
        SidebarReward,
        SidebarUserGroup,
        SidebarHotUserList,
        SidebarAccess
    },
    head(){
        return this.$seo(`${this.info.nickName}-${this.base.title}`,`${this.info.nickName}`,[{
            // hid:"piankr",
            name:"description",
            content:`${this.info.nickName}`
        }])
    },
    validate({ params }) {
        if (isNaN(params.id)) {
            return false // 如果参数有效
        }
        return true // 参数无效，Nuxt.js 停止渲染当前页面并显示错误页面
    },
    async asyncData({params,$axios,redirect}){
    
        const id = parseInt(params.id)
        const res = await $axios.get(api.getUserInfo,{params:{id:id}})
        if (res.code != 1) {
            redirect("/404")
        }

        return {
            id:id,
            info:res.data.info,
        }
    },
    data(){
        return{
            id:null
        }
    },
    computed:{
        ...mapState(["design","base"]),
        ...mapState("user",["token","userInfo"]),
    },
    methods:{
        async follow(){
            if (!this.token) {
                this.$Auth("login","登录","快速登录")
                return
            }
            if (this.info.isFollow) {
                this.info.fans -= 1
            }else{
                this.info.fans += 1
            }
            this.info.isFollow = !this.info.isFollow
            const res = await this.$axios.post(api.postUserFollow,{id:this.info.id})
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
                 if (this.info.isFollow) {
                this.info.fans -= 1
                }else{
                    this.info.fans += 1
                }
               this.info.isFollow = !this.info.isFollow
                return
            }
        },
    }
}
</script>

<style lang="less" scoped>
.profile-box{
    margin-top: 80px;
    min-height: 550px;
    display: flex;
    justify-content: center;
    .profile-info{
        .profile-info-cover{
            height: 200px;
            background-position: 50%;
            background-size: cover;
            transition: background-image .2s ease,background-size 1s ease;
            padding-top: 100px;
          
            .profile-info-user{
                display: flex;
                justify-content: space-between;
                padding: 20px;
                .profile-info-username-avatar{
                    display: flex;
                    align-items: center;
                    flex: 1;
                    margin-right: 20px;
                    .username{
                        color: white;
                        display: inline-block;
                        margin-right: 5px;
                        font-weight: 700;
                        line-height: 32px;
                        font-size: 18px;
                        vertical-align: middle;
                    }
                    .user-role{
                        font-size: 12px;
                        color: #090909;
                        // display: flex;
                        // align-items: center;
                        // img{
                        //     max-height: 20px;
                        //     max-width: 20px;
                        // }
                    }
                    .desc{
                        background: transparent;
                        border-radius: 4px;
                        border: none;
                        box-shadow: none;
                        color: #d6dee4;
                        font-size: 12px;
                        font-family: Microsoft Yahei;
                        line-height: 26px;
                        height: 26px;
                        margin-left: -5px;
                        padding: 0 5px;
                        font-weight: 400;
                    }
                }
            }
        }
        .profile-info-nav{
            height: 66px;
            background: #fff;
            box-shadow: 0 0 0 1px #eee;
            border-radius: 0 0 4px 4px;
            padding: 0 20px;
            display: flex;
            align-items: center;
            justify-content: space-between;
            .profile-info-nav-l{
                display: flex;
                align-items: center;
                flex: 1;
                height: 66px;
                li{
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    flex-direction: column;
                    height: 66px;
                    position: relative;
                    margin-right: 20px;
                   .nav-icon{
                        font-size: 16px;
                        margin-right: 5px;
                    }
                    .nav-title{
                        font-size: 16px;
                    }
                    .nav-num{
                        font-size: 12px;
                        color: #99a2aa;
                        font-family: Arial;
                    }
                    a{
                        color: #262626;
                    }
                    .nuxt-link-exact-active{
                        color: #262626;
                        font-weight: 600;
                    }
                    .nuxt-link-exact-active::after{
                        content: "";
                        position: absolute;
                        width: 20px;
                        height: 4px;
                        background-color: #4560c9;
                        bottom: 14px;
                        left: 50%;
                        transform: translateX(-50%);
                    } 
                }
            }
            .profile-info-nav-r{
                display: flex;
                align-items: center;
                height: 66px;
                li{
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    flex-direction: column;
                    height: 66px;
                    position: relative;
                    margin-right: 20px;
                    .item-link{
                        display: flex;
                        align-items: center;
                        justify-content: center;
                        flex-direction: column;
                        
                    }
                    .nav-title{
                        font-size: 14px;
                    }
                    .nav-num{
                        font-size: 12px;
                        color: #99a2aa;
                        font-family: Arial;
                    }
                    a{
                        color: #262626;
                    }
                    .nuxt-link-exact-active{
                        color: #262626;
                        font-weight: 600;
                        color: #00a1d6;
                    }
                }
            }
        }
        .profile-info-body{
            margin: 20px 0;
        }
    }
}
</style>
