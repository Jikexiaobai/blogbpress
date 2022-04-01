<template>
  <div class="main">
    <a-form-model 
      ref="loginForm" 
      :model="loginForm" 
      :rules="loginForm.rules" 
      @submit="handleSubmit"
      class="user-layout-login">

      <a-alert v-if="isLoginError" type="error" showIcon style="margin-bottom: 24px;" :message="$t('user.login.message-invalid-credentials')" />
        <a-form-model-item ref="account" prop="account">
          <a-input
            size="large"
            type="text"
            :placeholder="$t('user.login.username.placeholder')"
            v-model="loginForm.account"
          >
            <a-icon slot="prefix" type="user" :style="{ color: 'rgba(0,0,0,.25)' }"/>
          </a-input>
        </a-form-model-item>

        <a-form-model-item ref="passWord" prop="passWord">
          <a-input-password
            size="large"
            :placeholder="$t('user.login.password.placeholder')"
            v-model="loginForm.passWord"
          >
            <a-icon slot="prefix" type="lock" :style="{ color: 'rgba(0,0,0,.25)' }"/>
          </a-input-password>
        </a-form-model-item>

        <!-- 验证码 -->
        <a-row>
          <a-col :span="14">
            <a-form-model-item ref="captcha" prop="captcha">
              <a-input
                size="large"
                type="text"
                :placeholder="$t('user.login.captcha.placeholder')"
                v-model="loginForm.captcha"
              >
              </a-input>
            </a-form-model-item>
          </a-col>
          <a-col :span="9" :offset="1">
            <img @click="getImageCaptcha" :src="captchaImg" alt="验证码" class="captcha">
          </a-col>
        </a-row>

      <a-form-model-item>
        <a-checkbox v-decorator="['rememberMe', { valuePropName: 'checked' }]">{{ $t('user.login.remember-me') }}</a-checkbox>
        <router-link
          :to="{ name: 'recover', params: { user: 'aaa'} }"
          class="forge-password"
          style="float: right;"
        >{{ $t('user.login.forgot-password') }}</router-link>
      </a-form-model-item>

      <a-form-model-item style="margin-top:24px">
        <a-button
          size="large"
          type="primary"
          htmlType="submit"
          class="login-button"
          :loading="state.loginBtn"
          :disabled="state.loginBtn"
        >{{ $t('user.login.login') }}</a-button>
      </a-form-model-item>

      <div class="user-login-other">
        <span>{{ $t('user.login.sign-in-with') }}</span>
        <a>
          <a-icon class="item-icon" type="alipay-circle"></a-icon>
        </a>
        <a>
          <a-icon class="item-icon" type="taobao-circle"></a-icon>
        </a>
        <a>
          <a-icon class="item-icon" type="weibo-circle"></a-icon>
        </a>
        <!-- <router-link class="register" :to="{ name: 'register' }">{{ $t('user.login.signup') }}</router-link> -->
      </div>
    </a-form-model>
  </div>
</template>

<script>
import { mapActions } from 'vuex'
import { timeFix } from '@/utils/util'
import { login,getCaptcha } from '@/api/auth'
export default {
  components: {
   
  },
  data () {
    return {
      // customActiveKey: 'account',
      isLoginError: false,
      form: this.$form.createForm(this),
      loginForm:{
          account:"test@admin.com",
          passWord:"a121300",
          captcha:undefined,
          rules:{
              account:[
                { required: true, message: this.$t('user.userName.required'), trigger: 'change' },
                { validator: this.handlePhoneOrEmail ,trigger: 'change' },
              ],
              passWord:[
                  { required: true, message: this.$t('user.password.required'), trigger: 'blur' },
              ],
              captcha:[
                  { required: true, message: this.$t('user.captcha.required'), trigger: 'blur' },
              ],
          },
      },

      state: {
        time: 60,
        loginBtn: false,
        // login type: 0 email, 1 username, 2 telephone
        loginType: 0,
        smsSendBtn: false
      },
      // 验证码信息
      captchaImg:null,
      captchaKey:null
    }
  },
  created () {
    this.getImageCaptcha()

  },
  methods: {
    ...mapActions("user",['A_UPDATE_TOKEN']),
    async getImageCaptcha(){
      const res = await getCaptcha()
       if (res.code != 1) {
        this.$notification['error']({
          message: '错误',
          description: res.message,
          duration: 4
        })
      }
      this.captchaImg = res.data.imageCaptcha.B64
      this.captchaKey = res.data.imageCaptcha.Id
    },

    handlePhoneOrEmail (rule, value, callback) {
      const { state } = this
      const regex = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((\.[a-zA-Z0-9_-]{2,3}){1,2})$/
      if (regex.test(value)) {
        state.loginType = 0
      } else {
        state.loginType = 1
      }
      callback()
    },

    handleSubmit (e) {
      e.preventDefault()
      const {
        state,
      } = this
      
      state.loginBtn = true

      // const validateFieldsKey = customActiveKey === 'userName' ? ['userName', 'passWord','captcha'] : ['mobile', 'captcha']
    
      this.$refs.loginForm.validate(async (valid) => {
        if (valid) {
          const loginParams = { ...this.loginForm }
          
          // delete loginParams.userName
          // loginParams[!state.loginType ? 'email' : 'userName'] = this.loginForm.userName

          loginParams.key = this.captchaKey
          
          const res = await login(loginParams)
          if (res.code != 1) {
              this.$notification.error({
                message: '错误',
                description: res.message
              })
              setTimeout(() => {
                state.loginBtn = false
              }, 600)
              this.isLoginError = true
              return
          }
          this.A_UPDATE_TOKEN(res.data.token)
          this.$router.push({ path: '/' })
          // 延迟 1 秒显示欢迎信息
          setTimeout(() => {
            this.$notification.success({
              message: '欢迎',
              description: `${timeFix()}，欢迎回来`
            })
          }, 1000)
          this.isLoginError = false
        } else {
          setTimeout(() => {
            state.loginBtn = false
          }, 600)
        }
      })
    },
    getCaptcha (e) {
      e.preventDefault()
      const { form: { validateFields }, state } = this

      validateFields(['mobile'], { force: true }, (err, values) => {
        if (!err) {
          state.smsSendBtn = true

          const interval = window.setInterval(() => {
            if (state.time-- <= 0) {
              state.time = 60
              state.smsSendBtn = false
              window.clearInterval(interval)
            }
          }, 1000)

          const hide = this.$message.loading('验证码发送中..', 0)
          getSmsCaptcha({ mobile: values.mobile }).then(res => {
            setTimeout(hide, 2500)
            this.$notification['success']({
              message: '提示',
              description: '验证码获取成功，您的验证码为：' + res.result.captcha,
              duration: 8
            })
          }).catch(err => {
            setTimeout(hide, 1)
            clearInterval(interval)
            state.time = 60
            state.smsSendBtn = false
            this.requestFailed(err)
          })
        }
      })
    },
    stepCaptchaSuccess () {
      this.loginSuccess()
    },
    stepCaptchaCancel () {
      this.Logout().then(() => {
        this.loginBtn = false
        this.stepCaptchaVisible = false
      })
    },
  }
}
</script>

<style lang="less" scoped>
.user-layout-login {
  label {
    font-size: 14px;
  }

  .getCaptcha {
    display: block;
    width: 100%;
    height: 40px;
  }

  .forge-password {
    font-size: 14px;
  }

  button.login-button {
    padding: 0 15px;
    font-size: 16px;
    height: 40px;
    width: 100%;
  }

  .user-login-other {
    text-align: left;
    margin-top: 24px;
    line-height: 22px;

    .item-icon {
      font-size: 24px;
      color: rgba(0, 0, 0, 0.2);
      margin-left: 16px;
      vertical-align: middle;
      cursor: pointer;
      transition: color 0.3s;

      &:hover {
        color: #1890ff;
      }
    }

    .register {
      float: right;
    }
  }
}
</style>
