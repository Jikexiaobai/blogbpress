<template>
    <a-card :bordered="false">
        <div class="box">
            <a-form-model 
            ref="form" :model="form" :rules="form.rules">
                
                <!-- 昵称和密码 -->
                <a-row :gutter="[40,0]">
                    <a-col :span="12">
                         <a-form-model-item 
                            label="用户名"
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
                            label="手机号"
                            ref="phone" 
                            prop="phone">
                            <a-input
                                size="large"
                                placeholder="请输入手机号"
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

                <!-- 用户头像和封面 -->
                <a-row :gutter="[40,0]">
                    <a-col :span="12">
                        <a-form-model-item 
                            label="头像"
                            ref="avatar" 
                            prop="avatar">
                            <div class="upload-icon">
                                <a-input
                                    size="large"
                                    type="text"
                                    placeholder="请上传头像"
                                    v-model="form.avatar"
                                />
                                <a-button v-action="'/media/upload'" class="form-button" @click="uploadAvatar" size="large" type="primary" style="margin-left: 10px;">
                                    上传
                                </a-button>
                            </div>
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="12">
                        <a-form-model-item 
                            label="封面"
                            ref="cover" 
                            prop="cover">
                            <div class="upload-icon">
                                <a-input
                                    size="large"
                                    type="text"
                                    placeholder="请上传封面"
                                    v-model="form.cover"
                                />
                                <a-button v-action="'/media/upload'" class="form-button" @click="uploadCover" size="large" type="primary" style="margin-left: 10px;">
                                    上传
                                </a-button>
                            </div>
                        </a-form-model-item>
                    </a-col>
                </a-row>

                <!-- 性别 余额 积分 -->
                <a-row :gutter="[40,0]">
                    <a-col :span="8">
                        <a-form-model-item 
                            label="余额"
                            ref="balance" 
                            prop="balance">
                            <a-input-number
                                size="large" 
                                v-model="form.balance"
                                placeholder="请输入余额"
                                :style="{ width: '100%' }"
                                :min="0"
                                :precision="2"
                            />
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="8">
                       <a-form-model-item 
                            label="积分"
                            ref="integral" 
                            prop="integral">
                            <a-input-number
                                size="large" 
                                v-model="form.integral"
                                placeholder="请输入积分"
                                :style="{ width: '100%' }"
                                :min="0"
                            />
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
                </a-row>

                <!-- 用户关注数粉丝数，点赞数 -->
                <a-row :gutter="[40,0]">
                    <a-col :span="8">
                        <a-form-model-item 
                            label="关注数"
                            ref="follows" 
                            prop="follows">
                            <a-input-number
                                size="large" 
                                v-model="form.follows"
                                placeholder="请输入关注数"
                                :style="{ width: '100%' }"
                                :min="0"
                            />
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="8">
                        <a-form-model-item 
                            label="粉丝数"
                            ref="fans" 
                            prop="fans">
                            <a-input-number
                                size="large" 
                                v-model="form.fans"
                                placeholder="请输入粉丝数"
                                :style="{ width: '100%' }"
                                :min="0"
                            />
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="8">
                       <a-form-model-item 
                            label="获赞数"
                            ref="likes" 
                            prop="likes">
                            <a-input-number
                                size="large" 
                                v-model="form.likes"
                                placeholder="请输入获赞数"
                                :style="{ width: '100%' }"
                                :min="0"
                            />
                        </a-form-model-item>
                    </a-col>
                </a-row>

                <!-- 等级会员 -->
                <a-row :gutter="[40,0]">
                    <a-col :span="12">
                        <a-form-model-item 
                            label="等级"
                            ref="grade" 
                            prop="grade">
                            <a-select v-model="form.grade" 
                                size="large"
                                placeholder="请设置等级">
                                <a-select-option v-for="(item,index) in gradeList" :key="index+'grade'" :value="item.id">
                                    {{item.title}}
                                </a-select-option>
                            </a-select>
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="12">
                        <a-form-model-item 
                            label="会员"
                            ref="vip" 
                            prop="vip">
                            <a-select v-model="form.vip" 
                                size="large"
                                placeholder="请设置会员">
                                <a-select-option :value="0">
                                    无
                                </a-select-option>
                                <a-select-option v-for="(item,index) in vipList" :key="index+'vip'" :value="item.id">
                                    {{item.title}}
                                </a-select-option>
                            </a-select>
                        </a-form-model-item>
                    </a-col>
                </a-row>


                <a-form-model-item 
                    label="描述"
                    ref="description" 
                    prop="description">
                    <a-textarea
                        :maxLength="140"
                        :rows="5"
                        size="large"
                        type="text"
                        placeholder="请输入用户描述"
                        v-model="form.description"
                    />
                </a-form-model-item>

                <a-form-model-item>
                    <a-button @click="onSubmit" size="large" type="primary" >
                        提交
                    </a-button>
                </a-form-model-item>
            </a-form-model>
        </div>
    </a-card>
</template>

<script>
import { Create } from '@/api/user'
import { List as vipList } from '@/api/vip'
import { List as gradeList } from '@/api/grade'
export default {
    data() {
        return {
            form:{
                nickName:undefined,
                password:undefined,
                phone:undefined,
                email:undefined,
                avatar:undefined,
                cover:undefined,
                balance:undefined,
                integral:undefined,
                sex:undefined,
                follows:undefined,
                fans:undefined,
                likes:undefined,
                grade:undefined,
                vip:undefined,
                description:undefined,
                rules:{
                    nickName:[
                        { required: true, message: this.$t('user.form.nickName.rule.required'), trigger: 'blur' },
                    ],
                    password:[
                        { required: true, message: this.$t('user.form.password.rule.required'), trigger: 'blur' },
                    ],
                },
            },
            gradeList:[],
            vipList:[],
        }
    },
    mounted(){
        this.getData()
    },
    methods: {
        async getData(){
            try {
                const vipRes = await vipList()
                const gradeRes  = await gradeList()
                this.gradeList = gradeRes.data.list == null ? [] : gradeRes.data.list
                this.vipList = vipRes.data.list == null ? [] : vipRes.data.list
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

        uploadCover(){
            this.$Upload().then((res)=>{
                if (res != false) {
                    this.form.cover = res
                   
                }
            }).catch((err)=>{
                console.log(err)
            })
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
                title: "正在创建",
                onOk:() => {
                    this.$refs.form.validate(valid => {
                        if (valid) {
                        
                            if ((this.form.email == "" || this.form.email == undefined || this.form.email == null ) && (this.form.phone == "" || this.form.phone == undefined || this.form.phone == null) ) {
                                this.$message.error("请输入邮箱或手机号");
                                return
                            }
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
                this.$router.push({path: '/user/list'});
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
