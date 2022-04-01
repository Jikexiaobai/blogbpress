<template>
    <div class="sidear-box">
        <div class="clock-in-box">
            
            <div v-if="userInfo.sign == null" class="clock-in-btn">
                <a-button @click="clockIn" type="primary" icon="gift" block>
                    点击领取今天的签到奖励
                </a-button>
            </div>
            <div v-if="userInfo.sign != null" class="clock-in-desc">
                <a-icon type="gift" />
                恭喜！您今天获得了{{userInfo.sign.integral}}积分
            </div>
        </div>
        <ul class="menu">
            <li>
                <button 
                    :class="queryParam.type == 1 ? 'active': ''" 
                    @click="changeType(1)">
                    今日签到
                </button>
            </li>
            <li>
                <button 
                    :class="queryParam.type == 2 ? 'active': ''" 
                    @click="changeType(2)">
                    连续签到
                </button>
            </li>
        </ul>
        <ul>
            <li v-for="(item,index) in list" :key="index" @click="goProfile(item.id)">
                <div class="sidear-user">
                    <Avatar 
                        class="user-avatar"
                        :verifyRight="-5"
                        :verifyBottom="5"
                        :isVerify="item.isVerify"
                        shape="square" 
                        :src="item.avatar+'@w60_h60'" 
                        :size="40"
                    />
                    <div class="user-info">
                        <div class="user-info-l">
                            <h2 class="user-info-l-name">{{item.nickName}}</h2>
                            <div class="user-info-l-date">{{item.createTime | resetData}}</div>
                        </div>
                        <div class="user-info-r">
                            <div v-if="queryParam.type == 1" class="user-info-integral">
                                <!-- <a-icon type="fire" /> -->
                                <a-icon type="trademark" /> 
                                <b>{{item.integral | resetNum}}</b>
                            </div>
                            <div  v-if="queryParam.type == 2"  class="user-info-count">
                                连续{{item.count}}天
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
    .clock-in-box{
        .clock-in-btn{
            padding-bottom: 15px;
        }
        .clock-in-desc{
            padding-bottom: 15px;
            font-size: 13px;
            text-align: center;
            background-image: linear-gradient(
            90deg
            ,#673AB7 0%, #E91E63 50%);
            background-clip: text;
            -webkit-text-fill-color: transparent;
        }
        
    }
    .menu{
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        flex: 1;
        margin-bottom: 20px;
        li{
            font-size: 14px;
            margin-right: 8px;
            button{
                cursor: pointer;
                background: 0 0;
                border: 0;
                color: initial;
                padding: 5px 10px;
                border-radius: 2px;
                -webkit-appearance: none;
                outline: none;
                -webkit-tap-highlight-color: rgba(0,0,0,0);
                font-family: font-regular,'Helvetica Neue',sans-serif;
                // border: 1px solid #ccc;
                box-sizing: border-box;
                user-select: none;
            }
            .active{
                background-color: #4560c9;
                color: #fff;
            }
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
                    margin-bottom: 5px;
                }
                .user-info-l-date{
                    font-size: 12px;
                    color: #bcbcbc;
                }
            }
            .user-info-r{
                .user-info-integral{
                    padding: 2px 10px;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    border-radius: 20px;
                    background: #f2f3ff;
                    height: 26px;
                    font-size: 13px;
                    b{
                        
                        margin-left: 5px;
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
import { mapState, mapActions} from "vuex"
export default {
    components:{
        Avatar,
    },
    data(){
        return{
            isSign:false,
            integral:0,
            // locale: zhCN,
            queryParam:{
                page:1,
                limit: 6,
                type:1,
            },
            list:[],
        }
    },
    computed:{
        ...mapState("user",["token","userInfo"]),
    },
    mounted () {
        this.getList()
    },
    methods:{
        ...mapActions("user",["A_UPDATE_SIGN"]),
        async getList(){
            const res = await this.$axios.get(api.getUserSign,{params:this.queryParam})
            this.list = res.data.list == null ? [] : res.data.list
        },
        async clockIn(){
            if (this.token == null) {
                this.$Auth("login","登录","快速登录")
                return
            }

            const res = await this.$axios.post(api.postAccountSign)
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
            this.getList()
            this.A_UPDATE_SIGN({
                isSign:true,
                integral:res.data.integral
            })
           
        },
        goProfile(e){
            // console.log(e)
            this.$router.push({ path: `/profile/${e}`})
        },
        changeType(e){
            this.queryParam.type = e
            this.getList()
        }
    }
}
</script>