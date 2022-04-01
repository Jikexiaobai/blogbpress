<template>
   <div class="wave-player-box">
        <div ref="wave" class="wave-player"/>
        <a-icon v-if="!playActive" @click="play" theme="filled" class="play" type="play-circle" />
        <a-icon v-if="playActive" @click="pause" theme="filled" class="play" type="pause-circle" />
    </div>
</template>

<script>
if (process.client) {
    var WaveSurfer = require('wavesurfer.js')
}
export default {
    props:{ 
        link:{
            type: String, //指定传入的类型
            default: "/audio/toymachine.mp3" //这样可以指定默认的值
        },
    },
    data(){
        return{
            wavesurfer: '',
            playActive:false,
            // tmpwaveImage: null,	
            // isCreatingWave: false,	
        }
    },
    mounted(){
        if (process.client) {
            this.wavesurfer = WaveSurfer.create({
                container: this.$refs.wave,//绑定容器，第一种方法
                forceDecode: true,   
                waveColor: '#A8DBA8',
                progressColor: '#3B8686',
                cursorWidth:0,
                partialRender:true,
                barRadius:2,
                barHeight:1,
                barWidth:3,
                height:64,
                // xhr:  { 
                //     cache: 'default',
                //     mode: 'cors',
                //     method: 'GET', 
                //     credentials: 'same-origin', 
                //     redirect: 'follow', 
                //     referrer: 'client', 
                //     headers: [ { key: 'Authorization', value: 'my-token' },{ key: 'Access-Control-Allow-Origin', value: '*' } ]
                // }
            })
            this.wavesurfer.load(this.link)
            // this.wavesurfer.on('loading', () =>{
            //     this.isCreatingWave = true
            //     // setTimeout(() =>{
                    
            //     // }, 2000);
            // });
            // this.wavesurfer.on('ready', () =>{
            //     setTimeout(() =>{
            //         this.isCreatingWave = false
            //     }, 2000);
            // });
            //  this.isCreateWave = false
        }
    },
    destroyed(){
        this.wavesurfer.destroy()
    },
    methods:{
        play(){
              // this.tmpwaveImage = this.wavesurfer.exportImage("image/png")
                // console.log(this.tmpwaveImage)
            this.wavesurfer.play()
            this.playActive = true
        },
        pause(){
            this.wavesurfer.pause()
            this.playActive = false
        }
    }
}
</script>

<style lang="less" scoped>
.wave-player-box{
    display: flex;
    justify-content: space-between;
    align-items: center;
    .wave-player{
        flex: 1;
        margin-right: 20px;
    }
    /deep/ .ant-btn-link{
        color: #8590a6;
        padding: 0 5px;
    }
    .play{
        font-size: 50px;
    }
}
</style>