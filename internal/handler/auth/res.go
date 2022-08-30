package auth

import "time"

type LoginRes struct {
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	IsSuper   uint8     `json:"is_supper"`
	IsActive  uint8     `json:"is_active"`
	LastLogin time.Time `json:"last_login"`
	Token     string    `json:"token"`
}
