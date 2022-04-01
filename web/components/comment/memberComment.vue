<template>
    <div class="commnet-item">
        <a-avatar class="commnet-avatar" src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png" />
        <div class="commnet-info">
            <div class="comment-user-name">
                测试名称
            </div>
            <div class="comment-content">
                这是评论的内容
            </div>
            <div class="comment-meta">
                <a-space size="small">
                    <span class="comment-meta-date">2018-14-41</span>
                    <span>
                        <a-icon type="like"/>
                        <span>546</span>
                    </span>
                    <a-button type="link">
                        回复
                    </a-button>
                </a-space>
            </div>
            <!-- <div class="comment-repy-box">
                <a-textarea placeholder="Basic usage" :rows="4" />
                <div class="comment-repy-button">
                    <a-button type="primary">
                    发表回复
                    </a-button>
                </div>
            </div> -->
        </div>
        <div class="commnet-post">
            <!-- <img src="@/assets/img/0.jpg"> -->
            <div class="commnet-post-title">电子音乐常见的一些曲风写法【电子音乐教程】</div>
        </div>
    </div>
</template>

<style lang="less" scoped>
.commnet-item{
    width: 100%;
    padding: 10px;
    display: flex;
    border-bottom: 1px solid #e0e7eb;
    flex-direction: row;
    .commnet-info{
        margin: 0 10px;
        flex: 1;
        .comment-user-name{
            margin-bottom: 10px;
            color: #757575;
        }
        .comment-content{
            margin-bottom: 10px;
            color: #212121;
            font-size: 14px
        }
        .comment-meta{
            margin-bottom: 10px;
            .comment-meta-date{
                    font-size: 14px;
                    color: #999;
                    text-align: left;
                    display: inline-block;
            }
        }
        .comment-repy-box{
            .comment-repy-button{
                margin-top: 10px;
                display: flex;
                justify-content: flex-end;
            }
        }
    }
    .commnet-post{
        width: 100px;
        text-align: center;
        margin-left: 40px;
        img{
            width: 100px;
            height: 60px;
            border-radius: 4px;
        }
        .commnet-post-title{
            margin-top: 10px;
            width: 100%;
            font-size: 12px;
            color: #505050;
            overflow:hidden;
            text-overflow:ellipsis;
            white-space:nowrap;
        }
    }
}
</style>