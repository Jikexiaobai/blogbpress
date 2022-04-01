<template>
    <div class="modal_box center opacity" :class="[isTrue && 'is_back_show']">
        <div class="modal_box_container">
            <div class="modal_box_title">
                <div class="modal_box_title_l">
                    <span>举报</span>
                </div>
                <div class="modal_box_title_r">
                    <a-icon  type="close" @click="cancel"/>
                </div>
            </div>

            <div class="modal_box_content">
                <div class="modal_box_content_radio">
                    <a-radio-group buttonStyle="solid" v-model="type" >
                        <a-radio  :value="REPORTTYPE.GG">
                            广告垃圾
                        </a-radio>
                        <a-radio  :value="REPORTTYPE.WG">
                            违规内容
                        </a-radio>
                        <a-radio  :value="REPORTTYPE.GS">
                            恶意灌水
                        </a-radio>
                        <a-radio  :value="REPORTTYPE.CF">
                            内容重复
                        </a-radio>
                    </a-radio-group>
                </div>
                <div class="modal_box_content_desc">
                    <h2>举报描述</h2>
                    <a-textarea v-model="description" placeholder="Basic usage" :rows="4" />
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
const REPORTTYPE = {
    GG:1, //1广告垃圾
    WG:2, //2违规内容
    GS:3, //3恶意灌水
    CF:4, //4内容重复
}
export default {

    data() {
        return {
            REPORTTYPE,
            type:null,
            description:null,
            module:null,
            relatedId:null,
            // 登录输入框
            isTrue: false,
            state: null // 准备（prepare） 确定（ ascertain） 取消（cancel）
        };
    },
    methods: {
        async confirm(
            relatedId,
            module
        ) {
            this.relatedId = relatedId
            this.module = module
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
        cancel(){
            this.state.state = "cancel"
            this.close()
        },
        async ascertain(){
    
            const formData = {
                description:this.description,
                relatedId:this.relatedId,
                module:this.module,
                type:this.type
            }
            const res = await this.$axios.post(api.postReportCreate,formData)
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
            this.$message.success(
                res.message,
                3
            )
            this.type = null
            this.description = null
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
            width: 33rem;
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
                .modal_box_content_radio{
                }
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