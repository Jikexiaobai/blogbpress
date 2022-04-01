<template>
    <div class="container">
        <a-card :bordered="false">
            <Select  @select="select"/>
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
</template>

<script>
import { List } from '@/api/report'
import Select from "./components/list/select"
import Table from "./components/list/table"
export default {
    name:"ReportList",
    components: {
        Select,
        Table
    },
    data() {
        return {
            query:{
                page:1,
                limit:10,
                type:undefined,
                status:undefined,
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
     
            this.list = res.data.list == null ? [] :  res.data.list.map((item)=>{
                const tmp = {
                    id:item.id,
                    relatedId:item.relatedId,
                    title:item.title,
                    nickName:item.nickName,
                    module:item.module,
                    type:item.type,
                    description:item.description,
                    status:item.status,
                    createTime:item.createTime,
                }
                return tmp
            })
        
            this.total = res.data.total
        },
        changePage(e){
            this.query.page = e.current
            this.query.limit = e.pageSize
            this.getData()
        },
        select(e){
            this.query.type = e.type
            this.query.status = e.status
            this.getData()
        },
    }
}
</script>

<style scoped lang="less">

</style>