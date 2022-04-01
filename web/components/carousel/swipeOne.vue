<template>
    <div class="swipe-box">
        <swiper ref="mySwiper" :options="swiperOption">
            <swiper-slide>
                <div class="swiper-img">
                    <!-- <img src="/img/0.jpg"> -->
                    <div class="swiper-info">
                        <div>
                            <a-tag color="#f50">
                                圈子标题
                            </a-tag>
                        </div>
                        <h2>标题十多分钟就大发噶啥的</h2>
                        <div class="user-info">
                            <Avatar src="/img/2.jpg" shape="circle" class="user-avatar" :size="25" />
                            <span>用户名</span>
                        </div>
                    </div>
                </div>
            </swiper-slide>
            <div class="swiper-pagination" slot="pagination"></div>
        </swiper>
    </div>
</template>

<script>
import { mapState } from "vuex"
import Avatar from "@/components/avatar/avatar"
import { Swiper, SwiperSlide } from 'vue-awesome-swiper'
export default {
    components:{
        Swiper,
        SwiperSlide,
        Avatar
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
                grabCursor:true,
                // delay:1000
                pagination: {
                    // type: 'progressbar',
                    el: '.swiper-pagination',
                    dynamicBullets: true
                },
            },
        }
    },
    computed:{
        ...mapState(["allWidth"]),
        swiper() {
            console.log(this.$refs.mySwiper.swiper)
            return this.$refs.mySwiper.swiper
        }
    },
    
}
</script>

<style lang="less" scoped>
.swiper-img{
    width: 100%;
    height: 360px;
    img{
        
        object-fit: cover;
        width: 100%;
        height: 100%;
    }
    .swiper-info{
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
.swiper-container{
    --swiper-pagination-color:  white;/* 两种都可以 */
}
</style>