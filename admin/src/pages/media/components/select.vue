<template>
    <div class="select-box">
        <a-form-model ref="selectForm" :model="selectForm" :rules="selectForm.rules" layout="inline">
            <a-form-model-item ref="title" prop="title">
                <a-input
                    size="large"
                    type="text"
                    placeholder="请输入原名称"
                    v-model="selectForm.orName"
                >
                </a-input>
            </a-form-model-item>
            <a-form-model-item ref="status" prop="status">
                <a-select 
                    :style="{ width: 200 + 'px' }"
                    size="large"
                    v-model="selectForm.status" 
                    placeholder="请选择状态">
                    <a-select-option :value="1">
                        未使用
                    </a-select-option>
                    <a-select-option :value="2">
                        已使用
                    </a-select-option>
                </a-select>
            </a-form-model-item>
            <a-form-model-item>
                <a-button v-action="'/media/list'" @click="reset" style="margin-left: 10px;">
                    重置
                </a-button>
                <a-button v-action="'/media/list'" @click="select" type="primary" style="margin-left: 10px;">
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
            selectForm:{
                orName:undefined,
                status:undefined,
            }
        }
    },
    methods:{
        reset(){
            this.selectForm.orName = undefined
            this.selectForm.status = undefined
            this.$emit("select",this.selectForm)
        },
        select(){
            
            this.$emit("select",this.selectForm)
        },
    }
}
</script>

<style lang="less" scoped>
.select-box{
    margin-bottom: 10px;
}
</style>