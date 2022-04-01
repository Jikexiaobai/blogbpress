<template>
    <div class="sidebar-box">
        <div class="title">
            <h2>创建圈子</h2>
        </div>
        <ul>
            <li>
                <nuxt-link :to="{path:'#' }" class="item-link">   
                   什么是圈子
                </nuxt-link>
            </li>
            <li>
                <nuxt-link :to="{path:'#' }" class="item-link">   
                   我可以做什么？
                </nuxt-link>
            </li>
            <li>
                <nuxt-link :to="{path:'#' }" class="item-link">   
                   如何创建
                </nuxt-link>
            </li>
        </ul>
    </div>
</template>

<style lang="less" scoped>
.sidebar-box{
    background: white;
    margin-bottom: 10px;
    padding: 20px 20px 10px 20px;
    .title{
        margin-bottom: 10px;
        h2{
            font-size: 12px;
            line-height: 1;
            color: #999;
        }
    }
    ul{
        li{
            padding: 10px 0;
            border: 0;
            font-size: 13px;
            box-sizing: border-box;
            a{
                color: inherit;
                text-decoration: none;
                vertical-align: top;
                display: block;
                overflow: hidden;
                white-space: nowrap;
                text-overflow: ellipsis;
                font-size: 12px;
                position: relative;
                padding-left: 17px;
            }
            a::before {
                background-color: #3860f4;
                    content: '';
                    width: 6px;
                    height: 6px;
                    position: absolute;
                    left: 0;
                    top: 5px;
                    border-radius: 2px;
            }
        }
    }
}
</style>