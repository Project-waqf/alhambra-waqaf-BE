package paymentgateway

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
	"wakaf/features/wakaf/domain"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func Midtrans(input domain.PayWakaf) (*snap.Response, string) {
	environment := os.Getenv("MIDTRANS_ENV")
	var midtransEnv midtrans.EnvironmentType
	if environment == "1" {
		midtransEnv = 1
	} else {
		midtransEnv = 2
	}

	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtransEnv)
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
		EnabledPayments: []snap.SnapPaymentType{
			snap.PaymentTypeGopay,
			snap.PaymentTypeShopeepay,
			snap.PaymentTypeCreditCard,
			snap.PaymentTypeBankTransfer,
			snap.PaymentTypeBNIVA,
			snap.PaymentTypePermataVA,
			snap.PaymentTypeBCAVA,
			snap.PaymentTypeBRIVA,
			snap.PaymentTypeOtherVA,
			snap.PaymentTypeMandiriClickpay,
			snap.PaymentTypeCimbClicks,
			snap.PaymentTypeDanamonOnline,
			snap.PaymentTypeKlikBca,
			snap.PaymentTypeBCAKlikpay,
			snap.PaymentTypeBRIEpay,
			snap.PaymentTypeMandiriEcash,
			snap.PaymentTypeTelkomselCash,
			snap.PaymentTypeEChannel,
			snap.PaymentTypeIndomaret,
			snap.PaymentTypeKioson,
			snap.PaymentTypeAlfamart,
			snap.PaymentTypeConvenienceStore,
			"other_qris",
			// snap.PaymentTypeAkulaku,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: input.Name,
			Email: input.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)
	return snapResp, orderId
}

func DenyTransaction(input string) (string, error) {
	url := fmt.Sprintf("https://api.sandbox.midtrans.com/v2/%s/deny", input)

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	return string(body), nil

}
