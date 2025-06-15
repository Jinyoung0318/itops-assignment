package model

import (
	"time"
)

// 사용자 구조체
type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// 이슈 구조체
type Issue struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	User        *User     `json:"user,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// 상태값 정의
var ValidStatuses = map[string]bool{
	"PENDING":     true,
	"IN_PROGRESS": true,
	"COMPLETED":   true,
	"CANCELLED":   true,
}

// 임시 데이터
var PredefinedUsers = []User{
	{ID: 1, Name: "김개발"},
	{ID: 2, Name: "이디자인"},
	{ID: 3, Name: "박기획"},
}
