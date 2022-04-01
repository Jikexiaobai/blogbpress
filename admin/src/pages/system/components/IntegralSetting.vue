<template>
    <div class="setting">
        <div class="setting-item">
            <div class="setting-input">
                <span>新用户注册</span>
                <a-input-number class="setting-content"  
                v-model="registerIntegral" 
                  :min="0" :max="9999" :precision="0"/>
            </div>
            <p>奖励对象：注册者。首次注册时奖励。</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>签到</span>
                <a-input class="setting-content"  
                v-model="signInIntegral" 
                />
            </div>
            <p>奖励对象：当前登录用户，登录用户每日签到将会随机获得积分，<br/>如果是固定值请使用 xx-xx 例如 100-100（中间横杠不可缺失）</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>发布内容</span>
                <a-input-number class="setting-content"  
                v-model="contentIntegral" 
                  :min="0" :max="9999" :precision="0"/>
                  <a-input-number class="setting-content"  
                v-model="contentCount" 
                  :min="0" :max="9999" :precision="0"/>
            </div>
            <p>左边输入奖励积分，右边输入当日奖励次数 <br/>奖励对象：内容作者。<br/>内容发布包括：文章，视频，音频，资源，课程，话题,问题,圈子</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>回答问题</span>
                <a-input-number class="setting-content"  
                v-model="answerIntegral" 
                  :min="0" :max="9999" :precision="0"/>
                  <a-input-number class="setting-content"  
                v-model="answerCount" 
                  :min="0" :max="9999" :precision="0"/>
            </div>
            <p>左边输入奖励积分，右边输入当日奖励次数 <br/>奖励对象：答案作者。</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>发布评论</span>
                <a-input-number class="setting-content"  
                v-model="commentIntegral" 
                  :min="0" :max="9999" :precision="0"/>
                  <a-input-number class="setting-content"  
                v-model="commentCount" 
                  :min="0" :max="9999" :precision="0"/>
            </div>
            <p>左边输入奖励积分，右边输入当日奖励次数 <br/>奖励对象：评论作者。</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>点赞收藏</span>
                <a-input-number class="setting-content"  
                v-model="likefavoriteIntegral" 
                  :min="0" :max="9999" :precision="0"/>
                  <a-input-number class="setting-content"  
                v-model="likefavoriteCount" 
                  :min="0" :max="9999" :precision="0"/>
            </div>
            <p>左边输入奖励积分，右边输入当日奖励次数 <br/>奖励对象：被点赞收藏内容作者。</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>关注或被关注</span>
                <a-input-number class="setting-content"  
                v-model="followIntegral" 
                  :min="0" :max="9999" :precision="0"/>
                  <a-input-number class="setting-content"  
                v-model="followCount" 
                  :min="0" :max="9999" :precision="0"/>
            </div>
            <p>左边输入奖励积分，右边输入当日奖励次数 <br/>点击关注的人或被关注的人<br/>取消关注会扣除相应的积分。</p>
        </div>
        <div class="setting-item">
            <div class="setting-input">
                <span>举报内容</span>
                <a-input-number class="setting-content"  
                v-model="reportIntegral" 
                  :min="0" :max="9999" :precision="0"/>
                  <a-input-number class="setting-content"  
                v-model="reportCount" 
                  :min="0" :max="9999" :precision="0"/>
            </div>
            <p>左边输入奖励积分，右边输入当日奖励次数 <br/>奖励对象：举报成功删除内容的举报人。</p>
        </div>
        <div v-action="'/system/save'" class="setting-item">
            <a-button @click="handleSubmit" type="primary" class="setting-save">
                保存
            </a-button>
        </div>
    </div>
</template>
<script>
import { Save,Info } from '@/api/system'
export default {
    data(){
        return{
            registerIntegral:260,
            signInIntegral:"50-90",
            contentIntegral:50,
            contentCount:2,
            groupIntegral:50,
            groupCount:2,
            answerIntegral:50,
            answerCount:2,
            commentIntegral:50,
            commentCount:2,
            likefavoriteIntegral:50,
            likefavoriteCount:2,
            followIntegral:50,
            followCount:2,
            reportIntegral:50,
            reportCount:2,
        }
    },
    created(){
        this.getData()
    },
    methods:{
        
        async getData(){
            const res = await Info("IntegralSetting")  
            if (res.data.info != null) {
                const value = JSON.parse(res.data.info)
                // Object.assign(this,value)
                this.registerIntegral = value.registerIntegral
                this.signInIntegral = value.signInIntegral
                this.contentIntegral = value.contentIntegral
                this.contentCount = value.contentCount
                this.groupIntegral = value.groupIntegral
                this.groupCount = value.groupCount
                this.answerIntegral = value.answerIntegral
                this.answerCount = value.answerCount
                this.commentIntegral = value.commentIntegral
                this.commentCount = value.commentCount
                this.likefavoriteIntegral = value.likefavoriteIntegral
                this.likefavoriteCount = value.likefavoriteCount
                this.followIntegral = value.followIntegral
                this.followCount = value.followCount
                this.reportIntegral = value.reportIntegral
                this.reportCount = value.reportCount
                
            }

        },
        handleSubmit(){
            this.$confirm({
                content: this.$t('confirm.content.text'),
                onOk:() => {
                    const options = {
                        registerIntegral:this.registerIntegral,
                        signInIntegral:this.signInIntegral,
                        contentIntegral:this.contentIntegral,
                        contentCount:this.contentCount,
                        groupIntegral:this.groupIntegral,
                        groupCount:this.groupCount,
                        answerIntegral:this.answerIntegral,
                        answerCount:this.answerCount,
                        commentIntegral:this.commentIntegral,
                        commentCount:this.commentCount,
                        likefavoriteIntegral:this.likefavoriteIntegral,
                        likefavoriteCount:this.likefavoriteCount,
                        followIntegral:this.followIntegral,
                        followCount:this.followCount,
                        reportIntegral:this.reportIntegral,
                        reportCount:this.reportCount,
                    }
                    const config = {
                        "configValue":JSON.stringify(options),
                        "configKey": "IntegralSetting",
                        "configName": "积分设置",
                        "remark" : "积分设置"
                    }
                    this.CreatePost(config)
                },
                onCancel() {},
            });
        },
        async CreatePost(value){
            try {
                const data = await Save(value)
                if (data.code == 1) {
                    this.$message.success(
                        data.message,
                        3
                    )
                } else {
                    this.$message.error(
                        data.message,
                        3
                    )
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
    }
}
</script>

<style lang="less" scoped>
.setting{
    .setting-item{
        width: 520px;
        margin-top: 20px;
        .setting-input{
            display: flex;
            justify-content: space-between;
            align-items: center;
            span{
                width: 200px;
                font-size: 14px;
                font-weight: 700;
            }
            .setting-content{
                margin-left: 20px;
                width: 400px;
            }
        }
        .setting-upload{
            display: flex;
            justify-content: flex-end;
        }
        p{
            text-align: right;
            margin-top: 10px;
            font-size: 12px;
        }
        .setting-save{
            margin-right: 20px;
        }
    }
    
}
</style>
