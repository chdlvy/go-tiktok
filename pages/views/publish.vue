<template>
	<view class="publish">
		<view class="back">返回编辑</view>
		<view class="top">
			<textarea class="desc" placeholder="添加作品描述"></textarea>
			<view class="choose">
				<image class="cover" src="../../static/1.jpg"></image>
				<view>选封面</view>
			</view>
		</view>
		<view class="foot">
			<button class="save">
				存草稿
			</button>
			<button class="publish-btn">
				发布
			</button>
		</view>
	</view>
</template>

<script setup>
</script>

<style scoped>
	.top {
		display: flex;
		padding: 1rem 1rem 0 1rem;
	}
	.desc {
		flex: 2;
	}
	.choose {
		flex: 1;
		position: relative;
		border-radius: 5px;
		overflow: hidden;
	}
	.cover {
		width: 100%;
		height: 100%;
	}
	.choose>view {
		font-size: 14px;
		color: #fff;
		position: absolute;
		bottom: 0;
		left: 0;
		right: 0;
		height: 20px;
		line-height: 20px;
		background: rgba(0, 0, 0, .5);
		text-align: center;
		color: aliceblue;
	}
	
	.foot {
		display: flex;
		position: absolute;
		bottom: 10px;
		width: 100%;
		font-size: 14px
	}
	.save {
		background-color: aliceblue;
		color: #000;
		border-radius: 5px;
		width: 100%;
		margin-right: 10px;
		margin-left: .5rem;
	}
	.publish-btn {
		background-color: red;
		color: #fff;
		border-radius: 5px;
		width: 100%;
		margin-right: .5rem;
	}
</style>