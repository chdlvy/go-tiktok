<template>
	
	
	<view class="interact">
			<view class="subscribe">
				<view class="avatar">
					<image src="../../../static/logo.png"></image>
				</view>
				<image src="../../../static/interact/subscribe.png" class="subscribe-icon"></image>
			</view>
			<view class="like">
				<image src="../../../static/interact/like.png" @click="changeActive"  v-show="isActive"></image>
				<image src="../../../static/interact/like_active.png" @click="changeActive"  v-show="!isActive"></image>
				<view class="like-count count">111</view>
			</view>
			<view class="comment" @click="isShowComment">
				<image src="../../../static/interact/comment.png"></image>
				<view class="comment-count count">222</view>
			</view>
			<view class="collect">
				<image src="../../../static/interact/collect.png"></image>
				<view class="collect-count count">333</view>
			</view>
			<view class="share">
				<image src="../../../static/interact/share.png"></image>
				<view class="share-count count">444</view>
			</view>
	</view>
		<!-- <comment class="comment-mod" :showComment="showComment"></comment> -->
</template>

<script setup>
		import comment from "./childComps/comment/comment.vue"
		import {computed,ref} from "vue"
		
		const emit = defineEmits(["isShowComment"])
		let isActive = ref(true)
		function changeActive() {
			isActive.value  = !isActive.value
		}
		
		function isShowComment() {
			emit("isShowComment")
		}
</script>

<style scoped>
	.interact {
		position: fixed;
		right: 10px;
		bottom: 4rem;
		text-align: center;
	}
	.comment-mod {
		height: 50%;
	}
	.interact > view {
		margin: 1rem 0;
	}
	.subscribe {
		position: relative;
	}
	.avatar {
		border: 3px solid #fff;
		border-radius: 50%;
		overflow: hidden;
	}
	.subscribe-icon {
		width: 1.5rem;
		height: 1.5rem;
		position: absolute;
		left: 50%;
		transform: translateX(-50%);
		bottom: -.7rem;
		z-index: 1;
	}
	image {
		width: 35px;
		height: 35px;
	}
	
	.count {
		font-size: 14px;
		color: #fff;
	}
</style>