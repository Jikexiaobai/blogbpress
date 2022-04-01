<template>
    <div class="follow-box">
        <ul v-if="list.length > 0">
            <li v-for="(item,index) in list" :key="index">
                <div class="follow-item">
                    <nuxt-link :to="{path:'/profile/' + item.id }" class="item-link"> 
                        <Avatar  :src="item.avatar+'@w60_h60'" class="user-avatar" :size="60" />
                    </nuxt-link>
                    <div class="user-info">
                        <div class="user-name">
                            <nuxt-link :to="{path:'/profile/' + item.id }" class="item-link"> 
                                <h2>{{item.nickName}}</h2>
                            </nuxt-link>
                            <p>{{item.description}}</p>
                        </div>
                        <a-button @click="follow(index)" v-if="!item.isFollow && item.id != userInfo.userId" type="primary">
                            关注
                        </a-button>
                    </div>
                </div>
                <a-divider/>
            </li>
        </ul>
        <div class="content-pagination" v-if="list.length > 0" >
            <a-config-provider :locale="locale">
                <a-pagination
                    @change="changePage"
                    :pageSize="queryParam.limit"
                    :total="total"
                    show-quick-jumper
                />
            </a-config-provider>
        </div>
        <div v-if="list.length == 0" class="empty">
            <a-config-provider :locale="locale">
                <a-empty />
            </a-config-provider>
        </div>
    </div>
</template>

<script>
import zhCN from 'ant-design-vue/lib/locale-provider/zh_CN'
import Avatar from "@/components/avatar/avatar"
import api from "@/api/index"
import { mapState } from "vuex"
export default {
    components:{
        Avatar,
    },
    data(){
        return{
            locale: zhCN,
            queryParam:{
                page:1,
                limit: 8,
                userId:null,
                related:"fans",
            },
            id:null,
            list:[],
            total:0,
        }
    },
    computed:{
        ...mapState("user",["token","userInfo"]),
    },
    mounted () {
        this.queryParam.userId = parseInt(this.$route.params.id)
        this.getData()
    },
    methods:{
        async getData(){
            const res = await this.$axios.get(api.getUserFansOrFollows,{params:this.queryParam})
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
        async follow(i){
            if (!this.token) {
                this.$Auth("login","登录","快速登录")
                return
            }
            if (this.list[i].isFollow) {
                this.list[i].fans -= 1
            }else{
                this.list[i].fans += 1
            }
            this.list[i].isFollow = !this.list[i].isFollow
            const res = await this.$axios.post(api.postUserFollow,{id:this.list[i].id})
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
                if (this.list[i].isFollow) {
                    this.list[i].fans -= 1
                }else{
                    this.list[i].fans += 1
                }
                this.list[i].isFollow = !this.list[i].isFollow
                return
            }
        },
    }
}
</script>

<style lang="less" scoped>
.follow-box{
    margin: 10px 0;
    background-color: white;
    .empty{
        display: flex;
        justify-content: center;
        align-items: center;
        padding-bottom: 20px;
        height: 300px;
    }
    ul{
        padding: 20px;
        .follow-item{
            display: flex;
            justify-content: space-between;
            align-items: center;
            .user-info{
                flex: 1;
                display: flex;
                justify-content: space-between;
                align-items: center;
                .user-name{
                    h2{
                        font-weight: 400;
                        font-size: 18px;
                    }
                    p{
                        line-height: 14px;
                        font-size: 12px;
                        color: #6d757a;
                        overflow: hidden;
                        text-overflow: ellipsis;
                        white-space: nowrap;
                    }
                }
            }
        }
    }
    .content-pagination{
        margin-top: 20px;
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
    }
}
</style>