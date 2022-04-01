<template>
    <a-card :bordered="false">
        <div class="box">
            <a-form-model 
            ref="form" :model="form" :rules="form.rules">
                
                <a-row :gutter="[40,0]">
                    <a-col :span="12">
                        <a-form-model-item 
                            label="标题"
                            ref="title" 
                            prop="title">
                            <a-input
                                size="large"
                                placeholder="请输入标题"
                                v-model="form.title"
                            />
                           
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="12">
                        <a-form-model-item 
                            label="作者"
                            ref="userId" 
                            prop="userId">
                            <a-input
                                size="large"
                                type="text"
                                placeholder="请输入用户编号ID"
                                v-model="form.userId"
                            >
                            </a-input>
                        </a-form-model-item>
                    </a-col>
                </a-row>

                <a-row :gutter="[40,0]">
                    <a-col :span="12">
                        <a-form-model-item label="标签" ref="tags" prop="tags">
                            <a-select v-model="form.tags" 
                                :maxTagCount="5" 
                                style="width: 100%;"
                                size="large"
                                placeholder="请设置标签" mode="tags" :token-separators="[',']">
                                <a-select-option v-for="(item,index) in tagList" :key="index" :value="`${item.title}`">
                                    {{item.title}}
                                </a-select-option>
                            </a-select>
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="12">
                        <a-form-model-item 
                            label="封面"
                            ref="cover" 
                            prop="cover">
                            <div class="upload-icon">
                                <a-input
                                    size="large"
                                    type="text"
                                    placeholder="请上传封面"
                                    v-model="form.cover"
                                />
                                <a-button v-action="'/media/upload'" class="form-button" @click="upload" size="large" type="primary" style="margin-left: 10px;">
                                    上传
                                </a-button>
                            </div>
                        </a-form-model-item>
                    </a-col>
                </a-row>

                <a-row :gutter="[40,0]">
                    <a-col :span="12">
                        <a-form-model-item 
                            label="分类"
                            ref="cateId" 
                            prop="cateId">
                            <a-tree-select 
                                v-model="form.cateId"
                                :tree-data="cateList"
                                size="large" 
                                placeholder="请选择分类"
                                tree-default-expand-all
                            />
                        </a-form-model-item>
                    </a-col>
                </a-row>

                <a-row :gutter="[40,0]">
                    <a-col :span="4">
                        <a-form-model-item 
                            label="热度"
                            ref="hots" 
                            prop="hots">
                            <a-input-number
                                size="large" 
                                v-model="form.hots"
                                placeholder="请设置热度"
                                :style="{ width: '100%' }"
                                :min="0"
                            />
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="4">
                       <a-form-model-item 
                            label="收藏数"
                            ref="favorites" 
                            prop="favorites">
                            <a-input-number
                                size="large" 
                                v-model="form.favorites"
                                placeholder="请设置收藏数"
                                :style="{ width: '100%' }"
                                :min="0"
                            />
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="4">
                       <a-form-model-item 
                            label="点赞数"
                            ref="likes" 
                            prop="likes">
                            <a-input-number
                                size="large" 
                                v-model="form.likes"
                                placeholder="请设置收藏数"
                                :style="{ width: '100%' }"
                                :min="0"
                            />
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="6">
                       <a-form-model-item 
                            label="浏览数"
                            ref="views" 
                            prop="views">
                            <a-input-number
                                size="large" 
                                v-model="form.views"
                                placeholder="请设置浏览数"
                                :style="{ width: '100%' }"
                                :min="0"
                            />
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="6">
                       <a-form-model-item 
                            label="已报名人数"
                            ref="joins" 
                            prop="joins">
                            <a-input-number
                                size="large" 
                                v-model="form.joins"
                                placeholder="请输入已报名人数"
                                :style="{ width: '100%' }"
                                :min="0"
                            />
                        </a-form-model-item>
                    </a-col>
                </a-row>


                <a-form-model-item 
                    label="描述"
                    ref="description" 
                    prop="description">
                    <a-textarea
                        :maxLength="140"
                        :rows="5"
                        size="large"
                        type="text"
                        placeholder="请输入描述描述"
                        v-model="form.description"
                    />
                </a-form-model-item>
               
                <a-row :gutter="[40,0]">
                    <a-col :span="6">
                        <a-form-model-item 
                            label="报名方式"
                            ref="joinMode" 
                            prop="joinMode">
                            <a-radio-group 
                                v-model="form.joinMode"
                                size="large">
                                <a-radio-button :value="1">
                                    免费报名
                                </a-radio-button>
                                <a-radio-button :value="2">
                                    付费报名
                                </a-radio-button>
                            </a-radio-group>
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="6" v-if="form.joinMode == 2">
                        <a-form-model-item 
                            label="报名价格"
                            :ref="form.joinMode == 2 ? 'price' : ''" 
                            :prop="form.joinMode == 2 ? 'price' : ''" >
                            <a-input-number
                                size="large" 
                                v-model="form.price"
                                    placeholder="请输入下载价格"
                                    :style="{ width: '30%' }"
                                    :precision="2"
                                    :min="0"
                                />
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="6">
                        <a-form-model-item 
                            label="课程类型"
                            ref="type" 
                            prop="type">
                            <a-radio-group 
                                v-model="form.type"
                                size="large">
                                <a-radio-button :value="1">
                                    线下
                                </a-radio-button>
                                <a-radio-button :value="2">
                                   线上
                                </a-radio-button>
                            </a-radio-group>
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="6">
                        <a-form-model-item 
                            label="最大报名数"
                            ref="max" 
                            prop="max">
                            <a-input-number
                                v-model="form.max"
                                :style="{ width: '100%' }"
                                size="large"
                                placeholder="请输入最大报名人数"
                                :min="0"
                            />
                        </a-form-model-item>
                    </a-col>
                    
                </a-row>


                <!-- 上传课时 -->
                <a-form-model-item ref="section" prop="section">
                    <a-button @click="addSection" size="large">
                        添加章节
                    </a-button>
                    <ul class="section-box">
                        <li v-for="(item,index) in form.section" :key="index">
                            <div class="section-item">
                                <div class="section-ac">
                                    <a-input 
                                        v-if="item.isEdit"
                                        :maxLength="40" 
                                        class="section-title"  
                                        v-model="item.title" 
                                        placeholder="请输入章节标题"  />
                                    <span v-if="!item.isEdit" class="section-title span-title">{{item.title}}</span>
                                    <a-icon class="section-title" :type="item.isEdit ? 'check':'edit'" @click="editSection(index)"/>
                                    <a-icon :type="item.isShow ? 'up':'down'" @click="editSectionShowClass(index)"/>
                                </div>
                                <div>
                                    <a-space>
                                        <a-button @click="addClass(index)" type="link">
                                            添加课时
                                        </a-button>
                                        <a-button @click="removeSection(index)" type="danger">
                                            删除章节
                                        </a-button>
                                    </a-space>
                                </div>
                            </div>
                            <ul v-if="item.isShow && item.children.length > 0" class="class-box">
                                <li v-for="(jitem,jindex) in item.children" :key="jindex">
                                    <div class="class-item">
                                        <div class="class-ac">
                                            <a-input v-if="jitem.isEdit" 
                                                :maxLength="40" 
                                                class="class-title" 
                                                v-model="jitem.title"
                                                placeholder="添加课时标题" 
                                            />
                                            <a-input v-if="jitem.isEdit" 
                                                class="class-title" 
                                                v-model="jitem.link"
                                               placeholder="添加课时地址" 
                                            />
                                            <span v-if="!jitem.isEdit" class="class-title span-title">{{jitem.title}}</span>
                                            <a-icon class="class-title" :type="jitem.isEdit ? 'check':'edit'" 
                                            @click="editClass(index,jindex)"/>
                                            <a-tag v-if="jitem.link != null && form.type == 2" color="#87d068">
                                                已上传视频
                                            </a-tag>
                                            <a-tag v-if="jitem.isWatch && form.type == 2" color="#f50">
                                                已设置试看
                                            </a-tag>
                                        </div>
                                        <div>
                                            <a-space>
                                                <a-button v-if="form.type == 2" @click="setIsWatch(index,jindex)" type="link">
                                                    设置试看
                                                </a-button>
                                                <a-button v-if="form.type == 2" @click="uploadClass(index,jindex)" type="link">
                                                    上传视频
                                                </a-button>
                                                <a-button @click="removeClass(index,jindex)" type="danger">
                                                    删除课时
                                                </a-button>
                                            </a-space>
                                        </div>
                                    </div>
                                </li>
                            </ul>
                        </li>
                    </ul>
                </a-form-model-item>
              
               <a-form-model-item>
                    <TinyMceEditor @writeContent="writeContent"   
                    :valueContont="form.content"/>
                </a-form-model-item>
                
                <a-form-model-item>
                    <a-button v-action="'/edu/edit'" @click="onSubmit" size="large" type="primary" >
                        提交
                    </a-button>
                </a-form-model-item>
            </a-form-model>
        </div>
    </a-card>
</template>

<script>
import { ByModule } from '@/api/category'
import { Hots } from '@/api/tag'
import { Edit,EditInfo } from '@/api/edu'
import { loopCate } from '@/utils/tree/cate'
import { toTree } from '@/utils/tree/tree'
import TinyMceEditor from "@/components/Editor/Tinymce"
export default {
    components: {
        TinyMceEditor
    },
    data() {
        return {
            form:{
                title:undefined,
                userId:undefined,
                cateId:undefined,
                tags:[],
                description:undefined,
                hots:undefined,
                favorites:undefined,
                joins:undefined,
                likes:undefined,
                views:undefined,
                cover:undefined,
                link:undefined,
                price:0,
                joinMode:0,
                max:0,
                type:undefined,
                content:null,
                section:[
                    {
                        isShow:false,
                        isEdit:false,
                        title:"第1章节",
                        children:[
                            {
                                isWatch:false,
                                isEdit:false,
                                title:"第1课时",
                                link:null,
                            }
                        ]
                    }
                ],
                rules:{
                    title:[
                        { required: true, message: "请设置标题", trigger: 'blur' },
                    ],
                    cateId:[
                        { required: true, message: "请设置分类", trigger: 'blur' },
                    ],
                    joinMode:[
                        { required: true, message: "请设置报名方式", trigger: 'blur' },
                    ],
                    cover:[
                        { required: true, message: "请上传封面", trigger: 'blur' },
                    ],
                    max:[
                        { required: true, message: "请设置最大报名人数", trigger: 'blur' },
                    ],
                    price:[
                        { required: true, message: "请设置价格", trigger: 'blur' },
                    ],
                    type:[
                        { required: true, message: "请设置类型", trigger: 'blur' },
                    ],
                },
            },
            tagList:[],
            cateList:[],
            id:null,
        }
    },
    mounted(){
      if (isNaN(this.$route.query.id)) {
          this.$router.push({path:"/404"})
      }
      this.id = parseInt(this.$route.query.id)
      
      this.getData()
    },
    methods: {
        async getData(){
            try {
                const cateRes = await ByModule({"module":"edu"})
                const list =  toTree(cateRes.data.list != null ? cateRes.data.list : [],"cateId","parentId")
                this.cateList = loopCate(list)

                const tagRes = await Hots()
                this.tagList = tagRes.data.list

                const editInfo = await EditInfo({id:this.id})
                if (editInfo.code != 1) {
                    this.$router.push({path:'/404'})
                }
                this.form = Object.assign(this.form,editInfo.data.info)

                this.form.tags = editInfo.data.info.tagList == null ? [] :  editInfo.data.info.tagList.map((itme)=>{
                    return itme.title
                })

                if (editInfo.data.info.section != "") {
                    this.form.section = JSON.parse(editInfo.data.info.section)
                    this.form.section = this.form.section.map((item)=>{
                        const tmp = {
                            isShow:false,
                            isEdit:false,
                            title:item.title,
                            children:[]
                        }
                        tmp.children = item.children.map((jitem)=>{
                            const jtmp = {
                                isWatch:jitem.isWatch,
                                isEdit:false,
                                title:jitem.title,
                                link:jitem.link
                            }
                            return jtmp
                        })
                        return tmp
                    })
                }
                
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
        writeContent(e){
            this.form.content = e
        },
        editSectionShowClass(e){
            this.form.section[e].isShow = !this.form.section[e].isShow
        },
        editSection(e){
            this.form.section[e].isEdit = !this.form.section[e].isEdit
        },
        setIsWatch(e,ce){
             this.form.section[e].children[ce].isWatch = !this.form.section[e].children[ce].isWatch
        },
        editClass(e,ce){
            this.form.section[e].children[ce].isEdit = !this.form.section[e].children[ce].isEdit
        },
        addClass(e){
            this.form.section[e].isShow = true
            const form = {
                isWatch:false,
                isEdit:false,
                link:null,
                title:`第${this.form.section[e].children.length + 1}节课`,
            }
            if (this.form.section[e].children.length > 0) {
               this.form.section[e].children.unshift(form)
            } else {
                this.form.section[e].children.push(form)
            }
        },
        removeClass(e,ce){
           this.form.section[e].children.splice(ce,1)
        },
        addSection(){
            const form = {
                isShow:false,
                isEdit:false,
                title:`第${this.form.section.length + 1}章节`,
                children:[]
            }
            if (this.form.section.length > 0) {
                this.form.section.unshift(form)
            } else {
                this.form.section.push(form)
            }
        },
        removeSection(i){
           this.form.section.splice(i,1)
        },
        uploadClass(e,ce){
            this.$Upload("Video").then((res)=>{
                if (res != false) {
                    this.form.section[e].children[ce].link = res
                }
            }).catch((err)=>{
                this.form.section[e].children[ce].link = null
            })
        },
        upload(){
            this.$Upload().then((res)=>{
                if (res != false) {
                    console.log(res)
                    this.form.cover = res
                   
                }
            }).catch((err)=>{
                console.log(err)
            })
        },
        onSubmit(e){
            this.$confirm({
                content: "正在修改",
                onOk:() => {
                    this.$refs.form.validate(valid => {
                        if (valid) {
                        
                            if (this.form.section.length < 1) {
                                this.$message.error(
                                    "请创建章节",
                                    3
                                )
                                return
                            }

                            if (this.form.tags.length > 5) {
                                this.$message.error(
                                    "标签最多为5个",
                                    3
                                )
                                return
                            }

                            if (this.form.content == null || this.form.content == "") {
                                this.$message.error(
                                    "请填写详细介绍",
                                    3
                                )
                                return
                            }

                            if (this.form.joinMode != 2) {
                                this.form.price = 0
                            }

                             this.form.section = this.form.section.map((item)=>{
                                const tmp = {
                                    title:item.title,
                                    children:[]
                                }
                                tmp.children = item.children.map((jitem)=>{
                                    const jtmp = {
                                        isWatch:jitem.isWatch,
                                        title:jitem.title,
                                        link:jitem.link
                                    }
                                    return jtmp
                                })
                                return tmp
                            })
                        
                            let formData = {
                                title:undefined,
                                userId:undefined,
                                cateId:undefined,
                                tags:[],
                                description:undefined,
                                hots:undefined,
                                favorites:undefined,
                                likes:undefined,
                                views:undefined,
                                cover:undefined,
                            }

                            formData = Object.assign(formData,this.form)
                            formData.eduId = this.id
                            this.postEdit(formData)
                        } else {
                            return false;
                        }
                    });
                },
                onCancel() {},
            });
        },
        async postEdit(formData){
            try {
                const res = await Edit(formData)
                 if (res.code != 1) {
                    this.$message.error(
                        res.message,
                        3
                    )
                    return
                }
                this.$router.push({ path: "/edu/list"})
            } catch (error) {
                setTimeout(() => {
                    this.$notification.error({
                        message: '网络错误',
                        description: "请稍后再试"
                    })
                }, 1000)
            }
        },
    },
}
</script>

<style scoped lang="less">
.upload-icon{
    display: flex;
}
.section-box{
    margin-top: 20px;
    li{
        width: 100%;
        margin-bottom: 10px;
        .section-item{
            display: flex;
            align-items: center;
            justify-content: space-between;
            .section-ac{
                display: flex;
                align-items: center;
                .section-title{
                    margin-right: 10px;
                }
                .span-title{
                    font-size: 22px;
                    font-weight: bold;
                }
            }
        }
        .class-box{
            margin-left: 40px;
            background: #f2f4f5;
            
            .class-item{
                display: flex;
                align-items: center;
                justify-content: space-between;
                padding: 10px;
                .class-ac{
                    flex: 1;
                    display: flex;
                    align-items: center;
                    .class-title{
                        margin-right: 10px;
                    }
                    .span-title{
                        font-size: 17px;
                        font-weight: bold;
                    }
                }
            }
        }
    }
}
</style>
