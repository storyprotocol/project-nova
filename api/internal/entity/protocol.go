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
	Address:           "0x7c50c053f1d7f889c528dab717760a155bf3f543",
	Name:              "Forced Offline",
	VaultAddress:      "0x7c50c053f1d7f889c528dab717760a155bf3f543",
	CharacterRegistry: "0x29943be4cc35e5517c10e2917cecab8d96b53eea",
	CharacterContracts: []*CharacterCollection{
		{
			Address: "0xc6dc20dd779136555f6933a90bfd8defe0f6c8de",
			Name:    "Force Offline Main",
		},
		{
			Address: "0x31ed93fd8c4f656fa3e32693b3b8403687bc5503",
			Name:    "Forced offline FanFic",
		},
	},
	StoryRegistry: "0xbbbe9a2b68f8d2262ff608b615c2958df8b55f69",
	StoryContracts: []*StoryCollection{
		{
			Address: "0x6ff9cdefdaa8f58de2793b73a51cc6eb20f518da",
			IsCanon: true,
		},
		{
			Address: "0x5bcba74072a61c16ef19d966d0ed8c936131c3f0",
			IsCanon: false,
		},
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

var FranchiseMap = map[string]*Franchise{
	"0x7c50c053f1d7f889c528dab717760a155bf3f543": SingleFranchise,
}

var StoryContractMap = map[string]*StoryCollection{
	"0x6ff9cdefdaa8f58de2793b73a51cc6eb20f518da": {
		Address: "0x6ff9cdefdaa8f58de2793b73a51cc6eb20f518da",
		IsCanon: true,
	},
	"0x5bcba74072a61c16ef19d966d0ed8c936131c3f0": {
		Address: "0x5bcba74072a61c16ef19d966d0ed8c936131c3f0",
		IsCanon: false,
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
	Description *string     `json:"decription"`
	Image       *string     `json:"image"`
	Author      *string     `json:"author"`
	Attributes  []*NftTrait `json:"attributes"`
}

type StoryNftOnchainMeta struct {
	Name        *string `json:"name"`
	Description *string `json:"decription"`
	Image       *string `json:"image"`
}
