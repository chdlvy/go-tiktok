<template>
	<view class="card-tabbar">
		<view class="tabbar">
			<view 
			@click="isActive(index)"
			:class="index===curActive?'active':''" 
			v-for="(item,index) in text" 
			:key="index"
			>{{item}}</view>
		</view>
		
		<view class="category">
			<view class="product" v-show="curActive===0">
				<text>没有照片</text>
			</view>
			<view class="private" v-show="curActive===1">
				<text>没有私密</text>
			</view>
			<view class="collect" v-show="curActive===2">
				<text>没有收藏</text>
			</view>
			<view class="like" v-show="curActive===3">
				<like ref="isLike"></like>
				
			</view>
			
		</view>
	</view>
</template>

<script setup>
	import {reactive,ref} from "vue"
	import like from "./like.vue"
	let text = reactive(["作品","私密","收藏","喜欢"])
	let curActive = ref(0)
	function isActive(index) {
		curActive.value = index
		if(curActive.value==3) {
			isLike.value.getList()
		}
	}
	let isLike = ref()
	function showLikeList() {
		// this.$refs.isLike.getList()
		console.log(isLike.value);
		
	}
</script>

<style scoped>
	.tabbar {
		display: flex;
	}
	.tabbar view {
		font-size: 16px;
		color: gray;
		flex: 1;
		text-align: center;
		padding-bottom: 7px;
		padding-left: 0;
	}
	.active {
		border-bottom: 1px #000 solid;
		font-weight: bold;
		color: #000;
	}
	
	
</style>