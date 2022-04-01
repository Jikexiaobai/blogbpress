<template>
  <div class="sidear-box">
        <h2>热门文章</h2>
        <div class="news-flash-list">
            <a-timeline>
                <a-timeline-item>
                    <p class="flash-title">快讯标题快讯标题快讯标题快讯标题快讯标题快讯标题快讯标题快讯标题</p>
                    <div class="flash-cate-date">
                        <span>#快讯标签#</span>
                        <span>日期</span>
                    </div>
                </a-timeline-item>
                <a-timeline-item>
                    <p class="flash-title">快讯标题快讯标题快讯标题快讯标题快讯标题快讯标题快讯标题快讯标题</p>
                    <div class="flash-cate-date">
                        <span>#快讯标签#</span>
                        <span>日期</span>
                    </div>
                </a-timeline-item>
                <a-timeline-item>
                    <p class="flash-title">快讯标题快讯标题快讯标题快讯标题快讯标题快讯标题快讯标题快讯标题</p>
                    <div class="flash-cate-date">
                        <span>#快讯标签#</span>
                        <span>日期</span>
                    </div>
                </a-timeline-item>
            </a-timeline>
        </div>
    </div>
</template>

<script>
export default {

}
</script>

<style lang="less" scoped>
.sidear-box{
    background: white;
    margin-bottom: 10px;
    h2{
        font-size: 16px;
        padding: 10px 16px 8px;
        font-weight: 600;
        line-height: 1;
    }
    .news-flash-list{
        padding: 5px 16px 8px;
        .flash-title{
            font-size: 13px;
            font-weight: 600;
            display: -webkit-box;
            -webkit-box-orient: vertical;
            -webkit-line-clamp: 1;
            overflow: hidden;
            text-justify: inter-ideograph;
            word-break: break-all; 
        }
        .flash-cate-date{
            display: flex;
            justify-content: space-between;
            align-items: center;
            font-size: 13px;
            color: #999;
        }
    }
}
</style>