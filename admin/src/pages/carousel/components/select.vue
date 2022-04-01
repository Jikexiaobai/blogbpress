<template>
    <div class="select-box">
        <a-form-model ref="selectForm" :model="selectForm" :rules="selectForm.rules" layout="inline">
            <a-form-model-item ref="mode" prop="mode">
                <a-select v-model="selectForm.mode" 
                    :style="{width:200 + 'px'}"
                    size="large"
                    placeholder="请设置幻灯模式">
                    <a-select-option  :value="1">
                        平台内容
                    </a-select-option>
                    <a-select-option  :value="2">
                        外部内容
                    </a-select-option>
                </a-select>
            </a-form-model-item>
            <a-form-model-item ref="type" prop="type">
                <a-select v-model="selectForm.type" 
                    :style="{width:200 + 'px'}"
                    size="large"
                    placeholder="请设置幻灯类型">
                    <a-select-option  :value="1">
                        Web
                    </a-select-option>
                    <a-select-option  :value="2">
                        App
                    </a-select-option>
                </a-select>
            </a-form-model-item>
            <a-form-model-item>
                <a-button v-action="'/carousel/list'" @click="reset" style="margin-left: 10px;">
                    重置
                </a-button>
                <a-button v-action="'/carousel/list'" @click="select" type="primary" style="margin-left: 10px;">
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
                mode:undefined,
                type:undefined,
            }
        }
    },
    mounted(){
    },
    methods:{

        reset(){
            this.selectForm.mode = undefined
            this.selectForm.type = undefined
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