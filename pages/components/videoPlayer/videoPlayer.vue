<template>
  <view class="video-player">
	  <video :id="videoDetail.id" class="video"
	  :src="videoSrc" 
	  controls="false" 
	  autoplay="true" 
	  muted="false"
	  loop="true"
	  @timeupdate="timeUpdate"
	  @loadedmetadata="loadedmetadata"
	  ></video>
	  
	  <view class="video-foot">
	  	<view class="left">
	  		<view class="writer">@{{videoDetail.writerName}}</view>
	  		<view class="describe">{{videoDetail.description}}</view>
	  	</view>
	  	<view class="music-writer" :class="stopAnimation">
	  		<image src="../../../../static/logo.png"></image>
	  	</view>
	  </view>
	  
	  <view class="slider-view" v-show="isPause">
		  <slider 
			  v-show="showComment"
			  min="0" 
			  :max="duration"
			  step="0.01"
			  backgroundColor="gray" 
			  activeColor="#fff" 
			  block-size="14"
			  @changing="sliderChanging"
			  @change="sliderChange"
			  :value="currentTime"
		  ></slider>
	  </view>
	  <view class="mask-control" @click="controlVideo"></view>
  </view>
</template>

<script setup>
import {ref,reactive,onMounted,computed} from "vue"
	
const props = defineProps(['showComment',"videoDetail"])
	let videoDetail = ref(props.videoDetail)
	
	let videoSrc = ref("")
		videoSrc.value="/static/video.mp4"
	let videoContext
	// 当前时长
	let currentTime = ref(0)
	// 视频时长
	let duration = ref(0)
	let isPause = ref(false)
	
	// 是否拖拽
	let isDrag = ref(false)
	onMounted(()=> {
		videoContext = reactive(uni.createVideoContext(videoDetail.value.id,this))
	})
	// 点击暂停右下角图片旋转
	const stopAnimation = computed(()=> {
		return !isPause.value ? "runAnimation": "stopAnimation"
	})
	
	function timeUpdate(e) {
		if(!isDrag.value) {
			//获取当前时间
			currentTime.value = e.detail.currentTime
		}
	}
	
	// 视频加载完成
	function loadedmetadata(e) {
		duration.value = e.detail.duration
	}
	//拖动进度条时
	function sliderChanging(value) {
		isDrag.value=true
	}
	// 拖拽进度条结束
	function sliderChange(value) {
		isDrag.value = false
		console.log(value);
		 // 跳到指定时间点
		videoContext.seek(value.detail.value)
		videoContext.play()
	}
	// 通过遮罩控制视频播放暂停
	function controlVideo() {
		isPause.value=!isPause.value
		if(isPause.value) {
			videoContext.pause()
		}else {
			videoContext.play()
		}
	}
</script>

<style scoped>
.video-player {
		height: 100%;
		width: 100%;
		position: relative;
	}
.video {
		max-width: 100%;
		max-height: 100%;
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%,-50%);
	}
.video-foot {
	position: fixed;
	bottom: 10px;
	height: 10%;
	width: 100%;
}
.slider-view {
	position: absolute;
	bottom: -5px;
	width: 100%;
	z-index: 1;
}
.mask-control {
	position: absolute;
	top: 0;
	bottom: 0;
	left: 0;
	right: 0;
}

.video-foot {
	display: flex;
	justify-content: space-between;
	align-items: center;
	color: #fff;
	padding: 0 1rem;
	box-sizing: border-box;
	
}
.left {
	display: flex;
	flex-direction: column;
}
.writer {
	margin-bottom: 5px;
	font-size: 20px
}
.music-writer {
	width: 37px;
	height: 37px;
	border-radius: 50%;
	overflow: hidden;
	animation: avatarRotate 3s linear infinite;
	align-self: flex-end;
	margin-bottom: 15px;
}
.runAnimation {
	animation-play-state: running;
}
.stopAnimation {
	animation-play-state: paused;
}
.music-writer image{
	width: 100%;
	height: 100%;
}

@keyframes avatarRotate {
	from {
		transform: rotate(0);
	}
	to {
		transform: rotate(360deg);
	}
}
	
</style>
