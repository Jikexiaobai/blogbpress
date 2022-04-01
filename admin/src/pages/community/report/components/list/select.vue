<template>
    <div class="select-box">
        <a-form-model ref="selectForm" :model="selectForm" :rules="selectForm.rules" layout="inline">
            <a-form-model-item ref="type" prop="type">
               <a-select v-model="selectForm.type" 
                    :style="{width:200 + 'px'}"
                    size="large"
                    placeholder="类型">
                    <a-select-option  :value="1">
                        广告垃圾
                    </a-select-option>
                    <a-select-option  :value="2">
                        违规内容
                    </a-select-option>
                    <a-select-option  :value="3">
                        恶意灌水
                    </a-select-option>
                </a-select>
            </a-form-model-item>
            <a-form-model-item ref="status" prop="status">
               <a-select v-model="selectForm.status" 
                    :style="{width:200 + 'px'}"
                    size="large"
                    placeholder="状态">
                    <a-select-option  :value="1">
                        待处理
                    </a-select-option>
                    <a-select-option  :value="2">
                        已处理
                    </a-select-option>
                </a-select>
            </a-form-model-item>
            <a-form-model-item>
                <a-button v-action="'/report/list'" @click="reset"  size="large" style="margin-left: 10px;">
                    重置
                </a-button>
                <a-button v-action="'/report/list'" @click="select"  size="large" type="primary" style="margin-left: 10px;">
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
                type:undefined,
                status:undefined,
            }
        }
    },
    methods:{
        reset(){
            this.selectForm.type = undefined
            this.selectForm.status = undefined
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