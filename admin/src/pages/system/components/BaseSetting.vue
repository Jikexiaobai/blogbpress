<template>
    <div class="setting">
        <div class="setting-item">
            <div class="setting-input">
                <span>站点名称</span>
                <a-input v-model="title" class="setting-content" />
            </div>
            <p>网站的名称例如: Fiber</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>副标题</span>
                <a-input v-model="childTitle" class="setting-content"/>
            </div>
            <p>例如: 副标题: - 超级无敌的 Fiber 1v5</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>站点描述</span>
                <a-textarea  v-model="description" :maxLength="200" class="setting-content" :rows="4" />
            </div>
            <p>用于描述你站点的主营业务，一般不超过200个字。</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>站点地址</span>
                <a-input v-model="url" class="setting-content"/>
            </div>
            <p>网站地址如：https://www.baidu.com</p>
        </div>
        <!-- logo -->
        <div class="setting-item">
            <div class="setting-input">
                <span>站点LOGO</span>
                <a-input v-model="logo" class="setting-content" />
            </div>
            <p>设置站点 LOGO 后将会以设置的站点 LOGO 为准。<br/>推荐尺寸大小为 240px*64px。</p>
            <div class="setting-upload">
                <a-button v-action="'/media/upload'" @click="uploadLogo" type="primary">
                    上传
                </a-button>
            </div>
        </div>

        <div class="setting-item">
            <div class="setting-input">
                <span>站点图标</span>
                <a-input v-model="icon" class="setting-content" />
            </div>
            <p >在站点一些交互功能使用。<br/>推荐尺寸大小为 1 : 1。</p>
            <div class="setting-upload">
                <a-button v-action="'/media/upload'" @click="uploadIcon" type="primary">
                    上传
                </a-button>
            </div>
        </div>

        <div class="setting-item">
            <div class="setting-input">
                <span>网站货币符号</span>
                <a-input v-model="currencySymbol" class="setting-content"/>
            </div>
            <p>网站货币符号</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>语言</span>
                <a-select class="setting-content" default-value="zh" v-model="language">
                    <a-select-option value="zh">
                        中文
                    </a-select-option>
                    <a-select-option value="en">
                        英文
                    </a-select-option>
                </a-select>
            </div>
            <p>设置语言</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>备案号</span>
                <a-input v-model="recordNumber" class="setting-content" />
            </div>
            <p>设置备案号</p>
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
            title:null,
            childTitle:null,
            description:null,
            url:null,
            logo:null,
            adminLogo:null,
            icon:null,
            currencySymbol:null,
            language:null,
            recordNumber:null,
        }
    },
    created(){
        this.getData()
    },
    methods:{
         async getData(){
            const res = await Info("BaseSetting")  
            if (res.data.info != null) {
                const value = JSON.parse(res.data.info)
                this.title = value.title
                this.childTitle = value.childTitle
                this.description = value.description
                this.logo = value.logo
                this.adminLogo = value.adminLogo
                this.icon = value.icon
                this.url = value.url
                this.currencySymbol = value.currencySymbol
                this.language = value.language
                this.recordNumber = value.recordNumber
            }
            
        },
        async uploadLogo(){
            this.$Upload().then((res)=>{
                if (res != false) {
                    console.log(res)
                    this.logo = res
                    // this.createForm.cover = res
                }
            }).catch((err)=>{
                 console.log(err)
                // this.createForm.cover = undefined
            })
        },
        async uploadAdminLogo(){
            this.$Upload().then((res)=>{
                if (res != false) {
                    this.adminLogo = res
                }
            }).catch((err)=>{
                 console.log(err)
                // this.createForm.cover = undefined
            })
        },
        async uploadIcon(){
            this.$Upload().then((res)=>{
                if (res != false) {
                    console.log(res)
                    this.icon = res
                    // this.createForm.cover = res
                }
            }).catch((err)=>{
                 console.log(err)
                // this.createForm.cover = undefined
            })
        },
        handleSubmit(){
            this.$confirm({
                content: "正在保存配置，请检查",
                onOk:() => {
                    const options = {
                        title: this.title,
                        childTitle: this.childTitle,
                        description:this.description,
                        url:this.url,
                        logo:this.logo,
                        adminLogo:this.adminLogo,
                        icon:this.icon,
                        currencySymbol:this.currencySymbol,
                        language:this.language,
                        recordNumber: this.recordNumber,
                    }
                    const config = {
                        "configValue":JSON.stringify(options),
                        "configKey": "BaseSetting",
                        "configName": "基础设置",
                        "remark" : "基础设置"
                    }
                    this.PostCreate(config)
                },
                onCancel() {},
            });
        },
        async PostCreate(value){
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
            // align-items: center;
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
