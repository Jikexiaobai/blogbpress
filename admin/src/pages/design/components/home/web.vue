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


                                    <!-- 音频组件 -->
                                    <a-row 
                                        v-if="item.style == webWidget.audio" 
                                        class="setting-item" 
                                        :gutter="[{md:20}]">
                                        <a-col :span="24">
                                            <a-input 
                                                size="large"
                                                v-model="item.audioIds" 
                                                placeholder="请输音频id多个请英文逗号分开,例如:1,2,4" 
                                            />
                                        </a-col>
                                    </a-row>
                                    <!-- 视频组件 -->
                                    <a-row 
                                        v-if="item.style == webWidget.video" 
                                        class="setting-item" 
                                        :gutter="[{md:20}]">
                                        <a-col :span="24">
                                            <a-input 
                                                size="large"
                                                v-model="item.videoIds" 
                                                placeholder="请输视频id多个请英文逗号分开,例如:1,2,4" 
                                            />
                                        </a-col>
                                    </a-row>
                                    <!-- 资源组件 -->
                                    <a-row 
                                        v-if="item.style == webWidget.resource" 
                                        class="setting-item" 
                                        :gutter="[{md:20}]">
                                        <a-col :span="24">
                                            <a-input 
                                                size="large"
                                                v-model="item.resourceIds" 
                                                placeholder="请输资源id多个请英文逗号分开,例如:1,2,4" 
                                            />
                                        </a-col>
                                    </a-row>

                                    
                                    <!-- 轮播组件数据添加 -->
                                    <a-row 
                                        v-if="item.style == webWidget.swpie"
                                        class="setting-item" 
                                        :gutter="[{md:20},{md:20}]">
                                        <a-col :span="24">
                                            <a-button @click="addSwiper(index)" size="large" type="primary" block>
                                                添加轮播内容
                                            </a-button>
                                        </a-col>
                                    </a-row>
                                    <div v-if="item.style == webWidget.swpie">
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
                                            <a-col :span="4">
                                                <a-input 
                                                    size="large"
                                                    v-model="citem.title"
                                                    placeholder="内容标题" 
                                                />
                                            </a-col>
                                            <a-col :span="3">
                                                <a-select 
                                                    style="width: 100%;"
                                                    size="large"
                                                    placeholder="是否平台内跳转" 
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
                                        v-if="item.style == webWidget.community" 
                                        class="setting-item" 
                                        :gutter="[{md:20},{md:20}]">
                                        <a-col :span="12">
                                            <a-input 
                                                size="large"
                                                v-model="item.topicIds" 
                                                placeholder="请输话题帖子id多个请英文逗号分开,例如:1,2,4" 
                                            />
                                        </a-col>
                                        <a-col :span="12">
                                            <a-input 
                                                size="large"
                                                v-model="item.questionIds" 
                                                placeholder="请输问题帖子id多个请英文逗号分开,例如:1,2,4" 
                                            />
                                        </a-col>
                                    </a-row>
                                    
                                    <!-- 会员组件数据 -->
                                    <a-row 
                                        v-if="item.style == webWidget.vip" 
                                        class="setting-item" 
                                        :gutter="[{md:20},{md:20}]">
                                        <a-col :span="24">
                                            <a-input 
                                                size="large"
                                                v-model="item.vipIds" 
                                                placeholder="请输会员id多个请英文逗号分开,例如:1,2,4" 
                                            />
                                        </a-col>
                                    </a-row>

                                    <!-- 文章组件数据 -->
                                    <a-row 
                                        v-if="item.style == webWidget.article" 
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
                                        v-if="item.style == webWidget.edu" 
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

                                    <!-- 搜索组件背景图 -->
                                    <a-row 
                                        v-if="item.style == webWidget.search" 
                                        class="setting-item" 
                                        :gutter="[{md:20},{md:20}]">
                                        <a-col :span="24">
                                            <a-input 
                                                size="large"
                                                v-model="item.backgroundImage" 
                                                placeholder="请输入背景图链接" 
                                            />
                                        </a-col>
                                    </a-row>

                                    <!-- 自定义组件数据内 -->
                                    <a-row 
                                        v-if="item.style == 12"
                                        class="setting-item" 
                                        :gutter="[{md:20},{md:20}]">
                                        <a-col :span="24">
                                            <a-button @click="addCustom(index)" size="large" type="primary" block>
                                                添加自定义内容
                                            </a-button>
                                        </a-col>
                                    </a-row>
                                    <div v-if="item.style == 12">
                                        <a-row 
                                            v-for="(citem,cindex) in item.list" :key="cindex"
                                            class="setting-item" 
                                            :gutter="[{md:20},{md:20}]">
                                            <a-col :span="3">
                                                <a-input 
                                                    size="large"
                                                    v-model="citem.title"
                                                    placeholder="标题" 
                                                />
                                            </a-col>
                                            <a-col :span="3">
                                                <a-input 
                                                    size="large"
                                                    v-model="citem.desc"
                                                    placeholder="描述" 
                                                />
                                            </a-col>
                                            <a-col :span="3">
                                                <a-input 
                                                    size="large"
                                                    v-model="citem.bgColor"
                                                    placeholder="背景色" 
                                                />
                                            </a-col>
                                            <a-col :span="3">
                                                <a-input 
                                                    size="large"
                                                    v-model="citem.fontBgColor"
                                                    placeholder="字体背景色" 
                                                />
                                            </a-col>
                                            <a-col :span="3">
                                                <a-input 
                                                    size="large"
                                                    v-model="citem.fontColor"
                                                    placeholder="字体阴影色" 
                                                />
                                            </a-col>
                                            <a-col :span="3">
                                                <a-input 
                                                    size="large"
                                                    v-model="citem.link"
                                                    placeholder="跳转链接" 
                                                />
                                            </a-col>
                                            <a-col :span="3">
                                                <a-input 
                                                    size="large"
                                                    v-model="citem.icon"
                                                    placeholder="图标" 
                                                />
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
                                        v-if="item.style == webWidget.image" 
                                        class="setting-item" 
                                        :gutter="[{md:20},{md:20}]">
                                        <a-col :span="8">
                                            <a-input 
                                                size="large"
                                                v-model="item.content.cover" 
                                                placeholder="图片地址" 
                                            />
                                        </a-col>
                                        <a-col :span="8">
                                            <a-input 
                                                size="large"
                                                v-model="item.content.link" 
                                                :placeholder="item.isPlatform == 2 ? '请输入平台地址,无需携带域名开头' : '请输入携带完整域名地址'" 
                                            />
                                        </a-col>
                                        <a-col :span="8">
                                            <a-select 
                                                style="width: 100%;"
                                                size="large"
                                                placeholder="是否平台内跳转" 
                                                v-model="item.content.isPlatform">
                                                <a-select-option :value="1">
                                                    否
                                                </a-select-option>
                                                <a-select-option :value="2">
                                                    是
                                                </a-select-option>
                                            </a-select>
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
                    <h2>其他组件</h2>
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
import { webWidget } from "@/shared/widget"
export default {
    components: {
        Vuedraggable,
    },
    data(){
        return{
            webWidget,
            list: [],
            widgetList:[
                {
                    title:"轮播",
                    desc:"",
                    showTitle:undefined,
                    height:360,
                    style:webWidget.swpie,
                    list:[
                        {
                            link:undefined,
                            title:undefined,
                            cover:undefined
                        }
                    ],
                    isOpen:false
                },
                {
                    title:"音频组件",
                    desc:"",
                    showTitle:2,
                    module:"audio",
                    audioIds:"",
                    style:webWidget.audio,
                    isOpen:false
                },
                {
                    title:"视频组件",
                    desc:"",
                    showTitle:2,
                    module:"audio",
                    videoIds:"",
                    style:webWidget.video,
                    isOpen:false
                },
                {
                    title:"资源组件",
                    desc:"",
                    showTitle:2,
                    module:"audio",
                    resourceIds:"",
                    style:webWidget.resource,
                    isOpen:false
                },
                {
                    title:"社区组件",
                    desc:"",
                    showTitle:2,
                    style:webWidget.community,
                    topicIds:"",
                    questionIds:"",
                    isOpen:false
                },
                {
                    title:"会员组件",
                    desc:"",
                    showTitle:2,
                    style:webWidget.vip,
                    vipIds:"",
                    isOpen:false
                },
                {
                    title:"文章组件",
                    desc:"",
                    showTitle:2,
                    style:webWidget.article,
                    articleIds:"",
                    isOpen:false
                },
                {
                    title:"课程组件",
                    desc:"",
                    showTitle:2,
                    style:webWidget.edu,
                    eduIds:"",
                    isOpen:false
                },
                {
                    title:"搜索组件",
                    desc:"",
                    showTitle:2,
                    backgroundImage:"",
                    style:webWidget.search,
                    isOpen:false
                },
                {
                    title:"自定义内容组件",
                    desc:"",
                    showTitle:2,
                    list:[
                        {
                            title:"213",
                            desc:"123",
                            color:"123",
                            links:"123",
                            icon:"123",
                        }
                    ],
                    style:12,
                    isOpen:false
                },
                {
                    title:"自定义图片组件",
                    desc:"",
                    showTitle:2,
                    style:webWidget.image,
                    content:{
                        cover:"",
                        link:"",
                        isPlatform:undefined
                    },
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
            const res = await Info("HomeDesign")  
            if (res.data.info != null) {
                const value = JSON.parse(res.data.info)
                this.list = value.map((item)=>{
                    if (item.content != null && item.content != "") {
                         console.log(item.content)
                        item.content = JSON.parse(item.content)
                    }
    
                    const tmp = {
                        title:item.title,
                        showTitle:item.showTitle,
                        height:item.height,
                        backgroundImage:item.backgroundImage,
                        list:item.list,
                        vipIds:item.vipIds,
                        eduIds:item.eduIds,
                        articleIds:item.articleIds,
                        audioIds:item.audioIds,
                        videoIds:item.videoIds,
                        resourceIds:item.resourceIds,
                        topicIds:item.topicIds,
                        questionIds:item.questionIds,
                        style:item.style,
                        content:item.content,
                        isOpen:false
                    }
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
                            module:item.module,
                            vipIds:item.vipIds,
                            eduIds:item.eduIds,
                            audioIds:item.audioIds,
                            videoIds:item.videoIds,
                            resourceIds:item.resourceIds,
                            articleIds:item.articleIds,
                            topicIds:item.topicIds,
                            questionIds:item.questionIds,
                            style:item.style,
                            content:item.content,
                            isOpen:false
                        }
                        if (item.content != null && item.content != "") {
                            tmp.content = JSON.stringify(tmp.content)
                        }
                        return tmp
                    })
                    const config = {
                        "configValue":JSON.stringify(list),
                        "configKey": "HomeDesign",
                        "configName": "首页设计",
                        "remark" : "首页设计"
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
                title:"",
                cover:"",
                isPlatform: undefined,
            })
        },
        removeSwiper(e,ce){

            this.list[e].list.splice(ce,1)
            // this.list[e].list.push({text:""})
        },
        addCustom(e){
            this.list[e].list.push({text:""})
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
