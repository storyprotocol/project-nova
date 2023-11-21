package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/api/internal/service/thegraph"
	xhttp "github.com/project-nova/backend/pkg/http"
)

type ProtocolHandlerInterface interface {
	ListIpOrgsHandler(c *gin.Context)        // POST /iporg
	GetIpOrgHandler(c *gin.Context)          // GET /iporg/:ipOrgId
	GetIpAssetHandler(c *gin.Context)        // GET /ipasset/:ipAssetId
	ListIpAssetsHandler(c *gin.Context)      // POST /ipasset
	ListTransactionsHandler(c *gin.Context)  // POST /transaction
	GetRelationshipHandler(c *gin.Context)   // GET /relationship/:relationshipId
	ListRelationshipsHandler(c *gin.Context) // POST /relationship
	ListModulesHandler(c *gin.Context)       // POST /module
	GetModuleHandler(c *gin.Context)         // GET /module/:moduleId
	ListHooksHandler(c *gin.Context)         // POST /hook
	GetHookHandler(c *gin.Context)           // GET /hook/:hookId
	GetTransactionHandler(c *gin.Context)    // GET /transaction/:transactionId
	GetLicenseHandler(c *gin.Context)        // GET /license/:licenseId
	ListLicensesHandler(c *gin.Context)      // POST /license
}

type protocolHandler struct {
	alphaProtocolHandlers *AlphaProtocolHandler
}

func NewProtocolHandler(graphServiceAlpha thegraph.TheGraphServiceAlpha, httpClient xhttp.Client) ProtocolHandlerInterface {
	alphaProtocolHandlers := NewAlphaProtocolHandler(graphServiceAlpha)
	return &protocolHandler{
		alphaProtocolHandlers: alphaProtocolHandlers,
	}
}

func (ph *protocolHandler) ListIpOrgsHandler(c *gin.Context) {
	version := c.GetHeader(entity.HTTP_HEADER_VERSION)
	switch version {
	case entity.VERSION_V0_ALPHA:
		ph.alphaProtocolHandlers.ListIpOrgsHandler(c)
	default:
		ph.alphaProtocolHandlers.ListIpOrgsHandler(c)
	}
}

func (ph *protocolHandler) GetIpOrgHandler(c *gin.Context) {
	version := c.GetHeader(entity.HTTP_HEADER_VERSION)
	switch version {
	case entity.VERSION_V0_ALPHA:
		ph.alphaProtocolHandlers.GetIpOrgHandler(c)
	default:
		ph.alphaProtocolHandlers.GetIpOrgHandler(c)
	}
}

func (ph *protocolHandler) GetIpAssetHandler(c *gin.Context) {
	version := c.GetHeader(entity.HTTP_HEADER_VERSION)
	switch version {
	case entity.VERSION_V0_ALPHA:
		ph.alphaProtocolHandlers.GetIpAssetHandler(c)
	default:
		ph.alphaProtocolHandlers.GetIpAssetHandler(c)
	}
}

func (ph *protocolHandler) ListIpAssetsHandler(c *gin.Context) {
	version := c.GetHeader(entity.HTTP_HEADER_VERSION)
	switch version {
	case entity.VERSION_V0_ALPHA:
		ph.alphaProtocolHandlers.ListIpAssetsHandler(c)
	default:
		ph.alphaProtocolHandlers.ListIpAssetsHandler(c)
	}
}

func (ph *protocolHandler) ListTransactionsHandler(c *gin.Context) {
	version := c.GetHeader(entity.HTTP_HEADER_VERSION)
	switch version {
	case entity.VERSION_V0_ALPHA:
		ph.alphaProtocolHandlers.ListTransactionsHandler(c)
	default:
		ph.alphaProtocolHandlers.ListTransactionsHandler(c)
	}
}

func (ph *protocolHandler) GetRelationshipHandler(c *gin.Context) {
	version := c.GetHeader(entity.HTTP_HEADER_VERSION)
	switch version {
	case entity.VERSION_V0_ALPHA:
		ph.alphaProtocolHandlers.GetRelationshipHandler(c)
	default:
		ph.alphaProtocolHandlers.GetRelationshipHandler(c)
	}
}

func (ph *protocolHandler) ListRelationshipsHandler(c *gin.Context) {
	version := c.GetHeader(entity.HTTP_HEADER_VERSION)
	switch version {
	case entity.VERSION_V0_ALPHA:
		ph.alphaProtocolHandlers.ListRelationshipsHandler(c)
	default:
		ph.alphaProtocolHandlers.ListRelationshipsHandler(c)
	}
}

func (ph *protocolHandler) ListModulesHandler(c *gin.Context) {
	version := c.GetHeader(entity.HTTP_HEADER_VERSION)
	switch version {
	case entity.VERSION_V0_ALPHA:
		ph.alphaProtocolHandlers.ListModulesHandler(c)
	default:
		ph.alphaProtocolHandlers.ListModulesHandler(c)
	}
}

func (ph *protocolHandler) GetModuleHandler(c *gin.Context) {
	version := c.GetHeader(entity.HTTP_HEADER_VERSION)
	switch version {
	case entity.VERSION_V0_ALPHA:
		ph.alphaProtocolHandlers.GetModuleHandler(c)
	default:
		ph.alphaProtocolHandlers.GetModuleHandler(c)
	}
}

func (ph *protocolHandler) ListHooksHandler(c *gin.Context) {
	version := c.GetHeader(entity.HTTP_HEADER_VERSION)
	switch version {
	case entity.VERSION_V0_ALPHA:
		ph.alphaProtocolHandlers.ListHooksHandler(c)
	default:
		ph.alphaProtocolHandlers.ListHooksHandler(c)
	}
}

func (ph *protocolHandler) GetHookHandler(c *gin.Context) {
	version := c.GetHeader(entity.HTTP_HEADER_VERSION)
	switch version {
	case entity.VERSION_V0_ALPHA:
		ph.alphaProtocolHandlers.GetHookHandler(c)
	default:
		ph.alphaProtocolHandlers.GetHookHandler(c)
	}
}

func (ph *protocolHandler) GetTransactionHandler(c *gin.Context) {
	version := c.GetHeader(entity.HTTP_HEADER_VERSION)
	switch version {
	case entity.VERSION_V0_ALPHA:
		ph.alphaProtocolHandlers.GetTransactionHandler(c)
	default:
		ph.alphaProtocolHandlers.GetTransactionHandler(c)
	}
}

func (ph *protocolHandler) GetLicenseHandler(c *gin.Context) {
	version := c.GetHeader(entity.HTTP_HEADER_VERSION)
	switch version {
	case entity.VERSION_V0_ALPHA:
		ph.alphaProtocolHandlers.GetLicenseHandler(c)
	default:
		ph.alphaProtocolHandlers.GetLicenseHandler(c)
	}
}

func (ph *protocolHandler) ListLicensesHandler(c *gin.Context) {
	version := c.GetHeader(entity.HTTP_HEADER_VERSION)
	switch version {
	case entity.VERSION_V0_ALPHA:
		ph.alphaProtocolHandlers.ListLicensesHandler(c)
	default:
		ph.alphaProtocolHandlers.ListLicensesHandler(c)
	}
}
