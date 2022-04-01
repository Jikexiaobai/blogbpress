<template>
    <div class="sidebar-box">
        <div class="down-price">
            <a-space  v-if="info.downMode == PREMSTYPE.FF">
                <span style="color: #dadada; font-size: 15px; margin: 0;vertical-align:middle;">售价</span>
                <span class="price">{{info.price}}</span>
                <sup>{{base.currencySymbol}}</sup>
            </a-space>
            <a-space  v-if="info.downMode == PREMSTYPE.PL">
                <span class="price">评论获取</span>
            </a-space>
            <a-space  v-if="info.downMode == PREMSTYPE.DL">
                <span class="price">登录获取</span>
            </a-space>
            <a-space  v-if="info.downMode == 0">
                <span class="price">免费下载</span>
            </a-space>
        </div>
        <div v-if="vipPrice != 0" class="down-price">
            <a-space >
                <span v-if="info.downMode == PREMSTYPE.FF" 
                    style="color: red; font-size: 15px; margin: 0;vertical-align:middle;">
                    会员折扣价
                </span>
                <span style="color: red;" class="price">{{vipPrice}}</span>
                <sup v-if="info.downMode == PREMSTYPE.FF" style="color: red;">{{base.currencySymbol}}</sup>
            </a-space>
        </div>
        <ul class="down-vpi-price">
            <li>
                <a-space >
                    <a-icon type="dollar" />
                    <span>获取方式: {{info.downMode | restPrems}}</span>
                </a-space>
            </li>
            <div  v-if="info.purpose.length > 0">
                <li  v-for="(item,index) in info.purpose" :key="index">
                    <a-space >
                        <a-icon type="dollar" />
                        <span>{{item.key}} : {{item.val}}</span>
                    </a-space>
                </li>
            </div>
            
        </ul>
        <div v-if="!info.isDown">
            <a-button v-if="info.downMode == PREMSTYPE.FF && !info.isView" @click="goPay" size="large" type="primary" block>
                支付购买
            </a-button>
        </div>
        <ul v-if="info.isDown" class="down-url">
            <!-- <li >
                <a-button type="primary">
                    点击下载
                </a-button>
                <span>{{item.val}}</span>
            </li> -->
            <li v-for="(item,index) in info.downUrl" :key="index">
                <a-button-group>
                    <a-button @click="openUrl(item.key)" type="primary">{{item.title}}</a-button>
                    <a-button 
                        v-clipboard:copy="item.val" 
                        v-clipboard:success="onCopy"
                        type="primary">
                            提取码: {{item.val}}
                        </a-button>
                </a-button-group>
            </li>
            
        </ul>
        <ul v-if="info.attribute.length > 0" class="down-info">
            <li v-for="(item,index) in info.attribute" :key="index">
                <span>{{item.key}}</span>
                <span>{{item.val}}</span>
            </li>
        </ul>
    </div>
</template>

<script>
import { mapState } from "vuex"
import {ORDERTYPE} from "@/shared/order"
const PREMSTYPE = {
    FF:1, // 付费
    PL:2, // 评论
    DL:3, // 登录
}
export default {
    props:{ 
        info: {
            type: Object, //指定传入的类型
            default: {} //这样可以指定默认的值
        },
        module: {
            type: String, //指定传入的类型
            default: "" //这样可以指定默认的值
        },
    },
    filters: {
        restPrems(value) {
            switch (value) {
                case PREMSTYPE.FF:
                    return "付费购买"
                case PREMSTYPE.PL:
                    return "评论获取"
                case PREMSTYPE.DL:
                    return "登录获取"
                default:
                    return "开源使用"
            }
        }
    },
    computed:{
        ...mapState(["base"]),
        ...mapState("user",["token","userInfo"]),
    },
    data(){
        return{
            PREMSTYPE,
            vipPrice:0,
        }
    },
    mounted() {
       
        if (this.info.downMode == PREMSTYPE.FF) {
            let discount = 0
            if (this.userInfo.vip != null) {
                 discount = this.userInfo.vip.discount
                 
            }
            if (discount != 0) {
                let price = this.info.price
                
                let discountPrice = price - (price * discount)
                this.vipPrice = discountPrice
            }
        }
    },
    methods: {
        goPay(){
            if (this.token != null) {
                const product = {
                    detailId:this.info.id,
                    detailModule:this.module,
                    orderMoney:this.info.price,
                    orderType: ORDERTYPE.BUYZY,
                }
                if (this.vipPrice != 0) {
                    product.orderMoney = this.vipPrice
                }
                this.$Pay("购买内容",product).then(async (res)=>{
                    if (res != false) {
                        this.$emit('upadteView')
                        this.$message.success(
                            "成功购买",
                            3
                        )
                        // this.upadteView()
                    }
                }).catch((err)=>{
                    console.log(err)
                    // this.createForm.cover = undefined
                })
                
            } else {
                this.$Auth("login","登录","快速登录")
            }
        },
        openUrl(e){
            window.open(e); 
        },
        onCopy(){

            this.$message.success(
                "已复制提取码",
                3
            )
        }
    }
}
</script>

<style lang="less" scoped>
.sidebar-box{
    padding: 20px;
    background: white;
    margin-bottom: 10px;

    .down-price{
        padding: 10px 20px;
        border-radius: 4px;
        margin-top: 0px;
        margin-bottom: 10px;
        display: flex;
        justify-content: center;
        align-items: center;
        z-index: 11;
        -webkit-box-shadow: 0 4px 30px rgba(0, 0, 0, 0.1);
        box-shadow: 0 4px 30px rgba(0, 0, 0, 0.1);
        background: #FFF;
        .price{
            font-size: 28px;
            font-weight: bold;
        }
    }
    .down-url{
        li{
            padding: 10px 0;
        }
    }
    .down-vpi-price{
        li{
            padding: 10px 0;
            font-size: 14px;
            cursor: pointer;
            color: #8c8c8c;
        }
        li + li{
                border-top: 1px solid #e6e6e6;
        }
    }
    .code{
        margin-bottom: 10px;
    }
    .down-info{
        margin-top: 10px;
        padding: 10px;
        background-color: #f3f7ff;
        border-radius: 4px;
        li{
            display: flex;
            justify-content: space-between;
            align-items: center;
            color: #8c8c8c;
            font-size: 13px;
            line-height: 20px;
            padding: 5px 0;
            list-style: none;
        }
    }
}
</style>