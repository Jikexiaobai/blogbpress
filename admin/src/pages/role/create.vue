<template>
    <a-card :bordered="false">
        <div class="box">
            <a-form-model 
            ref="form" :model="form" :rules="form.rules">
                
                <a-row :gutter="[40,0]">
                    <a-col :span="12">
                        <a-form-model-item 
                            label="角色标题"
                            ref="title" 
                            prop="title">
                            <a-input
                                size="large"
                                v-model="form.title"
                            />
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="12">
                        <a-form-model-item 
                            label="角色状态"
                            ref="status" 
                            prop="status">
                            <a-select size="large" v-model="form.status">
                                <a-select-option :value="1">
                                    禁用
                                </a-select-option>
                                <a-select-option :value="2">
                                    启用
                                </a-select-option>
                            </a-select>
                        </a-form-model-item>
                    </a-col>
                </a-row>

                <a-form-model-item 
                    label="角色权限"
                    ref="authority" 
                    prop="authority">
                    <a-tree-select
                        size="large"
                        tree-checkable
                        show-checked-strategy="SHOW_ALL"
                        v-model="form.authority"
                        style="width: 100%"
                        :tree-data="tree"
                        >
                    </a-tree-select>
                </a-form-model-item>

                <a-form-model-item>
                    <a-button @click="onSubmit" size="large" type="primary" >
                        提交
                    </a-button>
                    <!-- <a-button @click="onSubmit" size="large" style="margin-left: 10px;">
                        {{$t('empty')}}
                    </a-button> -->
                </a-form-model-item>
            </a-form-model>
        </div>
    </a-card>
</template>

<script>
import { Create } from '@/api/role'
import { List } from '@/api/authority'
import { loopMenu } from '@/utils/tree/menu'
export default {
    data() {
        return {
            form:{
                title:undefined,
                status:2,
                authority:undefined,
                rules:{
                    title:[
                        { required: true, message: "请输入标题", trigger: 'blur' },
                    ],
                },
            },
            tree:[],
        }
    },
    mounted(){
        this.getData()
    },
    methods: {
        async getData(){
            try {
                const res = await List(this.query)

                if (res.code != 1) {
                    this.$message.error(res.message);
                    return
                }
                let menuLoot = loopMenu(res.data.list == null ? [] : res.data.list)
                
                this.tree = menuLoot
            } catch (error) {
                
                setTimeout(() => {
                    this.$notification.error({
                        message: '网络错误',
                        description: "请稍后再试"
                    })
                }, 1000)
            }
        },
        onSubmit(e){
            this.$confirm({
                okText: "确定",
                cancelText: "取消",
                title: "正在创建",
                onOk:() => {
                    this.$refs.form.validate(valid => {
                        if (valid) {
                            this.postCreate(this.form)
                        } else {
                            return false;
                        }
                    });
                },
                onCancel() {},
            });
        },
        async postCreate(formData){
            try {
                const res = await Create(formData)
                if (res.code != 1) {
                   this.$message.error(res.message);
                   return
                }
                this.$router.push({path: '/role/list'});
            } catch (error) {
                setTimeout(() => {
                    this.$notification.error({
                        message: '网络错误',
                        description: "请稍后再试"
                    })
                }, 1000)
            }
        },
    },
}
</script>

<style scoped lang="less">
.upload-icon{
    display: flex;
    
}
.create-purpose{
    margin-bottom: 10px;
    .create-purpose-ac{
        display: flex;
        justify-content: flex-end;
        align-items: center;
    }
}
</style>
