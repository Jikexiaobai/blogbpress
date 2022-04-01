<template>
    <div class="home-info">
        <div v-if="!isEmpty" class="home-body">
            <div class="home-box" v-if="resourceList.length > 0">
                <div class="home-title" >
                    <h2>发布的资源</h2>
                    <nuxt-link :to="{ path: `/profile/${queryParam.userId}/content/resource` }">
                        <span>更多</span>
                    </nuxt-link>
                </div>
                <div class="home-conten">
                    <a-row :gutter="{ md: '30'}">
                        <a-col v-for="(item,index) in resourceList" :key="index" :span="8">
                            <listTwo :info="item"/>
                        </a-col> 
                    </a-row>
                </div>
            </div> 
            <div class="home-box" v-if="audioList.length > 0">
                <div class="home-title">
                    <h2>发布的音频</h2>
                    <nuxt-link :to="{ path: `/profile/${queryParam.userId}/content/audio` }">
                        <span>更多</span>
                    </nuxt-link>
                </div>
                <div class="home-conten">
                    <a-row :gutter="[{md:30},{md:30}]">
                        <a-col v-for="(item,index) in audioList" :key="index" :span="8">
                            <listTwo :info="item"/>
                        </a-col> 
                    </a-row>
                </div>
            </div> 
            <div class="home-box" v-if="videoList.length > 0">
                <div class="home-title">
                    <h2>发布的视频</h2>
                    <nuxt-link :to="{ path: `/profile/${queryParam.userId}/content/video` }">
                        <span>更多</span>
                    </nuxt-link>
                </div>
                <div class="home-conten">
                    <a-row :gutter="[{md:30},{md:30}]">
                        <a-col v-for="(item,index) in videoList" :key="index" :span="8">
                            <listTwo :info="item"/>
                        </a-col> 
                    </a-row>
                </div>
            </div>
            <div class="home-box" v-if="eduList.length > 0">
                <div class="home-title">
                    <h2>发布的课程</h2>
                    <nuxt-link :to="{ path: `/profile/${queryParam.userId}/content/course` }">
                        <span>更多</span>
                    </nuxt-link>
                </div>
                <div class="home-conten">
                    <a-row :gutter="[{md:30},{md:30}]">
                        <a-col v-for="(item,index) in eduList" :key="index" :span="8">
                            <listTwo :info="item"/>
                        </a-col> 
                    </a-row>
                </div>
            </div> 
            <div class="home-box" v-if="articleList.length > 0">
                <div class="home-title">
                    <h2>发布的专栏</h2>
                    <nuxt-link :to="{ path: `/profile/${queryParam.userId}/content/article` }">
                        <span>更多</span>
                    </nuxt-link>
                </div>
                <div class="home-conten">
                    <a-row>
                        <a-col v-for="(item,index) in articleList" :key="index" :span="24">
                            <list-one :info="item"/>
                        </a-col> 
                    </a-row>
                </div>
            </div> 
        </div>
        <div v-if="isEmpty" class="empty">
            <a-config-provider :locale="locale">
                <a-empty />
            </a-config-provider>
        </div> 
    </div>
</template>
<script>
import ListTwo from "@/components/list/listTwo"
import ListOne from "@/components/list/listOne"


import zhCN from 'ant-design-vue/lib/locale-provider/zh_CN'

import {MODULE} from "@/shared/module"
import api from "@/api/index"
export default {
    components:{
        ListTwo,
        ListOne,
    },
    
    data(){
        return{
            locale: zhCN,
            MODULE,
            queryParam:{
                page:1,
                limit: 6,
                userId:null,
                mode: 1,
                module:""
            },
            articleList:[],
            audioList:[],
            videoList:[],
            resourceList:[],
            eduList:[],
            isEmpty:true
        }
    },
    mounted () {
        this.queryParam.userId = parseInt(this.$route.params.id)
        this.getData()
    },
    methods:{
        async getData(){
            const res = await this.$axios.get(api.getUserPosts,{params:this.queryParam})
            if (res.code != 1) {
                this.$router.push(`/404`)
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
            let articleList = []
            let audioList = []
            let videoList = []
            let resourceList = []
            let eduList = []

      
            if (res.data.list != null && res.data.list.length > 0) {
                
                res.data.list.map((item)=>{
                
                    if (item.module == MODULE.ARTICLE) {
                        articleList.push(item)
                    }
                    if (item.module == "course") {
                        eduList.push(item)
                    }
                    if (item.module == MODULE.VIDEO) {
                        videoList.push(item)
                    }
                    if (item.module == MODULE.RESOURCE) {
                        resourceList.push(item)
                    }
                    if (item.module == MODULE.AUDIO) {
                        audioList.push(item)
                    }
                })
                this.articleList = articleList || []
                this.audioList = audioList || []
                this.videoList = videoList || []
                this.resourceList = resourceList || []
                this.eduList = eduList || []
                this.isEmpty = false
            }else{
                this.isEmpty = true
            }
           
    
            
        },
    }
}
</script>

<style lang="less" scoped>
.home-info{
    
    .home-body{
        background: white;
        padding: 20px;
        .home-box{
            margin-bottom: 10px;
            border-bottom: 1px solid rgb(238, 238, 238);;
            .home-title{
                display: flex;
                justify-content: space-between;
                align-items: center;
                h2{
                    color: #000;
                    font-size: 20px;
                    font-weight: 400;
                    line-height: 33px;
                } 
            }
            .home-conten{
                padding: 20px 0;
            }
        }
        
    }
    .empty{
        background: white;
        height: 500px;
        display: flex;
        align-items: center;
        justify-content: center;
    }
    .sidebar{
        .sidebar-box{
            background: white;
            padding: 20px;
            .sidebar-title{
                border-bottom: 1px solid #e5e9ef;
                font-size: 14px;
                font-weight: 700;
                padding: 0;
                margin: -15px 0 10px;
                height: 45px;
                line-height: 45px;
            }
        } 
        .trigger{
            .info{
                display: flex;
                justify-content: space-between;
                align-items: center;
            }
            ul{
                margin-top: 10px;
                display: flex;
                flex-flow: wrap;
                li{
                    width: 10%;
                }
            }
        }
        .group{
            margin-top: 10px;
            ul{
                li{
                    display: flex;
                    align-items: center;
                    margin-bottom: 10px;
                    padding-bottom: 10px;
                    border-bottom: 1px solid #e5e9ef;
                }
                .more{
                    display: flex;
                    justify-content: center;
                    align-items: center;
                }
            }
        }
        .view{
            margin-top: 10px;
            .view-item{
                display: flex;
                align-items: center;
                justify-content: center;
            }
        }
    }
}
</style>