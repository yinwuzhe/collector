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
	key := r.URL.Query().Get("key")
	folder := r.URL.Query().Get("folder")
	fmt.Println("the key is:" + key)
	openid := r.Header.Get("X-WX-OPENID")
	fmt.Println("openid:" + openid)
	dao.CreateRecord(key, folder, r.URL.Query().Get("content"), openid)

	res := JsonResult{
		Code: 200,
	}
	shouldReturn := writeResultToResponse(res, w)
	if shouldReturn {
		return
	}
}
func DeleteObject(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	cli := db.Get()
	cli.Table("files").Where("`key` = ?", key).Delete(&model.FilesModel{})
	fmt.Println("删除成功:" + key)

	res := JsonResult{
		Code: 200,
	}

	shouldReturn := writeResultToResponse(res, w)
	if shouldReturn {
		return
	}
}

func UpdateObject(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	cli := db.Get()
	file := model.FilesModel{}

	cli.Where("`key` = ?", key).Take(&file)

	file.Name = r.URL.Query().Get("name")
	fmt.Println("UpdateObject:" + file.Name)
	fmt.Println(file)
	cli.Save(&file)

	res := JsonResult{
		Code: 200,
	}
	shouldReturn := writeResultToResponse(res, w)
	if shouldReturn {
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	//打印出header
	headers := r.Header
	// fmt.Println("headers {}", headers)
	//从headers里面尝试获取用户的名字和openid,查看是否登录过。没登陆过，则给创建对应的目录
	openid := headers.Get("X-WX-OPENID")
	//完全可信
	fmt.Println("X-WX-OPENID: " + openid)
	res := JsonResult{
		Code: 200,
		Data: openid,
	}
	shouldReturn := writeResultToResponse(res, w)
	if shouldReturn {
		return
	}
}

//

func ObjectList(w http.ResponseWriter, r *http.Request) {
	//改为直接从db里面读取，支持排序，按照创建时间逆序、分页
	prefix := r.URL.Query().Get("prefix")
	start := r.URL.Query().Get("start")
	size := r.URL.Query().Get("size")
	fmt.Println("the type is:" + prefix)
	openid := r.Header.Get("X-WX-OPENID")

	cli := db.Get()

	sizeint, _ := strconv.Atoi(size)
	startint, _ := strconv.Atoi(start)
	var files []model.FilesModel
	cli.Table("files").Where("folder = ? AND openid= ?", openid).Limit(sizeint).Offset(startint).Order("createdAt desc").Find(&files)
	// fmt.Printf(files)

	res := JsonResult{
		Code: 200,
		// Message: "success",
		Data: files,
	}

	shouldReturn := writeResultToResponse(res, w)
	if shouldReturn {
		return
	}
}

func GetObject(w http.ResponseWriter, r *http.Request) {
	// res := &JsonResult{}

	key := r.URL.Query().Get("key")
	fmt.Println("the key is:" + key)
	file := model.FilesModel{}
	db.Get().Table("files").Where("`key` = ?", key).Find(&file)
	fmt.Println("the item is:" + file.Content)
	res := JsonResult{
		Code: 200,
		Data: string(file.Content),
	}

	shouldReturn := writeResultToResponse(res, w)
	if shouldReturn {
		return
	}
}

func writeResultToResponse(res JsonResult, w http.ResponseWriter) bool {
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return true
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
	return false
}
