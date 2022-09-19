package campaign

type Response struct {
	ID             uint
	CampaignName   string          `json:"campaign_name"`
	Description    string          `json:"description"`
	Perks          string          `json:"perks"`
	BackerCount    int             `json:"backer_count"`
	GoalAmount     int             `json:"goal_amount"`
	CurrentAmount  int             `json:"current_amount"`
	Slug           string          `json:"slug"`
	CampaignImages []CampaignImage `json:"campaign_images"`
}
