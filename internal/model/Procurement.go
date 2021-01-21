package model

type Procurement struct {
	BaseUUIDModel
	SupplierID    string `gorm:"type:varchar(36);"` // 供应商ID
	Code          string `gorm:"type:varchar(32)"`
	Note          string `gorm:"type:varchar(256);"`
	Money         int    `gorm:"type:int"`
	IsSalesReturn bool   `gorm:"type:boolean;default:false"`
	HeadID        string `gorm:"type:varchar(36)"`
	FounderID     string `gorm:"type:varchar(36)"`

	// 用于返回
	Founder  User
	Head     User
	Supplier Supplier
}

func (Procurement) TableName() string {
	return TblNameProcurement
}
