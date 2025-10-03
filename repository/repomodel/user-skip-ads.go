package repomodel

type SkipAdsResult struct {
	UserID       string `gorm:"column:user_id" json:"user_id"`
	SkipAdsTotal int32  `gorm:"column:skip_ads_total" json:"skip_ads_total"`
}
