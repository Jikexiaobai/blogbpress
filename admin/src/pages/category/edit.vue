<template>
    <a-card :bordered="false">
        <div class="box">
            <a-form-model ref="form" :model="form" :rules="form.rules">
                <a-row :gutter="[40,0]">
                    <a-col :span="12">
                        <a-form-model-item 
                            ref="title" 
                            prop="title">
                                <a-input
                                    size="large"
                                    type="text"
                                    placeholder="请输入分类标题"
                                    v-model="form.title"
                                />
                            </a-form-model-item>
                    </a-col>
                    <a-col :span="12">
                        <a-form-model-item ref="module" prop="module">
                            <a-select 
                                    @change="changeData"
                                    size="large"
                                    v-model="form.module" 
                                    placeholder="请设置所属模块">
                                    <a-select-option value="article">
                                        文章
                                    </a-select-option>
                                    <a-select-option value="group">
                                        圈子
                                    </a-select-option>
                                    <a-select-option value="edu">
                                        课程
                                    </a-select-option>
                                    <a-select-option value="video">
                                        视频
                                    </a-select-option>
                                    <a-select-option value="resource">
                                        资源
                                    </a-select-option>
                                    <a-select-option value="audio">
                                        音频
                                    </a-select-option>
                                </a-select>
                            </a-form-model-item>
                    </a-col>
                </a-row>
                
                <a-row :gutter="[40,0]">
                    <a-col :span="12">
                        <a-form-model-item 
                        ref="parentId" prop="parentId">
                            <a-tree-select 
                                v-model="form.parentId"
                                :tree-data="list"
                                size="large" 
                                placeholder="请输入父级分类"
                                tree-default-expand-all
                            />
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="12">
                        <a-form-model-item 
                        ref="slug" prop="slug">
                            <a-input
                                size="large"
                                type="text"
                                placeholder="请输入分类别名"
                                v-model="form.slug"
                            >
                            </a-input>
                        </a-form-model-item>
                    </a-col>
                </a-row>
                
                <a-row :gutter="[40,0]">
                    <a-col :span="12">
                        <a-form-model-item ref="module" prop="module">
                            <a-select 
                                    @change="getData"
                                    size="large"
                                    v-model="form.isTop" 
                                    placeholder="请设置分类是否推荐">
                                    <a-select-option :value="1">
                                        否
                                    </a-select-option>
                                    <a-select-option :value="2">
                                        是
                                    </a-select-option>
                                </a-select>
                            </a-form-model-item>
                    </a-col>
                    <a-col :span="12">
                        <a-form-model-item
                            ref="description" prop="description">
                            <a-input
                                size="large"
                                type="text"
                                placeholder="请输入分类描述"
                                v-model="form.description"
                            >
                            </a-input>
                        </a-form-model-item>
                    </a-col>
                </a-row>
                
                
                
                <a-form-model-item 
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

                <a-form-model-item>
                    <a-button v-action="'/category/edit'" @click="onSubmit" icon="plus" type="primary" >
                        提交
                    </a-button>
                </a-form-model-item>
            </a-form-model>
        </div>
    </a-card>
</template>

<script>
import { ByModule,EditInfo,Edit  } from '@/api/category'
import { loopCate } from '@/utils/tree/cate'
import { toTree } from '@/utils/tree/tree'
export default {
    data() {
        return {
            list:[],
            form:{
                title:undefined,
                parentId:undefined,
                slug:undefined,
                isTop:undefined,
                description:undefined,
                module:undefined,
                cover:undefined,
                rules:{
                    title:[
                        { required: true, message: "请输入分类标题", trigger: 'blur' },
                    ],
                    slug:[
                        { required: true, message: "请输入分类别名", trigger: 'blur' },
                    ],
                    module:[
                        { required: true, message: "请设置所属模块", trigger: 'blur' },
                    ],
                },
            },
            id:undefined,
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
                const editInfo = await EditInfo({id:this.id})
                if (editInfo.code != 1) {
                    this.$router.push({path:'/404'})
                }
                this.form = Object.assign(this.form,editInfo.data.info)
                if (editInfo.data.info.parentId == 0) {
                   this.form.parentId = undefined
                }
                this.changeData()
            } catch (error) {
                setTimeout(() => {
                    this.$notification.error({
                        message: '网络错误',
                        description: "请稍后再试"
                    })
                }, 1000)
            }
        },
        async changeData(){
            try {
                const res = await ByModule({"module":this.form.module})
                const list =  toTree(res.data.list != null ? res.data.list : [],"cateId","parentId")
                this.list = loopCate(list)

            } catch (error) {
                setTimeout(() => {
                    this.$notification.error({
                        message: '网络错误',
                        description: "请稍后再试"
                    })
                }, 1000)
            }
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
                title: "正在创建分类",
                content: "请认真检查表单信息",
                onOk:() => {
                    this.$refs.form.validate(valid => {
                        if (valid) {
                            this.form.id = this.id
                            this.postEdit(this.form)
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
                this.$router.push({path: '/category/list'});
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
</style>
