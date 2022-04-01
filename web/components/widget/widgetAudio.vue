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
                <a-row :gutter="[{md:20},{md:20}]">
                    <a-col :span="count" v-for="(item,index) in info.list" :key="index" >
                        <ListTwo :info="item"/>
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
        ul{
            li{
                display: flex;
                .widget-media{
                    margin-bottom: 0;
                    margin-right: 20px;
                    width: 200px;
                    height: 200px;
                    img{
                        width: 100%;
                        
                    }
                }
                .widget-info{
                    flex: 1;
                }
            }
        }
        .widget-bg{
            background-image: url("/img/2.jpg");
            padding: 0 20px;
            color: #ffffff;
            height: 64px;
            line-height: 64px;
            font-size: 20px;
            border-radius: 4px 4px 0px 0px;
            background: rgb(253, 158, 62);
            background-size: 100%;
            background-position: center center;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }
        .widget-list{
            background: white;
            padding: 10px 20px 10px 20px; 
            li{
                display: flex;
                align-items: center;
                padding: 5px 0;
                span{
                    margin-right: 10px;
                }
                h2{
                    font-size: 14px;
                    color: #666;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: nowrap;
                }
                .count1{
                    color: red;
                }
                .count2{
                    color: #ff9900;
                }
                .count3{
                    color: #0099ff;
                }
                .count4{
                    color: #00ff00;
                }
                .count5{
                    color: #9900ff;
                }
                .count6{
                    color: #326786;
                }
            }
        }
    }
}
</style>


<script>
import { mapState } from "vuex"
import ListTwo from '@/components/list/listTwo'
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
        ListTwo
    },
    computed:{
        ...mapState(["design"]),
    },
    
}
</script>