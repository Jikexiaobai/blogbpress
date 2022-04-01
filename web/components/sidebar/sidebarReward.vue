<template>
    <div class="sidebar-box trigger">
        <div class="info">
            <div class="icon-b" @click="charge">
                <div class="text-b">
                    <a-icon type="gift" class="icon"/>
                    <div class="hi">给他充电</div>
                </div>
            </div>
            <div class="count">
                <svg viewBox="0 0 1028 385" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M1 77H234.226L307.006 24H790" stroke="#e5e9ef" stroke-width="20"></path> <path d="M0 140H233.035L329.72 71H1028" stroke="#e5e9ef" stroke-width="20"></path> <path d="M1 255H234.226L307.006 307H790" stroke="#e5e9ef" stroke-width="20"></path> <path d="M0 305H233.035L329.72 375H1028" stroke="#e5e9ef" stroke-width="20"></path> <rect y="186" width="236" height="24" fill="#e5e9ef"></rect> <ellipse cx="790" cy="25.5" rx="25" ry="25.5" fill="#e5e9ef"></ellipse> <circle r="14" transform="matrix(1 0 0 -1 790 25)" fill="white"></circle> <ellipse cx="790" cy="307.5" rx="25" ry="25.5" fill="#e5e9ef"></ellipse> <circle r="14" transform="matrix(1 0 0 -1 790 308)" fill="white"></circle></svg>
                <div class="mask">
                    <svg viewBox="0 0 1028 385" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M1 77H234.226L307.006 24H790" stroke="#f25d8e" stroke-width="20"></path> <path d="M0 140H233.035L329.72 71H1028" stroke="#f25d8e" stroke-width="20"></path> <path d="M1 255H234.226L307.006 307H790" stroke="#f25d8e" stroke-width="20"></path> <path d="M0 305H233.035L329.72 375H1028" stroke="#f25d8e" stroke-width="20"></path> <rect y="186" width="236" height="24" fill="#f25d8e"></rect> <ellipse cx="790" cy="25.5" rx="25" ry="25.5" fill="#f25d8e"></ellipse> <circle r="14" transform="matrix(1 0 0 -1 790 25)" fill="white"></circle> <ellipse cx="790" cy="307.5" rx="25" ry="25.5" fill="#f25d8e"></ellipse> <circle r="14" transform="matrix(1 0 0 -1 790 308)" fill="white"></circle></svg>
                </div>
                <div class="color-mask">
                    <svg viewBox="0 0 1028 385" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M1 77H234.226L307.006 24H790" stroke="#ffd52b" stroke-width="20"></path> <path d="M0 140H233.035L329.72 71H1028" stroke="#ffd52b" stroke-width="20"></path> <path d="M1 255H234.226L307.006 307H790" stroke="#ffd52b" stroke-width="20"></path> <path d="M0 305H233.035L329.72 375H1028" stroke="#ffd52b" stroke-width="20"></path> <rect y="186" width="236" height="24" fill="#ffd52b"></rect> <ellipse cx="790" cy="25.5" rx="25" ry="25.5" fill="#ffd52b"></ellipse> <circle r="14" transform="matrix(1 0 0 -1 790 25)" fill="white"></circle> <ellipse cx="790" cy="307.5" rx="25" ry="25.5" fill="#ffd52b"></ellipse> <circle r="14" transform="matrix(1 0 0 -1 790 308)" fill="white"></circle></svg>
                </div>
                <span>共{{total | resetNum}}</span>
            </div>
            <!-- <a-button v-if="id != userInfo.userId" @click="charge" size="large" type="primary">
                点击充电打赏
            </a-button>
            <span>已有{{total | resetNum}}人充电打赏</span> -->
        </div>
        <ul v-if="list.length > 0">
            <li v-for="(item,index) in list" :key="index" @click="goProfile(item.userId)">
                <Avatar shape="circle" :src="item.avatar+'@w60_h60'"/>
            </li>
        </ul>
    </div>
</template>

<script>
import api from "@/api/index"
import { mapState } from "vuex"
import {MODULE} from "@/shared/module"
import {ORDERTYPE} from "@/shared/order"
export default {
    props:{ 
        avatar: {
            type: String, //指定传入的类型
            default: "" //这样可以指定默认的值
        },
        nickName: {
            type: String, //指定传入的类型
            default: "" //这样可以指定默认的值
        },
    },
    data(){
        return{
            id:0,
            list:[],
            total:0,
        }
    },
    computed:{
        ...mapState("user",["token","userInfo"]),
    },
    mounted(){
        this.id = this.$route.params.id
        this.getData()
    },
    methods:{
        async getData(){
            const res = await this.$axios.get(api.getUserReward,{params:{id:this.id}})
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
            this.list = res.data.list == null ? [] : res.data.list
            this.total = res.data.total == 0 ? 0 : res.data.total
        },
        goProfile(e){
           this.$router.push({ path: `/profile/${e}`})
        },
        charge(){
            if (!this.token) {
                this.$Auth("login","登录","快速登录")
                return
            }
             if (this.id == this.userInfo.userId) {
                this.$message.error(
                    "哪有自己给自己打赏的？",
                    3
                )
                return
            }
            
            const product = {
                detailId:this.id,
                detailModule:MODULE.USER,
                orderType: ORDERTYPE.CD,
            }
            this.$Charge("充电",product).then(async (res)=>{
                if (res != false) {
                    this.$message.success(
                        "充电成功",
                        3
                    )
                    this.getData()
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
.sidebar-box{
    background: white;
    // padding: 20px;
    margin-bottom: 10px;
    .sidebar-title{
        border-bottom: 1px solid #e5e9ef;
        font-size: 14px;
        font-weight: 700;
        padding: 0;
        margin: -15px 0 10px;
        height: 45px;
        line-height: 45px;
    }
} 
.trigger{
    .info{
        // display: flex;
        // justify-content: space-between;
        // align-items: center;
        // width: 320px;
        height: 85px;
        position: relative;
        border-radius: 4px;
        .icon-b{
            width: 122px;
            height: 45px;
            background-color: #f25d8e;
            position: absolute;
            top: 50%;
            left: 5%;
            transform: translateY(-50%);
            border-radius: 4px;
            cursor: pointer;
            z-index: 2;
            .text-b{
                width: 100px;
                height: 100%;
                margin: 0 auto;
                position: relative;
                display: flex;
                align-items: center;
                color: #fff;
                .icon{
                    font-size: 15px;
                }
                .hi{
                    margin-left: 10px;
                    float: right;
                    line-height: 45px;
                    font-size: 15px;
                    color: #fff;
                }
            }
        }
        .count{
            width: 157px;
            height: 55px;
            position: absolute;
            right: 10px;
            top: 15px;
            .mask{
                width: 0;
                height: 100%;
                overflow: hidden;
                position: absolute;
                top: 0;
                left: 0;
                transition: all .5s;
            }
            .color-mask{
                width: 18px;
                height: 100%;
                overflow: hidden;
                position: absolute;
                left: -15px;
                top: 0;
            }
            span{
                position: absolute;
                right: 10px;
                top: 18px;
                font-size: 12px;
                font-family: "雅黑";
                color: #aaa;
            }
        }
        .icon-b:hover{
            background-color: #ff6b9a;
        }
    }
    ul{
        padding: 10px;
        display: flex;
        flex-flow: wrap;
        li{
            cursor: pointer;
            width: 10%;
        }
    }
}
</style>