<template>
	<view class="top-switch">
			<text :class="activeIndex === index?'top-words active' : 'top-words'" v-for="(item,index) in topArr" :key="index" @click="changeActive(index)">{{item}}</text>
	</view>
</template>

<script setup>
	import {reactive,computed,ref} from "vue"
	const emit = defineEmits(['change'])
	
	let topArr = reactive(["广州","关注","探索","商城","推荐"])
	let activeIndex = ref(4)
	function changeActive(index) {
		activeIndex.value = index
		emit('change',index)
	}
		
		
</script>

<style scoped>
	.top-switch {
		position: absolute;
		top: 15px;
		z-index: 1;
		width: 100%;
		display: flex;
		justify-content: space-around;
	}
	.top-words {
		font-family:"微软雅黑";
		color: #d5d5d5;
		position: relative;
	}
	.active {
		color: #fff;
	}
	.active::after {
		content: "";
		position: absolute;
		bottom: -5px;
		left: 5px;
		width: 22px;
		height: 2px;
		background-color: #fff;
	}
</style>