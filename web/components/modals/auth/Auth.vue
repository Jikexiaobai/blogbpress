<template>
    <div class="login_box center opacity"
        :class="[isTrue && 'is_back_show']">
        <div class="login_container">
            <a-icon class="login_close" type="close" @click="close"/>
            <a-row>
                <a-col :md="24" :span="24" class="login_right">
                    <!-- 标题 -->
                    <h2>{{title}}</h2>
                    <!-- 登录框 -->
                    <a-form-model v-show="type == 'login'" ref="loginForm" :model="loginForm" :rules="loginForm.rules" class="login_input">
                        <a-form-model-item ref="account" prop="account">
                            <a-input size="large" v-model="loginForm.account" placeholder="输入用户账号">
                                <a-icon slot="prefix" type="user" />
                            </a-input>
                        </a-form-model-item>
                        <a-form-model-item ref="password" prop="password">
                           <a-input-password @pressEnter="onSubmit" size="large" v-model="loginForm.password" placeholder="密码" />
                        </a-form-model-item>
                        <div class="login_register">
                            <p class="forget">
                                忘记密码
                            </p>
                            <p>
                                新用户？<span @click="goRegister" class="go_register">注册</span>
                            </p>
                        </div>
                    </a-form-model>
                    <!-- 注册框 -->
                    <a-form-model v-show="type == 'register'" ref="registerForm" :model="registerForm" :rules="registerForm.rules" class="register_input">
                        <a-form-model-item  
                            ref="account" 
                            prop="account">
                            <a-input size="large" v-model="registerForm.account" 
                            :placeholder="config.registerMode == 'email'?'请输入邮箱':'请输入手机号'">
                                <a-icon slot="prefix" type="user" />
                            </a-input>
                        </a-form-model-item>
                        <a-form-model-item ref="captcha" prop="captcha">
                            <div class="code">
                                <a-input :maxLength="6" class="code_input" size="large" v-model="registerForm.captcha" placeholder="验证码" />
                                <a-button  type="primary" @click="sendCode" :disabled="!registerForm.show">
                                    {{registerForm.content}}
                                </a-button>
                            </div>
                        </a-form-model-item>
                        <a-form-model-item ref="password" prop="password">
                            <a-input-password size="large" v-model="registerForm.password" placeholder="密码" />
                        </a-form-model-item>
                        <div class="go_login" @click="goLogin">
                            已有账号？登录
                        </div>
                    </a-form-model>
                    
                    <a-button @click="onSubmit" size="large" type="primary" block>
                        {{btnTitle}}
                    </a-button>
                    <!-- 社交登录 -->
                    <div v-if="type == 'login' && config.social.length > 0" class="login_open">
                        <div class="login_open_title">社交登录:</div>
                        <ul class="login_open_ul">
                            <li class="wechat">
                                <a-icon type="wechat" />
                                <span>微信</span>
                            </li>
                            <li class="qq">
                                <a-icon type="qq" />
                                <span>QQ</span>
                            </li>
                        </ul>
                    </div>

                    <!-- 协议 -->
                    <div v-if="type == 'register'" class="register_proxy">
                        <span>
                            注册登录即表示同意
                            <a href="/" target="_blank"><span>用户协议</span></a>
                            ,
                            <a href="/" target="_blank"><span>隐私协议</span></a>
                        </span>
                    </div>
                </a-col>
            </a-row>
        </div>
    </div>
</template>

<script>
import {mapActions} from "vuex"
import api from "@/api/index"
// import wsConnection from "@/service/websocket" 
// import router from '../../router'
export default {
    data() {
        return {
            type:"login",
            title:"登录",
            btnTitle:"快速登录",
            // 登录输入框
            loginForm:{
                account:null,
                password:null,
                rules:{
                    account:[
                        { required: true, message: ' 请输入账户', trigger: 'change' },
                    ],
                    password:[
                        { required: true, message: '请输入密码', trigger: 'change' },
                    ],
                },
            },
            registerForm:{
                show: true,
                count: 0,
                content: "发送验证码",
                timer: null,
                account:null,
                captcha:null,
                password:null,
                rules:{
                    account:[
                        // { required: true, message: '请输入邮箱', trigger: 'blur' },
                    ],
                    captcha:[
                        { required: true, message: '请输入验证码', trigger: 'change' },
                    ],
                    password:[
                        { required: true, message: '请输入密码', trigger: 'change' },
                        { min: 6, message: '用户名最小长度为6', trigger: 'change' },
                    ],
                },
            },
            // registerType: 0 email, 1 telephone
            registerType: 0,
            config: {
                registerMode: "",
                policyUrl: "",
                protocolUrl: "",
                social: ""
            },
            isTrue: false,
            state: null // 准备（prepare） 确定（ ascertain） 取消（cancel）
        };
    },
    methods: {
        ...mapActions("user",["A_UPDATE_USER","A_UPDATE_TOKEN"]),
        // ...mapActions("auth",["SendCaptcha","RegisterConfig","ToRegister","ToLogin"]),
        async confirm(
            type = "login",
            title = "登录",
            btnTitle = "快速登录"
        ) {
            this.type = type || this.type;
            this.title = title || this.title;
            this.btnTitle = btnTitle || this.btnTitle;

            const res = await this.$axios.get(api.getOption)     
            this.config.registerMode=res.data.info.registerMode
            this.config.policyUrl=res.data.info.policyUrl
            this.config.protocolUrl=res.data.info.protocolUrl
            this.config.social= res.data.info.social

            if (res.data.info.registerMode == "email") {
                this.registerForm.rules.account.push({
                    pattern:/^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((\.[a-zA-Z0-9_-]{2,3}){1,2})$/,
                    message: '请输入邮箱', trigger: 'blur'
                })
            }else{
                this.registerForm.rules.account.push({
                    pattern:/^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$/,
                    message: '请输入手机号', trigger: 'blur'
                })
            }
            
            // this.config = JSON.parse(data.config)
            this.open();
            return new Promise((resolve, reject) => {
                const target = { state: 'prepare' };
                let res = new Proxy(target, {
                    set(event, key, value) {
                        if (value === 'ascertain') {
                            resolve(true);
                        } else {
                            reject(false);
                        }
                        return true;
                    }
                });
                this.state = res;
            });
        },
        onSubmit(){
            switch(this.type) {
                case "register":
                    this.postRegister()
                    break;
                case "login":
                    this.postLogin()
                    break;
                case "forget":
                    console.log("忘记密码")
                    break;
            } 
        },
        sendCode(){
            this.$refs.registerForm.validateField("account",async (err)=>{
                if (err != "") {
                    return false
                }
                this.$Code(this.registerForm.account).then((res)=>{
                    if (res != false) {
                        const TIME_COUNT = 60;
                        if (!this.timer) {
                        this.registerForm.count = TIME_COUNT;
                        this.registerForm.show = false;
                        this.registerForm.timer = setInterval(() => {
                            if (this.registerForm.count > 0 && this.registerForm.count <= TIME_COUNT) {
                                this.registerForm.count--;
                                this.registerForm.content = `${this.registerForm.count}秒后重发`
                                } else {
                                this.registerForm.show = true;
                                this.registerForm.content = "发送验证码"
                                clearInterval(this.registerForm.timer);
                                this.registerForm.timer = null;
                                }
                            }, 1000)
                        }
                    }
                }).catch((err)=>{
                    console.log(err)
                })
            })
        },
        postRegister(){
             this.$refs.registerForm.validate(async (valid) => {
                try {
                    
                    if (!valid) {
                        return false;
                    } 

                    const formData = {
                        account: this.registerForm.account,
                        captcha: this.registerForm.captcha,
                        password: this.registerForm.password
                    }
                    
                    const res = await this.$axios.post(api.postRegister,formData)
                    console.log(res)
                    if (res.code != 1) {
                        this.$message.error(
                            res.message,
                            3
                        )
                        return
                    }
                   
                    this.loginForm.account = this.registerForm.account
                    this.loginForm.password = this.registerForm.password
                    
                    this.registerForm.account = null  
                    this.registerForm.captcha = null
                    this.registerForm.password = null
               
                    this.goLogin()
                } catch (error) {
                    console.log(error)
                    setTimeout(() => {
                        this.$notification.error({
                            message: '网络错误',
                            description: "请稍后再试"
                        })
                    }, 1000)
                }
            });
        },
        postLogin(){
             this.$refs.loginForm.validate(async (valid) => {
                try {
                    if (!valid) {
                        return false;
                    } 
                    const formData = {
                        "account":this.loginForm.account,
                        "password" : this.loginForm.password
                    }  
                    const {code,message,data} = await this.$axios.post(api.postLogin,formData)
             
                    if (code != 1) {
                        this.$message.error(
                            message,
                            3
                        )
                        return
                    }
                    this.A_UPDATE_TOKEN(data.token)
                    const res = await this.$axios.get(api.getAccountInfo)
                    // userInfo.avatar = res.data.info.avatar
                    // userInfo.nickName = res.data.info.nickName
                    // userInfo.userId = res.data.info.userId
                    this.A_UPDATE_USER(res.data.info)
                    this.loginForm.account = null
                    this.loginForm.password = null
                    this.$cookies.set("fiber-token",data.token,{
                        maxAge: 60 * 60 * 24 * 7,
                        path: '/'
                    })
                    // 刷新当前页面
                    location.reload()
                    // 连接ws
                    // this.$setWs.initWebSocket("ws://localhost:8199/ws/web")
                    // this.close()
                } catch (error) {
                    console.log(error)
                    setTimeout(() => {
                        this.$notification.error({
                            message: '网络错误',
                            description: "请稍后再试"
                        })
                    }, 1000)
                }
            });
        },
        // 切换注册
        goRegister(){
            this.type = "register"
            this.title = "用户注册"
            this.btnTitle = "立即注册"
        },
        // 切换登录
        goLogin(){
            this.type = "login"
            this.title = "登录"
            this.btnTitle = "快速登录"
        },
        cancel(){
            this.state.state = "cancel"
            this.close()
        },
        ascertain(){
            this.state.state = "ascertain"
            this.close()
        },
        open() {
            this.isTrue = true;
        },
        close() {
            this.isTrue = false;
        },
        
    }
};
</script>

<style lang="less" scoped>

.login_box {
    user-select: none;
    pointer-events: none;
    z-index: 20;
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    width: 100%;
    height: 100%;
    opacity: 0;
    visibility: hidden;
    transform: perspective(1px) scale(1.1);
    transition: visibility 0s linear .15s,opacity .15s 0s,transform .15s;
    display: flex;
    align-items: center;
    justify-content: center;
    .login_container{
        width: 340px;
        background-color: white;
        position: relative;  
        .login_close{
            position: absolute;
            right: -40px;
            font-size: 30px;
            color: #fff;
            z-index: 22;
            cursor: pointer;
        }
        // .login_left{
        //     .logo{
        //         width: 100%;
        //         height: 100%;
               
        //         img{
        //             width: 100%;
        //             height: 100%;
        //         }
        //     }
        // }
        .login_right{
            padding: 30px 24px 24px;
            background-image: url("/img/login.png");
            background-repeat: no-repeat;
            background-size: 100%;
            h2{
                text-align: center;
            }

            .login_input{
                margin: 40px 0 0 0;
                .login_register{
                    margin-bottom: 20px;
                    display: flex;
                    justify-content: space-between;
                    .forget{
                        color: #0084ff;
                    }
                    .go_register{
                        cursor: pointer;
                        color: #0084ff;
                    }
                }
            }

            .register_input{
                margin: 40px 0 0 0;
                .go_login{
                    margin-bottom: 20px;
                    color: #1890ff;
                    cursor: pointer;
                }
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

            .login_open{
                margin-top: 20px;
                padding: 10px;
                border: 1px dashed rgba(202, 202, 202, 0.7);
                display: flex;
                justify-content: space-between;
                .login_open_title{
                    font-size: 16px;
                }
                .login_open_ul{
                    display: flex;
                    .wechat{
                        background: rgba(0, 132, 255, 0.09);
                        padding: 0 6px;
                        margin-right: 10px;
                        justify-content: center;
                        align-items: center;
                        border-radius: 3px;
                        color: #17b31e;
                        font-size: 16px;
                        span{
                            margin-left: 5px;
                        }
                    }
                    .qq{
                        background: rgba(0, 132, 255, 0.09);
                        padding: 0 6px;
                        justify-content: center;
                        align-items: center;
                        border-radius: 3px;
                        color: #4cc0f2;
                        font-size: 16px;
                        span{
                            margin-left: 5px;
                        }
                    }
                }
            }

            .register_proxy{
                margin-top: 10px;
                padding: 10px 0;
                font-size: 12px;
                color: #999;
                border-top: 1px solid #f3f3f3;
                text-align: center;
                background: #f5f5f5;
                a{
                    color: #0084ff;
                }
            }
        }
    }
}
.is_back_show {
    opacity: 1 !important;
    background: rgba(42, 44, 48, 0.7);
    pointer-events: auto !important;
    visibility: visible;
    transform: perspective(1px) scale(1);
    transition: visibility 0s linear 0s,opacity .15s 0s,transform .15s;
}

@media only screen and (max-width: 768px) {
    .login_box{
        .login_container{
            margin: 0 20px;   
            .login_close{
                position: absolute;
                top: 0;
                right: 0;
                font-size: 20px;
                color: #000;
                padding: 10px;    
            }
        }
    }
}
</style>
