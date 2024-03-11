package service

import (
	"fmt"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	openid := r.Header.Get("X-WX-OPENID")
	fmt.Println("openid:" + openid)
	user := model.Users{Openid: openid}
	db.Get().Table("user").Where("`openid` = ?", openid).Find(&user)

	str := fmt.Sprintf("%v", user)
	res := JsonResult{
		Code: 200,
		Data: str,
	}

	shouldReturn := writeResultToResponse(res, w)
	if shouldReturn {
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	openid := r.Header.Get("X-WX-OPENID")
	fmt.Println("openid:" + openid)
	user := model.Users{Openid: openid}
	db.Get().Table("user").Where("`openid` = ?", openid).Find(&user)

	if user.Id == 0 {
		username := r.URL.Query().Get("username")
		fmt.Println("the username is:" + username)
		nickname := r.URL.Query().Get("nickname")
		fmt.Println("the nickname is:" + nickname)
		db.Get().Table("user").Create(&model.Users{
			Nickname: nickname,
			Name:     username,
			Openid:   openid,
		})
	}

	res := JsonResult{
		Code: 200,
		Data: true,
	}
	shouldReturn := writeResultToResponse(res, w)
	if shouldReturn {
		return
	}
}
