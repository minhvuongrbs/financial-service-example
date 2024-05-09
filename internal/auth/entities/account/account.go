package account

import (
	"fmt"
	"time"
)

type Account struct {
	Id int64 `json:"id"`

	Username    string `json:"username"`
	PhoneNumber int64  `json:"phone_number"`
	Email       string `json:"email"`

	PersonalInfo PersonalInfo `json:"personal_info"`

	FullName string `json:"full_name"`
	Password string `json:"password"` // hashed password
}

type PersonalInfo struct {
	BirthDay time.Time `json:"birth_day"`
}

func NewAccount(username string, phone int64, email string,
	birthday string, fullName string, password string) (a *Account, err error) {
	if isValidFullName(fullName) {
		return nil, fmt.Errorf("invalid fullname")
	}
	if IsValidPassword(password) {
		return nil, fmt.Errorf("invalid password")
	}

	if username == "" && phone == 0 && email == "" {
		return nil, fmt.Errorf("all identidy(phone, email, username) are empty")
	}
	if username != "" && !isValidUsername(username) {
		return nil, fmt.Errorf("invalid user name")
	}
	if phone != 0 && isValidVNPhone(phone) {
		return nil, fmt.Errorf("invalid VN phone")
	}
	if email != "" && isValidEmail(email) {
		return nil, fmt.Errorf("invalid email")
	}

	var birthDayTime time.Time
	if birthday != "" {
		birthDayTime, err = time.Parse(time.DateOnly, birthday)
		if err != nil {
			return nil, fmt.Errorf("invalid birthday format: %w", err)
		}
	}

	hpw, err := hashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	return &Account{
		Id:          0,
		Username:    username,
		PhoneNumber: phone,
		Email:       email,
		PersonalInfo: PersonalInfo{
			BirthDay: birthDayTime,
		},
		FullName: fullName,
		Password: hpw,
	}, nil
}

func (a *Account) WithId(accountId int64) {
	a.Id = accountId
}
