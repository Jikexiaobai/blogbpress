<template>
    <div class="sidebar-box">
        <div class="title">
            <h2>他的圈子</h2>
        </div>
        <ul>
            <li v-for="(item,index) in list" :key="index" @click="goGroup(item.id)">
                <div class="sidear-group">
                    <div class="group-name-img">
                        <div class="group-img">
                            <img :src="item.cover">
                        </div>
                        <div class="group-name-box">
                            <div class="group-name">
                                {{item.title}}
                            </div>
                            <div class="group-info">
                                <a-space size="small">
                                    <span>{{item.joins}}个圈友</span>
                                    <i>•</i>
                                    <span>{{item.contents}}个内容</span>
                                </a-space>
                            </div>
                        </div>
                    </div>
                    <div class="group-desc">
                        <p>{{item.description}}</p>
                    </div>
                </div>
            </li>
            <li v-if="list.length < total" class="more">
                查看更多
            </li>
        </ul>
    </div>
</template>

<script>
import api from "@/api/index"

import {MODULE} from "@/shared/module"

export default {
    data(){
        return{
            queryParam:{
                page:1,
                limit: 6,
                status: 2,
                userId:0,
                isJoin:true,
                module:[MODULE.GROUP]
            },
            list:[],
            total:0,
        }
    },
    mounted(){
        this.queryParam.userId = this.$route.params.id
        this.getData()
    },
    methods:{
        async getData(){
            const res = await this.$axios.get(api.getScreenList,{params:this.queryParam})
            if (res.code != 1) {
                this.$router.push(`/404`)
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
 
            this.list = res.data.list || []
            this.total = res.data.total || 0
        },
        goGroup(e){
           this.$router.push({ path: `/group/${e}`})
        }
    }
}
</script>

<style lang="less" scoped>
.sidebar-box{
    background: white;
    margin-bottom: 10px;
    padding: 20px 20px 10px 20px;
    .title{
        margin-bottom: 10px;
        h2{
            font-size: 12px;
            line-height: 1;
            color: #999;
        }
    }
    .sidear-group{
        margin-bottom: 10px;
        .group-name-img{
            cursor: pointer;
            display: flex;
            justify-content: space-between;
            .group-img{
                height: 40px;
                width: 40px;
                margin-right: 10px;
                box-shadow: 0 0 1px #888;
                border-radius: 5px;
                img{
                    width: 100%;
                    height: 100%;
                    border-radius: 5px;
                }
            }
            .group-name-box{
                flex: 1;
                .group-name{
                    font-weight: bold;
                    font-size: 14px;
                    color: #5C7390;
                    position: relative;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: nowrap;
                }
                .group-info{
                    font-size: 12px;
                    height: 15px;
                    color: #999;
                }
            }
            
        }
        .group-desc{
            margin-top: 10px;
            background: #f9f9f9;
            p{
                height: 26px;
                display: -webkit-box;
                -webkit-box-orient: vertical;
                -webkit-line-clamp: 1;
                overflow: hidden;
                font-size: 12px;
                color: #999;
                padding: 0 7px;
                line-height: 26px;
            }
        }
    }
    .more{
        cursor: pointer;
        padding: 10px;
        display: flex;
        justify-content: center;
    }
    
}
</style>