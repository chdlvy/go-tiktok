<template>
  <view class="video-player">
	  <video id="video" :src="videoSrc" controls="false" muted="true" autoplay="true" loop="true"></video>
  </view>
</template>

<script setup>
import {ref,reactive,nextTick} from "vue"
let videoSrc = ref("")
	let videoContext = reactive(uni.createVideoContext("video",this))

		videoSrc.value="/static/video.mp4"
		videoContext.play()

</script>

<style>
.video-player {
		height: 100%;
		width: 100%;
		position: relative;
	}
#video {
		max-width: 100%;
		max-height: 100%;
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%,-50%);
		
	}
</style>
