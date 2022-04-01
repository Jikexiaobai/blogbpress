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
                <a-row :gutter="[{md:12},{md:12}]">
                    <a-col :span="count" v-for="(item,index) in info.list" :key="index">
                        <nuxt-link :to="{path:`/${item.module}/${item.id}`}" class="link">
                            <div class="item">
                                <!-- <span class="item-free" v-if="item.cateInfo != null">
                                    {{item.cateInfo.title}}
                                </span> -->
                                 <img :src="item.cover | resetImage(240,170)" :alt="info.title">
                                <div class="item-title">
                                    <figcaption>
                                        <p>
                                            {{item.title}}
                                        </p>
                                    </figcaption>
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
            border-radius: 5px;
            height: 180px;
            position: relative;
            img{
                object-fit: cover;
                border-radius: 5px;
                width: 100%;
                height: 100%;
            }
            .item-free{
                position: absolute;
                top: 10px;
                left: 10px;
                background: rgba(0, 176, 240, 0.14);
                margin: 1px 5px 1px 0;
                padding: 2px 4px;
                color: #2997f7;
                border-radius: 4px;
                font-size: 12px;
                font-weight: bold;
            }
            .item-title{
                color: #fff;
                figcaption{
                    position: relative;
                    width: 100%;
                    p{
                        position: absolute;
                        bottom: 0;
                        left: 0;
                        padding: 2em;
                        width: 87%;
                        opacity: 1;
                        opacity: 0;
                        transition: opacity 0.35s, transform 0.35s;
                        -webkit-transform: translate3d(0,10px,0);
                        transform: translate3d(0,10px,0);
                        letter-spacing: 1px;
                        font-size: 68.5%;
                    }
                } 
            }
        }
        .item:hover{
            box-shadow: 0px 2px 5px rgb(228, 228, 228);
            // -webkit-box-shadow: 0 0 10px rgb(0 0 0 / 5%);
            // box-shadow: 0 0 10px rgb(0 0 0 / 5%);
            -webkit-transform: translateY(-3px);
            -ms-transform: translateY(-3px);
            transform: translateY(-3px);
            .item-title{
                p{
                    opacity: 1;
                    -webkit-transform: translate3d(0,0,0);
                    transform: translate3d(0,0,0);
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
    components:{
       
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