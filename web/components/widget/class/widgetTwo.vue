<template>
    <a-row class="content" :gutter="[{md:12}]">
        <a-col :span="4">
            <div class="item">
                <img src="https://oss-xpc0.xpccdn.com/Upload/edu/2020/09/225f69a07c387c1.png" alt="">
                <div class="text-wrap">
                    <h2>剪辑</h2>
                    <p>描述</p>
                </div>
            </div>
        </a-col>
        <a-col :span="4">
            <div class="item">
                <img src="https://oss-xpc0.xpccdn.com/Upload/edu/2020/09/225f69a07c387c1.png" alt="">
                <div class="text-wrap">
                    <h2>剪辑</h2>
                    <p>描述</p>
                </div>
            </div>
        </a-col>
        <a-col :span="4">
            <div class="item">
                <img src="https://oss-xpc0.xpccdn.com/Upload/edu/2020/09/225f69a07c387c1.png" alt="">
                <div class="text-wrap">
                    <h2>剪辑</h2>
                    <p>描述</p>
                </div>
            </div>
        </a-col>
        <a-col :span="4">
            <div class="item">
                <img src="https://oss-xpc0.xpccdn.com/Upload/edu/2020/09/225f69a07c387c1.png" alt="">
                <div class="text-wrap">
                    <h2>剪辑</h2>
                    <p>描述</p>
                </div>
            </div>
        </a-col>
        <a-col :span="4">
            <div class="item">
                <img src="https://oss-xpc0.xpccdn.com/Upload/edu/2020/09/225f69a07c387c1.png" alt="">
                <div class="text-wrap">
                    <h2>剪辑</h2>
                    <p>描述</p>
                </div>
            </div>
        </a-col>
        <a-col :span="4">
            <div class="item">
                <img src="https://oss-xpc0.xpccdn.com/Upload/edu/2020/09/225f69a07c387c1.png" alt="">
                <div class="text-wrap">
                    <h2>剪辑</h2>
                    <p>描述</p>
                </div>
            </div>
        </a-col>
    </a-row>
</template>

<script>
export default {

}
</script>

<style lang="less" scoped>
.content{
    margin-bottom: 30px;
    .item{
        height: 72px;
        background-image: none;
        background-color: #fff;
        box-shadow: 0 4px 8px 0 rgb(0 0 0 / 5%);
        display: flex;
        align-items: center;
        overflow: hidden;
        padding-left: 20px;
        img{
            width: 48px;
            height: 48px;
            object-fit: cover;
        }
        .text-wrap{
            color: #fff;
            margin-left: 14px;
            margin-top: 4px;
            h2{
                color: #333;
                font-size: 18px;
                font-weight: 500;
                letter-spacing: 5px;
            }
            p{
                color: #999;
                font-size: 12px;
                transform-origin: left;
                white-space: nowrap;
                padding-top: 4px;
                padding-left: 1px;
                -webkit-transform: scale(0.66);
                -moz-transform: scale(0.66);
                -ms-transform: scale(0.66);
                transform: scale(0.66);
            }
        }
    }
}
</style>