<template>
    <div class="info">
        <div class="user-info">
            <div class="user-info-l">
                <nuxt-link :to="{path:'/profile/' + info.userInfo.id }" class="info-link">   
                    <Avatar :src="info.userInfo.avatar+'@w60_h60'" :size="40"/>
                </nuxt-link>
                <div class="user-name-box">
                    <nuxt-link :to="{path:'/profile/' + info.userInfo.id }" class="info-link">   
                        <h2 class="user-name">{{info.userInfo.nickName}}</h2>
                    </nuxt-link>
                    <div class="user-role">
                        <a-space>
                            <img :src="info.userInfo.grade.icon" :alt="info.userInfo.grade.title">
                            <img v-if="info.userInfo.vip" :src="info.userInfo.vip.icon" :alt="info.userInfo.vip.title">
                        </a-space> 
                    </div>
                </div>
            </div>
            <div class="user-info-r">
                <span class="group-meta-date">
                    {{info.createTime | resetData}}
                </span>
            </div>
        </div> 
        <div class="content" @click="goInfo">
            <div class="cover">
                <img :src="info.cover" :alt="info.title">
                <a-icon  class="icon" type="play-circle" />
            </div>
           
            <div class="title-desc">
                <h2>{{info.title}}</h2>
                <p>{{info.description}}</p>
            </div>
        </div>
        <div class="mate">
            <a-space size="large">
                <a-space>
                    <a-icon type="eye" />
                    <span>{{info.views | resetNum}}</span>
                </a-space>
                <a-space>
                    <a-icon type="like" />
                    <span>{{info.likes | resetNum}}</span>
                </a-space>
            </a-space>
        </div>
        <div class="divider"></div>
    </div>
</template>

<script>
export default {
    props: {
        info:{
            type: Object,
            default: null
        }
    },
    methods:{
        goInfo(){
            console.log(this.info)
            this.$router.push({ path: `/${this.info.module}/${this.info.id}`})
        }
    }
}
</script>

<style lang="less" scoped>
.info{
    background-color: white;
    padding:10px;
    .user-info{
        display: flex;
        justify-content: space-between;
        .user-info-l{
            display: flex;
            align-items: center;
            margin-right: 10px;
            .user-name-box{
                .user-name{
                    font-size: 14px;
                }
                .user-role{
                    img{
                        max-height: 20px;
                        max-width: 20px;
                    }
                }
            }
        }
        .user-info-r{
            .group-meta-date{
                height: 20px;
                line-height: 20px;
                font-size: 12px;
                color: #b9b9b9;
            }
        }
    }
    .content{
        margin-top: 10px;
        cursor: pointer;
        .cover{
            height: 300px;
            width: 450px;
            border-radius: 5px;
            display: flex;
            align-items: center;
            justify-content: center;
            img{
                border-radius: 5px;
                display: block;
                max-width: 100%;
                max-height: 100%;
                background-color: #F5F6F7;
                width: 100%;
                height: 100%;
            }
            .icon{
                position: absolute;
                color: rgba(255, 255, 255, 0.4);
                font-size: 30px;
                transition: color .5s;
            }
            
        }
        .cover:hover{
            .icon{
                transition: color .5s;
                color: #fff;
            }
        }
        
        .title-desc{
            margin-top: 5px;
            h2{
                font-size: 18px;
                font-weight: bold;
                display: -webkit-box;
                -webkit-box-orient: vertical;
                -webkit-line-clamp: 1;
                overflow: hidden;
                text-justify: inter-ideograph;
                word-break: break-all;
            }
            p{
                font-size: 14px;
                font-weight: bold;
                display: -webkit-box;
                -webkit-box-orient: vertical;
                -webkit-line-clamp: 2;
                overflow: hidden;
                text-justify: inter-ideograph;
                word-break: break-all;
            }
        }
    }
    .mate{
        margin-top: 10px;
    }
    .divider{
        margin-top: 10px;
        height: 2px;
        background-color: #f8f8f8;
    }
}

</style>