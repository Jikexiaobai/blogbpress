<template>
    <a-card :bordered="false">
        <Table 
            :list="list"
        />
    </a-card>
</template>

<script>
import Table from "./components/table"
import { List } from '@/api/role'
export default {
    name: 'RoleList',
    components:{
        Table
    },
    data(){
        return{
            list:[],
        }
    },
    mounted(){
        this.getData()
    },
    methods: {
        async getData(){
            const res = await List()
            if (res.code != 1) {
                this.$message.error(res.message);
                return
            }
            this.list = res.data.list == null ? [] : res.data.list
        },
    }
}
</script>

<style lang="less" scoped>

</style>