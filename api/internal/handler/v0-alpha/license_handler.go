package v0alpha

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v0alpha "github.com/project-nova/backend/api/internal/entity/v0-alpha"
)

func HandleGetLicense(c *gin.Context) (*v0alpha.GetLicenseResponse, int, error) {
	return &v0alpha.GetLicenseResponse{
		License: &v0alpha.License{
			ID:          "49",
			IPAssetId:   "5",
			IPOrgId:     "7",
			Owner:       "0xd84316a1b6f40902c17b8177854cdaeb3c957daf",
			MetadataUri: "https://arweave.net/R7-xPDAMqOhUSw3CM_UwXI7zdpQkzCCCUq3smzxyAaU",
			CreatedAt:   "2023-11-14T00:29:13Z",
			TxHash:      "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
		},
	}, http.StatusOK, nil
}
