package v0alpha

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v0alpha "github.com/project-nova/backend/api/internal/entity/v0-alpha"
)

func HandleGetTransaction(c *gin.Context) (*v0alpha.GetTransactionResponse, int, error) {
	return &v0alpha.GetTransactionResponse{
		Transaction: &v0alpha.Transaction{
			ID:           "1",
			TxHash:       "0x00a1a14e0193144e1d7024428ee242c44e5cacdbd7458c629d17c6366f6c5cb6",
			IPOrgId:      "7",
			ResourceId:   "1",
			ResourceType: v0alpha.ResourceTypes.IPAsset,
			ActionType:   v0alpha.ActionTypes.Create,
			Creator:      "0x4f9693ac46f2c7e2f48dd14d8fe1ab44192cd57d",
			CreatedAt:    "0001-01-01T00:00:00Z",
		},
	}, http.StatusOK, nil
}
