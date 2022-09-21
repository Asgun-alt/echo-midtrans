package usecase

import (
	"context"
	"echo-midtrans/pkg/domain/transaction"
	"echo-midtrans/pkg/domain/users"
	"strconv"

	"github.com/spf13/viper"
	"github.com/veritrans/go-midtrans"
)

type PaymentUseCase struct{}

func NewPaymentUseCase() *PaymentUseCase {
	return &PaymentUseCase{}
}

func (uc *PaymentUseCase) GetPaymentURL(ctx context.Context, transaction *transaction.Transaction, user users.User) (string, error) {
	midtransClient := midtrans.NewClient()
	midtransClient.ServerKey = viper.GetString("midtrans.ServerKey")
	midtransClient.ClientKey = viper.GetString("midtrans.ClientKey")
	midtransClient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{Client: midtransClient}
	snapRequest := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Username,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(int(transaction.ID)),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResponse, err := snapGateway.GetToken(snapRequest)
	if err != nil {
		return "", err
	}
	return snapTokenResponse.RedirectURL, nil
}
