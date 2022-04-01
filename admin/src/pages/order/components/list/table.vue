<template>
    <div class="container">
        <a-space class="container-action" size="large">
            <a-button v-action="'/order/remove'" :disabled="selectedRowKeys.length > 0 ? false :true" @click="batchRemove" type="danger">
                批量删除
            </a-button>
        </a-space>
        <a-table 
            :rowKey="record=>record.id"
            :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
            :pagination="{
                showQuickJumper:true,
                showSizeChanger:true,
                showTotal:(total)=>`共${total}条`,
                pageSize:pageSize,
                current:current,
                total:total
            }"
            :columns="columns" 
            :dataSource="list"
            @change="pageChange"
        >
        <span slot="paymentMoney" slot-scope="paymentMoney">
            {{paymentMoney}} ￥
        </span>
        <span slot="orderMoney" slot-scope="orderMoney">
            {{orderMoney}} ￥
        </span>
        <span slot="status" slot-scope="status">
            <a-tag v-if="status == 1" color="#f50">
               未支付
            </a-tag>
            <a-tag v-if="status == 2" color="#87d068">
                已支付
            </a-tag>
        </span>
        <span slot="payMethod" slot-scope="payMethod">
            <a-tag v-if="payMethod == 1" color="#2db7f5">
                支付宝
            </a-tag>
            <a-tag v-if="payMethod == 2" color="#87d068">
                微信
            </a-tag>
            <a-tag v-if="payMethod == 3" color="#f50">
                余额
            </a-tag>
        </span>

        <span slot="orderType" slot-scope="orderType">
            <a-tag v-if="orderType == 1" color="#2db7f5">
                充值
            </a-tag>
            <a-tag v-if="orderType == 2" color="#87d068">
                打赏充电
            </a-tag>
            <a-tag v-if="orderType == 3" color="#f50">
                内容购买
            </a-tag>
            <a-tag v-if="orderType == 4" color="#f50">
                查看话题隐藏内容
            </a-tag>
            <a-tag v-if="orderType == 5" color="#f50">
                加入付费圈子
            </a-tag>
            <a-tag v-if="orderType == 6" color="#f50">
                购买付费课程
            </a-tag>
            <a-tag v-if="orderType == 7" color="#f50">
                查看付费答案
            </a-tag>
            <a-tag v-if="orderType == 8" color="#f50">
                开通vip
            </a-tag>
            <a-tag v-if="orderType == 9" color="#f50">
                认证付费
            </a-tag>
        </span>

        <span slot="action" slot-scope="text, record">
            <a v-action="'/order/edit'" @click="edit(record.id)">编辑</a> 
            <a v-action="'/order/info'" @click="view(record.id)">查看</a>     
            <a-divider type="vertical" />
            <a v-action="'/order/remove'"  @click="remove(record.id)">删除</a>   
        </span>
        </a-table>

        <a-modal v-model="visible" 
            title="请输入，处理通知信息"
            @ok="handleOk">
            <a-textarea v-model="remark" :placeholder="action == 'review' ? $t('modal.cancel.remark.placeholder') : $t('modal.remove.remark.placeholder')" :rows="4" />
        </a-modal>
    </div>
</template>

<script>
import { Remove,Review } from '@/api/order'
export default {
    props:{
        list:{
            type:Array,
                // 当为数组类型设置默认值时必须使用数组返回
            required:true,
            default: []
        },
        total:{
            type:Number,
            default: 0
        },
        pageSize:{
            type:Number,
            default: 10
        },
        current:{
            type:Number,
            default: 1
        },
    },
    data() {
        return {
            visible:false,
            remark:undefined,
            reviewFormData: null,
            // 表格
            columns:[
                {
                    title: "编号",
                    dataIndex: 'orderNum',
                    
                },
                {
                    title: "下单人",
                    dataIndex: 'nickName',
                    scopedSlots: { customRender: 'nickName' }
                },
                {
                    title: "类型",
                    dataIndex: 'orderType',
                    scopedSlots: { customRender: 'orderType' }
                },
                {
                    title: "订单金额",
                    dataIndex: 'orderMoney',
                    scopedSlots: { customRender: 'orderMoney' }
                },
                {
                    title: "支付金额",
                    dataIndex: 'paymentMoney',
                    scopedSlots: { customRender: 'paymentMoney' }
                },
                {
                    title: "支付方式",
                    dataIndex: 'payMethod',
                    scopedSlots: { customRender: 'payMethod' }
                },
                {
                    title: "状态",
                    dataIndex: 'status',
                    scopedSlots: { customRender: 'status' }
                },
                {
                    title: "发起时间",
                    dataIndex: 'createTime',
                    scopedSlots: { customRender: 'createTime' }
                },
                {
                    title: "支付时间",
                    dataIndex: 'payTime',
                    scopedSlots: { customRender: 'payTime' }
                },
                {
                    title: "操作",
                    width: '150px',
                    dataIndex: 'action',
                    scopedSlots: { customRender: 'action' }
                }
            ],
            selectedRowKeys: [],
            action:"review"
        }
    },
    methods: {
        onSelectChange(selectedRowKeys) {
            this.selectedRowKeys = selectedRowKeys;
        },
        goCreate(){
            this.$router.push({path:'/order/create'})
        },
        pageChange(e){
            this.$emit("changePage",e)
        },
     
        handleOk(){
            if (this.remark == null || this.remark == undefined ||this.remark == "") {
                 this.$message.error("请填写，处理信息");
                return
            }
            const formData = this.reviewFormData
            formData.remark = this.remark
            this.postReview(formData)
            this.visible = false
        },
        batchRemove(){
            this.$confirm({
                title: "正在删除",
                onOk:() => {
                    const formData = {
                        idList:this.selectedRowKeys
                    }
                    this.postRemove(formData)
                    this.selectedRowKeys = []
                },
                onCancel() {},
            });
        },
        remove(e){
            this.$confirm({
                title: "正在删除",
                onOk:() => {
                    const formData = {
                        "idList":[e]
                    }
                    this.postRemove(formData)
                },
                onCancel() {},
            });
        },
        view(e){
            this.$router.push({path: '/order/info',query:{id:e}});
        },
        edit(e){
            // this.$router.push({path: '/order/edit',query:{id:e}});
        },
        async postRemove(formData){
            try {
                const res = await Remove(formData)
               if (res.code != 1) {
                   this.$message.error(res.message);
                   return
                }
                this.$emit("getData")
            } catch (error) {
                console.log(error)
                setTimeout(() => {
                    this.$notification.error({
                        message: '网络错误',
                        description: "请稍后再试"
                    })
                }, 1000)
            }
        },
        async postReview(formData){
            try {
                const res = await Review(formData)
                if (res.code != 1) {
                   this.$message.error(res.message);
                   return
                }
                this.$emit("getData")
            } catch (error) {
                console.log(error)
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
.container-action{
    margin-bottom: 10px;
}
.avatar{
    width: 50px;
    height: 50px;
}
</style>
