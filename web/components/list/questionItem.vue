<template>
    <div class="ask-list-info">
        <!-- <div class="ask-rank">
            <div class="answere-count">
                {{info.answers | resetNum}} <small>回答</small>  
            </div>
            <div class="ask-view">
                {{info.views | resetNum}} <small>查看</small>  
            </div>
        </div> -->
        <div class="ask-info">
            <div class="ask-author-date">
                <nuxt-link :to="{path:'/profile/' + info.userInfo.id }" class="info-link">   
                    <h2 class="user-name">{{info.userInfo.nickName}}</h2>
                </nuxt-link>
                <span class="date">{{info.createTime | resetData}}</span>
            </div>
            <nuxt-link :to="{path:'/question/' + info.id }" class="info-link">   
                <h2 class="ask-title">{{info.title}}</h2>
            </nuxt-link>
        </div>
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
}
</script>

<style lang="less" scoped>
.ask-list-info{
    padding: .75rem 0;
    border-bottom: 1px solid rgba(0,0,0,.0625);
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    .ask-rank{
        display: flex;
        .answere-count{
            margin-right: 10px;
            text-align: center;
            line-height: 1.2;
            display: inline-block;
            padding-top: 3px;
            width: 40px;
            height: 40px;
            -moz-border-radius: 1px;
            -webkit-border-radius: 1px;
            border-radius: 3px;
            background: #0084ff;
            border: none;
            color: #fff;
            small{
                display: block;
                font-size: 12px;
            }
        }
        .ask-view{
            margin-right: 10px;
            text-align: center;
            line-height: 1.2;
            display: inline-block;
            padding-top: 3px;
            width: 40px;
            height: 40px;
            -moz-border-radius: 1px;
            -webkit-border-radius: 1px;
            border-radius: 3px;
            background-color: #f3f3f3;
            border-color: #f3f3f3;
            color: #7b7b7b;
            small{
                display: block;
                font-size: 12px;
            }
        }
    }
    .ask-info{
        flex: 1;
        .ask-author-date{
            display: flex;
            align-items: center;
            
            .user-name{
                font-size: 14px;
                color: #7b7b7b;
            }
            .date{
                margin-left: 10px;
                font-size: 12px;
                color: #7b7b7b;
            }
        }
        .ask-title{
            font-size: 14px;
        }
    }
}
</style>