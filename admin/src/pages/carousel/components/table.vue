<template>
    <div class="container">
        <a-space class="container-action" size="large">
            <a-button v-action="'/carousel/remove'" :disabled="selectedRowKeys.length > 0 ? false :true" @click="batchRemove" type="danger">
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
        <span slot="cover" slot-scope="cover,record">
            <img class="cover" :src="cover" :alt="record.title">
        </span>
        <span slot="mode" slot-scope="mode">
            <a-tag v-if="mode == 1" color="#2db7f5">
                平台内容
            </a-tag>
            <a-tag v-if="mode == 2" color="#87d068">
                外部内容
            </a-tag>
        </span>
        <span slot="type" slot-scope="type">
            <a-tag v-if="type == 1" color="#2db7f5">
                WEB
            </a-tag>
            <a-tag v-if="type == 2" color="#87d068">
                APP
            </a-tag>
        </span>
        <span slot="action" slot-scope="text, record">
            <a @click="view(record.link)">查看</a>
            <a-divider type="vertical" />
            <a v-action="'/carousel/remove'" @click="remove(record.id)">{{$t('remove')}}</a>
        </span>
        </a-table>
    </div>
</template>

<script>
import { Remove } from '@/api/carousel'
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
                    title: "编号",
                    dataIndex: 'id',
                    
                },
                {
                    title: "类型",
                    dataIndex: 'type',
                    scopedSlots: { customRender: 'type' }
                },
                {
                    title: "封面",
                    dataIndex: 'cover',
                    scopedSlots: { customRender: 'cover' }
                },
                {
                    title: "模式",
                    dataIndex: 'mode',
                    scopedSlots: { customRender: 'mode' }
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
        view(e){
           window.open(e,'_blank') // 新窗口打开外链接
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
    },
}
</script>

<style scoped lang="less">
.container-action{
    margin-bottom: 10px;
}
.cover{
    width: 50px;
    height: 50px;
}
</style>
