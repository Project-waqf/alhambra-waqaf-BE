package domain

import (
	"time"
)

type Wakaf struct {
	ID         uint
	Title      string
	Category   string
	Picture    string
	FileId     string
	Detail     string
	FundTarget int
	IsComplete bool
	DueDate    time.Time
	Collected  int
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Donors     []Donors
}

type Donors struct {
	Name       string
	Fund       int
	Doa        string
	Created_at time.Time `json:"created_at,omitempty"`
}

type PayWakaf struct {
	IdWakaf     int
	Name        string
	Email       string
	GrossAmount int
	Doa         string
	CreatedAt   time.Time
	RedirectURL string `json:"RedirectURL,omitempty"`
	OrderId     string
	Status      string
	PaymentType string `json:"PaymentType,omitempty"`
	Token       string `json:"Token,omitempty"`
	Payment     Payment
}

type Payment struct {
	Merchant string
	Tax      int
}

type UseCaseInterface interface {
	AddWakaf(input Wakaf) (Wakaf, error)
	GetAllWakaf(category string, page int, isUser bool, sort, filter, status string) ([]Wakaf, int, int, int, error)
	UpdateWakaf(id uint, input Wakaf) (Wakaf, error)
	DeleteWakaf(id uint) (Wakaf, error)
	GetFileId(id uint) (string, error)
	GetSingleWakaf(id uint) (Wakaf, error)
	PayWakaf(input PayWakaf) (PayWakaf, error)
	UpdatePayment(input PayWakaf) (PayWakaf, error)
	DenyTransaction(input string) error
	SearchWakaf(input string) ([]Wakaf, int, int, int, error)
	GetSummary() (int, int, int, error)
	GetSummaryDashboard() (int, int, int, error)
}

type RepoInterface interface {
	Insert(input Wakaf) (Wakaf, error)
	GetAllWakaf(category string, page int, isUser bool, sort, filter, status string) ([]Wakaf, int, int, int, error)
	Edit(id uint, input Wakaf) (Wakaf, error)
	Delete(id uint) (Wakaf, error)
	GetFileId(id uint) (string, error)
	GetSingleWakaf(id uint) (Wakaf, error)
	PayWakaf(input PayWakaf) (PayWakaf, error)
	UpdatePayment(input PayWakaf) (PayWakaf, error)
	Search(input string) ([]Wakaf, int, int, int, error)
	GetSummary() (int, int, int, error)
	SaveRedis(orderId string, data PayWakaf) error
	GetFromRedis(orderId string) (string, error)
	GetSummaryDashboard() (int, int, int, error)
}
