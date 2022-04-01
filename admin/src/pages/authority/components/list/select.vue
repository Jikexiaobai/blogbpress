<template>
    <div class="select-box">
        <a-form-model ref="selectForm" :model="selectForm" :rules="selectForm.rules" layout="inline">
            <a-form-model-item ref="title" prop="title">
                <a-input
                    size="large"
                    type="text"
                    placeholder="请输入权限标题"
                    v-model="selectForm.title"
                    />
            </a-form-model-item>
            <a-form-model-item>
                <a-button v-action="'/authority/create'" @click="add" style="margin-left: 10px;">
                    新建
                </a-button>
                <a-button v-action="'/authority/list'" @click="select" type="primary" style="margin-left: 10px;">
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
            }
        }
    },
    methods:{
        add(){
            this.$router.push({path: '/authority/create'});
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