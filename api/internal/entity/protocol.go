package entity

type Franchise struct {
	Address            string                 `json:"address"`
	Name               string                 `json:"name"`
	VaultAddress       string                 `json:"vaultAddress"`
	CharacterContracts []*CharacterCollection `json:"characterContract"`
	StoryContracts     []*StoryCollection     `json:"storyContract"`
}

var SingleFranchise = &Franchise{
	Address:      "0xForcedOffline",
	Name:         "Forced Offline",
	VaultAddress: "0xForcedOfflineVault",
	CharacterContracts: []*CharacterCollection{
		{
			Address: "0xForcedOfflineCharacter",
			Name:    "Force Offline Main",
		},
		{
			Address: "0xForcedOfflineFanFicCharacter",
			Name:    "Forced offline FanFic",
		},
	},
	StoryContracts: []*StoryCollection{
		{
			Address: "0xForcedOfflineStory",
			IsCanon: true,
		},
		{
			Address: "0xForcedOfflineFanFic",
			IsCanon: false,
		},
	},
}

var Franchises = []*Franchise{
	SingleFranchise,
	{
		Address:      "0xWhiteFountain",
		Name:         "White Fountain",
		VaultAddress: "0xWhiteFountainVault",
		CharacterContracts: []*CharacterCollection{
			{
				Address: "0xWhiteFountainCharacter",
				Name:    "White Fountain Main",
			},
			{
				Address: "0xWhiteFountainFanFicCharacter",
				Name:    "White Fountain FanFic",
			},
		},
		StoryContracts: []*StoryCollection{
			{
				Address: "0xWhiteFountainStory",
				IsCanon: true,
			},
			{
				Address: "0xWhiteFountainFanfic",
				IsCanon: false,
			},
		},
	},
}

type CharacterCollection struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type StoryCollection struct {
	Address string `json:"address"`
	IsCanon bool   `json:"isCanon"`
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

var SingleCharacter = &Character{
	CollectionAddress: "0xForcedOfflineCharacter",
	TokenId:           1,
	AuthorAddress:     "0xJosephEvans",
	OwnerAddress:      "jason.eth",
	Name:              "Rayze",
	Description:       "Rayze is the main character of force offline",
	ImageUrl:          "https://ipfs.io/ipfs/QmNvRAgJbWgcCvY2cJgvqvfgu6SwRKr5zsbZufm3owpaEc/images/145.png",
	Traits: []*NftTraitResponse{
		{
			TraitType: "hairColor",
			Value:     "red",
		},
	},
}

var Characters = []*Character{
	SingleCharacter,
	{
		CollectionAddress: "0xForcedOfflineCharacter",
		TokenId:           2,
		AuthorAddress:     "0xJosephEvans",
		OwnerAddress:      "allen.eth",
		Name:              "Myza",
		Description:       "Myza is the girlfriend of Rayze",
		ImageUrl:          "https://ipfs.io/ipfs/QmNvRAgJbWgcCvY2cJgvqvfgu6SwRKr5zsbZufm3owpaEc/images/11.png",
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

var SingleStory = &Story{
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
}

var Stories = []*Story{
	SingleStory,
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
