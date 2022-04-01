<template>
    <div class="widget-box">
        <div class="warper" :style="{width:`${width}px`}">
            <div class="tools-one">
                <div @click="goPath('/group')" class="icon-title">
                    <div class="icon" :style="{backgroundColor:'#59ca73'}">
                        <a-icon type="alert" />
                    </div>
                    <h2>社区</h2>
                </div>
                <div @click="goPath('/feed')" class="icon-title">
                    <div class="icon" :style="{backgroundColor:'#59ca73'}">
                        <a-icon type="alert" />
                    </div>
                    <h2>动态</h2>
                </div>
                <div @click="goPath('/question')" class="icon-title" >
                    <div class="icon" :style="{backgroundColor:'#f07775'}">
                        <a-icon type="question" />
                    </div>
                    <h2>问答</h2>
                </div>
            </div>
            <div class="tools-two">
                <a-row :gutter="[{md:12},{md:12}]"  >
                    <a-col  @click="goPath(`/group?id=${1}`)" :span="count" >
                        <div class="item">
                            招聘
                        </div>
                    </a-col>
                </a-row>
            </div>
            <div class="tools-three">
                <a-row :gutter="[{md:12},{md:12}]"  >
                    <a-col class="item" @click="goPath('/news')" :span="8" >
                        <a-space>
                            <a-icon type="book" />
                            <span>专栏</span>
                        </a-space>
                    </a-col>
                    <a-col class="item" @click="goPath('/discover')" :span="8" >
                        <a-space>
                            <a-icon type="bug" />
                            <span>资源</span>
                        </a-space>
                    </a-col>
                    <a-col class="item" @click="goPath('/edu')" :span="8" >
                        <a-space>
                            <a-icon type="bank" />
                            <span>学院</span>
                        </a-space>
                    </a-col>
                    <a-col class="item" @click="goPath('/vip')" :span="8" >
                        <a-space>
                            <a-icon type="crown" />
                            <span>会员开通</span>
                        </a-space>
                    </a-col>
                    <a-col class="item" @click="goPath('/verify')" :span="8" >
                        <a-space>
                            <a-icon type="contacts" />
                            <span>实名认证</span>
                        </a-space>
                    </a-col>
                    <a-col class="item" @click="goPath('/member')" :span="8" >
                        <a-space>
                            <a-icon type="project" />
                            <span>创作中心</span>
                        </a-space>
                    </a-col>
                </a-row>
            </div>
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
        count:{
            type: Number,
            default: 3
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
        goPath(e) {
           this.$router.push(e)
        },
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
        margin-top: 10px;
        width: 100%;
        height: 72px;
        display: flex;
        justify-content: space-between;
        align-items: center;
        .tools-one{
            display: flex;
            justify-content: space-between;
            width: 210px;
            .icon-title{
                cursor: pointer;
                display: flex;
                flex-direction: column;
                justify-content: center;
                align-items: center;
                .icon{
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    border-radius: 100%;
                    height: 46px;
                    width: 46px;
                    /deep/ .anticon{
                        font-size: 24px;
                        color: #fff;
                    }
                }
                h2{
                    color: #18191c;
                    font-size: 14px;
                    // font-weight: 700;
                }
            }
        }
        .tools-two{
            margin: 0 80px;
            flex:1;
            height: 72px;
            .item{
                // 
                cursor: pointer;
                height: 32px;
                display: flex;
                justify-content: center;
                align-items: center;
                border: 1px solid #f6f7f8;
                border-radius: 6px;
                background-color: #f6f7f8;
                color: #61666d;
                text-align: center;
                font-weight: 400;
                transition: background-color .3s,color .3s;
            }
        }
        .tools-three{
            display: flex;
            justify-content: space-between;
            width: 300px;
            .item{
                cursor: pointer;
            }
            /deep/ .anticon{
                font-size: 16px;
            }
            span{
                font-size: 13px;
                color: #61666d;
            }
        }
    } 
}

@media screen and (max-width: 768px) {
    .widget-box{
        .web{
            display: none;
        }
    }
}
</style>