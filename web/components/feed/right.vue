<template>
    <div class="right">
        <div class="content">

            <div class="create-box">
                <div class="create-top">
                    <div class="create-top-l">
                        <Avatar 
                            class="user-avatar"
                            :verifySize="10"
                            :verifyRight="-2"
                            :verifyBottom="2"
                            :isVerify="userInfo.isVerify"
                            shape="circle" 
                            :src="userInfo.avatar+'@w60_h60'" 
                            :size="30"
                        />
                        <span class="text">在</span>
                        <div class="select-group-box">
                            <span @click="openSelectGroup" class="group-text"># {{createForm.groupInfo == null ? '点击选择圈子' : createForm.groupInfo.title}}</span>
                            <div class="select-group" v-if="searchGroup.isOpen">
                                <a-input-search 
                                    placeholder="搜索圈子" 
                                    style="width: 100%"
                                    @search="changeSearch"
                                    v-model="searchGroup.searchGroupText" />
                                <ul v-if="searchGroup.searchGroupList.length > 0 || searchGroup.isSearchGroup">
                                    <li class="item" v-for="(item,index) in searchGroup.searchGroupList" :key="index" @click="selectGroup(item)">
                                        <a-space>
                                            <a-icon type="tag" />
                                            {{item.title}}
                                        </a-space>
                                    </li>
                                    <div v-if="searchGroup.searchGroupList.length == 0 && searchGroup.isSearchGroup">
                                        <a-empty :description="false" />
                                    </div>
                                </ul>
                                <ul v-if="searchGroup.searchGroupList.length == 0 && !searchGroup.isSearchGroup">
                                    <li class="item" v-for="(item,index) in searchGroup.list" :key="index"  @click="selectGroup(item)">
                                        <a-space>
                                            <a-icon type="tag" />
                                            {{item.title}}
                                        </a-space>
                                    </li>
                                    <div v-if="searchGroup.list.length == 0 && !searchGroup.isSearchGroup">
                                        <a-empty :description="false" />
                                    </div>
                                </ul>
                            </div>
                        </div>
                        <span class="text">发布</span>
                    </div>
                    <div class="create-top-r">
                        <div @click="changeType(1)" :class="createForm.type == 1 ? 'btn mrt10 active' : 'btn mrt10'">
                            <a-icon type="bug" />
                            <span>
                                帖子
                            </span>
                        </div>
                        <div @click="changeType(2)" :class="createForm.type == 2 ? 'btn mrt10 active' : 'btn mrt10'">
                            <a-icon type="question-circle" />
                            <span>
                                提问
                            </span>
                        </div>
                    </div>
                </div>
                <div class="create-content">
                    <a-textarea  
                    :disabled="this.token == null"
                    id="feedInput" 
                    :maxLength="256"
                     @change="changeTitle" 
                     placeholder="请写点内容" 
                     v-model="createForm.title" 
                     :rows="4" />
                    <div class="create-text-num">
                        <a-popover trigger="click"  placement="bottom">
                            <template slot="content">
                                <div class="emoji-box">
                                    <button v-for="(item,index) in emoji" 
                                        @click="selectEmoji(item)" 
                                        :key="index">
                                        {{item}}
                                    </button>
                                </div>
                            </template>
                            <a-icon type="smile" />
                        </a-popover>
                        <a-progress type="circle"  :percent="createForm.titleCount" :width="20" >
                            <template #format="percent">
                                <span v-if="createForm.titleCount < 100">{{ percent }}</span>
                                <span v-if="createForm.titleCount > 100">{{ createForm.titleCount }}</span>
                            </template>
                        </a-progress>
                    </div>
                </div>
                
                <div class="create-meta-box">
                    <div v-if="feedMetaOptions.imgVisible" class="create-meta img">
                        <p>最多上传5张图片</p>
                        <ul>
                            <li v-for="(item,index) in createForm.imgList" :key="index">
                                <div>
                                    <img :src="item" alt="xxx">
                                    <b @click="removeImg(index)" class="group-img-close"><a-icon type="close" /></b>
                                </div>
                            </li>
                        </ul>
                    </div>
                    <!-- <div v-if="feedMetaOptions.videoVisible" class="create-meta video">
                        <p>只能上传一个视频</p>
                        <ul>
                            <li>
                                <div>
                                    <video preload="auto" :src="createForm.video"></video>
                                    <b @click="removeVideo" class="group-img-close"><a-icon type="close" /></b>
                                </div>
                            </li>
                        </ul>
                    </div> -->
                </div>

                <div class="create-footer">
                    <ul class="create-footer-l">
                        <li class="create-footer-l-item">
                            <a-icon type="camera" @click="selectImg"/>
                        </li>
                        <!-- <li class="create-footer-l-item">
                            <a-icon type="video-camera" @click="selectVideo"/>
                        </li> -->
                        <!-- <li class="create-footer-l-item"  >
                            <span>@</span>
                        </li> -->
                    </ul>
                    <a-button @click="postTopic" type="primary">
                        发布
                    </a-button>
                </div>

                <div class="create-mask" v-if="token == null"></div>
                <div class="create-login" v-if="token == null">
                    <h2>你还未登录</h2>
                    <h3>登录后可阅读更多话题</h3>
                    <div class="login">
                        <a-button @click="postLogin" type="primary">
                            登录
                        </a-button>
                        <a-button @click="postLogin">注册</a-button>
                    </div>
                </div>
            </div>

            <!-- 置顶列表 -->
            <div v-if="topList.length > 0" class="top-list">
                 <ul v-if="!topLoading">
                    <li  v-for="(item,index) in topList" :key="index" class="item">
                        <span>置顶</span>
                        <h2>{{item.title}}</h2>
                    </li>
                </ul>
                 <ul  v-if="topLoading">
                    <li class="loading">
                        <a-skeleton :paragraph="{ rows: 2 }" />
                    </li>
                </ul>
            </div>

            <!-- 正常列表 -->
            <div class="list">
                <ul v-if="!loading">
                    <li v-for="(item,index) in list" :key="index" class="item">
                        <FeedItem :info="item"/>
                    </li>
                    <li v-if="noMore" class="no-more">
                        已经到底了！！没有东西了
                    </li>
                </ul>
                <ul v-if="loading">
                    <li class="loading">
                        <a-skeleton avatar :paragraph="{ rows: 4 }" />
                    </li>
                    <li class="loading">
                        <a-skeleton avatar :paragraph="{ rows: 4 }" />
                    </li>
                    <li class="loading">
                        <a-skeleton avatar :paragraph="{ rows: 4 }" />
                    </li>
                    <li class="loading">
                        <a-skeleton avatar :paragraph="{ rows: 4 }" />
                    </li>
                </ul>
            </div>
            
        </div>
        <div class="sidebbr-list">
            <SidebarUserInfo v-if="token != null" :info="userInfo"/>
                <SidebarClockIn/> 
                <!-- <SidebarQuestionList /> -->
                <SidebarHotTopic />
        </div>
    </div>
</template>

<script>
import FeedItem from "@/components/list/feedItem"
import Avatar from "@/components/avatar/avatar"
import SidebarHotTopic from "@/components/sidebar/sidebarHotTopic"
import SidebarHotUserList from "@/components/sidebar/sidebarHotUserList"
import SidebarUserInfo from "@/components/sidebar/sidebarUserInfo"
import SidebarClockIn from "@/components/sidebar/sidebarClockIn"
import SidbarAdv from "@/components/sidebar/sidbarAdv"


import api from "@/api/index"
import { mapState } from "vuex"
// import {MODE} from "@/shared/mode"
import {MODULE} from "@/shared/module"
export default {
    props:{
        hotGroupList:{
            type: Array, //指定传入的类型
            //type 也可以是一个自定义构造器函数，使用 instanceof 检测。
            default: [] //这样可以指定默认的值
        },
        topLoading:{
            type: Boolean, //指定传入的类型
            //type 也可以是一个自定义构造器函数，使用 instanceof 检测。
            default: false //这样可以指定默认的值
        },
        topList:{
            type: Array, //指定传入的类型
            //type 也可以是一个自定义构造器函数，使用 instanceof 检测。
            default: [] //这样可以指定默认的值
        },
        loading:{
            type: Boolean, //指定传入的类型
            //type 也可以是一个自定义构造器函数，使用 instanceof 检测。
            default: false //这样可以指定默认的值
        },
        list:{
            type: Array, //指定传入的类型
            //type 也可以是一个自定义构造器函数，使用 instanceof 检测。
            default: [] //这样可以指定默认的值
        },
        noMore: {
            type: Boolean, //指定传入的类型
            //type 也可以是一个自定义构造器函数，使用 instanceof 检测。
            default: false //这样可以指定默认的值
        },
    },
    components:{
        FeedItem,
        SidebarHotUserList,
        SidebarHotTopic,
        SidebarUserInfo,
        SidebarClockIn,
        Avatar,
        SidbarAdv
    },
    data(){
        return{
            emoji:[
                '😁',
                '😊',
                '😎',
                '😤',
                '😥',
                '😂',
                '😍',
                '😏',
                '😙',
                '😟',
                '😖',
                '😜',
                '😱',
                '😲',
                '😭',
                '😚',
                '💀',
                '👻',
                '👍',
                '💪',
                '👊',
            ],
            feedMetaOptions:{
                imgVisible:false,
                videoVisible:false,
            },
            createForm:{
                titleCount:0,
                imgList:[],
                groupInfo:null,
                title:"",
                type: 1,
                video:null,
            },
            searchGroup:{
                isSearchGroup: false,
                isOpen: false,
                list:[],
                searchGroupList:[],
                searchGroupText: null
            }
        }
    },
    computed:{
        ...mapState("user",["userInfo","token"]),
    },
    methods:{
        // 提交
        async postTopic(){
            if (this.token == null) {
                this.$Auth("login","登录","快速登录")
                return
            }
            if (this.createForm.title.length < 7 ) {
                this.$message.error(
                    "请写点内容吧最少6个字",
                    3
                )
                return
            }
            if (this.createForm.groupInfo == null) {
                this.$message.error(
                    "请设置圈子",
                    3
                )
                return
            }
            let formData = {
                title:this.createForm.title,
                type:this.createForm.type,
                files:"",
                groupId:this.createForm.groupInfo.id
            }
            if (this.createForm.imgList.length > 0 && this.createForm.video == null) {
                formData.files = JSON.stringify(this.createForm.imgList)
            }

            // if (this.createForm.imgList.length < 1 && this.createForm.video != null) {
            //     formData.files = this.createForm.video
            //     formData.type = 2
            // }

            try {
                let res = await this.$axios.post(api.postTopicCreate,formData)
                if (res.code != 1) {
                    this.$message.error(
                        res.message,
                        3
                    )
                    return
                }
                this.$message.success(
                    res.message,
                    3
                )
               
                this.createForm.imgList = []
                this.createForm.groupInfo = null
                this.createForm.title = ""
                this.createForm.video = null
                this.feedMetaOptions.imgVisible = false
                this.feedMetaOptions.videoVisible = false
                this.createForm.titleCount = 0


                this.$emit("resetList")
            } catch (error) {
                console.log(error)
                setTimeout(() => {
                    this.$notification.error({
                        message: '网络错误',
                        description: "请稍后再试"
                    })
                }, 1000)
            }
        },
        // 登录
        postLogin(){
            this.$Auth("login","登录","快速登录")
        },
        async openSelectGroup(){
            if (this.token == null) {
                this.$Auth("login","登录","快速登录")
                return
            }
            this.searchGroup.isOpen = !this.searchGroup.isOpen
            this.searchGroup.list = this.hotGroupList  
        },
        async changeSearch(){
            if (this.searchGroup.searchGroupText != null && this.searchGroup.searchGroupText != undefined) {
                this.searchGroup.isSearchGroup = true

                const queryParam = {
                    page:1,
                    limit: 7,
                    title: this.searchGroup.searchGroupText,
                    module: MODULE.GROUP,
                }
                const res = await this.$axios.get(api.getSystemFilter,{params: queryParam}) 
                if (res.code != 1) {
                    this.$message.error(
                        res.message,
                        3
                    )
                    return
                }
                this.searchGroup.searchGroupList = res.data.list != null ? res.data.list : []
                // this.searchGroup.searchGroupList = this.searchGroup.list.filter((item)=>{
                //     return item.title.indexOf(this.searchGroup.searchGroupText) != -1
                // })
            }else{
                this.searchGroup.isSearchGroup = false
                this.searchGroup.searchGroupList = []
            }
        },
        selectGroup(e){
            this.createForm.groupInfo = e
            this.searchGroup.isSearchGroup = false
            this.searchGroup.isOpen = false
            this.searchGroup.list = []
            this.searchGroup.searchGroupList = []
            this.searchGroup.searchGroupText = null
        },
        changeType(e){
            this.createForm.type = e
        },
        // 修改标题的字数
        changeTitle(){
            this.createForm.titleCount = this.createForm.title.length
        },
        selectEmoji(e){
            if (this.token == null) {
                this.$Auth("login","登录","快速登录")
                return
            }
            var elInput =document.getElementById("feedInput");
            var startPos = elInput.selectionStart;
            var endPos = elInput.selectionEnd;
            if(startPos ===undefined|| endPos ===undefined)return 
            var txt = this.createForm.title;
            var result = txt.substring(0, startPos) + e + txt.substring(endPos)    
            this.createForm.title = result;    
            elInput.focus();  
            this.$nextTick(() => {
                elInput.selectionStart = startPos + e.length;    
                elInput.selectionEnd = startPos + e.length;
            })
        },
        selectImg(){
            if (this.token == null) {
                this.$Auth("login","登录","快速登录")
                return
            }
            this.$Upload().then((res)=>{
                if (res != false) {
                    if (this.createForm.imgList.length <= 4) {
                        this.createForm.imgList.push(res)
                    }else{
                        this.$message.error(
                           "上传图片数量最多只能为5张",
                            3
                        )
                        return
                    }
                    this.createForm.video = null
                    this.feedMetaOptions.imgVisible = true
                    this.feedMetaOptions.videoVisible = false 
                }
            }).catch((err)=>{
               this.createForm.imgList = []
                // this.createForm.link = null
            })
            
        },
        // ---------- 删除
        removeImg(i){
            this.createForm.imgList.splice(i,1)
            if (this.createForm.imgList.length == 0) {
                 this.feedMetaOptions.imgVisible = false
            }
        },
        removeVideo(){
            this.createForm.video = null
            this.feedMetaOptions.videoVisible = false 
        },
    }
}
</script>


<style lang="less" scoped>
.right{
    display: flex;
    margin-left: 200px;
    .content{
        flex: 1;
        margin:0 20px;
        .create-box{
            position: relative;
            background-color: white;
            padding: 20px 20px 10px 20px;
            .create-top{
                display: flex;
                justify-content: space-between;
                .create-top-l{
                    display: flex;
                    align-items: center;
                    .text{
                        font-size: 12px;
                        line-height: 20px;
                        margin-right:10px;
                    }
                    .select-group-box{
                        z-index: 2;
                        .group-text{
                            font-size: 12px;
                            line-height: 20px;
                            margin-right:10px;
                            user-select: none;
                            cursor: pointer;
                            padding: 3px 5px;
                            background-color: #f5f6f7;
                            color: #1e80ff;
                            border-radius: 10px;
                        }
                        position: relative;
                        .select-group{
                            position: absolute;
                            padding: 10px;
                            // left: 120px;
                            top: 100%;
                            width: 300px;
                            // height: 200px;
                            margin-top: 10px;
                            background: #fff;
                            border-radius: 4px;
                            box-shadow: 0px 10px 12px 0px rgb(133 144 166 / 15%);
                            border: 1px solid rgba(133, 144, 166, 0.1);
                            ul{
                                height: 200px;
                                overflow: auto;
                                margin-top: 10px;
                                li{
                                    border-radius: 5px;
                                    cursor: pointer;
                                    padding: 5px;
                                    display: flex;
                                    margin-bottom: 10px;
                                    font-size: 12px;
                                }
                                .item:hover{
                                    background-color: #f6f7f8;
                                }
                            }
                        }
                        .select-group::before{
                            content: "";
                            position: absolute;
                            width: 12px;
                            height: 12px;
                            top: -8px;
                            left: 20px;
                            border-top: 1px solid rgba(133, 144, 166, 0.1);
                            border-left: 1px solid rgba(133, 144, 166, 0.1);
                            -webkit-transform: rotate(45deg);
                            transform: rotate(45deg);
                            background: #fff;
                        }
                    }
                }
                .create-top-r{
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    .btn{
                        cursor: pointer;
                        user-select: none;
                    }
                    .btn:hover{
                        color: #1e80ff;
                    }
                    .active{
                        color: #1e80ff;
                    }
                }
            }
            .create-content{
                margin: 10px 0;
                border: 1px solid #f5f5f5;
                /deep/ .ant-input{
                    resize : none;
                    border: 0;
                    outline: none;
                    -webkit-box-shadow: none !important;
                    box-shadow: none !important;
                }
                .create-text-num{
                    display: flex;
                    justify-content: space-between;
                    align-items: center;
                    padding: 0 10px;
                    height: 30px;
                    /deep/ .anticon{
                        font-size: 18.8px;
                    }
                    /deep/ .ant-progress{
                        /deep/ .ant-progress-text{
                            font-size:10px;
                                transform: translate(-50%, -50%) scale(0.8);
                        }
                        
                    }
                }
            }
            .create-meta-box{
                .create-meta{
                    border: 1px dashed #dfdfdf;
                    padding: 5px;
                    margin-bottom: 10px;
                    p{
                        font-size: 12px;
                        padding: 5px;
                        display: flex;
                        align-items: center;
                        line-height: 1;
                    }
                    .group-item{
                        cursor: pointer;
                    }
                }
                .img{
                
                    ul{
                        display: flex;
                        flex-flow: wrap;
                        li{
                            width: 16.6666666%;
                            padding: 5px;
                            position: relative;
                            div{
                                height: 0;
                                padding-top: 100%;
                                position: relative;
                                display: flex;
                                align-items: center;
                                justify-content: center;
                                font-size: 12px;
                                cursor: move;
                                img{
                                    position: absolute;
                                    left: 0;
                                    top: 0;
                                    width: 100%;
                                    height: 100%;
                                    box-shadow: inset 0 0 1px rgb(137, 137, 137);
                                    border-radius: 2px;
                                }
                                .group-img-close{
                                    position: absolute;
                                    right: 0;
                                    top: 9px;
                                    width: 14px;
                                    display: block;
                                    background: rgba(255, 255, 255, 0.88);
                                    text-align: center;
                                    height: 14px;
                                    border-radius: 100%;
                                    cursor: pointer;
                                    line-height: 14px;
                                    text-align: center;
                                    margin-right: 10px;
                                    font-size: 12px;
                                }
                            }
                            
                        }
                    }
                }
                .video{
                    padding: 5px;
                    ul{
                        display: flex;
                        flex-flow: wrap;
                        li{
                            width: 33.33333%;
                            padding: 5px;
                            position: relative;
                            div{
                                height: 0;
                                padding-top: 56%;
                                position: relative;
                                display: flex;
                                align-items: center;
                                justify-content: center;
                                font-size: 12px;
                                cursor: move;
                                video{
                                    position: absolute;
                                    left: 0;
                                    top: 0;
                                    width: 100%;
                                    height: 100%;
                                    box-shadow: inset 0 0 1px rgb(137, 137, 137);
                                    border-radius: 2px;
                                }
                                .group-img-close{
                                    position: absolute;
                                    right: 0;
                                    top: 9px;
                                    width: 14px;
                                    display: block;
                                    background: rgba(255, 255, 255, 0.88);
                                    text-align: center;
                                    height: 14px;
                                    border-radius: 100%;
                                    cursor: pointer;
                                    line-height: 14px;
                                    text-align: center;
                                    margin-right: 10px;
                                    font-size: 12px;
                                }
                            }
                            
                        }
                    }
                }
                .group{
                    ul{
                        padding: 5px;
                        display: flex;
                        flex-flow: wrap;
                        overflow-y: auto;
                        max-height: calc(100% - 30px);
                    }
                }
                .viewMode{
                    .viewMode-radio{
                        padding: 5px;
                    }
                    .viewMode-price{
                        padding: 5px;
                    }
                }
            }
            .create-footer{
                display: flex;
                justify-content: space-between;
                align-items: center;
                .create-footer-l{
                    display: flex;
                    align-items: center;
                    .create-footer-l-item{
                        font-size: 25px;
                        font-weight: 600;
                        margin-right: 20px;
                        cursor: pointer;
                        user-select: none;
                    }
                }
                .create-footer-r{
                    cursor: pointer;
                    font-size: 20px;
                }
            }
            .create-mask{
                background: rgba(0, 0, 0,0.5);
                top: 0;
                left: 0;
                z-index: 3;
                position: absolute;
                width: 100%;
                height: 100%;
                background: inherit;
                -webkit-filter: blur(5px);
                -moz-filter: blur(5px);
                -ms-filter: blur(5px);
                -o-filter: blur(5px);
                filter: blur(5px);
                filter: progid:DXImageTransform.Microsoft.Blur(PixelRadius=4, MakeShadow=false);
            }
            .create-login{
                top: 0;
                left: 0;
                z-index: 4;
                position: absolute;
                width: 100%;
                height: 100%;
                display: flex;
                justify-content: center;
                align-items: center;
                flex-direction: column;
                h3{
                    margin: 20px 0;
                }
            }
        }
        .top-list{
            background: white;
            padding: 10px 20px;
            .item{
                padding: 8px 0;
                display: flex;
                align-items: center;
                color: #8590a6;
                // font-size: 14px;
                cursor: pointer;
                user-select: none;
                h2{
                    margin-left: 20px;
                    flex: 1;
                    font-size: 14px;
                    // width: 100%;
                    display: -webkit-box;
                    -webkit-box-orient: vertical;
                    -webkit-line-clamp: 1;
                    overflow: hidden;
                    word-break: break-all;
                }
            }
        }
        .list{
            margin-top: 10px;
            .item{
                background: white;
                padding: 10px;
                border-radius: 4px;
                margin-bottom: 10px;
            }
            .no-more{
                margin-top: 20px;
                display: flex;
                justify-content: center;
                align-items: center;
            }
            .loading{
                background: white;
                padding: 10px;
                border-radius: 4px;
                margin-bottom: 10px;
            }
        }
    }
    .sidebbr-list{
        width: 260px;
        // margin: 20px 0;
    }
}
</style>