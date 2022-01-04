package domain

import "time"

type Task struct {
	ID int64 `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Title string `json:"title" gorm:"notNull"`
	Description string `json:"description" gorm:"type:text;notNull"`
	Status bool `json:"status" gorm:"notNull"`
	UserID int64 `json:"user_id" gorm:"notNull"`
	CategoryID int64 `json:"category_id" gorm:"notNull"`
	CreatedAt time.Time `json:"created_at" gorm:"notNull"`
	UpdatedAt time.Time `json:"updated_at" gorm:"notNull"`
}