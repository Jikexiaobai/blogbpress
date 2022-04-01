<template>
    <div class="modal_box center opacity" :class="[isTrue && 'is_back_show']">
        <div class="modal_box_container">
            <div class="modal_box_title">
                <div class="modal_box_title_l">
                    <span>报名信息</span>
                </div>
                <div class="modal_box_title_r">
                    <a-icon  type="close" @click="cancel"/>
                </div>
            </div>

            <div class="modal_box_content">
                <div class="modal_box_content_desc">
                     <h2>联系方式</h2>
                    <a-radio-group 
                        v-model="mode"
                        buttonStyle="solid" 
                        >
                        <a-radio  :value="1">
                            微信
                        </a-radio>
                        <a-radio  :value="2">
                            QQ
                        </a-radio>
                        <a-radio  :value="3">
                            手机号
                        </a-radio>
                    </a-radio-group>
                </div>
                <div class="modal_box_content_desc">
                    <h2>联系号码</h2>
                    <a-input v-model="number"/>
                </div>
                <div class="modal_box_content_desc">
                    <a-input v-model="name"  placeholder="请输入称呼" />
                </div>
                

                <div class="modal_box_content_ok">
                        
                    <a-space size="middle">
                        <a-button @click="cancel">取消</a-button>
                        <a-button @click="ascertain" type="primary">
                        确定
                        </a-button>
                    </a-space>
                </div>
            </div>
        </div>
    </div>
</template>

<script>

import api from "@/api/index"

export default {

    data() {
        return {
            mode:undefined,
            name:undefined,
            number:undefined,
            eduId:undefined,
            // 登录输入框
            isTrue: false,
            state: null // 准备（prepare） 确定（ ascertain） 取消（cancel）
        };
    },
    methods: {
        async confirm(
            eduId,
        ) {
            console.log(eduId)
            this.eduId = eduId
            this.open();
            return new Promise((resolve, reject) => {
                const target = { state: 'prepare'};
                const that = this
                let res = new Proxy(target, {
                    set(event, key, value) {
                        if (value === 'ascertain') {
                            resolve(true);
                        } else {
                            reject(false);
                        }
                        return true
                    }
                });
                this.state = res;
            });
        },
        cancel(){
            this.state.state = "cancel"
            this.close()
        },
        async ascertain(){
    
            const formData = {
                mode:this.mode,
                name:this.name,
                number:this.number,
                eduId:this.eduId
            }

            const res = await this.$axios.post(api.postEduJoin,formData)
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
            this.mode = undefined
            this.name = undefined
            this.number = undefined
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
            width: 22rem;
            margin: 0 auto;
            position: relative;
            background-image: url("/img/login.png");
            background-repeat: no-repeat;
            background-size: 100%;
            margin-top: -9%;

            .modal_box_title{
                font-size: 13px;
                display: flex;
                justify-content: space-between;
                align-items: center;
                padding: 10px 20px;
                .modal_box_title_l{
                    display: block;
                    align-items: center;
                    width: 80%;
                    overflow: hidden;
                    white-space: nowrap;
                    text-overflow: ellipsis;
                }
            }

            .modal_box_content{
                padding: 20px 20px;
                .modal_box_content_desc{
                    margin: 10px 0;
                    h2{
                        font-size: 16px;
                        margin-bottom: 5px;
                    }
                }
                .modal_box_content_ok{
                    display: flex;
                    justify-content: flex-end;
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