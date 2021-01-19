package model

type Supplier struct {
	BaseUUIDModel
	Name           string `gorm:"type:varchar(128);"`
	Level          int    `gorm:"type:int"`
	ContactName    string `gorm:"type:varchar(128);"`
	ContactPhone   string `gorm:"type:varchar(11);"`
	ContactAddress string `gorm:"type:varchar(128);"`
	Note           string `gorm:"type:varchar(256);"`
	HeadID         string `gorm:"type:varchar(36)"` // 负责人用户ID
	FounderID      string `gorm:"type:varchar(36)"` // 创建人用户ID

	// 用于返回
	Founder User
	Head    User
}

func (Supplier) TableName() string {
	return TblNameSupplier
}
