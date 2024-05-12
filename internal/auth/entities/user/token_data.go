package user

type TokenData struct {
	UserId int64 `json:"user_id"`

	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	PhoneNumber int64  `json:"phone_number"`
	FullName    string `json:"full_name"`

	GeneratedAt int64 `json:"generated_at"` //unix millisecond timestamp
	ExpiredAt   int64 `json:"expired_at"`   //unix millisecond timestamp
}
