<template>
    <div class="select-box">
        <a-form-model ref="selectForm" :model="selectForm" :rules="selectForm.rules" layout="inline">
            <a-form-model-item ref="nickName" prop="nickName">
                <a-input
                    size="large"
                    type="text"
                    placeholder="请输入昵称"
                    v-model="selectForm.nickName"
                >
                </a-input>
            </a-form-model-item>
            <a-form-model-item ref="email" prop="email">
                <a-input
                    size="large"
                    type="text"
                    placeholder="请输入邮箱"
                    v-model="selectForm.email"
                >
                </a-input>
            </a-form-model-item>
            <a-form-model-item ref="phone" prop="phone">
                <a-input
                    size="large"
                    type="text"
                    placeholder="请输入手机号"
                    v-model="selectForm.phone"
                >
                </a-input>
            </a-form-model-item>
            <a-form-model-item ref="status" prop="status">
                <a-select v-model="selectForm.status" 
                    :style="{width:200 + 'px'}"
                    size="large"
                    placeholder="请选择状态">
                    <a-select-option  :value="1">
                        禁用
                    </a-select-option>
                    <a-select-option  :value="2">
                        启用
                    </a-select-option>
                </a-select>
            </a-form-model-item>
            <a-form-model-item>
                <a-button v-action="'/manger/list'" @click="reset" style="margin-left: 10px;">
                    重置
                </a-button>
                <a-button v-action="'/manger/list'" @click="select" type="primary" style="margin-left: 10px;">
                    查询
                </a-button>
            </a-form-model-item>
        </a-form-model>
    </div>
</template>

<script>
export default {
    data(){
        return{
            cateList:[],
            selectForm:{
                status:undefined,
                nickName:undefined,
                email:undefined,
                phone:undefined,
            }
        }
    },
    mounted(){
    },
    methods:{
        reset(){
            this.selectForm.status = undefined
            this.selectForm.nickName = undefined
            this.selectForm.email = undefined
            this.selectForm.phone = undefined
            this.$emit("select",this.selectForm)
        },
        select(){
            this.$emit("select",this.selectForm)
        }
    }
}
</script>

<style lang="less" scoped>
.select-box{
    margin-bottom: 10px;
}
</style>