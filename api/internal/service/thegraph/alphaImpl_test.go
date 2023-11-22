package thegraph

import (
	"testing"

	"github.com/machinebox/graphql"
	"github.com/project-nova/backend/pkg/logger"
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
	assert.True(t, iporg.CreatedAt == "2023-11-18T06:44:12Z")
}

func TestGetIPOrgs_NotFound_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	iporg, err := service.GetIPOrg("0xde493e03d2de0cd7820b4f580beced57296b0011")
	assert.Nil(t, err)
	assert.Nil(t, iporg)
}

func TestListRelationships_MatchSource_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	relationships, err := service.ListRelationships("0x177175a4b26f6ea050676f8c9a14d395f896492c", "4", nil)
	assert.Nil(t, err)
	assert.True(t, len(relationships) > 0)
	assert.True(t, relationships[0].RegisteredAt == "2023-11-22T03:37:48Z")
}

func TestListRelationships_MatchDestination_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	relationships, err := service.ListRelationships("0x177175a4b26f6ea050676f8c9a14d395f896492c", "5", nil)
	assert.Nil(t, err)
	assert.True(t, len(relationships) > 0)
}

func TestListRelationships_NotMatch_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	relationships, err := service.ListRelationships("0x177175a4b26f6ea050676f8c9a14d395f896492c", "51", nil)
	assert.Nil(t, err)
	assert.True(t, len(relationships) == 0)
}

func TestGetRelationshipHandler_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	relationship, err := service.GetRelationship("1")
	assert.Nil(t, err)
	assert.True(t, relationship.ID == "1")
	assert.True(t, relationship.RegisteredAt == "2023-11-19T22:10:48Z")
}

func TestGetRelationshipHandler_NotFound_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	relationship, err := service.GetRelationship("123123")
	assert.Nil(t, err)
	assert.Nil(t, relationship)
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

func TestGetModule_NotFound_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	module, err := service.GetModule("0xa906e2589a7f8385a376babbb70a39dad551604b")
	assert.Nil(t, err)
	assert.Nil(t, module)
}

func TestListIPAssets_WithIpOrgId_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	ipOrgId := IP_ORG_ID
	ipAssets, err := service.ListIPAssets(&ipOrgId, nil)
	assert.Nil(t, err)
	assert.True(t, len(ipAssets) > 0)
}

func TestListIPAssets_WithIpOrgId_WithLimit_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	ipOrgId := IP_ORG_ID
	ipAssets, err := service.ListIPAssets(&ipOrgId, &TheGraphQueryOptions{
		First: 2,
		Skip:  0,
	})
	assert.Nil(t, err)
	assert.True(t, len(ipAssets) == 2)
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
	logger.Infof("ipAsset: %+v", ipAsset)
	assert.Nil(t, err)
	assert.True(t, ipAsset.ID == "1")
	assert.True(t, ipAsset.CreatedAt == "2023-11-19T18:46:24Z")
}

func TestGetIPAsset_NotFound_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	ipAsset, err := service.GetIPAsset("12312")
	assert.Nil(t, err)
	assert.Nil(t, ipAsset)
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
	assert.True(t, transaction.CreatedAt == "2023-11-19T18:46:24Z")
}

func TestGetTransaction_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	transaction, err := service.GetTransaction("0x158f74772af1bf9e5d1eb9d6633bb6a602eea97bbbd552b16696d7d1d3fa007703000000")
	assert.Nil(t, err)
	assert.Nil(t, transaction)
}

func TestListHooks_WithModuleID_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	moduleID := "0x091e5f55135155bb8cb5868adb39e5c34eb32cfd"
	hooks, err := service.ListHooks(&moduleID, nil)
	assert.Nil(t, err)
	assert.True(t, len(hooks) > 0)
}

func TestListHooks_WithoutModuleID_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	hooks, err := service.ListHooks(nil, nil)
	assert.Nil(t, err)
	assert.True(t, len(hooks) > 0)
}

func TestGetHook_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	hook, err := service.GetHook("0xc0f6e387ac0b324ec18eacf22ee7271207dce3d5")
	assert.Nil(t, err)
	assert.True(t, hook.ID == "0xc0f6e387ac0b324ec18eacf22ee7271207dce3d5")
	assert.True(t, hook.RegisteredAt == "2023-11-18T06:44:24Z")

}

func TestGetHook_NotFound_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	hook, err := service.GetHook("0xc0f6e387ac0b324ec18eacf22ee7271207dce2d5")
	assert.Nil(t, err)
	assert.Nil(t, hook)
}

func TestListLicenses_WithIpOrgIdAndIpAssetId_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	ipOrgId := "0xb422e54932c1dae83e78267a4dd2805aa64a8061"
	ipAssetId := "0"
	licenses, err := service.ListLicenses(&ipOrgId, &ipAssetId, nil)
	assert.Nil(t, err)
	assert.True(t, len(licenses) > 0)
}

func TestListLicenses_WithIpOrgId_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	ipOrgId := "0xb422e54932c1dae83e78267a4dd2805aa64a8061"
	licenses, err := service.ListLicenses(&ipOrgId, nil, nil)
	assert.Nil(t, err)
	assert.True(t, len(licenses) > 0)
}

func TestListLicenses_WithIpAssetId_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	ipAssetId := "0"
	licenses, err := service.ListLicenses(nil, &ipAssetId, nil)
	assert.Nil(t, err)
	assert.True(t, len(licenses) > 0)
}

func TestListLicenses_WithoutIpOrgIdAndIpAssetId_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	licenses, err := service.ListLicenses(nil, nil, nil)
	assert.Nil(t, err)
	assert.True(t, len(licenses) > 0)
}

func TestGetLicense_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	license, err := service.GetLicense("1")
	assert.Nil(t, err)
	assert.True(t, license.ID == "1")
	assert.True(t, license.CreatedAt == "2023-11-21T02:20:24Z")
}

func TestGetLicense_NotFound_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	license, err := service.GetLicense("1123")
	assert.Nil(t, err)
	assert.Nil(t, license)
}
