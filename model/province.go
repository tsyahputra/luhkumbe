package model

type Province struct {
	ID   int32  `json:"id" gorm:"primaryKey;autoIncrement"`
	Code string `json:"code" gorm:"type:varchar(2);not null"`
	Name string `json:"name" gorm:"type:varchar(64);not null"`
}

type ProvinceAPIResponse struct {
	Data []Province `json:"data"`
}
