<template>
    <div class="widget-box">
        <div class="warper" :style="{ width:'1500px' }">
            <div class="widget-title" v-if="info.showTitle == 2">
                <h2 class="title">{{info.title}}</h2>
            </div>
            <div class="widget-content">
                <a-row :gutter="[{md:20}]">
                    <a-col :span="12">
                        <div class="item">
                            <div class="item-header" :style="{ backgroundImage: `url(/img/0.jpg)` }">
                                <span>
                                    热门问题
                                </span>
                                <nuxt-link :to="{path:'/question'}" class="link">
                                    <a-icon type="right-circle" />
                                </nuxt-link>
                            </div>
                            <ul class="item-content">
                                <li v-for="(qitem,qindex) in info.list[0]" :key="qindex">
                                    <span  :class="`num num${qindex + 1}`">{{qindex + 1}}</span>
                                    <nuxt-link :to="{path:`/question/${qitem.id}`}" class="link">
                                        <span>{{qitem.title}}</span>
                                    </nuxt-link>
                                </li>
                            </ul>
                        </div>
                    </a-col>
                    <a-col :span="12">
                        <div class="item">
                            <div class="item-header" :style="{ backgroundImage: `url(/img/0.jpg)` }">
                                <span>
                                    热门话题
                                </span>
                                <nuxt-link :to="{path:'/feed'}" class="link">
                                    <a-icon type="right-circle" />
                                </nuxt-link>
                            </div>
                            <ul class="item-content">
                                <li v-for="(qitem,qindex) in info.list[1]" :key="qindex">
                                    <span  :class="`num num${qindex + 1}`">{{qindex + 1}}</span>
                                    <nuxt-link :to="{path:`/feed/${qitem.id}`}" class="link">
                                        <span>{{qitem.title}}</span>
                                    </nuxt-link>
                                </li>
                            </ul>
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
        justify-content: space-between;
        align-items: center;
        .title{
            font-size: 14px;
            letter-spacing: .8px;
            font-weight: 700;
            position: relative;
        }
        .title::after{
            background-color: #6c757d;
            bottom: -10px;
            content: "";
            height: 3px;
            left: 0;
            position: absolute;
            width: 20px;
        }
    }
    .widget-content{
        margin-top: 20px;
        .item{
            background-color: #fff;
            border-radius: 5px;
            .item-header{
                .link{
                    color:  #fff;
                }
                padding: 1rem 1.5rem;
                border-radius: 5px 5px 0 0!important;
                position: relative;
                overflow: hidden;
                background-repeat: no-repeat;
                background-position: 50%;
                background-size: 100%;
                text-shadow: 0 1px 3px rgb(0 0 0 / 30%);
                border: 0;
                color: #fff;
                display: flex;
                justify-content: space-between;
                align-items: center;
                span{
                    font-size: 1.15rem;
                    font-weight: 600;
                }
            }
            .item-header::after{
                bottom: 0;
                content: "";
                left: 0;
                position: absolute;
                right: 0;
                top: 0;
                background-color: rgb(0 0 0 / 30%);
                z-index: -1;
            }
            .item-content{
                margin-top: 10px;
                padding: 0 20px;
                li{
                    .link{
                        color: #666;
                    }
                    margin-bottom: 10px;
                    font-size: 14px;
                    line-height: 36px;
                    margin: 4px 0;
                    cursor: pointer;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: nowrap;
                    .num{
                        display: inline-block;
                        margin-right: 2px;
                        width: 21px;
                        height: 21px;
                        line-height: 21px;
                        text-align: center;
                        color: #999393;
                    }
                    .num1{
                        color: #f93131;
                    }
                    .num2{
                        color: #ff9900;
                    }
                    .num3{
                        color: #0066ff;
                    }
                    .num4{
                        color: #00ff43;
                    }
                    .num5{
                        color:#cc00ff;
                    }
                }
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
    data(){
        return{
            
        }
    },
    computed:{
        ...mapState(["design"]),
    },
}
</script>