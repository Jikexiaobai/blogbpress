<template>
    <div class="list-info-box">
        <div class="list-info-header">
            <img :src="listObj.cover | resetImage(195,160)" >
            <div class="list-info-header-status">
                <a-tag v-if="listObj.status == 3" color="#f50">
                        未通过
                    </a-tag >
                    <a-tag v-if="listObj.status == 1" color="#2db7f5">
                        待审核
                    </a-tag>
                    <a-tag v-if="listObj.status == 2" color="#87d068">
                        已通过
                    </a-tag>
            </div>
        </div>
        <div class="list-info-center">
            <h2>{{listObj.title}}</h2>
        </div>
        <div class="list-info-meta">
            <!-- <a-tag color="#2db7f5">
                分类
            </a-tag> -->
            <div class="list-info-meta-date">
                <span class="icon"><a-icon type="clock-circle" /></span>
                <span class="date">{{listObj.createTime | resetData}}</span>
            </div>
        </div>
        <div class="list-info-box-ac">
            <a-space size="large">
                <a-button @click="view(listObj.id)" type="primary" shape="circle" icon="search" />
                <a-button @click="edit(listObj.id)" type="primary" shape="circle" icon="edit" />
                <a-button  @click="remove(listObj.id)"  type="danger" shape="circle" icon="close" />
            </a-space>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        listObj:{
            type: Object,
            default: {}
        }
    },
    methods:{
        view(id){
            this.$emit('view',id);
        },
        edit(id){
            this.$emit('edit',id);
        },
        remove(id){
            this.$emit('remove',id);
        },
    }
}
</script>

<style lang="less" scoped>
.list-info-box{
    background: #f7f7f7;
    padding: 10px;
    margin: 5px;
    position: relative;
    .list-info-header{
        height: 160px;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        position: relative;
        font-weight: 500;
        color: #fff;
        cursor: pointer;
        img{
            width: 100%;
            height: 100%;
            object-fit: cover;
        }
        .list-info-header-status{
            position: absolute;
            top: 10px;
            left: 10px;
            z-index: 1;
        }
    }
    .list-info-center{
        h2{
            font-size: 18px;
            font-weight: bold;
            color: #425167;
            height: 32px;
            line-height: 32px;
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
        }
        p{
            line-height: 24px;
            font-size: 14px;
            height: 48px;
            color: #657786;
            overflow: hidden;
            text-overflow: ellipsis;
            text-align: justify;
            display: -webkit-box;
            -webkit-line-clamp: 2;
            -webkit-box-orient: vertical;
        }
    }
    .list-info-meta{
        margin-top: 10px;
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
    .list-info-box-ac{
        position: absolute;
        top: 0;
        left: 0;
        background: rgba(0,0,0,0.3);
        width: 100%;
        height: 100%;
        display: none;
        z-index: 2;
    }
}
.list-info-box:hover{
    box-shadow: 0 0 6px #dedede;
    .list-info-box-ac{
        display: flex;
        justify-content: center;
        align-items: center;
    }
}
</style>