<template>
    <div class="widget-box">
        <div class="warper" :style="{ width: `${width}px` }">
            <div class="widget-title" v-if="tmp.showTitle == 2">
                <h2 class="title">{{tmp.title}}</h2>
            </div>
            <div class="widget-content" @click="goLink">
                
                <img :src=" tmp.cover" :style="{ height: info.height + 'px' }" alt="">
            </div>
        </div>
    </div>
</template>

<script>
import { mapState } from "vuex"
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
    },
    data(){
        return{
            tmp:null
        }
    },
    created(){
        if (this.info.content != "") {
            this.tmp = JSON.parse(this.info.content)
        }   
    },
    computed:{
        ...mapState(["design"]),
    },
    methods:{
        goLink() {
            if (this.tmp.isPlatform == 1) {
                window.location.href = this.tmp.link;
            }
            if(this.tmp.isPlatform == 2){
                this.$router.push(this.tmp.link)
            }
        },
    }
}
</script>



<style lang="less" scoped>
.widget-box{
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 20px 0;
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
    .widget-content{
        cursor: pointer;
        margin-top: 20px;
        img{
            width: 100%;
            object-fit: cover;
        }
    }
}
</style>