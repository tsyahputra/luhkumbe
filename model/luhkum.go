package model

import "time"

type Luhkum struct {
	ID               int32          `json:"id" gorm:"primaryKey;autoIncrement"`
	ClassificationID int32          `json:"classification_id" gorm:"type:int;not null;index:classification_id"`
	Provinsi         string         `json:"provinsi" gorm:"type:varchar(64);not null"`
	Kabupaten        string         `json:"kabupaten" gorm:"type:varchar(64);not null"`
	Kecamatan        string         `json:"kecamatan" gorm:"type:varchar(64);not null"`
	Kelurahan        string         `json:"kelurahan" gorm:"type:varchar(64);not null"`
	Dir              string         `json:"dir" gorm:"type:varchar(100);not null"`
	Document         string         `json:"document" gorm:"type:varchar(50);not null"`
	Catatan          string         `json:"catatan" gorm:"null"`
	Modified         time.Time      `json:"modified" gorm:"type:datetime;autoUpdateTime:milli"`
	Classification   Classification `json:"classification" gorm:"foreignKey:ClassificationID"`
}

type AllLuhkums struct {
	Luhkums []Luhkum `json:"luhkums"`
	Total   int64    `json:"total"`
}
