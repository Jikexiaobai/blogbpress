<template>
    <div class="widget-box">
        <div class="widget-content" :style="{ width: '1500px' }">
            <div class="widget-info">
                <h1>精彩的人生就是无止境的折腾</h1>
                <div class="widget-info-btn">
                    <button @click="joinQQqun" class="btn">系统交流群：488741878</button>
                    <!-- <span></span>
                    <button class="btn2">申请签约接单</button> -->
                </div>
            </div>
            <div class="widget-am0 widget-am"></div>
            <div class="widget-am1 widget-am"></div>
            <div class="widget-am2 widget-am"></div>
            <div class="widget-am3 widget-am"></div>
            <div class="widget-am4 widget-am"></div>
            <div class="widget-am5 widget-am"></div>
            <div class="widget-am6 widget-am"></div>
            <div class="widget-am7 widget-am"></div>
            <div class="widget-am8 widget-am"></div>
            <div class="widget-am9 widget-am"></div>
        </div>
    </div>
</template>

<style lang="less" scoped>
.widget-box{
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 20px 0;
    width: 100%;
    height: 624px;
    margin-top: -75px;
    background-image: url("/img/tmp/header-bg.svg");
    background-size: auto 622px;
    background-position: 50% 0;
    background-repeat: no-repeat;
    overflow: hidden;
    .widget-content{
        position: relative;
        margin: auto;
        @keyframes toy {
            0% {
                -webkit-transform: scale(1);
                transform: scale(1);
            }
            50% {
                -webkit-transform: scale(.7);
                transform: scale(.7);
            }
            100% {
                -webkit-transform: scale(1);
                transform: scale(1);
            }
        }
        @keyframes upAnimation {
            0%, 100%, 20%, 50%, 80% {
                transition-timing-function: cubic-bezier(0.215,.61,.355,1);
                transform: translate3d(0,0,0);
            }
            40%, 43% {
                transition-timing-function: cubic-bezier(0.755,0.50,0.855,0.060);
                transform: translate3d(0,-30px,0);
            }
            70% {
                transition-timing-function: cubic-bezier(0.755,0.050,0.855,0.060);
                transform: translate3d(0,-15px,0);
            }
            90% {
                transform: translate3d(0,-4px,0);
            }
        }
        @keyframes zhuan {
            0% {
                -webkit-transform: rotate(
            0deg);
            }
            25% {
                -webkit-transform: rotate(
            90deg);
            }
            50% {
                -webkit-transform: rotate(
            180deg);
            }
            75% {
                -webkit-transform: rotate(
            270deg);
            }
            100% {
                -webkit-transform: rotate(
            360deg);
            }
        }
        @keyframes aroundAnimation {
            0%, 100%, 20%, 50%, 80% {
                transition-timing-function: cubic-bezier(0.215,.61,.355,1);
                transform: translate3d(0,0,0);
            }
            40%, 43% {
                transition-timing-function: cubic-bezier(0.755,0.50,0.855,0.060);
                transform: translate3d(-20px,0,0);
            }
            70% {
                transition-timing-function: cubic-bezier(0.755,0.050,0.855,0.060);
                transform: translate3d(-10px,0px,0);
            }
            90% {
                transform: translate3d(20px,0,0);
            }
        }
        @keyframes toy2 {
            0% {
                -webkit-transform: scale(.7);
                transform: scale(.7);
            }
            50% {
                -webkit-transform: scale(1);
                transform: scale(1);
            }
            100% {
                -webkit-transform: scale(.7);
                transform: scale(.7);
            }
        }
        .widget-am0 {
            left: 246px;
            top: 120px;
            width: 60px;
            height: 60px;
            background-image: url(/img/tmp/toy5.svg);
            -webkit-animation: toy 3s infinite;
            animation: toy 3s infinite;
        }
        .widget-am1 {
            left: -10px;
            top: 440px;
            width: 32px;
            height: 32px;
            background-image: url(/img/tmp/toy1.svg);
            // -webkit-animation: toy 3s infinite;
            animation: toy 3s infinite;
        }
        .widget-am2 {
            left: 100px;
            top: 302px;
            width: 29px;
            height: 29px;
            // -webkit-animation: upAnimation 3s infinite;
            animation: upAnimation 3s infinite;
            background-image: url(/img/tmp/toy2.svg);
        }
        .widget-am3 {
            left: 580px;
            top: 120px;
            width: 49px;
            width: 49px;
            height: 49px;
            // -webkit-animation: zhuan 3s .2s infinite;
            animation: zhuan 3s .2s infinite;
            background-image: url(/img/tmp/toy3.svg);
        }
        .widget-am4 {
            right: 400px;
            top: 330px;
            width: 24px;
            height: 24px;
            -webkit-animation: toy 3s 1s infinite;
            animation: toy 3s 1s infinite;
            background-image: url(/img/tmp/toy4.svg);
        }
        .widget-am5 {
            left: 980px;
            top: 70px;
            width: 54px;
            height: 54px;
            -webkit-animation: toy 3s .8s infinite;
            animation: toy 3s .8s infinite;
            background-image: url(/img/tmp/xin.svg);
        }
        .widget-am6 {
           left: 900px;
            top: -40px;
            width: 27px;
            -webkit-animation: zhuan .5s .2s infinite;
            animation: zhuan .5s .2s infinite;
            height: 23px;
            background-image: url(/img/tmp/toy6.svg);
        }
        .widget-am7 {
            left: 1271px;
            top: 373px;
            width: 40px;
            height: 40px;
            -webkit-animation: zhuan 2s 1.5s infinite;
            animation: zhuan 2s 1.5s infinite;
            background-image: url(/img/tmp/toy7.svg);
        }
        .widget-am8 {
            width: 83px;
            height: 84px;
            left: 1260px;
            top: -40px;
            -webkit-animation: toy2 3s .1s infinite;
            animation: toy2 3s .1s infinite;
            background-image: url(/img/tmp/toy2.svg);
        }
        .widget-am9 {
            width: 30px;
            height: 30px;
            right: -30px;
            top: 122px;
            -webkit-animation: aroundAnimation 3s .1s infinite;
            animation: aroundAnimation 3s .1s infinite;
            background-image: url(/img/tmp/toy2.svg);
        }
        .widget-am {
            position: absolute;
            background-size: cover;
            z-index: 0;
        }
        .widget-info{
            display: -ms-flexbox;
            display: flex;
            padding-top: 200px;
            -ms-flex-direction: column;
            flex-direction: column;
            -ms-flex-align: center;
            align-items: center;
            width: 1200px;
            height: 620px;
            margin: 0 auto;
            position: relative;
            z-index: 1;
            background-image: url("/img/tmp/tmpbg.png");
            background-size: cover;
            h1{
                font-size: 40px;
            }
            .widget-info-btn{
                display: flex;
                margin-top: 70px;
                .btn{
                    cursor: pointer;
                    box-shadow: 0 4px 8px 0 rgb(110, 192, 245,0.5);
                    border-radius: 25px;
                    background: #6ec0f5;
                    border: 1px solid #67baf0;
                    font-size: 16px;
                    padding: 8px 22px;
                    transition: all .2s;
                    color: #fff;
                }
                .btn2{
                    box-shadow: 0 4px 8px 0 rgb(110 192 245 / 50%);
                    border-radius: 25px;
                    background: #fff;
                    border: 1px solid #67baf0;
                    font-size: 16px;
                    padding: 8px 22px;
                    transition: all .2s;
                    color: #6ec0f5;
                }
                span{
                    width: 20px;
                }
            }
        }
    }
}
</style>


<script>
import { mapState } from "vuex"
export default {
    data(){
        return{
            
        }
    },
    computed:{
        ...mapState(["design"]),
    },
    methods: {
        joinQQqun(){
            window.open("https://qm.qq.com/cgi-bin/qm/qr?k=Spc0kkqEvVNoYMx6xgD2eeOBXDKafj_E&jump_from=webapi")
        }
    }, 
}
</script>