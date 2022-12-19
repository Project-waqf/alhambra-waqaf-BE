package domain

type Admin struct {
	ID uint
	Name string
	Username string
	Password string
}

type UseCaseInterface interface {
	Login(input Admin) (Admin, error)
}

type RepoInterface interface {
	Login(input Admin) (Admin, error)
}
