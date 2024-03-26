package dao

import (
	"fmt"
	"strings"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

// "wxcloudrun-golang/db"
// "wxcloudrun-golang/db/model"

const fileTable = "files"

// // ClearCounter 清除Counter
// func (imp *CounterInterfaceImp) ClearCounter(id int32) error {
// 	cli := db.Get()
// 	return cli.Table(tableName).Delete(&model.CounterModel{Id: id}).Error
// }

// // UpsertCounter 更新/写入counter
// func (imp *CounterInterfaceImp) UpsertCounter(counter *model.CounterModel) error {
// 	cli := db.Get()
// 	return cli.Table(tableName).Save(counter).Error
// }

// // GetCounter 查询Counter
// func (imp *CounterInterfaceImp) GetCounter(id int32) (*model.CounterModel, error) {
// 	var err error
// 	var counter = new(model.CounterModel)

// 	cli := db.Get()
// 	err = cli.Table(tableName).Where("id = ?", id).First(counter).Error

// 	return counter, err
// }

func CreateRecord(key string, folder string, content string, openid string) error {
	cli := db.Get()
	fmt.Println(cli)
	parts := strings.Split(key, "/")

	lastParts := parts[len(parts)-1]
	fmt.Println(lastParts)
	//处理name
	return cli.Table(fileTable).Create(&model.FilesModel{
		Key:     openid + "/" + key,
		Folder:  folder,
		Content: content,
		Name:    lastParts,
		Openid:  openid,
	}).Error

}
