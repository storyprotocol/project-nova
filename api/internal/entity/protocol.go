package entity

type Franchise struct {
	Address            string                 `json:"address"`
	Name               string                 `json:"name"`
	VaultAddress       string                 `json:"vaultAddress"`
	CharacterRegistry  string                 `json:"characterRegistry"`
	CharacterContracts []*CharacterCollection `json:"characterContract"`
	StoryRegistry      string                 `json:"storyRegistry"`
	StoryContracts     []*StoryCollection     `json:"storyContract"`
}

var SingleFranchise = &Franchise{
	Address:           "0xd5022d17c9a7768110c47075430ea6284f7e089a",
	Name:              "Forced Offline",
	VaultAddress:      "0xd5022d17c9a7768110c47075430ea6284f7e089a",
	CharacterRegistry: "0x1d2660e16164c9ccbdd8ebd8b7b358492c14c8db",
	CharacterContracts: []*CharacterCollection{
		{
			Address: "0x1751ba3d2b2c928ecdc24925fd8123c913d72873",
			Name:    "Force Offline Main",
		},
		{
			Address: "0x3c0b968f786485d2f9df4390de6c3a2360d1c1ff",
			Name:    "Forced offline FanFic",
		},
	},
	StoryRegistry: "0xd834353e6a1dbfa94ce8ff3ef5ec9452b86d66eb",
	StoryContracts: []*StoryCollection{
		{
			Address: "0xa0d090dac571a8ef51d10ec93d84d72f1db9552a",
			IsCanon: true,
		},
		{
			Address: "0x9f349f397db15b1bc669c420e224f1c694b9d9c6",
			IsCanon: false,
		},
	},
}

var FranchiseMap = map[string]*Franchise{
	"0xd5022d17c9a7768110c47075430ea6284f7e089a": SingleFranchise,
}

var StoryContractMap = map[string]*StoryCollection{
	"0xa0d090dac571a8ef51d10ec93d84d72f1db9552a": {
		Address: "0xa0d090dac571a8ef51d10ec93d84d72f1db9552a",
		IsCanon: true,
	},
	"0x9f349f397db15b1bc669c420e224f1c694b9d9c6": {
		Address: "0x9f349f397db15b1bc669c420e224f1c694b9d9c6",
		IsCanon: false,
	},
}

var Franchises = []*Franchise{
	SingleFranchise,
	/*
		{
			Address:           "0xWhiteFountain",
			Name:              "White Fountain",
			VaultAddress:      "0xWhiteFountainVault",
			CharacterRegistry: "",
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
			StoryRegistry: "",
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
	*/
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
	ImageUrl          string       `json:"imageUrl"`
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

type CharacterNftOnchainMeta struct {
	Name        *string     `json:"name"`
	Description *string     `json:"description"`
	Image       *string     `json:"image"`
	Author      *string     `json:"author"`
	Attributes  []*NftTrait `json:"attributes"`
}

type StoryNftOnchainMeta struct {
	Name        *string `json:"name"`
	Description *string `json:"decription"`
	Image       *string `json:"image"`
}
