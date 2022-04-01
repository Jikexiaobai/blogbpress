<template>
    <a-layout>
        <!-- 头部 -->
        <a-layout-header>
            <div class="menu">
                <div class="menu-buttom">
                    <div class="menu-buttom-box" :style="{ width: this.design.width + 'px' }">
                        <div class="menu-buttom-left">
                            <div @click="goHome" class="menu-buttom-left-logo">
                                <img :src="`${base.logo}`" :alt="`${base.title}`" >
                            </div>
                            <span>
                                用户创作中心    
                            </span>
                        </div>
                        <div class="menu-buttom-right">
                            <div v-if="token != null" class="menu-main-right-user">
                                <div class="messages" @click="goMessage">
                                    <a-badge   dot>
                                        <a-icon class="meta-icon meta-margin-left"  type="bell" />
                                    </a-badge>
                                </div>
                                <a-dropdown :trigger="['click']" placement="bottomCenter" >
                                    <a-avatar :src="userInfo.avatar" class="meta-avatar meta-margin-left" :size="30" icon="user" />
                                    <a-menu  slot="overlay">
                                        <a-menu-item @click="goProfile" key="0">
                                            <a-icon type="user" />
                                            我的主页
                                        </a-menu-item>
                                        <a-menu-item key="1">
                                            <a-icon type="wallet" />
                                            我的钱包
                                        </a-menu-item>
                                        <a-menu-item key="2">
                                            <a-icon type="barcode" />
                                            订单中心
                                        </a-menu-item>
                                        <a-menu-item key="3">
                                            <a-icon type="laptop" />
                                            我的课程
                                        </a-menu-item>
                                        <a-menu-item key="4" @click="goSetting">
                                            <a-icon type="setting" />
                                            <span>设置中心</span>
                                        </a-menu-item>
                                        <a-menu-divider />
                                        <a-menu-item key="5" >
                                            <a-icon type="logout" />
                                            退出
                                        </a-menu-item>
                                    </a-menu>
                                </a-dropdown>
                            </div>
                            <div v-if="token == null" class="menus-meta">
                                <a-button @click="toLogin" type="dashed" class="mmenus-meta-login">登录</a-button>
                                <a-button @click="toRegister" type="primary" class="mmenus-meta-login">注册</a-button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </a-layout-header>
        <!-- 主体内容 -->
        <a-layout-content>
            <transition name="page-transition">
                <Nuxt />
            </transition>
        </a-layout-content>
    </a-layout>
</template>

<script>
import { mapState } from "vuex"
// import MPlayer from "@/components/player/audioPlayer"
export default {
    computed:{
        ...mapState(["design","base"]),
        ...mapState("user",["token","userInfo"])
    },
    created(){
      // console.log("dasdasd")
    },
    methods:{
        // ...mapActions("auth",["LoginOut"]),
        
        toLogin(){
            this.$Auth("login","登录","快速登录")
        },
        logout(){
            this.LoginOut()
        },
        toRegister(){
            this.$Auth("register","用户注册","立即注册")
        },
        goSetting(){
            this.$router.push("/account/base")
        },
        goProfile(){
            this.$router.push(`/profile/${this.userInfo.userId}`)
        },
        goMessage(){
            this.$router.push("/messages/system")
        },
        goHome(){
            this.$router.push({
                path:"/"
            })
        },
    }
    
}
</script>


<style lang="less" scoped>
.menu{
    z-index: 900;
    position: fixed;
    width: 100%;
    background: white;
    box-shadow: 0 4px 8px 0 rgba(7, 17, 27, 0.1);
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    .menu-buttom{
        display: flex;
        justify-content: center;
        align-items: center;
        height: 64px;
        .menu-buttom-box{
            display: flex;
            justify-content: space-between;
            align-items: center;
            height: 64px;
            .menu-buttom-left{
                display: flex;
                align-items: center;
                height: 64px;
                .menu-buttom-left-logo{
                    cursor: pointer;
                    padding: 12px 0;
                    img{
                        display: block;
                        height: 34px;
                    }
                }
                span{
                    color: #1d7dfa;
                    font-weight: 700;
                    font-size: 20px;
                    font-family: PingFang SC;
                    line-height: 28px;
                    display: flex;
                    align-items: center;
                    position: relative;
                    margin-left: 16px;
                }
            }
            .menu-buttom-right{
                .menu-main-right-user{
                    .messages{
                        cursor: pointer;
                    }
                    display: flex;
                    align-items: center;
                    width: 180px;
                    justify-content: flex-end;
                    .meta-icon{
                        // margin-right: 10px;
                        font-size: 18px;
                    }
                    .meta-margin-left{
                        margin-left: 20px;
                    }
                    .meta-avatar{
                        cursor: pointer;
                    }
                }
                .menus-meta{
                    display: flex;
                    align-items: center;
                    .mmenus-meta-login{
                        margin-left: 10px;
                        line-height: 0.2;
                    }
                }
            }
        }
    }
}


.ant-layout-content{
  min-height: 100vh;
}

.ant-layout-header{
  padding: 0;
  line-height: 0px;
}

</style>