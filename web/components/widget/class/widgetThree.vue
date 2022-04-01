<template>
    <div class="widget-box">
        <div class="widget-title">
            <h2 class="title">
                最新课程
            </h2>
        </div>
        <a-row class="content" :gutter="[{md:20},{md:20}]">
            <a-col :span="6">
                <ClassItem />
            </a-col>
            <a-col :span="6">
                <ClassItem />
            </a-col>
            <a-col :span="6">
                <ClassItem />
            </a-col>
            <a-col :span="6">
                <ClassItem />
            </a-col>
        </a-row>
        <div class="more">
            <span>查看更多</span>
        </div>
    </div>
</template>

<script>
import ClassItem from "@/components/list/classItem"
export default {
    components:{
        ClassItem,
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
    .more{
        display: flex;
        justify-content: center;
        align-items: center;
        margin-top: 20px;
        span{
            background-color: white;
            padding: 5px 50px;
        }
    }
}
</style>