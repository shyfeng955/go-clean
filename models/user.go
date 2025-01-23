package models

type User struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
}

func (u *User) TableName() string {
	return "t_user"
}

type UserVo struct {
	Id int `json:"id"`
}
