package model

import "time"

type Password struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     Title     `json:"title" gorm:"foreignkey:TitleId";constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TitleId   uint      `json:"title_id" gorm:"not null"`
}

type PasswordResponse struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
