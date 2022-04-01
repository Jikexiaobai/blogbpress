<template>
    <div class="container">
        <a-card :bordered="false">
          <a-form-model 
          ref="query" 
          :model="query" 
          :rules="query.rules" 
          layout="inline">
                <a-form-model-item ref="userId" prop="userId">
                    <a-input
                        size="large"
                        type="text"
                        placeholder="请输入报名用户ID"
                        v-model="query.userId"
                    />
                </a-form-model-item>
                <a-form-model-item>
                    <a-button v-action="'/edu/joinList'" @click="reset" style="margin-left: 10px;">
                        重置
                    </a-button>
                    <a-button v-action="'/edu/joinList'" @click="select" type="primary" style="margin-left: 10px;">
                        查询
                    </a-button>
                </a-form-model-item>
          </a-form-model>
          <div class="container-t">
            <a-table 
              :rowKey="record=>record.id"
              :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
              :pagination="{
                  showQuickJumper:true,
                  showSizeChanger:true,
                  showTotal:(total)=>`共${total}条`,
                  pageSize:query.limit,
                  current:query.page,
                  total:total
              }"
              :columns="columns" 
              :dataSource="list"
              @change="pageChange"
            >
              <span slot="mode" slot-scope="mode">
                <a-tag v-if="mode == 1" color="#2db7f5">
                    QQ
                </a-tag>
                <a-tag v-if="mode == 2" color="#87d068">
                    WeChat
                </a-tag>
                <a-tag  v-if="mode == 3" color="#f50">
                    手机号码
                </a-tag>
              </span>

            </a-table>

          </div>
        </a-card>
    </div>
</template>

<script>
import { UserJoinList } from '@/api/edu'
export default {
  components: {

  },
  data() {
    return {
      query:{
          page:1,
          limit:10,
          userId:undefined,
          eduId:undefined,
      },
      list:[],
      total:0,

      columns:[
          {
            title: "用户编号",
            dataIndex: 'id',
          },
          {
              title: "名称",
              dataIndex: 'name',
              scopedSlots: { customRender: 'name' }
          },
          {
              title: "昵称",
              dataIndex: 'nickName',
              scopedSlots: { customRender: 'nickName' }
          },
          {
              title: "联系方式",
              dataIndex: 'mode',
              scopedSlots: { customRender: 'mode' }
          },
          {
              title: "名联系号码",
              dataIndex: 'number',
              scopedSlots: { customRender: 'number' }
          },
          {
              title: "报名时间",
              dataIndex: 'createTime',
          },
      ],
      selectedRowKeys: [],
    }
  },
  mounted(){
    if (isNaN(this.$route.query.id)) {
        this.$router.push({path:"/404"})
    }
    this.query.eduId = parseInt(this.$route.query.id)
    
    this.getData()
  },
  methods: {
    async getData(){
        try {
            const res = await UserJoinList(this.query)
            this.list = res.data.list == null ? [] :  res.data.list.map((item)=>{
                const tmp = {
                    id:item.userId,
                    name:item.name,
                    nickName:item.nickName,
                    mode:item.mode,
                    number:item.number,
                    createTime:item.createTime,
                }
                return tmp
            })
            this.total = res.data.total
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
    onSelectChange(selectedRowKeys) {
        this.selectedRowKeys = selectedRowKeys;
    },
    pageChange(e){
        this.$emit("changePage",e)
    },
    reset(){
        this.query.userId = undefined
        this.getData()
    },
    select(){
        this.getData()
    }
  },
}
</script>

<style scoped lang="less">
.container-t{
  margin-top: 10px;
}
</style>
