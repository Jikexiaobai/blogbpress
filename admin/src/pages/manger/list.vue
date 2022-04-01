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
import { List } from '@/api/manger'
import Select from "./components/list/select"
import Table from "./components/list/table"
export default {
    name:"MangerList",
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
                nickName:undefined,
                email:undefined,
                phone:undefined,
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
                    nickName:item.nickName,
                    avatar:item.avatar,
                    loginIp:item.loginIp,
                    loginTime:item.loginTime,
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
            this.query.nickName = e.nickName
            this.query.email = e.email
            this.query.phone = e.phone
            this.query.status = e.status
            this.getData()
        },
    }
}
</script>

<style scoped lang="less">

</style>
