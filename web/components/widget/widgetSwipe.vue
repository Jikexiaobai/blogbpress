<template>
    <div class="widget-box">
        <div class="warper" :style="{ width: `${width}px` }">
            <a-row v-if="info.list != null" 
            :gutter="[16,16]" 
            >
                <a-col :span="left">
                    <swiper  class="swiper-container" :style="{ height:  `${height}px` }" ref="widgetCarouselOne" :options="swiperOption"   >
                        <swiper-slide class="swiper-item"  v-for="(item,index) in llist" :key="index">
                            <div class="swiper-info">
                                <img class="swiper-cover" :src="item.cover">
                            </div>
                        </swiper-slide>
                        <div class="swiper-pagination" slot="pagination"></div>
                    </swiper>
                </a-col>
                <a-col :span="right">
                    <a-row :gutter="[16,16]">
                        <a-col v-for="(item,index) in rlist" :key="index" :span="rightCount">
                            <img :style="{ height:  `${height / 2 - 8}px` }" class="rcover" :src="item.cover" alt="">
                        </a-col>
                    </a-row>
                </a-col>
            </a-row>
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
        width:{
            type: Number,
            default: 1500
        },
        height:{
            type: Number,
            default: 370
        },
        left:{
            type: Number,
            default: 8
        },
        right:{
            type: Number,
            default: 16
        },
        rightCount:{
            type: Number,
            default: 6
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
                // grabCursor:true,
                preventClicksPropagation : false,
                // delay:1000
                on: {
                    click: function () {
                        vm.goLink(this.realIndex)
                    },
                    slideChange: function () {
                        vm.activeIndex = this.realIndex
                    }
                },
                pagination:false,
                // pagination: {
                //     // type: 'progressbar',
                //     el: '.swiper-pagination',
                //     dynamicBullets: false
                // },
            },
            activeIndex:0,
            list:[],
            llist:[],
            rlist:[],
        }
    },
    computed:{
        ...mapState(["design"]),
        swiper() {
            return this.$refs.widgetCarouselOne.swiper
        }
    },
    created(){
        vm = this
        this.list = JSON.parse(this.info.list)

        // this.list.forEach((item,index) => {
        //     this.llist.push(item)
        // });


        if (this.list.length == 1) {
            this.llist.push(this.list[0])
            this.rlist.push(this.list[0])
        }else if (this.list.length == 2) {
            this.llist.push(this.list[0])
            this.rlist.push(this.list[1])
        }else if (this.list.length == 3 ){
            this.llist.push(this.list[0])
            this.rlist.push(this.list[1])
            this.rlist.push(this.list[2])
        }else if (this.list.length == 4 ){
            this.llist.push(this.list[0])
            this.llist.push(this.list[1])
            this.rlist.push(this.list[2])
            this.rlist.push(this.list[3])
        }else if (this.list.length > 4){
            let tmpCount = this.list.length - 2
            this.list.forEach((item,index) => {
                if (index < tmpCount) {
                    this.llist.push(item)
                }
                if (index > tmpCount) {
                    this.rlist.push(item)
                }
            });
        }
      
        // console.log(this.rlist[0].contentInfo)
    },
    methods: {
        goLink() {
            if (this.list[this.activeIndex].isPlatform == 1) {
                window.open(this.list[this.activeIndex].link)
            }
            if(this.list[this.activeIndex].isPlatform == 2){
                this.$router.push(this.list[this.activeIndex].link)
            }
        },
    }, 
}
</script>

<style lang="less" scoped>
.widget-box{
    margin-bottom: 20px;
    display: flex;
    justify-content: center;
    .warper{
        
        .swiper-container{
            // position: relative;
            width: 100%;
            .swiper-item{
                cursor: pointer;
                height: 100%;
                width: 100%;
                .swiper-info{
                    height: 100%;
                    width: 100%;
                    // position: absolute;
                    .swiper-cover{
                        width: 100%;
                        height: 100%;
                        object-fit: cover;
                        // position: absolute;
                        // top: 0;
                        // left: 0;
                    }
                }
            }
            --swiper-pagination-color:  white;/* 两种都可以 */
        }
        .rcover{
            width: 100%;
            object-fit: fill;
        }
    } 
    .wap{
        display: none;
    }
}

@media screen and (max-width: 768px) {
    .widget-box{
        .web{
            display: none;
        }
    }
}
</style>