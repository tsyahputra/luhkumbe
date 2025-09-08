package model

import "time"

type User struct {
	ID                  int32     `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama                string    `json:"nama" gorm:"type:varchar(64);not null"`
	Username            string    `json:"username" gorm:"type:varchar(18);not null;unique"`
	Password            string    `json:"-" gorm:"type:varchar(255);not null"`
	RoleID              int32     `json:"role_id" gorm:"type:int;not null;index:role_id"`
	WorkunitID          int32     `json:"workunit_id" gorm:"type:int;not null;index:workunit_id"`
	CustomKey           string    `json:"-" gorm:"type:varchar(16);null"`
	TwoFASecret         string    `json:"-" gorm:"type:varchar(255);null"`
	TwoFAEnabled        bool      `json:"two_fa_enabled" gorm:"default:false"`
	ResetPasswordToken  string    `json:"-" gorm:"type:varchar(32);null"`
	ResetPasswordExpiry int64     `json:"-"`
	Modified            time.Time `json:"modified" gorm:"type:datetime;autoUpdateTime:milli"`
	Role                Role      `json:"role" gorm:"foreignKey:RoleID"`
	Workunit            Workunit  `json:"workunit" gorm:"foreignKey:WorkunitID"`
}

type UserToken struct {
	User         User   `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AllUsersWithTotal struct {
	Users []User `json:"users"`
	Total int64  `json:"total"`
}

type UserInput struct {
	Nama       string `json:"nama"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	RoleID     int32  `json:"role_id"`
	WorkunitID int32  `json:"workunit_id"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
}
