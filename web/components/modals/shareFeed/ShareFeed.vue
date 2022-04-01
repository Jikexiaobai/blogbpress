<template>
    <div class="modal_box center opacity" :class="[isTrue && 'is_back_show']">
        <div class="modal_box_container">
            <div class="modal_box_title">
                <div class="modal_box_title_l">
                    <span>转发到动态</span>
                </div>
                <div class="modal_box_title_r">
                    <a-icon  type="close" @click="cancel"/>
                </div>
            </div>
            <div class="modal_box_content">
                <div class="info">
                    <div class="cover">
                        <img :src="info.cover" :alt="info.title">
                    </div>
                    <h2>
                        {{info.title}}
                    </h2>
                </div>
                <div class="form">
                    <div class="feed-create-center">
                        <a-textarea  id="feedInput" 
                        :maxLength="140" 
                        @change="changeTitle" 
                        placeholder="请写点内容" 
                        v-model="createForm.title" :rows="4" />
                        <div class="feed-create-text-num">
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
                            <a-progress type="circle"  :percent="createData.titleCount" :width="20" >
                                <template #format="percent">
                                    <span v-if="createData.titleCount < 100">{{ percent }}</span>
                                    <span v-if="createData.titleCount > 100">{{ createData.titleCount }}</span>
                                </template>
                            </a-progress>
                        </div>
                    </div>
                    <div class="group">
                        <a-tag class="tag" @click="selectGroup">{{createForm.groupInfo != null ? createForm.groupInfo.title : "请选择圈子"}}</a-tag>
                        <div class="select-group" v-if="createData.isOpen">
                            <a-input-search 
                                placeholder="搜索圈子" 
                                style="width: 100%"
                                @change="changeSearch"
                                v-model="createData.searchGroupText" />
                            <ul v-if="createData.searchGroupList.length > 0 || createData.isSearchGroup">
                                <li class="item" v-for="(item,index) in createData.searchGroupList" :key="index" @click="changeGroup(item)">
                                    <a-space>
                                        <a-icon type="tag" />
                                        {{item.title}}
                                    </a-space>
                                </li>
                                <div v-if="createData.searchGroupList.length == 0 && createData.isSearchGroup">
                                    <a-empty :description="false" />
                                </div>
                            </ul>
                            <ul v-if="createData.searchGroupList.length == 0 && !createData.isSearchGroup">
                                <li class="item" v-for="(item,index) in createData.groupList" :key="index"  @click="changeGroup(item)">
                                    <a-space>
                                        <a-icon type="tag" />
                                        {{item.title}}
                                    </a-space>
                                </li>
                                <div v-if="createData.groupList.length == 0 && !createData.isSearchGroup">
                                    <a-empty :description="false" />
                                </div>
                            </ul>
                        </div>
                    </div>
                    <div @click="submit" class="post">
                        立即转发
                    </div>
                </div>
            </div>
            
        </div>
    </div>
</template>

<script>
import api from "@/api/index"
import FIcon from '@/components/icon/FIcon'
import { mapState } from "vuex"
import {MODULE} from "@/shared/module"
import {MODE} from "@/shared/mode"
export default {
    components:{
        FIcon
    },
    data() {
        return {
            module:"",
            info:{
                title:"",
                cover:"",
                id:0
            },
            createForm:{
                title:"",
                groupInfo:null,
            },
            createData:{
                groupList:[],
                searchGroupList:[],
                searchGroupText:"",
                isSearchGroup:false,
                titleCount:0,
                isOpen:false
            },
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
            // 登录输入框
            isTrue: false,
            state: null // 准备（prepare） 确定（ ascertain） 取消（cancel）
        };
    },
     computed:{
        ...mapState("user",["userInfo"]),
    },
    methods: {
        async confirm(
            info,
            module,
        ) {
            this.info = info || this.info
            this.module = module || this.module

            this.open();
            return new Promise((resolve, reject) => {
                const target = { state: 'prepare'};
                const that = this
                let res = new Proxy(target, {
                    set(event, key, value) {
                        if (value === 'ascertain') {
                            resolve(that.activeList);
                            
                        } else {
                            reject(false);
                        }
                        return true
                    }
                });
                this.state = res;
            });
        },
        async selectGroup(){
            this.createData.isOpen = !this.createData.isOpen

            const queryParam = {
                page:1,
                limit: 0,
                module: MODULE.GROUP,
                mode: MODE.HOT,
                isJoin: true,
                userId:this.userInfo.userId
            }
            const res = await this.$axios.get(api.getSystemFilter,{params: queryParam}) 
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
                this.cancel()
            }
            this.createData.groupList = res.data.list != null ? res.data.list : []
        },
        changeSearch(){
            if (this.createData.searchGroupText.length > 0) {
                this.createData.isSearchGroup = true
                this.createData.searchGroupList = this.createData.groupList.filter((item)=>{
                    return item.title.indexOf(this.createData.searchGroupText) != -1
                })
            }else{
                this.createData.isSearchGroup = false
                this.createData.searchGroupList = []
            }
        },
        // 修改标题的字数
        changeTitle(e){
            this.createData.titleCount = this.createForm.title.length
        },
        changeGroup(e){
            this.createForm.groupInfo = e
            this.createData.isOpen = !this.createData.isOpen
        },
        selectEmoji(e){
            
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
        async submit(){
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
                type:4,
                module:this.module,
                relatedId:this.info.id,
                groupId:this.createForm.groupInfo.id
            }
            let res = await this.$axios.post(api.postTopicCreate,formData)
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
            this.$message.success(
                "转发成功",
                3
            )
            this.createForm.title = ""
            this.createForm.groupInfo = null
            this.ascertain()
        },
        cancel(){
            this.state.state = "cancel"
            this.close()
        },
        ascertain(){
            this.state.state = "ascertain"
            this.close()
        },
        open() {
            this.isTrue = true;
        },
        close() {
            this.isTrue = false;
        },
        
    }
};
</script>

<style lang="less" scoped>
    .modal_box {
        user-select: none;
        pointer-events: none;
        z-index: 20;
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        width: 100%;
        height: 100%;
        opacity: 0;
        visibility: hidden;
        transform: perspective(1px) scale(1.1);
        transition: visibility 0s linear .15s,opacity .15s 0s,transform .15s;
        display: flex;
        align-items: center;
        justify-content: center;
        .modal_box_container{
            background-color: #fff;
            width: 25rem;
            margin: 0 auto;
            position: relative;
            background-image: url("/img/login.png");
            background-repeat: no-repeat;
            background-size: 100%;
            margin-top: -9%;
            .modal_box_title{
                font-size: 13px;
                display: flex;
                justify-content: space-between;
                align-items: center;
                padding: 10px 20px;
                .modal_box_title_l{
                    display: block;
                    align-items: center;
                    width: 80%;
                    overflow: hidden;
                    white-space: nowrap;
                    text-overflow: ellipsis;
                }
            }
            .modal_box_content{
                padding: 10px 20px;
                .info{
                    display: flex;
                    .cover{
                        width: 50px;
                        height: 50px;
                        img{
                            width: 100%;
                            height: 100%;
                            object-fit: cover;
                            border-radius: 5px;
                        }
                    }
                    h2{
                        margin-left: 10px;
                        flex: 1;
                        color: rgb(7, 7, 7);
                        font-size: 18px;
                        font-weight: bold;
                        display: -webkit-box;
                        -webkit-box-orient: vertical;
                        -webkit-line-clamp: 2;
                        overflow: hidden;
                        text-justify: inter-ideograph;
                        word-break: break-all;
                    }
                }
                .form{
                    margin: 10px 0;
                    .feed-create-center{
                        margin: 10px 0;
                        border: 1px solid #f5f5f5;
                        /deep/ .ant-input{
                            resize : none;
                            border: 0;
                            outline: none;
                            -webkit-box-shadow: none !important;
                            box-shadow: none !important;
                        }
                        .feed-create-text-num{
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
                    .group{
                        .tag{
                            cursor: pointer;
                        }
                        position: relative;
                        .select-group{
                            position: absolute;
                            padding: 10px;
                            // left: 120px;
                            top: 100%;
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
                    .post{
                        letter-spacing:8px;
                        cursor: pointer;
                        color: white;
                        font-size: 16px;
                        font-weight: 700;
                        border-radius: 20px;
                        margin-top: 10px;
                        display: flex;
                        justify-content: center;
                        align-items: center;
                        padding: 8px;
                        background-color: rgba(1, 122, 254, 0.5);
                    }
                }
            }
        }
        
    }
    .is_back_show {
        opacity: 1 !important;
        background: rgba(42, 44, 48, 0.7);
        pointer-events: auto !important;
        visibility: visible;
        transform: perspective(1px) scale(1);
        transition: visibility 0s linear 0s,opacity .15s 0s,transform .15s;
    }

    @media only screen and (max-width: 768px) {
        .modal_box{
            .modal_box_container{
                width: 100%;
            }
        }
    }
</style>