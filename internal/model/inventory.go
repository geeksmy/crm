package model

import (
	"time"
)

type Inventory struct {
	BaseUUIDModel
	ProductID   string    `gorm:"type:varchar(36)"`
	Number      int       `gorm:"type:int"`
	Code        string    `gorm:"type:varchar(32)"`
	WarehouseID string    `gorm:"type:varchar(36)"`
	Type        int       `gorm:"type:int"`
	Date        time.Time `msgpack:"-"`
	Note        string    `gorm:"type:varchar(256);"`
	HeadID      string    `gorm:"type:varchar(36)"`
	FounderID   string    `gorm:"type:varchar(36)"`
	InAndOut    int       `gorm:"type:int"`

	// 用于返回
	Product   Product
	Warehouse Warehouse
	Founder   User
	Head      User
}

func (Inventory) TableName() string {
	return TblNameInventory
}
