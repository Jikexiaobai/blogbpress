<template>
    <div class="container">
        <a-space class="container-action" size="large">
            <a-input-number 
                v-model="count" 
                size="large" 
                placeholder="请输入创建张数"
                :style="{ width: '100%' }"
                :precision="0"
                :min="1" 
                :max="100" />
            <a-input-number 
                v-model="money" 
                size="large" 
                placeholder="请输入卡密对应金额"
                :style="{ width: '100%' }"
                :precision="0"
                :min="1" 
                :max="999" />
            <a-button v-action="'/card/create'" @click="batchCreate" type="primary">
                批量创建
            </a-button>
            <a-button v-action="'/card/remove'" :disabled="selectedRowKeys.length > 0 ? false :true" @click="batchRemove" type="danger">
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
        <span slot="status" slot-scope="status">
            <a-tag v-if="status == 1" color="#f50">
                未使用
            </a-tag>
            <a-tag v-if="status == 2" color="#87d068">
               	已使用
            </a-tag>
        </span>
        <span slot="action" slot-scope="text, record">
            <a v-action="'/card/remove'" @click="remove(record.id)">删除</a>   
        </span>
        </a-table>
    </div>
</template>

<script>
import { Remove,Create } from '@/api/card'
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
            count:undefined,
            money:undefined,
            // 表格
            columns:[
                {
                    title: "编号",
                    dataIndex: 'id',
                    
                },
                {
                    title: "使用者",
                    dataIndex: 'nickName',
                    scopedSlots: { customRender: 'nickName' }
                },
                {
                    title: "密钥",
                    dataIndex: 'secretKey',
                    scopedSlots: { customRender: 'secretKey' }
                },
                {
                    title: "对应金额",
                    dataIndex: 'money',
                    scopedSlots: { customRender: 'money' }
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
        batchCreate(){
            this.$confirm({
                title: "请确认创建卡密信息是否正确",
                onOk:() => {
                    if (this.count == 0 || this.count == undefined) {
                        this.$message.error("请输入创建卡密张数");
                        return
                    }
                    if (this.money == 0 || this.money == undefined) {
                        this.$message.error("请输入卡密对应金额");
                        return
                    }

                    const formData = {
                        count:this.count,
                        money:this.money
                    }
                    this.postCreate(formData)
                    this.count = undefined
                    this.money = undefined
                    // this.selectedRowKeys = []
                },
                onCancel() {},
            });
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
        async postCreate(formData){
            try {
                const res = await Create(formData)
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
.remark{
    margin-top: 20px;
}
</style>
