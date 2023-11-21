package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/project-nova/backend/api/internal/service/thegraph"
)

func CreateTheGraphServiceAlpha() thegraph.TheGraphServiceAlpha {
	theGraphClientAlpha := graphql.NewClient("https://api.thegraph.com/subgraphs/name/edisonz0718/storyprotocol-v0-alpha")
	return thegraph.NewTheGraphServiceAlphaImpl(theGraphClientAlpha)
}

func MockGin(requestBody map[string]interface{}) (*gin.Context, *httptest.ResponseRecorder) {
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
