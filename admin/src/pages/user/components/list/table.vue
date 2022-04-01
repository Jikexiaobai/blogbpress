<template>
    <div class="container">
        <a-space class="container-action" size="large">
            <a-button v-action="'/user/create'" type="primary" @click="goCreate">
                新建
            </a-button>
            <a-button v-action="'/user/remove'" :disabled="selectedRowKeys.length > 0 ? false :true" @click="batchRemove" type="danger">
                批量删除
            </a-button>
            <a-popconfirm
                v-action="'/user/review'"
                :disabled="selectedRowKeys.length > 0 ? false :true"
                title="修改用户状态"
                ok-text="启用"
                cancel-text="禁用"
                @confirm="batchReview(2)"
                @cancel="batchReview(1)"
            >
                <a-button  :disabled="selectedRowKeys.length > 0 ? false :true">
                    批量修改用户状态
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
        <span slot="avatar" slot-scope="avatar,record">
            <img class="avatar" :src="avatar" :alt="record.title">
        </span>
        <span slot="integral" slot-scope="integral">
            {{integral}}
        </span>
        <span slot="balance" slot-scope="balance">
            {{balance}} ￥
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
            <a v-action="'/user/edit'" @click="edit(record.id)">编辑</a>   
            <a-divider type="vertical" />
            <a-dropdown>
                <a-menu slot="overlay">
                    <a-menu-item v-action="'/user/review'" >
                        <a-popconfirm
                            title="修改用户状态"
                            ok-text="启用"
                            cancel-text="禁用"
                            @confirm="review(2,record.id)"
                            @cancel="review(1,record.id)"
                        >
                            <a href="#">修改状态</a>
                        </a-popconfirm>
                    </a-menu-item>
                    <a-menu-item v-action="'/user/remove'">
                        <a @click="remove(record.id)">删除</a>   
                    </a-menu-item>
                </a-menu>
                <a>更多<a-icon type="down"/></a>
            </a-dropdown>
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
import { Remove,Review } from '@/api/user'
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
                    dataIndex: 'id',
                    
                },
                {
                    title: "昵称",
                    dataIndex: 'nickName',
                    scopedSlots: { customRender: 'nickName' }
                },
                {
                    title: "头像",
                    dataIndex: 'avatar',
                    scopedSlots: { customRender: 'avatar' }
                },
                {
                    title: "积分",
                    dataIndex: 'integral',
                    scopedSlots: { customRender: 'integral' }
                },
                {
                    title: "余额",
                    dataIndex: 'balance',
                    scopedSlots: { customRender: 'balance' }
                },
                {
                    title: "登录地址",
                    dataIndex: 'loginIp',
                    scopedSlots: { customRender: 'loginIp' }
                },
                {
                    title: "登录时间",
                    dataIndex: 'loginTime',
                    scopedSlots: { customRender: 'loginTime' }
                },
                {
                    title: "状态",
                    dataIndex: 'status',
                    scopedSlots: { customRender: 'status' }
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
            action:"review"
        }
    },
    methods: {
        onSelectChange(selectedRowKeys) {
            this.selectedRowKeys = selectedRowKeys;
        },
        goCreate(){
            this.$router.push({path:'/user/create'})
        },
        pageChange(e){
            this.$emit("changePage",e)
        },
        batchReview(e){
            const formData ={
                status:e,
                idList:this.selectedRowKeys
            }
            this.action = undefined
            this.remark = undefined
            this.reviewFormData = undefined
            this.action = "review"
            this.visible = true
            this.reviewFormData = formData
            this.selectedRowKeys = []
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
            if (this.remark == null || this.remark == undefined ||this.remark == "") {
                if (this.action == "review") {
                    this.$message.error("请输入处理信息");
                }
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
        edit(e){
            this.$router.push({path: '/user/edit',query:{id:e}});
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
