<template>
    <div class="container">
        <a-table 
            :rowKey="record=>record.perms"
            :pagination="false"
            :columns="columns" 
            :dataSource="list"
        >
        <span slot="type" slot-scope="type">
            <a-tag v-if="type == 1" color="#2db7f5">
                目录
            </a-tag>
            <a-tag v-if="type == 2" color="#87d068">
                菜单
            </a-tag>
            <a-tag v-if="type == 3" color="#f50">
                按钮
            </a-tag>
        </span>
        <span slot="icon" slot-scope="icon">
            <a-icon v-if="icon != ''" :type="icon" />
        </span>
        <span slot="action" slot-scope="text, record">
            <a v-action="'/authority/edit'" @click="edit(record)">{{$t('edit')}}</a>
            <a-divider type="vertical" />
            <a v-action="'/authority/remove'" @click="remove(record)">{{$t('remove')}}</a>
        </span>
        </a-table>
    </div>
</template>

<script>
import { Remove } from '@/api/authority'
export default {
    props:{
        list:{
            type:Array,
                // 当为数组类型设置默认值时必须使用数组返回
            required:true,
            default: []
        },
    },
    data() {
        return {
            // 表格
            columns:[
                {
                    title: "权限名称",
                    dataIndex: 'title',
                },
                {
                    title: "路由地址",
                    dataIndex: 'path',
                    scopedSlots: { customRender: 'path' }
                },
                {
                    title: "权限类型",
                    dataIndex: 'type',
                    scopedSlots: { customRender: 'type' }
                },
                {
                    title: "权限标识",
                    dataIndex: 'perms',
                    scopedSlots: { customRender: 'perms' }
                },
                {
                    title: "权限图标",
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
        }
    },
    methods: {
        edit(e){
            this.$router.push({path: '/authority/edit',query:{id:e.id}});
        },
        remove(e){
            this.$confirm({
                okText: "确定",
                cancelText: "取消",
                title: "正在删除",
                onOk:() => {
                    this.postRemove({"id":e.id})
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

</style>
