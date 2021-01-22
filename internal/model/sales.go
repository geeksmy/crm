package model

import (
	"time"
)

type Sales struct {
	BaseUUIDModel
	Name            string    `gorm:"type:varchar(128);"`
	Code            string    `gorm:"type:varchar(32)"`
	CustomerID      string    `gorm:"type:varchar(36)"`
	Money           int       `gorm:"type:int"`
	ConsignmentDate time.Time `msgpack:"-"`
	IsSalesReturn   bool      `gorm:"type:boolean;default:false"`
	Note            string    `gorm:"type:varchar(256);"`
	HeadID          string    `gorm:"type:varchar(36)"`
	FounderID       string    `gorm:"type:varchar(36)"`

	// 用于返回
	Customer Customer
	Founder  User
	Head     User
}

func (Sales) TableName() string {
	return TblNameSales
}
