package model

type Role struct {
	ID    int32  `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama  string `json:"nama" gorm:"type:varchar(100);not null"`
	Users []User `json:"users,omitempty" gorm:"foreignKey:RoleID"`
}
