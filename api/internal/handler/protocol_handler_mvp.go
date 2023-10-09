package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/api/internal/service"
	xhttp "github.com/project-nova/backend/pkg/http"
)

// GET /franchise
func NewGetFranchisesHandler(graphService service.TheGraphService, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.GetFranchisesResponseMVP{
			Data: []entity.FranchiseMVP{
				{
					ID:           "7",
					Name:         "Star Wars",
					OwnerAddress: "0x4f9693ac46f2c7e2f48dd14d8fe1ab44192cd57d",
					TokenUri:     "https://arweave.net/dnFJl1v8kgOx_6Z0gEsBce3D56cMP4-lxAcFqSsL0_w",
					TxHash:       "0xc80c23b7992cc94a271d1a56280ccc16a8f78a6d63ee34efdc35d8ffc71eda58",
				},
				{
					ID:           "10",
					Name:         "Divine Anarchy",
					OwnerAddress: "0x4f9693ac46f2c7e2f48dd14d8fe1ab44192cd57d",
					TokenUri:     "https://arweave.net/uwhn8-mPXjORkqHLBxPNe-rbUh_k9OV8OCCWkyNCijI",
					TxHash:       "0x95ed7d1d6fa5db08be22c7e58727729bf582a25ef82e94530a5e4cdf2d934a95",
				},
			},
		})
	}
}

// GET /franchise/:franchiseId
func NewGetFranchiseHandler(graphService service.TheGraphService, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.GetFranchiseResponseMVP{
			Data: entity.FranchiseMVP{
				ID:           "7",
				Name:         "Star Wars",
				OwnerAddress: "0x4f9693ac46f2c7e2f48dd14d8fe1ab44192cd57d",
				TokenUri:     "https://arweave.net/dnFJl1v8kgOx_6Z0gEsBce3D56cMP4-lxAcFqSsL0_w",
				TxHash:       "0xc80c23b7992cc94a271d1a56280ccc16a8f78a6d63ee34efdc35d8ffc71eda58",
			},
		})
	}
}

// GET /ipasset
func NewGetIpAssetsHandler(graphService service.TheGraphService, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.GetIpAssetsResponseMVP{
			Data: []entity.IpAssetMVP{
				{
					ID:           "1",
					FranchiseId:  "7",
					Type:         entity.IpAssetTypes.Story,
					Name:         "The Empire Strikes Back",
					OwnerAddress: "0x4f9693ac46f2c7e2f48dd14d8fe1ab44192cd57d",
					TokenUri:     "https://arweave.net/R7-xPDAMqOhUSw3CM_UwXI7zdpQkzCCCUq3smzxyAaU",
					TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
				},
				{
					ID:           "1000000000003",
					FranchiseId:  "7",
					Type:         entity.IpAssetTypes.Character,
					Name:         "Darth Vader",
					OwnerAddress: "0x69693d3234512ce8bfe17f7cb6c187dea51d0562",
					TokenUri:     "https://arweave.net/YiXTj7ps-hgV43JIAZOh7RM4xb_OQ918PAAiOV9RiLw",
					TxHash:       "0x1161c3b57913cbfa504220eb75dfcfa7bd5477d45ddfd8fb5fda44eae950186b",
				},
			},
		})
	}
}

// GET /ipasset/:ipAssetId
func NewGetIpAssetHandler(graphService service.TheGraphService, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.GetIpAssetResponseMVP{
			Data: entity.IpAssetMVP{
				ID:           "1",
				FranchiseId:  "7",
				Type:         entity.IpAssetTypes.Story,
				Name:         "The Empire Strikes Back",
				OwnerAddress: "0x4f9693ac46f2c7e2f48dd14d8fe1ab44192cd57d",
				TokenUri:     "https://arweave.net/R7-xPDAMqOhUSw3CM_UwXI7zdpQkzCCCUq3smzxyAaU",
				TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
			},
		})
	}
}

// GET /license
func NewGetLicensesHandler(graphService service.TheGraphService) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.GetLicensesResponseMVP{
			Data: []entity.LicenseMVP{
				{
					ID:           "50",
					IpAssetId:    "5",
					FranchiseId:  "7",
					OwnerAddress: "0xd84316a1b6f40902c17b8177854cdaeb3c957daf",
					Uri:          "https://arweave.net/R7-xPDAMqOhUSw3CM_UwXI7zdpQkzCCCUq3smzxyAaU",
					TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
				},
				{
					ID:           "49",
					IpAssetId:    "5",
					FranchiseId:  "7",
					OwnerAddress: "0xd84316a1b6f40902c17b8177854cdaeb3c957daf",
					Uri:          "https://arweave.net/R7-xPDAMqOhUSw3CM_UwXI7zdpQkzCCCUq3smzxyAaU",
					TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
				},
			},
		})
	}
}

// GET /license/:licenseId
func NewGetLicenseHandler(graphService service.TheGraphService) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.GetLicenseResponseMVP{
			Data: entity.LicenseMVP{
				ID:           "49",
				IpAssetId:    "5",
				FranchiseId:  "7",
				OwnerAddress: "0xd84316a1b6f40902c17b8177854cdaeb3c957daf",
				Uri:          "https://arweave.net/R7-xPDAMqOhUSw3CM_UwXI7zdpQkzCCCUq3smzxyAaU",
				TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
			},
		})
	}
}

// GET /collection
func NewGetCollectionsHandler(graphService service.TheGraphService) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.GetCollectionsResponseMVP{
			Data: []entity.CollectionMVP{
				{
					FranchiseId:    "7",
					IpAssetId:      "1",
					TotalCollected: 3,
				},
				{
					FranchiseId:    "7",
					IpAssetId:      "2",
					TotalCollected: 5,
				},
			},
		})
	}
}

// GET /transaction
func NewGetTransactionsHandler(graphService service.TheGraphService) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.GetTransactionsResponseMVP{
			Data: []entity.TransactionMVP{
				{
					ID:             "1",
					FranchiseId:    "7",
					CreatedAt:      "0001-01-01T00:00:00Z",
					ResourceType:   entity.ResourceTypes.IpAsset,
					ResourceId:     "1",
					CreatorAddress: "0x4f9693ac46f2c7e2f48dd14d8fe1ab44192cd57d",
					TxHash:         "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
				},
				{
					ID:             "2",
					FranchiseId:    "7",
					CreatedAt:      "0001-01-01T00:00:00Z",
					ResourceType:   entity.ResourceTypes.License,
					ResourceId:     "50",
					CreatorAddress: "0xd84316a1b6f40902c17b8177854cdaeb3c957daf",
					TxHash:         "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
				},
			},
		})
	}
}

// GET /transaction/:transactionId
func NewGetTransactionHandler(graphService service.TheGraphService) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.GetTransactionResponseMVP{
			Data: entity.TransactionMVP{
				ID:             "1",
				FranchiseId:    "7",
				CreatedAt:      "0001-01-01T00:00:00Z",
				ResourceType:   entity.ResourceTypes.IpAsset,
				ResourceId:     "1",
				CreatorAddress: "0x4f9693ac46f2c7e2f48dd14d8fe1ab44192cd57d",
				TxHash:         "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
			},
		})
	}
}
