package model

import "time"

type User struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Email     string    `json:"email" gorm:"type:varchar(100);unique_index"`
	Password  string    `json:"password" gorm:"type:varchar(100)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	Id    uint   `json:"id" gorm:"primary_key"`
	Email string `json:"email" gorm:"type:varchar(100);unique_index"`
}
