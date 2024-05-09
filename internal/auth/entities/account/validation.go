package account

import (
	"fmt"
	"net/mail"
	"strconv"

	"github.com/nyaruka/phonenumbers"
	"github.com/spf13/cast"
)

type IdentityType int

const (
	IdentityTypePhoneNumber IdentityType = iota
	IdentityTypeEmail
	IdentityTypeUsername
	IdentityTypeUnknown
)

func CheckIdentityType(identity string) (IdentityType, error) {
	if isValidEmail(identity) {
		return IdentityTypeEmail, nil
	}
	if isValidVNPhone(cast.ToInt64(identity)) {
		return IdentityTypePhoneNumber, nil
	}
	if isValidUsername(identity) {
		return IdentityTypeUsername, nil
	}
	return IdentityTypeUnknown, fmt.Errorf("invalid login type, identity: %s", identity)
}

func isValidVNPhone(phone int64) bool {
	_, err := phonenumbers.Parse(strconv.Itoa(int(phone)), "VN")
	if err != nil {
		return false
	}
	return true
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}

const (
	minLengthUserName = 6
	maxLengthUserName = 16
)

func isValidUsername(username string) bool {
	if len(username) < minLengthUserName || len(username) > maxLengthUserName {
		return false
	}
	return true
}

const (
	minimumLengthPassword = 1
	maximumLengthPassword = 30
)

func IsValidPassword(password string) bool {
	if len(password) < minimumLengthPassword || len(password) > maximumLengthPassword {
		return false
	}
	return true
}

func isValidFullName(fullName string) bool {
	if len(fullName) == 0 {
		return false
	}
	return true
}
