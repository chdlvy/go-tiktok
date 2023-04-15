<template>
		<view class="recommend">
			<swipercomp></swipercomp>
		</view>
</template>

<script setup>
	import {ref,onMounted,reactive,provide} from "vue"

	import swipercomp from "@/pages/components/swiper/swipercomp.vue"
	import {getVideos} from "@/network/home.js"
	
	// 从本地存储中获取id
	let user=JSON.parse(localStorage.getItem("user"))
	
	let videoList= ref([])
	// 将数据传入子组件和孙组件中，先发过去，后面异步再赋值即可同步数据
	provide("videoList",videoList)
	onMounted(()=> {
		getVideos(0,user.id).then(res=> {
			videoList.value = res.data.list
		})
	})
	
</script>

<style scoped>
	
	.recommend {
		height: 100%;
		position: relative;
		background-color: #000;
	}
	
</style>