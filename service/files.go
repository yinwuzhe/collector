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
	}
	shouldReturn := writeResultToResponse(res, w)
	if shouldReturn {
		return
	}
}
func DeleteObject(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	cli := db.Get()
	cli.Where("key = ?", key).Delete(&model.FilesModel{})
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
	oldKey := r.URL.Query().Get("old_key")
	cli := db.Get()
	file := model.FilesModel{}
	//只可以更新名字
	cli.Where("key = ?", oldKey).Take(&file)

	// 修改food模型的值
	file.Key = r.URL.Query().Get("new_key")

	// 等价于: UPDATE `foods` SET `title` = '可乐', `type` = '0', `price` = '100', `stock` = '26', `create_time` = '2018-11-06 11:12:04'  WHERE `foods`.`id` = '2'
	cli.Save(&file)

	res := JsonResult{
		Code: 200,
	}

	shouldReturn := writeResultToResponse(res, w)
	if shouldReturn {
		return
	}
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
