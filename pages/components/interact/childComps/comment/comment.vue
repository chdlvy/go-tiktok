<template>
	<transition>
		<scroll-view scroll-y="true" class="comment" >
			<view class="comment-top">
				<text>100条评论</text>
				<image src="/static/interact/close.png"></image>
			</view>
				<comment-item></comment-item>
				<comment-item></comment-item>
				<comment-item></comment-item>
				<comment-item></comment-item>
				<comment-item></comment-item>
				<comment-item></comment-item>
				<comment-item></comment-item>
				<comment-item></comment-item>
				<comment-item></comment-item>
				<comment-item></comment-item>
				<comment-item></comment-item>
				<comment-item></comment-item>
		</scroll-view>
	</transition>
</template>

<script setup>
	import {ref} from 'vue'
	import commentItem from "./commentItem"
	
</script>

<style scoped>
	
	.comment {
		background-color: #fff;
		position: relative;
		z-index: 1;
		border-top-left-radius: 10px;
		border-top-right-radius: 10px;
		overflow: hidden;
	}
	.comment-top {
		position: sticky;
		top: 0;
		width: 100%;
		height: 30px;
		line-height: 30px;
		text-align:center;
		
	}
	comment-item {
		margin-top: 1rem;
	}
	.comment-top text {
		font-size: 14px;
		font-weight: bold;
	}
	.comment-top image {
		width: 15px;
		height: 15px;
		position: absolute;
		top: 50%;
		right: 1.5rem;
		transform: translateY(-50%);
		
	}
	
	
	.v-enter-active,
	.v-leave-active {
	  transition: height .5s;
	}
	.v-enter-from,
	.v-leave-to {
	  height: 0;
	}
</style>