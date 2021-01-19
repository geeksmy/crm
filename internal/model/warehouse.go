package model

type Warehouse struct {
	BaseUUIDModel
	Name      string `gorm:"type:varchar(128);"`
	Code      string `gorm:"type:varchar(32)"`
	Address   string `gorm:"type:varchar(256);"`
	Type      int    `gorm:"type:int"`
	FounderID string `gorm:"type:varchar(36)"`

	// 用于返回
	Founder User
}

func (Warehouse) TableName() string {
	return TblNameWarehouse
}
