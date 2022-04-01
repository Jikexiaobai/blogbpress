<template>
    <div class="sidear-box">
        <div class="title">
            <h2>活跃用户</h2>
        </div>
        <ul>
            <li v-for="(item,index) in list" :key="index" @click="goProfile(item.userId)">
                <div class="sidear-user" v-if="item != null">
                    <Avatar 
                        class="user-avatar"
                        :verifyRight="-5"
                        :verifyBottom="5"
                        :isVerify="item.isVerify"
                        shape="circle" 
                        :src="item.avatar+'@w60_h60'" 
                        :size="45"
                    />
                    <div class="user-info">
                        <div class="user-info-l">
                            <h2 class="user-info-l-name">{{item.nickName}}</h2>
                            <ul class="user-info-l-meta">
                                <li  v-for="(ritem,rindex) in item.role" 
                                        :key="rindex">
                                        <img :src="ritem.icon" :alt="ritem.title">
                                    <!-- <span v-if="item.type == 3">
                                        {{item.icon}}
                                    </span> -->
                                </li>
                            </ul>
                        </div>
                        <div class="user-info-r">
                            <div class="user-info-integral">
                                <!-- <a-icon type="fire" /> -->
                                <a-icon type="trademark" /> 
                                <b>{{item.integral | resetNum}}</b>
                            </div>
                        </div>
                    </div>
                </div>
            </li>
        </ul>
    </div>
</template>

<style lang="less" scoped>
.sidear-box{
    background: white;
    margin-bottom: 10px;
    padding: 20px 20px 10px 20px;
    .title{
        margin-bottom: 10px;
        h2{
            font-size: 12px;
            line-height: 1;
            color: #999;
        }
    }
    .sidear-user{
        margin-bottom: 10px;
        display: flex;
        justify-content: space-between;
        align-items: center;
        cursor: pointer;
        .user-info{
            flex: 1;
            display: flex;
            justify-content: space-between;
            align-items: center;
            .user-info-l{
                .user-info-l-name{
                        width: 80px;
                        font-size: 13px;
                        display: -webkit-box;
                        -webkit-box-orient: vertical;
                        -webkit-line-clamp: 1;
                        overflow: hidden;
                        max-height: 52px;
                        margin-right: 5px;
                }
                .user-info-l-meta{
                    display: flex;
                    align-items: center;
                    li{
                        
                        img{
                            max-height: 24px;
                            max-width: 24px;
                        }
                    }
                }
            }
            .user-info-r{
                .user-info-integral{
                    padding: 2px 10px;
                    display: inline-block;
                    border-radius: 20px;
                    background: #f2f3ff;
                    height: 26px;
                    b{
                        font-weight: 400;
                    }
                }
            }
        }
    }
}
</style>

<script>
import Avatar from "@/components/avatar/avatar"
import api from "@/api/index"
import { mapState } from "vuex"
export default {
    components:{
        Avatar,
    },
    data(){
        return{
            // locale: zhCN,
            list:[],
        }
    },
    computed:{
        ...mapState("user",["token","userInfo"]),
    },
    mounted () {
        // this.queryParam.userId = parseInt(this.$route.params.id)
        this.getData()
    },
    methods:{
        async getData(){
            const res = await this.$axios.get(api.getSystemHotUser)
            if (res.code != 1) {
                this.$router.push(`/404`)
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
            this.list = res.data.list || []
           
        },
        goProfile(e){
            this.$router.push({ path: `/profile/${e}`})
        }
    }
}
</script>