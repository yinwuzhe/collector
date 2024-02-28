package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wxcloudrun-golang/db/dao"
)

func CreateObject(w http.ResponseWriter, r *http.Request) {
	// 获取上传的文件

	key := r.URL.Query().Get("key")
	folder := r.URL.Query().Get("folder")
	fmt.Println("the key is:" + key)
	dao.CreateRecord(key, folder)

	res := JsonResult{
		Code: 200,
		// Message: "success",
		// Data: string(body),
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)

}
