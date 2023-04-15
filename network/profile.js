import {request} from "@/network/request.js"
export function getLikeList(uid) {
	return request("get","/getLikeList",{uid})
}