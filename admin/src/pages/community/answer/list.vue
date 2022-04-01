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
import { List } from '@/api/answer'
import Select from "./components/list/select"
import Table from "./components/list/table"
export default {
    name:"AnswerList",
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
                userId:undefined,
                isPay:0,
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
                    content:item.content,
                    nickName:item.nickName,
                    price:item.price,
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
            this.query.userId = e.userId
            this.query.status = e.status
            this.query.isPay = e.isPay
            this.getData()
        },
    }
}
</script>

<style scoped lang="less">

</style>