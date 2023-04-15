<template>
	<view class="profile">	
		<view class="top">
			<view class="head">
				<image class="avatar" src="../../static/logo.png"></image>
				<view class="head-right">
					<view class="user-name">{{user.username}}</view>
					<view class="uid">抖音号：{{user.phone}}</view>
				</view>
			</view>
		</view>
		<card class="card"></card>
	</view>
</template>

<script setup>
	import {reactive} from "vue"
	import card from "@/pages/components/profile/card.vue"	
	let user =reactive(JSON.parse(localStorage.getItem("user")))
	
	
	
</script>

<style scoped>
	.profile {
		height: 100%;
		position: relative;
	}
	.top {
		position: relative;
		height: 30%;
		background-image: url("@/static/1.jpg");
	}
	.head {
		display: flex;
		position: absolute;
		bottom: 0;
		padding: 0 1rem 1.5rem 1rem;
	}
	
	.avatar {
		border-radius: 50%;
		width: 5rem;
		height: 5rem;
		margin-right: 1rem;
		border: 3px solid #fff;
	}
	.head-right {
		display: flex;
		flex-direction: column;
		justify-content: center;
		flex: 4;
	}
	.user-name {
		font-size: 20px;
		font-weight: bold;
		color: #fff;
		margin-bottom: 5px;
	}
	.uid {
		font-size: 14px;
		color: #000;
	}
	
	.card {
		height: 70%;
	}
</style>