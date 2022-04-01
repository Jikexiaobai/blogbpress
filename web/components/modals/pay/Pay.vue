<template>
    <div class="pay_box center opacity" :class="[isTrue && 'is_back_show']">
        <div class="pay_container">
            <div class="pay_title">
                <div class="pay_title_l">
                    <span>{{title}}</span>
                </div>
                <div class="pay_title_r">
                    <a-icon  type="close" @click="cancel"/>
                </div>
            </div>

            <div class="pay_content_dese">
                <span>
                    支付的金额
                </span>
            </div>

            <div class="pay_price">
                <i>$</i>
                <span>{{product.orderMoney}}</span>
            </div>

            <div class="pay_my_yu_e">
                <span>您当前的余额为：￥{{balance}}</span>
            </div>

            <div class="pay_options">
                <ul>
                    <li v-for="(item) in payMode" :key="item" :class="payActiveKey == item ? 'picked':''">
                        <a-button v-if="item == 1" @click="payPicked(item)" class="pay_button" icon="alipay-circle">
                            支付宝
                        </a-button>
                        <a-button v-if="item == 2" @click="payPicked(item)" class="pay_button"  icon="wechat">
                            微信
                        </a-button>
                        <a-button v-if="item == 3" @click="payPicked(item)" class="pay_button"  icon="wallet">
                            余额
                        </a-button>
                    </li>
                </ul>
            </div>

            <div class="pay_submit">
                <a-button v-if="isGoPay == 1" @click="payGo" size="large" type="primary" block>
                    支付
                </a-button>

                <a-button v-if="isGoPay == 2" type="primary" size="large" loading block>
                    创建订单
                </a-button>

                <a-button v-if="isGoPay == 3" type="primary" size="large" loading block>
                    正在支付
                </a-button>
            </div>

        </div>
    </div>
</template>

<script>
import { mapState } from "vuex"
import api from "@/api/index"
export default {
    data() {
        return {
            isGoPay:1,
            payMode:[],
            payActiveKey:0,

            title:null,
            product:{
                orderMoney:0,
            },
            balance:null, // 我的余额

            // 登录输入框
            isTrue: false,
            state: null // 准备（prepare） 确定（ ascertain） 取消（cancel）
        };
    },
    computed:{
        ...mapState(["pay"])
    },
    methods: {
        async confirm(
            title,
            product
        ) {
            this.title = title || "支付"
            this.product = product || null
            // 获取用户余额
            const {code,data,message} = await this.$axios.get(api.getAccountBalance)
            if(code != 1){
                this.$message.error(
                    message,
                    3
                )
                return
            }
            this.balance = data.balance
            this.payMode = this.pay.payMode


            this.open();
            return new Promise((resolve, reject) => {
                const target = { state: 'prepare' };
                let res = new Proxy(target, {
                    set(event, key, value) {
                        if (value === 'ascertain') {
                            resolve(true);
                        } else {
                            reject(false);
                        }
                        return true;
                    }
                });
                this.state = res;
            });
        },

        payPicked(i){
            this.payActiveKey = i
            
        },
        async payGo(){
            if (this.isGoPay == 2 || this.isGoPay == 3) {
                return false
            }
            
            if (this.product.orderMoney == 0 || this.product.orderMoney == null) {
                this.$message.error(
                    "请选择或输入充值金额",
                    3
                )
                return
            }

            if (this.payActiveKey == 0) {
                this.$message.error(
                    "请选择充值方式",
                    3
                )
                return
            }
            if (this.payActiveKey == 3) {
                if ((this.balance * 100) < (this.product.orderMoney * 100)) {
                    this.$message.error(
                        "余额不足",
                            3
                        )
                    return
                }
            }

           

            this.isGoPay = 2
            let formData = {
                payMethod:this.payActiveKey,
            }
            formData = Object.assign(this.product,formData)
            
            const res = await this.$axios.post(api.postOrderCreate,formData)

            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
            this.isGoPay = 3

            const qrRes = await this.$axios.post(api.postOrderPay,{orderNum:res.data.orderNum})
            if (qrRes.code != 1) {
                this.$message.error(
                    qrRes.message,
                    3
                )
                return
            }
            if(qrRes.data.info.isPay && qrRes.data.info.payMethod == 3){
                this.throttl(5000)
                this.ascertain()
                return
            }
            this.close()
            this.throttl(5000)
            if (this.payActiveKey != 3) {
                 this.$QrCode(this.title,formData.orderMoney,this.payActiveKey,qrRes.data.info.orderNum,qrRes.data.info.orderNum).then((res)=>{
                    if (res != false) {
                        this.throttl(5000)
                        this.ascertain()
                    }
                }).catch((err)=>{
                    this.cancel()
                })
            }
        },
        throttl(dey){ 
            let timeout;
            clearTimeout(timeout);  // 每次触发时先清除上一次的定时器,然后重新计时
            timeout = setTimeout(()=>{
                this.isGoPay = 1
            }, dey);  // 指定 xx ms 后触发真正想进行的操作 handler
        },
        cancel(){
            this.state.state = "cancel"
            this.close()
        },
        ascertain(){
            this.state.state = "ascertain"
            this.close()
        },
        open() {
            this.isTrue = true;
        },
        close() {
            this.isTrue = false;
        },
        
    }
};
</script>

<style lang="less" scoped>
    .pay_box {
        user-select: none;
        pointer-events: none;
        z-index: 20;
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        width: 100%;
        height: 100%;
        opacity: 0;
        visibility: hidden;
        transform: perspective(1px) scale(1.1);
        transition: visibility 0s linear .15s,opacity .15s 0s,transform .15s;
        display: flex;
        align-items: center;
        justify-content: center;
        .pay_container{
            background-color: #fff;
            width: 21rem;
            margin: 0 auto;
            position: relative;
            background-image: url("/img/login.png");
            background-repeat: no-repeat;
            background-size: 100%;
            margin-top: -9%;
            .pay_title{
                font-size: 13px;
                display: flex;
                justify-content: space-between;
                align-items: center;
                padding: 10px 20px;
                .pay_title_l{
                    display: block;
                    align-items: center;
                    width: 80%;
                    overflow: hidden;
                    white-space: nowrap;
                    text-overflow: ellipsis;
                }
            }

            .pay_content_dese{
                text-align: center;
                margin-top: 20px;
                font-size: 20px;
            }

            .pay_price{
                padding: 20px 0;
                font-size: 40px;
                color: green;
                display: flex;
                justify-content: center;
                i{
                    font-size: 21px;
                    font-style: normal;
                    top: 6px;
                    position: relative;
                }
            }
            .pay_my_yu_e{
                display: flex;
                justify-content: center;
                span{
                    display: inline-block;
                    font-size: 12px;
                    text-shadow: 0 0 1px #fff;
                    box-shadow: inset 0 0 10px #e0e0e0;
                    padding: 5px 10px;
                }
            }

            .pay_options{
                padding: 10px;
                height: 68px;
                ul{
                    display: flex;
                    flex-flow: nowrap;
                    align-items: center;
                    justify-content: center;
                    li{
                        width: 100%;
                        display: flex;
                        .pay_button{
                            display: flex;
                            align-items: center;
                            margin: 5px;
                            justify-content: center;
                            border: 1px solid #eee;
                            height: 38px;
                            font-size: 12px;
                            line-height: 1;
                            background: 0 0;
                            color: #333;
                            width: 100%;
                            padding: 0;
                        }
                    }
                    .picked{
                        .pay_button{
                            border-color: #f16b6f;
                            color: #f16b6f;
                        }
                    }
                }
            }

            .pay_submit{
                padding: 10px;
            }
        }
        
    }
    .is_back_show {
        opacity: 1 !important;
        background: rgba(42, 44, 48, 0.7);
        pointer-events: auto !important;
        visibility: visible;
        transform: perspective(1px) scale(1);
        transition: visibility 0s linear 0s,opacity .15s 0s,transform .15s;
    }

    @media only screen and (max-width: 768px) {
        .pay_box{
            // .pay_container{
            //     margin: 0 20px;   
            //     .login_close{
            //         position: absolute;
            //         top: 0;
            //         right: 0;
            //         font-size: 20px;
            //         color: #000;
            //         padding: 10px;    
            //     }
            // }
        }
    }
</style>