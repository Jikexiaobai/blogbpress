<template>
    <a-card :bordered="false">
        <a-row :gutter="[40,0]">
            <a-col :span="8">
                <Form 
                    @getData="getData"
                />
            </a-col>
            <a-col :span="16">
                <Table 
                :list="list"
                 @getData="getData"
                />
            </a-col>
        </a-row>
    </a-card>
</template>

<script>
// import Select from "./components/select"
import Table from "./components/table"
import Form from "./components/form"
import { List } from '@/api/vip'
export default {
    name: 'VipList',
    components:{
        Form,
        // Select,
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
            const res = await List(this.query)
            // const catelist =  toTree(res.data.list || [],"cateId","parentId")
            this.list = res.data.list == null ? [] : res.data.list
        }
    }
}
</script>

<style lang="less" scoped>

</style>