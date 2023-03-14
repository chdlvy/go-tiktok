<template>
	<view class="page">
		<top-switch @change="change"></top-switch>
		<recommend v-show="activeIndex===4"></recommend>
		<like v-show="activeIndex===1"></like>
		<nearby v-show="activeIndex===0"></nearby>
		<shop v-show="activeIndex===3"></shop>
		<explore v-show="activeIndex===2"></explore>
	</view>
</template>
<script setup>
	import {ref} from "vue"
	import topSwitch from "../components/topSwitch/topSwitch"
	import recommend from "../views/recommend.vue"
	import like from "../views/like.vue"
	import nearby from "../views/nearby.vue"
	import shop from "../views/shop.vue"
	import explore from "../views/explore.vue"
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