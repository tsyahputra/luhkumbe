package model

type Workunit struct {
	ID         int32  `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Nama       string `json:"nama" gorm:"type:varchar(255);not null"`
	NamaPendek string `json:"nama_pendek" gorm:"type:varchar(127);not null"`
	Kode       string `json:"kode" gorm:"type:varchar(15);not null"`
}

type WorkunitsAndRoles struct {
	Workunits []Workunit `json:"workunits"`
	Roles     []Role     `json:"roles"`
}
