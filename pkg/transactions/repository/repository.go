package repository

import (
	"context"
	"echo-midtrans/pkg/domain/common"
	"echo-midtrans/pkg/domain/transaction"
	"fmt"

	"gorm.io/gorm"
)

type TransactionDBRepository struct {
	DB *gorm.DB
}

func NewTransactionDBRepository(db *gorm.DB) *TransactionDBRepository {
	return &TransactionDBRepository{DB: db}
}

func (r *TransactionDBRepository) GetByTransactionID(ctx context.Context, transactionID uint) (*transaction.Transaction, error) {
	var (
		res transaction.Transaction
	)

	err := r.DB.WithContext(ctx).Find(&res, "id = ?", transactionID).Order("created_at ASC").Error
	if err != nil {
		return nil, fmt.Errorf("TransactionDBRepository.FindAll: %w", err)
	}
	return &res, nil
}

func (r *TransactionDBRepository) FindAll(ctx context.Context) ([]transaction.Transaction, error) {
	var (
		res []transaction.Transaction
	)

	err := r.DB.WithContext(ctx).Find(&res).Order("created_at DESC").Error
	if err != nil {
		return nil, fmt.Errorf("TransactionDBRepository.FindAll: %w", err)
	}
	if res == nil {
		return nil, common.ErrRecordNotFound
	}
	return res, nil
}

func (r *TransactionDBRepository) FindByCampaignID(ctx context.Context, campaignID uint) ([]transaction.Transaction, error) {
	var (
		res []transaction.Transaction
	)

	err := r.DB.WithContext(ctx).Find(&res, "campaign_id = ?", campaignID).Order("created_at DESC").Error
	if err != nil {
		return nil, fmt.Errorf("TransactionDBRepository.FindByCampaignID: %w", err)
	}
	if res == nil {
		return nil, common.ErrRecordNotFound
	}
	return res, nil
}

func (r *TransactionDBRepository) FindByUserID(ctx context.Context, userID uint) ([]transaction.Transaction, error) {
	var (
		res []transaction.Transaction
	)

	err := r.DB.WithContext(ctx).Find(&res, "user_id = ?", userID).Order("created_at DESC").Error
	if err != nil {
		return nil, fmt.Errorf("TransactionDBRepository.FindByUserID: %w", err)
	}
	if res == nil {
		return nil, common.ErrRecordNotFound
	}
	return res, nil
}

func (r *TransactionDBRepository) Create(ctx context.Context, req *transaction.Transaction) (*transaction.Transaction, error) {
	err := r.DB.WithContext(ctx).Save(req).Error
	if err != nil {
		return nil, fmt.Errorf("TransactionDBRepository.Create: %w", err)
	}
	return req, nil
}

func (r *TransactionDBRepository) Update(ctx context.Context, req *transaction.Transaction) (*transaction.Transaction, error) {
	err := r.DB.WithContext(ctx).Updates(req).Error
	if err != nil {
		return nil, fmt.Errorf("TransactionDBRepository.Update: %w", err)
	}
	return req, nil
}
