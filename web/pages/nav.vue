<template>
    <div class="container">
        
        <div class="left" :style="{width:collapsed?'200px':'80px'}">
            <div class="left-logo">
                <h2 v-if="collapsed">LOGO</h2>
                <h2 v-if="!collapsed">L</h2>
            </div>
            <div class="left-menu">
                <a-menu
                    :default-selected-keys="['1']"
                    :default-open-keys="['sub1']"
                    mode="inline"
                    theme="light"
                    :inline-collapsed="!collapsed"
                    >
                    <a-menu-item key="1" @click="goFexd(1)">
                            <a-icon type="pie-chart" />
                            <span>常用推荐</span>
                    </a-menu-item>
                    <a-menu-item key="2" @click="goFexd(2)">
                            <a-icon type="desktop" />
                            <span>社区资讯</span>
                    </a-menu-item>
                    <a-menu-item key="3">
                        <a-icon type="mail" />
                        <span>常用工具</span>
                    </a-menu-item>
                    <a-sub-menu key="sub1">
                        <span slot="title">
                            <a-icon type="mail" />
                            <span>资源素材</span>
                        </span>
                        <a-menu-item key="5">
                            音乐素材
                        </a-menu-item>
                        <a-menu-item key="6">
                            视频素材
                        </a-menu-item>
                    </a-sub-menu>
                    <a-sub-menu key="sub2">
                        <span slot="title">
                            <a-icon type="mail" />
                            <span>后端框架</span>
                        </span>
                        <a-menu-item key="5">
                            Gin
                        </a-menu-item>
                        <a-menu-item key="6">
                            Goframe
                        </a-menu-item>
                    </a-sub-menu>
                    <a-sub-menu key="sub3">
                        <span slot="title">
                            <a-icon type="mail" />
                            <span>前端框架</span>
                        </span>
                        <a-menu-item key="5">
                            Vue
                        </a-menu-item>
                        <a-menu-item key="6">
                            React
                        </a-menu-item>
                    </a-sub-menu>

                </a-menu>
            </div>
            <div class="left-about">
                sdf
            </div>
        </div>
        <div class="right" :style="{marginLeft:collapsed?'200px':'80px'}">
            
            <div class="right-search" style="background-image: linear-gradient(45deg, #8618db 0%, #d711ff 50%, #460fdd 100%);">
                <div class="right-header">
                    <div class="top-header-left-menu">
                        <div class="menu-fold" @click="changeSiber">
                            <a-icon :type="collapsed ? 'menu-fold':'menu-unfold'" />
                        </div>
                        <ul class="menu-list">
                            <li class="item">
                                首页
                            </li>
                        </ul>
                    </div>
                    <div class="top-header-right-menu">

                    </div>
                </div>
                
                <div class="search-box">
                    <div class="search-form">
                        <a-input-search 
                        size="large" :placeholder="navTop|searchText" />
                    </div>
                    <ul class="search-tab">
                        <li class="item" @click="changeNavTop(0)">
                            <span :class="navTop == 0 ? 'title active': 'title'">百度</span>
                        </li>
                        <li class="item" @click="changeNavTop(1)">
                            <span :class="navTop == 1 ? 'title active': 'title'">谷歌</span>
                        </li>
                        <li class="item" @click="changeNavTop(2)">
                            <span :class="navTop == 2 ? 'title active': 'title'">360</span>
                        </li>
                        <li class="item" @click="changeNavTop(3)">
                            <span :class="navTop == 3 ? 'title active': 'title'">搜狗</span>
                        </li>
                        <li class="item" @click="changeNavTop(4)">
                            <span :class="navTop == 4 ? 'title active': 'title'">必应</span>
                        </li>
                    </ul>
                </div>       
            </div>

            <div class="right-body">
                <ul>
                    <li class="item links">
                        <div class="cate-title">
                            <a-icon class="aicon" type="tag" />
                            <h4 class="title">分类名称</h4>
                        </div>
                        <div class="second-cate-box">
                            <div class="second-cate">
                                <SlidNav :list="[{
                                    label:'苹果',
                                    path:'/ping'
                                },{
                                    label:'香蕉',
                                    path:'/xiangj'
                                },{
                                    label:'橘子',
                                    path:'/juzi'
                                },]" value="/ping"/>
                            </div>
                            <div class="fill"></div>
                        </div>
                        <div class="info-list">
                            <a-row :gutter="[{md:12},{md:12}]">
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>

                            </a-row>
                        </div>
                    </li>
                    <li class="item links">
                        <div class="cate-title">
                            <a-icon class="aicon" type="tag" />
                            <h4 class="title">分类名称</h4>
                        </div>
                        <div class="second-cate-box">
                            <div class="second-cate">
                                <SlidNav :list="[{
                                    label:'苹果',
                                    path:'/ping'
                                },{
                                    label:'香蕉',
                                    path:'/xiangj'
                                },{
                                    label:'橘子',
                                    path:'/juzi'
                                },]" value="/ping"/>
                            </div>
                            <div class="fill"></div>
                        </div>
                        <div class="info-list">
                            <a-row :gutter="[{md:12},{md:12}]">
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>

                            </a-row>
                        </div>
                    </li>
                    <li class="item links">
                        <div class="cate-title">
                            <a-icon class="aicon" type="tag" />
                            <h4 class="title">分类名称</h4>
                        </div>
                        <div class="second-cate-box">
                            <div class="second-cate">
                                <SlidNav :list="[{
                                    label:'苹果',
                                    path:'/ping'
                                },{
                                    label:'香蕉',
                                    path:'/xiangj'
                                },{
                                    label:'橘子',
                                    path:'/juzi'
                                },]" value="/ping"/>
                            </div>
                            <div class="fill"></div>
                        </div>
                        <div class="info-list">
                            <a-row :gutter="[{md:12},{md:12}]">
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>

                            </a-row>
                        </div>
                    </li>
                    <li class="item links">
                        <div class="cate-title">
                            <a-icon class="aicon" type="tag" />
                            <h4 class="title">分类名称</h4>
                        </div>
                        <div class="second-cate-box">
                            <div class="second-cate">
                                <SlidNav :list="[{
                                    label:'苹果',
                                    path:'/ping'
                                },{
                                    label:'香蕉',
                                    path:'/xiangj'
                                },{
                                    label:'橘子',
                                    path:'/juzi'
                                },]" value="/ping"/>
                            </div>
                            <div class="fill"></div>
                        </div>
                        <div class="info-list">
                            <a-row :gutter="[{md:12},{md:12}]">
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>

                            </a-row>
                        </div>
                    </li>
                    <li class="item links">
                        <div class="cate-title">
                            <a-icon class="aicon" type="tag" />
                            <h4 class="title">分类名称</h4>
                        </div>
                        <div class="second-cate-box">
                            <div class="second-cate">
                                <SlidNav :list="[{
                                    label:'苹果',
                                    path:'/ping'
                                },{
                                    label:'香蕉',
                                    path:'/xiangj'
                                },{
                                    label:'橘子',
                                    path:'/juzi'
                                },]" value="/ping"/>
                            </div>
                            <div class="fill"></div>
                        </div>
                        <div class="info-list">
                            <a-row :gutter="[{md:12},{md:12}]">
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>

                            </a-row>
                        </div>
                    </li>
                    <li class="item 6xx" id="6xx">
                        <div class="cate-title">
                            <a-icon class="aicon" type="tag" />
                            <h4 class="title">分类名称</h4>
                        </div>
                        <div class="second-cate-box">
                            <div class="second-cate">
                                <SlidNav :list="[{
                                    label:'苹果',
                                    path:'/ping'
                                },{
                                    label:'香蕉',
                                    path:'/xiangj'
                                },{
                                    label:'橘子',
                                    path:'/juzi'
                                },]" value="/ping"/>
                            </div>
                            <div class="fill"></div>
                        </div>
                        <div class="info-list">
                            <a-row :gutter="[{md:12},{md:12}]">
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>
                                <a-col  :span="4">
                                    <a-tooltip placement="bottom">
                                        <template slot="title">
                                            的分公司的粉红色的还是
                                        </template>
                                        <div class="item">
                                            <div class="item-top">
                                                <img class="item-cover" src="/img/xxxx1.jpg" alt="xxxx">
                                                <div class="item-top-title">
                                                    <strong class="title">测试测试测试测试测试测试测试测试测试测试</strong>
                                                    <p>的分公司的粉红色的还是</p>
                                                </div>
                                                <a-icon class="aicon" type="right-circle" />
                                            </div>
                                        </div>
                                    </a-tooltip>
                                </a-col>

                            </a-row>
                        </div>
                    </li>
                </ul>
                

            </div>
        </div>
   
    </div>  
</template>
<script>
import SlidNav from "@/components/nav/slidNav"
import { mapState } from "vuex"
import { Swiper, SwiperSlide } from 'vue-awesome-swiper'
var vm  =null
export default {
    components:{
        Swiper,
        SwiperSlide,
        SlidNav
    },
    layout: "nav",
    filters: {
        searchText: function (value) {
           switch (value) {
                case 0: 
                    return "百度搜索";
                case 1: 
                    return "谷歌搜索";
                case 2: 
                    return "360搜索";
                case 3: 
                    return "搜狗搜索";
                case 4: 
                    return "必应搜索";
           }
        }
    },
    data() {
        return {
            collapsed: true,
            navTop:0,
            swiperOption: {
                
                //循环
                loop:true,
                //每张播放时长3秒，自动播放
                autoplay:true,
                //滑动速度
                speed:1000,
                // grabCursor:true,
                preventClicksPropagation : false,
                // delay:1000
                // on: {
                //     click: function () {
                //         vm.goLink(this.realIndex)
                //     },
                //     slideChange: function () {
                //         vm.activeIndex = this.realIndex
                //     }
                // },
                pagination:false,
                // pagination: {
                //     // type: 'progressbar',
                //     el: '.swiper-pagination',
                //     dynamicBullets: false
                // },
            },
            swiperList:[],
        };
    },
    computed:{
        ...mapState(["design"]),
        swiper() {
            return this.$refs.widgetCarouselOne.swiper
        }
    },
    created(){
        vm = this
        // this.list = JSON.parse(this.info.list)
        // if (this.list.length == 1) {
        //     this.llist.push(this.list[0])
        //     this.rlist.push(this.list[0])
        // }else if (this.list.length == 2) {
        //     this.llist.push(this.list[0])
        //     this.rlist.push(this.list[1])
        // }else if (this.list.length == 3 ){
        //     this.llist.push(this.list[0])
        //     this.rlist.push(this.list[1])
        //     this.rlist.push(this.list[2])
        // }else if (this.list.length == 4 ){
        //     this.llist.push(this.list[0])
        //     this.llist.push(this.list[1])
        //     this.rlist.push(this.list[2])
        //     this.rlist.push(this.list[3])
        // }else if (this.list.length > 4){
        //     let tmpCount = this.list.length - 2
        //     this.list.forEach((item,index) => {
        //         if (index < tmpCount) {
        //             this.llist.push(item)
        //         }
        //         if (index > tmpCount) {
        //             this.rlist.push(item)
        //         }
        //     });
        // }
      
    },
    methods:{
        changeNavTop(e){
            this.navTop = e
        },
        changeSiber(){
            this.collapsed = !this.collapsed
        },
        goFexd(e){
            let jump = document.querySelectorAll(`.links`)
            // 获取需要滚动的距离
            let total = jump[e].offsetTop
            console.log(total)

            let distance = document.documentElement.scrollTop || document.body.scrollTop
                // 平滑滚动，时长500ms，每10ms一跳，共50跳
                let step = total / 50
                if (total > distance) {
                smoothDown()
                } else {
                let newTotal = distance - total
                step = newTotal / 50
                    smoothUp()
                }
                function smoothDown () {
                    if (distance < total) {
                        distance += step
            　　　　　　　document.body.scrollTop = distance
                        document.documentElement.scrollTop = distance
                        setTimeout(smoothDown, 6)
                    } else {
                        document.body.scrollTop = total
                        document.documentElement.scrollTop = total
                    }
                }
                function smoothUp () {
                    if (distance > total) {
                        distance -= step
            　　　　　　　document.body.scrollTop = distance
                        document.documentElement.scrollTop = distance
                        setTimeout(smoothUp, 6)
                    } else {
                        document.body.scrollTop = total
                        document.documentElement.scrollTop = total
                    }
            } 
        },
    }
};
</script>


<style lang="less" scoped>
.container{
    // display: flex;
    .left{
        transition: width .25s;
        position: fixed;
        height: 100vh;
        background-color: white;
        
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        .left-logo{
            height: 60px;
            background-color: white;
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 28px;
            font-weight: bold;
        }
        .left-menu{
            flex: 1;
            width: 100%;
            /deep/ .ant-menu-inline{
                border: none;
            }
        }
        .left-about{

        }
    }
    .right{
        transition: margin .25s;
        .right-search{
            // z-index: 2;
            position: relative;
            // margin-top: -75px;
            // padding-bottom: 180px;
            height: 490px;
            background-size: 400%;
            background-position: 0% 100%;
            -webkit-animation: gradient 7.5s ease-in-out infinite;
            animation: gradient 7.5s ease-in-out infinite;
            @keyframes gradient {
                50% {
                    background-position: 100% 0;
                }
            }
            .right-header{
                background: rgba(255, 255, 255, 0);
                z-index: 20;
                height: 60px;
                display: flex;
                justify-content: space-between;
                align-items: center;
                padding: 0 20px;
                .top-header-left-menu{
                    display: flex;
                    justify-content: space-between;
                    align-items: center;
                    .menu-fold{
                        color: white;
                        font-size: 26px;
                        margin-right: 10px;
                        cursor: pointer;
                    }
                    .menu-list{
                        margin-left: 10px;
                        display: flex;
                        justify-content: space-between;
                        align-items: center;
                        .item{
                            color: white;
                            font-size: 16px;
                            font-weight: 600;
                        }
                    }  
                }
            }

            .search-box{
                width: 100%;
                height: 430px;
                display: flex;
                justify-content: center;
                align-items: center;
                flex-direction: column;
                .search-tab{
                    display: flex;
                    .item{
                        cursor: pointer;
                        padding: 10px 20px;
                        display: flex;
                        justify-content: center;
                        align-items: center;
                        flex-direction: column;
                        .title{
                            position: relative;
                            opacity: .5;
                            transition: opacity .3s;
                            font-size: 14px;
                            color: #ffffff;
                            // font-weight: 700;
                            margin-bottom: 10px;
                        }
                        .active{
                            opacity: 1;
                        }
                        .active::before{
                            content: '';
                            border-width: 8px 8px 0px 8px;
                            border-style: solid;
                            border-color: rgba(0,0,0,.2) transparent transparent;
                            position: absolute;
                            left: 50%;
                            top: -10px;
                            margin-left: -8px;
                        }

                        .boder{
                            width: 100%;
                            height: 3px;
                            background-color: white;
                            border-radius: 8px;
                        }
                    }
                    
                }
                .search-form{
                    padding: 0 500px;
                    width: 100%;
                    // background:none;

                    /deep/.ant-input{
                        padding-left: 30px;
                        border-radius: 50px;
                        height: 50px;
                        color: #fff;
                        fill: #fff;
                        background-color: rgba(204, 204, 204, 0.2);
                        outline: none;
                        border:none;
                        box-shadow: none;
                    }
                    
                    /deep/ .ant-input-search-icon{
                        padding-right: 10px;
                        font-size: 18px;
                        font-weight: 700;
                        color: #fff;
                    }
                }
            }
        } 
        .right-body{
            // height: 100%;
            padding: 0 20px;
            background-color: #f9f9f9;
            // position: relative;
            ul{
                .item{
                    padding: 20px 0;
                    .cate-title{
                        display: flex;
                        align-items: center;
                        font-size: 1.125rem;
                        color: #555;
                        margin-bottom: 10px;
                        .aicon{
                            font-size: 22px;
                            margin-right: 10px;
                        }
                    }
                    .second-cate-box{
                        display: flex;
                        justify-content: space-between;
                        align-items: center;
                        .fill{
                            flex: 1;
                        }
                    }
                    .info-list{
                        margin-top: 20px;
                        .item{
                            cursor: pointer;
                            padding: 20px 10px;
                            background: #fff;
                            border-radius: 6px;
                            transition: all .3s ease;
                            transform: translateY(0px);
                            box-shadow: 0px 0px 20px -5px rgba(158,158,158, 0.2);
                            .item-top{
                                display: flex;
                                align-items: center;
                                justify-content: space-between;
                                .item-cover{
                                    object-fit: cover;
                                    height: 40px;
                                    width: 40px;
                                    border-radius: 6px;
                                }
                                .item-top-title{
                                    flex: 1;
                                    margin: 0 5px;
                                    .title{
                                        color: #000;
                                        font-size: 16px;
                                        font-weight: bold;
                                        overflow: hidden;
                                        text-overflow: ellipsis;
                                        display: -webkit-box;
                                        -webkit-line-clamp: 1;
                                        -webkit-box-orient: vertical;
                                    }
                                    p{
                                        color: #646a73;
                                        font-size: 14px;
                                        overflow: hidden;
                                        text-overflow: ellipsis;
                                        display: -webkit-box;
                                        -webkit-line-clamp: 1;
                                        -webkit-box-orient: vertical;
                                    }
                                }
                                .aicon{
                                    color: #e9eaeb;
                                    font-size: 18px;
                                    font-weight: 700;
                                }
                            }
                            
                        }
                        .item:hover{
                            transform: translateY(-6px);
                            box-shadow: 0 26px 40px -24px rgb(0 36 100 / 30%);
                            .item-top-title{
                                .title{
                                    color: red;
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
</style>
