package model

type Product struct {
	BaseUUIDModel
	Name        string `gorm:"type:varchar(128);"`
	Unit        int    `gorm:"type:int"`
	CostPrice   string `gorm:"type:varchar(32)"`
	MarketPrice string `gorm:"type:varchar(32)"`
	Note        string `gorm:"type:varchar(256);"`
	Image       string `gorm:"type:varchar(128);"`
	Code        string `gorm:"type:varchar(32)"`
	Size        string `gorm:"type:varchar(32)"`
	Type        int    `gorm:"type:int"`
	IsShelves   bool   `gorm:"type:boolean;default:false"`
	FounderID   string `gorm:"type:varchar(36)"` // 创建人ID

	// 用于返回
	Founder User
}

func (Product) TableName() string {
	return TblNameProduct
}
