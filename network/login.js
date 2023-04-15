import {request} from "./request.js"

export function login(phone,password) {
	return request("post","/login",{
		phone,
		password
	})
}