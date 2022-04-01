<template>
    <div class="select-box">
        <a-form-model ref="selectForm" :model="selectForm" :rules="selectForm.rules" layout="inline">
            <a-form-model-item ref="title" prop="title">
                <a-input
                    size="large"
                    type="text"
                    placeholder="请输入分类标题"
                    v-model="selectForm.title"
                >
                </a-input>
            </a-form-model-item>
            <a-form-model-item ref="module" prop="module">
                <a-select 
                    :style="{ width: 200 + 'px' }"
                    size="large"
                    v-model="selectForm.module" 
                    placeholder="请选择所属模块">
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
            <a-form-model-item>
                <a-button v-action="'/category/list'" @click="reset" style="margin-left: 10px;">
                    重置
                </a-button>
                <a-button v-action="'/category/list'" @click="select" type="primary" style="margin-left: 10px;">
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
                title:undefined,
                module:undefined,
                rules:{},
            }
        }
    },
    methods:{
        reset(){
            this.selectForm.title = undefined
            this.selectForm.module = undefined
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