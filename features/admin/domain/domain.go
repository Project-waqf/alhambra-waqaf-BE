package domain

type Admin struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Token    string
	Image    string
	FileId   string
}

type UseCaseInterface interface {
	Login(input Admin) (Admin, error)
	Register(input Admin) error
	UpdatePassword(input Admin) error
	ForgotSendEmail(input Admin) (Admin, error)
	ForgotUpdate(token, password string) error
	UpdateProfile(input Admin) (Admin, error)
	UpdateImage(input Admin) error
	GetProfile(id uint) (Admin, error)
}

type RepoInterface interface {
	Login(input Admin) (Admin, error)
	Register(input Admin) error
	GetUser(input Admin) error
	UpdatePassword(input Admin) error
	GetFromRedis(email string) (string, error)
	SaveRedis(email, token string) error
	UpdatePasswordByEmail(input Admin) error
	DeleteToken(token string) error
	UpdateProfile(input Admin) (Admin, error)
	GetUserById(id uint) (Admin, error)
	UpdateImage(input Admin) error
}
