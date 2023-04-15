<template>
	<view class="like">
		<view class="img-item" v-for="item in likeList">
			<view class="img-box">
				<image src="../../../../static/1.jpg"></image>
				<view class="show-count">
					<image src="../../../../static/interact/great.png"></image>
					<text>{{item.likeCount}}</text>
				</view>
			</view>
		</view>
	</view>
</template>

<script setup>
	import {ref} from "vue"
	import {getLikeList} from "@/network/profile.js"
	defineExpose({
		getList
	})
	let likeList = ref([])
	let user = JSON.parse(localStorage.getItem("user"))
	function getList() {
		getLikeList(user.id).then(res=> {
			likeList.value = res.data.likeList
		})
	}
	
	
	
</script>

<style scoped>
	.like {
		width: 100%;
		height: 100%;
		display: flex;
		flex-wrap: wrap;
	}
	.img-item{
		flex-basis: calc(33.33% - 1px);
		height: 200px;
	}
	.img-item:nth-child(3n-1) {
		padding: 0 1px;
	}
	.img-item:nth-child(n+4) {
		margin-top: -2px;
	}
	.img-box {
		position: relative;
		height: 100%;
	}
	.img-item image {
		width: 100%;
		height: 100%;
	}
	.show-count {
		position: absolute;
		bottom: 10px;
		left: 10px;
		color: #fff;
	}
	.show-count text {
		margin-left: 3px;
	}
	.show-count image {
		margin-bottom:2px;
		width: 15px;
		height:15px;
		vertical-align: bottom;
	}
</style>