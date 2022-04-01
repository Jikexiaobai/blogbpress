<template>
    <div class="select-box">
        <a-form-model ref="selectForm" :model="selectForm" :rules="selectForm.rules" layout="inline">
            <a-form-model-item ref="secretKey" prop="secretKey">
                <a-input
                    size="large"
                    type="text"
                    placeholder="密钥"
                    v-model="selectForm.secretKey"
                >
                </a-input>
            </a-form-model-item>
            <a-form-model-item ref="status" prop="status">
                <a-select v-model="selectForm.status" 
                    :style="{width:200 + 'px'}"
                    size="large"
                    placeholder="状态">
                    <a-select-option  :value="1">
                        未使用
                    </a-select-option>
                    <a-select-option  :value="2">
                        已使用
                    </a-select-option>
                </a-select>
            </a-form-model-item>
            <a-form-model-item>
                <a-button v-action="'/card/list'" @click="reset" style="margin-left: 10px;">
                    重置
                </a-button>
                <a-button v-action="'/card/list'" @click="select" type="primary" style="margin-left: 10px;">
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
                secretKey:undefined,
            }
        }
    },
    mounted(){
    },
    methods:{
        reset(){
            this.selectForm.status = undefined
            this.selectForm.secretKey = undefined
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