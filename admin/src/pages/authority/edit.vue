<template>
    <a-card :bordered="false">
        <div class="box">
            <a-form-model 
            ref="form" :model="form" :rules="form.rules">
                
                <a-row :gutter="[40,0]">
                    <a-col :span="12">
                        <a-form-model-item 
                            label="权限标题"
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
                            label="上级菜单"
                            ref="parentId" 
                            prop="parentId">
                            <a-tree-select
                                size="large"
                                v-model="form.parentId"
                                style="width: 100%"
                                :tree-data="tree"
                                >
                            </a-tree-select>
                        </a-form-model-item>
                    </a-col>
                </a-row>

                <a-row :gutter="[40,0]">
                    <a-col :span="12">
                        <a-form-model-item 
                            label="权限类型"
                            ref="type" 
                            prop="type">
                            <a-radio-group v-model="form.type">
                                <a-radio :value="1">
                                    目录
                                </a-radio>
                                <a-radio :value="2">
                                    菜单
                                </a-radio>
                                <a-radio :value="3">
                                    按钮
                                </a-radio>
                            </a-radio-group>
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="12" v-if="form.type == 1">
                        <a-form-model-item 
                            label="重定向"
                            ref="redirect" 
                            prop="redirect">
                            <a-input
                                size="large"
                                v-model="form.redirect"
                            />
                        </a-form-model-item>
                    </a-col>
                </a-row>

                <a-row :gutter="[40,0]" v-if="form.type != 3">
                    <a-col :span="8">
                        <a-form-model-item 
                            label="排序"
                            ref="orderBy" 
                            prop="orderBy">
                            <a-input-number 
                            :precision="0"
                            size="large"
                            v-model="form.orderBy" 
                            style="width: 100%;"
                            :min="1" :max="100" 
                            />
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="8">
                        <a-form-model-item 
                            label="打开方式"
                            ref="target" 
                            prop="target">
                            <a-select size="large" v-model="form.target">
                                <a-select-option :value="1">
                                    原窗口
                                </a-select-option>
                                <a-select-option :value="2">
                                    新窗口
                                </a-select-option>
                            </a-select>
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="8">
                        <a-form-model-item 
                            label="是否隐藏"
                            ref="hidden" 
                            prop="hidden">
                            <a-select size="large" v-model="form.hidden">
                                <a-select-option :value="1">
                                    隐藏
                                </a-select-option>
                                <a-select-option :value="2">
                                    显示
                                </a-select-option>
                            </a-select>
                        </a-form-model-item>
                    </a-col>
                </a-row>

                <a-row :gutter="[40,0]">
                    <a-col :span="12" v-if="form.type != 3">
                        <a-form-model-item 
                            label="路由地址"
                            ref="path" 
                            prop="path">
                            <a-input
                                size="large"
                                v-model="form.path"
                            />
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="12">
                        <a-form-model-item 
                            label="权限标识"
                            ref="perms" 
                            prop="perms">
                            <a-input
                                size="large"
                                v-model="form.perms"
                            />
                        </a-form-model-item>
                    </a-col>
                </a-row>

                <a-row :gutter="[40,0]" v-if="form.type != 3">
                    <a-col :span="12">
                        <a-form-model-item 
                            label="组件"
                            ref="component" 
                            prop="component">
                            <a-input
                                size="large"
                                v-model="form.component"
                            />
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="12">
                        <a-form-model-item 
                            label="图标"
                            ref="icon" 
                            prop="icon">
                            <a-input
                                size="large"
                                v-model="form.icon"
                            />
                        </a-form-model-item>
                    </a-col>
                </a-row>

                
        
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
import { List,EditInfo,Edit } from '@/api/authority'
import { loopMenu } from '@/utils/tree/menu'
export default {
    data() {
        return {
            form:{
                parentId:0,
                title:undefined,
                type:1,
                target:undefined,
                hidden:undefined,
                redirect:undefined,
                orderBy:undefined,
                path:undefined,
                perms:undefined,
                icon:undefined,
                component:undefined,
                rules:{
                    title:[
                        { required: true, message: "请输入标题", trigger: 'blur' },
                    ],
                    parentId:[
                        { required: true, message: "请设置上级菜单", trigger: 'blur' },
                    ],
                    type:[
                        { required: true, message: "请设置权限类型", trigger: 'blur' },
                    ],
                    target:[
                        { required: true, message: "请设置菜单打开方式", trigger: 'blur' },
                    ],
                    hidden:[
                        { required: true, message: "请设置菜单是否隐藏", trigger: 'blur' },
                    ],
                    path:[
                        { required: true, message: "请设置路由地址", trigger: 'blur' },
                    ],
                    perms:[
                        { required: true, message: "请设置权限标识", trigger: 'blur' },
                    ],
                    component:[
                        { required: true, message: "请设置组件", trigger: 'blur' },
                    ],
                },
            },
            tree:[
                {
                    title: '主菜单',
                    value: 0,
                    key: 0,
                    children: [],
                },
            ],
            id:null,
        }
    },
    mounted(){
        if (isNaN(this.$route.query.id)) {
            this.$router.push({path:"/404"})
        }
        this.id = parseInt(this.$route.query.id)
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
                this.tree[0].children = menuLoot

                 const editInfo = await EditInfo({id:this.id})
                if (editInfo.code != 1) {
                    this.$router.push({path:'/404'})
                }
                this.form = Object.assign(this.form,editInfo.data.info)

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
                okText: this.$t('confirm.ok.text'),
                cancelText:this.$t('confirm.cancel.text'),
                title: this.$t('confirm.create.title.text'),
                content: this.$t('confirm.content.text'),
                onOk:() => {
                    this.$refs.form.validate(valid => {
                        if (valid) {

                            this.form.id = this.id
                            this.postEdit(this.form)
                            
                        } else {
                            return false;
                        }
                    });
                },
                onCancel() {},
            });
        },
        async postEdit(formData){
            try {
                const res = await Edit(formData)
                if (res.code != 1) {
                   this.$message.error(res.message);
                   return
                }
                
                this.$router.push({path: '/authority/list'});
         
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