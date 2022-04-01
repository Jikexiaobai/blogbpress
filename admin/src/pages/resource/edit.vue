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
                                placeholder="请输入用户Id"
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
                                placeholder="请设置分类"
                                tree-default-expand-all
                            />
                        </a-form-model-item>
                    </a-col>
                </a-row>

                <a-row :gutter="[40,0]">
                    <a-col :span="6">
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
                    <a-col :span="6">
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
                    <a-col :span="6">
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
                    <a-col :span="12">
                        <a-form-model-item 
                        label="下载模块" 
                        ref="hasDown" prop="hasDown">
                            <a-select v-model="form.hasDown" 
                                style="width: 100%;"
                                size="large"
                                placeholder="请设置是否添加下载模块">
                                <a-select-option :value="1">
                                    否
                                </a-select-option>
                                <a-select-option :value="2">
                                    是
                                </a-select-option>
                            </a-select>
                        </a-form-model-item>
                    </a-col>
                </a-row>


                <!-- 下载权限 -->
                <a-form-model-item 
                    v-if="form.hasDown == 2" 
                    label="下载方式"
                    class="perms-radio" ref="downMode" prop="downMode">
                    <a-radio-group size="large"  
                        v-model="form.downMode" 
                        button-style="solid">
                        <a-radio-button :value="0">
                            免费下载
                        </a-radio-button>
                        <a-radio-button :value="1">
                            支付下载
                        </a-radio-button>
                        <a-radio-button :value="2">
                            评论下载
                        </a-radio-button>
                        <a-radio-button :value="3">
                            登录下载
                        </a-radio-button>
                    </a-radio-group>
                </a-form-model-item>

                <a-form-model-item 
                v-if="form.downMode == 1"
                class="price" ref="price" prop="price">
                    <a-input-number
                    size="large" 
                    v-model="form.price"
                        placeholder="请输入下载价格"
                        :style="{ width: '30%' }"
                        :precision="2"
                        :min="0"
                    />
                </a-form-model-item>

                <a-form-model-item 
                v-if="form.hasDown == 2" 
                ref="downUrl" 
                prop="downUrl">
                    <a-row v-for="(item,index) in form.downUrl" :key="index" :gutter="[40,0]" 
                    class="create-purpose">
                        <a-col :span="7">
                            <a-input
                            v-model="item.title" 
                            :maxLength="40" size="large"  
                            placeholder="请输入平台" />
                        </a-col>
                        <a-col :span="7">
                            <a-input 
                            v-model="item.key" 
                            size="large"  
                            placeholder="请输入下载链接" />
                        </a-col>
                        <a-col :span="7">
                            <a-input v-model="item.val" 
                            :maxLength="40" size="large"  
                            placeholder="请输入提取码" />
                        </a-col>
                        <a-col :span="3" class="create-purpose-ac">
                            <a-button @click="removeDownUrl(index)" size="large" type="danger">
                                删除
                            </a-button>
                        </a-col>
                    </a-row>
                    <a-button @click="addDownUrl" size="large" type="primary" block>
                        添加下载地址
                    </a-button>
                </a-form-model-item>

                <a-form-model-item 
                v-if="form.hasDown == 2" 
                ref="purpose" prop="purpose">
                    <a-row v-for="(item,index) in form.purpose" 
                    :key="index" :gutter="[40,0]" 
                    class="create-purpose">
                        <a-col :span="10">
                            <a-input v-model="item.key" size="large"  
                            placeholder="请输入用途类型" />
                        </a-col>
                        <a-col :span="10">
                            <a-input v-model="item.val" 
                            :maxLength="40" size="large" 
                            placeholder="请输入用途" />
                        </a-col>
                        <a-col :span="4" class="create-purpose-ac">
                            <a-button @click="removePurpose(index)" 
                            size="large" type="danger">
                                删除
                            </a-button>
                        </a-col>
                    </a-row>
                    <a-button @click="addPurpose" size="large" type="primary" block>
                        添加用途
                    </a-button>
                </a-form-model-item>

                <a-form-model-item 
                v-if="form.hasDown == 2" 
                ref="attribute" 
                prop="attribute">
                    <a-row v-for="(item,index) in form.attribute" :key="index" :gutter="[40,0]" class="create-purpose">
                        <a-col :span="10">
                            <a-input v-model="item.key" 
                            :maxLength="40" size="large"  
                            placeholder="请输入属性标题" />
                        </a-col>
                        <a-col :span="10">
                            <a-input v-model="item.val" :maxLength="40" size="large" 
                             placeholder="请输入属性值" />
                        </a-col>
                        <a-col :span="4" class="create-purpose-ac">
                            <a-button @click="removeAttribute(index)" size="large" type="danger">
                                删除
                            </a-button>
                        </a-col>
                    </a-row>
                    <a-button @click="addAttribute" size="large" type="primary" block>
                        添加属性
                    </a-button>
                </a-form-model-item>

                <a-form-model-item>
                    <TinyMceEditor @writeContent="writeContent"   
                    :valueContont="form.content"/>
                </a-form-model-item>
                
                <a-form-model-item>
                    <a-button @click="onSubmit" size="large" type="primary" >
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
import { Edit,EditInfo } from '@/api/resource'
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
                likes:undefined,
                views:undefined,
                cover:undefined,
                content:undefined,
                hasDown:1,
                price:0,
                downMode:0,
                downUrl:[],
                attribute:[],
                purpose:[],
                rules:{
                    title:[
                        { required: true, message: "请输入标题", trigger: 'blur' },
                    ],
                    cateId:[
                        { required: true, message: "请设置分类", trigger: 'blur' },
                    ],
                    cover:[
                        { required: true, message: "请设置封面", trigger: 'blur' },
                    ],
                    hasDown:[
                        { required: true, message: "请设置是否有下载模块", trigger: 'blur' },
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
                const cateRes = await ByModule({"module":"resource"})
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

                if (editInfo.data.info.purpose != "") {
                    this.form.purpose = JSON.parse(editInfo.data.info.purpose)
                }
                
                if (editInfo.data.info.attribute != "") {
                    this.form.attribute = JSON.parse(editInfo.data.info.attribute)
                }

                if (editInfo.data.info.downUrl != "") {
                    this.form.downUrl = JSON.parse(editInfo.data.info.downUrl)
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
        addDownUrl(){
            const form = {
                title:undefined,
                key:undefined,
                val:undefined
            }
            this.form.downUrl.push(form)
        },
        removeDownUrl(i){
           this.form.downUrl.splice(i,1)
        },
        addPurpose(){
            const form = {
                key:undefined,
                val:undefined
            }
            this.form.purpose.push(form)
        },
        removePurpose(i){
           this.form.purpose.splice(i,1)
        },
        addAttribute(){
            const form = {
                key:undefined,
                val:undefined
            }
            this.form.attribute.push(form)
           
        },
        removeAttribute(i){
            this.form.attribute.splice(i,1)
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
                title: "正在修改",
                onOk:() => {
                    this.$refs.form.validate(valid => {
                        if (valid) {
                        
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
                                content:null,
                            }

                            formData = Object.assign(formData,this.form)
                            formData.resourceId = this.id
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
                   this.$message.error(res.message);
                   return
                }
                this.$router.push({path: '/resource/list'});
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
.create-purpose{
    margin-bottom: 10px;
    .create-purpose-ac{
        display: flex;
        justify-content: flex-end;
        align-items: center;
    }
}
</style>
