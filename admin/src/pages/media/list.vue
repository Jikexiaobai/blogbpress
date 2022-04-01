<template>
    <a-card :bordered="false">
        <div class="container">
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
        </div>
    </a-card>
</template>

<script>
import Select from "./components/select"
import Table from "./components/table"
import { List } from '@/api/media'

export default {
    name: 'MediaList',
    components:{
        Select,
        Table
    },
    data(){
        return{
            query:{
                page:1,
                limit:10,
                orName:"",
                status:0
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
            this.list = res.data.list
            this.total = res.data.total
        },
        changePage(e){
            this.query.page = e.current
            this.query.limit = e.pageSize
            this.getData()
        },
        select(e){
            this.query.orName = e.orName
            this.query.status = e.status
            this.getData()
        },
    }
}
</script>

<style lang="less" scoped>

</style>