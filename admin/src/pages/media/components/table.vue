<template>
    <div class="container">
        <a-space class="container-action" size="large">
            <a-button v-action="'/media/remove'" :disabled="selectedRowKeys.length > 0 ? false :true" @click="batchRemove" type="danger">
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
            <a-tag v-if="status == 1" color="#2db7f5">
                未使用
            </a-tag>
            <a-tag v-if="status == 2" color="#87d068">
                已使用
            </a-tag>
        </span>
        <span slot="size" slot-scope="size">
            {{size | resetSize}} 
        </span>
        <span slot="uploadKey" slot-scope="uploadKey">
            <a-tag v-if="uploadKey == 1" color="#2db7f5">
                本地
            </a-tag>
            <a-tag v-if="uploadKey == 2" color="#87d068">
                阿里云
            </a-tag>
            <a-tag v-if="uploadKey == 3" color="#87d068">
                七牛云
            </a-tag>
        </span>
        <span slot="action" slot-scope="text, record">
            <a v-action="'/media/remove'" @click="remove(record.id)">{{$t('remove')}}</a>
            <a-divider v-action="'/media/remove'" type="vertical" />
            <a @click="view(record.link)">{{$t('view')}}</a>
        </span>
        </a-table>
    </div>
</template>

<script>
import { Remove } from '@/api/media'
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
                    title: "媒体编号",
                    dataIndex: 'id',
                },
                {
                    title: "媒体原名",
                    dataIndex: 'orName',
                    scopedSlots: { customRender: 'orName' }
                },
                {
                    title: "媒体后缀",
                    dataIndex: 'ext',
                    scopedSlots: { customRender: 'ext' }
                },
                {
                    title: "媒体大小",
                    dataIndex: 'size',
                    scopedSlots: { customRender: 'size' }
                },
                {
                    title: "上传者",
                    dataIndex: 'nickName',
                    scopedSlots: { customRender: 'nickName' }
                },
                {
                    title: "所属引擎",
                    dataIndex: 'uploadKey',
                    scopedSlots: { customRender: 'uploadKey' }
                },
                {
                    title: "状态",
                    dataIndex: 'status',
                    scopedSlots: { customRender: 'status' }
                },
                {
                    title: "上传时间",
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
    filters: {
        resetSize: function(v){
            let tmp = v / 1024
            if (tmp < 1024) {
                return tmp.toFixed(2) + "kb"
            }
            if (tmp >= 1024) {
                return (tmp / 1024).toFixed(2) + "mb"
            }
        }
    },
    methods: {
        onSelectChange(selectedRowKeys) {
            this.selectedRowKeys = selectedRowKeys;
        },
        pageChange(e){
            this.$emit("changePage",e)
        },
        view(e){
            window.open(e,'_blank') // 新窗口打开外链接
        },
        batchRemove(){
            this.$confirm({
                title: "批量删除",
                onOk:() => {
                    this.postRemove({"idList":this.selectedRowKeys})
                    this.selectedRowKeys = []
                },
                onCancel() {},
            });
        },
        remove(e){
            this.$confirm({
                title: "删除",
                onOk:() => {
                    const formData = {
                        idList:[e]
                    }
                    this.postRemove(formData)
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
    },
}
</script>

<style scoped lang="less">
.container-action{
    margin-bottom: 10px;
}
</style>
