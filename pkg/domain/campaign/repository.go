package campaign

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]Campaign, error)
	FindWithLimit(ctx context.Context, limit int) ([]Campaign, error)
	FindByUserID(ctx context.Context, userID uint) ([]Campaign, error)
	FindByCampaignID(ctx context.Context, campaignID uint) (*Campaign, error)
	FindCampaignImage(ctx context.Context, campaignID uint) (*Campaign, error)
	FindBySlug(ctx context.Context, slug string) (*Campaign, error)
	Create(ctx context.Context, req *Campaign) (*Campaign, error)
	Update(ctx context.Context, req *Campaign) error
	Delete(ctx context.Context, req *Campaign) error
	CreateImage(ctx context.Context, req *CampaignImage) (*CampaignImage, error)
}
