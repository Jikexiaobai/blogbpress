<template>
    <div class="pay_box center opacity" :class="[isTrue && 'is_back_show']">
        <div class="pay_container">
            <div class="pay_title">
                <div class="pay_title_l">
                    <Avatar :src="avatar+'@w60_h60'" shape="circle" :size="28" />
                    <span>{{nickName}}{{title}}</span>
                </div>
                <div class="pay_title_r">
                    <a-icon  type="close" @click="close"/>
                </div>
            </div>

            <div class="pay_content">
                <div class="pay_content_desc">
                    给你喜欢的用户充电加油
                </div>
                <ul class="pay_content_xx">
                    <li @click="picked(item,index)" :class="index == priceActiveKey ? 'picked':''" v-for="(item,index) in priceList" :key="index">
                        <div class="cz_item">
                            <span>{{item}}</span>
                            <b>￥</b>
                        </div>
                    </li>

                    <li @click="picked(null,6)" :class="priceActiveKey == 6 ? 'picked':''">
                        <div v-if="priceActiveKey != 6" class="cz_item cz_custom">
                            自定义
                        </div>
                        <div v-else class="cz_item">
                            <a-input-number class="cz_input" v-model="price" :style="{ width: '100%' }"  :min="5" :max="9999" :precision="2"/>
                        </div>
                    </li>
                </ul>
            </div>

            <div class="pay_number">
                <i>￥</i>
                <span>{{price}}</span>
            </div>

            <div class="pay_my_yu_e">
                <span>您当前的余额为：￥{{balance}}</span>
            </div>

            <div class="pay_options">
                <ul>
                    <li :class="payActiveKey == 1 ? 'picked':''">
                        <a-button @click="payPicked(1)" class="pay_button" icon="alipay-circle">
                            支付宝
                        </a-button>
                    </li>
                    <li :class="payActiveKey == 2 ? 'picked':''">
                        <a-button  @click="payPicked(2)" class="pay_button"  icon="wechat">
                            微信
                        </a-button>
                    </li>
                    <li :class="payActiveKey == 3 ? 'picked':''">
                        <a-button @click="payPicked(3)" class="pay_button"  icon="wallet">
                            余额
                        </a-button>
                    </li>
                </ul>
            </div>

            <div class="pay_submit">
                <a-button v-if="isGoPay == 1" @click="goPay" size="large" type="primary" block>
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
import Avatar from "@/components/avatar/avatar"
import api from "@/api/index"

// import wsConnection from "@/service/websocket" 
// import router from '../../router'
export default {
    components:{
        Avatar,
    },
    data() {
        return {
            isGoPay:1,

            priceActiveKey:0,
            priceList:[20,50,100,500,1000],


            payActiveKey:null,
            payType: null,
            title:null,
            avatar:"",
            nickName:"",
            product:null,
            price:0,
            balance:null, // 我的余额

            // 登录输入框
            isTrue: false,
            state: null // 准备（prepare） 确定（ ascertain） 取消（cancel）
        };
    },
    methods: {

        async confirm(
            title,
            product,
            avatar,
            nickName,
        ) {
            this.price = this.priceList[this.priceActiveKey]
            this.title = title || "充电"
            this.nickName = nickName || ""
            this.avatar = avatar || ""
            this.product = product || null

           const res = await this.$axios.get(api.getAccountBalance)
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
            }
            this.balance = res.data.balance

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
                this.price = i
            }
            this.priceActiveKey = index
        },
        payPicked(i){
            this.payActiveKey = i
            this.payType = i
        },
        async goPay(){
            //  = false
            if (!this.isGoPay) {
                return false
            }
            
            if (this.price == 0 || this.price == null) {
                this.$message.error(
                    "请选择或输入支付金额",
                    3
                )
                return
            }

            if (this.payType == null) {
                this.$message.error(
                    "请选择支付方式",
                    3
                )
                return
            }
            if (this.payType == 3) {
                if ((this.balance * 100) < (this.price * 100)) {
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
                orderMoney:this.price,
                authorId:this.product.detailId
            }
            formData = Object.assign(this.product,formData)
            
            const res = await this.$axios.post(api.postOrderCreate,formData)
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
            }
            if (res.data.orderNum != null || res.data.orderNum != "") {
                this.isGoPay = 3
                const qrRes = await this.$axios.post(api.postOrderPay,{orderNum:res.data.orderNum})
                if (qrRes.code != 1) {
                    this.$message.error(
                        qrRes.message,
                        3
                    )
                    this.cancel()
                    return
                }
                if(qrRes.data.info.isPay && qrRes.data.info.payMethod == 3){
                    this.throttl(5000)
                    this.ascertain()
                    return
                }

                if (this.payActiveKey != 3) {
                    this.$QrCode(this.title,formData.orderMoney,this.payActiveKey,qrRes.data.info.orderNum,qrRes.data.info.orderNum).then((res)=>{
                        if (res != false) {
                            // this.throttl(5000)
                            this.ascertain()
                        }
                    }).catch((err)=>{
                        this.cancel()
                    })
                }
            }else{
                this.cancel()
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
                    display: flex;
                    align-items: center;
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
