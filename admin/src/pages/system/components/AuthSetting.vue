<template>
    <div class="setting">

        <div class="setting-item">
            <div class="setting-input">
                <span>注册方式</span>
                    <a-radio-group v-model="registerMode" class="setting-input">
                    <a-radio value="email">
                        邮箱
                    </a-radio>
                    <a-radio value="phone" name="phone">
                        手机
                    </a-radio>
                </a-radio-group>
            </div>
            <p>请配置好对应的配置信息</p>
        </div>

        <div class="setting-item">
            <div class="setting-input">
                <span>社交注册登录</span>
                <a-checkbox-group v-model="social" class="setting-input">
                    <a-checkbox value="qq">
                        QQ
                    </a-checkbox>
                    <a-checkbox value="wechat">
                        微信
                    </a-checkbox>
                </a-checkbox-group>
            </div>
            <p>请配置好对应的配置信息，（目前还没实现）</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>隐私网址</span>
                <a-input v-model="policyUrl" class="setting-content" />
            </div>
            <p>设置后将会在注册框显示如：http://www.xxxxx.com</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>用户协议</span>
                <a-input v-model="protocolUrl" class="setting-content"/>
            </div>
            <p>设置后将会在注册框显示如：http://www.xxxxx.com</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>注册通知</span>
                <a-input v-model="register" class="setting-content"/>
            </div>
            <p>用户注册后的欢迎提示语例如：恭喜您注册成为我们的用户</p>
        </div>
        <div class="setting-item">
            <a-button v-action="'/system/save'" @click="handleSubmit" type="primary" class="setting-save">
                保存
            </a-button>
        </div>


    </div>
</template>
<script>
import { Save,Info } from '@/api/system'
// import { getByType } from '@/api/dict-data'
export default {
    data(){
        return{
            registerMode: undefined,
            policyUrl:undefined,
            protocolUrl: undefined,
            social: [],
            register:undefined,
        }
    },
    created(){
        this.getData()
    },
    methods:{

        async getData(){
            const res = await Info("AuthSetting")  
            if (res.data.info != null) {
                const value = JSON.parse(res.data.info)
                this.registerMode=value.registerMode
                this.policyUrl=value.policyUrl
                this.protocolUrl=value.protocolUrl
                this.social= value.social
                this.register = value.register
            }
        },
        handleSubmit(){
            this.$confirm({
                content: this.$t('confirm.content.text'),
                onOk:() => {
                    const options = {
                        registerMode: this.registerMode,
                        policyUrl: this.policyUrl,
                        protocolUrl: this.protocolUrl,
                        social: this.social,
                        register: this.register,
                    }
                    // options.authMode = options.authMode.join(",")
                    // options.social = options.social.join(",")
                    const config = {
                        "configValue":JSON.stringify(options),
                        "configKey": "AuthSetting",
                        "configName": "登录注册设置",
                        "remark" : "登录注册设置"
                    }
                    this.CreatePost(config)
                },
                onCancel() {},
            });
        },
        async CreatePost(value){
            try {
                const data = await Save(value)
                if (data.code == 1) {
                    this.$message.success(
                        data.message,
                        3
                    )
                } else {
                    this.$message.error(
                        data.message,
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
        },
    }
}
</script>

<style lang="less" scoped>
.setting{
    .setting-item{
        width: 520px;
        margin-top: 20px;
        .setting-input{
            display: flex;
            justify-content: space-between;
            align-items: center;
            span{
                width: 200px;
                font-size: 14px;
                font-weight: 700;
            }
            .setting-content{
                margin-left: 20px;
                width: 400px;
            }
        }
        .setting-upload{
            display: flex;
            justify-content: flex-end;
        }
        p{
            text-align: right;
            margin-top: 10px;
            font-size: 12px;
        }
        .setting-save{
            margin-right: 20px;
        }
    }
    
}
</style>
