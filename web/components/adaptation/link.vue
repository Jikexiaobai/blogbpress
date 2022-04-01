<template>
    <div class="link">
        <div class="content">
            <img class="cover" src="/img/test1.jpg" alt="">
            <div class="info">
                <h2 class="title">
                    连接标题连接标题连接标
                </h2>
                <p class="desc">连接描述连接描述连接描述连接描述连接描述连接描述连接描述连接描述连接描述连接描述</p>
            </div>
        </div>
    </div>
</template>

<style lang="less" scoped>
.link{
    margin-top: 10px;
    .content{
        cursor: pointer;
        padding: 10px;
        border-radius: 4px;
        background: #f5f6f7;
        display: flex;
        .cover{
            border-radius: 4px;
            width: 50px;
            height: 50px;
            object-fit: cover;
            margin-right: 10px;
        }
        .info{
            flex: 1;
            .title{
                line-height: 20px;
                color: #797c80;
                font-size: 16px;
                font-weight: 600;
                margin-bottom: 8px;
                display: -webkit-box;
                -webkit-box-orient: vertical;
                -webkit-line-clamp: 1;
                overflow: hidden;
                word-break: break-all;
            }
            .desc{
                line-height: 15px;
                color: #797c80;
                font-size: 13px;
                margin-bottom: 8px;
                display: -webkit-box;
                -webkit-box-orient: vertical;
                -webkit-line-clamp: 1;
                overflow: hidden;
                word-break: break-all;
            }
        }
    }
}
</style>