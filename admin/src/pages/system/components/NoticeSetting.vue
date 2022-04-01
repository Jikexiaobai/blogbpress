<template>
    <div class="setting">
        <div class="setting-item">
            <div class="setting-input">
                <span>{{$t('system.setting.notice.register.lable')}}</span>
                <a-input v-model="register" class="setting-content"/>
            </div>
            <p>{{$t('system.setting.notice.var.siteTitle.description')}}</p>
            <p>{{$t('system.setting.notice.register.description')}}</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>{{$t('system.setting.notice.content.lable')}}</span>
                <a-input v-model="create" class="setting-content" />
            </div>
            <p>{{$t('system.setting.notice.var.reason.description')}}</p>
            <p>{{$t('system.setting.notice.var.title.description')}}</p>
            <p>{{$t('system.setting.notice.content.description')}}</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>{{$t('system.setting.notice.remove.lable')}}</span>
                <a-input v-model="remove" class="setting-content" />
            </div>
            <p>{{$t('system.setting.notice.var.title.description')}}</p>
            <p>{{$t('system.setting.notice.var.reason.description')}}</p>
            <p>{{$t('system.setting.notice.remove.description')}}</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>{{$t('system.setting.notice.groupCreate.lable')}}</span>
                <a-input v-model="groupCreate" class="setting-content" />
            </div>
            <p>{{$t('system.setting.notice.var.reason.description')}}</p>
            <p>{{$t('system.setting.notice.var.title.description')}}</p>
            <p>{{$t('system.setting.notice.groupCreate.description')}}</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>{{$t('system.setting.notice.groupRemove.lable')}}</span>
                <a-input v-model="groupRemove" class="setting-content" />
            </div>
            <p>{{$t('system.setting.notice.var.reason.description')}}</p>
            <p>{{$t('system.setting.notice.var.title.description')}}</p>
            <p>{{$t('system.setting.notice.groupRemove.description')}}</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>{{$t('system.setting.notice.report.lable')}}</span>
                <a-input v-model="report" class="setting-content" />
            </div>
            <p>{{$t('system.setting.notice.report.description')}}</p>
            <p>{{$t('system.setting.notice.var.reason.description')}}</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>{{$t('system.setting.notice.userProhibit.lable')}}</span>
                <a-input v-model="userProhibit" class="setting-content" />
            </div>
            <p>{{$t('system.setting.notice.userProhibit.description')}}</p>
            <p>{{$t('system.setting.notice.var.reason.description')}}</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>{{$t('system.setting.notice.verify.lable')}}</span>
                <a-input v-model="verify" class="setting-content" />
            </div>
            <p>{{$t('system.setting.notice.verify.description')}}</p>
            <p>{{$t('system.setting.notice.var.reason.description')}}</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>{{$t('system.setting.notice.cash.lable')}}</span>
                <a-input v-model="cash" class="setting-content" />
            </div>
            <p>{{$t('system.setting.notice.var.code.description')}}</p>
            <p>{{$t('system.setting.notice.cash.description')}}</p>
        </div>
        <div class="setting-item">
            <a-button @click="handleSubmit" type="primary" class="setting-save">
                {{$t('save')}}
            </a-button>
        </div>
    </div>
</template>
<script>
import { save,getInfo } from '@/api/system'
export default {
    data(){
        return{
            register:undefined,
            create:undefined,
            remove:undefined,
            groupCreate:undefined,
            groupRemove:undefined,
            report:undefined,
            userProhibit:undefined,
            verify:undefined,
            cash:undefined,
        }
    },
    created(){
        this.getData()
    },
    methods:{
        
        async getData(){
            const info = await getInfo("NoticeSetting")  
            if (info.data.config != null) {
                const value = JSON.parse(info.data.config)
                this.register = value.register
                this.create = value.create
                this.remove = value.remove
                this.groupCreate = value.groupCreate
                this.groupRemove = value.groupRemove
                this.report = value.report
                this.userProhibit = value.userProhibit
                this.verify = value.verify
                this.cash = value.cash
            }

        },
        handleSubmit(){
            this.$confirm({
                content: this.$t('confirm.content.text'),
                onOk:() => {
                    const options = {
                        register:this.register,
                        create:this.create,
                        remove:this.remove,
                        groupCreate:this.groupCreate,
                        groupRemove:this.groupRemove,
                        report:this.report,
                        userProhibit: this.userProhibit,
                        verify: this.verify,
                        cash: this.cash,
                        cashRemove:this.cashRemove,
                    }
                    const config = {
                        "config_value":JSON.stringify(options),
                        "config_key": "NoticeSetting",
                        "config_name": "通知设置",
                        "remark" : "通知设置"
                    }
                    this.CreatePost(config)
                },
                onCancel() {},
            });
        },
        async CreatePost(value){
            try {
                const data = await save(value)
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
