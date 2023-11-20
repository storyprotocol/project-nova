package thegraph

import (
	"testing"

	"github.com/machinebox/graphql"
	"github.com/stretchr/testify/assert"
)

const IP_ORG_ID string = "0xde493e03d2de0cd7820b4f580beced57296b0009"

func CreateTheGraphServiceAlpha() TheGraphServiceAlpha {
	theGraphClientAlpha := graphql.NewClient("https://api.thegraph.com/subgraphs/name/edisonz0718/storyprotocol-v0-alpha")
	return NewTheGraphServiceAlphaImpl(theGraphClientAlpha)
}

func TestListIPOrgs_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	iporgs, err := service.ListIPOrgs(&TheGraphQueryOptions{
		First: 10,
	})

	assert.Nil(t, err)
	assert.True(t, len(iporgs) > 0)
}

func TestGetIPOrgs_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	iporg, err := service.GetIPOrg("0xde493e03d2de0cd7820b4f580beced57296b0009")

	assert.Nil(t, err)
	assert.NotNil(t, iporg)
	assert.True(t, iporg.ID == "0xde493e03d2de0cd7820b4f580beced57296b0009")
}

func TestListRelationships_MatchSource_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	relationships, err := service.ListRelationships("0x177175a4b26f6ea050676f8c9a14d395f896492c", "4", nil)
	assert.Nil(t, err)
	assert.True(t, len(relationships) > 0)
}

func TestListRelationships_MatchDestination_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	relationships, err := service.ListRelationships("0x177175a4b26f6ea050676f8c9a14d395f896492c", "5", nil)
	assert.Nil(t, err)
	assert.True(t, len(relationships) > 0)
}

func TestListModules_WithIpOrgId_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	ipOrgId := "0x0000000000000000000000000000000000000000"
	modules, err := service.ListModules(&ipOrgId, nil)
	assert.Nil(t, err)
	assert.True(t, len(modules) > 0)
}

func TestListModules_WithoutIpOrgId_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	modules, err := service.ListModules(nil, nil)
	assert.Nil(t, err)
	assert.True(t, len(modules) > 0)
}

func TestGetModule_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	module, err := service.GetModule("0xa906e2589a7f8385a376babbb70a39dad551603b")
	assert.Nil(t, err)
	assert.True(t, module.ModuleKey == "LICENSING_MODULE")
}

func TestListIPAssets_WithIpOrgId_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	ipOrgId := IP_ORG_ID
	ipAssets, err := service.ListIPAssets(&ipOrgId, nil)
	assert.Nil(t, err)
	assert.True(t, len(ipAssets) > 0)
}

func TestListIPAssets_WithoutIpOrgId_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	ipAssets, err := service.ListIPAssets(nil, nil)
	assert.Nil(t, err)
	assert.True(t, len(ipAssets) > 0)
}

func TestGetIPAsset_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	ipAsset, err := service.GetIPAsset("1")
	assert.Nil(t, err)
	assert.True(t, ipAsset.ID == "0x158f74772af1bf9e5d1eb9d6633bb6a602eea97bbbd552b16696d7d2d3fa007703000000")
}

func TestListTransactions_WithIpOrgId_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	ipOrgId := IP_ORG_ID
	transactions, err := service.ListTransactions(&ipOrgId, nil)
	assert.Nil(t, err)
	assert.True(t, len(transactions) > 0)
}

func TestListTransactions_WithoutIpOrgId_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	transactions, err := service.ListTransactions(nil, nil)
	assert.Nil(t, err)
	assert.True(t, len(transactions) > 0)
}

func TestGetTransaction_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	transaction, err := service.GetTransaction("0x158f74772af1bf9e5d1eb9d6633bb6a602eea97bbbd552b16696d7d2d3fa007703000000")
	assert.Nil(t, err)
	assert.True(t, transaction.ID == "0x158f74772af1bf9e5d1eb9d6633bb6a602eea97bbbd552b16696d7d2d3fa007703000000")
}
