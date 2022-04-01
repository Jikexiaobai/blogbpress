<template>
    <div class="upload_file_box center opacity" :class="[isTrue && 'is_back_show']">
        <div class="upload_file_container">
            <div class="upload_file_title">
                <div class="upload_file_title_l">
                    <span>插入文件</span>
                </div>
                <div class="upload_file_title_r">
                    <a-icon  type="close" @click="cancel"/>
                </div>
            </div>


            <div class="upload_file_content">
                <div  class="upload_file_content_box">
                    <a-upload-dragger
                        name="file"
                        :customRequest="uploadFlie"
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
            </div>


        </div>
    </div>
</template>

<script>
import { postUploadFile} from '@/api/upload'
import {getType} from "@/utils/tools/fileType"
import { mapState } from "vuex"
export default {
    data() {
        return {
            mediaSize:1000, // 文件大小
            mediaType:null,
            imageType:[".png",".jpg",".gif",".jpeg"], // 图片类型
            audioType:[".mp3",".wav"], // 图片类型
            videoType:[".mp4"], // 图片类型
            otherType:[".crt"], // 其他类型

            activeLink: null,
            isTrue: false,
            state: null // 准备（prepare） 确定（ ascertain） 取消（cancel）
        };
    },
    computed:{
        ...mapState("system",["file"])
    },
    methods: {
        async confirm(
            mediaType
        ) {
            this.open();
            this.mediaType = mediaType || "Image"
            this.imageType = this.file.imageType
            this.audioType = this.file.audioType
            this.videoType = this.file.videoType
            this.otherType = this.file.otherType
            this.mediaSize = this.file.fileSize
            // this.queryParam.ext = this.mediaType
  
            return new Promise((resolve, reject) => {
                const target = { state: 'prepare'};
                const that = this
                let res = new Proxy(target, {
                    set(event, key, value) {
                        if (value === 'ascertain') {
                            resolve(that.activeLink);
                        } else {
                            reject(false);
                        }
                        return true
                    }
                });
                this.state = res;
            });
        },
        

        async uploadFlie(file){
            if (!file) return;

            //  
            if ((this.mediaSize * 1024 * 1024) < file.file.size) {
                this.$message.error(
                    "文件太大了",
                    3
                )
                return false
            }
            
            const type = getType(file.file.name)
            if (this.mediaType === "Image" && this.imageType.indexOf(type) == -1 ) {
                this.$message.error(
                    "文件类型不正确",
                    3
                )
                return
            }
            if (this.mediaType === "Audio" && this.audioType.indexOf(type) == -1 ) {
                this.$message.error(
                    "文件类型不正确",
                    3
                )
                return
            }
            if (this.mediaType === "Video" && this.videoType.indexOf(type) == -1 ) {
                this.$message.error(
                    "文件类型不正确",
                    3
                )
                return
            }
            if (this.mediaType === "Other" && this.otherType.indexOf(type) == -1 ) {
                this.$message.error(
                    "文件类型不正确",
                    3
                )
                return
            }

            let formData = new FormData();
            formData.append("file", file.file);

            const res = await postUploadFile(formData)
            this.activeLink = res.data.link[0]
            this.ascertain()
            return
        },

        cancel(){
            this.state.state = "cancel"
            this.activeLink = null
            this.close()
        },
        ascertain(){
            this.state.state = "ascertain"
            this.close()
            this.activeLink = null
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
    .upload_file_box {
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
        .upload_file_container{
            background-color: #fff;
            width: 22rem;
            margin: 0 auto;
            position: relative;
            background-image: url("/img/login.png");
            background-repeat: no-repeat;
            background-size: 100%;
            margin-top: -9%;

            .upload_file_title{
                font-size: 13px;
                display: flex;
                justify-content: space-between;
                align-items: center;
                padding: 10px 20px;
                .upload_file_title_l{
                    display: block;
                    align-items: center;
                    width: 80%;
                    overflow: hidden;
                    white-space: nowrap;
                    text-overflow: ellipsis;
                }
            }
            .upload_file_tabs{
                .upload_file_tabs_l{
                    margin-right: 10px;
                }
                .upload_file_tabs_item{
                    text-align: center;
                    cursor: pointer;
                }
                .picked{
                    color: brown;
                }
            }
            .upload_file_content{
                padding: 20px 20px;
                .upload_file_content_box{
                    width: 100%;
                    height: 200px;
                    .hash-percentage{
                        margin: 10px 0;
                    }
                }

                .upload_file_content_list{
                    
                    ul{
                        display: flex;
                        flex-flow: wrap;
                        overflow-y: auto;
                        max-height: 326px;
                        li{
                            width: 50%;
                            border: 1px solid #fff;
                            position: relative;
                            cursor: pointer;
                            .editor_image{
                                height: 0;
                                padding-top: 60%;
                                position: relative;
                                background-color: #f5f5f5;
                                video{
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
                    .upload_file_content_list_pagination{
                        margin-top: 20px;
                        display: flex;
                        justify-content: flex-end;
                    }
                    .upload_file_content_list_ok{
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