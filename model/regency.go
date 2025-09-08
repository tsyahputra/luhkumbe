package model

type Regency struct {
	ID         int32  `json:"id" gorm:"primaryKey;autoIncrement"`
	ProvinceID string `json:"province_code" gorm:"type:varchar(2);not null;index:province_id"`
	Code       string `json:"code" gorm:"type:varchar(5);not null"`
	Name       string `json:"name" gorm:"type:varchar(64);not null"`
}

type RegencyAPIResponse struct {
	Data []Regency `json:"data"`
}
