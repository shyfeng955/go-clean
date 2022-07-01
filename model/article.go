package model

import "time"

/**
 * 封装了各种实体类对象，与数据库交互的、与UI交互的等等，任何的实体类都应该放在这里。如：
 */

type Article struct {
	Id       int64     `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	UpdateAt time.Time `json:"update_at"`
	CreateAt time.Time `json:"create_at"`
}

type ArticleVo struct {
	Num        int       `json:"num"`
	CreateData time.Time `json:"create_data"`
}
