<template>
    <div class="container">
        <a-card :bordered="false">
            <Select  @select="select"/>
            <Table 
                :list="list"
                :total="total"
                :pageSize="query.limit"
                :current="query.page"
                :status="query.status"
                @changePage="changePage"
                @getData="getData"
            />
        </a-card>
    </div>
</template>

<script>
import { List } from '@/api/topic'
import Select from "./components/list/select"
import Table from "./components/list/table"
export default {
    name:"TopicList",
    components: {
        Select,
        Table
    },
    data() {
        return {
            query:{
                page:1,
                limit:10,
                status:undefined,
                isTop:undefined,
                title:undefined,
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
                    title:item.title,
                    author:item.nickName,
                    status:item.status,
                    isTop:item.isTop,
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
            this.query.title = e.title
            this.query.status = e.status
            this.query.isTop = e.isTop
            this.getData()
        },
    }
}
</script>

<style scoped lang="less">

</style>
