<template>
    <div class="container">
        <a-card :bordered="false">
            <a-descriptions :title="info.orderNum" bordered>
                <a-descriptions-item label="下单人">
                    {{info.nickName}}
                </a-descriptions-item>
                <a-descriptions-item label="收益人">
                    {{info.authorName}}
                </a-descriptions-item>
                <a-descriptions-item label="	支付方式">
                    <a-tag v-if="info.payMethod == 1" color="#2db7f5">
                        支付宝
                    </a-tag>
                    <a-tag v-if="info.payMethod == 2" color="#87d068">
                        微信
                    </a-tag>
                    <a-tag v-if="info.payMethod == 3" color="#f50">
                        余额
                    </a-tag>
                </a-descriptions-item>
                <a-descriptions-item label="订单金额">
                    {{info.orderMoney}} {{base.currencySymbol}}
                </a-descriptions-item>
                <a-descriptions-item label="优惠金额">
                    {{info.districtMoney}} {{base.currencySymbol}}
                </a-descriptions-item>
                <a-descriptions-item label="收益金额">
                    {{info.authorMoney}} {{base.currencySymbol}}
                </a-descriptions-item>
                <a-descriptions-item label="服务费用">
                    {{info.serviceMoney}} {{base.currencySymbol}}
                </a-descriptions-item>
                <a-descriptions-item label="支付金额">
                    {{info.paymentMoney}} {{base.currencySymbol}}
                </a-descriptions-item>
                <a-descriptions-item label="支付积分">
                    {{info.orderPoint}}
                </a-descriptions-item>
                <a-descriptions-item label="订单类型">
                    <span  v-if="info.orderType ==1">
                        充值
                    </span>
                    <span  v-if="info.orderType ==2">
                        充电打赏
                    </span>
                    <span  v-if="info.orderType ==3">
                        内容购买
                    </span>
                    <span  v-if="info.orderType ==4">
                        查看话题隐藏内容
                    </span>
                    <span  v-if="info.orderType ==5">
                        加入付费圈子
                    </span>
                    <span  v-if="info.orderType ==6">
                        报名付费课程
                    </span>
                    <span  v-if="info.orderType ==7">
                        查看付费答案
                    </span>
                    <span  v-if="info.orderType ==8">
                        开通会员
                    </span>
                    <span  v-if="info.orderType ==9">
                        支付认证费用
                    </span>
                </a-descriptions-item>
                <a-descriptions-item label="订单模式">
                    <span  v-if="info.orderMode ==1">
                        虚拟商品
                    </span>
                    <span  v-if="info.orderMode ==2">
                        实体商品
                    </span>
                </a-descriptions-item>
                <a-descriptions-item label="运输费用">
                    {{info.shippingMoney}} {{base.currencySymbol}}
                </a-descriptions-item>
                <a-descriptions-item label="收货地址">
                    {{info.shippingAddress}}
                </a-descriptions-item>
                <a-descriptions-item label="快递公司">
                    {{info.shippingCompName}}
                </a-descriptions-item>
                <a-descriptions-item label="联系方式">
                    {{info.shippingPhone}}
                </a-descriptions-item>
                <a-descriptions-item label="收货人姓名">
                    {{info.shippingName}}
                </a-descriptions-item>
                <a-descriptions-item label="快递单号">
                    {{info.shippingSn}}
                </a-descriptions-item>
                <a-descriptions-item label="发货时间">
                    {{info.shippingTime}}
                </a-descriptions-item>
                <a-descriptions-item label="商品名称">
                    {{info.title}}
                </a-descriptions-item>
                <a-descriptions-item label="所属模块">
                    <span  v-if="info.detailModule == 'audio'">
                        音频
                    </span>
                    <span  v-if="info.detailModule == 'video'">
                        视频
                    </span>
                    <span  v-if="info.detailModule == 'resource'">
                        资源
                    </span>
                    <span  v-if="info.detailModule == 'edu'">
                        课程
                    </span>
                    <span  v-if="info.detailModule == 'answer'">
                        答案
                    </span>
                    <span  v-if="info.detailModule == 'group'">
                        圈子
                    </span>
                </a-descriptions-item>
                <a-descriptions-item label="订单状态">
                    <a-tag v-if="info.status == 1" color="#2db7f5">
                        未支付
                    </a-tag>
                    <a-tag v-if="info.status == 2" color="#87d068">
                        已支付
                    </a-tag>
                </a-descriptions-item>
                <a-descriptions-item label="收货时间">
                    {{info.receiveTime}}
                </a-descriptions-item>
                <a-descriptions-item label="支付时间">
                    {{info.payTime}}
                </a-descriptions-item>
                <a-descriptions-item label="创建时间">
                    {{info.createTime}}
                </a-descriptions-item>
            </a-descriptions>
        </a-card>
    </div>
</template>

<script>
import { Info } from '@/api/order'
import { mapState } from 'vuex'
export default {
    data() {
        return {
            info:{
                orderNum:null
            }
        }
    },
    computed:{
        ...mapState("system",["base"]),
    },
    mounted(){
        if (isNaN(this.$route.query.id)) {
            this.$router.push({path:"/404"})
        }
        this.id = parseInt(this.$route.query.id)
        
        this.getData()
    },
    methods: {
        async getData(){
            try {
    
                const res = await Info({id:this.id})
                if (res.code != 1) {
                    this.$router.push({path:'/404'})
                }
                this.info = res.data.info
            } catch (error) {
    
                setTimeout(() => {
                    this.$notification.error({
                        message: '网络错误',
                        description: "请稍后再试"
                    })
                }, 1000)
            }
        },
    },
}
</script>

<style scoped lang="less">

</style>
