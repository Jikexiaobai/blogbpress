<template>
    <div class="author-avatar-box" :style="{width: size + 'px',height: size + 'px',marginRight:margin  + 'px'}">
        <a-avatar :src="src" :size="size" :shape="shape" icon="user" class="author-avatar"/>
        <!-- <img v-if="isVerify" 
            src="/img/avatar/verify.png" 
            class="author-avatar-verify" 
            :style="{
                width: verifySize + 'px',
                height: verifySize + 'px',
                bottom: verifyBottom + 'px',
                right: verifyRight + 'px'
            }"> -->
    </div>
</template>
<script>
export default {
    props:{ 
        src:{
            type: String, //指定传入的类型
            //type 也可以是一个自定义构造器函数，使用 instanceof 检测。
            default: "" //这样可以指定默认的值
        },
        size: {
            type: Number, //指定传入的类型
            //type 也可以是一个自定义构造器函数，使用 instanceof 检测。
            default: 32 //这样可以指定默认的值
        },
        margin: {
            type: Number, //指定传入的类型
            //type 也可以是一个自定义构造器函数，使用 instanceof 检测。
            default: 10 //这样可以指定默认的值
        },
        verifyBottom:{
            type: Number, //指定传入的类型
            //type 也可以是一个自定义构造器函数，使用 instanceof 检测。
            default: 2 //这样可以指定默认的值
        },
        verifyRight:{
            type: Number, //指定传入的类型
            //type 也可以是一个自定义构造器函数，使用 instanceof 检测。
            default: -2 //这样可以指定默认的值
        },
        verifySize:{
            type: Number, //指定传入的类型
            //type 也可以是一个自定义构造器函数，使用 instanceof 检测。
            default: 16 //这样可以指定默认的值
        },
        isVerify:{
            type: Boolean, //指定传入的类型
            //type 也可以是一个自定义构造器函数，使用 instanceof 检测。
            default: false //这样可以指定默认的值
        },
        shape:{
            type: String, //指定传入的类型
            //type 也可以是一个自定义构造器函数，使用 instanceof 检测。
            default: "square" //这样可以指定默认的值
        }
    }

}
</script>
<style lang="less" scoped>
.author-avatar-box{
    position: relative;
    .author-avatar{
        border: 2px solid #fff;
    }
    .author-avatar-verify{
        position: absolute;
        font-weight: 400;
        line-height: 16px;
        background: #fff;
        border-radius: 100%;
        font-size: 17px;
        color: #ff3a55!important;
    }
}
</style>