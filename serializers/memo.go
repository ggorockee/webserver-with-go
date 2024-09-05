package serializers

import (
	"time"
)

type Memo struct {
	Id   uint               `json:"id"`
	User TinyUserSerializer `json:"user"`

	Title   string `json:"title"`
	Content string `json:"content"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateMemoSerializer struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

//
//func (m *Memo) MemoListSerializer(
//	memo models.Memo,
//	user TinyUserSerializer,
//) Memo {
//	return Memo{
//		Id:        memo.Id,
//		User:      user,
//		Title:     memo.Title,
//		Content:   memo.Content,
//		CreatedAt: memo.CreatedAt,
//		UpdatedAt: memo.UpdatedAt,
//	}
//}
