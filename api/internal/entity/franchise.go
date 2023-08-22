package entity

type Franchise struct {
	ID           string          `json:"franchiseId,omitempty"`
	Name         string          `json:"franchiseName,omitempty"`
	OwnerAddress string          `json:"ownerAddress,omitempty"`
	TokenUri     string          `json:"tokenUri,omitempty"`
	ImageUrl     *string         `json:"image,omitempty"`
	BannerUrl    *string         `json:"banner,omitempty"`
	Txhash       *string         `json:"txhash,omitempty"`
	Metrics      FranchiseMetric `json:"metrics,omitempty"`
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
	TxHash          string `json:"transactionHash"`
}

func (f *FranchiseTheGraph) ToFranchise() *Franchise {
	franchise := &Franchise{
		ID:           f.FranchiseId,
		OwnerAddress: f.Owner,
		Name:         f.Name,
		TokenUri:     f.TokenURI,
		Txhash:       &f.TxHash,
	}
	return franchise
}
