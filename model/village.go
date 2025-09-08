package model

type Village struct {
	ID         int32  `json:"id" gorm:"primaryKey;autoIncrement"`
	DistrictID string `json:"district_code" gorm:"type:varchar(8);not null;index:district_id"`
	Code       string `json:"code" gorm:"type:varchar(14);not null"`
	Name       string `json:"name" gorm:"type:varchar(64);not null"`
}

type VillageAPIResponse struct {
	Data []Village `json:"data"`
}
