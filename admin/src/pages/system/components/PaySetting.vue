<template>
    <div class="setting">
        <div class="setting-item">
            <div class="setting-input">
                <span>支付宝</span>
                <a-select class="setting-content"  v-model="alyPay">
                    <a-select-option :value="1">
                        关闭
                    </a-select-option>
                    <a-select-option :value="2">
                        官方
                    </a-select-option>
                </a-select>
            </div>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>微信支付</span>
                <a-select class="setting-content"  v-model="weChatPay">
                     <a-select-option :value="1">
                        关闭
                    </a-select-option>
                    <a-select-option :value="2">
                        官方
                    </a-select-option>
                </a-select>
            </div>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>支付选项</span>
                <a-checkbox-group v-model="payMode" :options="payOptions"   />
            </div>
            <p>请设置支付选项内容</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>充值选项</span>
                <a-checkbox-group v-model="recharge" :options="options"   />
            </div>
            <p>请设置充值选项内容</p>
        </div>
      
        <div class="setting-item">
            <div class="setting-input">
                <span>平台分成比例</span>
                <a-input-number class="setting-content" 
                v-model="servicePercent" 
                 :min="0" :max="1" :precision="2"/>
            </div>
            <p>默认 0.05 （5%），如果网站不抽成，请设置为0。</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>最小提现金额</span>
                <a-input-number class="setting-content" 
                 v-model="cashMin" 
                   :min="0" :max="9999" :precision="2"/>
            </div>
            <p>设置最小提现金额例如: 50</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>提现服务费</span>
                <a-input-number class="setting-content"  
                v-model="cashServicePercent" 
                  :min="0" :max="1" :precision="2"/>
            </div>
            <p>默认 0.05 （5%），如果不收取服务费，请设置为0。</p>
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
            alyPay:"close",
            weChatPay:"close",
            payMode:[],
            recharge:[],
            cashMin:50,
            cashServicePercent:0.05,
            servicePercent:0.05,
            options: [
                { label: '支付宝', value: 1 },
                { label: '微信', value: 2 },
                { label: '卡密', value: 3 },
                { label: '人工转账', value: 4 },
            ],
            payOptions: [
                { label: '支付宝', value: 1 },
                { label: '微信', value: 2 },
                { label: '余额', value: 3 },
            ]
        }
    },
    created(){
        this.getData()
    },
    methods:{
        
        async getData(){
            const res = await Info("PaySetting")  
            if (res.data.info != null) {
                const value = JSON.parse(res.data.info)
                this.alyPay = value.alyPay
                this.weChatPay = value.weChatPay
                this.payMode = value.payMode
                this.recharge = value.recharge
                this.cashMin = value.cashMin
                this.cashServicePercent = value.cashServicePercent
                this.weChaservicePercenttPay = value.servicePercent
            }

        },
        handleSubmit(){
            this.$confirm({
                content: this.$t('confirm.content.text'),
                onOk:() => {
                    const options = {
                        alyPay: this.alyPay,
                        weChatPay: this.weChatPay,
                        payMode:  this.payMode,
                        recharge: this.recharge,
                        cashMin: this.cashMin,
                        cashServicePercent: this.cashServicePercent,
                        servicePercent: this.servicePercent,
                    }
                    const config = {
                        "configValue":JSON.stringify(options),
                        "configKey": "PaySetting",
                        "configName": "支付设置",
                        "remark" : "支付设置"
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
