export function getCommentTime(timestamp) {
	let nowDate= new Date()
	
	let date = new Date(timestamp*1000)
	
	let Y = date.getFullYear()
	let M =
		(date.getMonth() + 1 < 10 ?
			"0" + (date.getMonth() + 1) :
			date.getMonth() + 1)
	let D = (date.getDate() < 10 ? "0" + date.getDate() : date.getDate()) + " ";
	let h = date.getHours() + ":";
	let m = date.getMinutes();
	let s = date.getSeconds();
	
	if(nowDate.getFullYear()>date.getFullYear()) {
		return Y+"-"+M+"-"+D
	}else{
		if(nowDate.getMonth()>date.getMonth()){
			return M+"-"+D
		}else {
			let timeDiff = nowDate.getDate()-date.getDate()
			if(timeDiff==0) {
				return h+m
			}else if(timeDiff==1){
				return "昨天"+h+m
			}else {
				return timeDiff+"天前"
			}
 		}
		
	}
}
