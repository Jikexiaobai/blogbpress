<template>
    <div class="sidebar-box">
        <div class="title">
            <h2>热门课程</h2>
        </div>
        <ul>
            <li v-for="(item,index) in list" :key="index" @click="goResource(item.id,item.module)">
                <div class="article-cover">
                    <img :src="item.cover" :alt="item.title">
                </div>
                <div class="article-info">
                    <h2>
                        {{item.title}}
                    </h2>
                </div>
            </li>
        </ul>
    </div>
</template>

<script>
import api from "@/api/index"

import {MODULE} from "@/shared/module"
import {MODE} from "@/shared/mode"

export default {
    data(){
        return{
            queryParam:{
                page:1,
                limit: 2,
                status: 2,
                mode: MODE.HOT,
                module:MODULE.EDU
            },
            list:[],
        }
    },
    mounted(){
        this.getData()
    },
    methods:{
        async getData(){
            const res = await this.$axios.get(api.getSystemFilter,{params:this.queryParam})
            if (res.code != 1) {
                this.$router.push(`/404`)
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
           this.list = res.data.list != null ? res.data.list : []
        },
        goResource(e,m){
           this.$router.push({ path: `/${m}/${e}`})
        }
    }
}
</script>

<style lang="less" scoped>
.sidebar-box{
    padding: 20px;
    background: white;
    margin-bottom: 10px;
   .title{
       h2{
            font-size: 12px;
            line-height: 1;
            color: #999;
       }
   }
   ul{
       li{
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-top: 10px;
            cursor: pointer;
            .article-cover{
                width: 83px;
                height: 51px;
                margin-right: 10px;
                img{
                    width: 100%;
                    height: 100%;
                    border-radius: 5px;
                }
            }
            .article-info{
                flex: 1;
                height: 51px;
                h2{
                    font-size: 13px;
                    display: -webkit-box;
                    -webkit-box-orient: vertical;
                    -webkit-line-clamp: 2;
                    overflow: hidden;
                    max-height: 52px;
                }
            }
       }
   }
}
</style>