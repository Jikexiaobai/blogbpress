<template>
    <div class="box">
        <a-row  v-if="list.length > 0" :gutter="{ md: '30'}">
            <a-col v-for="(item,index) in list" :key="index" :span="8">
                <listTwo :info="item"/>
            </a-col> 
        </a-row>
        <div  v-if="list.length > 0" class="content-pagination" >
            <a-config-provider :locale="locale">
                <a-pagination
                    @change="changePage"
                    :pageSize="queryParam.limit"
                    :total="total"
                    show-quick-jumper
                >
                </a-pagination>
            </a-config-provider>
        </div>
        <div v-if="list.length == 0" class="empty">
            <a-config-provider :locale="locale">
                <a-empty />
            </a-config-provider>
        </div>
    </div>
</template>

<style lang="less" scoped>
.box{
    padding: 20px;
    .content-pagination{
        margin: 20px 0;
    }
}
</style>

<script>
import ListTwo from "@/components/list/listTwo"

import zhCN from 'ant-design-vue/lib/locale-provider/zh_CN'
import {MODULE} from "@/shared/module"
import api from "@/api/index"
export default {
    components:{
       ListTwo,
    },
    data(){
        return{
            locale: zhCN,
            MODULE,
            queryParam:{
                page:1,
                limit: 8,
                userId:null,
                mode: 1,
                module:MODULE.AUDIO
            },
            id:null,
            list:[],
            total:0,
        }
    },
    mounted () {
        this.queryParam.userId = parseInt(this.$route.params.id)
        this.getData()
    },
    methods:{
        async getData(){
            const res = await this.$axios.get(api.getUserPosts,{params:this.queryParam})
            if (res.code != 1) {
                this.$router.push(`/404`)
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
            this.list = res.data.list || []
            this.total = res.data.total || 0
        },
        changePage(page,limit){
            this.queryParam.limit = limit
            this.queryParam.page = page
            this.getData()
        },
    }
}
</script>