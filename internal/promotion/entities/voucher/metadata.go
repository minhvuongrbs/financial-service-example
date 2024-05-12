package voucher

import "time"

type Metadata struct {
	Strategy       Strategy  `json:"strategy"`
	ExpirationTime time.Time `json:"expiration_time"`
}

type StrategyName string

const (
	StrategyNameFixedAmountForAppIds        = "FixedAmountForAppIds"
	StrategyNamePercentageDiscountForAppIds = "PercentageDiscountForAppIds"
)

type Strategy interface {
}

type StrategyFixedAmountForAppIds struct {
	Name   StrategyName `json:"name"`
	Amount int64        `json:"amount"`
	AppIds []int64      `json:"app_ids"`
}

func NewStrategyFixedAmountForAppIds(amount int64, appIds []int64) StrategyFixedAmountForAppIds {
	return StrategyFixedAmountForAppIds{
		Name:   StrategyNameFixedAmountForAppIds,
		Amount: amount,
		AppIds: appIds,
	}
}

type StrategyPercentageDiscountForAppIds struct {
	Name       StrategyName `json:"name"`
	Percentage int64        `json:"percentage"`
	AppIds     []int64      `json:"app_ids"`
}

func NewStrategyPercentageDiscountForAppIds(percentage int64, appIds []int64) StrategyPercentageDiscountForAppIds {
	return StrategyPercentageDiscountForAppIds{
		Name:       StrategyNamePercentageDiscountForAppIds,
		Percentage: percentage,
		AppIds:     appIds,
	}
}
