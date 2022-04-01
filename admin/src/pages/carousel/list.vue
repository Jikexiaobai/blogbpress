<template>
    <div class="container">
        <a-card :bordered="false">
            <a-row :gutter="[40,0]">
                <a-col :span="8">
                    <Form 
                        @getData="getData"
                    />
                </a-col>
                <a-col :span="16">
                    <Select  @select="select"/>
                    <Table 
                        :list="list"
                        :total="total"
                        :pageSize="query.limit"
                        :current="query.page"
                        @changePage="changePage"
                        @getData="getData"
                    />
                </a-col>
            </a-row>
        </a-card>
    </div>
</template>

<script>
import { List } from '@/api/carousel'
import Select from "./components/select"
import Table from "./components/table"
import Form from "./components/form"
export default {
    name:"AudioList",
    components: {
        Select,
        Table,
        Form
    },
    data() {
        return {
            query:{
                page:1,
                limit:10,
                mode:undefined,
                type:undefined,
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
                    type:item.type,
                    mode:item.mode,
                    link:item.link,
                    cover:item.cover,
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
            this.query.mode = e.mode
            this.query.type = e.type
            this.getData()
        },
    }
}
</script>

<style scoped lang="less">

</style>
