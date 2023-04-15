<template>
		<view class="comment-item">
			<view class="left avatar"><image src="/static/logo.png"></image></view>
			<view class="right">
				<view class="user-name">{{commentItemMsg.userMsg.username}}</view>
				<view class="content">{{commentItemMsg.content}}</view>
				<view class="bottom">
					<view class="bottom-left">
						<text class="time">{{commentTime}} · 广东</text>
						<text class="reply">回复</text>
					</view>
					<!-- 点赞和踩 -->
					<view class="bottom-right">
						<view class="great" @click="greate">
							<image v-show="!greateActive" src="../../../../../static/interact/great.png"></image>
							<image v-show="greateActive" src="../../../../../static/interact/great_active.png"></image>
							<text>{{commentItemMsg.likeCount}}</text>
						</view>
						<view class="bad" @click="bad">
							<image v-show="!badActive" src="../../../../../static/interact/bad.png"></image>
							<image v-show="badActive" src="../../../../../static/interact/bad_active.png"></image>
						</view>
					</view>
				</view>
				
				
			</view>
		</view>
</template>

<script setup>
	import {ref} from "vue"
	import {getCommentTime} from "@/utils/utils.js"
	const props = defineProps(["commentItemMsg"])
	let commentItemMsg = ref(props.commentItemMsg)	
	let greateActive = ref(false)
	let badActive = ref(false)
	
	let commentTime = ref(getCommentTime(commentItemMsg.value.createAt))
	function greate() {
		greateActive.value = !greateActive.value
	}
	function bad() {
		badActive.value = !badActive.value
	}
</script>

<style scoped>
	.comment-item {
		background: #fff;
		display: flex;
		margin: .5rem 0;
	}
	.left {
		flex: 1;
		text-align: center;
	}
	.avatar image {
		border-radius: 50%;
		width: 40px;
		height: 40px;
	}
	.right {
		flex: 4;
	}
	
	.user-name,.time,.reply{
		font-size: 15px;
		color: #a6a6a6;
	}
	.content {
		margin: .2rem 2rem .2rem 0;
	}
	
	.bottom {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		
		padding-right: 1rem;
	}
	.reply {
		margin-left: 2rem;
		color: #0055ff;
	}
	.bottom-left {
		flex: 3;
	}
	.bottom-right {
		display: flex;
		flex: 1.5;
		justify-content:space-around;
		align-items: center;
	}
	.bottom-right image {
		width: 20px;
		height: 20px;
	}
	.great {
		display: flex;
	}
	.great text {
		font-size: 14px;
		color: #cdcdcd;
		margin-left: 2px;
	}
	
</style>