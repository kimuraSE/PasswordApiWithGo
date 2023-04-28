package model

import "time"

type Title struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignkey:UserId";constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}

type TitleResponse struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
