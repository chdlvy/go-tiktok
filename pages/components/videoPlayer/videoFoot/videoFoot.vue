<template>
	<view class="video-foot">
		<view class="left">
			<view class="writer">@作者名</view>
			<view class="describe">有趣啊有趣啊</view>
		</view>
		
		<view class="music-writer" :class="stopAnimation">
			<image src="../../../../static/logo.png"></image>
		</view>
		
	</view>
</template>

<script setup>
	import {computed} from "vue"
	const props = defineProps(['videoPause'])
	const stopAnimation = computed(()=> {
		return !props.videoPause ? "runAnimation": "stopAnimation"
	})
	
</script>

<style scoped>
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