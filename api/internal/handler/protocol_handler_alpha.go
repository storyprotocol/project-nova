package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v0alpha_entity "github.com/project-nova/backend/api/internal/entity/v0-alpha"
	"github.com/project-nova/backend/api/internal/service/thegraph"
	"github.com/project-nova/backend/pkg/logger"
)

type AlphaProtocolHandler struct {
	graphServiceAlpha thegraph.TheGraphServiceAlpha
}

func NewAlphaProtocolHandler(graphServiceAlpha thegraph.TheGraphServiceAlpha) *AlphaProtocolHandler {
	return &AlphaProtocolHandler{
		graphServiceAlpha: graphServiceAlpha,
	}
}

// POST /ipOrg
func (p *AlphaProtocolHandler) ListIpOrgsHandler(c *gin.Context) {
	var requestBody v0alpha_entity.ListIpOrgsRequest
	if err := c.BindJSON(&requestBody); err != nil {
		logger.Errorf("Failed to read request body: %v", err)
		requestBody = v0alpha_entity.ListIpOrgsRequest{}
	}

	iporgs, err := p.graphServiceAlpha.ListIPOrgs(thegraph.FromRequestQueryOptions(requestBody.Options))
	if err != nil {
		logger.Errorf("Failed to get iporgs: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
		return
	}

	c.JSON(http.StatusOK, v0alpha_entity.ListIpOrgsResponse{
		IPOrgs: iporgs,
	})
}

// GET /ipOrg/:ipOrgId
func (p *AlphaProtocolHandler) GetIpOrgHandler(c *gin.Context) {
	ipOrdId := c.Param("ipOrgId")
	ipOrg, err := p.graphServiceAlpha.GetIPOrg(ipOrdId)
	if err != nil {
		logger.Errorf("Failed to get iporg: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
		return
	}

	if ipOrg == nil {
		c.JSON(http.StatusNotFound, ErrorMessage("Not found"))
		return
	}

	c.JSON(http.StatusOK, v0alpha_entity.GetIpOrgResponse{
		IPOrg: ipOrg,
	})
}

// GET /ipasset/:ipAssetId
func (p *AlphaProtocolHandler) GetIpAssetHandler(c *gin.Context) {
	ipAssetId := c.Param("ipAssetId")
	ipAsset, err := p.graphServiceAlpha.GetIPAsset(ipAssetId)
	if err != nil {
		logger.Errorf("Failed to get ipasset: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
		return
	}

	if ipAsset == nil {
		c.JSON(http.StatusNotFound, ErrorMessage("Not found"))
		return
	}

	c.JSON(http.StatusOK, v0alpha_entity.GetIpAssetResponse{
		IPAsset: ipAsset,
	})
}

// POST /ipAsset
func (p *AlphaProtocolHandler) ListIpAssetsHandler(c *gin.Context) {
	var requestBody v0alpha_entity.ListIpAssetsRequest
	if err := c.BindJSON(&requestBody); err != nil {
		logger.Errorf("Failed to read request body: %v", err)
		requestBody = v0alpha_entity.ListIpAssetsRequest{}
	}

	logger.Infof("requestBody: %+v", requestBody)
	ipAssets, err := p.graphServiceAlpha.ListIPAssets(&requestBody.IpOrgId, thegraph.FromRequestQueryOptions(requestBody.Options))
	if err != nil {
		logger.Errorf("Failed to get ip assets: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
		return
	}

	c.JSON(http.StatusOK, v0alpha_entity.ListIpAssetsResponse{
		IPAssets: ipAssets,
	})
}

// GET /relationship/:relationshipId
func (p *AlphaProtocolHandler) GetRelationshipHandler(c *gin.Context) {
	relationshipId := c.Param("relationshipId")
	relationship, err := p.graphServiceAlpha.GetRelationship(relationshipId)
	if err != nil {
		logger.Errorf("Failed to get relationship: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
		return
	}

	if relationship == nil {
		c.JSON(http.StatusNotFound, ErrorMessage("Not found"))
		return
	}

	c.JSON(http.StatusOK, v0alpha_entity.GetRelationshipResponse{
		Relationship: relationship,
	})
}

// POST /relationship
func (p *AlphaProtocolHandler) ListRelationshipsHandler(c *gin.Context) {
	var requestBody v0alpha_entity.ListRelationshipRequest
	if err := c.BindJSON(&requestBody); err != nil {
		logger.Errorf("Failed to read request body: %v", err)
		c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
		return
	}
	relationships, err := p.graphServiceAlpha.ListRelationships(requestBody.Contract, requestBody.TokenId, thegraph.FromRequestQueryOptions(requestBody.Options))
	if err != nil {
		logger.Errorf("Failed to get relationships: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
		return
	}

	c.JSON(http.StatusOK, v0alpha_entity.ListRelationshipsResponse{
		Relationships: relationships,
	})
}

// GET /relationship-type
func (p *AlphaProtocolHandler) GetRelationshipTypeHandler(c *gin.Context) {
	var requestBody v0alpha_entity.GetRelationshipTypeRequest
	if err := c.BindQuery(&requestBody); err != nil {
		logger.Errorf("Failed to read request body: %v", err)
		c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
		return
	}

	relationshipType, err := p.graphServiceAlpha.GetRelationshipType(&requestBody.RelType, &requestBody.IpOrgId)
	if err != nil {
		logger.Errorf("Failed to get relationship types: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
		return
	}

	if relationshipType == nil {
		c.JSON(http.StatusNotFound, ErrorMessage("Not found"))
		return
	}

	c.JSON(http.StatusOK, v0alpha_entity.GetRelationshipTypeResponse{
		RelationshipType: relationshipType,
	})
}

// POST /relationship-type
func (p *AlphaProtocolHandler) ListRelationshipTypesHandler(c *gin.Context) {
	var requestBody v0alpha_entity.ListRelationshipTypesRequest
	if err := c.BindJSON(&requestBody); err != nil {
		logger.Errorf("Failed to read request body: %v", err)
		c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
		return
	}
	relationshipTypes, err := p.graphServiceAlpha.ListRelationshipTypes(&requestBody.IpOrgId, thegraph.FromRequestQueryOptions(requestBody.Options))
	if err != nil {
		logger.Errorf("Failed to get relationship types: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
		return
	}

	c.JSON(http.StatusOK, v0alpha_entity.ListRelationshipTypesResponse{
		RelationshipTypes: relationshipTypes,
	})
}

// POST /module
func (p *AlphaProtocolHandler) ListModulesHandler(c *gin.Context) {
	var requestBody v0alpha_entity.ListModulesRequest
	if err := c.BindJSON(&requestBody); err != nil {
		logger.Errorf("Failed to read request body: %v", err)
		requestBody = v0alpha_entity.ListModulesRequest{}
	}

	modules, err := p.graphServiceAlpha.ListModules(requestBody.IpOrgId, thegraph.FromRequestQueryOptions(requestBody.Options))
	if err != nil {
		logger.Errorf("Failed to get modules: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
		return
	}

	c.JSON(http.StatusOK, v0alpha_entity.ListModulesResponse{
		Modules: modules,
	})
}

// GET /module/:moduleId
func (p *AlphaProtocolHandler) GetModuleHandler(c *gin.Context) {
	moduleId := c.Param("moduleId")
	module, err := p.graphServiceAlpha.GetModule(moduleId)
	if err != nil {
		logger.Errorf("Failed to get module: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
		return
	}

	if module == nil {
		c.JSON(http.StatusNotFound, ErrorMessage("Not found"))
		return
	}

	c.JSON(http.StatusOK, v0alpha_entity.GetModuleResponse{
		Module: module,
	})
}

// POST /hook
func (p *AlphaProtocolHandler) ListHooksHandler(c *gin.Context) {
	var requestBody v0alpha_entity.ListHooksRequest
	if err := c.BindJSON(&requestBody); err != nil {
		logger.Errorf("Failed to read request body: %v", err)
		requestBody = v0alpha_entity.ListHooksRequest{}
	}

	hooks, err := p.graphServiceAlpha.ListHooks(requestBody.ModuleId, thegraph.FromRequestQueryOptions(requestBody.Options))
	if err != nil {
		logger.Errorf("Failed to get hooks: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
		return
	}

	c.JSON(http.StatusOK, v0alpha_entity.ListHooksResponse{
		Hooks: hooks,
	})
}

// GET /hook/:hookId
func (p *AlphaProtocolHandler) GetHookHandler(c *gin.Context) {
	hookId := c.Param("hookId")
	hook, err := p.graphServiceAlpha.GetHook(hookId)
	if err != nil {
		logger.Errorf("Failed to get hook: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
		return
	}

	if hook == nil {
		c.JSON(http.StatusNotFound, ErrorMessage("Not found"))
		return
	}

	c.JSON(http.StatusOK, v0alpha_entity.GetHookResponse{
		Hook: hook,
	})
}

// GET /transaction/:transactionId
func (p *AlphaProtocolHandler) GetTransactionHandler(c *gin.Context) {
	transactionId := c.Param("transactionId")
	logger.Infof("transactionId: %s", transactionId)
	transaction, err := p.graphServiceAlpha.GetTransaction(transactionId)
	if err != nil {
		logger.Errorf("Failed to get transaction: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
		return
	}

	if transaction == nil {
		c.JSON(http.StatusNotFound, ErrorMessage("Not found"))
		return
	}

	c.JSON(http.StatusOK, v0alpha_entity.GetTransactionResponse{
		Transaction: transaction,
	})
}

// GET /license/:licenseId
func (p *AlphaProtocolHandler) GetLicenseHandler(c *gin.Context) {
	licenseId := c.Param("licenseId")
	license, err := p.graphServiceAlpha.GetLicense(licenseId)
	if err != nil {
		logger.Errorf("Failed to get license: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
		return
	}

	if license == nil {
		c.JSON(http.StatusNotFound, ErrorMessage("Not found"))
		return
	}

	c.JSON(http.StatusOK, v0alpha_entity.GetLicenseResponse{
		License: license,
	})
}

// POST /license
func (p *AlphaProtocolHandler) ListLicensesHandler(c *gin.Context) {
	var requestBody v0alpha_entity.ListLicensesRequest
	if err := c.BindJSON(&requestBody); err != nil {
		logger.Errorf("Failed to read request body: %v", err)
		requestBody = v0alpha_entity.ListLicensesRequest{}
	}

	licenses, err := p.graphServiceAlpha.ListLicenses(requestBody.IpOrgId, requestBody.IpAssetId, thegraph.FromRequestQueryOptions(requestBody.Options))
	if err != nil {
		logger.Errorf("Failed to get licenses: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
		return
	}

	c.JSON(http.StatusOK, v0alpha_entity.ListLicensesResponse{
		Licenses: licenses,
	})
}

// POST /transaction
func (p *AlphaProtocolHandler) ListTransactionsHandler(c *gin.Context) {
	var requestBody v0alpha_entity.ListTransactionsRequest
	if err := c.BindJSON(&requestBody); err != nil {
		logger.Errorf("Failed to read request body: %v", err)
		requestBody = v0alpha_entity.ListTransactionsRequest{}
	}

	transactions, err := p.graphServiceAlpha.ListTransactions(requestBody.IpOrgId, thegraph.FromRequestQueryOptions(requestBody.Options))
	if err != nil {
		logger.Errorf("Failed to get transactions: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
		return
	}

	c.JSON(http.StatusOK, v0alpha_entity.ListTransactionsResponse{
		Transactions: transactions,
	})
}
