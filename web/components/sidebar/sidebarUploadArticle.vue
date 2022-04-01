<template>
  <div class="sidear-box">
        <div class="info">
            <div class="btn">
                <h2>专栏投递</h2>  
                <a-button @click="goArticle" type="primary">
                    立即投递
                </a-button>
            </div>
            <div class="desc">
                欢迎投递行业相关资讯线索。
            </div>
        </div>
    </div>
</template>

<script>
export default {
    methods:{
        goArticle(){
           this.$router.push({ path: `/member/article/list`})
        }
    }
}
</script>

<style lang="less" scoped>
.sidear-box{
    background: white;
    margin-bottom: 10px;
    .info{
        padding: 20px;
        .btn{
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 10px;
            h2{
                font-weight: 700;
                font-size: 16px;
                color: #1d2129;
            }
        }
        .desc{
            font-size: 14px;
            line-height: 22px;
            color: #86909c;
        }
    }
}
</style>