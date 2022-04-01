<template>
    <div class="audio-item">
        <div class="audio-item-info">
            <div class="audio-item-cover">
                <img src="/img/2.jpg">
            </div>
            <div class="audio-item-meta">
                <div class="audio-item-meta-info">
                    <div  class="audio-item-meta-info-p">
                        <a-icon theme="filled" type="pause-circle" />
                        <div class="audio-item-meta-info-title">
                            <h2>音频名称</h2>
                            <p>音频分类</p>
                        </div>
                    </div>
                    <div class="audio-item-meta-att">
                        <a-space>
                            <span>速度：128</span>
                            <span>类型：作品</span>
                        </a-space>
                    </div>
                </div>
                <div class="audio-item-meta-wave">
                    <WaveSurfer />
                </div>
            </div>
        </div>
        <a-divider />
        <div class="audio-item-user">
            <div class="audio-item-user-buy">
                <a-button  type="danger">
                    我要购买
                </a-button>
            </div>
            <div class="audio-item-user-info">
                <Avatar :size="40"/>
                <div class="audio-item-user-name">
                    <h2>用户名称</h2>
                    <p>sdfasdfsdf</p>
                </div>
            </div>
            <div class="audio-item-data">
                <div class="audio-item-like">
                    <a-space>
                        <span><a-icon type="eye" /><span class="value">42454</span></span>
                        <span><a-icon type="like" /><span class="value">42454</span></span>
                        <span><a-icon type="message" /><span class="value">42454</span></span>
                    </a-space>
                </div>
                <div class="audio-item-date">
                    <p>2012/412/54</p>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import WaveSurfer from "@/components/player/wavesurfer"
import Avatar from "@/components/avatar/avatar"
export default {
    components:{
       WaveSurfer,
       Avatar
    },
}
</script>

<style lang="less" scoped>
.audio-item{
    background: #fff;
    box-shadow: 0 0 10px #eaeaea;
    -webkit-box-shadow: 0 0 10px #eaeaea;
    padding: 20px;
    margin-bottom: 10px;
    .audio-item-info{
        display: flex;
        justify-content: space-between;
        .audio-item-cover{
            width: 120px;
            height: 120px;
            margin-right: 10px;
            border-radius: 5px;
            img{
                width: 100%;
                height: 100%;
                border-radius: 5px;
            }
        }
        .audio-item-meta{
            flex: 1;
            .audio-item-meta-info{
                display: flex;
                justify-content: space-between;
                align-items: center;
                .audio-item-meta-info-p{
                    font-size: 40px;
                    display: flex;
                    justify-content: space-between;
                    align-items: center;
                    .audio-item-meta-info-title{
                        margin-left: 10px;
                        h2{
                            font-size: 16px;
                        } 
                        p{
                        font-size: 12px; 
                        }
                    }
                }
                .audio-item-meta-att{
                    flex: 1;
                    display: flex;
                    justify-content: flex-end;
                }
            }
            .audio-item-meta-wave{
            margin-top: 10px;
            }
        }
    }
    .audio-item-user{
        display: flex;
        justify-content: space-around;
        align-items: center;
        .audio-item-user-buy{
            width: 120px;
            display: flex;
            justify-content: center;
            align-items: center;
        }
        .audio-item-user-info{
            flex: 1;
            display: flex;
            margin: 0 10px;
            .audio-item-user-name{
                display: flex;
                justify-content: center;
                flex-direction: column;

                h2{
                    font-size: 14px;
                }
                p{
                    font-size: 12px;
                }
            }
        }
    }
    .audio-item-data{
        .audio-item-like{
            .value{
                margin-left: 5px;
                font-size: 12px;
            }
        }
        .audio-item-date{
            display: flex;
            justify-content: flex-end;
        }
    }
}
</style>