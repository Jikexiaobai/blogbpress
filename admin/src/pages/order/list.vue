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
import { List } from '@/api/order'
import Select from "./components/list/select"
import Table from "./components/list/table"
export default {
    name:"OrderList",
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
                orderNum:undefined,
                orderType:undefined,
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
                    orderNum:item.orderNum,
                    nickName:item.nickName,
                    orderType:item.orderType,
                    orderMoney:item.orderMoney,
                    paymentMoney:item.paymentMoney,
                    payMethod:item.payMethod,
                    status:item.status,
                    createTime:item.createTime,
                    status:item.status,
                    payTime:item.payTime,
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
            this.query.orderNum = e.orderNum
            this.query.orderType = e.orderType
            this.query.status = e.status
            this.getData()
        },
    }
}
</script>

<style scoped lang="less">

</style>
