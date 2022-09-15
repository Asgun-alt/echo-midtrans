package usecase

import (
	"context"
	"echo-midtrans/pkg/domain/campaign"
)

type CampaignsUseCase struct {
	DBRepo campaign.Repository
}

func NewCampaignsUseCase(dbrepo campaign.Repository) *CampaignsUseCase {
	return &CampaignsUseCase{DBRepo: dbrepo}
}

func (uc *CampaignsUseCase) GetCampaigns(ctx context.Context) ([]campaign.Campaign, error) {
	campaigns, err := uc.DBRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return campaigns, nil
}

func (uc *CampaignsUseCase) GetCampaignByID(ctx context.Context, req *campaign.Campaign) (*campaign.Campaign, error) {
	return nil, nil
}

func (uc *CampaignsUseCase) CreateCampaign(ctx context.Context, req *campaign.Campaign) (*campaign.Campaign, error) {
	var (
		err error
	)

	req, err = uc.DBRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (uc *CampaignsUseCase) UpdateCampaign(ctx context.Context, req *campaign.Campaign) error {
	err := uc.DBRepo.Update(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (uc *CampaignsUseCase) DeleteCampaign(ctx context.Context, req *campaign.Campaign) error {
	err := uc.DBRepo.Delete(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (uc *CampaignsUseCase) CreateCampaignImage(ctx context.Context, req *campaign.CampaignImage, fileLocation string) (*campaign.CampaignImage, error) {
	return nil, nil
}
