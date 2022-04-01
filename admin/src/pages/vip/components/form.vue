<template>
    <div class="box">
        <a-form-model 
            ref="form" :model="form" :rules="form.rules">
            <a-form-model-item
                ref="title" 
                prop="title">
                <a-input
                    size="large"
                    type="text"
                    placeholder="请输入会员标题"
                    v-model="form.title"
                >
                </a-input>
            </a-form-model-item>
            <a-form-model-item 
                ref="day" 
                prop="day">
                <a-input-number
                    size="large" 
                    v-model="form.day"
                    placeholder="过期天数"
                    :style="{ width: '100%' }"
                    :min="0"
                />
            </a-form-model-item>
            <a-form-model-item 
                ref="price" 
                prop="price">
                <a-input-number
                    size="large" 
                    v-model="form.price"
                    placeholder="开通价格"
                    :style="{ width: '100%' }"
                    :min="0"
                    :precision="2"
                />
            </a-form-model-item>
            <a-form-model-item 
                ref="discount" 
                prop="discount">
                <a-input-number
                    size="large" 
                    v-model="form.discount"
                    placeholder="会员折扣"
                    :style="{ width: '100%' }"
                    :min="0"
                    :max="0.99"
                    :precision="2"
                />
            </a-form-model-item>
            <a-form-model-item 
                ref="color" 
                prop="color">
                <a-input
                    size="large"
                    type="text"
                    placeholder="自定义颜色"
                    v-model="form.color"
                >
                </a-input>
            </a-form-model-item>
            <a-form-model-item 
                ref="icon" 
                prop="icon">
                <div class="upload-icon">
                    <a-input
                        size="large"
                        type="text"
                        placeholder="会员图标"
                        v-model="form.icon"
                    />
                    <a-button v-action="'/media/upload'" class="form-button" @click="upload" size="large" type="primary" style="margin-left: 10px;">
                        上传
                    </a-button>
                </div>
            </a-form-model-item>


            <a-form-model-item>
                <a-button v-action="'/vip/create'" size="large" type="primary" >
                    提交
                </a-button>
            </a-form-model-item>
        </a-form-model>
    </div>
</template>

<script>
import { Create } from '@/api/vip'

export default {
    data() {
        return {
            form:{
                title:undefined,
                day:undefined,
                price:undefined,
                discount:undefined,
                color:undefined,
                icon:undefined,
                rules:{
                    title:[
                        { required: true, message: "请输入会员标题", trigger: 'blur' },
                    ],
                    price:[
                        { required: true, message: "请设置开通价格", trigger: 'blur' },
                    ],
                    icon:[
                        { required: true, message: "请上传图标", trigger: 'blur' },
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
                    this.form.icon = res
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
                this.form.title = undefined
                this.form.day = undefined
                this.form.price = undefined
                this.form.discount = undefined
                this.form.color = undefined
                this.form.icon = undefined
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

