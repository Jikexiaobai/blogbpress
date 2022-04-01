<template>
    <div class="footer">
        <!-- 背景 -->
        <div class="footer_content">
            <!-- 内容 -->
            <div class="footer_info" :style="{ width: design.layout == 'home' ? `1500px`: `${design.width}px`}">
                <a-row>
                    <a-col :span="12">
                        <div class="footer_about">
                            <!-- <img src="" alt=""> -->
                            <img class="logo" :src="base.logo" :alt="base.title">
                        </div>
                        <div class="footer_about_content">
                            {{base.description}}
                        </div>
                    </a-col>
                    <a-col :span="12">
                        <a-row type="flex" justify="end" class="footer_about_link">
                            <a-col :span="8">
                                <h4>关于我们</h4>
                                <!-- <ul>
                                    <li>关于我们</li>
                                    <li>关于我们</li>
                                    <li>关于我们</li>
                                    <li>关于我们</li>
                                </ul> -->
                            </a-col>
                            <a-col :span="8">
                                <h4>网站协议</h4>
                                <!-- <ul>
                                    <li>关于我们</li>
                                    <li>关于我们</li>
                                    <li>关于我们</li>
                                    <li>关于我们</li>
                                </ul> -->
                            </a-col>
                        </a-row>
                    </a-col>
                </a-row>
            </div>
            <!-- nav 导航 -->
            <div class="footer_nav">
                <!-- <div class="footer_nav_links" :style="{ width: design.layout == 'home' ? `1500px`: `${design.width}px`}">
                    <ul>
                        <li>
                            音频内容
                        </li>
                        <li>
                            音频内容
                        </li>
                        <li>
                            音频内容
                        </li>
                        <li>
                            音频内容
                        </li>
                        <li>
                            音频内容
                        </li>
                    </ul>
                </div> -->
                <div class="footer_bottom">
                    <div class="footer_bottom_nav" :style="{ width: design.layout == 'home' ? `1500px`: `${design.width}px`}">
                        <div class="footer_bottom_opyright">
                            <span>Copyright © 2022 </span>
                            <span>{{base.title}}</span>
                            <span class="bottom_opyright">{{base.recordNumber}}</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import { mapState } from "vuex"
export default {
    computed:{
        ...mapState(["design","base"])
    },
}
</script>
<style lang="less" scoped>
.footer{
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  .footer_content{
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        width: 100%; 
        background-color: #232323;
        .footer_info{
            padding: 50px 0;
            color: #979b9e;
            .footer_about{
               img{
                   margin-right: 20px;
                   width: 140px;
                   height: 40px;
               } 
               span{
                   font-size: 19px;
                   font-weight: bold;
               }
            }
            .footer_about_content{
                margin-top: 20px;
                width: 400px;
                font-size: 12px;
                font-weight: bold;
                color: #979b9e;
            }
            .footer_about_link{
                h4{
                    color: #979b9e;
                    font-size: 16px;
                    font-weight: bold;
                    margin-bottom: 10px;
                    display: flex;
                    justify-content: flex-end;
                }
                li{
                    display: flex;
                    justify-content: flex-end;
                }
            }
        }
        .footer_nav{
            width: 100%;
            background-color: #0a0a0a;
            display: flex;
            justify-content: center;
            align-items: center;
            flex-direction: column;
            .footer_nav_links{
                margin-top: 40px;
                margin-bottom: 20px;
                color: #939393;
                border-color: #939393;
                opacity: .5;
                text-align: center;
                display: flex;
                justify-content: center;
                border: 1px dashed #2f2f2f;
                padding: 20px;
                border-radius: 3px;
                position: relative;
                ul{
                    display: flex;
                    flex-wrap: wrap;
                    li{
                        padding: 5px 8px;
                    }
                }
            }
            .footer_bottom{
                bottom: 0;
                left: 0;
                letter-spacing: .4px;
                line-height: 1.3;
                // z-index: 0;
                width: 100%;
                background: url(https://dl.zhutix.net/2018/11/footer_bottom_bg.png) repeat-x left bottom #000;
                padding: 25px 0 25px 0;
                overflow: hidden;
                position: relative;
                display: flex;
                justify-content: center;
                align-items: center;
                .footer_bottom_nav{
                    color: #939393;
                    .bottom_opyright{
                        margin-left: 20px;
                    }
                }
            }
        }
    }
}
// 响应式处理
@media only screen and (max-width: 768px) {
  .footer{
      display: none;
  }
}

</style>