<template>
    <div class="pay-center">
        <h2>我的钱包</h2>
        <div class="setting-container">
            <div class="wallet-top">
                <a-row :gutter="[{md:12}]">
                    <a-col :span="8">
                        <div class="wallet-bl">
                            <div class="margbtm">可用余额</div>
                            <div class="amount">
                                <span>{{base.currencySymbol}}</span>
                                <span>{{balance.toFixed(2)}}</span>
                                <a-button  @click="cash" type="link">
                                    提现
                                </a-button>
                                <a-button @click="recharge" type="link">
                                    充值
                                </a-button>
                            </div>
                        </div>
                    </a-col>
                    <!-- <a-col :span="8">
                        <div class="wallet-d">
                            <div class="margbtm">冻结金额</div>
                            <div class="amount">
                                <span>￥</span>
                                <span>124.00</span>
                            </div>
                        </div>
                    </a-col>
                    <a-col :span="8">
                       <div class="wallet-pass">
                            <div class="margbtm">支付密码</div>
                            <div class="amount">
                                <span>
                                    设置/修改
                                </span>
                            </div>
                        </div>
                    </a-col> -->
                </a-row>
            </div>
            <ul class="wallet-menu">
                <li>
                    <button 
                        :class="type == 0 ? 'active': ''" 
                        @click="changeType(0)">
                        提现记录
                    </button>
                </li>
                <li>
                    <button 
                        :class="type == 1 ? 'active': ''" 
                        @click="changeType(1)">
                        充值记录
                    </button>
                </li>
                <li>
                    <button 
                        :class="type == 2 ? 'active': ''" 
                        @click="changeType(2)">
                        订单明细
                    </button>
                </li>
            </ul>
            <div v-if="type == 0" class="wallet-tb">
                <div class="wallet-select">
                    <a-select placeholder="筛选提现状态"  style="width: 220px" @change="changeOrderStatus">
                        <a-select-option :value="0">
                            全部
                        </a-select-option>
                        <a-select-option :value="2">
                            已打款
                        </a-select-option>
                        <a-select-option :value="1">
                            待审核
                        </a-select-option>
                        <a-select-option :value="3">
                            提现失败
                        </a-select-option>
                    </a-select>
                </div>
                <div class="wallet-table">
                    <a-table 
                    :pagination="{
                        pageSize: queryParam.limit,
                        total:total,
                    }"
                    @change="changePage"    
                    :columns="cashColumn" 
                    size="middle"  
                    :data-source="list" 
                    >
                        
                        <span slot="createTime" slot-scope="createTime">{{ createTime | resetData }}</span>
                        <span slot="status" slot-scope="status">
                            <a-tag v-if="status ==1" color="#f50">
                                待审核
                            </a-tag>
                            <a-tag v-if="status ==2" color="#87d068">
                                已打款
                            </a-tag>
                            <a-tag v-if="status == 3" color="#87d068">
                                提现失败
                            </a-tag>
                        </span>
                        <span slot="payMethod" slot-scope="payMethod">
                            <a-tag v-if="payMethod ==1" color="#2db7f5">
                                支付宝
                            </a-tag>
                            <a-tag v-if="payMethod ==2" color="#87d068">
                                微信
                            </a-tag>
                        </span>
                        <span slot="serviceMoney" slot-scope="serviceMoney">
                            {{base.currencySymbol}} {{serviceMoney}}
                        </span>
                        <span slot="money" slot-scope="money">
                            {{base.currencySymbol}} {{money}}
                        </span>
                        <span slot="cashMoney" slot-scope="cashMoney">
                            {{base.currencySymbol}} {{cashMoney}}
                        </span>
                        <span slot="action" slot-scope="text, record">
                            <a @click="viewRemark(record.remark)" v-if="record.status == 3">
                                查看原因 
                            </a>
                        </span>
                    </a-table>
                </div>
            </div>
            <div v-if="type == 1" class="wallet-recharge">
                <div class="wallet-select">
                    <a-select placeholder="筛选充值状态"  style="width: 220px" @change="changeOrderStatus">
                        <a-select-option :value="0">
                            全部
                        </a-select-option>
                        <a-select-option :value="1">
                            待审核
                        </a-select-option>
                        <a-select-option :value="2">
                            已充值
                        </a-select-option>
                        <a-select-option :value="3">
                            充值失败
                        </a-select-option>
                    </a-select>
                </div>
                <div class="wallet-table">
                    <a-table 
                    :pagination="{
                        pageSize: queryParam.limit,
                        total:total,
                    }"
                    @change="changePage"    
                    :columns="rechargeColumn" 
                    size="middle"  
                    :data-source="list" 
                    >
                        <span slot="createTime" slot-scope="createTime">{{ createTime | resetData }}</span>
                        <span slot="status" slot-scope="status">
                            <a-tag v-if="status == 1" color="#f50">
                                待审核
                            </a-tag>
                            <a-tag v-if="status == 2" color="#2db7f5">
                                已充值
                            </a-tag>
                            <a-tag v-if="status == 3" color="#87d068">
                                充值失败
                            </a-tag>
                        </span>
                        <span slot="mode" slot-scope="mode">
                            <a-tag v-if="mode == 1" color="#f50">
                                支付宝
                            </a-tag>
                            <a-tag v-if="mode == 2" color="#2db7f5">
                                微信
                            </a-tag>
                            <a-tag v-if="mode == 3" color="#87d068">
                                卡密
                            </a-tag>
                            <a-tag v-if="mode == 4" color="#108ee9">
                                人工转账
                            </a-tag>
                        </span>
                        <span slot="action" slot-scope="text, record">
                            <a @click="viewRemark(record.remark)" v-if="record.status == 3">
                                查看原因 
                            </a>
                        </span>
                    </a-table>
                </div>
            </div>
            <div v-if="type == 2" class="wallet-order">
                <div class="wallet-select">
                    <a-select placeholder="筛选订单状态"  style="width: 220px" @change="changeOrderStatus">
                        <a-select-option :value="0">
                            全部
                        </a-select-option>
                        <a-select-option :value="2">
                            已支付
                        </a-select-option>
                        <a-select-option :value="1">
                            未支付
                        </a-select-option>
                    </a-select>
                    <a-select placeholder="筛选订单类型"  style="width: 220px" @change="changeOrderType">
                        <a-select-option :value="0">
                            全部
                        </a-select-option>
                        <a-select-option :value="ORDERTYPE.CD">
                            用户充电打赏
                        </a-select-option>
                        <a-select-option :value="ORDERTYPE.BUYZY">
                            购买内容
                        </a-select-option>
                        <a-select-option :value="ORDERTYPE.VIEWANSWER">
                            查看付费答案
                        </a-select-option>
                        <a-select-option :value="ORDERTYPE.JOINCOURSE">
                            加入付费课程
                        </a-select-option>
                        <a-select-option :value="ORDERTYPE.JOINGROUP">
                            加入付费圈子
                        </a-select-option>
                    </a-select>
                </div>
                <div class="wallet-table">
                    <a-table 
                    :pagination="{
                        pageSize: queryParam.limit,
                        total:total,
                    }"
                    @change="changePage"    
                    :columns="orderColumn" 
                    size="middle"  
                    :data-source="list" 
                    >
                        <span slot="orderType" slot-scope="orderType">{{ orderType | orderTypeRestTitle }}</span>
                        <span slot="createTime" slot-scope="createTime">{{ createTime | resetData }}</span>
                        <span slot="status" slot-scope="status">
                            <a-tag v-if="status ==1" color="#f50">
                                未支付
                            </a-tag>
                            <a-tag v-if="status ==2" color="#87d068">
                                已支付
                            </a-tag>
                        </span>
                        <div slot="money" slot-scope="text,record">
                            <span v-if="record.isIncome|| record.orderType == 1"><span>+</span>{{record.money}}</span>
                            <span v-else><span>-</span>{{record.money}}</span>
                        </div>
                        <div slot="description" slot-scope="text,record">
                            <span >{{record.description}}</span>
                        </div>
                    </a-table>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import api from "@/api/index"
import {ORDERTYPE} from "@/shared/order"
import { mapState } from "vuex"

export default {
    middleware: 'auth',
    data(){
        return{
            ORDERTYPE,
            orderColumn:[
                {
                    title: '类型',
                    dataIndex: 'orderType',
                    key: 'orderType',
                    scopedSlots: { customRender: 'orderType' },
                },
                {
                    title: '金额',
                    dataIndex: 'money',
                    key: 'money',
                    scopedSlots: { customRender: 'money' },
                },
                {
                    title: '状态',
                    key: 'status',
                    dataIndex: 'status',
                    scopedSlots: { customRender: 'status' },
                },
                {
                    title: '时间',
                    key: 'createTime',
                    dataIndex: 'createTime',
                    scopedSlots: { customRender: 'createTime' },
                },
                {
                    title: '描述',
                    key: 'description',
                    dataIndex: 'description',
                    scopedSlots: { customRender: 'description' },
                },
                {
                    title: '订单号',
                    key: 'orderNum',
                    dataIndex: 'orderNum',
                    scopedSlots: { customRender: 'orderNum' },
                },
            ],
            cashColumn:[
                {
                    title: '提现单号',
                    dataIndex: 'code',
                    key: 'code',
                    scopedSlots: { customRender: 'code' },
                },
                {
                    title: '可到账金额',
                    dataIndex: 'cashMoney',
                    key: 'cashMoney',
                    scopedSlots: { customRender: 'cashMoney' },
                },
                {
                    title: '提现金额',
                    key: 'money',
                    dataIndex: 'money',
                    scopedSlots: { customRender: 'money' },
                },
                {
                    title: '服务费',
                    key: 'serviceMoney',
                    dataIndex: 'serviceMoney',
                    scopedSlots: { customRender: 'serviceMoney' },
                },
                {
                    title: '收款方式',
                    key: 'payMethod',
                    dataIndex: 'payMethod',
                    scopedSlots: { customRender: 'payMethod' },
                },
                {
                    title: '状态',
                    key: 'status',
                    dataIndex: 'status',
                    scopedSlots: { customRender: 'status' },
                },
                {
                    title: '申请时间',
                    key: 'createTime',
                    dataIndex: 'createTime',
                    scopedSlots: { customRender: 'createTime' },
                },
                {
                    title: "查看原因",
                    dataIndex: 'action',
                    scopedSlots: { customRender: 'action' }
                }
            ],
            rechargeColumn:[
                {
                    title: '充值单号',
                    dataIndex: 'code',
                    key: 'code',
                    scopedSlots: { customRender: 'code' },
                },
                {
                    title: '充值方式',
                    dataIndex: 'mode',
                    key: 'mode',
                    scopedSlots: { customRender: 'mode' },
                },
                {
                    title: '充值金额',
                    dataIndex: 'money',
                    key: 'money',
                    scopedSlots: { customRender: 'money' },
                },
                {
                    title: '状态',
                    key: 'status',
                    dataIndex: 'status',
                    scopedSlots: { customRender: 'status' },
                },
                {
                    title: '申请时间',
                    key: 'createTime',
                    dataIndex: 'createTime',
                    scopedSlots: { customRender: 'createTime' },
                },
                {
                    title: "查看原因",
                    dataIndex: 'action',
                    scopedSlots: { customRender: 'action' }
                }
            ],
            type:0,
            balance:0,
            total: 0,
            list: [],
            queryParam: {
                page: 1,
                limit: 10,
                orderType:0,
                status:0,
            },
        }
    },
    head(){
        return this.$seo(`用户中心-${this.base.title}`,`用户中心`,[{
            hid:"fiber",
            name:"description",
            content:`用户中心`
        }])
    },
    computed:{
        ...mapState("user",["userInfo"]),
        ...mapState(["base"])
    },
    mounted(){
        this.getBlance()
        this.getData()
    },
    filters: {
        orderTypeRestTitle(value) {
            switch (value) {
                case ORDERTYPE.CD:
                    return "打赏用户"
                case ORDERTYPE.BUYZY:
                    return "购买内容"
                case ORDERTYPE.JOINGROUP:
                    return "加入圈子"
                case ORDERTYPE.JOINCOURSE:
                    return "报名课程"
                case ORDERTYPE.VIEWANSWER:
                    return "查看付费答案"
                case ORDERTYPE.OPENVIP:
                    return "开通付费会员"
                case ORDERTYPE.VERIFY:
                    return "认证服务费"
            }
        }
    },
    methods:{
        async getBlance(){
            const res = await this.$axios.get(api.getAccountBalance)
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
            }
            this.balance = res.data.balance
        },
        async getData(){
            if (this.type == 2) {
                const res = await this.$axios.get(api.getOrderList,{params: this.queryParam})
                if (res.code != 1) {
                    this.$message.error(
                        res.message,
                        3
                    )
                }
                if (res.data.list != null) {
                    res.data.list = res.data.list.map((item)=>{
                        const tmp = {
                            key: item.orderNum,
                            orderType: item.orderType,
                            money: item.money.toFixed(2),
                            createTime: item.createTime,
                            status:item.status,
                            description: item.title,
                            isIncome:item.isIncome,
                            orderNum:item.orderNum
                        }
                        return tmp
                    })
                }

                this.list = res.data.list || []
                this.total = res.data.total || 0

                return
            }
            
            if (this.type == 0) {
                const res = await this.$axios.get(api.getCashList,{params: this.queryParam})
                if (res.code != 1) {
                    this.$message.error(
                        res.message,
                        3
                    )
                }
                if (res.data.list != null) {
                    res.data.list = res.data.list.map((item)=>{
                        const tmp = {
                            key: item.code,
                            cashMoney: item.cashMoney.toFixed(2),
                            money: item.money.toFixed(2),
                            serviceMoney: item.serviceMoney.toFixed(2),
                            payMethod:item.payMethod,
                            status: item.status,
                            createTime:item.createTime,
                            code:item.code,
                            remark:item.remark
                        }
                        return tmp
                    })
                }

                this.list = res.data.list || []
                this.total = res.data.total || 0

                return
            }
           
            if (this.type == 1) {
                const res = await this.$axios.get(api.getRechargeList,{params: this.queryParam})
                if (res.code != 1) {
                    this.$message.error(
                        res.message,
                        3
                    )
                }
                if (res.data.list != null) {
                    res.data.list = res.data.list.map((item)=>{
                        const tmp = {
                            key: item.code,
                            mode: item.mode,
                            money: item.money.toFixed(2),
                            status: item.status,
                            createTime:item.createTime,
                            code:item.code,
                            remark:item.remark
                        }
                        return tmp
                    })
                }

                this.list = res.data.list || []
                this.total = res.data.total || 0

                return
            }
       
        },
        changePage(e){
            this.queryParam.limit = e.pageSize
            this.queryParam.page = e.current
            this.getData()
        },
        recharge(){
            this.$Recharge().then((res)=>{
                if (res) {
                    this.getBlance()
                    if (this.type == 2) {
                        this.getData()
                    }
                }
            }).catch((err)=>{
                this.getData()
            })
        },
        cash(){
            this.$Cash(this.balance).then((res)=>{
                if (res) {
                    this.$message.success(
                        "提现申请提交成功",
                        3
                    )
                    this.getData()
                    return
                }
            }).catch((err)=>{
                console.log(err)
            })
        },
        changeOrderType(e){
            this.queryParam.orderType = e
            this.getData()
        },
        changeOrderStatus(e){
            this.queryParam.status = e
            this.getData()
        },
        changeType(e){
            this.type = e
            this.getData()
        },
        viewRemark(e){
            this.$message.error(
                e,
                10
            )
        }
    }
}
</script>

<style lang="less" scoped>
.pay-center{
    background-color: #fff;
    padding: 20px;
    h2{
        color: #bcbcbc;
        font-size: 18px;
    }
    .setting-container{
        .wallet-top{
            margin-top: 10px;
            .margbtm{
                font-size: 12px;
                margin-bottom: 5px;
            }
            .amount{
                display: flex;
                align-items: center;
                span{
                    font-size: 18px;
                    font-weight: 700;
                }
            }
            .wallet-d{
                display: flex;
                flex-direction: column;
                justify-content: center;
                align-items: center;
            }
            .wallet-pass{
                display: flex;
                justify-content: flex-end;
                align-items: flex-end;
                flex-direction: column;
            }
        }
        .wallet-menu{
            display: flex;
            flex-wrap: wrap;
            flex: 1;
            margin-top: 20px;
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
        .wallet-tb{
            margin-top: 10px;  
        }
        .wallet-select{
            margin: 10px 0;
        }
    }
}
</style>
