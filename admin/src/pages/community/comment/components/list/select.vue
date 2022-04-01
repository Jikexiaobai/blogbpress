<template>
    <div class="select-box">
        <a-form-model ref="selectForm" :model="selectForm" :rules="selectForm.rules" layout="inline">
            <a-form-model-item ref="content" prop="content">
                <a-input
                    size="large"
                    type="text"
                    placeholder="内容"
                    v-model="selectForm.content"
                >
                </a-input>
            </a-form-model-item>
            <a-form-model-item ref="module" prop="module">
                <a-select 
                    :style="{width:200 + 'px'}"
                    size="large"
                    v-model="selectForm.module" 
                    placeholder="所属模块">
                    <a-select-option value="article">
                        文章
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
                    <a-select-option value="audio">
                        音频
                    </a-select-option>
                    <a-select-option value="topic">
                        帖子
                    </a-select-option>
                    <a-select-option value="answer">
                        回答
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
                <a-button v-action="'/comment/list'" @click="reset"  size="large" style="margin-left: 10px;">
                    重置
                </a-button>
                <a-button v-action="'/comment/list'" @click="select"  size="large" type="primary" style="margin-left: 10px;">
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
                module:undefined,
                content:undefined,
            }
        }
    },
    methods:{
        reset(){
            this.selectForm.status = undefined
            this.selectForm.module = undefined
            this.selectForm.content = undefined
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