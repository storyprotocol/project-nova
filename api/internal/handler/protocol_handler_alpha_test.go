package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/project-nova/backend/api/internal/service/thegraph"
	"github.com/stretchr/testify/assert"
)

func mockGin(requestBody map[string]interface{}) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// test request, must instantiate a request first
	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header), // if you need to test headers
	}

	// convert requestBody to json
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBody))

	// example: req.Header.Add("Accept", "application/json")

	// // request query
	// testQuery := weldprogs.QueryParam{ /* init fields */ }

	// q := req.URL.Query()
	// for _, s := range testQuery.Basematgroup_id {
	// 	q.Add("basematgroup_id", s)
	// }
	// ... repeat for other fields as needed

	// must set this, since under the hood c.BindQuery calls
	// `req.URL.Query()`, which calls `ParseQuery(u.RawQuery)`
	// req.URL.RawQuery = q.Encode()

	// finally set the request to the gin context
	c.Request = req

	return c, w
}

func createTheGraphServiceAlpha() thegraph.TheGraphServiceAlpha {
	theGraphClientAlpha := graphql.NewClient("https://api.thegraph.com/subgraphs/name/edisonz0718/storyprotocol-v0-alpha")
	return thegraph.NewTheGraphServiceAlphaImpl(theGraphClientAlpha)
}

func TestListIpOrgsHandler_Success(t *testing.T) {
	c, w := mockGin(map[string]interface{}{
		"offset": 0,
		"limit":  10,
	})
	ph := NewAlphaProtocolHandler(createTheGraphServiceAlpha())
	ph.ListIpOrgsHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetIpOrgHandler_Success(t *testing.T) {
	c, w := mockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "ipOrgId", Value: "0xde493e03d2de0cd7820b4f580beced57296b0009"},
	}
	ph := NewAlphaProtocolHandler(createTheGraphServiceAlpha())
	ph.GetIpOrgHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetIpAssetHandler_Success(t *testing.T) {
	c, w := mockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "ipAssetId", Value: "1"},
	}
	ph := NewAlphaProtocolHandler(createTheGraphServiceAlpha())
	ph.GetIpAssetHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListIpAssethandler_Success(t *testing.T) {
	c, w := mockGin(map[string]interface{}{
		"ipOrgId": "0xde493e03d2de0cd7820b4f580beced57296b0009",
		"queryOptions": map[string]interface{}{
			"offset": 0,
			"limit":  10,
		},
	})
	ph := NewAlphaProtocolHandler(createTheGraphServiceAlpha())
	ph.ListIpAssetsHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetRelationshipHandler_Success(t *testing.T) {
	c, w := mockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "relationshipId", Value: "1"},
	}
	ph := NewAlphaProtocolHandler(createTheGraphServiceAlpha())
	ph.GetRelationshipHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListRelationshipsHandler_Success(t *testing.T) {
	c, w := mockGin(map[string]interface{}{
		"contract": "0x177175a4b26f6ea050676f8c9a14d395f896492c",
		"tokenId":  "5",
	})
	ph := NewAlphaProtocolHandler(createTheGraphServiceAlpha())
	ph.ListRelationshipsHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListModulesHandler_Success(t *testing.T) {
	c, w := mockGin(map[string]interface{}{
		"ipOrgId": "0x0000000000000000000000000000000000000000",
		"queryOptions": map[string]interface{}{
			"offset": 0,
			"limit":  10,
		},
	})
	ph := NewAlphaProtocolHandler(createTheGraphServiceAlpha())
	ph.ListModulesHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetModuleHandler_Success(t *testing.T) {
	c, w := mockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "moduleId", Value: "0x091e5f55135155bb8cb5868adb39e5c34eb32cfd"},
	}
	ph := NewAlphaProtocolHandler(createTheGraphServiceAlpha())
	ph.GetModuleHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListHooksHandler_Success(t *testing.T) {
	c, w := mockGin(map[string]interface{}{
		"moduleId": "0x091e5f55135155bb8cb5868adb39e5c34eb32cfd",
		"queryOptions": map[string]interface{}{
			"offset": 0,
			"limit":  10,
		},
	})
	ph := NewAlphaProtocolHandler(createTheGraphServiceAlpha())
	ph.ListHooksHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetHookHandler_Success(t *testing.T) {
	c, w := mockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "hookId", Value: "0xc0f6e387ac0b324ec18eacf22ee7271207dce3d5"},
	}
	ph := NewAlphaProtocolHandler(createTheGraphServiceAlpha())
	ph.GetHookHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTransactionHandler_Success(t *testing.T) {
	c, w := mockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "transactionId", Value: "0x158f74772af1bf9e5d1eb9d6633bb6a602eea97bbbd552b16696d7d2d3fa007703000000"},
	}
	ph := NewAlphaProtocolHandler(createTheGraphServiceAlpha())
	ph.GetTransactionHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetLicenseHandler_Success(t *testing.T) {
	c, w := mockGin(map[string]interface{}{})
	c.Params = gin.Params{
		{Key: "licenseId", Value: "1"},
	}
	ph := NewAlphaProtocolHandler(createTheGraphServiceAlpha())
	ph.GetLicenseHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListLicensesHandler_Success(t *testing.T) {
	c, w := mockGin(map[string]interface{}{
		"ipOrgId":   "0xb422e54932c1dae83e78267a4dd2805aa64a8061",
		"ipAssetId": "0",
		"queryOptions": map[string]interface{}{
			"offset": 0,
			"limit":  10,
		},
	})
	ph := NewAlphaProtocolHandler(createTheGraphServiceAlpha())
	ph.ListLicensesHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListTransactionsHandler_Success(t *testing.T) {
	c, w := mockGin(map[string]interface{}{
		"ipOrgId": "0xde493e03d2de0cd7820b4f580beced57296b0009",
		"queryOptions": map[string]interface{}{
			"offset": 0,
			"limit":  10,
		},
	})

	ph := NewAlphaProtocolHandler(createTheGraphServiceAlpha())
	ph.ListTransactionsHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

// func TestListTransactionsHandler_WithIpAssetId_Success(t *testing.T) {
// 	c, w := mockGin(map[string]interface{}{
// 		"ipAssetId": "1",
// 		"queryOptions": map[string]interface{}{
// 			"offset": 0,
// 			"limit":  10,
// 		},
// 	})

// 	ph := NewAlphaProtocolHandler(createTheGraphServiceAlpha())
// 	ph.ListTransactionsHandler(c)
// 	assert.Equal(t, http.StatusOK, w.Code)
// }
