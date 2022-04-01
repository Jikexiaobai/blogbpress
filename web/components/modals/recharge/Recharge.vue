<template>
    <div class="pay_box center opacity" :class="[isTrue && 'is_back_show']">
        <div class="pay_container">
            <div class="pay_title">
                <div class="pay_title_l">
                    <span>充值</span>
                </div>
                <div class="pay_title_r">
                    <a-icon  type="close" @click="cancel"/>
                </div>
            </div>

            <div class="pay_content" v-if="mode == 1 || mode == 2">
                <div class="pay_content_desc">
                    请输入充值金额：
                </div>
                <ul class="pay_content_xx">
                    <li @click="picked(item,index)" :class="index == moneyActiveKey ? 'picked':''" v-for="(item,index) in moneyList" :key="index">
                        <div class="cz_item">
                            <span>{{item}}</span>
                            <b>{{base.currencySymbol}}</b>
                        </div>
                    </li>

                    <li @click="picked(null,6)" :class="moneyActiveKey == 6 ? 'picked':''">
                        <div v-if="moneyActiveKey != 6" class="cz_item cz_custom">
                            自定义
                        </div>
                        <div v-else class="cz_item">
                            <a-input-number class="cz_input" v-model="money" :style="{ width: '100%' }"  :min="5" :max="9999" :precision="2"/>
                        </div>
                    </li>
                </ul>
            </div>

            <div class="pay_content" v-if="mode == 3">
                <div class="pay_content_desc">
                    请输入卡密：
                </div>
                <div class="crad-key">
                    <a-input-password v-model="cardKey" size="large" placeholder="请输入密钥" />
                </div>
            </div>

            <div class="pay_content" v-if="mode == 4">
                <div class="pay_content_desc">
                    请输入转账信息：
                </div>
                <div class="crad-key">
                    
                    <a-input v-model="name" size="large" placeholder="请输入账户名称" />
                    <a-select v-model="type" class="crad-key" size="large" style="width: 100%" placeholder="请选择转账方式">
                        <a-select-option :value="1">
                            支付宝
                        </a-select-option>
                        <a-select-option :value="2">
                            微信
                        </a-select-option>
                    </a-select>
                    <a-input v-model="number" class="crad-key" size="large" placeholder="请输入转账单号" />
                    <a-input-number placeholder="请输入充值金额" class="crad-key" size="large"
                     v-model="money"
                    :style="{ width: '100%' }" 
                     :min="5" :max="9999" 
                    :precision="2"/>
                </div>
            </div>

            <div class="pay_number"  v-if="mode != 3 && mode != 4">
                <i>{{base.currencySymbol}}</i>
                <span>{{money}}</span>
            </div>

            <div class="pay_options">
                <ul>
                    <li v-for="(item) in rechargeType" :key="item"  :class="mode == item ? 'picked':''">
                        <a-button v-if="item == 1" @click="payPicked(item)" class="pay_button" icon="alipay-circle">
                            支付宝
                        </a-button>
                        <a-button v-if="item == 2" @click="payPicked(item)" class="pay_button" icon="alipay-circle">
                            微信
                        </a-button>
                        <a-button v-if="item == 3" @click="payPicked(item)" class="pay_button" icon="alipay-circle">
                            卡密
                        </a-button>
                        <a-button v-if="item == 4" @click="payPicked(item)" class="pay_button" icon="alipay-circle">
                            人工充值
                        </a-button>
                    </li>
                </ul>
            </div>

            <div class="pay_desc" v-if="mode != 3 && mode != 4">
                即时到账
            </div>
            <div class="pay_desc" v-if="mode == 3 || mode == 4">
                审核到账
            </div>

            <div class="pay_submit">
                <a-button v-if="isGoPay == 1" @click="czPay" size="large" type="primary" block>
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

import api from "@/api/index"
import { mapState } from "vuex"
export default {
    data() {
        return {
            isGoPay:1,

            rechargeType:[],

            moneyActiveKey:0,
            moneyList:[20,50,100,500,1000],

            mode:1,
            money: 0,
            cardKey: "",
            name: "",
            type: 1,
            number: "",

            

            isTrue: false,
            state: null // 准备（prepare） 确定（ ascertain） 取消（cancel）
        };
    },
    computed:{
        ...mapState(["base","pay"])
    },
    methods: {
        async confirm(
        ) {
            this.money = this.moneyList[this.moneyActiveKey]

            this.rechargeType = this.pay.recharge
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
        picked(i,index){
            if(i != null){
                this.money = i
            }
            this.moneyActiveKey = index
        },
        payPicked(i){
            this.money = 0
            this.mode = i
        },
        async czPay(){
            if (this.isGoPay == 2 || this.isGoPay == 3) {
                return false
            }
            
            if ((this.mode == 1 || this.mode == 2) && this.money == 0) {
                this.$message.error(
                    "请选择或输入充值金额",
                    3
                )
                return
            }

            if (this.mode == 3 && (this.cardKey == "")) {
                this.$message.error(
                    "请输入卡密",
                    3
                )
                return
            }

            if (this.mode == 4 && (this.type == "" || this.name == "" || this.number == "")) {
                this.$message.error(
                    "请输入转账信息",
                    3
                )
                return
            }


            this.isGoPay = 2
            const formData = {
                mode: this.mode,
                money:this.money,
                cardKey:this.cardKey,
                name:this.name,
                type:this.type,
                number:this.number,
            }

            const res = await this.$axios.post(api.postRechargeCreate,formData)

            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
            
            this.isGoPay = 3
           
            if (this.mode == 1 || this.mode == 2) {
                const qrRes = await this.$axios.post(api.postRechargePay,{code:res.data.code})
                if (qrRes.code != 1) {
                    this.$message.error(
                        qrRes.message,
                        3
                    )
                    return
                }
                this.ascertain()
                this.throttl(5000)
                this.$RechargeQr("用户充值",formData.money,this.mode,qrRes.data.info.qrCode,qrRes.data.info.code).then((res)=>{
                    if (res != false) {
                        this.throttl(5000)
                        this.ascertain()
                    }
                }).catch((err)=>{
                    this.cancel()
                })
            }else{
                this.ascertain()
                this.throttl(5000)
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
            this.mode = 1
            this.money = 0
            this.cardKey = ""
            this.name = ""
            this.type = 1
            this.number = ""
            this.state.state = "cancel"
            this.close()
        },
        ascertain(){
            this.mode = 1
            this.money = 0
            this.cardKey = ""
            this.name = ""
            this.type = 1
            this.number = ""
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
            width: 25rem;
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
            .pay_content{
                display: flex;
                flex-flow: column;
                justify-content: center;
                margin-bottom: 10px;
                align-items: center;
                padding: 10px;
                .pay_content_desc{
                    font-size: 18px;
                    font-weight: bold;
                    text-align: center;
                    margin-top: 20px;
                    margin-bottom: 10px;
                }
                .pay_content_xx{
                    display: flex;
                    flex-flow: wrap;
                    width: 100%;
                    li{
                        width: 33.333%;
                        .cz_item{
                            padding: 10px;
                            font-size: 10px;
                            margin: 5px;
                            border: 1px solid #ddd;
                            text-align: center;
                            line-height: 1;
                            display: flex;
                            justify-content: center;
                            align-items: flex-end;
                            border-radius: 3px;
                            cursor: pointer;
                            height: 50px;
                            span{
                                font-size: 30px;
                            }
                        }
                        .cz_input{
                            border: 0;
                            font-size: 18px;
                            padding: 0;
                            color: #f16b6f
                        }
                        .cz_custom{
                            font-size: 15px;
                            align-items: center;
                        }
                    }
                    li:hover{
                        background-color: #f7f8fa;
                    }
                    .picked{
                        .cz_item{
                            border-color: #f16b6f;
                            color: #f16b6f;
                        }
                    }
                }
                .crad-key{
                    margin-top: 10px;
                }
            }

            .pay_number{
                font-size: 40px;
                color: green;
                display: flex;
                justify-content: center;
                padding-right: 10px;
                i{
                    font-size: 21px;
                    font-style: normal;
                    top: 6px;
                    position: relative;
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
            .pay_desc{
                display: flex;
                justify-content: center;
                align-items: center;
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
        .login_box{
            .login_container{
                margin: 0 20px;   
                .login_close{
                    position: absolute;
                    top: 0;
                    right: 0;
                    font-size: 20px;
                    color: #000;
                    padding: 10px;    
                }
            }
        }
    }
</style>
