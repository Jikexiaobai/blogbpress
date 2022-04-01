<template>
    <div>
        <editor id="editor-box" 
            v-model="content"
            :init="init"  
            :disabled="disabled"></editor>
    </div>
</template>
<script>
import tinymce from 'tinymce/tinymce'
import Editor from '@tinymce/tinymce-vue'
import 'tinymce/themes/silver'
import 'tinymce/icons/default'
import 'tinymce/plugins/image'
import 'tinymce/plugins/table'
import 'tinymce/plugins/lists'
import 'tinymce/plugins/code'
import 'tinymce/plugins/autoresize'
import 'tinymce/plugins/codesample'
import 'tinymce/plugins/fullscreen'
import 'tinymce/plugins/link'
import 'tinymce/plugins/hr'
import 'tinymce/plugins/quickbars'
import 'tinymce/plugins/paste'
import 'tinymce/plugins/advlist'
export default {
    name: "tinymce-editor",
    components: {
        Editor
    },
    props: {
      valueContont: {
        type: String,
        default: ''
      },
      height: {
        type: Number,
        default: 400
      },
      disabled: {
        type: Boolean,
        default: false
      },
      plugins: {
        type: [String, Array],
        default: 'code,lists,advlist,autoresize,codesample,link,hr,fullscreen,quickbars,image,paste'
      },
      toolbar: {
        type: [String, Array],
        default: 'undo redo | formatselect  | h2Title hr bold italic forecolor backcolor | alignleft aligncenter alignright alignjustify | bullist numlist outdent indent | lists imagelibrary  | removeformat'
      }
    },
    data(){
        return{
            content:"",
            init: {
                body_class:'entry-content',
                language_url: '/tinymce/lang/zh_CN.js',// 语言包的路径
                language: 'zh_CN',//语言
                skin_url: '/tinymce/skins/ui/oxide',// skin路径
                content_css:'/tinymce/skins/content/default/content.css',
                min_height:this.height,//编辑器高度
                branding: false,//是否禁用“Powered by TinyMCE”
                menubar: false,//顶部菜单栏显示
                content_style: "body {padding: 10px}",
                body_style: '.entry-content {margin:16px}',
                statusbar: false,
                toolbar_sticky:true,
                relative_urls: false,
                remove_script_host: false,
                image_advtab:true,
                plugins: this.plugins,
                toolbar1: this.toolbar, 
                block_formats:"Paragraph=p;标题1=h1;标题2=h2;",
                // plugins: 'code,lists,advlist,autoresize,codesample,link,hr,fullscreen,quickbars,image,paste',
                // toolbar1:'bold italic forecolor backcolor | h2Title blockquote codesample numlist bullist | link imagelibrary hr | removeformat  | fullscreen',
                image_caption: true,
                content_style: 'img {max-width: 100%;}',
                oninit : "setPlainText",
                quickbars_selection_toolbar: false,
                quickbars_insert_toolbar: false,
                paste_auto_cleanup_on_paste : true,
                paste_remove_styles: true,
                paste_remove_styles_if_webkit: true,
                paste_strip_class_attributes: true,
                setup: (editor)=>{
                //重写标题按钮
                editor.ui.registry.addIcon('h2-title', '<svg t="1604508549782" class="icon" viewBox="0 0 1216 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="6794" width="24" height="24"><path d="M285.803 419.24V155.595h87.328a23.73 23.73 0 1 0 0-47.461H142.945a23.73 23.73 0 1 0 0 47.46h87.328v263.673a27.765 27.765 0 0 0 55.53 0z m183.199-246.058c10.441 0 18.984-3.323 26.103-9.967s10.916-15.662 10.916-26.104-3.796-18.984-10.44-25.629a37.705 37.705 0 0 0-26.579-9.993c-10.441 0-18.984 3.323-26.104 9.967s-10.441 15.188-10.441 25.629 3.322 18.984 10.441 26.103c7.12 6.645 15.663 9.967 26.104 9.967z m27.053 246.797V228.685a27.053 27.053 0 1 0-54.106 0v191.267a27.053 27.053 0 1 0 54.106 0z m192.216 4.746a22.307 22.307 0 0 0-22.306-22.307h-9.967c-5.695 0-9.492-1.424-12.34-4.272-2.847-3.322-3.797-7.593-3.797-13.289V246.245h32.274a22.307 22.307 0 1 0 0-44.613H639.86v-73.09a27.053 27.053 0 1 0-54.105 0v73.116h-21.832a22.307 22.307 0 1 0 0 44.613h21.832v138.586c0 19.934 4.746 34.647 14.238 45.088 9.967 11.39 25.629 17.086 46.986 17.086h18.985a22.307 22.307 0 0 0 22.306-22.306z m105.364-4.746V128.542a27.053 27.053 0 1 0-54.106 0v291.437a27.053 27.053 0 1 0 54.106 0z m168.486 33.697c32.748 0 59.8-9.492 81.633-28.002a108.712 108.712 0 0 0 26.763-36.15c5.748-12.498 1.898-21.278-11.866-21.278h-23.07c-12.183 0-14.687 7.067-16.533 10.204a68.977 68.977 0 0 1-17.06 19.696c-10.441 7.594-23.73 11.391-40.342 11.391-19.459 0-34.646-6.17-45.087-18.035-10.442-11.865-16.612-29.426-18.51-52.207h158.598c10.125-0.264 26.315-2.716 23.705-27.264-3.375-31.377-12.789-56.847-28.055-76.2-21.358-27.528-52.207-40.817-92.549-40.817-36.545 0-65.496 12.34-86.379 37.494-21.832 24.68-32.273 55.054-32.273 91.6 0 40.816 11.39 72.615 34.172 95.87 21.357 22.307 50.308 33.698 86.853 33.698z m54.527-155.25H899.051c3.296-19.46 9.94-34.172 19.907-44.139 10.441-10.441 24.205-15.187 42.24-15.187 34.014 0 54.686 15.609 62.042 47.487 2.32 10.02 0.712 11.575-6.592 11.812z m45.563 252.808a26.367 26.367 0 0 1 3.059 52.55l-3.06 0.185H152.86a26.367 26.367 0 0 1-3.086-52.55l3.085-0.185h909.352z m0 158.203a26.367 26.367 0 0 1 3.059 52.55l-3.06 0.185H152.86a26.367 26.367 0 0 1-3.086-52.55l3.085-0.184h909.352zM798.539 867.641a26.367 26.367 0 0 1 3.059 52.55l-3.059 0.184H152.86a26.367 26.367 0 0 1-3.085-52.55l3.085-0.184h645.68z" p-id="6795"></path></svg>');
                editor.ui.registry.addToggleButton('h2Title', {
                    icon: 'h2-title',
                    tooltip:'Title',
                    onAction: function (_) {
                    editor.execCommand('mceToggleFormat', false, 'h2');
                    },
                    onSetup: function (api) {
                    editor.formatter.formatChanged('h2', function (state) {
                        api.setActive(state);
                    });
                    }
                });


                //图片按钮
                editor.ui.registry.addButton('imagelibrary', {
                    tooltip: "Insert image",
                    icon: 'image',
                    onAction: ()=>{

                        this.$Upload().then((res)=>{
                            if (res != false) {
                                let html = `<p><img src="${res}" /></p>`
                                editor.insertContent(html)
                            }
                        }).catch((err)=>{
                            console.log(err)
                        })
                    }
                });
        
                //视频按钮按钮
                editor.ui.registry.addIcon('b2-video', '<svg class="Zi Zi--InsertVideo" width="24" height="24"><path d="M10.5 15c-.4.3-.8 0-.8-.5v-5c0-.6.4-.8.8-.5l4.3 2.5c.5.3.5.7 0 1L10.5 15zM5 5S3 5 3 7v10s0 2 2 2h14c2 0 2-2 2-2V7c0-2-2-2-2-2H5z" fill-rule="evenodd"/></svg>');
                editor.ui.registry.addButton('b2video', {
                    tooltip: "Insert video",
                    icon: 'b2-video',
                    onAction: this.showVideoBoxAc
                });
                
                },
                init_instance_callback(editor){
                    
                    // if(b2_write_data.edit_content){
                    //     editor.setContent(b2_write_data.edit_content);
                    // }
                }
            },
        }
    },
    // created(){
        
    // },
    mounted(){ 
        tinymce.init({});
    },
    watch: {
        valueContont(newValue) {
            this.content = newValue
        },
        content(newValue) {
            this.$emit('writeContent', newValue);
        }
    }
}
</script>

<style>

</style>