package model

type User struct {
	BaseModel
	Username string `gorm:"type:varchar(50);uniqueIndex;not null;comment:用户名"`
	Password string `gorm:"type:varchar(100);not null;comment:加密密码"`
	UUID     string `json:"uuid" gorm:"type:char(36);uniqueIndex;not null;comment:用户唯一标识UUID"`
}
