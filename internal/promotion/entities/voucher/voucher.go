package voucher

type Voucher struct {
	Id int64

	Name        string
	Description string

	Metadata Metadata
	IsActive bool
}

func NewVoucher(id int64, name, description string, isActive bool, metadata Metadata) *Voucher {
	return &Voucher{
		Id:          id,
		Name:        name,
		Description: description,
		Metadata:    metadata,
		IsActive:    isActive,
	}
}
