<template>
    <div class="widget-box">
        <div class="warper" :style="{ width: '1500px' }">
            <div class="widget-title" v-if="info.showTitle == 2">
                <h2 class="title">{{info.title}}</h2>
            </div>
            <div class="widget-body" >
                <a-row :gutter="[{md:20,xs:10},{xs:10,md:20}]">
                    <a-col v-for="(item,index) in list" :key="index" :md="4" :sm="4" :xs="8">
                        <div class="item">
                            <div class="item-bg" :style="{ background: `${item.bgColor}` }">
                                <div class="item-info" :style="{ color: `${item.color}` }">
                                    <h2 :style="{ textShadow: `5px 5px 5px ${item.textColor}`}">{{item.title}}</h2>
                                    <p>{{item.desc}}</p>
                                </div>
                                <div class="link">
                                    <a v-for="(aitem,aindex) in item.links" :key="aindex" :href="aitem.link">
                                        {{aitem.title}}
                                    </a>
                                </div>
                            </div>
                        </div>
                    </a-col>
                </a-row>
            </div>
            
            <div class="widget-empty" v-if="info.list == null">
                <a-config-provider :locale="locale">
                    <a-empty />
                </a-config-provider>
            </div>
        </div>
    </div>
</template>

<script>
import { mapState } from "vuex"
import zhCN from 'ant-design-vue/lib/locale-provider/zh_CN'
import Avatar from "@/components/avatar/avatar"
export default {
    props: {
        info:{
            type: Object,
            default: {}
        },
    },
    components:{
        Avatar
    },
    data(){
        return{
             locale: zhCN,
            list:[]
        }
    },
    computed:{
        ...mapState(["design"]),
    },
    created(){
        
        if (this.info.list.length>0) {
            const tmplist= JSON.parse(this.info.list)
            const list = tmplist.map((item)=>{
                let tmpStr = item.text.split("|");
                const tmp = {
                    title:tmpStr[0],
                    desc:tmpStr[1],
                    bgColor:tmpStr[2],
                    color:tmpStr[3],
                    textColor:tmpStr[4],
                    links:tmpStr[5],
                    icon:tmpStr[6],
                }
                let linkStr = tmpStr[5].split(",");
                const tmpLinkList = linkStr.map((litem)=>{
                    let tmplinkStr = litem.split("-");
                    const linkInfo = {
                        title:tmplinkStr[0],
                        link:tmplinkStr[1],
                    }
                    return linkInfo
                })
                tmp.links = tmpLinkList
                return tmp
            })
            this.list = list
        }
       
    },
    methods: {
       
    }, 
}
</script>

<style lang="less" scoped>
.widget-box{
    margin: 10px 0;
    display: flex;
    justify-content: center;
    align-items: center; 
    .widget-title{
        display: flex;
        justify-content: space-between;
        align-items: center;
        .title{
            font-size: 14px;
            letter-spacing: .8px;
            font-weight: 700;
            position: relative;
        }
        .title::after{
            background-color: #6c757d;
            bottom: -10px;
            content: "";
            height: 3px;
            left: 0;
            position: absolute;
            width: 20px;
        }
    }
    .widget-body{
        margin-top: 20px;
        .item{
            height: 100px;
            position: relative;
            z-index: 3;
            .item-bg{
                position: relative;
                height: 100px;
                font-size: 14px;
                border-radius: 10px;
                transition: 0.2s;
                overflow: hidden;
                .item-info{
                    display: block;
                    position: relative;
                    box-sizing: border-box;
                    padding: 22px 0 0 16px;
                    border-radius: 5px;
                    overflow: hidden;
                    height: 100px;
                    h2{
                        color: #fff;
                        font-weight: 900;
                    }
                    p{
                        position: relative;
                        z-index: 1;
                        text-shadow: 0 0 5px rgb(255 255 255 / 80%);
                        background: linear-gradient(to right, #ffffffa6, #fefefd00);
                        margin-left: -18px;
                        padding-left: 18px;
                    }
                }
                .link{
                    text-align: center;
                    line-height: 26px;
                    font-size: 14px;
                    a{
                        margin: 10px 5px 0;
                        display: inline-block;
                        background: rgba(255,255,255,0.50);
                        width: 77px;
                        height: 26px;
                        border-radius: 13px;
                        font-size: 13px;
                        color: #842100;
                    }
                }
            }
            .item-bg:hover{
                height: 180px;
                box-shadow: 0px 2px 5px rgb(228, 228, 228);
            }
            .item-bg::before{
                background-position: 0 -100px;
                content: '';
                position: absolute;
                right: 0;
                top: 0;
                width: 168px;
                height: 100px;
            }
        }
    }
    
}
@media screen and (max-width: 768px){
    .widget-box{
        margin: 0 10px;
    }
}
</style>