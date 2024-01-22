package model

type User struct {
	Username    string  `json:"username"`
	Password    string  `json:"password"`
	ID          uint    `gorm:"primaryKey;autoIncrement;not null'"`
	Email       *string `json:"email" gorm:"default:null"`
	PhoneNumber *uint   `json:"phone_number" gorm:"default:null"`
}

type UserInfo struct {
	Username string `json:"username" gorm:"default:null"`
	Gender   string `json:"gender" gorm:"default:null"`
	Sign     string `json:"sign" gorm:"default:null"`
}
