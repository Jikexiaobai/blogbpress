<template>
    <div class="pay_qr_box center opacity" :class="[isTrue && 'is_back_show']">
        <div class="pay_qr_container">
            <div class="pay_title">
                <div class="pay_title_l">
                    <span>{{title}}</span>
                </div>
                <div class="pay_title_r">
                    <a-icon  type="close" @click="cancel"/>
                </div>
            </div>

            <div  class="pay_qr_content">

                <div class="qrcode-img">
                    <img v-if="qrCode != null" :src="qrCode">
                    <a-icon v-if="qrCode == null" type="loading" class="loading"/>
                </div>

                <div class="qrcode-money">
                    <span>￥{{money}}</span>
                </div>

                <div class="qrcode-desc">
                   {{mode | restDesc}}
                </div>

                <div class="qrcode-time">
                    「 <span v-html="timesec"></span> 」
                </div>

            </div>
        </div>
    </div>
</template>

<script>
import api from "@/api/index"
import Qrious from 'qrious'
const PAYTYPE = {
    ALY:1,
    WX:2,
    BL:3
}

export default {
    filters: {
        restDesc(value) {
            switch (value) {
                case PAYTYPE.ALY:
                    return "请打开手机使用支付宝扫码支付"
                case PAYTYPE.WX:
                    return "请打开手机使用微信扫码支付"
            }
        },
    },
    data() {
        return {
            title: "", // 标题
            money: 0,
            mode: null,
            qrCode: null,
            code:null,


            // 轮询定时器
            endTime:"",
            currTime:"",
            timeOut:300,
            timesec:'',
            sTime:'',
            success:'',
            checkTime:'',
            // 登录输入框
            isTrue: false,
            state: null // 准备（prepare） 确定（ ascertain） 取消（cancel）
        };
    },
    watch:{
        isTrue(val){
            if(val){
                this.sTime = ''
      
                this.checkTime = ''
            }else{
                this.sTime = null
               
                this.checkTime = null
            }
        },
    },
    methods: {
        async confirm(
            title,
            money,
            mode,
            qrCode,
            code,
        ) {
            this.open();

           
            this.title = title || this.title;
            this.money = money || this.money;
            this.mode = mode || this.mode;
            this.qrCode = qrCode || this.qrCode;
            this.code = code || this.code;
            this.checkbill()

            this.currTime = parseInt(Date.parse(new Date())/1000);
            this.endTime = parseInt(this.currTime + this.timeOut);
            this.setTime()
    
            if (this.mode != 3) {
                    var qr = new Qrious({
                        value: qrCode,
                        size:200,
                        level:'L'
                });
                this.qrCode = qr.toDataURL('image/jpeg')
            }

            
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
        

        async checkbill(){
            if(this.sTime === null || this.checkTime === null || this.isTrue == false){
                this.checkTime = null
                return
            }
            
            const formData = {
                code:this.code,
            }
            const res = await this.$axios.post(api.postRechargeCheckStatus,formData)
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
            }
            if(res.data.status){
                this.checkTime = null;
                this.$message.success("支付成功", 5)
                this.ascertain()
            }else{
                this.checkTime = setTimeout(()=>{
                    this.checkbill()
                },1000)
            }

        },

        setTime(){

            if(this.isTrue == false) return
            let diff_time = parseInt(this.endTime-this.currTime);
            let m = Math.floor((diff_time / 60 % 60));
            let s = Math.floor((diff_time % 60));
            this.timesec = (m > 0 ? m + '<b>分</b>' : '') + s + '<b>秒</b>';
            if(diff_time > 0){
                this.sTime = setTimeout(()=>{
                    this.endTime = this.endTime - 1;
                    this.setTime()
                },1000)
            }else{
                this.sTime = null
                this.$message.error("支付失败", 5)
                this.ascertain()
            }
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
    .pay_qr_box {
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
        .pay_qr_container{
            background-color: #fff;
            width: 18rem;
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
            .pay_qr_content{
                padding: 20px;
                display: flex;
                justify-content: center;
                flex-flow: column;
                align-items: center;
                padding: 20px;
                .qrcode-img{
                    width: 186px;
                    height: 186px;
                   
                    background-color: #f3f3f3;
                    background-repeat: no-repeat;
                    background-position: center;
                    background-size: 18px;
                    padding: 10px;
                    border: 1px solid #ddd;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    .loading{
                        font-size: 60px;
                    }
                    img{
                        max-width: 100%;
                        height: auto;
                        object-fit: cover;
                        vertical-align: bottom;
                        image-rendering: -webkit-optimize-contrast;
                    }
                }
                .qrcode-money{
                    margin: 20px 0;
                    span{
                        display: inline-block;
                        color: #fff;
                        background-color: green;
                        padding: 5px 14px;
                        border-radius: 3px;
                        font-size: 15px;
                        line-height: 1;
                    }
                }
                .qrcode-desc {
                    font-size: 15px;
                    display: flex;
                    align-items: center;
                    color: #777;
                }

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