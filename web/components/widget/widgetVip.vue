<template>
    <div class="widget-box">
        <div class="warper" :style="{ width: '1500px' }">
            <div class="widget-title" v-if="info.showTitle == 2">
                <h1 class="title">{{info.title}}</h1>
                <h2>升级VIP会员享受海量资源免费下载</h2>
            </div>
            <div class="widget-content">
                <a-row :gutter="[{md:20}]" type="flex" justify="space-between">
                    <a-col v-for="(item,index) in info.list" :key="index"  :span="(24 / info.list.length)">
                        <div class="item">
                            <div class="item-header">
                                <h1 class="item-price">
                                    <span>￥{{item.price.toFixed(2)}}</span>
                                </h1>
                                <div class="item-title-box">
                                    <a-space>
                                        <img class="item-icon" :src="item.icon" :alt="item.title">
                                        <span class="item-title">{{item.title}}</span>
                                        <span class="item-title" v-if="item.day != 0">/</span>
                                        <span class="item-title" v-if="item.day != 0">{{item.day}}天</span>
                                    </a-space>
                                </div>
                            </div>
                            <ul class="item-content">
                                <li>
                                    {{`享受${item.title}所有特权`}}
                                </li>
                                <li>
                                    享受VIP专属标志
                                </li>
                                <li>
                                    享受{{item.discount}}%专属折扣
                                </li>
                            </ul>
                            <div class="item-push">
                                <a-button size="large" @click="goVip" type="primary" block>
                                    立即加入
                                </a-button>
                            </div>
                        </div>
                    </a-col>
                </a-row>
            </div>
        </div>
    </div>
</template>

<style lang="less" scoped>
.widget-box{
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 20px 0;
    .widget-title{
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        .title{
            font-size: 1.5rem;
            letter-spacing: .8px;
            font-weight: 700;
            position: relative;
        }
   
    }
    .widget-content{
        margin-top: 20px;
        .item{
            background-color: #fff;
            padding: 20px;
            border-radius: 5px;
            .item-header{
                display: flex;
                justify-content: center;
                align-items: center;
                flex-direction: column;
                padding-bottom: 30px;
                .item-price{
                    color: #ff5b5b!important;
                    // font-weight: 500;
                    font-size: 2.5rem;
                    line-height: 1.2;
                    font-weight: bolder!important;
                }
                .item-title-box{
                    color: #ff5b5b!important;
                    margin-top: 10px;
                    .item-icon{
                        width: 24px;
                        height: 24px;
                    }
                    .item-title{
                        font-size: 18px;
                        font-weight: bolder!important;
                    }
                }
            }
            .item-content{
                display: flex;
                justify-content: center;
                align-items: center;
                flex-direction: column;
                li{
                    padding: .5rem 0;
                    color: #6c757d!important;
                }
            }
            .item-push{
                margin-top: 30px;
                padding: 0 40px;
            }
        }
    }
}
</style>


<script>
import { mapState } from "vuex"
export default {
    props: {
        info:{
            type: Object,
            default: {}
        },
    },
    filters: {
       resetService(v){
           let notPayServices = JSON.parse(v)
           let str = ""
           notPayServices.forEach((item) => {
                if (item == 1) {
                    str +=  " 认证"
                }
                if (item == 2) {
                    str +=  "提现"
                }
           });
           return str
       }
    },
    data(){
        return{
            
        }
    },
    computed:{
        ...mapState(["design"]),
        ...mapState("user",["token"]),
    },
    methods: {
        goVip(e) {
            if (this.token == null) {
                this.$Auth("login","登录","快速登录")
                return
            }
            this.$router.push("/account/vip")
        },
    }, 
}
</script>