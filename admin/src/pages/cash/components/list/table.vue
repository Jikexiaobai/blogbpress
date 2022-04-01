<template>
    <div class="container">
        <a-space class="container-action" size="large">
            <a-button v-action="'/cash/remove'" :disabled="selectedRowKeys.length > 0 ? false :true" @click="batchRemove" type="danger">
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
        <span slot="cashMoney" slot-scope="cashMoney">
            {{cashMoney}} ￥
        </span>
        <span slot="serviceMoney" slot-scope="serviceMoney">
            {{serviceMoney}} ￥
        </span>
        <span slot="money" slot-scope="money">
            {{money}} ￥
        </span>
        <span slot="status" slot-scope="status">
            <a-tag v-if="status == 1" color="#f50">
                待审核
            </a-tag>
            <a-tag v-if="status == 2" color="#87d068">
               	提现成功
            </a-tag>
            <a-tag v-if="status == 3" color="#87d068">
               	提现失败
            </a-tag>
        </span>
         <span slot="mode" slot-scope="mode">
            <a-tag v-if="mode == 1" color="#f50">
                人工打款
            </a-tag>
            <a-tag v-if="mode == 2" color="#87d068">
                第三方到账
            </a-tag>
        </span>
        <span slot="payMethod" slot-scope="payMethod">
            <a-tag v-if="payMethod == 1" color="#2db7f5">
                支付宝
            </a-tag>
            <a-tag v-if="payMethod == 2" color="#87d068">
                微信
            </a-tag>
        </span>
        <div slot="expandedRowRender" slot-scope="record" style="margin: 0">
            <span>收据号：</span> {{record.receiptNum}}
            <div v-if="record.status == 3">
                <span>失败原因：{{record.remark}}</span>
            </div>
        </div>
        <span slot="action" slot-scope="text, record">
            <a-popconfirm
                title="正在审核"
                ok-text="通过"
                cancel-text="拒绝"
                @confirm="review(2,record.code)"
                @cancel="review(3,record.code)"
            >
                <a v-action="'/cash/review'" v-if="record.status == 1">审核</a>
            </a-popconfirm>
            <a-divider v-action="'/cash/review'" type="vertical" v-if="record.status == 1"/>
            <a v-action="'/cash/remove'" @click="remove(record.id)">删除</a>   
        </span>
        </a-table>

        <a-modal v-model="visible" 
            title="请填写下面信息"
            @ok="handleOk">
            <a-input v-if="reviewFormData.status == 2" v-model="receiptNum" size="large"
            placeholder="请输入收据号" />
            <a-textarea placeholder="请输入通知信息" class="remark" v-model="remark" :rows="4" />
        </a-modal>
    </div>
</template>

<script>
import { Remove,Review } from '@/api/cash'
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
            receiptNum:undefined,
            remark:undefined,
            reviewFormData:{
                status:null
            },
            action:"review",
            // 表格
            columns:[
                {
                    title: "编号",
                    dataIndex: 'code',
                    
                },
                {
                    title: "申请人",
                    dataIndex: 'nickName',
                    scopedSlots: { customRender: 'nickName' }
                },
                {
                    title: "到账金额",
                    dataIndex: 'cashMoney',
                    scopedSlots: { customRender: 'cashMoney' }
                },
                {
                    title: "服务费",
                    dataIndex: 'serviceMoney',
                    scopedSlots: { customRender: 'serviceMoney' }
                },
                {
                    title: "提现金额",
                    dataIndex: 'money',
                    scopedSlots: { customRender: 'money' }
                },
                {
                    title: "到账平台",
                    dataIndex: 'payMethod',
                    scopedSlots: { customRender: 'payMethod' }
                },
                {
                    title: "提现方式",
                    dataIndex: 'mode',
                    scopedSlots: { customRender: 'mode' }
                },
                {
                    title: "状态",
                    dataIndex: 'status',
                    scopedSlots: { customRender: 'status' }
                },
                {
                    title: "创建时间",
                    dataIndex: 'createTime',
                    scopedSlots: { customRender: 'createTime' }
                },

                {
                    title: "操作",
                    width: '150px',
                    dataIndex: 'action',
                    scopedSlots: { customRender: 'action' }
                }
            ],
            selectedRowKeys: []
        }
    },
    methods: {
        onSelectChange(selectedRowKeys) {
            this.selectedRowKeys = selectedRowKeys;
        },
        pageChange(e){
            this.$emit("changePage",e)
        },
        handleOk(){
            if (this.reviewFormData.status == 2) {
                if ((this.receiptNum == null || this.receiptNum == undefined || this.receiptNum == "") || (this.remark == null || this.remark == undefined ||this.remark == "") ) {
                    this.$message.error("请输入收据号和通知信息");
                    return
                }
            }
            
            const formData = this.reviewFormData
            formData.receiptNum = this.receiptNum
            formData.remark = this.remark

            this.postReview(formData)
            this.selectedRowKeys = []
            this.reviewFormData = {
                status:null
            }
            this.receiptNum = null
            this.remark = null
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
        review(e,i){
            const formData = {
                "status":e,
                "code":i
            }
            this.action = undefined
            this.remark = undefined
            this.reviewFormData = {
                status:null
            }
            this.action = "review"
            this.visible = true
            this.reviewFormData = formData
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
.remark{
    margin-top: 20px;
}
</style>
