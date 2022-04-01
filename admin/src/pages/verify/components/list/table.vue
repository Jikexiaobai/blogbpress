<template>
    <div class="container">
        <a-space class="container-action" size="large">
            <a-button v-action="'/verify/remove'" :disabled="selectedRowKeys.length > 0 ? false :true" @click="batchRemove" type="danger">
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

        <span slot="mode" slot-scope="mode">
            <a-tag v-if="mode == 1" color="#2db7f5">
                QQ
            </a-tag>
            <a-tag v-if="mode == 2" color="#87d068">
                微信
            </a-tag>
        </span>

        <span slot="status" slot-scope="status">
            <a-tag v-if="status == 1" color="#2db7f5">
                待审核
            </a-tag>
            <a-tag v-if="status == 2" color="#87d068">
                已认证
            </a-tag>
            <a-tag v-if="status == 3" color="#f50">
                未通过
            </a-tag>
        </span>
       
        <span slot="action" slot-scope="text, record">
            <a-popconfirm
                v-action="'/verify/review'"
                v-if="record.status == 1"
                title="审核认证内容"
                ok-text="通过"
                cancel-text="驳回"
                @confirm="review(2,record.id)"
                @cancel="review(3,record.id)"
            >
                <a href="#">审核</a>
            </a-popconfirm>
            <a-divider v-if="record.status == 1" type="vertical" />
            <a  v-action="'/verify/remove'" @click="remove(record.id)">删除</a>
        </span>
        </a-table>

        <a-modal v-model="visible" 
            title="请填写，处理信息"
            @ok="handleOk">
            <a-textarea v-model="remark" :rows="4" />
        </a-modal>
    </div>
</template>

<script>
import { Remove,Review } from '@/api/verify'
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
                    title: "真实姓名",
                    dataIndex: 'name',
                    
                },
                {
                    title: "身份证号码",
                    dataIndex: 'code',
                    scopedSlots: { customRender: 'code' }
                },
                {
                    title: "认证用户",
                    dataIndex: 'nickName',
                    scopedSlots: { customRender: 'nickName' }
                },
                {
                    title: "联系方式",
                    dataIndex: 'mode',
                    scopedSlots: { customRender: 'mode' }
                },
                {
                    title: "联系号码",
                    dataIndex: 'number',
                    scopedSlots: { customRender: 'number' }
                },
                {
                    title: "认证状态",
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
            selectedRowKeys: [],
            action:"review"
        }
    },
    methods: {
        onSelectChange(selectedRowKeys) {
            this.selectedRowKeys = selectedRowKeys;
        },
        pageChange(e){
            this.$emit("changePage",e)
        },

        review(e,i){
            const formData = {
                "status":e,
                "idList":[i]
            }
            this.action = undefined
            this.remark = undefined
            this.reviewFormData = undefined
            this.action = "review"
            this.visible = true
            this.reviewFormData = formData
        },
        handleOk(){
            const formData = this.reviewFormData
            formData.remark = this.remark
            if (this.action == "review") {
                this.postReview(formData)
            }
            if (this.action == "recover") {
                this.postRecover(formData)
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
