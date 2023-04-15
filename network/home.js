import {request} from "@/network/request.js"
export function getVideos(page,uid) {
	return request("get","/getVideoList",{page,uid})
}

// 点赞
export function likeThisVideo(uid,videoId,writerId) {
	return request("get","/like",{uid,videoId,writerId})
}
// 取消点赞
export function cancelLikeThisVideo(uid,videoId,writerId) {
	return request("get","/cancelLike",{uid,videoId,writerId})
}





// 评论
export function getComment(page,videoId) {
	return request("get","/getComment",{page,videoId})
}
