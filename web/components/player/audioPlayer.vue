<template>
    <div class="m-player">
        <a-row class="warper" :style="{ width: this.allWidth + 'px' }">
            <a-col class="m-player-audio-info" :span="5">
                <div class="m-player-cover">
                    <img src="/img/1.jpg">
                </div>
                <div class="m-player-title">
                    <h2>
                        音乐标题
                    </h2>
                </div>
            </a-col>
            <a-col class="m-player-audio-ac" :span="14">
                <div class="m-player-audio-button">
                    <a-space size="middle">
                        <a-icon type="step-backward" />
                        <a-icon class="play-b" type="play-circle" />
                        <a-icon type="step-forward" />
                    </a-space>
                </div>
                <div class="m-player-audio-progess">
                    <span class="current-time">
                        0:00
                    </span>
                    <div class="m-player-audio-progess-info">
                        <a-progress class="audio-progess" :show-info="false" :percent="30" />
                        <a-slider class="audio-slider" id="test" :default-value="30" />
                    </div>
                    <span  class="duration-time">
                        3:24
                    </span>
                </div>
            </a-col>
            <a-col class="m-player-audio-r" :span="5">
                <div class="m-player-audio-volume">
                    <a-icon type="sound" />
                    <a-slider  class="slider" id="test" :default-value="30" />
                </div>
                <div class="m-player-audio-loop">
                    <a-icon type="retweet" />
                    <a-icon type="bars" />
                </div>
            </a-col>
        </a-row>
    </div>
</template>
<style lang="less" scoped>
.m-player{
    position: fixed;
    bottom: 0;
    width: 100%;
    height: 80px;
    z-index: 99;
    background-color: #222225;
    display: flex;
    justify-content: center;
    align-items: center;
    .warper{
        color: white;
        height: 100%;
        
        .m-player-audio-info{
            height: 100%;
            display: flex;
            align-items: center;
            .m-player-cover{
                width: 55px;
                min-width: 55px;
                height: 55px;
                float: left;
                border-radius: 5px;
                margin-right: 10px;
                img{
                    width: 55px;
                    height: 55px;
                    object-fit: cover;
                    display: block;
                    border-radius: 5px;
                    background-color: #3e3e42;
                    box-shadow: 0 2px 10px -2px rgba(0, 0, 0, 0.3);
                }
            }
            .m-player-title{
                h2{
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: nowrap;
                    color: #fff;
                    display: block;
                    font-size: 17px;
                    margin: 0 0 2px;
                    transition: all 0.2s;
                }
            }
        }
        .m-player-audio-ac{
            height: 100%;
            display: flex;
            flex-direction: column;
            justify-content: center;
            .m-player-audio-button{
                display: flex;
                justify-content: center;
                font-size: 24px;
                .play-b{
                    font-size: 40px;
                }
            }
            .m-player-audio-progess{
                display: flex;
                justify-content: space-around;
                align-items: center;
                .m-player-audio-progess-info{
                    margin: 0 10px;
                    width: 100%;
                    position: relative;
                    .audio-slider{
                        height: 8px;
                        padding: 0;
                        margin: 0;
                        position: absolute;
                        left: 0;
                        top: -4px;
                        z-index: 90;
                        right: 0;
                        /deep/ .ant-slider-rail{
                            height: 8px;
                            border-radius: 100px;
                            background-color: transparent;
                        }
                        /deep/ .ant-slider-track{
                            height: 8px;
                            background-color: #1890ff;
                            border-radius: 100px;
                        }
                        /deep/ .ant-slider-step{
                            height: 8px;
                            border-radius: 100px;
                        }
                        /deep/ .ant-slider-handle{
                            margin-top: -2.5px;
                        }
                    }
                    .audio-progess{
                        position: absolute;
                        z-index: 60;
                        top: -12px;
                        /deep/ .ant-progress-bg{
                            background: transparent;
                            background-color: #D7DEE3;
                        }
                    }
                }
            }
            
        }
        .m-player-audio-r{
            height: 100%;
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 20px;
            .m-player-audio-volume{
                flex: 1;
                margin-right: 10px;
                display: flex;
                align-items: center;
                .slider{
                    margin-left: 10px;
                    flex: 1;
                }
            }
        }
    }
}
</style>


<script>
import { mapState } from "vuex"

export default {
    components:{
    //    ChartsSwipe
    },
    computed:{
        ...mapState(["allWidth"])
    },
}
</script>