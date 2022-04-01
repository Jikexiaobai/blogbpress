<template>
    <div class="widget-box">
        <div class="warper">
            <!-- <div class="widget-title" v-if="info.showTitle == 2">
                <h2 class="title">{{info.title}}</h2>
            </div> -->
            <div :style="{ height: `${this.info.height + 20}px`, backgroundImage:`url(${info.backgroundImage})`}" class="widget-body" >
                <div class="widget-bg-overlay"></div>
                <div class="widget-info">
                    <h2><a-icon type="coffee" /> {{info.title}} - 这是一个吸引你的搜索框</h2>
                    <div class="widget-info-sreach-box">
                        <a-input-search
                            placeholder="优质内容等你搜索"
                            enter-button="搜索"
                            size="large"
                            />
                        <span class="bta">或</span>
                        <a-button type="primary" size="large">
                            开通会员
                        </a-button>
                    </div>
                    <div class="widget-info-sreach-hot" v-if="info.list.length > 0">
                        <span><a-icon type="fire" theme="filled"/> 热门搜索：</span>
                        <ul>
                            <li v-for="(item) in info.list" :key="item">
                                {{item}}
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>

export default {
    props: {
        info:{
            type: Object,
            default: {}
        },
    },
}
</script>

<style lang="less" scoped>
.widget-box{

    margin: 20px 0;
    // margin-bottom: 500px;
    display: flex;
    justify-content: center;
    align-items: center; 
    width: 100%;
    .warper{
        width: 100%;
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
        .widget-body{
            position: relative;
            background-repeat: no-repeat;
            background-position: 50%;
            background-size: 150%;
            .widget-bg-overlay{
                height: 100%;
                width: 100%;
                top: 0;
                left: 0;
                position: absolute;
                filter: brightness( 100% ) contrast( 100% ) saturate( 100% ) blur( 0px ) hue-rotate(0deg);
                background-color: #191e4b;
                opacity: 0.5;
                transition: background 0.3s, border-radius 0.3s, opacity 0.3s;
            }
            .widget-info{
                z-index: 1;
                height: 100%;
                width: 100%;
                position: absolute;
                top: 0;
                left: 0;
                display: flex;
                justify-content: center;
                align-items: center;
                flex-direction: column;
                h2{
                    color: #fff;
                    letter-spacing: 2px;
                    font-size: 24px;
                    font-weight: 700;
                    line-height: 1.2;
                    margin: 0 0 15px;
                }
                .widget-info-sreach-box{
                    margin: 20px 0;
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    /deep/ .ant-input-search{
                        width: 600px;
                    }
                    .bta{
                        margin: 0 15px;
                        color: white;
                        font-size: 18px;
                        font-weight: 700;
                    }
                }
                .widget-info-sreach-hot{
                    font-size: 18px;
                    color: white;
                    display: flex;
                    justify-content: flex-start;
                    align-items: center;
                    ul{
                        display: flex;
                        justify-content: center;
                        align-items: center;
                        li{
                            margin-right: 10px;
                        }
                    }
                }
            }    
        }
    }
    
    
}
</style>