<template>
    <div class="recaptcha_box center opacity"
        :class="[isTrue && 'is_back_show']">
        <div class="recaptcha_container">
            <a-form-model ref="codeForm" :model="codeForm" :rules="codeForm.rules" class="codeForm">
                <div class="recaptcha">
                    <img @click="restCaptcha" :src="captcha" alt="">
                    <h2>请输入验证码</h2>
                    <p>请输入图片中的验证码 <br/>点击发送按钮获取验证码</p>
                    <a-form-model-item ref="code" prop="code">
                        <a-input size="large" v-model="codeForm.code" placeholder="输入验证码" />
                    </a-form-model-item>
                </div>
                <div class="recaptcha-button">
                    <a-button @click="cancel" type="dashed">
                        取消
                    </a-button>
                    <a-button @click="sureSend" type="primary">
                        发送
                    </a-button>
                </div>
            </a-form-model>
        </div>
    </div>
</template>
<style lang="less" scoped>
    .recaptcha_box {
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
        .recaptcha_container{

            background-color: white;
            position: relative;  
            padding: 30px 24px 24px;
            background-image: url("/img/login.png");
            background-repeat: no-repeat;
            background-size: 100%;
            .codeForm{
                .recaptcha{
                    display: flex;
                    flex-direction: column;
                    justify-content: center;
                    align-items: center;
                    text-align: center;
                    img{
                        border-radius: 3px;
                        width: 186px;
                        height: 50px;
                        cursor: pointer;
                        max-width: 100%;
                        height: auto;
                        object-fit: cover;
                        vertical-align: bottom;
                    }
                    h2{
                        text-align: center;
                        margin-bottom: 10px;
                        margin-top: 16px;
                        font-size: 18px;
                        font-weight: 700;
                    }
                    p{
                        color: #b2bac2;
                        font-size: 13px;
                        line-height: 1.5;
                        margin: 10px 0 20px;
                    }
                }
            }
            .recaptcha-button{
                display: flex;
                justify-content: space-between;
            }
        }
    }
    .is_back_show {
        opacity: 1 !important;
        background: rgba(42, 44, 48, 0.7);
        pointer-events: auto !important;
        opacity: 1;
        visibility: visible;
        transform: perspective(1px) scale(1);
        transition: visibility 0s linear 0s,opacity .15s 0s,transform .15s;
    }
    @media only screen and (max-width: 768px) {
        .recaptcha_box{
            .recaptcha_container{
                margin: 0 20px;
            }
        }
    }
</style>

<script>
import api from "@/api/index"
export default {
    data(){
        return{
            isTrue: false,
            captcha: "",
            captchaKey:null,
            state: null, // 准备（prepare） 确定（ ascertain） 取消（cancel）
            codeForm:{
                code:null,
                rules:{
                    code:[
                        { required: true, message: '请输入验证码', trigger: 'blur' }
                    ],
                },
            },
            account:null
        }
    },
    methods:{
        async confirm(
            account = null,
        ) {
            this.account = account || this.account;
            // this.title = title || this.title;
            const {data} = await this.$axios.get(api.getImageCaptcha)     
  
            this.captcha = data.imageCaptcha.B64
            this.captchaKey = data.imageCaptcha.Id
            this.open();
            return new Promise((resolve, reject) => {
                const target = { state: 'prepare' };
                let res = new Proxy(target, {
                    set(event, key, value) {
                        if (value === 'ascertain') {
                            resolve(true);
                        } else {
                            resolve(false);
                        }
                        return true;
                    }
                });
                this.state = res;
            });
        },
        async restCaptcha(){
            try {
                const {data} = await this.$axios.get(api.getImageCaptcha)
                this.captcha = data.imageCaptcha.B64
                this.captchaKey = data.imageCaptcha.Id
            } catch (error) {
                setTimeout(() => {
                    this.$notification.error({
                        message: '网络错误',
                        description: "请稍后再试"
                    })
                }, 1000)
            }
        },
        async sureSend() {
            this.$refs.codeForm.validateField("code",async (err)=>{
                if (err != "") {
                    console.log(err)
                    return false
                }
                // {"email":this.registerForm.email}
                const formData = {
                    account:this.account,
                    key:this.captchaKey,
                    captcha:this.codeForm.code
                }
                const res = await this.$axios.post(api.sendCaptcha,formData)
                if (res.code != 1) {
                    this.$message.error(
                        res.message,
                        3
                    )
                    return
                }
                this.ascertain()
                // this.state.state = 'ascertain';
                // this.close();
                // this.$Auth("register","用户注册","立即注册")
            })
        },
        cancel(){
            this.state.state = 'cancel'
            this.close()
            // this.$Auth("register","用户注册","立即注册")
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
        }
    }
}
</script>