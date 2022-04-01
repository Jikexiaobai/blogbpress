<template>
    <div class="swipe-box">
        <swiper class="swiper-container" ref="articleSwiper" :options="swiperOption">
            <swiper-slide class="swiper-item">
                <div class="swiper-info">
                    <img src="/img/0.jpg">
                    <div  class="swiper-content-info">
                        <h2>把对方客观角度大发噶</h2>
                        <div class="user-info">
                            <Avatar src="/img/0.jpg" shape="circle" class="user-avatar" :size="25" />
                            <span>士大夫撒旦工</span>
                        </div>
                    </div>
                </div>
            </swiper-slide>
            <div class="swiper-pagination" slot="pagination"></div>
        </swiper>
    </div>
</template>

<script>
import { Swiper, SwiperSlide } from 'vue-awesome-swiper'
export default {
    components:{
        Swiper,
        SwiperSlide
    },

    data(){
        return{
            swiperOption: {
                //循环
                loop:true,
                //每张播放时长3秒，自动播放
                autoplay:true,
                //滑动速度
                speed:1000,
                // delay:1000
                pagination: {
                    type: 'progressbar',
                    el: '.swiper-pagination',
                    dynamicBullets: true
                },
            },
        }
    },
    computed:{
        swiper() {
            console.log(this.$refs.articleSwiper.swiper)
            return this.$refs.articleSwiper.swiper
        }
    },
    
}
</script>

<style lang="less" scoped>
.swipe-box{
    .swiper-container{
        margin-bottom: 30px;
        position: relative;
        .swiper-item{
            cursor: pointer;
            .swiper-info{
                width: 100%;
                height: 300px;
                img{
                    object-fit: cover;
                    width: 100%;
                    height: 100%;
                }
                .swiper-content-info{
                    position: absolute;
                    padding: 10px;
                    bottom: 0;
                    h2{
                        font-size: 22px;
                        color: #fff;
                        margin: 10px 0;
                        display: -webkit-box;
                        -webkit-box-orient: vertical;
                        -webkit-line-clamp: 1;
                        white-space: normal;
                        position: relative;
                        z-index: 2;
                        overflow: hidden;
                    }
                    .user-info{
                        display: flex;
                        color: #fff;
                        align-items: center;
                        position: relative;
                        z-index: 2;
                        span{
                            font-size: 12px;
                            color: white;
                        }
                    }
                }
            }
        }
    }
    
    .swiper-container{
        --swiper-pagination-color: #1890ff;/* 两种都可以 */
    }
}
</style>