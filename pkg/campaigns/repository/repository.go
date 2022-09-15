package repository

import (
	"context"
	"echo-midtrans/pkg/domain/campaign"
	"echo-midtrans/pkg/domain/common"
	"errors"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type CampaignsDBRepository struct {
	DB *gorm.DB
}

func NewCampaignsDBRepository(db *gorm.DB) *CampaignsDBRepository {
	return &CampaignsDBRepository{DB: db}
}

func (r *CampaignsDBRepository) FindAll(ctx context.Context) ([]campaign.Campaign, error) {
	var (
		res []campaign.Campaign
	)

	err := r.DB.WithContext(ctx).Find(&res).Order("created_at DESC").Error
	if err != nil {
		return nil, fmt.Errorf("CampaignsDBRepository.FindAll: %w", err)
	}
	if res == nil {
		return nil, common.ErrRecordNotFound
	}
	return res, nil
}

func (r *CampaignsDBRepository) FindWithLimit(ctx context.Context, limit int) ([]campaign.Campaign, error) {
	return nil, nil
}

func (r *CampaignsDBRepository) FindByUserID(ctx context.Context, userID uint) ([]campaign.Campaign, error) {
	return nil, nil
}

func (r *CampaignsDBRepository) FindByID(ctx context.Context, ID uint) (*campaign.Campaign, error) {
	return nil, nil
}

func (r *CampaignsDBRepository) FindBySlug(ctx context.Context, slug string) (*campaign.Campaign, error) {
	return nil, nil
}

func (r *CampaignsDBRepository) Create(ctx context.Context, req *campaign.Campaign) (*campaign.Campaign, error) {
	var (
		res campaign.Campaign
	)
	err := r.DB.WithContext(ctx).First(&res, "campaign_name = ?", req.CampaignName).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("CampaignsDBRepository.Create: %w", err)
	}

	// DeepEqual is used to check two interfaces are equal or not
	if !reflect.DeepEqual(res, campaign.Campaign{}) {
		return nil, common.ErrCampaignAlreadyCreated
	}

	err = r.DB.WithContext(ctx).Save(req).Error
	if err != nil {
		return nil, fmt.Errorf("CampaignsDBRepository.Create: %w", err)
	}

	return req, nil
}

func (r *CampaignsDBRepository) Update(ctx context.Context, req *campaign.Campaign) error {
	err := r.DB.WithContext(ctx).Updates(req).Error
	if err != nil {
		return fmt.Errorf("CampaignsDBRepository.Update: %w", err)
	}
	return nil
}
func (r *CampaignsDBRepository) Delete(ctx context.Context, req *campaign.Campaign) error {
	err := r.DB.WithContext(ctx).Delete(req).Error
	if err != nil {
		return fmt.Errorf("CampaignsDBRepository.DeleteByID: %w", err)
	}
	return nil
}

func (r *CampaignsDBRepository) CreateImage(ctx context.Context, req *campaign.CampaignImage) (*campaign.CampaignImage, error) {
	return nil, nil
}
