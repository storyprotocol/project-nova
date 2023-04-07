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
	Address:           "0x624bdd3b5d4f67fef15880a8f2cb0a0703d6ec0c",
	Name:              "Forced Offline",
	VaultAddress:      "0x624bdd3b5d4f67fef15880a8f2cb0a0703d6ec0c",
	CharacterRegistry: "0x402d3e9c136dfefaae3b04ea086632688951eccd",
	CharacterContracts: []*CharacterCollection{
		{
			Address: "0xf90bf1f50b71baae6bedcaf92a3a63d97200382c",
			Name:    "Force Offline Main",
		},
		{
			Address: "0xd7431ef1cbd4b5bcc568def48d4480f4e08d2224",
			Name:    "Forced offline FanFic",
		},
	},
	StoryRegistry: "0x13687aafb5accecb358f207f58614e01b09dec72",
	StoryContracts: []*StoryCollection{
		{
			Address: "0x6a84bcceebcd42ee0d73fa95f75862f635fd96ca",
			IsCanon: true,
		},
		{
			Address: "0x72f822e6b4a752b11b4275aad841b86a3f5266ab",
			IsCanon: false,
		},
	},
}

var FranchiseMap = map[string]*Franchise{
	"0x624bdd3b5d4f67fef15880a8f2cb0a0703d6ec0c": SingleFranchise,
}

var StoryContractMap = map[string]*StoryCollection{
	"0x6a84bcceebcd42ee0d73fa95f75862f635fd96ca": {
		Address: "0x6a84bcceebcd42ee0d73fa95f75862f635fd96ca",
		IsCanon: true,
	},
	"0x72f822e6b4a752b11b4275aad841b86a3f5266ab": {
		Address: "0x72f822e6b4a752b11b4275aad841b86a3f5266ab",
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
	Description *string `json:"description"`
	Image       *string `json:"image"`
}
