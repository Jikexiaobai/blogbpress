<template>
    <div class="item">
        <div class="item-cover">
            <i class="cover" style="background-image:url(/img/0.jpg)"></i>
        </div>
        <div class="user-info">
            <div class="user-avatar">
                <Avatar
                    :verifyRight="-2"
                    :verifyBottom="2"
                    :verifySize="15"
                    isVerify
                    shape="circle" 
                    src="/img/0.jpg" 
                    :size="50"
                />
                <div class="user-nickName-meta">
                    <h2>用户名称</h2>
                    <div class="meta">
                        <a-space size="large">
                            <a-space>
                                <span>关注:</span>
                                <span>80</span>
                            </a-space>
                            <a-space>
                                <span>课程:</span>
                                <span>80</span>
                            </a-space>
                        </a-space>
                    </div>
                </div>
            </div>
            <div class="user-desc">
                <a-space>
                    <a-icon type="profile" />
                    <span class="desc">鹅厂射鸡师，不定时分享原创设计干货文章，希望大家可以少...</span>
                </a-space>
            </div>
            <div class="user-follow">
                关注
            </div>
        </div>
        <div class="item-list">
            <h3>
                最新课程
            </h3>
            <ul>
                <li>
                    <div class="cover">
                        <img src="/img/0.jpg" alt="">
                    </div>
                    <div class="title">
                        <h2>民国时期的平面设计到底有多潮？穿着旗袍开机车！</h2>
                        <div  class="mt">
                            <!-- <span class="cate">免费</span> -->
                            <span class="price">
                                <i class="price-symbol">￥</i>80
                            </span>
                        </div>
                    </div>
                </li>
                <li>
                    <div class="cover">
                        <img src="/img/0.jpg" alt="">
                    </div>
                    <div class="title">
                        <h2>民国时期的平面设计到底有多潮？穿着旗袍开机车！</h2>
                        <div  class="mt">
                            <!-- <span class="cate">免费</span> -->
                            <span class="price">
                                <i class="price-symbol">￥</i>80
                            </span>
                        </div>
                    </div>
                </li>
            </ul>
        </div>
    </div>
</template>

<script>
export default {

}
</script>

<style lang="less" scoped>
    .item{
        padding: 50px 20px 0 20px;
        background-color: #fff;
        border-radius: 10px;
        overflow: hidden;
        position: relative;
        -webkit-transition: all .2s;
        transition: all .2s;
        .item-cover{
            position: absolute;
            left: 0;
            right: 0;
            top: 0;
            padding-top: 100%;
            overflow: hidden;
            .cover{
                -webkit-transition: all .2s;
                transition: all .2s;
                padding-top: 100%;
                position: absolute;
                left: -50px;
                right: -50px;
                top: -50px;
                width: auto;
                background-size: 200%;
                -webkit-filter: blur(100px);
                filter: blur(100px);
            }
        }
        .item-cover::before{
            content: '';
            display: block;
            position: absolute;
            left: 0;
            right: 0;
            bottom: 0;
            padding-top: 54%;
            background-color: #fff;
            z-index: 1;
        }
        .item-cover::after{
            content: '';
            display: block;
            position: absolute;
            left: 0;
            right: 0;
            bottom: 54%;
            padding-top: 11%;
            background-image: -webkit-linear-gradient(top, rgba(255,255,255,0.7) 0%, #FFFFFF 100%);
            background-image: linear-gradient(
            180deg, rgba(255,255,255,0.7) 0%, #FFFFFF 100%);
        }
        .user-info{
            position: relative;
            box-shadow: 0 20px 40px 0 rgb(0 0 0 / 2%);
            padding: 15px;
            background-color: #ffffff;
            border-radius: 10px;
            z-index: 1;
            .user-avatar{
                display: flex;
                align-content: center;
                .user-nickName-meta{
                    h2{
                        font-size: 16px;
                        margin-bottom: .4em;
                    }
                    .meta{
                        font-size: 12px;
                        color: #9a9a9a;
                    }
                }
            }
            .user-desc{
                font-size: 12px;
                padding-top: 1.15em;
                .desc{
                    color: #b5b5b5;
                    display: -webkit-box;
                    -webkit-box-orient: vertical;
                    -webkit-line-clamp: 1;
                    overflow: hidden;
                    text-justify: inter-ideograph;
                    word-break: break-all;
                }
                
            }
            .user-follow{
                display: none;
                position: absolute;
                right: 0;
                top: 0;
                background: #ff6000;
                color: #fff;
                padding: 3px .8rem;
                border-radius: 0 10px 0 10px;
                -webkit-transition: all .2s;
                transition: all .2s;
                cursor: pointer;

                outline: none;
                border: none;
                text-align: center;
            }
        }
        .item-list{
            font-size: 14px;
            position: relative;
            z-index: 1;
            padding-top: 20px;
            height: auto;
            overflow: hidden;
            box-sizing: border-box;
            h3{
                font-size: 14px;
                color: #b5b5b5;
                font-weight: normal;
                margin-bottom: 1.25em;
            }
            ul{
                li{
                    display: flex;
                    margin-bottom: 1.25em;
                    .cover{
                        width: 100px;
                        height: 75px;
                        border-radius: 5px;
                        img{
                            width: 100%;
                            height: 100%;
                            border-radius: 5px;
                            object-fit: cover;
                        }
                    }
                    .title{
                        margin-left: 10px;
                        flex: 1;
                        display: flex;
                        justify-content: space-between;
                        flex-direction: column;
                        h2{
                            font-size: 14px;
                            color: #323232;
                            display: -webkit-box;
                            -webkit-box-orient: vertical;
                            -webkit-line-clamp: 2;
                            overflow: hidden;
                            text-justify: inter-ideograph;
                            word-break: break-all; 
                        }
                        .mt{
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
                    }
                }
            }
        }
    }
    .item:hover{
        box-shadow: 0 0 30px rgb(0 0 0 / 5%);
        .user-info{
            .user-follow{
                display: inline-block;
            }
        }
    }
</style>