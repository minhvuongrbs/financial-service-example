package campaign

type Strategy struct {
	StrategyDetail StrategyDetail `json:"strategy_detail"`
}

func NewStrategy(detail StrategyDetail) Strategy {
	return Strategy{
		StrategyDetail: detail,
	}
}

type StrategyDetail interface{}

type StrategyDetailTopRegister struct {
	Name              StrategyDetailName   `json:"name"`
	TopRegisterNumber int64                `json:"top_register_number"` // apply for top x user join campaign, example: 100
	Vouchers          []VoucherApplication `json:"vouchers"`            // list vouchers receive for each user
}

type StrategyDetailName string

const (
	StrategyDetailNameTopRegister = "StrategyDetailTopRegister"
	StrategyDetailNameSpending    = "StrategyDetailSpending"
)

type StrategyDetailSpending struct {
	Name     StrategyDetailName   `json:"name"`
	Vouchers []VoucherApplication `json:"vouchers"`
	Amount   int                  `json:"amount"`
}

func NewStrategyDetailTopRegister(topNumber int64, vouchers []VoucherApplication) StrategyDetailTopRegister {
	return StrategyDetailTopRegister{
		Name:              StrategyDetailNameTopRegister,
		TopRegisterNumber: topNumber,
		Vouchers:          vouchers,
	}
}
