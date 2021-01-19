package model

type Group struct {
	BaseUUIDModel
	Name string `gorm:"type:varchar(36);not null;"`
}

func (Group) TableName() string {
	return TblNameGroup
}
