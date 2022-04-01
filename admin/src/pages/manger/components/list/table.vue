<template>
    <div class="container">
        <a-space class="container-action" size="large">
            <a-button v-action="'/manger/create'" type="primary" @click="goCreate">
                新建
            </a-button>
            <a-button v-action="'/manger/remove'" :disabled="selectedRowKeys.length > 0 ? false :true" @click="batchRemove" type="danger">
                删除
            </a-button>
            <a-popconfirm
                :disabled="selectedRowKeys.length > 0 ? false :true"
                title="修改管理状态"
                ok-text="启用"
                cancel-text="禁用"
                @confirm="batchReview(2)"
                @cancel="batchReview(1)"
            >
                <a-button v-action="'/manger/review'" :disabled="selectedRowKeys.length > 0 ? false :true">
                    批量修改状态
                </a-button>
            </a-popconfirm>
        </a-space>
        <a-table 
            :rowKey="record=>record.id"
            :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
            :pagination="{
                showQuickJumper:true,
                showSizeChanger:true,
                showTotal:(total)=>`共${total}$条`,
                pageSize:pageSize,
                current:current,
                total:total
            }"
            :columns="columns" 
            :dataSource="list"
            @change="pageChange"
        >
        <span slot="avatar" slot-scope="avatar,record">
            <img class="avatar" :src="avatar" :alt="record.title">
        </span>
        <span slot="status" slot-scope="status">
            <a-tag v-if="status == 1" color="#f50">
                禁用
            </a-tag>
            <a-tag v-if="status == 2" color="#87d068">
                启用
            </a-tag>
        </span>

        <span slot="action" slot-scope="text, record">
            <a v-action="'/manger/edit'" @click="edit(record.id)">编辑</a>   
            <a-divider type="vertical" />
            <a-dropdown>
                <a-menu slot="overlay">
                    <a-menu-item v-action="'/manger/review'">
                        <a-popconfirm
                            title="修改状态"
                            ok-text="启用"
                            cancel-text="禁用"
                            @confirm="review(2,record.id)"
                            @cancel="review(1,record.id)"
                        >
                            <a href="#">修改状态</a>
                        </a-popconfirm>
                    </a-menu-item>
                    <a-menu-item v-action="'/manger/remove'">
                        <a @click="remove(record.id)">删除</a>   
                    </a-menu-item>
                </a-menu>
                <a>{{$t('more')}}<a-icon type="down"/></a>
            </a-dropdown>
        </span>
        </a-table>
    </div>
</template>

<script>
import { Remove,Review } from '@/api/manger'
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
                    title: this.$t('user.list.table.userId.title'),
                    dataIndex: 'id',
                    
                },
                {
                    title: this.$t('user.list.table.nickName.title'),
                    dataIndex: 'nickName',
                    scopedSlots: { customRender: 'nickName' }
                },
                {
                    title: this.$t('user.list.table.avatar.title'),
                    dataIndex: 'avatar',
                    scopedSlots: { customRender: 'avatar' }
                },
                {
                    title: this.$t('user.list.table.loginIp.title'),
                    dataIndex: 'loginIp',
                    scopedSlots: { customRender: 'loginIp' }
                },
                {
                    title: this.$t('user.list.table.loginTime.title'),
                    dataIndex: 'loginTime',
                    scopedSlots: { customRender: 'loginTime' }
                },
                {
                    title: this.$t('user.list.table.status.title'),
                    dataIndex: 'status',
                    scopedSlots: { customRender: 'status' }
                },
                {
                    title: this.$t('user.list.table.createTime.title'),
                    dataIndex: 'createTime',
                },
                {
                    title: this.$t('table.action.title'),
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
        goCreate(){
            this.$router.push({path:'/manger/create'})
        },
        edit(e){
            this.$router.push({path: '/manger/edit',query:{id:e}});
        },
        pageChange(e){
            this.$emit("changePage",e)
        },
        batchRemove(){
            this.$confirm({
                title: "正在删除",
                content: "是否确定要删除",
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
                content: "是否确定要删除",
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
