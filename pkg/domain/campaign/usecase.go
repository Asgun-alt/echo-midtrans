package campaign

import (
	"context"
	"mime/multipart"
)

type UseCase interface {
	GetCampaigns(ctx context.Context) ([]Campaign, error)
	GetCampaignDetails(ctx context.Context, req *Campaign) (*Campaign, error)
	CreateCampaign(ctx context.Context, req *Campaign) (*Campaign, error)
	UpdateCampaign(ctx context.Context, req *Campaign) error
	DeleteCampaign(ctx context.Context, req *Campaign) error
	CreateCampaignImage(ctx context.Context, req *CampaignImage, fileLocation string, image multipart.File) (*CampaignImage, error)
}
