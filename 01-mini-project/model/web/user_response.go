package web

import (
	"time"
)

type UserResponse struct {
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty" metadata:",optional"`
}

type UserRefreshAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type UserPostListResponse struct {
	Title     string               `json:"title"`
	Slug      string               `json:"slug"`
	Category  PostCategoryResponse `json:"category"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}
