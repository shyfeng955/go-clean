package models

type User struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
}

func (u *User) TableName() string {
	return "user"
}

type UserVo struct {
	Id int `json:"id"`
}

type AddUserVo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
