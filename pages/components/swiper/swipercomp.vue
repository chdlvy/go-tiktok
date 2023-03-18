<template>
	<swiper vertical="true" class="swiper" @disTouch="disTouch">
		<swiper-item>
			<swiper-itemcomp></swiper-itemcomp>
		</swiper-item>
		
	</swiper>
</template>

<script setup>
	import {ref} from "vue"
	import swiperItemcomp from "@/pages/components/swiper/swiperItemcomp.vue"
	
	
	let isTouch = ref(true)
	function disTouch(value) {
		isTouch.value = value
	}
</script>

<style scoped>
	.swiper {
		height: 100%;
	}
	
</style>