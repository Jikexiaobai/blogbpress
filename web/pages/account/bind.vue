<template>
    <div class="bingding-setting">
        <h2>账户设置</h2>
        <ul class="setting-contet">
            <!-- <li class="setting-item">
                <div class="setting-input-title">
                用户名
                </div>
                <div class="setting-input">
                    {{userName}}
                </div>
            </li> -->
            <li class="setting-item">
                <div class="setting-input-title">
                    登录密码
                </div>
                <div class="setting-input">
                    <span class="setting-input-span">已设置</span>
                    <a-button @click="showPassWordModal" type="dashed">
                        修改
                    </a-button>
                </div>
            </li>
            <li class="setting-item">
                <div class="setting-input-title">
                    验证邮箱
                </div>
                <div class="setting-input">
                    <span v-if="email != ''">{{email}}</span>
                    <span v-else class="setting-input-span">未设置</span>
                    <!-- <a-button @click="showEmailModal" type="dashed">
                       {{email != '' ? "修改" : "设置"}}
                    </a-button> -->
                </div>   
            </li>
            <li class="setting-item">
                <div class="setting-input-title">
                    验证手机
                </div>
                <div class="setting-input">
                     <span v-if="phone != ''">{{phone}}</span>
                    <span v-else class="setting-input-span">未设置</span>
                    <!-- <a-button  v-if="phone == ''" @click="showPhoneModal" type="dashed">
                       {{phone != '' ? "修改" : "设置"}}
                    </a-button> -->
                </div>
            </li>
            <li class="setting-item">
                <div class="setting-input-title">
                    QQ
                </div>
                <div class="setting-input">
                    <span class="setting-input-span">未设置</span>
                    <a-button type="dashed">
                        绑定
                    </a-button>
                </div>
            </li>
            <li class="setting-item">
                <div class="setting-input-title">
                    微信
                </div>
                <div class="setting-input">
                    <span class="setting-input-span">未设置</span>
                    <a-button type="dashed">
                        绑定
                    </a-button>
                </div>
            </li>
        </ul>

        <!-- 密码模态 -->
        <a-modal
            title="修改密码"
            :visible="passWordModal.visible"
            @ok="passWordhandleOk"
            @cancel="passWordhandleCancel"
            okText="确定"
            cancelText="取消"
            class="binding-modal"
            >
            <a-form-model ref="passWordForm" :model="passWordModal" :rules="passWordModal.rules" class="passWordModal_input">
                <a-form-model-item  ref="oldPassWord" prop="oldPassWord">
                    <a-input-password size="large" v-model="passWordModal.oldPassWord" placeholder="旧密码" />
                </a-form-model-item>
                <a-form-model-item  ref="newPassWord" prop="newPassWord">
                    <a-input-password size="large" v-model="passWordModal.newPassWord" placeholder="新密码" />
                </a-form-model-item>
            </a-form-model>
        </a-modal>

        <!-- 邮箱模态 -->
        <a-modal
            title="设置邮箱"
            :visible="emailModal.visible"
            @ok="emailhandleOk"
            @cancel="emailhandleCancel"
            okText="确定"
            cancelText="取消"
            class="binding-modal"
            >
            <a-form-model ref="emailModalForm" :model="emailModal" :rules="emailModal.rules" class="emailModal_input">
                <a-form-model-item  ref="email" prop="email">
                    <a-input size="large" v-model="emailModal.email" placeholder="输入邮箱" />
                </a-form-model-item>
                <a-form-model-item ref="captcha" prop="captcha">
                    <div class="code">
                        <a-input :maxLength="6" class="code_input" size="large" v-model="emailModal.captcha" placeholder="验证码" />
                        <a-button @click="emailsendCode"  type="primary" :disabled="!emailModal.show">
                             {{emailModal.content}}
                        </a-button>
                    </div>
                </a-form-model-item>
            </a-form-model>
            
        </a-modal>

        <!-- 手机模态 -->
        <a-modal
            title="设置手机"
            :visible="phoneModal.visible"
            @ok="phonehandleOk"
            @cancel="phonehandleCancel"
            okText="确定"
            cancelText="取消"
            class="binding-modal"
            >
            <a-form-model ref="phoneModalForm" :model="phoneModal" :rules="phoneModal.rules" class="phoneModal_input">
                <a-form-model-item  ref="phone" prop="phone">
                    <a-input size="large" v-model="phoneModal.phone" placeholder="输入手机">
                    </a-input>
                </a-form-model-item>
                <a-form-model-item ref="captcha" prop="captcha">
                    <div class="code">
                        <a-input :maxLength="6" class="code_input" size="large" v-model="phoneModal.captcha" placeholder="验证码" />
                        <a-button  type="primary" >
                            发送验证码
                        </a-button>
                    </div>
                </a-form-model-item>
            </a-form-model>

        </a-modal>
    </div>
</template>

<script>
import { mapState } from "vuex"
import api from "@/api/index"
export default {
    middleware: 'auth',
    head(){
        return this.$seo(`用户中心-${this.base.title}`,`用户中心`,[{
            hid:"fiber",
            name:"description",
            content:`用户中心`
        }])
    },
    computed:{
        ...mapState(["base"])
    },
    data(){
        return {
            email: null,
            phone: null,
            passWordModal:{
                visible:false,
                oldPassWord: null,
                newPassWord: null,
                rules:{
                    oldPassWord:[
                        { required: true, message: '请输入旧密码', trigger: 'blur' }
                    ],
                    newPassWord:[
                        { required: true, message: '请输入新密码', trigger: 'blur' },
                         { max: 30, message: '最大长度为30', trigger: 'blur'},
                        { min: 6, message: '最小长度为6', trigger: 'blur'}
                    ],
                },
            },
            emailModal:{
                show: true,
                count: 0,
                content: "发送验证码",
                timer: null,
                visible:false,
                email: null,
                captcha: null,
                rules:{
                    email:[
                        { required: true, message: '请输入邮箱', trigger: 'blur' },
                        { pattern:/^[A-Za-z0-9]+([_\.][A-Za-z0-9]+)*@([A-Za-z0-9\-]+\.)+[A-Za-z]{2,6}$/, message: '邮箱格式不正确', trigger: 'blur'}
                    ],
                    // captcha:[
                    //     { required: true, message: '请输入验证码', trigger: 'blur' },
                    // ],
                },
            },
            phoneModal:{
                visible:false,
                phone: null,
                captcha: null,
                rules:{
                    phone:[
                        { required: true, message: '请输入手机号', trigger: 'blur' },
                        {pattern:/^1[3456789]\d{9}$/, message: '手机号码格式不正确', trigger: 'blur'}
                    ]
                },
            }
        }
    },
    mounted(){
        this.getData()
    },
    methods:{
        async getData(){
            try {
                const res = await this.$axios.get(api.getAccountSecurity)
                
                this.email = res.data.info.email
                this.phone = res.data.info.phone
            } catch (error) {
                setTimeout(() => {
                    this.$notification.error({
                        message: '网络错误',
                        description: "请稍后再试"
                    })
                }, 1000)
            }
        },
        // 密码框modal
        showPassWordModal(){
            this.passWordModal.visible = true
        },
        passWordhandleOk(){
            this.$refs.passWordForm.validate(async valid => {
                if (valid) {     
                    try {
                        const formData = {
                            "oldPass":this.passWordModal.oldPassWord,
                            "newPass":this.passWordModal.newPassWord
                        }
                        const res = await this.$axios.post(api.postAccountUpdatePassWord,formData)
                        if (res.code == 1) {
                            this.$message.success(
                                res.message,
                                3
                            )
                            this.passWordModal.visible = false
                            this.passWordModal.oldPassWord = null
                            this.passWordModal.newPassWord = null
                        }else{
                             this.$message.error(
                                res.message,
                                3
                            )
                        }
                    } catch (error) {
                        // console.log(error)
                        setTimeout(() => {
                            this.$notification.error({
                                message: '网络错误',
                                description: "请稍后再试"
                            })
                        }, 1000)
                    }
                } else {
                    console.log('error submit!!');
                    return false;
                }
            });
        },
        passWordhandleCancel(){
            this.passWordModal.visible = false;
        },

        //  邮箱设置框
        showEmailModal(){
            this.emailModal.visible = true
        },
        emailhandleOk(){
            this.$refs.emailModalForm.validate(async valid => {
                if (valid) {     
                    try {
                        const formData = {
                            "email":this.emailModal.email,
                            "captcha":this.emailModal.captcha
                        }
                        const res = await this.$axios.post(api.postAccountUpdateEmail,formData)
                        if (res.code == 1) {
                            this.$message.success(
                                res.message,
                                3
                            )
                            this.passWordModal.visible = false
                            this.passWordModal.oldPassWord = null
                            this.passWordModal.newPassWord = null
                        }else{
                             this.$message.error(
                                res.message,
                                3
                            )
                        }
                    } catch (error) {
                        setTimeout(() => {
                            this.$notification.error({
                                message: '网络错误',
                                description: "请稍后再试"
                            })
                        }, 1000)
                    }
                } else {
                    console.log('error submit!!');
                    return false;
                }
            });
        },
        emailhandleCancel(){
            this.emailModal.visible = false;
        },
        emailsendCode(){
             this.$refs.emailModalForm.validate(async valid => {
                    if (!valid) {
                        // console.log("sdfasdf")
                        return false
                    }
                    this.emailhandleCancel()
                    this.$Code(this.emailModal.email).then((res)=>{
                        if (res != false) {
                            this.showEmailModal()
                            const TIME_COUNT = 60;
                            if (!this.timer) {
                            this.emailModal.count = TIME_COUNT;
                            this.emailModal.show = false;
                            this.emailModal.timer = setInterval(() => {
                                if (this.emailModal.count > 0 && this.emailModal.count <= TIME_COUNT) {
                                    this.emailModal.count--;
                                    this.emailModal.content = `${this.emailModal.count}秒后重发`
                                    } else {
                                    this.emailModal.show = true;
                                    this.emailModal.content = "发送验证码"
                                    clearInterval(this.emailModal.timer);
                                    this.emailModal.timer = null;
                                    }
                                }, 1000)
                            }
                        }
                    }).catch((err)=>{
                        console.log(err)
                        // this.createForm.cover = undefined
                    })
                })
        },

        //  手机设置框
        showPhoneModal(){
            this.phoneModal.visible = true
        },
        phonehandleOk(){

        },
        phonehandleCancel(){
            this.phoneModal.visible = false;
        }

    }
}
</script>

<style lang="less" scoped>
.bingding-setting{
    background-color: #fff;
    padding: 20px;
    h2{
        color: #bcbcbc;
        font-size: 18px;
    }
    .setting-contet{
        margin: 20px 0;
        .setting-item{
            display: flex;
            align-items: center;
            border-bottom: 1px solid #ebeef5;
            padding: 20px;
            .setting-input-title{
                font-weight: 700;
                font-size: 17px;
                width: 100px;
                flex-grow: 0;
            }
            .setting-input{
                width: 100%;
                display: flex;
                justify-content: space-between;
                align-items: center;
                .setting-input-span{
                    color: green;
                }
            }
        }
    } 
}
.binding-modal{
    .code{
        display: flex;
        align-items: center;
        justify-content: space-between;
        .code_input{
            // width: 100px;
            margin-right: 20px;
        }
    }
}
</style>
