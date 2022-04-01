<template>
    <div class="select-box">
        <a-form-model ref="selectForm" :model="selectForm" :rules="selectForm.rules" layout="inline">
            <a-form-model-item ref="title" prop="title">
                <a-input
                    size="large"
                    type="text"
                    placeholder="标题"
                    v-model="selectForm.title"
                >
                </a-input>
            </a-form-model-item>
            <a-form-model-item ref="cateId" prop="cateId">
                <a-tree-select 
                        :style="{width:200 + 'px'}"
                        v-model="selectForm.cateId"
                        :tree-data="cateList"
                        size="large" 
                        placeholder="请选择分类"
                        tree-default-expand-all
                    />
            </a-form-model-item>
            <a-form-model-item ref="status" prop="status">
               <a-select v-model="selectForm.status" 
                    :style="{width:200 + 'px'}"
                    size="large"
                    placeholder="请选择状态">
                    <a-select-option  :value="1">
                        未审核
                    </a-select-option>
                    <a-select-option  :value="2">
                        已通过
                    </a-select-option>
                    <a-select-option  :value="3">
                        未通过
                    </a-select-option>
                    <a-select-option  :value="4">
                        回收站
                    </a-select-option>
                </a-select>
            </a-form-model-item>
            <a-form-model-item>
                <a-button v-action="'/video/list'" @click="reset"  size="large" style="margin-left: 10px;">
                    重置
                </a-button>
                <a-button v-action="'/video/list'" @click="select"  size="large" type="primary" style="margin-left: 10px;">
                    查询
                </a-button>
            </a-form-model-item>
        </a-form-model>
    </div>
</template>

<script>
import { ByModule } from '@/api/category'
import { loopCate } from '@/utils/tree/cate'
import { toTree } from '@/utils/tree/tree'
export default {
    data(){
        return{
            cateList:[],
            selectForm:{
                status:undefined,
                cateId:undefined,
                title:undefined,
            }
        }
    },
    mounted(){
        this.getData()
    },
    methods:{
        async getData(){
            try {
                const res = await ByModule({"module":"video"})
                const list =  toTree(res.data.list != null ? res.data.list : [],"cateId","parentId")
                this.cateList = loopCate(list)
            } catch (error) {
                console.log(error)
                setTimeout(() => {
                    this.$notification.error({
                        message: '网络错误',
                        description: "请稍后再试"
                    })
                }, 1000)
            }
        },
        reset(){
            this.selectForm.status = undefined
            this.selectForm.cateId = undefined
            this.selectForm.title = undefined
            this.$emit("select",this.selectForm)
        },
        select(){
            this.$emit("select",this.selectForm)
        }
    }
}
</script>

<style lang="less" scoped>
.select-box{
    margin-bottom: 10px;
}
</style>