<template>
    <a-card :bordered="false">
        <div class="setting">
            <div class="setting-item">
                <div class="setting-input">
                    <span>邮箱主机地址</span>
                    <a-input v-model="host" class="setting-content" />
                </div>
                <p>邮箱地址例如：smtp.qq.com</p>
            </div>
            <div class="setting-item">
                <div class="setting-input">
                    <span>邮箱主机端口</span>
                    <a-input v-model="port" class="setting-content" />
                </div>
                <p>例如：587</p>
            </div>
            <div class="setting-item">
                <div class="setting-input">
                    <span>邮箱发送者</span>
                    <a-input v-model="user" class="setting-content" />
                </div>
                <p>邮箱地址例如：xxxx</p>
            </div>
            <div class="setting-item">
                <div class="setting-input">
                    <span>邮箱密钥</span>
                    <a-input v-model="pass" class="setting-content"/>
                </div>
                <p>请注意你配置的邮箱密钥</p>
            </div>
            <div class="setting-item">
                <div class="setting-input">
                    <span>发送邮箱</span>
                    <a-input v-model="email" class="setting-content"/>
                </div>
                <p>邮箱地址例如：xxx@qq.com</p>
            </div>
            <div v-action="'/system/save'" class="setting-item">
                <a-button @click="handleSubmit" type="primary" class="setting-save">
                    保存
                </a-button>
            </div>
        </div>
    </a-card>
</template>
<script>
import { Save,Info } from '@/api/system'
export default {
    data(){
        return{
            host:null,
            port:null,
            user:null,
            pass:null,
            email:null,
        }
    },
    created(){
        this.getData()
    },
    methods:{
         async getData(){
            const res = await Info("EmailOptions")  
            if (res.data.info != null) {
                const value = JSON.parse(res.data.info)
                this.host=value.host
                this.port=value.port
                this.user=value.user
                this.pass= value.pass
                this.email= value.email
            }
        },
        handleSubmit(){
            this.$confirm({
                title: '是否修改',
                content: '是否要修改配置选项，请注意',
                onOk:() => {
                    const options = {
                        host: this.host,
                        port: this.port,
                        user:this.user,
                        pass: this.pass,
                        email: this.email,
                    }
                    const config = {
                        "configValue":JSON.stringify(options),
                        "configKey": "EmailOptions",
                        "configName": "邮箱配置",
                        "remark" : "邮箱配置"
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
