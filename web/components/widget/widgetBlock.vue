<template>
    <div class="widget-box">
        <div class="warper" :style="{ width: width + 'px' }">
            <a-row :gutter="[16,16]">
                <a-col :span="6">
                    <div class="tools-box">
                        <ul class="one-block">
                            <li class="item">
                                <div class="item-am">
                                    <img src="/img/tools/bz.svg">
                                    <span>
                                        软件工具
                                    </span>
                                </div>
                            </li>
                            <li class="item">
                                <div class="item-am">
                                    <img src="/img/tools/ca.svg">
                                    <span>
                                        资源下载
                                    </span>
                                </div>
                            </li>
                            <li class="item">
                                <div class="item-am">
                                    <img src="/img/tools/qz.svg">
                                    <span>
                                        在线学习
                                    </span>
                                </div>
                            </li>
                            <li class="item">
                                <div class="item-am">
                                    <img src="/img/tools/ng.svg">
                                    <span>
                                        有问必答
                                    </span>
                                </div>
                            </li>
                            <li class="item">
                                <div class="item-am">
                                    <img src="/img/tools/hlb.svg">
                                    <span>
                                        畅所欲言
                                    </span>
                                </div>
                            </li>
                            <li class="item">
                                <div class="item-am">
                                    <img src="/img/tools/hb.svg">
                                    <span>
                                        行业资讯
                                    </span>
                                </div>
                            </li>
                        </ul>
                    </div>
                </a-col>
                <a-col :span="6">
                    <div class="tools-box">
                        <div class="two-block">
                            <div class="title">
                                <h2>
                                    <strong class="titlestrong">
                                        热门课程
                                    </strong>
                                </h2>
                                <span>
                                    More
                                </span>
                            </div>
                        </div>
                        <div class="">

                        </div>
                    </div>
                </a-col>
                <a-col :span="6">
                    <div class="tools-box">
                        
                    </div>
                </a-col>
                <a-col :span="6">
                    <div class="tools-box">
                        
                    </div>
                </a-col>
            </a-row>
        </div>
    </div>
</template>

<script>
import { mapState } from "vuex"
export default {
    props: {
        info:{
            type: Object,
            default: {}
        },
        width:{
            type: Number,
            default: 1500
        },
    },
    data(){
        return{
        }
    },
    computed:{
        ...mapState(["design"]),
    },

    methods: {
    }, 
}
</script>

<style lang="less" scoped>
.widget-box{
    margin-bottom: 10px;
    display: flex;
    justify-content: center;
    align-items: center; 
    .warper{
        height: 160px;
        .tools-box{
            width: 100%;
            height: 160px;
            background: #fff;
            margin-bottom: 10px;
            border-radius: 3px;
            box-shadow: 0 0 10px rgb(0 0 0 / 10%);
            .one-block{
                margin: 0;
                list-style: none;
                overflow: hidden;
                padding: 0;
                border: 0;
                font-size: 100%;
                font: inherit;
                vertical-align: baseline;
                word-wrap: break-word;
                word-break: break-all;
                box-sizing: border-box;
                -webkit-box-sizing: border-box;
                font-family: Optima-Regular, PingFang SC, Microsoft YaHei, WenQuanYi Micro Hei, Arial, Helvetica, sans-serif;
                .item{
                    float: left;
                    width: 33.333333%;
                    text-align: center;
                    display: block;
                    text-align: center;
                    font-size: 14px;
                    color: #606060;
                    position: relative;
                    line-height: 1.5;
                    padding: 10px 5px;
                    box-sizing: border-box;
                    height: 80px;
                    transition: all 0.1s;
                    cursor: pointer;
                    
                    .item-am{
                        display: flex;
                        justify-content: center;
                        align-items: center;
                        flex-direction: column;
                        img{
                            height: 40px;
                            width: 40px;
                        }
                    }
                }
                .item:hover{
                    color: white;
                    background: #8c6feb;
                }
                .item:nth-child(1)::before{
                    content: "";
                    display: block;
                    position: absolute;
                    right: 0;
                    bottom: 0;
                    height: 85%;
                    border-right: 1px solid #ddd;
                }
                .item:nth-child(1)::after{
                    content: "";
                    display: block;
                    position: absolute;
                    right: 0;
                    bottom: 0;
                    width: 85%;
                    border-bottom: 1px solid #ddd;
                }
                .item:nth-child(2)::before{
                    content: "";
                    display: block;
                    position: absolute;
                    right: 0;
                    bottom: 0;
                    height: 85%;
                    border-right: 1px solid #ddd;
                }
                .item:nth-child(2)::after{
                    content: "";
                    display: block;
                    position: absolute;
                    right: 0;
                    bottom: 0;
                    width: 100%;
                    border-bottom: 1px solid #ddd;
                }
                .item:nth-child(3)::after{
                    content: "";
                    display: block;
                    position: absolute;
                    right: auto;
                    left: 0;
                    bottom: 0;
                    width: 85%;
                    border-bottom: 1px solid #ddd;
                }
                .item:nth-child(4)::before{
                    content: "";
                    display: block;
                    position: absolute;
                    right: 0;
                    bottom: auto;
                    top: 0;
                    height: 85%;
                    border-right: 1px solid #ddd;
                }
                .item:nth-child(5)::before{
                    content: "";
                    display: block;
                    position: absolute;
                    right: 0;
                    bottom: auto;
                    top: 0;
                    height: 85%;
                    border-right: 1px solid #ddd;
                }
            }
            .two-block{
               .title{
                   display: flex;
                   justify-content: space-between;
                    h2{
                        font-size: 12px;
                        padding: 10px 0 20px;
                        margin-bottom: 0;
                        height: 16px;
                        overflow: hidden;
                        position: relative;
                        margin: 0;
                        border: 0;
                        font: inherit;
                        vertical-align: baseline;
                        word-wrap: break-word;
                        word-break: break-all;
                        font-family: Optima-Regular, PingFang SC, Microsoft YaHei, WenQuanYi Micro Hei, Arial, Helvetica, sans-serif;
                        .titlestrong{
                            font-weight: 400;
                            color: #f580af;
                            margin-right: 5px;
                            font-size: 16px;
                            margin-left: 10px;
                        }
                    }
                    span{
                        display: block;
                        position: absolute;
                        right: 20px;
                        top: 20px;
                        font-size: 10px;
                        text-transform: uppercase;
                        padding: 0 5px 0;
                        border-radius: 3px;
                        border: 1px solid #b1afaf;
                        color: #b1afaf;
                    }
                   .titlestrong::before{
                        content: "";
                        display: inline-block;
                        width: 3px;
                        height: 16px;
                        vertical-align: middle;
                        background: #f580af;
                        margin-right: 6px;
                        margin-top: -2px;
                   }
               } 
            }
        }
    }
}
</style>