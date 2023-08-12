package entity

type Franchise struct {
	ID           string          `json:"franchiseId"`
	Name         string          `json:"franchiseName"`
	OwnerAddress string          `json:"ownerAddress"`
	TokenUri     string          `json:"tokenUri"`
	ImageUrl     *string         `json:"image"`
	BannerUrl    *string         `json:"banner"`
	Txhash       *string         `json:"txhash"`
	Metrics      FranchiseMetric `json:"metrics"`
}

type FranchiseMetric struct {
	StoryCount   int64  `json:"storyCount"`
	LicenseCount int64  `json:"licenseCount"`
	Revenue      string `json:"revenue"`
}

type FranchiseMetadata struct {
	ImageUrl  *string `json:"imageUrl"`
	BannerUrl *string `json:"bannerUrl"`
}

type FranchisesTheGraphResponse struct {
	FranchisesRegistered []*FranchiseTheGraph `json:"franchiseRegistereds"`
}

type FranchiseTheGraph struct {
	ID              string `json:"id"`
	FranchiseId     string `json:"franchiseId"`
	Owner           string `json:"owner"`
	IpAssetRegistry string `json:"ipAssetRegistry"`
	Name            string `json:"name"`
	TokenURI        string `json:"tokenURI"`
}

func (f *FranchiseTheGraph) ToFranchise() *Franchise {
	franchise := &Franchise{
		ID:           f.FranchiseId,
		OwnerAddress: f.Owner,
		Name:         f.Name,
		TokenUri:     f.TokenURI,
	}
	return franchise
}
