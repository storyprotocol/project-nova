package handler

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestListIpOrgsHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"pagination": map[string]interface{}{
			"offset": 0,
			"limit":  10,
		},
	})
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListIpOrgsHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetIpOrgHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "ipOrgId", Value: "0xde493e03d2de0cd7820b4f580beced57296b0009"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetIpOrgHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetIpOrgHandler_NotFound_Failure(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "ipOrgId", Value: "0xde493e03d2de0cd7820b4f580beced57296b0011"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetIpOrgHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetIpAssetHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "ipAssetId", Value: "1"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetIpAssetHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetIpAssetHandler_NotFound_Failure(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "ipAssetId", Value: "10000000"},
	}

	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetIpAssetHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestListIpAssethandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"ipOrgId": "0xde493e03d2de0cd7820b4f580beced57296b0009",
		"queryOptions": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  1,
			},
		},
	})
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListIpAssetsHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetRelationshipHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "relationshipId", Value: "1"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetRelationshipHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetRelationshipHandler_NotFound_Failure(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "relationshipId", Value: "10000000"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetRelationshipHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestListRelationshipsHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"contract": "0x177175a4b26f6ea050676f8c9a14d395f896492c",
		"tokenId":  "5",
	})
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListRelationshipsHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListModulesHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"ipOrgId": "0x0000000000000000000000000000000000000000",
		"queryOptions": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  10,
			},
		},
	})
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListModulesHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetModuleHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "moduleId", Value: "0x091e5f55135155bb8cb5868adb39e5c34eb32cfd"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetModuleHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetModuleHandler_NotFound_Failure(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "moduleId", Value: "0x091e5f55135155bb8cb5868adb39e5c34eb32cfe"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetModuleHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestListHooksHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"moduleId": "0x091e5f55135155bb8cb5868adb39e5c34eb32cfd",
		"queryOptions": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  10,
			},
		},
	})
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListHooksHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetHookHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "hookId", Value: "0xc0f6e387ac0b324ec18eacf22ee7271207dce3d5"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetHookHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetHookHandler_NotFound_Failure(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "hookId", Value: "0xc0f6e387ac0b324ec18eacf22ee7271207dce2d5"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetHookHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetTransactionHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "transactionId", Value: "0x158f74772af1bf9e5d1eb9d6633bb6a602eea97bbbd552b16696d7d2d3fa007703000000"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetTransactionHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTransactionHandler_NotFound_Failure(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "transactionId", Value: "0x158f74772af1bf9e5d1eb9d6633bb6a602eea97bbbd552b16696d7d2d3fa007703000001"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetTransactionHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetLicenseHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "licenseId", Value: "1"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetLicenseHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetLicenseHandler_NotFound_Failure(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "licenseId", Value: "10000000"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetLicenseHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestListLicensesHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"ipOrgId":   "0xb422e54932c1dae83e78267a4dd2805aa64a8061",
		"ipAssetId": "0",
		"queryOptions": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  10,
			},
		},
	})
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListLicensesHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListTransactionsHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"ipOrgId": "0xde493e03d2de0cd7820b4f580beced57296b0009",
		"queryOptions": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  10,
			},
		},
	})

	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListTransactionsHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}
