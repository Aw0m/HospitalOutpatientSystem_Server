package cqe

import (
	"time"
)

type RegisterRequest struct {
	ID           string    `json:"id"`                 // 用户ID
	Name         string    `json:"name"`               // 用户姓名
	Birthday     time.Time `json:"birthday,omitempty"` // 出生日期
	Password     string    `json:"password"`           // 密码
	Position     string    `json:"position"`           // 职位
	Profile      string    `json:"profile"`
	IsExpert     bool      `json:"is_expert"`
	DepartmentID string    `json:"department_id"`
	Gender       string    `json:"gender"`
}

type LoginRequest struct {
	UserID   string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
