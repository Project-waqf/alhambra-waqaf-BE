package domain

type Admin struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Token    string
}

type UseCaseInterface interface {
	Login(input Admin) (Admin, error)
	Register(input Admin) error
	UpdatePassword(input Admin) error
	ForgotSendEmail(input Admin) (Admin, error)
}

type RepoInterface interface {
	Login(input Admin) (Admin, error)
	Register(input Admin) error
	GetUser(input Admin) error 
	UpdatePassword(input Admin) error
}
