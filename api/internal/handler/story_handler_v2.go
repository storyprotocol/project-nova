package handler

import (
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/api/internal/repository"
	"github.com/project-nova/backend/pkg/gateway"
	xhttp "github.com/project-nova/backend/pkg/http"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/utils"
	"github.com/project-nova/backend/proto/v1/web3_gateway"
)

// GET /franchise/:franchiseId/authors/:authorId
func NewGetAuthorDetailHandlerV2() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.Author{
			ID:           "1",
			Name:         "Charles Dickens",
			Introduction: "Charles John Huffam Dickens FRSA was an English writer and social critic. He created some of the world's best-known fictional characters and is regarded by many as the greatest novelist of the Victorian era. His works enjoyed unprecedented popularity during his lifetime, and by the 20th century, critics and scholars had recognised him as a literary genius. His novels and short stories are still widely read today.",
			Image:        "https://upload.wikimedia.org/wikipedia/commons/thumb/a/aa/Dickens_Gurney_head.jpg/220px-Dickens_Gurney_head.jpg",
			Addresses: []entity.ContentAddress{
				{
					Type:    "ethereum",
					Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
				},
			},
			Stories: []entity.StoryV2{
				{
					ID:           "1",
					Title:        "A Tale of Two Cities",
					Introduction: "A novel by Charles Dickens, is set against the backdrop of the French Revolution and explores themes of resurrection, sacrifice, and the struggle between tyranny and freedom. The story unfolds in two cities: London and Paris. It opens with the famous lines, 'It was the best of times, it was the worst of times,' reflecting the contradictions of the era. Central characters include Charles Darnay, a French aristocrat who renounces his heritage for a modest life in England, and Sydney Carton, a dissipated English lawyer whose love for Darnay's wife, Lucie Manette, inspires profound self-sacrifice. Lucie's father, Dr. Manette, unjustly imprisoned in the Bastille for 18 years, embodies the trauma of the old regime's brutality. The novel is a rich tapestry of love, honor, and redemption, played out against the grim realities of the Revolution.",
					Image:        "https://cdn.discordapp.com/attachments/1091875369864413214/1173056228990136340/Screenshot_2023-11-11_at_4.26.16_PM.png?ex=656290bf&is=65501bbf&hm=5f7a5f69e14c29c0c9f9172243b9bda49d03ea05e74c7c0625eddd030ae9eada&",
					WordCount:    1000,
					Addresses: []entity.ContentAddress{
						{
							Type:    "arweave",
							Address: "https://arweave.net/1",
						},
						{
							Type:    "owner",
							Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
						},
						{
							Type:    "nft",
							Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
						},
					},
				},
			},
		})
	}
}

// GET /franchise/:franchiseId/assets/:assetId
func NewGetAssetDetailHandlerV2() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.Asset{
			ID:    "1",
			Type:  "character",
			Name:  "Charles Darnay",
			Image: "https://static.wikia.nocookie.net/classical-literature/images/5/50/Darnay1958.png/revision/latest?cb=20230222192722",
			Addresses: []entity.ContentAddress{
				{
					Type:    "arweave",
					Address: "https://arweave.net/1",
				},
				{
					Type:    "ethereum",
					Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
				},
				{
					Type:    "nft",
					Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
				},
			},
			Author: entity.Author{
				ID:           "1",
				Name:         "Charles Dickens",
				Introduction: "Charles John Huffam Dickens FRSA was an English writer and social critic. He created some of the world's best-known fictional characters and is regarded by many as the greatest novelist of the Victorian era. His works enjoyed unprecedented popularity during his lifetime, and by the 20th century, critics and scholars had recognised him as a literary genius. His novels and short stories are still widely read today.",
				Image:        "https://upload.wikimedia.org/wikipedia/commons/thumb/a/aa/Dickens_Gurney_head.jpg/220px-Dickens_Gurney_head.jpg",
				Addresses: []entity.ContentAddress{
					{
						Type:    "ethereum",
						Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
					},
				},
			},
			Stories: []entity.StoryV2{
				{
					ID:           "1",
					Title:        "A Tale of Two Cities",
					Introduction: "A novel by Charles Dickens, is set against the backdrop of the French Revolution and explores themes of resurrection, sacrifice, and the struggle between tyranny and freedom. The story unfolds in two cities: London and Paris. It opens with the famous lines, 'It was the best of times, it was the worst of times,' reflecting the contradictions of the era. Central characters include Charles Darnay, a French aristocrat who renounces his heritage for a modest life in England, and Sydney Carton, a dissipated English lawyer whose love for Darnay's wife, Lucie Manette, inspires profound self-sacrifice. Lucie's father, Dr. Manette, unjustly imprisoned in the Bastille for 18 years, embodies the trauma of the old regime's brutality. The novel is a rich tapestry of love, honor, and redemption, played out against the grim realities of the Revolution.",
					Image:        "https://cdn.discordapp.com/attachments/1091875369864413214/1173056228990136340/Screenshot_2023-11-11_at_4.26.16_PM.png?ex=656290bf&is=65501bbf&hm=5f7a5f69e14c29c0c9f9172243b9bda49d03ea05e74c7c0625eddd030ae9eada&",
					WordCount:    1000,
					Addresses: []entity.ContentAddress{
						{
							Type:    "arweave",
							Address: "https://arweave.net/1",
						},
						{
							Type:    "owner",
							Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
						},
						{
							Type:    "nft",
							Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
						},
					},
				},
			},
		})
	}
}

// GET /franchise/:franchiseId/stories/:storyId
func NewGetStoryDetailHandlerV2() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.StoryV2{
			ID:           "1",
			Title:        "A Tale of Two Cities",
			Introduction: "A novel by Charles Dickens, is set against the backdrop of the French Revolution and explores themes of resurrection, sacrifice, and the struggle between tyranny and freedom. The story unfolds in two cities: London and Paris. It opens with the famous lines, 'It was the best of times, it was the worst of times,' reflecting the contradictions of the era. Central characters include Charles Darnay, a French aristocrat who renounces his heritage for a modest life in England, and Sydney Carton, a dissipated English lawyer whose love for Darnay's wife, Lucie Manette, inspires profound self-sacrifice. Lucie's father, Dr. Manette, unjustly imprisoned in the Bastille for 18 years, embodies the trauma of the old regime's brutality. The novel is a rich tapestry of love, honor, and redemption, played out against the grim realities of the Revolution.",
			Image:        "https://cdn.discordapp.com/attachments/1091875369864413214/1173056228990136340/Screenshot_2023-11-11_at_4.26.16_PM.png?ex=656290bf&is=65501bbf&hm=5f7a5f69e14c29c0c9f9172243b9bda49d03ea05e74c7c0625eddd030ae9eada&",
			WordCount:    1000,
			Addresses: []entity.ContentAddress{
				{
					Type:    "arweave",
					Address: "https://arweave.net/1",
				},
				{
					Type:    "owner",
					Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
				},
				{
					Type:    "nft",
					Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
				},
			},
			Author: entity.Author{
				ID:           "1",
				Name:         "Charles Dickens",
				Introduction: "Charles John Huffam Dickens FRSA was an English writer and social critic. He created some of the world's best-known fictional characters and is regarded by many as the greatest novelist of the Victorian era. His works enjoyed unprecedented popularity during his lifetime, and by the 20th century, critics and scholars had recognised him as a literary genius. His novels and short stories are still widely read today.",
				Image:        "https://upload.wikimedia.org/wikipedia/commons/thumb/a/aa/Dickens_Gurney_head.jpg/220px-Dickens_Gurney_head.jpg",
				Addresses: []entity.ContentAddress{
					{
						Type:    "ethereum",
						Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
					},
				},
			},
			Assets: []entity.Asset{
				{
					ID:    "1",
					Type:  "character",
					Name:  "Charles Darnay",
					Image: "https://static.wikia.nocookie.net/classical-literature/images/5/50/Darnay1958.png/revision/latest?cb=20230222192722",
					Addresses: []entity.ContentAddress{
						{
							Type:    "arweave",
							Address: "https://arweave.net/1",
						},
						{
							Type:    "ethereum",
							Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
						},
						{
							Type:    "nft",
							Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
						},
					},
				},
				{
					ID:    "2",
					Type:  "character",
					Name:  "Sydney Carton",
					Image: "https://static.wikia.nocookie.net/classical-literature/images/5/50/Darnay1958.png/revision/latest?cb=20230222192722",
					Addresses: []entity.ContentAddress{
						{
							Type:    "arweave",
							Address: "https://arweave.net/1",
						},
						{
							Type:    "ethereum",
							Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
						},
						{
							Type:    "nft",
							Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
						},
					},
				},
			},
			Chapters: []entity.Chapter{
				{
					ID:        "1",
					Title:     "Chapter 1",
					Content:   "It was the best of times, it was the worst of times, it was the age of wisdom, it was the age of foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of Darkness, it was the spring of hope, it was the winter of despair, we had everything before us, we had nothing before us, we were all going direct to Heaven, we were all going direct the other way—in short, the period was so far like the present period, that some of its noisiest authorities insisted on its being received, for good or for evil, in the superlative degree of comparison only.",
					WordCount: 1000,
				},
				{
					ID:        "2",
					Title:     "Chapter 2",
					Content:   "It was the best of times, it was the worst of times, it was the age of wisdom, it was the age of foolishness, it was the epoch of belief, it was the epoch of incredulity",
					WordCount: 1000,
				},
			},
		})
	}
}

// GET /franchise/:franchiseId/stories
func NewListFranchiseStoriesHandlerV2() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.StoriesV2{
			Stories: []*entity.StoryV2{
				{
					ID:           "1",
					Title:        "A Tale of Two Cities",
					Introduction: "A novel by Charles Dickens, is set against the backdrop of the French Revolution and explores themes of resurrection, sacrifice, and the struggle between tyranny and freedom. The story unfolds in two cities: London and Paris. It opens with the famous lines, 'It was the best of times, it was the worst of times,' reflecting the contradictions of the era. Central characters include Charles Darnay, a French aristocrat who renounces his heritage for a modest life in England, and Sydney Carton, a dissipated English lawyer whose love for Darnay's wife, Lucie Manette, inspires profound self-sacrifice. Lucie's father, Dr. Manette, unjustly imprisoned in the Bastille for 18 years, embodies the trauma of the old regime's brutality. The novel is a rich tapestry of love, honor, and redemption, played out against the grim realities of the Revolution.",
					Image:        "https://cdn.discordapp.com/attachments/1091875369864413214/1173056228990136340/Screenshot_2023-11-11_at_4.26.16_PM.png?ex=656290bf&is=65501bbf&hm=5f7a5f69e14c29c0c9f9172243b9bda49d03ea05e74c7c0625eddd030ae9eada&",
					WordCount:    1000,
					Addresses: []entity.ContentAddress{
						{
							Type:    "arweave",
							Address: "https://arweave.net/1",
						},
						{
							Type:    "ethereum",
							Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
						},
						{
							Type:    "nft",
							Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
						},
					},
					Author: entity.Author{
						ID:   "1",
						Name: "Charles Dickens",
					},
				},
				// do the same format but for another novel
				{
					ID:           "2",
					Title:        "The Great Gatsby",
					Introduction: "The Great Gatsby is a 1925 novel by American writer F. Scott Fitzgerald. Set in the Jazz Age on Long Island, the novel depicts narrator Nick Carraway's interactions with mysterious millionaire Jay Gatsby and Gatsby's obsession to reunite with his former lover, Daisy Buchanan.",
					Image:        "https://cdn.discordapp.com/attachments/1091875369864413214/1173056228990136340/Screenshot_2023-11-11_at_4.26.16_PM.png?ex=656290bf&is=65501bbf&hm=5f7a5f69e14c29c0c9f9172243b9bda49d03ea05e74c7c0625eddd030ae9eada&",
					WordCount:    1000,
					Addresses: []entity.ContentAddress{
						{
							Type:    "arweave",
							Address: "https://arweave.net/1",
						},
						{
							Type:    "owner",
							Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
						},
						{
							Type: "nft",

							Address: "0x8fbad875c3485f64bc1fad98595dec4f7b0f225d",
						},
					},

					Author: entity.Author{
						ID:   "2",
						Name: "F. Scott Fitzgerald",
					},
				},
			},
		})
	}
}

// GET /story/:franchiseId/:storyId/:chapterId
func NewGetStoryContentHandlerV2(
	contentRepo repository.ProtocolStoryContentRepository,
	httpClient xhttp.Client,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 1. verify addresses
		franchiseAddress, err := utils.SanitizeAddress(c.Param("franchiseId"))
		if err != nil {
			logger.Errorf("Invalid franchise address: %s", c.Param("franchiseId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise address"))
			return
		}

		collectionAddress, err := utils.SanitizeAddress(c.Param("storyId"))
		if err != nil {
			logger.Errorf("Invalid story address: %s", c.Param("storyId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid story address"))
			return
		}

		tokenId, err := strconv.Atoi(c.Param("chapterId"))
		if err != nil {
			logger.Errorf("Invalid chapter id: %s", c.Param("chapterId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid chapter id"))
			return
		}

		// 2. call db to get content uri
		content, err := contentRepo.GetContentByAddresses(franchiseAddress, collectionAddress, tokenId)
		if err != nil {
			logger.Errorf("Failed to read content from database: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		// 3. call the uri to get the content
		var result entity.ContentV2
		_, err = httpClient.Request("GET", *content.ContentUri, nil, &result)
		if err != nil {
			logger.Errorf("Failed to read content from remote storage: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

// GET /story/:franchiseId/:storyId/:chapterId
func NewAdminUploadStoryContentHandlerV2(
	contentRepo repository.ProtocolStoryContentRepository,
	web3Gateway gateway.Web3GatewayClient,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 1. verify addresses
		franchiseAddress, err := utils.SanitizeAddress(c.Param("franchiseId"))
		if err != nil {
			logger.Errorf("Invalid franchise address: %s", c.Param("franchiseId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise address"))
			return
		}

		collectionAddress, err := utils.SanitizeAddress(c.Param("storyId"))
		if err != nil {
			logger.Errorf("Invalid story address: %s", c.Param("storyId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid story address"))
			return
		}

		tokenId, err := strconv.Atoi(c.Param("chapterId"))
		if err != nil {
			logger.Errorf("Invalid chapter id: %s", c.Param("chapterId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid chapter id"))
			return
		}

		var requestBody struct {
			Content string `json:"content"`
		}
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
			return
		}

		// 2. call web3 gateway to upload content
		contentBase64 := base64.StdEncoding.EncodeToString([]byte(requestBody.Content))
		resp, err := web3Gateway.UploadContent(&web3_gateway.UploadContentReq{
			Storage:     web3_gateway.StorageType_ARWEAVE,
			Content:     []byte(contentBase64),
			ContentType: "text/markdown",
			Tags: []*web3_gateway.Tag{
				{
					Name:  "franchise",
					Value: franchiseAddress,
				},
				{
					Name:  "story",
					Value: collectionAddress,
				},
				{
					Name:  "chapter",
					Value: strconv.Itoa(tokenId),
				},
				{
					Name:  "Content-Type",
					Value: "application/json",
				},
			},
		})
		if err != nil {
			logger.Errorf("Failed to upload content to web3-gateway: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		// 3. store content uri and other data in db
		err = contentRepo.CreateContent(&entity.ProtocolStoryContentModel{
			ID:                uuid.New().String(),
			FranchiseAddress:  &franchiseAddress,
			CollectionAddress: &collectionAddress,
			TokenId:           &tokenId,
			ContentJson:       requestBody.Content,
			ContentUri:        &resp.ContentUrl,
		})
		if err != nil {
			logger.Errorf("Failed to create content in the database: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"contentUrl": resp.ContentUrl,
		})
	}
}
