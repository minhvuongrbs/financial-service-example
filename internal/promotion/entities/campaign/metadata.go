package campaign

import "time"

type Metadata struct {
	Strategy Strategy `json:"strategy"`

	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
}

func NewMetadata(startAt, endAt time.Time, s Strategy) Metadata {
	return Metadata{
		Strategy: s,
		StartAt:  startAt,
		EndAt:    endAt,
	}
}

type VoucherApplication struct {
	VoucherId    int64 `json:"voucher_id"`    // nap dien thoai
	TotalVoucher int64 `json:"total_voucher"` // 1 voucher
}

func NewVoucher(id, totalVoucher int64) VoucherApplication {
	return VoucherApplication{
		VoucherId:    id,
		TotalVoucher: totalVoucher,
	}
}
