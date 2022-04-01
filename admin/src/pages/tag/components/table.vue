<template>
    <div class="container">
        <a-space class="container-action" size="large">
            <a-button v-action="'/tag/remove'" :disabled="selectedRowKeys.length > 0 ? false :true" @click="batchRemove" type="danger">
                批量删除
            </a-button>
            <a-popconfirm
                v-action="'/tag/top'"
                title="是否设置推荐标签"
                ok-text="是"
                cancel-text="否"
                @confirm="batchTop(2)"
                @cancel="batchTop(1)"
            >
                <a-button :disabled="selectedRowKeys.length > 0 ? false :true">
                    批量推荐
                </a-button>
            </a-popconfirm>
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
        <span slot="top" slot-scope="top">
            <a-tag v-if="top == 2" color="#f50">
                推荐
            </a-tag>
            <a-tag v-if="top == 1" color="#2db7f5">
                不推荐
            </a-tag>
        </span>
        <span slot="action" slot-scope="text, record">
            <a-popconfirm
                v-action="'/tag/top'"
                title="是否设置推荐标签"
                ok-text="是"
                cancel-text="否"
                @confirm="top(2,record.tagId)"
                @cancel="top(1,record.tagId)"
            >
                <a href="#">{{$t('top')}}</a>
            </a-popconfirm>
           
            <a-divider type="vertical" />
            <a v-action="'/tag/remove'" @click="remove(record.tagId)">删除</a>
        </span>
        </a-table>
    </div>
</template>

<script>
import { Remove,Top } from '@/api/tag'
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
            // 表格
            columns:[
                {
                    title: "标签编号",
                    dataIndex: 'id',
                },
                {
                    title: "标题",
                    dataIndex: 'title',
                    scopedSlots: { customRender: 'title' }
                },
                {
                    title: "是否推荐",
                    dataIndex: 'top',
                    scopedSlots: { customRender: 'top' }
                },
                {
                    title: "创建时间",
                    dataIndex: 'createTime',
                },
                {
                    title: "操作",
                    width: '150px',
                    dataIndex: 'action',
                    scopedSlots: { customRender: 'action' }
                }
            ],
            selectedRowKeys: [],
        }
    },
    methods: {
        onSelectChange(selectedRowKeys) {
            this.selectedRowKeys = selectedRowKeys;
        },
        pageChange(e){
            this.$emit("changePage",e)
        },
        top(e,id){
            const formData = {
                idList:[id],
                isTop: e
            }
            this.postTop(formData)
        },
        batchTop(e){
            this.postTop({"idList":this.selectedRowKeys,isTop:e})
            this.selectedRowKeys = []
        },
        batchRemove(){
            this.$confirm({
                title: "正在删除",
                onOk:() => {
                    this.postRemove({"idList":this.selectedRowKeys})
                    this.selectedRowKeys = []
                },
                onCancel() {},
            });
        },
        remove(e){
            this.$confirm({
                title: "正在删除",
                onOk:() => {
                    this.postRemove({"idList":[e]})
                },
                onCancel() {},
            });
        },
        async postRemove(formData){
            try {
                await Remove(formData)
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
        async postTop(formData){
            try {
                await Top(formData)
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
</style>
