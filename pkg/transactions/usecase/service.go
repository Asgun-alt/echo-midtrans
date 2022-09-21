package usecase

import (
	"context"
	"echo-midtrans/pkg/domain/campaign"
	"echo-midtrans/pkg/domain/payment"
	"echo-midtrans/pkg/domain/transaction"
)

type TransactionUseCase struct {
	PaymentService payment.UseCase
	CampaignRepo   campaign.Repository
	DBRepo         transaction.Repository
}

func NewTransactionUseCase(dbrepo transaction.Repository, paymentService payment.UseCase, campaignRepo campaign.Repository) *TransactionUseCase {
	return &TransactionUseCase{
		DBRepo:         dbrepo,
		CampaignRepo:   campaignRepo,
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
	_, err = uc.DBRepo.Update(ctx, newTransaction)
	if err != nil {
		return nil, err
	}

	return newTransaction, nil
}

func (uc *TransactionUseCase) ProcessPayment(ctx context.Context, req *transaction.TransactionNotificationRequest) error {
	transactionID := req.OrderID

	transaction, err := uc.DBRepo.GetByTransactionID(ctx, transactionID)
	if err != nil {
		return err
	}

	if req.PaymentType == "credit_card" && req.TransactionStatus == "capture" && req.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if req.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if req.TransactionStatus == "deny" || req.TransactionStatus == "expire" || req.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := uc.DBRepo.Update(ctx, transaction)
	if err != nil {
		return err
	}

	campaign, err := uc.CampaignRepo.FindByCampaignID(ctx, updatedTransaction.CampaignID)
	if err != nil {
		return nil
	}

	if updatedTransaction.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount

		err := uc.CampaignRepo.Update(ctx, campaign)
		if err != nil {
			return err
		}
	}

	return nil
}
