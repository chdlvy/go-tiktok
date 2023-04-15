const BASEURL = "http://localhost:8000"
export function request(method,url,data={},config={header:{}}) {
	
	if(method=="post" || method=="POST") {
		if(!config.header["content-type"]) {
			config.header["content-type"] = "application/x-www-form-urlencoded"
		}
	}
	// 请求头添加token
	let t=localStorage.getItem("token")
	config.header["authorization"] = t?t:""
	
	return new Promise((reslove,reject)=> {
		uni.request({
			url:BASEURL+url,
			method:method,	
			timeout:5000,
			data:data,
			...config,
			success(res) {
				reslove(res)
			},
			fail(err) {
				reject(err)
			}
		})
	})
}