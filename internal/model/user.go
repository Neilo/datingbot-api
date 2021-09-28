package model

import "context"

//User user of dating bot
type User struct {
	UserNameTG  string   `json:"id" db:"id"`
	ChatID      *string  `json:"chat_id" db:"chat_id"`
	Name        *string  `json:"name,omitempty" db:"name"`
	Description *string  `json:"description,omitempty" db:"description"`
	Age         *int     `json:"age,omitempty" db:"age"`
	Gender      *string  `json:"gender,omitempty" db:"gender"`
	FindSex     *string  `json:"find_sex,omitempty" db:"find_sex"`
	Longitude   *float64 `json:"longitude,omitempty" db:"longitude"`
	Latitude    *float64 `json:"latitude,omitempty" db:"latitude"`
	Rating      *float64 `json:"rating" db:"rating"`
}

//UserRepository интерфейс для реализации хранения пользователей
type UserRepository interface {
	//return id and error
	Store(ctx context.Context, user *User) (string, error)
	Update(ctx context.Context, id *string, user *User) error
	Find(ctx context.Context, id *string) (*User, error)
	//выбрать следующего чела для лайка
	Pick(ctx context.Context, id *string) (*User, error)
}
