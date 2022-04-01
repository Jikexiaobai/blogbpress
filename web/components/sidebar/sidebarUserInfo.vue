<template>
    <div class="sidear-box">
        <div class="user-cover" :style="{ backgroundImage: `url(${info.cover}@w300_h100)` }"></div>
        <div class="user-info">
            <nuxt-link :to="{path:'/profile/' + info.userId }" class="item-link"> 
                <!-- <Avatar  :src="info.avatar" class="user-avatar" :size="60" /> -->
                <Avatar 
                    class="user-avatar"
                    :verifyRight="-5"
                    :verifyBottom="5"
                    :isVerify="info.isVerify"
                    shape="square" 
                    :src="info.avatar+'@w60_h60'" 
                    :size="55"
                />
            </nuxt-link>
            <div class="user-name">
                <nuxt-link :to="{path:'/profile/' + info.userId }" class="item-link"> 
                    <h2>{{info.nickName}}</h2>
                </nuxt-link>
                <div class="user-info-l-meta">
                    <a-space>
                        <span class="grade-title">
                            {{info.grade.title}}
                        </span>
                        <!-- <span class="vip-title" v-if="info.vip">
                            {{info.vip.title}}
                        </span> -->
                    </a-space>
                </div>
            </div>
        </div>
        <ul class="user-meta">
            <!-- <li><p>投稿</p><span>{{info.posts | resetNum}}</span></li> -->
            <li><p>获赞</p><span>{{info.likes | resetNum}}</span></li>
            <li><p>关注</p><span>{{info.follows | resetNum}}</span></li>
            <li><p>粉丝</p><span>{{info.fans | resetNum}}</span></li>
        </ul>
        <div v-if="userInfo.userId != info.userId || info.isFollow" class="user-ac" >
            <a-button @click="follow" type="primary">
                {{info.isFollow ? "取消关注" : "关注"}}
            </a-button>
            <a-button @click="charge" type="dashed">
                给他充电
            </a-button>
        </div>
    </div>
</template>

<script>
import Avatar from "@/components/avatar/avatar"
import { mapState } from "vuex"
import api from "@/api/index"
import {MODULE} from "@/shared/module"
import {ORDERTYPE} from "@/shared/order"
export default {
    props:{ 
        info: {
            type: Object, //指定传入的类型
            default: {
                cover:""
            } //这样可以指定默认的值
        },
    },
    components:{
        Avatar,
    },
    computed:{
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
            const res = await this.$axios.post(api.postUserFollow,{id:this.info.userId})
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
        charge(){
            if (!this.token) {
                this.$Auth("login","登录","快速登录")
                return
            }
            
            const product = {
                detailId:this.info.userId,
                detailModule:MODULE.USER,
                orderType: ORDERTYPE.CD,
            }

            this.$Charge("充电",product,this.info.avatar,this.info.nickName).then(async (res)=>{
                if (res != false) {
                    this.$message.success(
                        "充电成功",
                        3
                    )
                    
                }
            }).catch((err)=>{
                this.$message.error(
                    "充电失败",
                    3
                )
                // this.createForm.cover = undefined
            })
            
            
            
        }
    }
}
</script>
<style lang="less" scoped>
.sidear-box{
    margin-bottom: 10px;
    background: white;
    .user-cover{
        // background-image: url("/img/2.jpg");
        width: 100%;
        padding-bottom: 90px;
        background-position: center center;
        background-size: cover;
        // border-radius: 5px;
        margin-right: 20px;
        // box-shadow: rgba(0, 0, 0, 0.2) 0px 10px 15px 0px;
    }
    .user-info{
        position: relative;
        .user-avatar{
            position: absolute;
            top: -15px;
            left: 12px;
        }
        .user-name{
            margin-left: 70px;
            padding-top: 0px;
            // height: 55px;
            h2{
                font-size: 16px;
                font-weight: 600;
                line-height: 20px;
                max-width: 100%;
                display: -webkit-box;
                -webkit-box-orient: vertical;
                -webkit-line-clamp: 1;
                overflow: hidden;
                height: 20px;

            }
            .user-info-l-meta{
                .grade-title{
                    font-size: 12px;
                    font-style: normal;
                    display: inline-block;
                    background-color: rgba(173,173,173,0.16);
                    transform: scale(1);
                    height: 17px;
                    line-height: 17px;
                    padding: 0 6px;
                    border-radius: 2px;
                    text-transform: capitalize;
                }
            }
        }
    }
    .user-meta{
        display: flex;
        justify-content: space-between;
        border-top: 1px solid #ebeef5;
        margin: 10px 20px;
        padding-top: 10px;
        height: 50px;
        p{
            font-size: 13px;
            color: #bcbcbc;
            margin-bottom: 2px;
        }
        span{
            font-size: 12px;
        }
    }
    .user-ac{
        margin: 0 20px;
        padding-bottom: 10px;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }
}
</style>