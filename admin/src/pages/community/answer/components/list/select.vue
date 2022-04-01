<template>
    <div class="select-box">
        <a-form-model ref="selectForm" :model="selectForm" :rules="selectForm.rules" layout="inline">
            <a-form-model-item ref="userId" prop="userId">
                <a-input
                    size="large"
                    type="text"
                    placeholder="用户编号"
                    v-model="selectForm.userId"
                >
                </a-input>
            </a-form-model-item>
            <a-form-model-item ref="isPay" prop="isPay">
               <a-select v-model="selectForm.isPay" 
                    :style="{width:200 + 'px'}"
                    size="large"
                    placeholder="是否付费查看">
                    <a-select-option  :value="1">
                        否
                    </a-select-option>
                    <a-select-option  :value="2">
                        是
                    </a-select-option>
                </a-select>
            </a-form-model-item>
            <a-form-model-item ref="status" prop="status">
               <a-select v-model="selectForm.status" 
                    :style="{width:200 + 'px'}"
                    size="large"
                    placeholder="状态">
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
                <a-button v-action="'/answer/list'" @click="reset"  size="large" style="margin-left: 10px;">
                    重置
                </a-button>
                <a-button v-action="'/answer/list'" @click="select"  size="large" type="primary" style="margin-left: 10px;">
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
                isPay:undefined,
                userId:undefined,
            }
        }
    },
    methods:{
        reset(){
            this.selectForm.status = undefined
            this.selectForm.isPay = undefined
            this.selectForm.userId = undefined
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