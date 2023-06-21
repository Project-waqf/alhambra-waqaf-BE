package delivery

import (
	"strconv"
	"strings"
	"time"
	"wakaf/features/wakaf/domain"
)

type WakafResponse struct {
	ID            uint    `json:"id"`
	Title         string  `json:"title"`
	Detail        string  `json:"detail"`
	Category      string  `json:"category"`
	Picture       string  `json:"picture"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
	Collected     int     `json:"collected"`
	FundTarget    int     `json:"fund_target"`
	DueDate       int     `json:"due_date"`
	DueDateString string  `json:"due_date_string"`
	Status        string  `json:"status"`
	IsComplete    bool    `json:"is_complete"`
	Donors        []Donor `json:"donors"`
}

type Donor struct {
	Name      string `json:"name"`
	Fund      int    `json:"fund"`
	Doa       string `json:"doa"`
	CreatedAt string `json:"created_at"`
}

type PayWakafRes struct {
	IdWakaf     uint   `json:"id_wakaf"`
	Name        string `json:"name"`
	GrossAmount int    `json:"gross_amount"`
	Doa         string `json:"doa"`
	CreatedAt   string `json:"created_at"`
	RedirectURL string `json:"redirect_url"`
	Token       string `json:"token"`
}

type SummaryWakafRes struct {
	TotalProgram int `json:"total_program"`
	TotalWakaf   int `json:"total_wakaf"`
	TotalWakif   int `json:"total_wakif"`
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func daysBetween(a, b time.Time) int {
	// fmt.Printf("a=%d-%d-%d b=%d-%d-%d\n", a.Year(), int(a.Month()), a.Day(), b.Year(), int(b.Month()), b.Day())
	days := a.Sub(b).Hours() / 24
	if days < 0 {
		days *= -1
	}
	return int(days)
}

func FromDomainAdd(input domain.Wakaf) WakafResponse {

	// Days between now and due date
	dueDate := input.DueDate.Format("2006-1-2")
	dt := strings.Split(dueDate, "-")
	timeNow := time.Now()
	var date1 []int
	for _, v := range dt {
		time, err := strconv.Atoi(v)
		if err != nil {
			logger.Error("Failed to convert date")
		}
		date1 = append(date1, time)
	}

	t1 := date(timeNow.Year(), int(timeNow.Month()), timeNow.Day())
	t2 := date(date1[0], date1[1], date1[2])
	days := daysBetween(t1, t2)

	if input.DueDate.Before(timeNow) {
		days = 0
	}

	return WakafResponse{
		ID:         input.ID,
		Title:      input.Title,
		Category:   input.Category,
		Picture:    input.Picture,
		CreatedAt:  input.CreatedAt.Format("02 Jan 2006"),
		UpdatedAt:  input.UpdatedAt.Format("Monday, 02-01-2006 T15:04:05"),
		Collected:  input.Collected,
		FundTarget: input.FundTarget,
		Detail:     input.Detail,
		DueDate:    days,
		Status:     input.Status,
	}
}

func FromDomainGetAll(input []domain.Wakaf) []WakafResponse {
	var res []WakafResponse

	for _, v := range input {
		if v.Collected == v.FundTarget {
			v.IsComplete = true
		}
		resDetail := FromDomainGet(v)
		res = append(res, resDetail)
	}
	return res
}

func FromDomainGet(input domain.Wakaf) WakafResponse {
	// Days between now and due date
	dueDate := input.DueDate.Format("2006-1-2")
	dt := strings.Split(dueDate, "-")
	timeNow := time.Now()
	var date1 []int
	for _, v := range dt {
		time, err := strconv.Atoi(v)
		if err != nil {
			logger.Error("Failed to convert date")
		}
		date1 = append(date1, time)
	}

	t1 := date(timeNow.Year(), int(timeNow.Month()), timeNow.Day())
	t2 := date(date1[0], date1[1], date1[2])
	days := daysBetween(t1, t2)

	if input.DueDate.Before(timeNow) {
		days = 0
	}

	var newDonors []Donor

	for _, v := range input.Donors {
		tmp := Donor{
			Name:      v.Name,
			Fund:      v.Fund,
			Doa:       v.Doa,
			CreatedAt: v.Created_at.Format("02/01/2006"),
		}
		newDonors = append(newDonors, tmp)
	}

	return WakafResponse{
		ID:            input.ID,
		Title:         input.Title,
		Detail:        input.Detail,
		Category:      input.Category,
		Picture:       input.Picture,
		CreatedAt:     input.CreatedAt.Format("02 January 2006"),
		UpdatedAt:     input.UpdatedAt.Format("Monday, 02-01-2006 T15:04:05"),
		DueDate:       days,
		DueDateString: input.DueDate.Format("2006-01-02"),
		Collected:     input.Collected,
		FundTarget:    input.FundTarget,
		IsComplete:    input.IsComplete,
		Status:        input.Status,
		Donors:        newDonors,
	}
}

func FromDomainPaywakaf(input domain.PayWakaf) PayWakafRes {
	return PayWakafRes{
		IdWakaf:     uint(input.IdWakaf),
		Name:        input.Name,
		GrossAmount: input.GrossAmount,
		Doa:         input.Doa,
		CreatedAt:   input.CreatedAt.Format("02/01/2006 15:04"),
		RedirectURL: input.RedirectURL,
		Token:       input.Token,
	}
}

func SummaryWakaf(count, sum, wakif int) SummaryWakafRes {
	return SummaryWakafRes{
		TotalProgram: count,
		TotalWakaf:   sum,
		TotalWakif:   wakif,
	}
}
