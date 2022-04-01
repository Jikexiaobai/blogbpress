<template>
    <div class="select-box">
        <a-form-model ref="selectForm" :model="selectForm" :rules="selectForm.rules" layout="inline">
            <a-form-model-item ref="title" prop="title">
                <a-input
                    size="large"
                    type="text"
                    placeholder="请输入标题"
                    v-model="selectForm.title"
                >
                </a-input>
            </a-form-model-item>
            <a-form-model-item ref="status" prop="status">
               <a-select v-model="selectForm.isTop" 
                    :style="{width:200 + 'px'}"
                    size="large"
                    placeholder="请选择是否置顶">
                    <a-select-option  :value="1">
                        不置顶
                    </a-select-option>
                    <a-select-option  :value="2">
                        置顶
                    </a-select-option>
                </a-select>
            </a-form-model-item>
            <a-form-model-item ref="status" prop="status">
               <a-select v-model="selectForm.status" 
                    :style="{width:200 + 'px'}"
                    size="large"
                    placeholder="请选择状态">
                    <a-select-option  :value="1">
                        未审核
                    </a-select-option>
                    <a-select-option  :value="2">
                        已通过
                    </a-select-option>
                    <a-select-option  :value="3">
                        未通过
                    </a-select-option>
                    <a-select-option  :value="4">
                        回收站
                    </a-select-option>
                </a-select>
            </a-form-model-item>
            <a-form-model-item>
                <a-button v-action="'/topic/list'" @click="reset"  size="large" style="margin-left: 10px;">
                    重置
                </a-button>
                <a-button v-action="'/topic/list'" @click="select"  size="large" type="primary" style="margin-left: 10px;">
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
                status:undefined,
                isTop:undefined,
                title:undefined,
            }
        }
    },
    methods:{
        reset(){
            this.selectForm.status = undefined
            this.selectForm.isTop = undefined
            this.selectForm.title = undefined
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