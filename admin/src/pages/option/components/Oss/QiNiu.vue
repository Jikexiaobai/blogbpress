<template>
<div class="setting">
    <div class="setting-item">
        <div class="setting-input">
            <span>上传地区</span>
            <a-input v-model="address" class="setting-content"/>
        </div>
        <p>请填写你所在的区域</p>
    </div>
    <div class="setting-item">
        <div class="setting-input">
            <span>oss域名</span>
            <a-input v-model="endpoint" class="setting-content" />
        </div>
        <p>可以是OSS的默认域名，必须携带http://开头。（支持https）</p>
    </div>
    <div class="setting-item">
        <div class="setting-input">
            <span>AccessKeyId</span>
            <a-input v-model="accessKeyId" class="setting-content" />
        </div>
 
    </div>
    <div class="setting-item">
        <div class="setting-input">
            <span>AccessKeySecret</span>
            <a-input v-model="accessKeySecret" class="setting-content" />
        </div>

    </div>
    <div class="setting-item">
        <div class="setting-input">
            <span>Bucket</span>
            <a-input v-model="bucketName" class="setting-content" />
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
            address:null,
            accessKeyId:null,
            accessKeySecret:null,
            bucketName:null,
            endpoint:null,
        }
    },
    created(){
        this.getData()
    },
    methods:{
        async getData(){
             const res = await Info("QiNiuOssOption")  
            if (res.data.info != null || res.data.info != "") {
                const value = JSON.parse(res.data.info)
                this.endpoint=value.endpoint
                this.accessKeyId=value.accessKeyId
                this.accessKeySecret=value.accessKeySecret
                this.bucketName= value.bucketName
                this.address= value.address
            }
        },
        handleSubmit(){
            this.$confirm({
                title: '是否修改',
                content: '是否要修改配置选项，请注意',
                onOk:() => {
                    const options = {
                        endpoint: this.endpoint,
                        accessKeyId: this.accessKeyId,
                        accessKeySecret:this.accessKeySecret,
                        bucketName: this.bucketName,
                        address: this.address,
                    }
                    const config = {
                        "configValue":JSON.stringify(options),
                        "configKey": "QiNiuOssOption",
                        "configName": "七牛oss配置",
                        "remark" : "七牛oss配置"
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
