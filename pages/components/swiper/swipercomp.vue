<template>
	<swiper vertical="true" class="swiper" >
		<swiper-item @touchmove="disSwiper" v-for="(item,index) in videoList" :key="item.id">
			<swiper-itemcomp :video-detail="item"  @disTouch="disTouch"></swiper-itemcomp>
		</swiper-item>
	</swiper>
</template>

<script setup>
	import {ref,onMounted,inject,watch } from "vue"
	import swiperItemcomp from "@/pages/components/swiper/swiperItemcomp.vue"
	let videoList = inject("videoList")

	// 打开评论区之后禁止滑动swiper
	function disSwiper(e) {
		if(!canSwiper.value) {
			e.stopPropagation()
		}
	}
	
	let canSwiper = ref(true)
	function disTouch(value) {
		canSwiper.value = !value
		console.log(canSwiper.value);
	}
</script>

<style scoped>
	.swiper {
		height: 100%;
		
	}
	
</style>