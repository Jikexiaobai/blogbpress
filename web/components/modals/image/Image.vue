<template>
    <div class="img_pkg_box center opacity" :class="[isTrue && 'is_back_show']">
        <div class="img_pkg_container">
            <div class="img_pkg_title">
                <div class="img_pkg_title_l">
                    <span>插入图片</span>
                </div>
                <div class="img_pkg_title_r">
                    <a-icon  type="close" @click="cancel"/>
                </div>
            </div>

            <div class="img_pkg_tabs">
                <a-row type="flex" justify="center">
                    <a-col :span="4" @click="tabsChange(1)"  :class="tabsKey == 1 ? 'picked':''" class="img_pkg_tabs_item img_pkg_tabs_l">
                        上传图片
                    </a-col>
                    <a-col :span="4" @click="tabsChange(2)"  :class="tabsKey == 2 ? 'picked':''" class="img_pkg_tabs_item">
                        插入图片
                    </a-col>
                </a-row>
            </div>

            <div class="img_pkg_content">
                <div v-if="tabsKey == 1" class="img_pkg_content_upload">
                    <a-upload-dragger
                        name="file"
                        :customRequest="uploadImg"
                        :showUploadList="false"
                    >
                        <p class="ant-upload-drag-icon">
                        <a-icon type="inbox" />
                        </p>
                        <p class="ant-upload-text">
                            点击或者拖入文件到此处
                        </p>
                    </a-upload-dragger>
                </div>
                <div v-if="tabsKey == 2" class="img_pkg_content_list">
                   <ul>
                        <li  v-for="(item,index) in imgList" 
                            :key="index"
                            @click="handerActiveImg(item)"

                            :class="multiple != true ? activeImg == item ? 'picked':'' : activeImgList.indexOf(item) > -1 ? 'picked':''">
                            <div class="editor_image">
                                <img :src="item">
                            </div>
                        </li>
                   </ul>
                   <div class="img_pkg_content_list_pagination">
                        <a-pagination simple  :total="total" :page-size="4" @change="pageChange" />
                   </div>
                   <div class="img_pkg_content_list_ok">
                        
                        <a-space size="middle">
                            <a-button @click="cancel">取消</a-button>
                            <a-button @click="ascertain" type="primary">
                            确定
                            </a-button>
                        </a-space>
                   </div>
                </div>
            </div>


        </div>
    </div>
</template>

<script>
import {mapActions} from "vuex"
import api from "@/api/index"
export default {
    data() {
        return {
            multiple:false,

            mediaSize:8, // 文件大小
            mediaType:[".png",".jpg","jpeg","gif"], // 文件类型

            tabsKey:1,

            total: 0,
            queryParam: {
                page: 1,
                limit: 8,
                ext: null
            },
            imgList: [],

            activeImg: null,
            activeImgList: [],
            // 登录输入框
            isTrue: false,
            state: null // 准备（prepare） 确定（ ascertain） 取消（cancel）
        };
    },
    methods: {
        async confirm(
             multiple
        ) {
            this.open();
            this.multiple = multiple || false
            this.queryParam.ext = this.mediaType
            const res = await this.$axios.get(api.getfileList,{params: this.queryParam})
            if (res.code != 1) {
                this.$message.error(
                    res.message,
                    3
                )
                return
            }
            console.log(res)
            this.imgList = res.data.list || []
            this.total = res.data.total || 0
            return new Promise((resolve, reject) => {
                const target = { state: 'prepare'};
                const that = this
                let res = new Proxy(target, {
                    set(event, key, value) {
                        if (value === 'ascertain') {
                            if (!that.multiple) {
                                resolve(that.activeImg);
                            } else {
                                resolve(that.activeImgList);
                            }
                        } else {
                            reject(false);
                        }
                        return true
                    }
                });
                this.state = res;
            });
        },
        tabsChange(e){
            this.tabsKey = e
        },
        cancel(){
            this.state.state = "cancel"
            this.tabsKey = 1
            this.imgList = []
            this.activeImg = null
            this.activeImgList = []
            this.close()
        },
        ascertain(){
            this.state.state = "ascertain"
            this.close()
            this.tabsKey = 1
            this.imgList = []
            this.activeImg = null
        },
        async pageChange(page, limit){
            this.queryParam.page = page
            const res = await this.$axios.get(api.getfileList,{params: this.queryParam})
            this.imgList = res.data.list
        },
        handerActiveImg(e){
            if (!this.multiple) {
                this.activeImg = e
            } else {
                if (this.activeImgList.includes(e)) {
                    this.activeImgList.splice(this.activeImgList.indexOf(e),1)
                } else {
                    this.activeImgList.push(e)
                }
            }
        },
        async uploadImg(file){
            const link = await this.upload(file)
            this.activeImg = link[0]
            
            if (this.imgList.length > 7) {
                this.imgList.unshift(link[0])
                this.imgList.pop()
            } else {
                if (this.imgList.length > 0) {
                     this.imgList.unshift(link[0])
                }else{
                    this.imgList.push(link[0])
                }
            }
            
            this.tabsKey = 2
        },
        async upload(file){
            try {
                if ((this.mediaSize * 1024 * 1024) < file.file.size) {
                    this.$message.error(
                        "文件太大了",
                        3
                    )
                    return false
                }
                
                const type = this.$getType(file.file.name)
                if (this.mediaType.indexOf(type) == -1) {
                    this.$message.error(
                        "文件类型不正确",
                        3
                    )
                    return
                }

                let formData = new FormData();
                formData.append("file", file.file);

                const data = await this.$axios.post(api.postuploadFile,formData)
                if (data.code == 1) {
                    return data.data.link
                }
            } catch (error) {
                setTimeout(() => {
                    this.$notification.error({
                        message: '网络错误',
                        description: "请稍后再试"
                    })
                }, 1000)
            }
        },
        open() {
            this.isTrue = true;
        },
        close() {
            this.isTrue = false;
        },
        
    }
};
</script>

<style lang="less" scoped>
    .img_pkg_box {
        user-select: none;
        pointer-events: none;
        z-index: 20;
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        width: 100%;
        height: 100%;
        opacity: 0;
        visibility: hidden;
        transform: perspective(1px) scale(1.1);
        transition: visibility 0s linear .15s,opacity .15s 0s,transform .15s;
        display: flex;
        align-items: center;
        justify-content: center;
        .img_pkg_container{
            background-color: #fff;
            width: 22rem;
            margin: 0 auto;
            position: relative;
            background-image: url("/img/login.png");
            background-repeat: no-repeat;
            background-size: 100%;
            margin-top: -9%;

            .img_pkg_title{
                font-size: 13px;
                display: flex;
                justify-content: space-between;
                align-items: center;
                padding: 10px 20px;
                .img_pkg_title_l{
                    display: block;
                    align-items: center;
                    width: 80%;
                    overflow: hidden;
                    white-space: nowrap;
                    text-overflow: ellipsis;
                }
            }
            .img_pkg_tabs{
                .img_pkg_tabs_l{
                    margin-right: 10px;
                }
                .img_pkg_tabs_item{
                    text-align: center;
                    cursor: pointer;
                }
                .picked{
                    color: brown;
                }
            }
            .img_pkg_content{
                padding: 20px 20px;
                .img_pkg_content_upload{
                    width: 100%;
                }
                .img_pkg_content_list{
                    
                    ul{
                        display: flex;
                        flex-flow: wrap;
                        overflow-y: auto;
                        max-height: 326px;
                        li{
                            width: 25%;
                            border: 1px solid #fff;
                            position: relative;
                            cursor: pointer;
                            .editor_image{
                                height: 0;
                                padding-top: 100%;
                                position: relative;
                                background-color: #f5f5f5;
                                img{
                                    position: absolute;
                                    height: 100%;
                                    width: 100%;
                                    top: 0;
                                    left: 0;
                                }
                            }
                        }
                        .picked::after{
                            content: '✓';
                            position: absolute;
                            width: 20px;
                            height: 20px;
                            background: #f44336;
                            color: #fff;
                            top: 5px;
                            right: 5px;
                            display: flex;
                            align-items: center;
                            justify-content: center;
                            border-radius: 100%;
                            border: 1px solid #fff;
                        }
                    }
                    .img_pkg_content_list_pagination{
                        margin-top: 20px;
                        display: flex;
                        justify-content: flex-end;
                    }
                    .img_pkg_content_list_ok{
                        margin-top: 20px;
                        display: flex;
                        justify-content: flex-end;
                    }
                }
            }
        }
        
    }
    .is_back_show {
        opacity: 1 !important;
        background: rgba(42, 44, 48, 0.7);
        pointer-events: auto !important;
        visibility: visible;
        transform: perspective(1px) scale(1);
        transition: visibility 0s linear 0s,opacity .15s 0s,transform .15s;
    }

    @media only screen and (max-width: 768px) {
    }
</style>