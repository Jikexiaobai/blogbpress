<template>
    <a-card :bordered="false">
        <Select 
            @select="select"
        />
        <Table 
            :list="list"
            :total="total"
            :pageSize="query.limit"
            :current="query.page"
            @changePage="changePage"
            @getData="getData"
        />
    </a-card>
</template>

<script>
import Select from "./components/select"
import Table from "./components/table"
import { List } from '@/api/category'
import { toTree } from '@/utils/tree/tree'
export default {
    name: 'CategoryList',
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
                title:"",
                module:""
            },
            list:[],
            total:0
        }
    },
    mounted(){
        this.getData()
    },
    methods: {
        async getData(){
            const res = await List(this.query)
            const catelist =  toTree(res.data.list || [],"id","parentId")
      
            this.list = catelist
            this.total = res.data.total
        },
        changePage(e){
            this.query.page = e.current
            this.query.limit = e.pageSize
            this.getData()
        },
        select(e){
            this.query.title = e.title
            this.query.module = e.module
            this.getData()
        },
        // changeFormType(e){
        //     this.formType = e.type
        //     // this.formType = e
        // }
    }
}
</script>

<style lang="less" scoped>

</style>