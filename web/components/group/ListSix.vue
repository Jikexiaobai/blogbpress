<template>
    <div class="user-info">
        <div class="user-info-cover">
            <img :src="info.cover" :alt="info.nickName">
        </div>
        <div class="user-info-avatar-name">
            <Avatar @click="goInfo" :src="info.avatar+'@w60_h60'" :size="56"/>
            <div class="user-info-name">
                <h2 @click="goInfo">{{info.nickName}}</h2>
                <div>
                    <a-space>
                        <img :src="info.grade.icon" :alt="info.grade.title">
                        <img v-if="info.vip" :src="info.vip.icon" :alt="info.vip.title">
                    </a-space>
                </div>
            </div>
        </div>
        <ul class="user-info-meta">
            <li>
                <span>粉丝</span>
                <p>{{info.fans | resetNum}}</p>
            </li>
            <li>
                <span>获赞</span>
                <p>{{info.likes | resetNum}}</p>
            </li>
            <li>
                <span>关注</span>
                <p>{{info.follows | resetNum}}</p>
            </li>
        </ul>
        <div class="user-info-desc">
            {{info.description}}
        </div>
    </div>
</template>


<script>
import Avatar from "@/components/avatar/avatar"
export default {
    components:{
       Avatar
    },
    props: {
        info:{
            type: Object,
            default: null
        }
    },
    methods:{
        goInfo(){
            this.$router.push({ path: `/profile/${this.info.id}`})
        }
    }
}
</script>

<style lang="less" scoped>
.user-info{
    
    background: white;
    .user-info-cover{
        position: relative;
        height: 0;
        padding-top: 38%;
        overflow: hidden;
        img{
            position: absolute;
            top: 0;
            left: 0;
            height: 100%;
            width: 100%;
        }          
    }
    .user-info-avatar-name{
        cursor: pointer;
        display: flex;
        flex-flow: wrap;
        margin-top: -14px;
        position: relative;
        .user-info-name{
            margin-top: 10px;
            display: flex;
            flex-flow: column;
            // margin-left: 10px;
            h2{
                font-size: 16px;
                font-weight: 700;
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
                overflow: hidden;
                max-width: 110px;
            }
            img{
                max-height: 20px;
                max-width: 20px;
            }
        }
    }
    .user-info-meta{
        display: flex;
        flex-flow: wrap;
        justify-content: space-between;
        padding: 0 20px;
        font-size: 12px;
        color: #bcbcbc;
        text-align: center;
        margin-top: 20px;
    }
    .user-info-desc{
        font-size: 12px;
        margin-top: 10px;
        display: -webkit-box;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 2;
        overflow: hidden;
        width: 100%;
        background: #f9f9f9;
        padding: 10px;
        color: #bcbcbc;
        height: 50px;
    }
}
</style>
