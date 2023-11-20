package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	v0alpha_entity "github.com/project-nova/backend/api/internal/entity/v0-alpha"
	v0alpha "github.com/project-nova/backend/api/internal/handler/v0-alpha"
	"github.com/project-nova/backend/api/internal/service/thegraph"
	xhttp "github.com/project-nova/backend/pkg/http"
	"github.com/project-nova/backend/pkg/logger"
)

// POST /iporg
func NewListIpOrgsHandlerAlpha(graphService thegraph.TheGraphServiceAlpha, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody v0alpha_entity.ListIpOrgsRequest
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
			return
		}

		iporgs, err := graphService.ListIPOrgs(thegraph.FromRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to get iporgs: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, v0alpha_entity.ListIpOrgsResponse{
			IPOrgs: iporgs,
		})
	}
}

// GET /iporg/:ipOrgId
func NewGetIpOrgHandlerAlpha(graphService thegraph.TheGraphServiceAlpha, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		ipOrdId := c.Param("ipOrgId")
		ipOrg, err := graphService.GetIPOrg(ipOrdId)
		if err != nil {
			logger.Errorf("Failed to get iporg: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, v0alpha_entity.GetIpOrgResponse{
			IPOrg: ipOrg,
		})
	}
}

// GET /ipasset/:ipAssetId
func NewGetIpAssetHandlerAlpha(graphService thegraph.TheGraphServiceAlpha, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		ipAssetId := c.Param("ipAssetId")
		ipAsset, err := graphService.GetIPAsset(ipAssetId)
		if err != nil {
			logger.Errorf("Failed to get ipasset: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, v0alpha_entity.GetIpAssetResponse{
			IPAsset: ipAsset,
		})
	}
}

// POST /ipasset
func NewListIpAssetsHandlerAlpha(graphService thegraph.TheGraphServiceAlpha, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId := c.DefaultQuery("franchiseId", "")
		_, err := strconv.ParseInt(franchiseId, 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", franchiseId)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		var requestBody v0alpha_entity.ListIpAssetsRequest
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
			return
		}

		ipAssets, err := graphService.ListIPAssets(&franchiseId, thegraph.FromRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to get ip assets: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, v0alpha_entity.ListIpAssetsResponse{
			IPAssets: ipAssets,
		})
	}
}

// POST /transaction
func NewListTransactionsHandlerAlpha(graphService thegraph.TheGraphServiceAlpha) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody v0alpha_entity.ListTransactionsRequest
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
			return
		}
		franchiseId := c.DefaultQuery("franchiseId", "")
		_, err := strconv.ParseInt(franchiseId, 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", franchiseId)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		transactions, err := graphService.ListTransactions(&franchiseId, thegraph.FromRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to get transactions: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, v0alpha_entity.ListTransactionsResponse{
			Transactions: transactions,
		})
	}
}

// GET /relationship/:relationshipId
func NewGetRelationshipHandlerAlpha(graphService thegraph.TheGraphServiceAlpha) func(c *gin.Context) {
	return func(c *gin.Context) {
		relationshipId := c.Param("relationshipId")
		relationship, err := graphService.GetRelationship(relationshipId)
		if err != nil {
			logger.Errorf("Failed to get relationship: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, v0alpha_entity.GetRelationshipResponse{
			Relationship: relationship,
		})
	}
}

// POST /relationship
func NewListRelationshipsHandlerAlpha(graphService thegraph.TheGraphServiceAlpha) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody v0alpha_entity.ListRelationshipRequest
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
			return
		}
		relationships, err := graphService.ListRelationships(requestBody.Contract, requestBody.TokenId, thegraph.FromRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to get relationships: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, v0alpha_entity.ListRelationshipsResponse{
			Relationships: relationships,
		})
	}
}

// POST /module
func NewListModulesHandlerAlpha(graphService thegraph.TheGraphServiceAlpha) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody v0alpha_entity.ListTransactionsRequest
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
			return
		}
		franchiseId := c.DefaultQuery("franchiseId", "")
		_, err := strconv.ParseInt(franchiseId, 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", franchiseId)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		modules, err := graphService.ListModules(&franchiseId, thegraph.FromRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to get modules: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, v0alpha_entity.ListModulesResponse{
			Modules: modules,
		})
	}
}

// GET /module/:moduleId
func NewGetModuleHandlerAlpha(graphService thegraph.TheGraphServiceAlpha) func(c *gin.Context) {
	return func(c *gin.Context) {
		moduleId := c.Param("moduleId")
		module, err := graphService.GetModule(moduleId)
		if err != nil {
			logger.Errorf("Failed to get module: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, v0alpha_entity.GetModuleResponse{
			Module: module,
		})
	}
}

// POST /hook
func NewListHooksHandlerAlpha(graphService thegraph.TheGraphServiceAlpha) func(c *gin.Context) {
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
func NewGetHookHandlerAlpha(graphService thegraph.TheGraphServiceAlpha) func(c *gin.Context) {
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

// GET /transaction/:transactionId
func NewGetTransactionHandlerAlpha(graphService thegraph.TheGraphServiceAlpha) func(c *gin.Context) {
	return func(c *gin.Context) {
		transactionId := c.Param("transactionId")
		transaction, err := graphService.GetTransaction(transactionId)
		if err != nil {
			logger.Errorf("Failed to get transaction: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, v0alpha_entity.GetTransactionResponse{
			Transaction: transaction,
		})
	}
}

// GET /license/:licenseId
func NewGetLicenseHandlerAlpha(graphService thegraph.TheGraphServiceAlpha) func(c *gin.Context) {
	return func(c *gin.Context) {
		response, statusCode, err := v0alpha.HandleGetLicense(c)
		if err != nil {
			c.JSON(statusCode, ErrorMessage(err.Error()))
			return
		}
		c.JSON(statusCode, response)

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
func NewListLicensesHandlerAlpha(graphService thegraph.TheGraphServiceAlpha) func(c *gin.Context) {
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
