package entity

type Franchise struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	OwnerAddress string          `json:"ownerAddress"`
	ImageUrl     *string         `json:"imageUrl"`
	BannerUrl    *string         `json:"bannerUrl"`
	Txhash       *string         `json:"txhash"`
	Metrics      FranchiseMetric `json:"metrics"`
}

type FranchiseMetric struct {
	StoryCount   int64   `json:"storyCount"`
	LicenseCount int64   `json:"licenseCount"`
	Revenue      float64 `json:"revenue"`
}
