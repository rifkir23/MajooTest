package repository

type MerchantRepository interface {
}

type merchantRepo struct {
}

func NewMerchantRepo() MerchantRepository {
	return &merchantRepo{}
}
