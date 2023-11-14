package v0mvp

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/api/internal/service/thegraph"
	"github.com/project-nova/backend/pkg/logger"
)

func HandleGetIpAsset(c *gin.Context, graphService thegraph.TheGraphServiceMvp) (*entity.GetIpAssetResponseMVP, int, error) {
	ipAssetId := c.Param("ipAssetId")
	_, err := strconv.ParseInt(ipAssetId, 10, 64)
	if err != nil {
		logger.Errorf("Invalid ip asset id: %s", ipAssetId)
		return nil, http.StatusBadRequest, fmt.Errorf("Invalid ip asset id")
	}

	franchiseId := c.DefaultQuery("franchiseId", "")
	_, err = strconv.ParseInt(franchiseId, 10, 64)
	if err != nil {
		logger.Errorf("Invalid franchise id: %s", franchiseId)
		return nil, http.StatusBadRequest, fmt.Errorf("invalid franchise id")
	}

	ipAsset, err := graphService.GetIpAsset(franchiseId, ipAssetId)
	if err != nil {
		logger.Errorf("Failed to get ip asset: %v", err)
		return nil, http.StatusInternalServerError, fmt.Errorf("Internal server error")
	}

	return &entity.GetIpAssetResponseMVP{
		Data: ipAsset,
	}, http.StatusOK, nil
}
