package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
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

func ObjectList(w http.ResponseWriter, r *http.Request) {
	//改为直接从db里面读取，支持排序，按照创建时间逆序、分页
	prefix := r.URL.Query().Get("prefix")
	start := r.URL.Query().Get("start")
	size := r.URL.Query().Get("size")
	fmt.Println("the type is:" + prefix)

	cli := db.Get()

	sizeint, _ := strconv.Atoi(size)
	startint, _ := strconv.Atoi(start)
	var files []model.FilesModel
	cli.Table("files").Where("folder = ?", prefix).Limit(sizeint).Offset(startint).Order("createdAt desc").Find(&files)
	// fmt.Printf(files)

	res := JsonResult{
		Code: 200,
		// Message: "success",
		Data: files,
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
