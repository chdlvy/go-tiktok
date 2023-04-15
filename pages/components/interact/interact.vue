<template>
	
	
	<view class="interact">
			<view class="subscribe">
				<view class="avatar">
					<image src="../../../static/logo.png"></image>
				</view>
				<image src="../../../static/interact/subscribe.png" class="subscribe-icon"></image>
			</view>
			<view class="like">
				<image src="../../../static/interact/like.png" @click="likeVideo"  v-show="isLike"></image>
				<image src="../../../static/interact/like_active.png" @click="cancelLikeVideo"  v-show="!isLike"></image>
				<view class="like-count count">{{videoDetail.likeCount}}</view>
			</view>
			<view class="comment" @click="isShowComment">
				<image src="../../../static/interact/comment.png"></image>
				<view class="comment-count count">{{videoDetail.commentCount}}</view>
			</view>
			<view class="collect">
				<image src="../../../static/interact/collect.png"></image>
				<view class="collect-count count">{{videoDetail.collectCount}}</view>
			</view>
			<view class="share">
				<image src="../../../static/interact/share.png"></image>
				<view class="share-count count">{{videoDetail.shareCount}}</view>
			</view>
	</view>
		
</template>

<script setup>
		import {computed,ref,onMounted} from "vue"
		import {likeThisVideo,cancelLikeThisVideo} from "@/network/home.js"
		const props = defineProps(["videoDetail"])
		const emit = defineEmits(["isShowComment"])
		let isLike = ref(true)
		let videoDetail = ref(props.videoDetail)
		onMounted(()=> {
			if(videoDetail.value.isLike==1) {
				isLike.value=false
			}
		})
		
		// 从本地存储中获取id
		let user=JSON.parse(localStorage.getItem("user"))
		
		function likeVideo() {
			isLike.value  = false
			likeThisVideo(user.id,videoDetail.value.id,videoDetail.value.writerId).then(res=>{
				
				videoDetail.value.likeCount=res.data.likeCount
			}).catch(err=> {
				isLike.value  = true
			})
		}
		function cancelLikeVideo() {
			// isLike.value  = true
			cancelLikeThisVideo(user.id,videoDetail.value.id,videoDetail.value.writerId).then(res=>{
				isLike.value  = true
				videoDetail.value.likeCount=res.data.likeCount
			})
		}
		
		function isShowComment() {
			emit("isShowComment")
		}
</script>

<style scoped>
	.interact {
		position: fixed;
		right: .7rem;
		bottom: 4rem;
		text-align: center;
	}
	.comment-mod {
		height: 50%;
	}
	.like,.comment,.collect,.share {
		margin: 1rem 0;
	}
	.subscribe {
		position: relative;
		margin-bottom: 30px;
	}
	.avatar {
		border: 3px solid #fff;
		border-radius: 50%;
		overflow: hidden;
		width: 35px;
		height: 35px;
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