<template>
    <div class="select-box">
        <a-form-model ref="selectForm" :model="selectForm" :rules="selectForm.rules" layout="inline">
            <a-form-model-item ref="orderNum" prop="orderNum">
                <a-input
                    size="large"
                    type="text"
                    placeholder="订单编号"
                    v-model="selectForm.orderNum"
                >
                </a-input>
            </a-form-model-item>
            <a-form-model-item ref="orderType" prop="orderType">
                <a-select v-model="selectForm.orderType" 
                    :style="{width:200 + 'px'}"
                    size="large"
                    placeholder="订单类型">
                    <a-select-option  :value="1">
                        充值
                    </a-select-option>
                    <a-select-option  :value="2">
                        打赏充电
                    </a-select-option>
                    <a-select-option  :value="3">
                        内容购买
                    </a-select-option>
                    <a-select-option  :value="4">
                        查看话题
                    </a-select-option>
                    <a-select-option  :value="5">
                        加入付费圈子
                    </a-select-option>
                    <a-select-option  :value="6">
                        购买付费课程
                    </a-select-option>
                    <a-select-option  :value="7">
                        查看付费答案
                    </a-select-option>
                    <a-select-option  :value="8">
                        开通vip
                    </a-select-option>
                    <a-select-option  :value="9">
                        认证付费
                    </a-select-option>
                </a-select>
            </a-form-model-item>
            <a-form-model-item ref="status" prop="status">
                <a-select v-model="selectForm.status" 
                    :style="{width:200 + 'px'}"
                    size="large"
                    placeholder="订单状态">
                    <a-select-option  :value="1">
                        未支付
                    </a-select-option>
                    <a-select-option  :value="2">
                        已支付
                    </a-select-option>
                </a-select>
            </a-form-model-item>
            <a-form-model-item>
                <a-button v-action="'/order/list'" @click="reset" style="margin-left: 10px;">
                    重置
                </a-button>
                <a-button v-action="'/order/list'" @click="select" type="primary" style="margin-left: 10px;">
                    查询
                </a-button>
            </a-form-model-item>
        </a-form-model>
    </div>
</template>

<script>
export default {
    data(){
        return{
            cateList:[],
            selectForm:{
                status:undefined,
                orderNum:undefined,
                orderType:undefined,
            }
        }
    },
    mounted(){
    },
    methods:{
        reset(){
            this.selectForm.status = undefined
            this.selectForm.orderNum = undefined
            this.selectForm.orderType = undefined
            this.$emit("select",this.selectForm)
        },
        select(){
            this.$emit("select",this.selectForm)
        }
    }
}
</script>

<style lang="less" scoped>
.select-box{
    margin-bottom: 10px;
}
</style>