<template>
    <div class="content-box">
        <div class="warper" >
            <div class="title">
                <h2>发布问题</h2>
                <!-- <a-divider /> -->
                <a-alert message="尝试在社区找找看看是否有答案呢？" banner />
            </div>
            <a-form-model ref="createForm" :model="createForm" :rules="createForm.rules">
                <a-form-model-item ref="title" prop="title">
                    <a-input :maxLength="40" class="create-title" size="large" v-model="createForm.title" placeholder="请输入问题标题">
                    </a-input>
                </a-form-model-item>

                <a-row :gutter="[40,0]">
                    <a-col :span="12">
                        <a-form-model-item ref="groupId" prop="groupId">
                            <a-select
                            show-search
                            @inputKeydown="searchGroup"
                            v-model="createForm.groupId" 
                            size="large" 
                            placeholder="请选择关注的圈子">
                                <a-select-option v-for="(item,index) in groupList" :key="index" :value="item.id">
                                {{item.title}}
                                </a-select-option>
                            </a-select>
                        </a-form-model-item>
                    </a-col>
                    <a-col :span="12">
                        <a-form-model-item class="article-cates-group" ref="cateId" prop="cateId">
                            <a-select

                                v-model="createForm.anonymous" 
                                size="large" 
                                placeholder="请选择是否匿名">
                                <a-select-option :value="ANONYMOUS.yes">
                                    是
                                </a-select-option>
                                <a-select-option :value="ANONYMOUS.no">
                                    否
                                </a-select-option>
                            </a-select>
                        </a-form-model-item>
                    </a-col>
                </a-row>

                <tinymceEditor 
                    toolbar="undo redo | imagelibrary  | h2Title hr bold italic forecolor backcolor  | bullist numlist | lists  | removeformat"
                    @writeContent="writeContent"   
                    v-model="createForm.content"/>
                    <a-button @click="onSubmit" class="post" type="primary">
                        发布
                    </a-button>
            </a-form-model>
        </div>
    </div>
</template>

<style lang="less" scoped>
.content-box{
    margin: 80px 0;
    display: flex;
    justify-content: center;
    min-height: 550px;
    .warper{
        padding: 20px;
        width: 900px;
        background: white;
        .title{
            font-size: 17px;
            margin-bottom: 10px;
            h2{
                margin-bottom: 10px;
            }
        }
        .post{
            margin-top: 10px;
        }
    }
}
</style>

<script>
import tinymceEditor from "@/components/editor/tinymceEditor"
import api from "@/api/index"

const ANONYMOUS = {
    yes:1,
    no:2
}
export default {
    components:{
        tinymceEditor
    },
    data(){
        return{
            ANONYMOUS,
            groupList:[],
            createForm:{
                title:undefined,
                content:null,
                groupId:undefined,
                anonymous:undefined,
                rules:{
                    title:[
                        { required: true, message: '请输入标题', trigger: 'change' },
                    ],
                    groupId:[
                        { required: true, message: '请设置圈子', trigger: 'change' },
                    ],
                },
            },
            queryParam:{
                page:1,
                limit: 7,
                module: "group",
                mode: 1,
                title: "",
            }
        }
    },
    mounted(){
        this.getData()
    },
    methods:{
        async getData(){
            try {
                const res = await this.$axios.get(api.getSystemFilter,{params: this.queryParam}) 
                if (res.code != 1) {
                    this.$router.push(`/404`)
                    this.$message.error(
                        res.message,
                        3
                    )
                    return
                }
                this.groupList = res.data.list != null ? res.data.list : []
            } catch (error) {
                setTimeout(() => {
                    this.$notification.error({
                        message: '网络错误',
                        description: "请稍后再试"
                    })
                }, 1000)
            }
        },
        searchGroup(e){
            
            if (e.keyCode == 13) {
                console.log(this.queryParam.title)
                // this.queryParam.title = this.createForm.
            }
        },
        writeContent(e){
            this.createForm.content = e
        },
        onSubmit(e){
            this.$confirm({
                okText:"确定",
                cancelText:"取消",
                title: '发布',
                content: '请注意，您填写的信息是否正确',
                onOk:() => {
                    this.$refs.createForm.validate(valid => {
                        if (valid) {

                            let formData = {
                                content:null,
                                title:null,
                                groupId:null,
                                anonymous:null,
                            }

                            formData = Object.assign(formData,this.createForm)
                            
                            this.postCreate(formData)
                        } else {
                            return false;
                        }
                    });
                },
                onCancel() {},
            });
        },
        async postCreate(formData){
            try {
                const res = await this.$axios.post(api.postQuestionCreate,formData)
                if (res.code != 1) {
                    this.$message.error(
                        res.message,
                        3
                    )
                    return
                }
                this.$router.push({ path: "/question"})
            } catch (error) {
                setTimeout(() => {
                    this.$notification.error({
                        message: '网络错误',
                        description: "请稍后再试"
                    })
                }, 1000)
            }
        },
    }
}
</script>