<template>
    <a-card :bordered="false">
        <div class="box">
            <a-form-model 
            ref="form" :model="form" :rules="form.rules">
                
                <!-- 编号和下单人 -->
                <a-row :gutter="[40,0]">
                    <a-col :span="12">
                         <a-form-model-item 
                            :label="$t('order.form.orderNum.label')"
                            ref="orderNum" 
                            prop="orderNum">
                            <a-input
                                disabled
                                size="large"
                                type="text"
                                v-model="form.orderNum"
                            >
                            </a-input>
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="12">
                        <a-form-model-item 
                            :label="$t('order.form.user.label')"
                            ref="nickName" 
                            prop="nickName">
                            <a-input
                                disabled
                                size="large"
                                v-model="form.nickName"
                            />
                           
                        </a-form-model-item>
                    </a-col>
                </a-row>

                <!-- 邮箱和手机号 -->
                <a-row :gutter="[40,0]">
                    <a-col :span="12">
                        <a-form-model-item 
                            :label="$t('order.form.phone.label')"
                            ref="phone" 
                            prop="phone">
                            <a-input
                                size="large"
                                :placeholder="$t('order.form.phone.placeholder')"
                                v-model="form.phone"
                            />
                           
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="12">
                         <a-form-model-item 
                            :label="$t('order.form.email.label')"
                            ref="email" 
                            prop="email">
                            <a-input
                                size="large"
                                type="text"
                                :placeholder="$t('order.form.email.placeholder')"
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
                            :label="$t('order.form.avatar.label')"
                            ref="avatar" 
                            prop="avatar">
                            <div class="upload-icon">
                                <a-input
                                    size="large"
                                    type="text"
                                    :placeholder="$t('order.form.cover.placeholder')"
                                    v-model="form.avatar"
                                />
                                <a-button class="form-button" @click="uploadAvatar" size="large" type="primary" style="margin-left: 10px;">
                                    {{$t('upload')}}
                                </a-button>
                            </div>
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="12">
                        <a-form-model-item 
                            :label="$t('order.form.cover.label')"
                            ref="cover" 
                            prop="cover">
                            <div class="upload-icon">
                                <a-input
                                    size="large"
                                    type="text"
                                    :placeholder="$t('order.form.cover.placeholder')"
                                    v-model="form.cover"
                                />
                                <a-button class="form-button" @click="uploadCover" size="large" type="primary" style="margin-left: 10px;">
                                    {{$t('upload')}}
                                </a-button>
                            </div>
                        </a-form-model-item>
                    </a-col>
                </a-row>

                <!-- 性别 余额 积分 -->
                <a-row :gutter="[40,0]">
                    <a-col :span="8">
                        <a-form-model-item 
                            :label="$t('order.form.balance.label')"
                            ref="balance" 
                            prop="balance">
                            <a-input-number
                                size="large" 
                                v-model="form.balance"
                                :placeholder="$t('order.form.balance.placeholder')"
                                :style="{ width: '100%' }"
                                :min="0"
                                :precision="2"
                            />
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="8">
                       <a-form-model-item 
                            :label="$t('order.form.integral.label')"
                            ref="integral" 
                            prop="integral">
                            <a-input-number
                                size="large" 
                                v-model="form.integral"
                                :placeholder="$t('order.form.integral.placeholder')"
                                :style="{ width: '100%' }"
                                :min="0"
                            />
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="8">
                        <a-form-model-item 
                        :label="$t('order.form.sex.label')" 
                            ref="sex" 
                            prop="sex">
                            <a-select v-model="form.sex" 
                                size="large"
                                :placeholder="$t('order.form.sex.placeholder')">
                                <a-select-option :value="1">
                                    {{$t('order.form.sex.one.label')}}
                                </a-select-option>
                                <a-select-option :value="2">
                                    {{$t('order.form.sex.tow.label')}}
                                </a-select-option>
                                <a-select-option :value="3">
                                    {{$t('order.form.sex.three.label')}}
                                </a-select-option>
                            </a-select>
                        </a-form-model-item>
                    </a-col>
                </a-row>

                <!-- 用户关注数粉丝数，点赞数 -->
                <a-row :gutter="[40,0]">
                    <a-col :span="8">
                        <a-form-model-item 
                            :label="$t('order.form.follows.label')"
                            ref="follows" 
                            prop="follows">
                            <a-input-number
                                size="large" 
                                v-model="form.follows"
                                :placeholder="$t('order.form.follows.placeholder')"
                                :style="{ width: '100%' }"
                                :min="0"
                            />
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="8">
                        <a-form-model-item 
                            :label="$t('order.form.fans.label')"
                            ref="fans" 
                            prop="fans">
                            <a-input-number
                                size="large" 
                                v-model="form.fans"
                                :placeholder="$t('order.form.fans.placeholder')"
                                :style="{ width: '100%' }"
                                :min="0"
                            />
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="8">
                       <a-form-model-item 
                            :label="$t('order.form.likes.label')"
                            ref="likes" 
                            prop="likes">
                            <a-input-number
                                size="large" 
                                v-model="form.likes"
                                :placeholder="$t('order.form.likes.placeholder')"
                                :style="{ width: '100%' }"
                                :min="0"
                            />
                        </a-form-model-item>
                    </a-col>
                </a-row>

                <!-- 等级会员 -->
                <a-row :gutter="[40,0]">
                    <a-col :span="8">
                        <a-form-model-item 
                            :label="$t('order.form.admin.label')"
                            ref="admin" 
                            prop="admin">
                            <a-select v-model="form.admin" 
                                size="large"
                                :placeholder="$t('order.form.admin.placeholder')">
                                <a-select-option :value="0">
                                    {{$t('order.form.role.placeholder')}}
                                </a-select-option>
                                <a-select-option v-for="(item,index) in adminList" :key="index+'admin'" :value="item.roleId">
                                    {{item.title}}
                                </a-select-option>
                            </a-select>
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="8">
                        <a-form-model-item 
                            :label="$t('order.form.grade.label')"
                            ref="grade" 
                            prop="grade">
                            <a-select v-model="form.grade" 
                                size="large"
                                :placeholder="$t('order.form.grade.placeholder')">
                                <a-select-option :value="0">
                                    {{$t('order.form.role.placeholder')}}
                                </a-select-option>
                                <a-select-option v-for="(item,index) in gradeList" :key="index+'grade'" :value="item.roleId">
                                    {{item.title}}
                                </a-select-option>
                            </a-select>
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="8">
                        <a-form-model-item 
                            :label="$t('order.form.vip.label')"
                            ref="vip" 
                            prop="vip">
                            <a-select v-model="form.vip" 
                                size="large"
                                :placeholder="$t('order.form.vip.placeholder')">
                                <a-select-option :value="0">
                                    {{$t('order.form.role.placeholder')}}
                                </a-select-option>
                                <a-select-option v-for="(item,index) in vipList" :key="index+'vip'" :value="item.roleId">
                                    {{item.title}}
                                </a-select-option>
                            </a-select>
                        </a-form-model-item>
                    </a-col>
                </a-row>


                <a-form-model-item 
                    :label="$t('order.form.description.label')"
                    ref="description" 
                    prop="description">
                    <a-textarea
                        :maxLength="140"
                        :rows="5"
                        size="large"
                        type="text"
                        :placeholder="$t('order.form.description.placeholder')"
                        v-model="form.description"
                    />
                </a-form-model-item>

                <a-form-model-item>
                    <a-button @click="onSubmit" size="large" type="primary" >
                        {{$t('submit')}}
                    </a-button>
                    <a-button @click="onSubmit" size="large" style="margin-left: 10px;">
                        {{$t('empty')}}
                    </a-button>
                </a-form-model-item>
            </a-form-model>
        </div>
    </a-card>
</template>

<script>
import { Edit,CreateMeta,EditInfo } from '@/api/user'
export default {
    data() {
        return {
            labelCol: { span: 2},
            wrapperCol: { span: 10 },
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
                admin:undefined,
                grade:undefined,
                vip:undefined,
                description:undefined,
                rules:{
                    nickName:[
                        { required: true, message: this.$t('order.form.nickName.rule.required'), trigger: 'blur' },
                    ],
                    password:[
                        { required: true, message: this.$t('order.form.password.rule.required'), trigger: 'blur' },
                    ],
                },
            },
            adminList:[],
            gradeList:[],
            vipList:[],
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
                const res = await CreateMeta()
                if (res.code != 1) {
                    this.$router.push({path:'/404'})
                }
                this.adminList = res.data.roleList == null ? [] : res.data.roleList.filter((item)=>{
                    return item.type == 1
                }) 
                this.gradeList = res.data.roleList == null ? [] : res.data.roleList.filter((item)=>{
                    return item.type == 3
                }) 
                this.vipList = res.data.roleList == null ? [] : res.data.roleList.filter((item)=>{
                    return item.type == 2
                }) 
                
                const editInfo = await EditInfo({id:this.id})
                if (editInfo.code != 1) {
                    this.$router.push({path:'/404'})
                }
                this.form = Object.assign(this.form,editInfo.data.info)
                if (this.form.admin == 0) {
                    this.form.admin = undefined
                }
                if (this.form.grade == 0) {
                    this.form.grade = undefined
                }
                if (this.form.vip == 0) {
                    this.form.vip = undefined
                }
                if (this.form.sex == 0) {
                    this.form.sex = undefined
                }
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
                okText: this.$t('confirm.ok.text'),
                cancelText:this.$t('confirm.cancel.text'),
                title: this.$t('confirm.create.title.text'),
                content: this.$t('confirm.content.text'),
                onOk:() => {
                    this.$refs.form.validate(valid => {
                        if (valid) {
                            if ((this.form.email == "" || this.form.email == undefined || this.form.email == null ) && (this.form.phone == "" || this.form.phone == undefined || this.form.phone == null) ) {
                                this.$message.error(this.$t('order.form.emailOrphone.rule.required'));
                                return
                            }
                            let formData = {}
                            formData = Object.assign(formData,this.form)
                            formData.orderId = this.id
                            this.postEdit(formData)
                            
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
                this.$router.push({path: '/order/list'});
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
