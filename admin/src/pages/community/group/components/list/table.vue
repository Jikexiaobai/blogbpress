<template>
    <div class="container">
        <a-space class="container-action" size="large">
            <a-popconfirm
                v-action="'/group/review'"
                :disabled="selectedRowKeys.length > 0 ? false :true"
                title="正在审核"
                ok-text="通过"
                cancel-text="拒绝"
                @confirm="batchReview(2)"
                @cancel="batchReview(3)"
            >
                <a-button v-action="'/group/review'" :disabled="selectedRowKeys.length > 0 ? false :true">
                    批量审核
                </a-button>
            </a-popconfirm>
            <a-button
                v-action="'/group/recover'"
                v-if="status != 4"
                :disabled="selectedRowKeys.length > 0 ? false :true" 
                @click="batchRecover" type="danger">
                批量移入回收站
            </a-button>
           <a-button
                v-action="'/group/reduction'"
                v-if="status == 4"
                :disabled="selectedRowKeys.length > 0 ? false :true" 
                @click="batchReduction" type="primary">
                批量还原
            </a-button>
            <a-button
                v-action="'/group/remove'"
                v-if="status == 4"
                :disabled="selectedRowKeys.length > 0 ? false :true" 
                @click="batchRemove" type="danger">
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
        <span slot="status" slot-scope="status">
            <a-tag v-if="status == 1" color="#2db7f5">
                未审核
            </a-tag>
            <a-tag v-if="status == 2" color="#87d068">
                已通过
            </a-tag>
            <a-tag  v-if="status == 3" color="#f50">
                未通过
            </a-tag>
        </span>
        <span slot="action" slot-scope="text, record">
            <!-- <a @click="edit(record.id)"   v-if="status != 4">{{$t('edit')}}</a> -->
            <a @click="reduction(record.id)"  v-action="'/group/reduction'"  v-if="status == 4">还原</a>   
            <a-divider type="vertical" />
            <a @click="remove(record.id)"  v-action="'/group/remove'"  v-if="status == 4">删除</a> 
            <a-dropdown v-if="status != 4">
                <a-menu slot="overlay">
                    <a-menu-item v-action="'/group/review'"  v-if="record.status == 1">
                        <a-popconfirm
                            title="审核"
                            @confirm="review(2,record.id)"
                            @cancel="review(3,record.id)"
                        >
                            <a href="#">审核</a>
                        </a-popconfirm>
                    </a-menu-item>
                    <a-menu-item v-action="'/group/recover'">
                        <a @click="recover(record.id)">移入回收站</a>
                    </a-menu-item>
                </a-menu>
                <a>更多<a-icon type="down"/></a>
            </a-dropdown>
        </span>
        </a-table>

        <a-modal v-model="visible" 
            :title="$t('modal.remark.title')"
            @ok="handleOk">
            <a-textarea v-model="remark" :rows="4" />
        </a-modal>
    </div>
</template>

<script>
import { Remove,Review,Reduction,Recover } from '@/api/group'
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
        status:{
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
                    title: "标题",
                    dataIndex: 'title',
                    scopedSlots: { customRender: 'title' }
                },
                {
                    title: "分类",
                    dataIndex: 'cate',
                    scopedSlots: { customRender: 'cate' }
                },
                {
                    title: "圈主",
                    dataIndex: 'author',
                    scopedSlots: { customRender: 'author' }
                },
                {
                    title: "封面",
                    dataIndex: 'cover',
                    scopedSlots: { customRender: 'cover' }
                },
                {
                    title: "已加入人数",
                    dataIndex: 'joins',
                    scopedSlots: { customRender: 'joins' }
                },
                {
                    title: "内容数",
                    dataIndex: 'contents',
                    scopedSlots: { customRender: 'contents' }
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
            this.$router.push({path:'/group/create'})
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
                    this.$message.error("请输入审核信息");
                }
                if (this.action == "recover") {
                    this.$message.error("请输入回收信息");
                }
                return
            }
            const formData = this.reviewFormData
            formData.remark = this.remark
            if (this.action == "review") {
                this.postReview(formData)
                this.selectedRowKeys = []
            }
            if (this.action == "recover") {
                this.postRecover(formData)
                this.selectedRowKeys = []
            }
           
            this.visible = false
        },
        batchRecover(){
            this.$confirm({
                title: "正在移到回收站",
                onOk:() => {
                    const formData = {
                        idList:this.selectedRowKeys
                    }
                    this.action = undefined
                    this.remark = undefined
                    this.reviewFormData = undefined
                    this.action = "recover"
                    this.visible = true
                    this.reviewFormData = formData
                    this.selectedRowKeys = []
                },
                onCancel() {},
            });
        },
        recover(e){
            this.$confirm({
                title: "正在移到回收站",
                onOk:() => {
                    const formData = {
                        "idList":[e]
                    }
                    this.action = undefined
                    this.remark = undefined
                    this.reviewFormData = undefined
                    this.action = "recover"
                    this.visible = true
                    this.reviewFormData = formData
                },
                onCancel() {},
            });
        },
       
        batchReduction(){
            this.$confirm({
                title: "正在还原",
                onOk:() => {
                    const formData = {
                        idList:this.selectedRowKeys
                    }
                    this.postReduction(formData)
                    this.selectedRowKeys = []
                },
                onCancel() {},
            });
        },
        reduction(e){
            this.$confirm({
                title: "正在还原",
                onOk:() => {
                    const formData = {
                        "idList":[e]
                    }
                    this.postReduction(formData)
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

        // edit(e){
        //     this.$router.push({path: '/group/edit',query:{id:e}});
        // },
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
        async postReduction(formData){
            try {
                const res = await Reduction(formData)
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
        async postRecover(formData){
            try {
                const res = await Recover(formData)
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
