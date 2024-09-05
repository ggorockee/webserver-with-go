package database

import "time"

type Memo struct {
	Id uint `json:"id" gorm:"primaryKey"`

	//UserId int  `json:"user_id"`
	//User   User `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`

	Title   string `json:"title"`
	Content string `json:"content"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
