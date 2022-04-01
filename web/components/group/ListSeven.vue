<template>
    <div class="list-group-item" @click="goInfo">
        <div class="list-group-item-img">
            <img :src="info.cover" :alt="info.title">
        </div>
        <div class="item-info">
            <h2>{{info.title}}</h2>
            <div class="item-meta">
                <span><b>{{info.joins | resetNum}}</b>个圈友</span>
                <i>•</i>
                <span><b>{{info.contents | resetNum}}</b>个话题</span>
            </div>
        </div>
        <div class="item-info-desc">
            <p>{{info.description}}</p>
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
    methods:{
        goInfo(){
            this.$router.push({ path: `/${this.info.module}/${this.info.id}`})
        }
    }
}
</script>


<style lang="less" scoped>
.list-group-item{
    position: relative;
    font-size: 13px;
    padding: 16px;
    border: 1px solid #f3f3f3;
    cursor: pointer;
    .list-group-item-img{
        position: absolute;
        width: 40px;
        height: 40px;
        overflow: hidden;
        box-shadow: 0 0 1px #888;
        img{
            width: 100%;
            height: 100%;
        }
    }
    .item-info{
        margin-left: 50px;
        h2{
            height: 17px;
            margin-right: 55px;
            font-size: 14px;
            display: -webkit-box;
            overflow: hidden;
            font-weight: 700;
        }
        .item-meta{
            height: 17px;
            margin-top: 4px;
            span{
                display: inline-block;
                font-size: 12px;
                height: 15px;
                color: #999;
                b{
                    font-weight: 400;
                }
            }
            i{
                margin: 0 5px;
                font-style: normal;
                color: #ccc;
            }
        }
    }
    .item-info-desc{
        margin-top: 10px;
        position: relative;
        background: #f9f9f9;
    }
    .item-info-desc::before{
        left: 17px;
        right: inherit;
        content: "";
        position: absolute;
        border-style: solid;
        border-color: transparent transparent #f9f9f9;
        border-width: 7px;
        top: -12px;
        z-index: 1;
    }
}
</style>