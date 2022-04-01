<template>
    <div class="modal_box center opacity" :class="[isTrue && 'is_back_show']">
        <div class="modal_box_container">
            <div class="modal_box_title">
                <div class="modal_box_title_r">
                    <a-icon  type="close" @click="cancel"/>
                </div>
            </div>
            <div class="modal_box_content">
                <div class="social">
                    <div>
                        <a-space :size="20">
                            <div class="social_item" @click="share('qq')">
                                <div class="icon" style="background-color: #4dafea;">
                                    <f-icon type="iconQQ" :size="25"></f-icon>
                                </div>
                                <h5>QQ</h5>
                            </div>
                            <div class="social_item" @click="share('qzone')">
                                <div class="icon" style="background-color: #eecf3d;">
                                    <f-icon type="iconqqkongjian" :size="25"></f-icon>
                                </div>
                                <h5>QQ空间</h5>
                            </div>
                            <div class="social_item" @click="share">
                                <div class="icon" style="background-color: #e6162d;">
                                    <f-icon type="iconweibo" :size="25"></f-icon>
                                </div>
                                <h5>微博</h5>
                            </div>
                        </a-space>
                    </div>
                    <div class="link">
                        <h6>链接地址</h6>
                        <div class="copy-link">
                            <div class="text">
                                {{link}}
                            </div>
                            <a-button 
                            v-clipboard:copy="link" 
                            v-clipboard:success="onCopy"
                            type="dashed" size="small">
                                复制
                            </a-button>
                        </div>
                    </div>
                </div>
                <div class="wechat">
                    <img :src="qrImg" alt="title">
                </div>
            </div>
            
        </div>
    </div>
</template>

<script>
import FIcon from '@/components/icon/FIcon'

import Qrious from 'qrious'

export default {
    components:{
        FIcon
    },
    data() {
        return {
            link:null,
            title:null,
            description:null,
            cover:null,
            qrImg:null,
            // 登录输入框
            isTrue: false,
            state: null // 准备（prepare） 确定（ ascertain） 取消（cancel）
        };
    },
    methods: {
        async confirm(
            link,
            title,
            description,
            cover
        ) {
            this.link = link || this.link
            this.title = title || this.title
            this.cover = cover || this.cover
            this.description = description || this.description
             var qr = new Qrious({
                    value: link,
                    size:100,
                    level:'L'
            });
            this.qrImg = qr.toDataURL('image/jpeg')

            this.open();
            return new Promise((resolve, reject) => {
                const target = { state: 'prepare'};
                const that = this
                let res = new Proxy(target, {
                    set(event, key, value) {
                        if (value === 'ascertain') {
                            resolve(that.activeList);
                            
                        } else {
                            reject(false);
                        }
                        return true
                    }
                });
                this.state = res;
            });
        },
        onCopy(e){
            this.$message.success(
                "复制成功",
                3
            )
            this.ascertain()
        },
        share(e){
            switch (e) {
                case "qq":
                    const qqurl = `http://connect.qq.com/widget/shareqq/index.html?url=${this.link}&sharesource=qzone&title=${this.title}&pics=${this.cover}&summary=${this.description}&desc=${this.description}`
                    window.open(qqurl,'newwindow','height=100,width=100,top=100,left=100');  
                    this.ascertain()
                    return;
                case "qzone":
                    const qzoneurl = `https://sns.qzone.qq.com/cgi-bin/qzshare/cgi_qzshare_onekey?url=${this.link}&sharesource=qzone&title=${this.title}&pics=${this.cover}&summary=${this.description}`
                    console.log(qzoneurl)
                    window.open(qzoneurl,'newwindow','height=100,width=100,top=100,left=100');  
                    this.ascertain()
                    return;
            }
        },
        cancel(){
            this.state.state = "cancel"
            this.close()
        },
        async ascertain(){
            this.state.state = "ascertain"
            this.close()
        },
        open() {
            this.isTrue = true;
        },
        close() {
            this.isTrue = false;
        },
        
    }
};
</script>

<style lang="less" scoped>
    .modal_box {
        user-select: none;
        pointer-events: none;
        z-index: 20;
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        width: 100%;
        height: 100%;
        opacity: 0;
        visibility: hidden;
        transform: perspective(1px) scale(1.1);
        transition: visibility 0s linear .15s,opacity .15s 0s,transform .15s;
        display: flex;
        align-items: center;
        justify-content: center;
        .modal_box_container{
            background-color: #fff;
            width: 25rem;
            margin: 0 auto;
            position: relative;
            background-image: url("/img/login.png");
            background-repeat: no-repeat;
            background-size: 100%;
            margin-top: -9%;
            .modal_box_title{
                font-size: 15px;
                display: flex;
                justify-content: flex-end;
                padding: 10px 20px;
            }
            .modal_box_content{
                display: flex;
                align-items: center;
                padding: 0px 20px;
                .social{
                    flex: 1;
                    margin-right: 20px;
                    .social_item{
                        cursor: pointer;
                        display: flex;
                        flex-direction: column;
                        align-items: center;
                        .icon{
                            width: 40px;
                            height: 40px;
                            padding: 10px;
                            border-radius: 80%;
                            display: flex;
                            justify-content: center;
                            align-items: center;
                        }
                        h5{
                            margin-top: 5px;
                        }
                    }
                    .link{
                        margin: 20px 0;
                        
                        .copy-link{
                            margin-top: 10px;
                            display: flex;
                            align-items: center;
                            .text{
                                width: 150px;
                                margin-right: 10px;
                                padding: 5px;
                                font-size: 12px;
                                text-overflow: ellipsis;
                                overflow: hidden;
                                white-space: nowrap;
                                background-color: #f8f8f8;
                            }
                        }
                    }
                }
                .wechat{
                    margin-bottom: 20px;
                    width: 100px;
                    height: 100px;
                    img{
                        width: 100%;
                        height: 100%;
                    }
                }
            }
        }
        
    }
    .is_back_show {
        opacity: 1 !important;
        background: rgba(42, 44, 48, 0.7);
        pointer-events: auto !important;
        visibility: visible;
        transform: perspective(1px) scale(1);
        transition: visibility 0s linear 0s,opacity .15s 0s,transform .15s;
    }

    @media only screen and (max-width: 768px) {
        .modal_box{
            .modal_box_container{
                width: 100%;
            }
        }
    }
</style>