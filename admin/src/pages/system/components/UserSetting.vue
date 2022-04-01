<template>
<div class="setting">
    <div class="setting-item">
        <div class="setting-input">
            <span>用户默认头像</span>
            <a-input v-model="defaultAvatar"
             class="setting-content" />
        </div>
        <p>用户默认头像，尺寸建议1：1</p>
        <div v-action="'/media/upload'" class="setting-upload">
            <a-button @click="avatarUpload" type="primary">
                上传
            </a-button>
        </div>
    </div>
    <div class="setting-item">
        <div class="setting-input">
            <span>用户默认背景</span>
            <a-input v-model="defaultCover" class="setting-content" />
        </div>
        <p>用户默认背景，尺寸建议4：1</p>
        <div v-action="'/media/upload'" class="setting-upload">
            <a-button @click="coverUpload" type="primary">
                上传
            </a-button>
        </div>
    </div>

    <div class="setting-item">
        <div class="setting-input">
            <span>用户默认等级</span>
            <a-input v-model="defaultGrade" class="setting-content" />
        </div>
        <p>输入用户默认的等级例如: 1</p>
    </div>
    <div class="setting-item">
        <div class="setting-input">
            <span>认证图标</span>
            <a-input v-model="verifyIcon"
             class="setting-content" />
        </div>
        <p>认证后显示的图标建议尺寸1:1</p>
        <div v-action="'/media/upload'" class="setting-upload">
            <a-button @click="verifyIconUpload" type="primary">
                上传
            </a-button>
        </div>
    </div>
    <div class="setting-item">
        <div class="setting-input">
            <span>认证费用</span>
            <a-input-number class="setting-content"  v-model="verifyPrice" placeholder="请输入认证费用"   :min="0" :max="9999" :precision="2"/>
        </div>
        <p>0 则表示免费</p>
    </div>

    <div class="setting-item">
        <div class="setting-input">
            <span>认证模式</span>
            <a-select class="setting-content"  v-model="verifyMode"  placeholder="请设置认证模式" >
                <a-select-option value="0">
                    人工认证
                </a-select-option>
                <!-- <a-select-option value="1">
                    {{$t('system.setting.user.verifyMode.two.text')}}
                </a-select-option> -->
            </a-select>
        </div>
        <p>默认为人工认证，后期会接入第三方认证接口</p>
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
            verifyIcon: null,
            verifyPrice: 0.00,
            verifyMode: '0',
            defaultAvatar:null,
            defaultCover:null,
            defaultGrade:null,
        }
    },
    created(){
        this.getData()
    },
    methods:{

        async getData(){
            const res = await Info("UserSetting")  
            if (res.data.info != null) {
                const value = JSON.parse(res.data.info)
                this.verifyIcon = value.verifyIcon
                this.verifyPrice = value.verifyPrice
                this.verifyMode = value.verifyMode
                this.defaultAvatar = value.defaultAvatar
                this.defaultCover = value.defaultCover
                this.defaultGrade = value.defaultGrade
            }
        },
        async avatarUpload(){
            this.$Upload().then((res)=>{
                if (res != false) {
                    console.log(res)
                    this.defaultAvatar = res
                    // this.createForm.cover = res
                }
            }).catch((err)=>{
                 console.log(err)
                // this.createForm.cover = undefined
            })
        },
        async coverUpload(){
            this.$Upload().then((res)=>{
                if (res != false) {
                    this.defaultCover = res
                    // this.createForm.cover = res
                }
            }).catch((err)=>{
                 console.log(err)
                // this.createForm.cover = undefined
            })
        },
        async verifyIconUpload(){
            this.$Upload().then((res)=>{
                if (res != false) {

                    this.verifyIcon = res
                    // this.createForm.cover = res
                }
            }).catch((err)=>{
                 console.log(err)
                // this.createForm.cover = undefined
            })
        },
        handleSubmit(){
            this.$confirm({
                content: this.$t('confirm.content.text'),
                onOk:() => {
                    const options = {
                        defaultAvatar: this.defaultAvatar,
                        defaultCover: this.defaultCover,
                        defaultGrade:this.defaultGrade,
                        verifyMode: this.verifyMode,
                        verifyPrice: this.verifyPrice,
                        verifyIcon: this.verifyIcon,
                    }
                    const config = {
                        "configValue":JSON.stringify(options),
                        "configKey": "UserSetting",
                        "configName": "用户设置",
                        "remark" : "用户设置"
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
