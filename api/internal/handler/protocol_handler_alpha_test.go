package handler

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	v0alpha_entity "github.com/project-nova/backend/api/internal/entity/v0-alpha"
	"github.com/project-nova/backend/api/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestListIpOrgsHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  10,
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListIpOrgsHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)

	var response v0alpha_entity.ListIpOrgsResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response)
	assert.True(t, len(response.IPOrgs) <= 10)
}

func TestListIpOrgsHandler_EmptyResponse_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 5000,
				"limit":  0,
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListIpOrgsHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)

	var response v0alpha_entity.ListIpOrgsResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Empty(t, response.IPOrgs)
}

func TestListIpOrgsHandler_InvalidPaginationParams_Failure(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": "invalid",
				"limit":  10,
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListIpOrgsHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetIpOrgHandler_Success(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "ipOrgId", Value: "0x0dad65978b6c637598674ea03b1c6f3333d00f5b"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetIpOrgHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)

	var response v0alpha_entity.GetIpOrgResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response)
	assert.NotNil(t, response.IPOrg)

}

func TestGetIpOrgHandler_InvalidId_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "ipOrgId", Value: "invalid_id"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetIpOrgHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetIpOrgHandler_EmptyId_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "ipOrgId", Value: ""},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetIpOrgHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetIpOrgHandler_WrongIpOrgId_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "ipOrgId", Value: "0x1234567890abcdef"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetIpOrgHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetIpOrgHandler_NotFound_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "ipOrgId", Value: "0x0dad65978b6c637598674ea03b1c6f4333d00f5b"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetIpOrgHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetIpAssetHandler_Success(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "ipAssetId", Value: "1"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetIpAssetHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetIpAssetHandler_EmptyId_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "ipAssetId", Value: ""},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetIpAssetHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetIpAssetHandler_InvalidId_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "ipAssetId", Value: "invalid_id"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetIpAssetHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetIpAssetHandler_NotFound_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "ipAssetId", Value: "10000000"},
	}

	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetIpAssetHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestListIpAssetsHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"ipOrgId": "0xde493e03d2de0cd7820b4f580beced57296b0009",
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  1,
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListIpAssetsHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListIpAssetsHandler_EmptyIpOrgId_Failure(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"ipOrgId": "",
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  1,
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListIpAssetsHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListIpAssetsHandler_InvalidPaginationParams_Failure(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"ipOrgId": "0xde493e03d2de0cd7820b4f580beced57296b0009",
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": "invalid_offset",
				"limit":  1,
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListIpAssetsHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListIpAssetsHandler_NoResults_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"ipOrgId": "0xde493e03d2de0cd7820b4f580beced57296b0009",
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  1,
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListIpAssetsHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
	var response v0alpha_entity.ListIpAssetsResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	// Assert that the response body is empty
	assert.True(t, len(response.IPAssets) == 0)
}

func TestListIpAssetsHandler_InvalidIpOrgId_Failure(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"ipOrgId": "invalid_id",
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  1,
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListIpAssetsHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetRelationshipHandler_Success(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "relationshipId", Value: "1"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetRelationshipHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetRelationshipHandler_EmptyId_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "relationshipId", Value: ""},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetRelationshipHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetRelationshipHandler_InvalidId_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "relationshipId", Value: "invalid_id"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetRelationshipHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetRelationshipHandler_NotFound_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "relationshipId", Value: "10000000"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetRelationshipHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetRelationshipTypeHandler_FindByRelType_AndIpOrgId_Success(t *testing.T) {
	c, w := test.MockGin(nil, map[string]interface{}{
		"ipOrgId": "0xb422e54932c1dae83e78267a4dd2805aa64a8061",
		"relType": "0xc12a5f0d1e5a95f4fc32ff629c53defa11273a372e29ae51ab24323e4af84fc3",
	})
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetRelationshipTypeHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetRelationshipTypeHandler_MissingParams_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetRelationshipTypeHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetRelationshipTypeHandler_NotFound_Failure(t *testing.T) {
	c, w := test.MockGin(nil, map[string]interface{}{
		"ipOrgId": "0xb422e54932c1dae83e78267a4dd2805aa64a8062",
		"relType": "0xc12aaf0d1e5a95f4fc32ff629c53dafa11273a372e29ae51ab24323e4af84fc3",
	})
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetRelationshipTypeHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestListRelationshipTypesHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"ipOrgId": "0xb422e54932c1dae83e78267a4dd2805aa64a8061",
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  1,
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListRelationshipTypesHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListRelationshipTypesHandler_WithoutParams_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  1,
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListRelationshipTypesHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListRelationshipsHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"contract": "0x177175a4b26f6ea050676f8c9a14d395f896492c",
		"tokenId":  "5",
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListRelationshipsHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListModulesHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"ipOrgId": "0x0000000000000000000000000000000000000000",
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  10,
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListModulesHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetModuleHandler_Success(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "moduleId", Value: "0xd692de739fe1c1aaa31c3d0847dc17976afc05ff"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetModuleHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetModuleHandler_EmptyId_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "moduleId", Value: ""},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetModuleHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetModuleHandler_InvalidId_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "moduleId", Value: "invalid_module_id"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetModuleHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetModuleHandler_NotFound_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "moduleId", Value: "0x091e5f55135155bb8cb5868adb39e5c34eb32cfe"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetModuleHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestListHooksHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"moduleId": "0x948f67e1c4f75bc89c5fb42147d96356fb4b359f",
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  10,
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListHooksHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListHooksHandler_EmptyModuleId_Failure(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"moduleId": "",
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  10,
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListHooksHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListHooksHandler_InvalidModuleId_Failure(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"moduleId": "invalid_module_id",
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  10,
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListHooksHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListHooksHandler_InvalidPaginationParams_Failure(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"moduleId": "0x948f67e1c4f75bc89c5fb42147d96356fb4b359f",
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": "invalid_offset",
				"limit":  "invalid_limit",
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListHooksHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListHooksHandler_WithourParams_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  10,
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListHooksHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetHookHandler_Success(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "hookId", Value: "0xa26ba8224fb6173063f63388685f80708a6f4d96"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetHookHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetHookHandler_EmptyId_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "hookId", Value: ""},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetHookHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetHookHandler_InvalidId_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "hookId", Value: "invalid_hook_id"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetHookHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetHookHandler_InvalidHookId_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "hookId", Value: "0x1234567890abcdef"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetHookHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetHookHandler_NotFound_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "hookId", Value: "0xc0f6e387ac0b324ec18eacf22ee7271207dce2d5"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetHookHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetTransactionHandler_Success(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "transactionId", Value: "0x07da84387bbd29bf5476b0684677628f95d6b551fdb145c4fccb27b6342cdfd12e000000"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetTransactionHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTransactionHandler_EmptyId_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "transactionId", Value: ""},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetTransactionHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetTransactionHandler_InvalidId_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "transactionId", Value: "invalid_id"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetTransactionHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetTransactionHandler_LongId_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "transactionId", Value: "0x07da84387bbd29bf5476b0684677628f95d6b551fdb145c4fccb27b6342cdfd12e00000000000000000000000000000000000000000000000000000000000000000"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetTransactionHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetTransactionHandler_NotFound_Failure(t *testing.T) {
	c, w := test.MockGin(nil, nil)
	c.Params = gin.Params{
		{Key: "transactionId", Value: "0x07da84387bbd29bf5476b0684677628f95d6b551fdb145c4accb27b6342cdfd12e000000"},
	}
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.GetTransactionHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

// func TestGetLicenseHandler_Success(t *testing.T) {
// 	c, w := test.MockGin(nil, nil)
// 	c.Params = gin.Params{
// 		{Key: "licenseId", Value: "1"},
// 	}
// 	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
// 	ph.GetLicenseHandler(c)
// 	assert.Equal(t, http.StatusOK, w.Code)
// }

// func TestGetLicenseHandler_NotFound_Failure(t *testing.T) {
// 	c, w := test.MockGin(nil, nil)
// 	c.Params = gin.Params{
// 		{Key: "licenseId", Value: "10000000"},
// 	}
// 	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
// 	ph.GetLicenseHandler(c)
// 	assert.Equal(t, http.StatusNotFound, w.Code)
// }

// func TestListLicensesHandler_Success(t *testing.T) {
// 	c, w := test.MockGin(map[string]interface{}{
// 		"ipOrgId":   "0xb422e54932c1dae83e78267a4dd2805aa64a8061",
// 		"ipAssetId": "0",
// 		"options": map[string]interface{}{
// 			"pagination": map[string]interface{}{
// 				"offset": 0,
// 				"limit":  10,
// 			},
// 		},
// 	})
// 	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
// 	ph.ListLicensesHandler(c)
// 	assert.Equal(t, http.StatusOK, w.Code)
// }

func TestListTransactionsHandler_Success(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"ipOrgId": "0xb422e54932c1dae83e78267a4dd2805aa64a8062",
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  10,
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListTransactionsHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListTransactionsHandler_EmptyIpOrgId_Failure(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"ipOrgId": "",
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": 0,
				"limit":  10,
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListTransactionsHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListTransactionsHandler_InvalidPaginationParams_Failure(t *testing.T) {
	c, w := test.MockGin(map[string]interface{}{
		"ipOrgId": "0xb422e54932c1dae83e78267a4dd2805aa64a8062",
		"options": map[string]interface{}{
			"pagination": map[string]interface{}{
				"offset": "invalid",
				"limit":  10,
			},
		},
	}, nil)
	ph := NewAlphaProtocolHandler(test.CreateTheGraphServiceAlpha())
	ph.ListTransactionsHandler(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
