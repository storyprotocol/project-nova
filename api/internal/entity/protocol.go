package entity

type Franchise struct {
	Address           string `json:"address"`
	Name              string `json:"name"`
	VaultAddress      string `json:"vaultAddress"`
	CharacterContract string `json:"characterContract"`
	StoryContract     string `json:"storyContract"`
}

var Franchises = []*Franchise{
	{
		Address:           "0xForcedOffline",
		Name:              "Forced Offline",
		VaultAddress:      "0xForcedOfflineVault",
		CharacterContract: "0xForcedOfflineCharacter",
		StoryContract:     "0xForcedOfflineStory",
	},
	{
		Address:           "0xWhiteFountain",
		Name:              "White Fountain",
		VaultAddress:      "0xWhiteFountainVault",
		CharacterContract: "0xWhiteFountainCharacter",
		StoryContract:     "0xWhiteFoundtainStory",
	},
}

type Character struct {
	CollectionAddress string              `json:"collectionAddress"`
	TokenId           int                 `json:"tokenId"`
	AuthorAddress     string              `json:"authorAddress"`
	OwnerAddress      string              `json:"ownerAddress"`
	Name              string              `json:"name"`
	Description       string              `json:"description"`
	ImageUrl          string              `json:"imageUrl"`
	Traits            []*NftTraitResponse `json:"traits"`
}

var Characters = []*Character{
	{
		CollectionAddress: "0xForcedOfflineCharacter",
		TokenId:           1,
		AuthorAddress:     "0xJosephEvans",
		OwnerAddress:      "jason.eth",
		Name:              "Rayze",
		Description:       "Rayze is the main character of force offline",
		Traits: []*NftTraitResponse{
			{
				TraitType: "hairColor",
				Value:     "red",
			},
		},
	},
	{
		CollectionAddress: "0xForcedOfflineCharacter",
		TokenId:           2,
		AuthorAddress:     "0xJosephEvans",
		OwnerAddress:      "allen.eth",
		Name:              "Myza",
		Description:       "Myza is the girlfriend of Rayze",
		Traits: []*NftTraitResponse{
			{
				TraitType: "hairColor",
				Value:     "purple",
			},
		},
	},
}

type Story struct {
	CollectionAddress string       `json:"collectionAddress"`
	TokenId           int          `json:"tokenId"`
	AuthorAddress     []string     `json:"authorAddress"`
	OwnerAddress      string       `json:"ownerAddress"`
	Title             string       `json:"title"`
	ContentUrl        string       `json:"contentUrl"`
	Characters        []*Character `json:"characters"`
	IsCanon           bool         `json:"isCanon"`
}

var Stories = []*Story{
	{
		CollectionAddress: "0xForcedOfflineStory",
		TokenId:           1,
		AuthorAddress:     []string{"0xJosephEvans", "jasonlevy.eth"},
		OwnerAddress:      "0xJosephEvans",
		Title:             "Forced Offline",
		ContentUrl:        "https://stag.api.storyprotocol.net/v1/story/1/1/",
		Characters: []*Character{
			{
				CollectionAddress: "0xForcedOfflineCharacter",
				TokenId:           1,
				AuthorAddress:     "0xJosephEvans",
				OwnerAddress:      "jason.eth",
				Name:              "Rayze",
				Description:       "Rayze is the main character of force offline",
				Traits: []*NftTraitResponse{
					{
						TraitType: "hairColor",
						Value:     "red",
					},
				},
			},
			{
				CollectionAddress: "0xForcedOfflineCharacter",
				TokenId:           2,
				AuthorAddress:     "0xJosephEvans",
				OwnerAddress:      "allen.eth",
				Name:              "Myza",
				Description:       "Myza is the girlfriend of Rayze",
				Traits: []*NftTraitResponse{
					{
						TraitType: "hairColor",
						Value:     "purple",
					},
				},
			},
		},
		IsCanon: true,
	},
}

type Collector struct {
	Address string
}

var Collectors = []*Collector{
	{
		Address: "allen.eth",
	},
	{
		Address: "leo.eth",
	},
}
