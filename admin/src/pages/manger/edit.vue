<template>
    <a-card :bordered="false">
        <div class="box">
            <a-form-model 
            ref="form" :model="form" :rules="form.rules">
                
                <!-- 用户名和昵称 -->
                <a-row :gutter="[40,0]">
                    <a-col :span="12">
                         <a-form-model-item 
                            label="昵称"
                            ref="nickName" 
                            prop="nickName">
                            <a-input
                                size="large"
                                type="text"
                                placeholder="请输入昵称"
                                v-model="form.nickName"
                            >
                            </a-input>
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="12">
                        <a-form-model-item 
                            label="密码"
                            ref="password" 
                            prop="password">
                            <a-input
                                size="large"
                                placeholder="请输入密码"
                                v-model="form.password"
                            />
                           
                        </a-form-model-item>
                    </a-col>
                </a-row>

                <!-- 邮箱和手机号 -->
                <a-row :gutter="[40,0]">
                    <a-col :span="12">
                        <a-form-model-item 
                            label="手机"
                            ref="phone" 
                            prop="phone">
                            <a-input
                                size="large"
                                placeholder="请输入手机"
                                v-model="form.phone"
                            />
                           
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="12">
                         <a-form-model-item 
                            label="邮箱"
                            ref="email" 
                            prop="email">
                            <a-input
                                size="large"
                                type="text"
                                placeholder="请输入邮箱"
                                v-model="form.email"
                            >
                            </a-input>
                        </a-form-model-item>
                    </a-col>
                </a-row>

                <!-- 用户头像和 -->
                <a-row :gutter="[40,0]">
                    <a-col :span="8">
                        <a-form-model-item 
                            label="头像"
                            ref="avatar" 
                            prop="avatar">
                            <div class="upload-icon">
                                <a-input
                                    size="large"
                                    type="text"
                                    placeholder="头像地址"
                                    v-model="form.avatar"
                                />
                                <a-button  class="form-button" @click="uploadAvatar" size="large" type="primary" style="margin-left: 10px;">
                                    上传
                                </a-button>
                            </div>
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="8">
                        <a-form-model-item 
                            label="性别" 
                            ref="sex" 
                            prop="sex">
                            <a-select v-model="form.sex" 
                                size="large"
                                placeholder="请设置性别">
                                <a-select-option :value="1">
                                    男
                                </a-select-option>
                                <a-select-option :value="2">
                                    女
                                </a-select-option>
                                <a-select-option :value="3">
                                    未知
                                </a-select-option>
                            </a-select>
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="8">
                        <a-form-model-item 
                            label="角色" 
                            ref="roleId" 
                            prop="roleId">
                            <a-select 
                                mode="multiple"
                                v-model="form.roleId" 
                                size="large"
                                placeholder="请设置角色">
                                <a-select-option v-for="(item,index) in list" :key="index" :value="item.id">
                                    {{item.title}}
                                </a-select-option>
                            </a-select>
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
import { List } from '@/api/role'
import { Edit,EditInfo } from '@/api/manger'
export default {
    data() {
        return {
            form:{
                nickName:undefined,
                password:undefined,
                phone:undefined,
                email:undefined,
                avatar:undefined,
                sex:3,
                roleId:undefined,
                rules:{
                    nickName:[
                        { required: true, message: "请输入昵称", trigger: 'blur' },
                    ],
                    password:[
                        { required: true, message: "请输入密码", trigger: 'blur' },
                    ],
                    roleId:[
                        { required: true, message: "请设置角色", trigger: 'blur' },
                    ],
                },
            },
            list:[],
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
                const res = await List()
                if (res.code != 1) {
                    this.$message.error(res.message);
                    return
                }
                this.list = res.data.list != null ? res.data.list : []


                const editInfo = await EditInfo({id:this.id})
                // if (editInfo.code != 1) {
                //     this.$router.push({path:'/404'})
                // }
                this.form = Object.assign(this.form,editInfo.data.info)
            } catch (error) {
                console.log(error)
                setTimeout(() => {
                    this.$notification.error({
                        message: '网络错误',
                        description: "请稍后再试"
                    })
                }, 1000)
            }
        },
        uploadAvatar(){
            this.$Upload().then((res)=>{
                if (res != false) {
                    this.form.avatar = res
                   
                }
            }).catch((err)=>{
                console.log(err)
            })
        },
        onSubmit(e){
            this.$confirm({
                title: "正在修改",
                content: "请认真检查表单信息",
                onOk:() => {
                    this.$refs.form.validate(valid => {
                        if (valid) {
                            if ((this.form.email == "" || this.form.email == undefined || this.form.email == null ) && (this.form.phone == "" || this.form.phone == undefined || this.form.phone == null) ) {
                                this.$message.error("请设置邮箱或手机号");
                                return
                            }
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
                this.$router.push({path: '/manger/list'});
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
