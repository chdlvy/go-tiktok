<template>
	<view class="page">
		<top-switch @change="change" v-show="!isShowTop"></top-switch>
		<recommend v-show="activeIndex===4"></recommend>
		<like v-show="activeIndex===1"></like>
		<nearby v-show="activeIndex===0"></nearby>
		<shop v-show="activeIndex===3"></shop>
		<explore v-show="activeIndex===2"></explore>
		
	</view>
</template>
<script setup>
	import {ref,provide} from "vue"
	import topSwitch from "../components/topSwitch/topSwitch"
	import recommend from "../views/recommend.vue"
	import like from "../views/like.vue"
	import nearby from "../views/nearby.vue"
	import shop from "../views/shop.vue"	
	import explore from "../views/explore.vue"
	
	// 控制头顶显示,打开评论区就隐藏头顶，关闭评论区就显示头顶
	let isShowTop = ref(false)
	function GrandsonGrandfatherChannel(value) {
		isShowTop.value = value
	}  
	// 发送函数给孙
	provide('sendHandle', GrandsonGrandfatherChannel)
	
	let activeIndex = ref(4)
	function change(index) {
		activeIndex.value = index
	}
</script>

<style scoped>
	.page {
		height: 100%;
	}
</style>