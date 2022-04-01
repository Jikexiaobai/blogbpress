<template>
<div class="setting">
    <div class="setting-item">
        <div class="setting-input">
            <span>APPID</span>
            <a-input v-model="APPID" class="setting-content" placeholder="如：huadong" />
        </div>
        <p>(打开链接： https://open.alipay.com <br/>账户中心->密钥管理->开放平台密钥，<br/>填写您支付的应用的APPID)</p>
    </div>
    <div class="setting-item">
        <div class="setting-input">
            <span>商户id</span>
            <a-input v-model="siteTitle" class="setting-content" placeholder="请输入" />
        </div>
    </div>
    <div class="setting-item">
        <div class="setting-input">
            <span class="public-key">商户支付密钥</span>
            <a-textarea v-model="welcomeMsg" :rows="10" class="setting-content" placeholder="请输入" />
        </div>
    </div>
    <div class="setting-item">
        <a-button @click="handleSubmit" type="primary" class="setting-save">
        保存
        </a-button>
    </div>
</div>
</template>
<script>
import { mapActions } from 'vuex'
export default {
    data(){
        return{
            APPID:null,
            childTitle:null,
            siteLogo:null,
            recordNumber:null,
        }
    },
    created(){
        this.getData()
    },
    methods:{
        ...mapActions("setting",["settingSave","settingInfo"]),
         async getData(){
            const info = await this.settingInfo("BaseOptions")  
            const value = JSON.parse(info.data.config.config_value)
            this.siteTitle = value.siteTitle
            this.childTitle = value.childTitle
            this.siteLogo = value.siteLogo
            this.recordNumber = value.recordNumber
        },
        handleSubmit(){
            this.$confirm({
                title: '是否修改',
                content: '是否要修改配置选项，请注意',
                onOk:() => {
                    const options = {
                        siteTitle: this.siteTitle,
                        childTitle: this.childTitle,
                        siteLogo:this.siteLogo,
                        recordNumber: this.recordNumber,
                    }
                    const config = {
                        "config_value":JSON.stringify(options),
                        "config_key": "BaseOptions",
                        "config_name": "基础设置",
                        "remark" : "基础设置"
                    }
                    this.CreatePost(config)
                },
                onCancel() {},
            });
        },
        async CreatePost(value){
            try {
                const data = await this.settingSave(value)
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
            .public-key{
                display: flex;
                align-self: flex-start;
            }
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
