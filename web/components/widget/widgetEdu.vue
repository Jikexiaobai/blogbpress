<template>
    <div class="widget-box">
        <div class="warper" :style="{ width: `${width}px` }">
            <div class="widget-title" v-if="info.showTitle == 2">
                <h2 class="title">{{info.title}}</h2>
            </div>
            <div class="widget-content">
                <a-row :gutter="[{md:12}]">
                    <a-col v-for="(item,index) in info.list" :key="index"  :span="count">
                        <ClassItem :info="item" :height="info.height"/>
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
        .group-item{
            position: relative;
            background: white;
            .group-info{
                border-bottom: 1px solid #efefef;
                .group-cover{
                    width: 100%;
                    height: 0;
                    padding-top: 40%;
                    position: relative;
                    overflow: hidden;
                    img{
                        position: absolute;
                        top: 0;
                        left: 0;
                        width: 100%;
                        max-width: 100%;
                        height: auto;
                        object-fit: cover;
                        image-rendering: -webkit-optimize-contrast;
                    }
                }
                .group-title{
                    margin: 10px;
                    h2{
                        font-size: 20px;
                        font-weight: 700;
                    }
                }
            }
            .group-joins{
                padding: 8px 18px;
                position: absolute;
                z-index: 2;
                left: 10px;
                top: 30px;
                border-radius: 0 4px 4px 0;
                font-size: 13px;
                background-color: #e32c28;
                display: inline-block;
                color: #fff;
            }
            .group-content{
                padding: 20px 0;
                li{
                    padding: 5px 15px;
                    font-size: 13px;
                    display: flex;
                    line-height: 1.4;
                    align-items: center;
                    h3{
                        -webkit-box-flex: 1;
                        -ms-flex: 1 1;
                        flex: 1 1;
                        overflow: hidden;
                        text-overflow: ellipsis;
                        white-space: nowrap;
                    }
                }
            }
        }
    }
}
</style>


<script>
import ClassItem from "@/components/list/classItem"
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
        ClassItem,
    },
    computed:{
        ...mapState(["design"]),
    },
    
}
</script>