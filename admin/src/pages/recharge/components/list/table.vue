<template>
    <div class="container">
        <a-space class="container-action" size="large">
            <a-button v-action="'/recharge/remove'" :disabled="selectedRowKeys.length > 0 ? false :true" @click="batchRemove" type="danger">
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
        <span slot="status" slot-scope="status, record">
            <a-tag v-if="status == 1 && (record.mode == 3 || record.mode ==4)" color="#f50">
                待审核
            </a-tag>
            <a-tag v-if="status == 1 && (record.mode == 1 || record.mode ==2)" color="#f50">
                待支付
            </a-tag>
            <a-tag v-if="status == 2" color="#2db7f5">
                充值成功
            </a-tag>
            <a-tag v-if="status == 3" color="#87d068">
                充值失败
            </a-tag>
        </span>
        <span slot="mode" slot-scope="mode">
            <a-tag v-if="mode == 1" color="#f50">
                支付宝
            </a-tag>
            <a-tag v-if="mode == 2" color="#2db7f5">
                微信
            </a-tag>
            <a-tag v-if="mode == 3" color="#87d068">
                卡密
            </a-tag>
            <a-tag v-if="mode == 4" color="#108ee9">
                人工充值
            </a-tag>
        </span>
        <div slot="expandedRowRender" slot-scope="record" style="margin: 0">
            <div v-if="record.mode == 4">
                    <a-space :size="50">
                        <span>转账人名称：{{record.name}}</span>
                        <span v-if="record.type == 1">转账类型：支付宝</span>
                        <span v-if="record.type == 2">转账类型：微信</span>
                        <span>转账订单号：{{record.number}}</span>
                    </a-space>
            </div>
            <div v-if="record.mode == 3">
                <span>卡密：{{record.cardKey}}</span>
            </div>
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
                <a v-action="'/recharge/review'" v-if="(record.mode == 3 || record.mode == 4) && record.status == 1">审核</a>
            </a-popconfirm>
            
            <a-divider v-action="'/recharge/review'" type="vertical" v-if="(record.mode == 3 || record.mode == 4) && record.status == 1"/>
            <a v-action="'/recharge/remove'" @click="remove(record.id)">删除</a>   
        </span>
        </a-table>

        <a-modal v-model="visible" 
            title="请输入审核信息"
            @ok="handleOk">
            <a-textarea placeholder="请输入审核信息" class="remark" v-model="remark" :rows="4" />
        </a-modal>
    </div>
</template>

<script>
import { Remove,Review } from '@/api/recharge'
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
            reviewFormData:undefined,
            action:"review",
            // 表格
            columns:[
                {
                    title: "编号",
                    dataIndex: 'code',
                    
                },
                {
                    title: "充值用户",
                    dataIndex: 'nickName',
                    scopedSlots: { customRender: 'nickName' }
                },
                {
                    title: "充值金额",
                    dataIndex: 'money',
                    scopedSlots: { customRender: 'money' }
                },
                {
                    title: "充值方式",
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
        goCreate(){
            this.$router.push({path:'/recharge/create'})
        },
        pageChange(e){
            this.$emit("changePage",e)
        },
     
        handleOk(){
            if (this.remark == null || this.remark == undefined ||this.remark == "") {
                if (this.action == "review") {
                    this.$message.error("请输入审核信息");
                }

                return
            }
            const formData = this.reviewFormData
            formData.remark = this.remark
            if (this.action == "review") {
                this.postReview(formData)
                this.selectedRowKeys = []
            }
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
            this.reviewFormData = undefined
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
