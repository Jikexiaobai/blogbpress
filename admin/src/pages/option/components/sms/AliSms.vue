<template>
    <div class="setting">
        <div class="setting-item">
            <div class="setting-input">
                <span>AccessKey ID</span>
                <a-input v-model="id" class="setting-content"/>
            </div>
            <p>阿里云账户AccessKey管理下获取ID</p>
        </div>
        <div class="setting-item">
            <div class="setting-input ">
                <span>AccessKey Secret</span>
                <a-input v-model="secret" class="setting-content" />
                <!-- <a-textarea class="setting-content"  :rows="4" /> -->
            </div>
            <p>阿里云账户AccessKey管理下获取Secret</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>应用公钥证书</span>
                <a-input v-model="appPublicKey" class="setting-content" />
            </div>
            <div v-action="'/media/upload'" class="setting-upload">
                <a-button @click="handleUploadAppPublicKey" type="primary">
                    上传
                </a-button>
            </div>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>支付宝公钥证书</span>
                <a-input v-model="alyCertPublicKey" class="setting-content" placeholder="请上传" />
            </div>
            <div v-action="'/media/upload'" class="setting-upload">
                <a-button @click="handleUploadAlyPayPublicKey" type="primary">
                   上传
                </a-button>
            </div>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>应用根证书</span>
                <a-input v-model="alyRootCert" class="setting-content" placeholder="请上传" />
            </div>
            <div v-action="'/media/upload'" class="setting-upload">
                <a-button @click="handleUploadAliPayRootCert" type="primary">
                    上传
                </a-button>
            </div>
        </div>
        <div v-action="'/system/save'" class="setting-item">
            <a-button @click="handleSubmit" type="primary" class="setting-save">
                保存
            </a-button>
        </div>
    </div>
</template>
<script>
import { Save,Info } from '@/api/system'
export default {
    data(){
        return{
            id:null,
            secret:null,
            publicKey:null,
            appPublicKey:null,
            alyCertPublicKey:null,
            alyRootCert:null,
        }
    },
    created(){
        this.getData()
    },
    methods:{
        
        async getData(){
            const res = await Info("AlySmsOptions")  
            const value = JSON.parse(res.data.info)
            this.id = value.id
            this.secret = value.secret
            this.publicKey = value.publicKey
            
            this.appPublicKey = value.appPublicKey
            this.alyCertPublicKey = value.alyCertPublicKey
            this.alyRootCert = value.alyRootCert
        },
        async handleUploadAppPublicKey(){
            this.$Upload("Other").then((res)=>{
                this.appPublicKey = res
            }).catch(err =>{
                console.log(err)
            }) 
        },
        async handleUploadAlyPayPublicKey(){
            this.$Upload("Other").then((res)=>{
                this.alyCertPublicKey = res
            }).catch(err =>{
                console.log(err)
            }) 
        },
        async handleUploadAliPayRootCert(){
            this.$Upload("Other").then((res)=>{
                this.alyRootCert = res
            }).catch(err =>{
                console.log(err)
            }) 
        },
        handleSubmit(){
            this.$confirm({
                title: '是否修改',
                content: '是否要修改配置选项，请注意',
                onOk:() => {
                    const options = {
                        id: this.id,
                        secret : this.secret,
                        publicKey: this.publicKey,
                        appPublicKey: this.appPublicKey,
                        alyCertPublicKey:this.alyCertPublicKey,
                        alyRootCert:this.alyRootCert,
                    }
                    const config = {
                        "configValue":JSON.stringify(options),
                        "configKey": "AlySmsOptions",
                        "configName": "阿里云短信配置",
                        "remark" : "阿里云短信配置"
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
            margin-top: 10px;
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
