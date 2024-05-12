package voucher

import "time"

type Voucher struct {
	Id int64

	Type           Type
	Amount         int64
	IsActive       bool
	ExpirationTime time.Time
	AppId          string
}

type Type int

const (
	FixedAmount Type = iota
	PercentageDiscount
)
