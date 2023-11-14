package v0alpha

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v0alpha "github.com/project-nova/backend/api/internal/entity/v0-alpha"
)

func HandleGetIpAsset(c *gin.Context) (*v0alpha.GetIpAssetResponse, int, error) {
	return &v0alpha.GetIpAssetResponse{
		IPAsset: &v0alpha.IPAsset{
			ID:          "1",
			IPOrgId:     "7",
			Type:        v0alpha.IpAssetTypes.Story,
			Name:        "The Empire Strikes Back",
			Owner:       "0x4f9693ac46f2c7e2f48dd14d8fe1ab44192cd57d",
			MetadataUrl: "https://arweave.net/R7-xPDAMqOhUSw3CM_UwXI7zdpQkzCCCUq3smzxyAaU",
			CreatedAt:   "2023-11-14T00:29:13Z",
			TxHash:      "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
		},
	}, http.StatusOK, nil
}
