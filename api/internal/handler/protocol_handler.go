package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/entity"
	v0alpha_entity "github.com/project-nova/backend/api/internal/entity/v0-alpha"
	v0alpha "github.com/project-nova/backend/api/internal/handler/v0-alpha"
	v0mvp "github.com/project-nova/backend/api/internal/handler/v0-mvp"
	"github.com/project-nova/backend/api/internal/service/thegraph"
	xhttp "github.com/project-nova/backend/pkg/http"
	"github.com/project-nova/backend/pkg/logger"
)

// GET /franchise
func NewGetFranchisesHandler(graphService thegraph.TheGraphServiceMvp, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchises, err := graphService.GetFranchises()
		if err != nil {
			logger.Errorf("Failed to get franchises: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, entity.GetFranchisesResponseMVP{
			Data: franchises,
		})
		/*
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
		*/
	}
}

// GET /franchise/:franchiseId
func NewGetFranchiseHandler(graphService thegraph.TheGraphServiceMvp, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId := c.Param("franchiseId")
		_, err := strconv.ParseInt(franchiseId, 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", franchiseId)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		franchise, err := graphService.GetFranchise(franchiseId)
		if err != nil {
			logger.Errorf("Failed to get franchise: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, entity.GetFranchiseResponseMVP{
			Data: franchise,
		})
		/*
			c.JSON(http.StatusOK, entity.GetFranchiseResponseMVP{
				Data: &entity.FranchiseMVP{
					ID:           "7",
					Name:         "Star Wars",
					OwnerAddress: "0x4f9693ac46f2c7e2f48dd14d8fe1ab44192cd57d",
					TokenUri:     "https://arweave.net/dnFJl1v8kgOx_6Z0gEsBce3D56cMP4-lxAcFqSsL0_w",
					TxHash:       "0xc80c23b7992cc94a271d1a56280ccc16a8f78a6d63ee34efdc35d8ffc71eda58",
				},
			})
		*/
	}
}

// POST /iporg
func NewListIpOrgsHandler(graphService thegraph.TheGraphServiceMvp, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, v0alpha_entity.ListIpOrgsResponse{
			IPOrgs: []*v0alpha_entity.IPOrg{
				{
					ID:          "7",
					Name:        "Star Wars",
					Symbol:      "STAR",
					Owner:       "0x4f9693ac46f2c7e2f48dd14d8fe1ab44192cd57d",
					MetadataUrl: "https://arweave.net/dnFJl1v8kgOx_6Z0gEsBce3D56cMP4-lxAcFqSsL0_w",
					TxHash:      "0xc80c23b7992cc94a271d1a56280ccc16a8f78a6d63ee34efdc35d8ffc71eda58",
					CreatedAt:   "2023-11-14T00:29:13Z",
				},
				{
					ID:          "10",
					Name:        "Divine Anarchy",
					Symbol:      "DA",
					Owner:       "0x4f9693ac46f2c7e2f48dd14d8fe1ab44192cd57d",
					MetadataUrl: "https://arweave.net/uwhn8-mPXjORkqHLBxPNe-rbUh_k9OV8OCCWkyNCijI",
					TxHash:      "0x95ed7d1d6fa5db08be22c7e58727729bf582a25ef82e94530a5e4cdf2d934a95",
					CreatedAt:   "2023-11-14T00:29:13Z",
				},
			},
		})
	}
}

// GET /iporg/:ipOrgId
func NewGetIpOrgHandler(graphService thegraph.TheGraphServiceMvp, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, v0alpha_entity.GetIpOrgResponse{
			IPOrg: &v0alpha_entity.IPOrg{
				ID:          "7",
				Name:        "Star Wars",
				Symbol:      "STAR",
				Owner:       "0x4f9693ac46f2c7e2f48dd14d8fe1ab44192cd57d",
				MetadataUrl: "https://arweave.net/dnFJl1v8kgOx_6Z0gEsBce3D56cMP4-lxAcFqSsL0_w",
				TxHash:      "0xc80c23b7992cc94a271d1a56280ccc16a8f78a6d63ee34efdc35d8ffc71eda58",
				CreatedAt:   "2023-11-14T00:29:13Z",
			},
		})
	}
}

// GET /ipasset
func NewGetIpAssetsHandler(graphService thegraph.TheGraphServiceMvp, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId := c.DefaultQuery("franchiseId", "")
		_, err := strconv.ParseInt(franchiseId, 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", franchiseId)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		ipAssets, err := graphService.GetIpAssets(franchiseId)
		if err != nil {
			logger.Errorf("Failed to get ip assets: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, entity.GetIpAssetsResponseMVP{
			Data: ipAssets,
		})
		/*
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
		*/
	}
}

// GET /ipasset/:ipAssetId
func NewGetIpAssetHandler(graphService thegraph.TheGraphServiceMvp, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		version := c.GetHeader(entity.HTTP_HEADER_VERSION)
		switch version {
		case entity.VERSION_V0_ALPHA:
			response, statusCode, err := v0alpha.HandleGetIpAsset(c)
			if err != nil {
				c.JSON(statusCode, ErrorMessage(err.Error()))
				return
			}
			c.JSON(statusCode, response)
		default:
			response, statusCode, err := v0mvp.HandleGetIpAsset(c, graphService)
			if err != nil {
				c.JSON(statusCode, ErrorMessage(err.Error()))
				return
			}
			c.JSON(statusCode, response)
		}

		/*
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
		*/
	}
}

// POST /ipasset
func NewListIpAssetsHandler(graphService thegraph.TheGraphServiceMvp, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, v0alpha_entity.ListIpAssetsResponse{
			IPAssets: []*v0alpha_entity.IPAsset{
				{
					ID:          "1",
					IPOrgId:     "7",
					Type:        v0alpha_entity.IpAssetTypes.Story,
					Name:        "The Empire Strikes Back",
					Owner:       "0x4f9693ac46f2c7e2f48dd14d8fe1ab44192cd57d",
					MetadataUrl: "https://arweave.net/R7-xPDAMqOhUSw3CM_UwXI7zdpQkzCCCUq3smzxyAaU",
					CreatedAt:   "2023-11-14T00:29:13Z",
					TxHash:      "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
				},
				{
					ID:          "1000000000003",
					IPOrgId:     "7",
					Type:        v0alpha_entity.IpAssetTypes.Character,
					Name:        "Darth Vader",
					Owner:       "0x69693d3234512ce8bfe17f7cb6c187dea51d0562",
					MetadataUrl: "https://arweave.net/YiXTj7ps-hgV43JIAZOh7RM4xb_OQ918PAAiOV9RiLw",
					CreatedAt:   "2023-11-14T00:29:13Z",
					TxHash:      "0x1161c3b57913cbfa504220eb75dfcfa7bd5477d45ddfd8fb5fda44eae950186b",
				},
			},
		})
	}
}

// GET /license
func NewGetLicensesHandler(graphService thegraph.TheGraphServiceMvp) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId := c.DefaultQuery("franchiseId", "")
		_, err := strconv.ParseInt(franchiseId, 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", franchiseId)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		ipAssetId := c.DefaultQuery("ipAssetId", "")
		_, err = strconv.ParseInt(ipAssetId, 10, 64)
		if err != nil {
			logger.Errorf("Invalid ip asset id: %s", ipAssetId)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid ip asset id"))
			return
		}

		licenses, err := graphService.GetLicenses(franchiseId, ipAssetId)
		if err != nil {
			logger.Errorf("Failed to get licenses: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, entity.GetLicensesResponseMVP{
			Data: licenses,
		})
		/*
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
		*/
	}
}

// GET /license/:licenseId
func NewGetLicenseHandler(graphService thegraph.TheGraphServiceMvp) func(c *gin.Context) {
	return func(c *gin.Context) {
		version := c.GetHeader(entity.HTTP_HEADER_VERSION)
		switch version {
		case entity.VERSION_V0_ALPHA:
			response, statusCode, err := v0alpha.HandleGetLicense(c)
			if err != nil {
				c.JSON(statusCode, ErrorMessage(err.Error()))
				return
			}
			c.JSON(statusCode, response)
		default:
			response, statusCode, err := v0mvp.HandleGetLicense(c, graphService)
			if err != nil {
				c.JSON(statusCode, ErrorMessage(err.Error()))
				return
			}
			c.JSON(statusCode, response)
		}

		/*
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
		*/
	}
}

// POST /license
func NewListLicensesHandler(graphService thegraph.TheGraphServiceMvp) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, v0alpha_entity.ListLicensesResponse{
			Licenses: []*v0alpha_entity.License{
				{
					ID:          "50",
					IPAssetId:   "5",
					IPOrgId:     "7",
					Owner:       "0xd84316a1b6f40902c17b8177854cdaeb3c957daf",
					MetadataUri: "https://arweave.net/R7-xPDAMqOhUSw3CM_UwXI7zdpQkzCCCUq3smzxyAaU",
					CreatedAt:   "2023-11-14T00:29:13Z",
					TxHash:      "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
				},
				{
					ID:          "49",
					IPAssetId:   "5",
					IPOrgId:     "7",
					Owner:       "0xd84316a1b6f40902c17b8177854cdaeb3c957daf",
					MetadataUri: "https://arweave.net/R7-xPDAMqOhUSw3CM_UwXI7zdpQkzCCCUq3smzxyAaU",
					CreatedAt:   "2023-11-14T00:29:13Z",
					TxHash:      "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
				},
			},
		})
	}
}

// GET /collection
func NewGetCollectionsHandler(graphService thegraph.TheGraphServiceMvp) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId := c.DefaultQuery("franchiseId", "")
		_, err := strconv.ParseInt(franchiseId, 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", franchiseId)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		collections, err := graphService.GetCollections(franchiseId)
		if err != nil {
			logger.Errorf("Failed to get collections: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, entity.GetCollectionsResponseMVP{
			Data: collections,
		})
		/*
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
		*/
	}
}

// GET /transaction
func NewGetTransactionsHandler(graphService thegraph.TheGraphServiceMvp) func(c *gin.Context) {
	return func(c *gin.Context) {
		transactions, err := graphService.GetTransactions()
		if err != nil {
			logger.Errorf("Failed to get transactions: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, entity.GetTransactionsResponseMVP{
			Data: transactions,
		})
		/*
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
		*/
	}
}

// GET /transaction/:transactionId
func NewGetTransactionHandler(graphService thegraph.TheGraphServiceMvp) func(c *gin.Context) {
	return func(c *gin.Context) {
		version := c.GetHeader(entity.HTTP_HEADER_VERSION)
		switch version {
		case entity.VERSION_V0_ALPHA:
			response, statusCode, err := v0alpha.HandleGetTransaction(c)
			if err != nil {
				c.JSON(statusCode, ErrorMessage(err.Error()))
				return
			}
			c.JSON(statusCode, response)
		default:
			response, statusCode, err := v0mvp.HandleGetTransaction(c, graphService)
			if err != nil {
				c.JSON(statusCode, ErrorMessage(err.Error()))
				return
			}
			c.JSON(statusCode, response)
		}

		/*
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
		*/
	}
}

// POST /transaction
func NewListTransactionsHandler(graphService thegraph.TheGraphServiceMvp) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, &v0alpha_entity.ListTransactionsResponse{
			Transactions: []*v0alpha_entity.Transaction{
				{
					ID:           "1",
					TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
					IPOrgId:      "7",
					ResourceId:   "1",
					ResourceType: v0alpha_entity.ResourceTypes.IPAsset,
					ActionType:   v0alpha_entity.ActionTypes.Create,
					Creator:      "0x4f9693ac46f2c7e2f48dd14d8fe1ab44192cd57d",
					CreatedAt:    "0001-01-01T00:00:00Z",
				},
				{
					ID:           "2",
					TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
					IPOrgId:      "7",
					ResourceId:   "50",
					ResourceType: v0alpha_entity.ResourceTypes.License,
					ActionType:   v0alpha_entity.ActionTypes.Create,
					Creator:      "0xd84316a1b6f40902c17b8177854cdaeb3c957daf",
					CreatedAt:    "0001-01-01T00:00:00Z",
				},
			},
		})
	}
}

// GET /relationship/:relationshipId
func NewGetRelationshipHandler(graphService thegraph.TheGraphServiceMvp) func(c *gin.Context) {
	return func(c *gin.Context) {
		ttl := int64(1000)
		c.JSON(http.StatusOK, &v0alpha_entity.GetRelationshipResponse{
			Relationship: &v0alpha_entity.Relationship{
				ID:           "1",
				Type:         "APPEAR_IN",
				TypeId:       "2",
				TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
				SrcContract:  "0x123",
				SrcTokenId:   "4",
				SrcName:      "Darth Vader",
				DstContract:  "0x456",
				DstTokenId:   "7",
				DstName:      "Star War",
				RegisteredAt: "0001-01-01T00:00:00Z",
				TTL:          &ttl,
			},
		})
	}
}

// POST /relationship
func NewListRelationshipsHandler(graphService thegraph.TheGraphServiceMvp) func(c *gin.Context) {
	return func(c *gin.Context) {
		ttl := int64(1000)
		c.JSON(http.StatusOK, &v0alpha_entity.ListRelationshipsResponse{
			Relationships: []*v0alpha_entity.Relationship{
				{
					ID:           "1",
					Type:         "APPEAR_IN",
					TypeId:       "2",
					TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
					SrcContract:  "0x123",
					SrcTokenId:   "4",
					SrcName:      "Darth Vader",
					DstContract:  "0x456",
					DstTokenId:   "7",
					DstName:      "Star War",
					RegisteredAt: "0001-01-01T00:00:00Z",
					TTL:          &ttl,
				},
				{
					ID:           "2",
					Type:         "APPEAR_IN",
					TypeId:       "2",
					TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
					SrcContract:  "0x123",
					SrcTokenId:   "4",
					SrcName:      "Darth Vader",
					DstContract:  "0x456",
					DstTokenId:   "7",
					DstName:      "Star War",
					RegisteredAt: "0001-01-01T00:00:00Z",
				},
			},
		})
	}
}

// POST /module
func NewListModulesHandler(graphService thegraph.TheGraphServiceMvp) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, &v0alpha_entity.ListModulesResponse{
			Modules: []*v0alpha_entity.Module{
				{
					ID:        "1",
					IPOrgId:   "7",
					Interface: "(address,uint)",
					PreHooks: []*v0alpha_entity.Hook{
						{
							ID:           "1",
							TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
							ModuleId:     "0x1234514e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
							Interface:    "(uint,uint)",
							RegisteredAt: "0001-01-01T00:00:00Z",
						},
					},
					PostHooks: []*v0alpha_entity.Hook{
						{
							ID:           "1",
							TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
							ModuleId:     "0x1234514e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
							Interface:    "(uint,uint)",
							RegisteredAt: "0001-01-01T00:00:00Z",
						},
					},
				},
				{
					ID:        "2",
					IPOrgId:   "7",
					Interface: "(uint,uint)",
					PreHooks: []*v0alpha_entity.Hook{
						{
							ID:           "1",
							TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
							ModuleId:     "0x1234514e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
							Interface:    "(uint,uint)",
							RegisteredAt: "0001-01-01T00:00:00Z",
						},
					},
					PostHooks: []*v0alpha_entity.Hook{
						{
							ID:           "1",
							TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
							ModuleId:     "0x1234514e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
							Interface:    "(uint,uint)",
							RegisteredAt: "0001-01-01T00:00:00Z",
						},
					},
				},
			},
		})
	}
}

// GET /module/:moduleId
func NewGetModuleHandler(graphService thegraph.TheGraphServiceMvp) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, &v0alpha_entity.GetModuleResponse{
			Module: &v0alpha_entity.Module{
				ID:        "2",
				IPOrgId:   "7",
				Interface: "(uint,uint)",
				PreHooks: []*v0alpha_entity.Hook{
					{
						ID:           "1",
						TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
						ModuleId:     "0x1234514e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
						Interface:    "(uint,uint)",
						RegisteredAt: "0001-01-01T00:00:00Z",
					},
				},
				PostHooks: []*v0alpha_entity.Hook{
					{
						ID:           "1",
						TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
						ModuleId:     "0x1234514e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
						Interface:    "(uint,uint)",
						RegisteredAt: "0001-01-01T00:00:00Z",
					},
				},
			},
		})
	}
}

// POST /hook
func NewListHooksHandler(graphService thegraph.TheGraphServiceMvp) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, &v0alpha_entity.ListHooksResponse{
			Hooks: []*v0alpha_entity.Hook{
				{
					ID:           "1",
					TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
					ModuleId:     "0x1234514e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
					Interface:    "(uint,uint)",
					RegisteredAt: "0001-01-01T00:00:00Z",
				},
				{
					ID:           "2",
					TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
					ModuleId:     "0x1234514e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
					Interface:    "(uint,uint)",
					RegisteredAt: "0001-01-01T00:00:00Z",
				},
			},
		})
	}
}

// GET /hook/:hookId
func NewGetHookHandler(graphService thegraph.TheGraphServiceMvp) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, &v0alpha_entity.GetHookResponse{
			Hook: &v0alpha_entity.Hook{
				ID:           "2",
				TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
				ModuleId:     "0x1234514e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
				Interface:    "(uint,uint)",
				RegisteredAt: "0001-01-01T00:00:00Z",
			},
		})
	}
}
