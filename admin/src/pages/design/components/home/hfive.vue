<template>
    <div class="setting">
        <a-row :gutter="[{md:12}]">
            <a-col :span="24" :md="18">
                <div class="web-box">
                    <div class="web-header">
                        网站顶部
                    </div>
                    <div class="web-body">
                        <h2>请拖动组件库到下方</h2>
                        <Vuedraggable class="wrapper" group="site" v-model="list">
                            <div v-for="(item,index) in list" :key="3234+index"  class="item">
                                <div class="item-top">
                                    <div>
                                        <a-icon @click="remove(index)" type="close-circle" class="item-close"/>
                                        <span class="item-title">{{item.title}}</span>
                                    </div>
                                    <a-icon  @click="open(index)" type="down" />
                                </div>
                                <div v-if="item.isOpen" class="item-bottom">
                                    <!-- 组件标题 -->
                                    <a-row class="setting-item" :gutter="[{md:20}]">
                                        <a-col :span="8">
                                            <a-input 
                                                size="large"
                                                v-model="item.title" 
                                                placeholder="请输入组件标题" 
                                            />
                                        </a-col>
                                        <a-col :span="8">
                                            <a-input 
                                                size="large"
                                                v-model="item.desc" 
                                                placeholder="请输入组件描述" 
                                            />
                                        </a-col>
                                        <a-col :span="8">
                                            <a-select 
                                                size="large"
                                                style="width: 100%;"
                                                placeholder="请设置是否显示组件标题和组件描述" 
                                                :default-value="2" 
                                                v-model="item.showTitle">
                                                <a-select-option :value="1">
                                                    否
                                                </a-select-option>
                                                <a-select-option :value="2">
                                                    是
                                                </a-select-option>
                                            </a-select>
                                        </a-col>
                                    </a-row>

                                    <!-- 组件高度 -->
                                    <a-row class="setting-item" 
                                        :gutter="[{md:20}]">
                                        <a-col :span="24">
                                            <a-input 
                                                size="large"
                                                v-model="item.height" 
                                                placeholder="请输组件高度" 
                                            />
                                        </a-col>
                                    </a-row>
                                    <!-- 轮播组件数据添加 -->
                                    <a-row 
                                        v-if="item.style == h5Widget.swpie"
                                        class="setting-item" 
                                        :gutter="[{md:20},{md:20}]">
                                        <a-col :span="24">
                                            <a-button @click="addSwiper(index)" size="large" type="primary" block>
                                                添加轮播内容
                                            </a-button>
                                        </a-col>
                                    </a-row>
                                    <div v-if="item.style == h5Widget.swpie">
                                        <a-row 
                                            v-for="(citem,cindex) in item.list" :key="cindex"
                                            class="setting-item" 
                                            :gutter="[{md:20},{md:20}]">
                                            <a-col :span="7">
                                                <a-input 
                                                    size="large"
                                                    v-model="citem.link"
                                                    :placeholder="citem.isPlatform == 2 ? '请输入平台地址,无需携带域名开头' : '请输入携带完整域名地址'" 
                                                />
                                            </a-col>
                                            <a-col :span="5">
                                                <a-input 
                                                    size="large"
                                                    v-model="citem.cover"
                                                    placeholder="封面地址" 
                                                />
                                            </a-col>
                                            <a-col :span="5">
                                                <a-select 
                                                    style="width: 100%;"
                                                    size="large"
                                                    placeholder="平台跳转" 
                                                    v-model="citem.isPlatform">
                                                    <a-select-option :value="1">
                                                        否
                                                    </a-select-option>
                                                    <a-select-option :value="2">
                                                        是
                                                    </a-select-option>
                                                </a-select>
                                            </a-col>
                                            <a-col :span="4">
                                                <a-button block @click="removeSwiper(index,cindex)" size="large" type="danger">
                                                    删除
                                                </a-button>
                                            </a-col>
                                        </a-row>
                                    </div>

                                    
                                    <!-- 社区组件数据编号 -->
                                    <a-row 
                                        v-if="item.style == h5Widget.community" 
                                        class="setting-item" 
                                        :gutter="[{md:20},{md:20}]">
                                        <a-col :span="24">
                                            <a-input 
                                                size="large"
                                                v-model="item.topicIds" 
                                                placeholder="请输话题帖子id多个请英文逗号分开,例如:1,2,4" 
                                            />
                                        </a-col>
                                        <!-- <a-col :span="12">
                                            <a-input 
                                                size="large"
                                                v-model="item.questionIds" 
                                                placeholder="请输问题帖子id多个请英文逗号分开,例如:1,2,4" 
                                            />
                                        </a-col> -->
                                    </a-row>
                                    
                                    <!-- 文章组件数据 -->
                                    <a-row 
                                        v-if="item.style == h5Widget.article" 
                                        class="setting-item" 
                                        :gutter="[{md:20},{md:20}]">
                                        <a-col :span="24">
                                            <a-input 
                                                size="large"
                                                v-model="item.articleIds" 
                                                placeholder="请输文章编号id多个请英文逗号分开,例如:1,2,4" 
                                            />
                                        </a-col>
                                    </a-row>

                                    <!-- 课程组件数据 -->
                                    <a-row 
                                        v-if="item.style == h5Widget.edu" 
                                        class="setting-item" 
                                        :gutter="[{md:20},{md:20}]">
                                        <a-col :span="24">
                                            <a-input 
                                                size="large"
                                                v-model="item.eduIds" 
                                                placeholder="请输课程编号id多个请英文逗号分开,例如:1,2,4" 
                                            />
                                        </a-col>
                                    </a-row>

                                    <!-- 自定义组件数据 -->
                                    <a-row 
                                        v-if="item.style == h5Widget.custom"
                                        class="setting-item" 
                                        :gutter="[{md:20},{md:20}]">
                                        <a-col :span="24">
                                            <a-button @click="addCustom(index)" size="large" type="primary" block>
                                                添加自定义内容
                                            </a-button>
                                        </a-col>
                                    </a-row>
                                    <div v-if="item.style == h5Widget.custom">
                                        <a-row 
                                            v-for="(citem,cindex) in item.list" :key="cindex"
                                            class="setting-item" 
                                            :gutter="[{md:20},{md:20}]">
                                            <a-col :span="4">
                                                <a-input 
                                                    size="large"
                                                    v-model="citem.title"
                                                    placeholder="描述标题" 
                                                />
                                            </a-col>
                                            <a-col :span="4">
                                                <a-input 
                                                    size="large"
                                                    v-model="citem.icon"
                                                    placeholder="图标地址" 
                                                />
                                            </a-col>
                                            <a-col :span="4">
                                                <a-input 
                                                    size="large"
                                                    v-model="citem.link"
                                                    :placeholder="citem.isPlatform == 2 ? '请输入平台地址,无需携带域名开头' : '请输入携带完整域名地址'" 
                                                />
                                            </a-col>
                                            <a-col :span="4">
                                                <a-select 
                                                    style="width: 100%;"
                                                    size="large"
                                                    placeholder="是否为平台链接" 
                                                    v-model="citem.isPlatform">
                                                    <a-select-option :value="1">
                                                        否
                                                    </a-select-option>
                                                    <a-select-option :value="2">
                                                        是
                                                    </a-select-option>
                                                </a-select>
                                            </a-col>
                                            <a-col :span="2">
                                                <a-button @click="removeCustom(index,cindex)" size="large" type="danger">
                                                    删除
                                                </a-button>
                                            </a-col>
                                        </a-row>
                                    </div>

                                    <!-- 自定义图片组件 -->
                                    <a-row 
                                        v-if="item.style == h5Widget.image" 
                                        class="setting-item" 
                                        :gutter="[{md:20},{md:20}]">
                                        <a-col :span="8">
                                            <a-input 
                                                size="large"
                                                v-model="item.image" 
                                                placeholder="图片地址" 
                                            />
                                        </a-col>
                                        <a-col :span="8">
                                            <a-input 
                                                size="large"
                                                v-model="item.link" 
                                                :placeholder="item.isPlatform == 2 ? '请输入平台地址,无需携带域名开头' : '请输入携带完整域名地址'" 
                                            />
                                        </a-col>
                                        <a-col :span="8">
                                            <a-select 
                                                style="width: 100%;"
                                                size="large"
                                                placeholder="是否为平台内地址" 
                                                v-model="item.isPlatform">
                                                <a-select-option :value="1">
                                                    否
                                                </a-select-option>
                                                <a-select-option :value="2">
                                                    是
                                                </a-select-option>
                                            </a-select>
                                        </a-col>
                                    </a-row>

                                    <!-- 自定义投稿内容组件 -->
                                    <a-row 
                                        v-if="item.style == h5Widget.content" 
                                        class="setting-item" 
                                        :gutter="[{md:20},{md:20}]">
                                        <a-col :span="8">
                                            <a-input 
                                                size="large"
                                                v-model="item.audioIds" 
                                                placeholder="音频id列表" 
                                            />
                                        </a-col>
                                        <a-col :span="8">
                                            <a-input 
                                                size="large"
                                                v-model="item.videoIds" 
                                                placeholder="视频id列表" 
                                            />
                                        </a-col>
                                        <a-col :span="8">
                                            <a-input 
                                                size="large"
                                                v-model="item.resourceIds" 
                                                placeholder="资源id列表" 
                                            />
                                        </a-col>
                                    </a-row>
                                </div>
                            </div>
                        </Vuedraggable>
                    </div>
                    <div class="web-footer">
                        网站底部
                    </div>
                </div>
            </a-col>
            <a-col :span="24" :md="6">
                <div class="save">
                    <a-button @click="handleSubmit" type="primary" block>
                        保存外观
                    </a-button>
                </div>
                <div class="widget-lib">
                    <h2>组件模块</h2>
                    <div class="widget-lib-box">
                        <Vuedraggable :options="{group:{name: 'site',pull:'clone'},sort: false}" v-model="widgetList">
                            <transition-group class="wrapper">
                                <div v-for="(item,index) in widgetList" 
                                :key="index+77" class="widget-lib-item">
                                    {{item.title}}
                                </div>
                            </transition-group>
                        </Vuedraggable>
                    </div>
                </div>
            </a-col>
        </a-row>
    </div>
</template>

<script>
import Vuedraggable from 'vuedraggable';
import { Save,Info } from '@/api/system'
import {h5Widget} from "@/shared/widget"
export default {
    components: {
        Vuedraggable,
    },
    data(){
        return{
            h5Widget,
            list: [],
            widgetList:[
                {
                    title:"轮播1",
                    desc:"",
                    showTitle:undefined,
                    height:360,
                    style:h5Widget.swpie,
                    list:[
                        {
                            link:undefined,
                            cover:undefined
                        }
                    ],
                    isOpen:false
                },
                {
                    title:"社区组件",
                    desc:"",
                    showTitle:2,
                    style:h5Widget.community,
                    topicIds:"",
                    questionIds:"",
                    isOpen:false
                },
                {
                    title:"文章组件",
                    desc:"",
                    showTitle:2,
                    style:h5Widget.article,
                    ids:"",
                    isOpen:false
                },
                {
                    title:"课程组件",
                    desc:"",
                    showTitle:2,
                    style:h5Widget.edu,
                    ids:"",
                    isOpen:false
                },
                {
                    title:"自定义内容组件",
                    desc:"",
                    showTitle:2,
                    list:[],
                    style:h5Widget.custom,
                    isOpen:false
                },
                {
                    title:"自定义图片组件",
                    desc:"",
                    showTitle:2,
                    style:h5Widget.image,
                    content:"",
                    isOpen:false
                },
                {
                    title:"投稿组件",
                    desc:"",
                    showTitle:2,
                    style:h5Widget.content,
                    content:"",
                    isOpen:false
                },
            ]
        }
    },
    created(){
        this.getData()
    },
    methods:{
        async getData(){
            const res = await Info("H5Design")  
            if (res.data.info != null) {
                const value = JSON.parse(res.data.info)
                this.list = value.map((item)=>{
                        
                        const tmp = {
                            title:item.title,
                            showTitle:item.showTitle,
                            height:item.height,
                            backgroundImage:item.backgroundImage,
                            list:item.list,
                            eduIds:item.eduIds,
                            audioIds:item.audioIds,
                            videoIds:item.videoIds,
                            resourceIds:item.resourceIds,
                            articleIds:item.articleIds,
                            topicIds:item.topicIds,
                            questionIds:item.questionIds,
                            style:item.style,
                            image:item.image,
                            link:item.link,
                            isOpen:false,
                            isPlatform:item.isPlatform,
                        }
                        // if (item.style == 9) {
                        //     const roleIds = JSON.parse(item.roleIds)
                        //     tmp.roleIds = roleIds
                        // }
                    return tmp
                })
            }
            
        },
        handleSubmit(){
            this.$confirm({
                title: '是否修改',
                content: '是否要修改外观，请注意',
                onOk:() => {
                    let list = this.list.map((item)=>{
                        const tmp = {
                            title:item.title,
                            showTitle:item.showTitle,
                            height:item.height,
                            backgroundImage:item.backgroundImage,
                            list:item.list,
                            image:item.image,
                            link:item.link,
                            eduIds:item.eduIds,
                            audioIds:item.audioIds,
                            videoIds:item.videoIds,
                            resourceIds:item.resourceIds,
                            articleIds:item.articleIds,
                            topicIds:item.topicIds,
                            questionIds:item.questionIds,
                            style:item.style,
                            isPlatform:item.isPlatform,
                        }
                        return tmp
                    })
                    const config = {
                        "configValue":JSON.stringify(list),
                        "configKey": "H5Design",
                        "configName": "H5首页设计",
                        "remark" : "H5首页设计"
                    }
                    this.CreatePost(config)
                },
                onCancel() {},
            });
        },
        async CreatePost(value){
            try {
                const data = await Save(value)
                if (data.code == 1) {
                    this.$message.success(
                        data.message,
                        3
                    )
                } else {
                    this.$message.error(
                        data.message,
                        3
                    )
                }
            } catch (error) {
                setTimeout(() => {
                    this.$notification.error({
                        message: '网络错误',
                        description: "请稍后再试"
                    })
                }, 1000)
            }
        },
        add(){
            const form = {
                title:"标题",
                key:"carousel",
                count:6,
                style:1,
                mode:1,
                cateId:[],
                isOpen:false
            }
            this.list.push(form)
        },
        remove(index) {
            this.list.splice(index,1)
        },
        open(index) {
            this.list[index].isOpen = !this.list[index].isOpen
        },
        // ------------------
        addSwiper(e){
            this.list[e].list.push({
                link:"",
                cover:"",
                isPlatform:2,
            })
        },
        removeSwiper(e,ce){

            this.list[e].list.splice(ce,1)
            // this.list[e].list.push({text:""})
        },
        addCustom(e){
            this.list[e].list.push({
                title:"",
                icon:"",
                link:"",
                isPlatform:2,
            })
        },
        removeCustom(e,ce){

            this.list[e].list.splice(ce,1)
            // this.list[e].list.push({text:""})
        },
    }
}
</script>

<style lang="less" scoped>
.setting{
    .web-box{
        background: rgb(247, 247, 247);
        .web-header{
            height: 80px;
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 26px;
            font-weight: 700;
        }
        .web-body{
            background: white;
            padding: 20px 0;
            h2{
                display: flex;
                justify-content: center;
                align-items: center;
                font-size: 21px;
                font-weight: 700;
            }
            .wrapper{
                min-height: 300px;
                border: 3px solid rgb(247, 247, 247);
            }
            .item{
                
                margin-bottom: 10px;
                .item-top{
                    cursor: move;
                    background: rgb(247, 247, 247);
                    padding: 10px;
                    display: flex;
                    justify-content: space-between;
                    align-items: center;
                    .item-close{
                        color: red;
                    }
                    .item-title{
                        margin-left: 10px;
                
                        font-weight: 700;
                    }
                }
                .item-bottom{
                    border: 1px solid  rgb(247, 247, 247);
                    padding: 10px;
                    .setting-item{
                        margin-bottom: 30px;
                    }
                }
            }
            .item:hover{
                .item-top{
                    background: #fdfdfd;
                }
            }
        }
        .web-footer{
            height: 120px;
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 26px;
            font-weight: 700;
        }
    }
    .widget-lib{
        margin-top: 10px;
        .widget-lib-box{
            border: 1px solid rgb(247, 247, 247);
            .wrapper{
                display: flex;
                flex-wrap: wrap;
                .widget-lib-item{
                    cursor: move;
                    padding: 10px;
                    margin: 5px;
                    border: 1px solid rgb(247, 247, 247);
                }
            }
                    
        }
       
    }
}
</style>
