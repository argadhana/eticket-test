package request

type LoginRequest struct {
	Username string `json:"username" validate:"required"` // bisa username/email/nip/nim
	Password string `json:"password" validate:"required"`
}

