package domain

import "time"

type Wakaf struct {
	ID         uint
	Title      string
	Category   string
	Picture    string
	FileId     string
	Detail     string
	FundTarget int
	IsComplete bool
	DueDate    *time.Time
	Collected  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Donors     []Donors
}

type Donors struct {
	Name string
	Fund int
	Doa  string
}

type PayWakaf struct {
	IdWakaf     int
	Name        string
	GrossAmount int
	Doa         string
	CreatedAt   time.Time
	RedirectURL string
	OrderId     string
	Status      string
	PaymentType string
}

type UseCaseInterface interface {
	AddWakaf(input Wakaf) (Wakaf, error)
	GetAllWakaf(category string, page int) ([]Wakaf, int, error)
	UpdateWakaf(id uint, input Wakaf) (Wakaf, error)
	DeleteWakaf(id uint) (Wakaf, error)
	GetFileId(id uint) (string, error)
	GetSingleWakaf(id uint) (Wakaf, error)
	PayWakaf(input PayWakaf) (PayWakaf, error)
	UpdatePayment(input PayWakaf) (PayWakaf, error)
	DenyTransaction(input string) error
	SearchWakaf(input string) ([]Wakaf, int, error)
	GetSummary() (int, int, int, error)
}

type RepoInterface interface {
	Insert(input Wakaf) (Wakaf, error)
	GetAllWakaf(category string, page int) ([]Wakaf, int, error)
	Edit(id uint, input Wakaf) (Wakaf, error)
	Delete(id uint) (Wakaf, error)
	GetFileId(id uint) (string, error)
	GetSingleWakaf(id uint) (Wakaf, error)
	PayWakaf(input PayWakaf) (PayWakaf, error)
	UpdatePayment(input PayWakaf) (PayWakaf, error)
	Search(input string) ([]Wakaf, error)
	GetSummary() (int, int, int, error)
}
