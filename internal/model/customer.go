package model

type Customer struct {
	BaseUUIDModel
	Name     string `gorm:"type:varchar(128);"`
	Src      int    `gorm:"type:int"`
	Mobile   string `gorm:"type:varchar(11);"`
	Url      string `gorm:"type:varchar(128);"`
	Email    string `gorm:"type:varchar(128);"`
	Industry int    `gorm:"type:int"`
	Level    int    `gorm:"type:int"`
	Note     string `gorm:"type:varchar(256);"`
	Address  string `gorm:"type:varchar(256);"`
}

func (Customer) TableName() string {
	return TblNameCustomer
}
