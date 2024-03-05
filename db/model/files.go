package model

import "time"

// CounterModel 计数器模型
type FilesModel struct {
	Id        int32     `gorm:"column:id" json:"id"`
	Key       string    `gorm:"column:key" json:"key"`
	Folder    string    `gorm:"column:folder" json:"folder"`
	Content   string    `gorm:"column:content" json:"content"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}
