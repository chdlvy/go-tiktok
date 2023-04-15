<template>
	<view class="card">
		<view class="about">
			<text><text class="about-num">0</text>获赞</text>
			<text><text class="about-num">2</text>朋友</text>
			<text><text class="about-num">5</text>关注</text>
			<text><text class="about-num">7</text>粉丝</text>
		</view>
		
		<view class="introduce">
			<text>点击添加介绍，让大家认识你...</text>
			<ul class="tag">
				<li>21岁</li>
				<li>广东·珠海</li>
				<li>+ 添加学校等标签</li>
			</ul>
		</view>
		
		<view class="shop">
			<view class="shop-icon">
				<image src="../../../static/interact/collect.png" mode=""></image>
			</view>
			<view class="shop-right">
				<view class="shop-title">抖音商城</view>
				<view class="shop-ad">美好生活触手可得</view>
			</view>
		</view>
		
		<view class="card-foot">
			<view class="edit btn">
				编辑资料 70%
			</view>
			<view class="add-friend btn">
				添加朋友
			</view>
		</view>
		
		<card-tabbar class="card-tabbar"></card-tabbar>
		
	</view>
</template>

<script setup>
	import cardTabbar from "./childComp/cardTabbar.vue"
</script>

<style scoped>
	.card {
		position: relative;
	}
	.card::before {
		content: "";
		position: absolute;
		top: -10px;
		left: 0;
		width: 100%;
		height: 20px;
		background-color:#fff;
		border-radius: 20px;
	}
	
	.about {
		padding-top: 10px;
		width: 60%;
		display: flex;
		padding-left: 20px;
	}
	.about text {
		flex: 1;
	}
	.about-num {
		font-weight: 600;
		font-size: 18px;
		margin-right: 4px;
	}
	
	.introduce {
		margin-top: 10px;
		padding-left: 20px;
	}
	.introduce>text {
		font-size: 14px;
		color: gray;
	}
	.tag {
		margin-top: 5px;
	}
	.tag>li {
		display: inline-block;
		font-size: 14px;
		margin-right: 5px;
		background-color: #ebebeb;
		color: gray;
		padding:2px 5px;
		border-radius: 5px;
	}
	
	.shop {
		display: flex;
		margin-top: 1.5rem;
		padding-left: 20px;
	}
	.shop-icon {
		width: 30px;
		height: 30px;
		background: gray;
		padding: .6rem;
		margin-right: 10px;
	}
	
	.shop-icon image {
		height: 100%;
		width: 100%;
	}
	.shop-title {
		margin-bottom: .6rem;
	}
	.shop-ad {
		color: gray;
		font-size: 14px;
	}
	
	.card-foot {
		display: flex;
		padding-right: 20px;
		margin-top: 1.5rem;
		padding-left: 20px;
	}
	.btn {
		background: #ebebeb;
		padding: .5rem 1rem;
		border-radius: 5px;
		text-align: center;
	}
	.edit {
		flex: 1;
		margin-right: 5px;
	}
	.add-friend {
		flex: 1;
	}
	.card-tabbar {
		margin-top: 1.5rem;
	}
</style>