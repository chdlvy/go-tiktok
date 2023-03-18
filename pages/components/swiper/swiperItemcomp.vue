<template>
	
	<view class="swiper-itemcomp">
		
		<view class="item">
			<video-player class="video" :style="{height:changeHeight}"></video-player>
			<comment v-show="showComment" class="comment"></comment>
		</view>
		<view class="mask" @click="mask" ></view>
		<interact v-show="!showComment"  @isShowComment="isShowComment" class="interact"></interact>
	</view>
</template>

<script setup>
	import {ref,inject,computed} from "vue"
	import videoPlayer from "@/pages/components/videoPlayer/videoPlayer.vue"
	import interact from "@/pages/components/interact/interact.vue"
	import comment from "@/pages/components/interact/childComps/comment/comment.vue"
	
	const emit = defineEmits(["disTouch"])
	// 动态改变视频高度，打开评论区高度变小
	let changeHeight = computed(()=> {
		return showComment.value ? "28%" : "100%"
	})
	
	// 拿到爷发送的函数，通过函数引用的特性在这里调用函数，实际上是执行爷的方法
	const GrandsonGrandfatherChannel = inject("sendHandle")
	let showComment = ref(false)
	function isShowComment() {
		// 打开评论
		showComment.value = true
		// 发送信息给父,禁用用户滑动
		emit("disTouch",true)
		
		// 发送信息给爷
		GrandsonGrandfatherChannel(showComment.value)
	}
	
	
	function mask() {
		// 点击暂停视频·····
			
		// 点击关闭评论
		showComment.value = false
		// 发送信息给父,开启用户滑动
		emit("disTouch",false)
		
		// 发送信息给爷
		GrandsonGrandfatherChannel(false)
	}
</script>

<style scoped>
	
	.swiper-itemcomp {
		height: 100%;
	}
	.item {
		height: 100%;
	}
	.video {
		height: 100%;
		transition: height .5s;
	}
	.comment {
		height: 72%;
	}
	
	.mask {
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
	}
</style>