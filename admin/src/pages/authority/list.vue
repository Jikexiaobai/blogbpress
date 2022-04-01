<template>
    <a-card :bordered="false">
        <Select 
            @select="select"
        />
        <Table 
        :list="list"
        @getData="getData"
        />
    </a-card>
</template>

<script>
import Select from "./components/list/select"
import Table from "./components/list/table"
import { List } from '@/api/authority'
import { loopMenu } from '@/utils/tree/menu'
export default {
    name: 'AuthorityList',
    components:{
        Select,
        Table
    },
    data(){
        return{
            // formType:"create",
            query:{
                page:1,
                limit:10,
                title:""
            },
            list:[],
        }
    },
    mounted(){
        this.getData()
    },
    methods: {
        async getData(){
            const res = await List(this.query)

            if (res.code != 1) {
                this.$message.error(res.message);
                return
            }
            let menuLoot = loopMenu(res.data.list == null ? [] : res.data.list)
            this.list = menuLoot
        },
        select(e){
            this.query.title = e.title
            this.getData()
        },
    }
}
</script>

<style lang="less" scoped>

</style>