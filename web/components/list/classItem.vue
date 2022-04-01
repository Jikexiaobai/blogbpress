<template>
    <div class="item">
        <div class="cover" :style="{ paddingTop: this.height + '%' }">
            <nuxt-link :to="{path:`/course/${info.id}` }" class="cover-link">
              <img :src="info.cover | resetImage(240,170)" :alt="info.title">
            </nuxt-link>
        </div>
        <div class="info">
            <h2 class="list-box-info-title">{{info.title}}</h2>
            <p>{{info.description}}</p>
            <div class="list-box-meta">
                <div  class="list-box-cate">
                    <span v-if="info.joinMode == 1" class="cate">免费</span>
                    <span v-if="info.joinMode == 2" class="price">
                        <i class="price-symbol">￥</i>{{info.price}}
                    </span>
                </div>
                <div  class="list-box-joins">
                    <span class="value">{{info.joins | resetNum}}人已报名</span>
                </div>
            </div>
            <div class="list-box-user-date">
                <nuxt-link  :to="{path:'/profile/' + 1 }" class="item-link">
                    <Avatar 
                        class="user-avatar"
                        :verifyRight="-4"
                        :verifyBottom="4"
                        :verifySize="10"
                        :isVerify="info.userInfo.isVerify"
                        shape="circle" 
                        :src="info.userInfo.avatar+'@w60_h60'" 
                        :size="28"
                    /> 
                </nuxt-link>
                <div class="list-box-user-date">
                    <nuxt-link  :to="{path:'/profile/' + info.userInfo.id }" class="item-link">   
                        <h2 class="user-name">{{info.userInfo.nickName}}</h2>
                    </nuxt-link>
                    <span>{{info.createTime | resetData}}</span>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        height: {
            type: Number,
            default: 70
        },
        info:{
            type: Object,
            default: {}
        }
    },
}
</script>

<style lang="less" scoped>
.item{
    border-radius: 8px;
    background-color: #fff; 
    .cover{
        position: relative;
        height: 0;
        overflow: hidden;
        border-top-right-radius: 8px;
        border-top-left-radius: 8px;
        .cover-link{
            display: flex;
            align-items: center;
            justify-content: center;
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            img{
                object-fit: cover;
                border-top-right-radius: 8px;
                border-top-left-radius: 8px;
                display: block;
                max-width: 100%;
                max-height: 100%;
                background-color: #fff;
                width: 100%;
                height: 100%;
            }
        }
    }
    .info{
        padding: 10px;
        .list-box-info-title{
            display: -webkit-box;
            -webkit-box-orient: vertical;
            -webkit-line-clamp: 1;
            overflow: hidden;
            // line-height: 44px;
            font-weight: 700;
            font-size: 14px;
        }
        p{
            font-size: 14px;
            color: #999;
            display: -webkit-box;
            -webkit-box-orient: vertical;
            -webkit-line-clamp: 1;
            overflow: hidden;
            text-justify: inter-ideograph;
            word-break: break-all;
        }
        .list-box-meta{
            margin-top: 5px;
            display: flex;
            justify-content: space-between;
            align-items: center;
            .list-box-cate{
                height: 30px;
                .price{
                    font-size: 18px;
                    font-weight: 600;
                    color: #ff4038;
                    .price-symbol{
                        font-style: normal;
                        font-size: 10px;
                        vertical-align: middle;
                    }
                }
                .price::after{
                    display: block;
                    content: '';
                    clear: both;
                }
                .cate{
                    background: #3d7eff;
                    margin: 1px 5px 1px 0;
                    padding: 4px 8px;
                    color: #fff;
                    border-radius: 20px;
                    font-size: 12px;
                    font-weight: bold;
                }
            }
            .list-box-joins{
                font-size: 12px;
            }
        }
        .list-box-user-date{
            display: flex;
            justify-content: space-between;
            align-items: center;
            .list-box-user-date{
                flex: 1;
                display: flex;
                justify-content: space-between;
                align-items: center;
                h2{
                    font-size: 12px;
                    font-weight: 700;
                }
                span{
                    font-size: 12px;
                }
            }
        }
    }
}
</style>