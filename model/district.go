package model

type District struct {
	ID        int32  `json:"id" gorm:"primaryKey;autoIncrement"`
	RegencyID string `json:"regency_code" gorm:"type:varchar(5);not null;index:regency_id"`
	Code      string `json:"code" gorm:"type:varchar(8);not null"`
	Name      string `json:"name" gorm:"type:varchar(64);not null"`
}

type DistrictAPIResponse struct {
	Data []District `json:"data"`
}
