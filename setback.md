### 记录遇到的问题
#### swiper禁止滑动问题
描述：swiper有个孙组件comment(评论区)，当打开评论区时仍可以滑动切换视频
目的：打开评论区swiper禁止滑动(不能切换视频)，直至评论区关闭
过程：发现是冒泡事件的问题，当打开评论区并滑动的时候事件会冒泡到swiper-item，就会执行swiper-item的滑动切换视频的效果，随即为swiper-item监听`@touchmove.stop`发现可以实现无法滑动，但同时其子组件也会无法滑动，最重要的是.stop修饰符无法动态切换(一旦设定无法改变)
**解决方法**：为swiper-item组件监听`@touchmove="disSwiper"`事件，并设置flag保存评论区的开闭状态，当打开评论区时发送事件给swiper-item让它根据flag来动态调用原生的禁止冒泡方法
```vue
<template>
	<swiper vertical="true" class="swiper" >
		<swiper-item @touchmove="disSwiper">
			<swiper-itemcomp  @disTouch="disTouch"></swiper-itemcomp>
		</swiper-item>
		
	</swiper>
</template>

<script setup>
	import {ref,onMounted} from "vue"
	import swiperItemcomp from "@/pages/components/swiper/swiperItemcomp.vue"

	// 打开评论区之后禁止滑动swiper
	function disSwiper(e) {
		if(!canSwiper.value) {
			e.stopPropagation()
		}
	}
	定义一个flag
	let canSwiper = ref(true)
	function disTouch(value) {
		canSwiper.value = !value
	}
</script>
```