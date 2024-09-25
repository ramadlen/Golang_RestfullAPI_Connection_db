package migration

import (
	"time"

	"gorm.io/gorm"
)

type Biodata struct {
	ID       uint64 `json:"id" gorm:"primaryKey;column:id"`
	Username string `json:"username" gorm:"column:username"`
	FullName string `json:"fullname" gorm:"column:fullname"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
	CreateAt time.Time 	`json:"create_at" gorm:"column:create_at;autoCreateTime:true;"`
	UpdateAt time.Time 	`json:"update_at" gorm:"column:update_at;autoCreateTime:true;autoUpdateTime:true"`
	DeleteAt gorm.DeletedAt `json:"-" gorm:"index;column:deleted_at"`
}
func (bio *Biodata) TableName() string {
	return "biodata"
}