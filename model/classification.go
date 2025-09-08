package model

type Classification struct {
	ID        int32  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" gorm:"type:varchar(64);not null"`
	ShortName string `json:"short_name" gorm:"type:varchar(64);not null"`
}
