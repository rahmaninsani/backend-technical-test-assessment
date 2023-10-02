package web

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRefreshAccessTokenRequest struct {
	RefreshToken string
}

type UserProfileRequest struct {
	Username string `json:"username"`
}
