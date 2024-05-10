// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package da_generated

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type CampaignStatus string

const (
	CampaignStatusActive   CampaignStatus = "active"
	CampaignStatusInactive CampaignStatus = "inactive"
)

func (e *CampaignStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CampaignStatus(s)
	case string:
		*e = CampaignStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for CampaignStatus: %T", src)
	}
	return nil
}

type NullCampaignStatus struct {
	CampaignStatus CampaignStatus `json:"campaign_status"`
	Valid          bool           `json:"valid"` // Valid is true if CampaignStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCampaignStatus) Scan(value interface{}) error {
	if value == nil {
		ns.CampaignStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CampaignStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCampaignStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CampaignStatus), nil
}

type VoucherVoucherType string

const (
	VoucherVoucherTypeFixedAmount        VoucherVoucherType = "fixed_amount"
	VoucherVoucherTypePercentageDiscount VoucherVoucherType = "percentage_discount"
)

func (e *VoucherVoucherType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = VoucherVoucherType(s)
	case string:
		*e = VoucherVoucherType(s)
	default:
		return fmt.Errorf("unsupported scan type for VoucherVoucherType: %T", src)
	}
	return nil
}

type NullVoucherVoucherType struct {
	VoucherVoucherType VoucherVoucherType `json:"voucher_voucher_type"`
	Valid              bool               `json:"valid"` // Valid is true if VoucherVoucherType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullVoucherVoucherType) Scan(value interface{}) error {
	if value == nil {
		ns.VoucherVoucherType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.VoucherVoucherType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullVoucherVoucherType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.VoucherVoucherType), nil
}

type AccountVoucher struct {
	ID        int64        `json:"id"`
	AccountID int64        `json:"account_id"`
	VoucherID int64        `json:"voucher_id"`
	IsActive  bool         `json:"is_active"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Campaign struct {
	ID          int64          `json:"id"`
	CampaignKey string         `json:"campaign_key"`
	Status      CampaignStatus `json:"status"`
	Name        string         `json:"name"`
	StartAt     time.Time      `json:"start_at"`
	EndAt       time.Time      `json:"end_at"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}

type CampaignAccount struct {
	ID         int64        `json:"id"`
	AccountID  int64        `json:"account_id"`
	CampaignID int64        `json:"campaign_id"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
}

type CampaignVoucher struct {
	ID            int64        `json:"id"`
	CampaignID    int64        `json:"campaign_id"`
	VoucherID     int64        `json:"voucher_id"`
	TotalVouchers int32        `json:"total_vouchers"`
	CreatedAt     sql.NullTime `json:"created_at"`
	UpdatedAt     sql.NullTime `json:"updated_at"`
}

type Voucher struct {
	ID             int64              `json:"id"`
	VoucherType    VoucherVoucherType `json:"voucher_type"`
	Amount         string             `json:"amount"`
	IsActive       bool               `json:"is_active"`
	ExpirationTime time.Time          `json:"expiration_time"`
	// * will apply to all app_id
	AppID     string       `json:"app_id"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}