<template>
    <div class="widget-box">
        <div class="warper" :style="{ width: '1500px' }">
            <div class="widget-title" v-if="info.showTitle == 2">
                <h2 class="title">{{info.title}}</h2>
            </div>
            <div class="widget-body" v-if="info.list != null || info.list.length > 1" :style="{ height: `${this.info.height + 20}px` }">
                <swiper class="swiper-container" :style="{ height: `${this.info.height + 20}px` }" ref="widgetCarouselTow" :options="swiperOption"   >
                    <swiper-slide class="swiper-item"  v-for="(item,index) in info.list" :key="index">
                        <div class="swiper-info">
                            <img class="swiper-cover" :src="item.cover">
                            <div v-if="item.contentInfo != null" class="swiper-content-info">
                                <div>
                                    <a-tag v-for="(gitem,gindex) in item.contentInfo.groupList" :key="gindex" color="#f50">
                                        {{gitem.title}}
                                    </a-tag>
                                </div>
                                <h2>{{item.contentInfo.title}}</h2>
                                <div class="user-info">
                                    <Avatar :src="item.contentInfo.userInfo.avatar+'@w60_h60'" shape="circle" class="user-avatar" :size="25" />
                                    <span>{{item.contentInfo.userInfo.nickName}}</span>
                                </div>
                            </div>
                        </div>
                    </swiper-slide>
                    <div class="swiper-pagination" slot="pagination"></div>
                </swiper>
            </div>
            
            <div class="widget-empty" v-if="info.list == null || info.list.length < 1">
                <a-config-provider :locale="locale">
                    <a-empty />
                </a-config-provider>
            </div>
        </div>
    </div>
</template>

<script>
import { mapState } from "vuex"
import zhCN from 'ant-design-vue/lib/locale-provider/zh_CN'
import Avatar from "@/components/avatar/avatar"
import { Swiper, SwiperSlide } from 'vue-awesome-swiper'
var vm  =null
export default {
    props: {
        info:{
            type: Object,
            default: {}
        },
    },
    components:{
        Swiper,
        SwiperSlide,
        Avatar
    },
    data(){
        return{
            locale: zhCN,
            swiperOption: {
                
                //循环
                loop:true,
                //每张播放时长3秒，自动播放
                autoplay:true,
                //滑动速度
                speed:1000,
                on: {
                    click: function () {
                        vm.goLink(this.realIndex)
                    }
                },
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
        ...mapState(["design"]),
        swiper() {
            console.log(this.$refs.widgetCarouselTow.swiper)
            return this.$refs.widgetCarouselTow.swiper
        }
    },
    created(){
       vm = this
    },
     methods: {
        goLink(e) {
            window.location.href = this.list[e].link;
        },
    }, 
    
}
</script>

<style lang="less" scoped>
.widget-box{
    margin: 10px 0;
    display: flex;
    justify-content: center;
    align-items: center; 
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
        .swiper-container{
            border-radius: 5px;
            position: relative;
            width: 100%;
            .swiper-item{
                cursor: pointer;
                height: 100%;
                width: 100%;
                border-radius: 5px;
                .swiper-info{
                    border-radius: 5px;
                    height: 100%;
                    width: 100%;
                    position: absolute;
                    .swiper-cover{
                        border-radius: 5px;
                        width: 100%;
                        height: 100%;
                        object-fit: cover;
                        position: absolute;
                        top: 0;
                        left: 0;
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
            --swiper-pagination-color:  white;/* 两种都可以 */
        }
    
    }
    

}
</style>