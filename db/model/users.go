package model

import "time"

// Users 模型
type Users struct {
	Id        int32     `gorm:"column:id" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Openid    string    `gorm:"column:openid" json:"openid"`
	Nickname  string    `gorm:"column:nickname" json:"nickname"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

func (f *Users) TableName() string {
	return "users"
}
