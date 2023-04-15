<template>
	<transition>
		<view class="comment" >
				<view class="comment-top">
					<text>{{commentDetail.count}}条评论</text>
					<image src="/static/interact/close.png" @click="closeComment"></image>
				</view>
				<scroll-view scroll-y="true" style="height: 100%;">
					<comment-item 
					v-for="(item,index) in commentDetail.list" 
					:key="index"
					:commentItemMsg="item"
					></comment-item>
				</scroll-view>
		</view>
	</transition>
</template>

<script setup>
	import {ref,onMounted,reactive} from 'vue'
	import commentItem from "./commentItem"
	import {getComment} from "@/network/home.js"
	const props = defineProps(["commentDetail","videoId"])
	const emit = defineEmits(['closeComment'])
	
	let commentDetail = reactive({
		count:0,
		list:[]
	})
	
	onMounted(()=> {
		// 发送请求获取评论信息
		getComment(0,props.videoId).then(res=>{
			commentDetail.count=res.data.count
			commentDetail.list= res.data.list
		})
	})
	
	function closeComment() {
		emit('closeComment')
	}
	
</script>

<style scoped>
	.comment {
		background-color: #fff;
		/* position: absolute;
		bottom: 0; */
		position: relative;
		z-index: 1;
		border-top-left-radius: 15px;
		border-top-right-radius: 15px;
		overflow: hidden;
	}
	.comment-top {
		position: sticky;
		top: 0;
		width: 100%;
		height: 30px;
		line-height: 30px;
		text-align:center;
		background-color: #fff;
		
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