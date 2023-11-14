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

func HandleGetLicense(c *gin.Context, graphService thegraph.TheGraphServiceMvp) (*entity.GetLicenseResponseMVP, int, error) {
	licenseId := c.Param("licenseId")
	_, err := strconv.ParseInt(licenseId, 10, 64)
	if err != nil {
		logger.Errorf("Invalid license id: %s", licenseId)
		return nil, http.StatusBadRequest, fmt.Errorf("invalid license id")
	}

	license, err := graphService.GetLicense(licenseId)
	if err != nil {
		logger.Errorf("Failed to get license: %v", err)
		return nil, http.StatusInternalServerError, fmt.Errorf("internal server error")
	}

	return &entity.GetLicenseResponseMVP{
		Data: license,
	}, http.StatusOK, nil
}
