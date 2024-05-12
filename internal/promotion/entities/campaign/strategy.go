package campaign

type Strategy interface{}

type StrategyTopRegister struct {
	Name              StrategyName         `json:"name"`
	TopRegisterNumber int64                `json:"top_register_number"` // apply for top x user join campaign, example: 100
	Vouchers          []VoucherApplication `json:"vouchers"`            // list vouchers receive for each user
}

type StrategyName string

const (
	StrategyNameTopRegister = "StrategyTopRegister"
	StrategyNameSpending    = "StrategySpending"
)

type StrategySpending struct {
	Name     StrategyName         `json:"name"`
	Vouchers []VoucherApplication `json:"vouchers"`
	Amount   int                  `json:"amount"`
}

func NewStrategyTopRegister(topNumber int64, vouchers []VoucherApplication) StrategyTopRegister {
	return StrategyTopRegister{
		Name:              StrategyNameTopRegister,
		TopRegisterNumber: topNumber,
		Vouchers:          vouchers,
	}
}
