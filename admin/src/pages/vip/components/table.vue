<template>
    <div class="container">
        <a-table 
            :rowKey="record=>record.id"
            :pagination="false"
            :columns="columns" 
            :dataSource="list"
            @change="pageChange"
        >
        <span slot="icon" slot-scope="icon,record">
            <img class="icon" :src="icon" :alt="record.title">
        </span>

        <span v-action="'/vip/edit'" slot="action" slot-scope="text, record">
            <a @click="edit(record.id)">编辑</a>
        </span>
        </a-table>
        <a-modal v-model="visible" 
            title="修改内容"
            @ok="handleOk">
            <a-form-model 
                ref="form" :model="form" :rules="form.rules">
                <a-form-model-item
                    ref="title" 
                    prop="title">
                    <a-input
                        size="large"
                        type="text"
                        placeholder="请输入会员标题"
                        v-model="form.title"
                    >
                    </a-input>
                </a-form-model-item>
                <a-form-model-item 
                    ref="day" 
                    prop="day">
                    <a-input-number
                        size="large" 
                        v-model="form.day"
                        placeholder="过期天数"
                        :style="{ width: '100%' }"
                        :min="0"
                    />
                </a-form-model-item>
                <a-form-model-item 
                    ref="price" 
                    prop="price">
                    <a-input-number
                        size="large" 
                        v-model="form.price"
                        placeholder="开通价格"
                        :style="{ width: '100%' }"
                        :min="0"
                        :precision="2"
                    />
                </a-form-model-item>
                <a-form-model-item 
                    ref="discount" 
                    prop="discount">
                    <a-input-number
                        size="large" 
                        v-model="form.discount"
                        placeholder="会员折扣"
                        :style="{ width: '100%' }"
                        :min="0"
                        :max="0.99"
                        :precision="2"
                    />
                </a-form-model-item>
                <a-form-model-item 
                    ref="color" 
                    prop="color">
                    <a-input
                        size="large"
                        type="text"
                        placeholder="自定义颜色"
                        v-model="form.color"
                    >
                    </a-input>
                </a-form-model-item>
                <a-form-model-item 
                    ref="icon" 
                    prop="icon">
                    <div class="upload-icon">
                        <a-input
                            size="large"
                            type="text"
                            placeholder="会员图标"
                            v-model="form.icon"
                        />
                        <a-button v-action="'/media/upload'" class="form-button" @click="upload" size="large" type="primary" style="margin-left: 10px;">
                            上传
                        </a-button>
                    </div>
                </a-form-model-item>
            </a-form-model>
        </a-modal>
    </div>
</template>

<script>
import { Remove,EditInfo,Edit } from '@/api/vip'
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
            // 表格
            columns:[
                {
                    title: "编号",
                    dataIndex: 'id',
                },
                {
                    title: "标题",
                    dataIndex: 'title',
                },
                {
                    title: "开通价格",
                    dataIndex: 'price',
                    scopedSlots: { customRender: 'price' }
                },
                {
                    title: "会员折扣",
                    dataIndex: 'discount',
                    scopedSlots: { customRender: 'discount' }
                },
                {
                    title: "过期天数",
                    dataIndex: 'day',
                    scopedSlots: { customRender: 'day' }
                },
                {
                    title: "自定义颜色",
                    dataIndex: 'color',
                    scopedSlots: { customRender: 'color' }
                },
                {
                    title: "图标",
                    dataIndex: 'icon',
                    scopedSlots: { customRender: 'icon' }
                },
                {
                    title: "操作",
                    width: '150px',
                    dataIndex: 'action',
                    scopedSlots: { customRender: 'action' }
                }
            ],
            form:{
                title:undefined,
                day:undefined,
                price:undefined,
                discount:undefined,
                color:undefined,
                icon:undefined,
                vipId:undefined,
                rules:{
                    title:[
                        { required: true, message: "请输入会员标题", trigger: 'blur' },
                    ],
                    price:[
                        { required: true, message: "请设置开通价格", trigger: 'blur' },
                    ],
                    icon:[
                        { required: true, message: "请上传图标", trigger: 'blur' },
                    ],
                },
            }
        }
    },
    methods: {
        pageChange(e){
            this.$emit("changePage",e)
        },
        async edit(e){
            const editInfo = await EditInfo({id:e})
            if (editInfo.code != 1) {
                this.$router.push({path:'/404'})
            }
           
            this.form = Object.assign(this.form,editInfo.data.info)
            this.form.vipId = e
            this.visible = true
        },
        upload(){
            this.visible = false
            this.$Upload().then((res)=>{
                if (res != false) {
                    this.form.icon = res
                    this.visible = true
                }
            }).catch((err)=>{
                console.log(err)
            })
        },
        handleOk(){
            this.$confirm({
                title: "正在修改",
                onOk:() => {
                    this.$refs.form.validate(valid => {
                        if (valid) {
                            this.postEdit(this.form)
                        } else {
                            return false;
                        }
                    });
                },
                onCancel() {},
            });
            this.visible = false
        },
        batchRemove(){
            this.$confirm({
                okText: this.$t('confirm.ok.text'),
                cancelText:this.$t('confirm.cancel.text'),
                title: this.$t('confirm.remove.title.text'),
                content: this.$t('confirm.remove.content.text'),
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
                okText: this.$t('confirm.ok.text'),
                cancelText:this.$t('confirm.cancel.text'),
                title: this.$t('confirm.remove.title.text'),
                content: this.$t('confirm.remove.content.text'),
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
        async postEdit(formData){
            try {
                const res = await Edit(formData)
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
.upload-icon{
    display: flex;
}
.icon{
    width: 50px;
    height: 50px;
}
</style>
