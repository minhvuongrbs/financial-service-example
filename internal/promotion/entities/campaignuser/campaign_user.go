package campaignuser

type CampaignUser struct {
	CampaignId int64
	UserId     int64
}

func NewCampaignUser(campaignId, userId int64) *CampaignUser {
	return &CampaignUser{
		CampaignId: campaignId,
		UserId:     userId,
	}
}
