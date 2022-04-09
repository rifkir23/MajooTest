package repository

type OutletRepository interface {
}

type outletRepo struct {
}

func NewOutletRepo() OutletRepository {
	return &outletRepo{}
}
