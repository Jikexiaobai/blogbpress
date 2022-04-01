<template>
    <div class="info">
        <div class="user-info">
            <div class="user-info-l">
                <nuxt-link :to="{path:'/profile/' + info.userInfo.id }" class="info-link">   
                    <Avatar :src="info.userInfo.avatar+'@w60_h60'" :size="40"/>
                </nuxt-link>
                <div class="user-name-box">
                    <nuxt-link :to="{path:'/profile/' + info.userInfo.id }" class="info-link">   
                        <h2 class="user-name">{{info.userInfo.nickName}}</h2>
                    </nuxt-link>
                    <div class="user-role">
                        <a-space>
                            <span class="grade-title">
                                {{info.userInfo.grade.title}}
                            </span>
                            <span class="vip-title" v-if="info.userInfo.vip">
                                {{info.userInfo.vip.title}}
                            </span>
                        </a-space> 
                    </div>
                </div>
            </div>
            <div class="user-info-r">
                <span class="group-meta-date">
                    {{info.createTime | resetData}}
                </span>
            </div>
        </div> 
        <div class="content" @click="goInfo">
            <p class="feed-text">{{info.title}}</p>
            <div>
                <div v-if="info.type == 1" class="feed-img-box">
                    <div v-if="info.files.length == 1" class="feed-img-one">
                        <img :src="info.files[0]">
                    </div>
                    <ul v-if="info.files.length > 1">
                        <li v-for="(item,index) in info.files.slice(0, 3)" :key="index">
                            <div>
                                <img :src="item" alt="">
                                <span v-if="index === 2 && info.files.length > 3" class="image-number">
                                    +<b v-text="info.files.length - 3"></b>
                                </span>
                            </div>
                        </li>
                    </ul>
                </div>
                <!-- <div v-if="info.type == 2" class="feed-video-box">
                    
                </div> -->
            </div>
        </div>
        <div class="mate">
            <a-space size="large">
                <a-space>
                    <a-icon type="eye" />
                    <span>{{info.views | resetNum}}</span>
                </a-space>
                <a-space>
                    <a-icon type="like" />
                    <span>{{info.likes | resetNum}}</span>
                </a-space>
            </a-space>
        </div>
        <div class="divider"></div>
    </div>
</template>

<script>
const PREMSTYPE = {
    FF:1, // 付费
    PL:2, // 评论
    DL:3, // 登录
}
export default {
    props: {
        info:{
            type: Object,
            default: null
        }
    },
    filters: {
        restPrems(value) {
            switch (value) {
                case PREMSTYPE.FF:
                    return "付费可见"
                    break;
                case PREMSTYPE.PL:
                    return "评论可见"
                    break;
                case PREMSTYPE.DL:
                    return "登录可见"
                    break;
            }
        },
        restPremsDesc(value) {
            switch (value) {
                case PREMSTYPE.FF:
                    return "支付费用阅读隐藏内容"
                    break;
                case PREMSTYPE.PL:
                    return "评论后阅读隐藏内容"
                    break;
                case PREMSTYPE.DL:
                    return "登录之后阅读隐藏内容"
                    break;
            }
        },
    },
    data(){
        return{
            PREMSTYPE,
            // ORDERTYPE
            // width:0,
            // height:0,
        }
    },
    methods:{
        goInfo(){
            this.$router.push({ path: `/feed/${this.info.id}`})
        }
    }
}
</script>

<style lang="less" scoped>
.info{
    background-color: white;
    padding:10px;
    .user-info{
        display: flex;
        justify-content: space-between;
        .user-info-l{
            display: flex;
            align-items: center;
            margin-right: 10px;
            .user-name-box{
                .user-name{
                    font-size: 14px;
                }
                .user-role{
                    font-size: 12px;
                    .vip-title{
                        font-size: 12px;
                    }
                }
            }
        }
        .user-info-r{
            .group-meta-date{
                height: 20px;
                line-height: 20px;
                font-size: 12px;
                color: #b9b9b9;
            }
        }
    }
    .content{
        margin-top: 10px;
        cursor: pointer;
        margin: 10px 0;
        .feed-text{
            font-size: 14px!important;
        }
        .feed-img-box{
            margin: 10px 0;
            .feed-img-one{
                width: 190px;
                height: 190px;
                border-radius: 8px;
                img{
                    border-radius: 8px;
                    width: 100%;
                    height: 100%;
                }
            }
            ul{
                display: flex;
                margin: -5px;
                li{
                    width: 33.333333%;
                    padding: 5px;
                    div{
                        height: 0;
                        padding-top: 100%;
                        cursor: pointer;
                        overflow: hidden;
                        position: relative;
                        transition: padding-top .2s;
                        max-width: 100%;
                        img{
                            position: absolute;
                            left: 0;
                            top: 0;
                            background-color: #f5f5f5;
                            width: 100%;
                            height: 100%;
                            display: block;
                        }
                        .image-number{
                            position: absolute;
                            right: 10px;
                            top: 10px;
                            height: 24px;
                            line-height: 24px;
                            border-radius: 15px;
                            -webkit-backdrop-filter: blur(5px);
                            backdrop-filter: blur(5px);
                            padding: 0 8px;
                            font-size: 13px;
                            font-weight: 500;
                            color: #fff;
                            white-space: nowrap;
                            background-color: rgba(26,26,26,.3);
                        }
                        
                    }
                }
            }
        }
        .feed-video-box{
            margin: 10px 0;
        }
        .feed-show{
            position: relative;
            background: #f5f5f5;
            max-width: 390px;
            width: 100%;
            margin-top: 10px;
            padding: 16px;
            .feed-show-desc{
                padding-bottom: 10px;
                border-bottom: 1px solid #e5e5e5;
                margin-bottom: 10px;
                font-size: 13px;
                line-height: 1;
            }
            .feed-show-ac{
                p{
                    margin-bottom: 16px;
                    font-size: 12px;
                    color: #878787;
                }
            }
        }
    }
    .mate{
        margin-top: 10px;
    }
    .divider{
        margin-top: 10px;
        height: 2px;
        background-color: #f8f8f8;
    }
}

</style>