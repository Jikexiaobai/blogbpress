<template>
    <div class="widget-box">
        <div class="warper" :style="{ width: `${width}px` }">
            <div class="widget-title" v-if="info.showTitle == 2">
                <h2 class="title">{{info.title}}</h2>
                <nuxt-link :to="{path:'/screen',query:{cateId:0,tagId:0,mode:0,isDown:2}}" class="link">
                    <span>查看更多</span>
                    <a-icon type="right" />
                </nuxt-link>
            </div>
            <div class="widget-content">
                <a-row :gutter="[{md:12}]">
                    <a-col v-for="(item,index) in info.list" :key="index"  :span="count">
                        <nuxt-link :to="{path:`/${item.module}/${item.id}`}" class="link">
                            <div class="item">
                                <div class="item-top">
                                    <img class="item-top-icon" :src="item.cover | resetImage(60,60)" :alt="info.title">
                                    <div class="item-top-title">
                                        <h2>{{item.title}}</h2>
                                        <span class="cate" v-if="item.cateInfo != null">{{item.cateInfo.title}}</span>
                                    </div>
                                </div>
                                <div class="item-bottom">
                                    <div class="item-desc">
                                        {{item.description}}
                                    </div>
                                    <div class="item-meta">
                                        <div>
                                            <a-space>
                                                <a-space>
                                                    <a-icon type="calendar" />
                                                    <span>{{item.createTime | resetData}}</span>
                                                </a-space>
                                                <a-space>
                                                    <a-icon type="eye" />
                                                    <span>{{item.views | resetNum}}</span>
                                                </a-space>
                                            </a-space>
                                        </div>
                                        <div class="item-get">
                                            获取
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </nuxt-link>
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
        .link{
            color: #999;
            cursor: pointer;
            letter-spacing: 1px;
            line-height: 22px;
            outline: none;
            padding: 5px 10px;
            border-radius: 5px;
            font-weight: 400;
            background: #fff;
            border: 1px solid #eee;
            font-size: 14px;
            display: flex;
            justify-content: center;
            align-items: center;
        }
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
            display: block;
            cursor: pointer;
            padding: 10px;
            background: #fff;
            border: 1px solid #dee0e3;
            border-radius: 6px;
            overflow: hidden;
            .item-top{
                display: flex;
                .item-top-icon{
                    object-fit: cover;
                    height: 54px;
                    width: 54px;
                    border-radius: 6px;
                }
                .item-top-title{
                    margin-left: 10px;
                    .cate{
                        background: rgba(0, 176, 240, 0.14);
                        margin: 1px 5px 1px 0;
                        padding: 2px 4px;
                        color: #2997f7;
                        border-radius: 4px;
                        font-size: 12px;
                        font-weight: bold;
                    }
                    h2{
                        font-size: 16px;
                        font-weight: 600;
                        color: #606060;
                        overflow: hidden;
                        text-overflow: ellipsis;
                        display: -webkit-box;
                        -webkit-line-clamp: 1;
                        -webkit-box-orient: vertical;
                    }
                }
            }
            .item-bottom{
                margin-top: 15px;
                .item-desc{
                    color: #646a73;
                    font-size: 14px;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    display: -webkit-box;
                    -webkit-line-clamp: 2;
                    -webkit-box-orient: vertical;
                    height: 42px;
                }
                .item-meta{
                    color: #646a73;
                    margin-top: 10px;
                    display: flex;
                    justify-content: space-between;
                    align-items: center;
                    .item-get{
                        height: 24px;
                        min-width: 60px;
                        padding: 0 12px;
                        font-size: 14px;
                        color: #0dce9e;
                        line-height: 22px;
                        text-align: center;
                        border: 1px solid #0dce9e;
                        border-radius: 12px;
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
        width:{
            type: Number,
            default: 1500
        },
        count:{
            type: Number,
            default: 6
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