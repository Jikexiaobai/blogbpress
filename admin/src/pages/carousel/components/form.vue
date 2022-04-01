<template>
    <div class="box">
        <a-form-model ref="form" :model="form" :rules="form.rules">
            <a-form-model-item 
            ref="link" 
            prop="link">
                <a-input
                    size="large"
                    type="text"
                    placeholder="请输入链接可携带https"
                    v-model="form.link"
                />
            </a-form-model-item>

            <a-form-model-item 
                ref="cover" 
                prop="cover">
                <div class="upload-icon">
                    <a-input
                        size="large"
                        type="text"
                        placeholder="请设置封面 "
                        v-model="form.cover"
                    />
                    <a-button v-action="'/media/upload'" class="form-button" @click="upload" size="large" type="primary" style="margin-left: 10px;">
                        上传
                    </a-button>
                </div>
            </a-form-model-item>

            <a-form-model-item ref="type" prop="type">
               <a-select 
                    size="large"
                    v-model="form.type" 
                    placeholder="请选择类型">
                    <a-select-option :value="1">
                        Web
                    </a-select-option>
                    <a-select-option :value="2">
                        App
                    </a-select-option>     
                </a-select>
            </a-form-model-item>

            <a-form-model-item ref="mode" prop="mode">
               <a-select 
                    size="large"
                    v-model="form.mode" 
                    placeholder="请选择模式">
                    <a-select-option :value="1">
                        平台内容
                    </a-select-option>
                    <a-select-option :value="2">
                        外部内容
                    </a-select-option>     
                </a-select>
            </a-form-model-item>

            <a-row :gutter="[40,0]" v-if="form.mode == 1">
                <a-col :span="12">
                    <a-form-model-item 
                    :ref="form.mode == 1 ? 'module':''" 
                    :prop="form.mode == 1 ? 'module':''">
                        <a-select 
                            size="large"
                            v-model="form.module" 
                            placeholder="请设置模块">
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
                            <a-select-option value="auido">
                                音频
                            </a-select-option>
                            
                        </a-select>
                    </a-form-model-item>
                </a-col>
                <a-col :span="12">
                    <a-form-model-item 
                        :ref="form.mode == 1 ? 'relatedId':''" 
                        :prop="form.mode == 1 ? 'relatedId':''">
                        <a-input
                            size="large"
                            type="text"
                            placeholder="请输入模块ID"
                            v-model="form.relatedId"
                        />
                    </a-form-model-item>
                </a-col>
            </a-row>
            

            <a-form-model-item>
                <a-button v-action="'/carousel/create'" @click="onSubmit" icon="plus" type="primary" >
                    提交
                </a-button>
            </a-form-model-item>
        </a-form-model>
    </div>
</template>

<script>
import { Create } from '@/api/carousel'

export default {
    data() {
        return {
            form:{
                link:undefined,
                type:undefined,
                cover:undefined,
                mode:undefined,
                module:undefined,
                relatedId:undefined,
                rules:{
                    link:[
                        { required: true, message: "请设置链接", trigger: 'blur' },
                    ],
                    type:[
                        { required: true, message: "请设置类型", trigger: 'blur' },
                    ],
                    cover:[
                        { required: true, message: "请设置封面", trigger: 'blur' },
                    ],
                    mode:[
                        { required: true, message: "请设置方式", trigger: 'blur' },
                    ],
                    module:[
                        { required: true, message: "请设置模块", trigger: 'blur' },
                    ],
                    relatedId:[
                        { required: true, message: "请设置关联Id", trigger: 'blur' },
                    ],
                },
            }
        }
    },
    methods: {
        onSubmit(e){
            this.$confirm({
                title: "正在创建",
                onOk:() => {
                    this.$refs.form.validate(valid => {
                        if (valid) {
                            this.postCreate(this.form)
                        } else {
                            return false;
                        }
                    });
                },
                onCancel() {},
            });
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
        async postCreate(formData){
            try {
                const res = await Create(formData)
                if (res.code != 1) {
                   this.$message.error(res.message);
                   return
                }
                this.form.link = undefined
                this.form.type = undefined
                this.form.cover = undefined
                this.form.mode = undefined
                this.form.module = undefined
                this.form.relatedId = undefined
                this.$emit("getData")
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
