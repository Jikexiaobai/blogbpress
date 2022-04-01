<template>
    <div class="widget-box">
        <div class="widget-title">
            <h2 class="title">
                专家讲师
            </h2>
        </div>
        <a-row class="content" :gutter="[{md:20},{md:20}]">
            <a-col :span="8">
                <UserItem />
            </a-col>
            <a-col :span="8">
                <UserItem />
            </a-col>
            <a-col :span="8">
                <UserItem />
            </a-col>
            <a-col :span="8">
                <UserItem />
            </a-col>
            <a-col :span="8">
                <UserItem />
            </a-col>
            <a-col :span="8">
                <UserItem />
            </a-col>
        </a-row>
    </div>
</template>

<script>
import UserItem from "@/components/list/userItem"
export default {
    components:{
        UserItem,
    },
}
</script>

<style lang="less" scoped>
.widget-box{
    margin-bottom: 30px;
    .widget-title{
        margin-bottom: 10px;
        h2{
            color: #3c3e43;
            font-size: 18px;
            font-weight: 700;
        }
    }
}
</style>