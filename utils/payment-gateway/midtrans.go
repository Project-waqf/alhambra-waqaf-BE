package paymentgateway

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"wakaf/features/wakaf/domain"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func PayBill(input domain.PayWakaf) (string, string) {
	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).
	var orderId = "ORDER-" + (time.Now().Format("02-Jan-06 15:04")) + fmt.Sprintf("%f", rand.Float64())

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: int64(input.GrossAmount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: input.Name,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)
	return snapResp.RedirectURL, orderId
}
