<template>
	<view class="login">
		<view class="head">手机号密码登录</view>
		<form>
			<view class="account">
				<text class="phone-head">+86</text>
				<label>
					<input type="text" :value="phone" id="phone" focus @input="inputPhone" maxlength="13"  />
				</label>
			</view>
			<view class="psw">
				<label>
					<input id="password" type="password" placeholder="请输入密码" v-model="password" />
				</label>
			</view>
			<text class="err" v-show="errMsg.length!=0">{{errMsg}}</text>
			<view class="protocol">
				<label @click="radioClick">
					<radio :checked="isRead"/><text>已阅读并同意<text class="highlight">用户协议</text>和<text class="highlight">隐私政策</text></text>
				</label>
			</view>
			<view class="tips">未注册的手机号验证通过后将自动注册</view>
			<button class="btn" @click="login_btn">
				<view class="mask" v-show="phone.length==0"></view>
				<text>登录</text>
			</button>
			<view class="other">
				<view class="forget">忘记了？<text class="highlight">找回密码</text></view>
				<view class="email-login"><text class="highlight">邮箱密码登录</text></view>
			</view>
		</form>
	</view>
	
</template>
f
<script setup>
	import {ref} from "vue"
	import {login} from "@/network/login.js"
	let isRead = ref(false)
	let phone = ref("")
	let password = ref("")
	let errMsg = ref("")
	
	function login_btn() {
		let newphone = phone.value.replace(/\s*/g,"");
		if(!isRead.value) {
			errMsg.value = "请同意用户协议"
			return
		}
		if(newphone.length<11) {
			errMsg.value = "手机号填写错误"
			return
		}
		errMsg.value=""
		login(newphone,password.value).then(res=> {
			console.log(res);
			// 登录成功
			if(res.data.status==1) {
				// 保存token
				localStorage.setItem("token",res.data.token)
				// 保存用户信息
				localStorage.setItem("user",JSON.stringify(res.data.user))
				// 跳转页面
				uni.switchTab({
					url:"/pages/tabbar/home"
				})
				
			}
		})
	}
	function radioClick() {
		isRead.value = !isRead.value
	}
	
	// 在第3位和第7位后添加空格
	function inputPhone(e) {
	  const oldValue = phone.value;
	  phone.value = e.detail.value;
	  if (phone.value.length >= 4 && oldValue.length < phone.value.length) {
	    phone.value = phone.value.replace(/^(\d{3})(\d+)/g, "$1 $2");
	  }
	  if (phone.value.length === 4 && oldValue.length > phone.value.length) {
	    phone.value = phone.value.slice(0, 3);
	  }
	  if (phone.value.length >= 8 && oldValue.length < phone.value.length) {
	     phone.value = phone.value.replace(/^(\d{3}\s\d{4})(\d+)/g, "$1 $2");
	  }
	  if (phone.value.length === 8 && oldValue.length > phone.value.length) {
	    phone.value = phone.value.slice(0, 7);
	  }
	}
	
	
</script>

<style scoped>
	
	.login {
		padding: 0 2rem;
	}
	.head {
		font-size: 20px;
		font-weight: 500;
		margin-bottom: 1.3rem;
	}
	.account,.psw {
		position: relative;
		display: flex;
		align-items: center;
		border-radius: 5px;
		background-color: #ececec;
		margin-bottom: .5rem;
	}
	#phone,#password {
		padding: 10px 0;
		caret-color: red;
	}
	.phone-head {
		padding: 0 10px;
	}
	#password {
		padding-left: 10px;
	}
	.err {
		color: red;
		font-size: 14px;
	}
	.protocol {
		font-size: 14px;
		
		margin-top: .3rem;
		margin-bottom: .3rem;
	}
	.protocol radio {
		scale: .8;
	}
	.tips {
		font-size: 14px;
		color: #b4b4b4;
		margin-bottom: .5rem;
	}
	.btn {
		background-color: #ff0059;
		color: #fff;
		position: relative;
		margin-bottom: 1rem;
	}
	.mask {
		position: absolute;
		left: 0;
		right: 0;
		top: 0;
		bottom: 0;
		opacity: .6;
		background-color: #fff;
	}
	.other {
		display: flex;
		justify-content: space-between;
		font-size: 14px;
	}
	.highlight {
		color: #094fe8;
	}
</style>