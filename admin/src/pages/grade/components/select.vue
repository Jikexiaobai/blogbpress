<template>
    <div class="select-box">
        <a-form-model ref="selectForm" :model="selectForm" :rules="selectForm.rules" layout="inline">
            <a-form-model-item ref="title" prop="title">
                <a-input
                    size="large"
                    type="text"
                    placeholder="等级标题"
                    v-model="selectForm.title"
                >
                </a-input>
            </a-form-model-item>
            <a-form-model-item ref="type" prop="type">
                <a-select 
                    :style="{ width: 200 + 'px' }"
                    size="large"
                    v-model="selectForm.type" 
                    placeholder="等级类型">
                    <a-select-option :value="1">
                        {{$t('role.form.type.one.text')}}
                    </a-select-option>
                    <a-select-option :value="2">
                        {{$t('role.form.type.two.text')}}
                    </a-select-option>
                    <a-select-option :value="3">
                        {{$t('role.form.type.three.text')}}
                    </a-select-option>
                </a-select>
            </a-form-model-item>
            <a-form-model-item>
                <a-button @click="empty" style="margin-left: 10px;">
                    {{$t('empty')}}
                </a-button>
                <a-button @click="select" type="primary" style="margin-left: 10px;">
                    {{$t('select')}}
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
                type:undefined,
            }
        }
    },
    methods:{
        empty(){
            this.selectForm.title = undefined
            this.selectForm.type = undefined
            this.$emit("select",this.selectForm)
        },
        select(){
            console.log("qweqw")
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