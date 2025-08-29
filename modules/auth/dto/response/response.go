package response

type LoginResponse struct {
	AccessToken string        `json:"access_token"`
	TokenType   string        `json:"token_type"`
	ExpiresIn   int           `json:"expires_in"`
	User        UserShortInfo `json:"user"`
}

type UserShortInfo struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
	Role     string `json:"role"`
}
