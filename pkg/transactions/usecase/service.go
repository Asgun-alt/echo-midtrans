package usecase

import (
	"context"
	"echo-midtrans/pkg/domain/payment"
	"echo-midtrans/pkg/domain/transaction"
)

type TransactionUseCase struct {
	PaymentService payment.UseCase
	DBRepo         transaction.Repository
}

func NewTransactionUseCase(dbrepo transaction.Repository, paymentService payment.UseCase) *TransactionUseCase {
	return &TransactionUseCase{
		DBRepo:         dbrepo,
		PaymentService: paymentService,
	}
}

func (uc *TransactionUseCase) GetTransactionsByCampaignID(ctx context.Context, campaignID uint) ([]transaction.Transaction, error) {
	campaigns, err := uc.DBRepo.FindByCampaignID(ctx, campaignID)
	if err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (uc *TransactionUseCase) GetTransactionsByUserID(ctx context.Context, userID uint) ([]transaction.Transaction, error) {
	campaigns, err := uc.DBRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (uc *TransactionUseCase) CreateTransaction(ctx context.Context, req *transaction.CreateTransactionRequest) (*transaction.Transaction, error) {
	transaction := transaction.Transaction{}
	transaction.Amount = req.Amount
	transaction.CampaignID = req.CampaignID
	transaction.UserID = req.UserID

	newTransaction, err := uc.DBRepo.Create(ctx, &transaction)
	if err != nil {
		return nil, err
	}

	paymentURL, err := uc.PaymentService.GetPaymentURL(ctx, newTransaction, transaction.User)
	if err != nil {
		return nil, err
	}

	newTransaction.PaymentURL = paymentURL
	err = uc.DBRepo.Update(ctx, newTransaction)
	if err != nil {
		return nil, err
	}

	return newTransaction, nil
}

func (uc *TransactionUseCase) ProcessPayment(ctx context.Context, req *transaction.TransactionNotificationRequest) error {
	// transactionID := req.OrderID

	return nil
}
