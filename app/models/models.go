package models

import (
	"Mugen-Spiegel/findmo/config"
	"time"
)

type Action interface {
	Create() struct{}
	Update() struct{}
	Delete() bool
}

type User struct {
	ID         int       `json:"id" gorm:"primary_key"`
	FirstName  string    `json:"first_name"`
	MiddleName string    `json:"middle_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (user User) Create() User {

	result := config.DB.Create(&user)
	if result.Error != nil {
		// error handling...
		panic(result.Error)
	}

	return user
}

func (user User) Find() User {

	config.DB.Find(&user)
	return user
}

func (user User) Update() User {

	result := config.DB.Updates(&user)
	if result.Error != nil {
		// error handling...
		panic(result.Error)
	}
	return user
}

func (user User) Delete() bool {
	result := config.DB.Delete(&user)
	if result.Error != nil {
		// error handling...
		panic(result.Error)
	}

	return true
}

type Service struct {
	ID          int    `json:"id" gorm:"primary_key"`
	ServiceName string `json:"service_name"`
	UserID      int    `json:"user_id" gorm:"primaryKey"`
	User        User
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (user Service) Create() Service {
	return user
}

func (user Service) Update() Service {
	return user
}

func (user Service) Delete() Service {
	return user
}

type Post struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Post      string `json:"service_name" gorm:"type:text"`
	UserID    int    `json:"user_id" gorm:"primaryKey"`
	User      User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (post Post) Create() Post {
	return post
}

func (post Post) Update() Post {
	return post
}

func (post Post) Delete() Post {
	return post
}
