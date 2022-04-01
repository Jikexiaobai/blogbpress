<template>
    <div class="swipe-box">
        <swiper ref="mySwiper" :options="swiperOption">
            <swiper-slide>
                <div class="swiper-img">
                    xx
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
            console.log(this.$refs.mySwiper.swiper)
            return this.$refs.mySwiper.swiper
        }
    },
    
}
</script>

<style lang="less" scoped>
.swipe-box{
    .swiper-img{
        width: 100%;
        height: 260px;
        img{
            width: 100%;
            height: 100%;
            object-fit: cover;
        }
    }
    .swiper-container{
        --swiper-pagination-color: #fca7f0;/* 两种都可以 */
    }
}
</style>