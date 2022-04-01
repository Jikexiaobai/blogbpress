<template>
    <div class="container">
        <a-space v-action="'/role/create'" class="container-action" size="large">
            <a-button v-action="'/role/create'" type="primary" @click="goCreate">
                新建
            </a-button>
        </a-space>
        <a-table 
            :rowKey="record=>record.id"
            :pagination="false"
            :columns="columns" 
            :dataSource="list"
        >
        <span slot="status" slot-scope="status">
            <a-tag v-if="status == 1" color="#f50">
                禁用
            </a-tag>
            <a-tag v-if="status == 2" color="#2db7f5">
                启用
            </a-tag>
        </span>
        <span slot="action" slot-scope="text, record">
            <a v-action="'/role/edit'" @click="edit(record.id)">编辑</a>
        </span>
        </a-table>
    </div>
</template>

<script>
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
                    title: "角色标题",
                    dataIndex: 'title',
                },
                {
                    title: "角色状态",
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
        }
    },
    methods: {
        goCreate(){
            this.$router.push({path:'/role/create'})
        },
        edit(e){
            this.$router.push({path: '/role/edit',query:{id:e}});
        },
    },
}
</script>

<style scoped lang="less">
.container-action{
    margin-bottom: 10px;
}
</style>
