package usecase

import (
	"context"
	"echo-midtrans/pkg/domain/campaign"
	"io"
	"mime/multipart"
	"os"
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

func (uc *CampaignsUseCase) GetCampaignDetails(ctx context.Context, req *campaign.Campaign) (*campaign.Campaign, error) {
	campaign, err := uc.DBRepo.FindByCampaignID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	res, err := uc.DBRepo.FindCampaignImage(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	campaign.CampaignImages = res.CampaignImages

	return campaign, nil
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

func (uc *CampaignsUseCase) CreateCampaignImage(ctx context.Context, req *campaign.CampaignImage, fileLocation string, image multipart.File) (*campaign.CampaignImage, error) {
	var (
		response campaign.CampaignImage
	)

	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, image); err != nil {
		return nil, err
	}

	response.CampaignID = req.CampaignID
	response.IsPrimary = req.IsPrimary
	response.FileName = fileLocation

	res, err := uc.DBRepo.CreateImage(ctx, &response)
	if err != nil {
		return nil, err
	}

	return res, nil
}
