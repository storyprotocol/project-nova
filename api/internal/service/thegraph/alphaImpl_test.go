package thegraph

import (
	"testing"

	"github.com/machinebox/graphql"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/stretchr/testify/assert"
)

const (
	THE_GRAPH_URL_ALPHA           string = "https://api.thegraph.com/subgraphs/name/edisonz0718/storyprotocol-v0-alpha"
	IP_ORG_ID_CORRECT             string = "0x1ebb43775fcc45cf05eaa96182c8762220e17941"
	IP_ORG_ID_WRONG               string = "0x1ebb43775fcc45cf05eaa96182c8762220e17942"
	RELATIONSHIP_TYPE_CORRECT     string = "0xc12a5f0d1e5a95f4fc32ff629c53defa11273a372e29ae51ab24323e4af84fc3"
	SOURCE_CONTRACT_CORRECT       string = "0x309c205347e3826472643f9b7ebd8a50d64ccd9e"
	SOURCE_TOKEN_ID_CORRECT       string = "2"
	DST_CONTRACT_CORRECT          string = "0x309c205347e3826472643f9b7ebd8a50d64ccd9e"
	DST_TOKEN_ID_CORRECT          string = "16"
	MODULE_ID_CORRECT             string = "0xd692de739fe1c1aaa31c3d0847dc17976afc05ff"
	TRANSACTION_ID_CORRECT        string = "0x07da84387bbd29bf5476b0684677628f95d6b551fdb145c4fccb27b6342cdfd12e000000"
	HOOK_LOOKUP_MODULE_ID_CORRECT string = "0x948f67e1c4f75bc89c5fb42147d96356fb4b359f"
	HOOK_ID_CORRECT               string = "0xa26ba8224fb6173063f63388685f80708a6f4d96"
)

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
func TestListIPOrgs_SuccessWithQueryOptions(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	options := &TheGraphQueryOptions{
		First:          10,
		Skip:           0,
		OrderBy:        "name",
		OrderDirection: "asc",
	}
	iporgs, err := service.ListIPOrgs(options)
	assert.Nil(t, err)
	assert.True(t, len(iporgs) > 0)
}

func TestListIPOrgs_SuccessWithNilQueryOptions(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	iporgs, err := service.ListIPOrgs(nil)
	assert.Nil(t, err)
	assert.True(t, len(iporgs) > 0)
}

func TestListIPOrgs_SuccessWithEmptyQueryOptions(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	options := &TheGraphQueryOptions{}
	iporgs, err := service.ListIPOrgs(options)
	assert.Nil(t, err)
	assert.True(t, len(iporgs) > 0)
}

func TestListIPOrgs_SuccessWithSkip(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	options := &TheGraphQueryOptions{
		First: 10,
		Skip:  10,
	}
	iporgs, err := service.ListIPOrgs(options)
	assert.Nil(t, err)
	assert.True(t, len(iporgs) > 0)
}

func TestListIPOrgs_SuccessWithOrderByDesc(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	options := &TheGraphQueryOptions{
		First:          10,
		OrderBy:        "name",
		OrderDirection: "desc",
	}
	iporgs, err := service.ListIPOrgs(options)
	assert.Nil(t, err)
	assert.True(t, len(iporgs) > 0)
}

func TestListIPOrgs_SuccessWithCustomQueryOptions(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	options := &TheGraphQueryOptions{
		First:          5,
		Skip:           5,
		OrderBy:        "createdAt",
		OrderDirection: "asc",
	}
	iporgs, err := service.ListIPOrgs(options)
	assert.Nil(t, err)
	assert.True(t, len(iporgs) > 0)
}

func TestGetRelationship_InvalidRelationshipId_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	_, err := service.GetRelationship("0xde493e03d2de0cd7820b4f580beced572a6b0011")
	assert.NotNil(t, err)
}

func TestGetRelationship_EmptyID_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	relationship, err := service.GetRelationship("")
	assert.NotNil(t, err)
	assert.Nil(t, relationship)
}

func TestGetRelationship_InvalidID_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	relationship, err := service.GetRelationship("invalid_id")
	assert.NotNil(t, err)
	assert.Nil(t, relationship)
}

func TestGetRelationshipType_NotFound_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	relType, err := service.GetRelationshipType("0xde493e03d2de0cd7820b4f580beced572a6b0011", IP_ORG_ID_CORRECT)
	assert.Nil(t, err)
	assert.Nil(t, relType)
}

func TestGetRelationshipType_EmptyID_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	relType, err := service.GetRelationshipType("", IP_ORG_ID_CORRECT)
	assert.NotNil(t, err)
	assert.Nil(t, relType)
}

func TestGetRelationshipType_InvalidID_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	relType, err := service.GetRelationshipType("invalid_id", IP_ORG_ID_CORRECT)
	assert.NotNil(t, err)
	assert.Nil(t, relType)
}

func TestListRelationships_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	options := &TheGraphQueryOptions{
		First: 10,
	}
	relationships, err := service.ListRelationships(SOURCE_CONTRACT_CORRECT, SOURCE_TOKEN_ID_CORRECT, options)
	assert.Nil(t, err)
	assert.True(t, len(relationships) > 0)
}

func TestListHooks_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	options := &TheGraphQueryOptions{
		First: 10,
	}
	hooks, err := service.ListHooks(nil, options)
	assert.Nil(t, err)
	assert.True(t, len(hooks) > 0)
}

func TestGetHook_EmptyID_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	hook, err := service.GetHook("")
	assert.Nil(t, err)
	assert.Nil(t, hook)
}

func TestGetHook_InvalidID_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	hook, err := service.GetHook("invalid_id")
	assert.NotNil(t, err)
	assert.Nil(t, hook)
}

func TestListModules_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	options := &TheGraphQueryOptions{
		First: 10,
	}
	modules, err := service.ListModules(nil, options)
	assert.Nil(t, err)
	assert.True(t, len(modules) > 0)
}

func TestGetModule_EmptyID_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	module, err := service.GetModule("")
	assert.Nil(t, err)
	assert.Nil(t, module)
}

func TestGetModule_InvalidID_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	module, err := service.GetModule("invalid_id")
	assert.NotNil(t, err)
	assert.Nil(t, module)
}

func TestListIPAssets_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	options := &TheGraphQueryOptions{
		First: 10,
	}
	ipAssets, err := service.ListIPAssets(nil, options)
	assert.Nil(t, err)
	assert.True(t, len(ipAssets) > 0)
}

func TestGetIPAsset_EmptyID_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	ipAsset, err := service.GetIPAsset("")
	assert.NotNil(t, err)
	assert.Nil(t, ipAsset)
}

func TestGetIPAsset_InvalidID_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	ipAsset, err := service.GetIPAsset("invalid_id")
	assert.NotNil(t, err)
	assert.Nil(t, ipAsset)
}

func TestListTransactions_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	options := &TheGraphQueryOptions{
		First: 10,
	}
	transactions, err := service.ListTransactions(nil, options)
	assert.Nil(t, err)
	assert.True(t, len(transactions) > 0)
}

func TestGetTransaction_NotFound_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	transaction, err := service.GetTransaction("0xde493e03d2de0cd7820b4f580beced572a6b0011")
	assert.Nil(t, err)
	assert.Nil(t, transaction)
}

func TestGetTransaction_EmptyID_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	transaction, err := service.GetTransaction("")
	assert.Nil(t, err)
	assert.Nil(t, transaction)
}

func TestGetTransaction_InvalidID_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	transaction, err := service.GetTransaction("invalid_id")
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
}

// func TestListLicenses_Success(t *testing.T) {
// 	service := CreateTheGraphServiceAlpha()
// 	options := &TheGraphQueryOptions{
// 		First: 10,
// 	}
// 	licenses, err := service.ListLicenses(nil, nil, options)
// 	assert.Nil(t, err)
// 	assert.True(t, len(licenses) > 0)
// }

// func TestGetLicense_NotFound_Failure(t *testing.T) {
// 	service := CreateTheGraphServiceAlpha()
// 	license, err := service.GetLicense("0xde493e03d2de0cd7820b4f580beced572a6b0011")
// 	assert.Nil(t, err)
// 	assert.Nil(t, license)
// }

// func TestGetLicense_EmptyID_Failure(t *testing.T) {
// 	service := CreateTheGraphServiceAlpha()
// 	license, err := service.GetLicense("")
// 	assert.NotNil(t, err)
// 	assert.Nil(t, license)
// }

// func TestGetLicense_InvalidID_Failure(t *testing.T) {
// 	service := CreateTheGraphServiceAlpha()
// 	license, err := service.GetLicense("invalid_id")
// 	assert.NotNil(t, err)
// 	assert.Nil(t, license)
// }

func TestGetIPOrgs_NotFound_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	iporg, err := service.GetIPOrg("0xde493e03d2de0cd7820b4f580beced572a6b0011")
	assert.Nil(t, err)
	assert.Nil(t, iporg)
}

func TestGetIPOrgs_EmptyID_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	iporg, err := service.GetIPOrg("")
	assert.Nil(t, err)
	assert.Nil(t, iporg)
}

func TestGetIPOrgs_InvalidID_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	iporg, err := service.GetIPOrg("invalid_id")
	assert.NotNil(t, err)
	assert.Nil(t, iporg)
}

func TestListRelationships_MatchSource_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	relationships, err := service.ListRelationships(SOURCE_CONTRACT_CORRECT, SOURCE_TOKEN_ID_CORRECT, nil)
	assert.Nil(t, err)
	assert.True(t, len(relationships) > 0)
}

func TestGetRelationshipType_FindByRelType_AndIpOrgId_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	relType := RELATIONSHIP_TYPE_CORRECT
	ipOrgId := IP_ORG_ID_CORRECT
	relationshipType, err := service.GetRelationshipType(relType, ipOrgId)
	assert.Nil(t, err)
	assert.True(t, relationshipType.IpOrgId == ipOrgId)
}

func TestGetRelationshipType_FindByIpOrgId_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	ipOrgId := IP_ORG_ID_CORRECT
	_, err := service.GetRelationshipType("", ipOrgId)
	assert.NotNil(t, err)
}

func TestGetRelationshipType_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	relType := "0x177175a4b26f6ea050676f8c9a14d395f896492c"
	_, err := service.GetRelationshipType(relType, "")
	assert.NotNil(t, err)
}

func TestListRelationshipTypes_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	relationshipTypes, err := service.ListRelationshipTypes(nil, nil)
	assert.Nil(t, err)
	assert.True(t, len(relationshipTypes) > 0)
	for i := 0; i < len(relationshipTypes)-1; i++ {
		assert.True(t, relationshipTypes[i].RegisteredAt >= relationshipTypes[i+1].RegisteredAt)
	}
}

func TestListRelationshipTypes_WithIpOrgId_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	ipOrgId := IP_ORG_ID_CORRECT
	relationshipTypes, err := service.ListRelationshipTypes(&ipOrgId, nil)
	assert.Nil(t, err)
	assert.True(t, len(relationshipTypes) > 0)
}

func TestListRelationships_MatchDestination_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	relationships, err := service.ListRelationships(DST_CONTRACT_CORRECT, DST_TOKEN_ID_CORRECT, nil)
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
	relationship, err := service.GetRelationship("2")
	assert.Nil(t, err)
	assert.True(t, relationship.ID == "2")
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
	module, err := service.GetModule(MODULE_ID_CORRECT)
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
	ipOrgId := IP_ORG_ID_CORRECT
	ipAssets, err := service.ListIPAssets(&ipOrgId, nil)
	assert.Nil(t, err)
	assert.True(t, len(ipAssets) > 0)
	assert.True(t, ipAssets[0].CreatedAt >= ipAssets[1].CreatedAt)
	assert.NotNil(t, ipAssets[0].Type)
}

func TestListIPAssets_WithIpOrgId_WithLimit_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	ipOrgId := IP_ORG_ID_CORRECT
	ipAssets, err := service.ListIPAssets(&ipOrgId, &TheGraphQueryOptions{
		First: 2,
		Skip:  0,
	})
	assert.Nil(t, err)
	assert.True(t, len(ipAssets) == 2)
	assert.True(t, ipAssets[0].CreatedAt >= ipAssets[1].CreatedAt)
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
}

func TestGetIPAsset_NotFound_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	ipAsset, err := service.GetIPAsset("12312")
	assert.Nil(t, err)
	assert.Nil(t, ipAsset)
}

func TestListTransactions_WithIpOrgId_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	ipOrgId := IP_ORG_ID_CORRECT
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
	transaction, err := service.GetTransaction(TRANSACTION_ID_CORRECT)
	assert.Nil(t, err)
	assert.True(t, transaction.ID == TRANSACTION_ID_CORRECT)
}

func TestGetTransaction_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	transaction, err := service.GetTransaction("0x158f74772af1bf9e5d1eb9d6633bb6a602eea97bbbd552b16696d7d1d3fa007703000000")
	assert.Nil(t, err)
	assert.Nil(t, transaction)
}

func TestListHooks_WithModuleID_Success(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	moduleID := HOOK_LOOKUP_MODULE_ID_CORRECT
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
	hook, err := service.GetHook(HOOK_ID_CORRECT)
	assert.Nil(t, err)
	assert.True(t, hook.ID == HOOK_ID_CORRECT)

}

func TestGetHook_NotFound_Failure(t *testing.T) {
	service := CreateTheGraphServiceAlpha()
	hook, err := service.GetHook("0xc0f6e387ac0b324ec18eacf22ee7271207dce2d5")
	assert.Nil(t, err)
	assert.Nil(t, hook)
}

// func TestListLicenses_WithIpOrgIdAndIpAssetId_Success(t *testing.T) {
// 	service := CreateTheGraphServiceAlpha()
// 	ipOrgId := "0xb422e54932c1dae83e78267a4dd2805aa64a8061"
// 	ipAssetId := "0"
// 	licenses, err := service.ListLicenses(&ipOrgId, &ipAssetId, nil)
// 	assert.Nil(t, err)
// 	assert.True(t, len(licenses) > 0)
// }

// func TestListLicenses_WithIpOrgId_Success(t *testing.T) {
// 	service := CreateTheGraphServiceAlpha()
// 	ipOrgId := "0xb422e54932c1dae83e78267a4dd2805aa64a8061"
// 	licenses, err := service.ListLicenses(&ipOrgId, nil, nil)
// 	assert.Nil(t, err)
// 	assert.True(t, len(licenses) > 0)
// }

// func TestListLicenses_WithIpAssetId_Success(t *testing.T) {
// 	service := CreateTheGraphServiceAlpha()
// 	ipAssetId := "0"
// 	licenses, err := service.ListLicenses(nil, &ipAssetId, nil)
// 	assert.Nil(t, err)
// 	assert.True(t, len(licenses) > 0)
// }

// func TestListLicenses_WithoutIpOrgIdAndIpAssetId_Success(t *testing.T) {
// 	service := CreateTheGraphServiceAlpha()
// 	licenses, err := service.ListLicenses(nil, nil, nil)
// 	assert.Nil(t, err)
// 	assert.True(t, len(licenses) > 0)
// }

// func TestGetLicense_Success(t *testing.T) {
// 	service := CreateTheGraphServiceAlpha()
// 	license, err := service.GetLicense("1")
// 	assert.Nil(t, err)
// 	assert.True(t, license.ID == "1")
// 	assert.True(t, license.CreatedAt == "2023-11-21T02:20:24Z")
// }

// func TestGetLicense_NotFound_Failure(t *testing.T) {
// 	service := CreateTheGraphServiceAlpha()
// 	license, err := service.GetLicense("1123")
// 	assert.Nil(t, err)
// 	assert.Nil(t, license)
// }
