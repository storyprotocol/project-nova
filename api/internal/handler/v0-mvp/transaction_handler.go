package v0mvp

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/api/internal/service/thegraph"
	"github.com/project-nova/backend/pkg/logger"
)

func HandleGetTransaction(c *gin.Context, graphService thegraph.TheGraphServiceMvp) (*entity.GetTransactionResponseMVP, int, error) {
	transactionId := c.Param("transactionId")
	if !strings.HasPrefix(transactionId, "0x") {
		logger.Errorf("Invalid transaction id: %s", transactionId)
		return nil, http.StatusBadRequest, fmt.Errorf("Invalid transaction id")
	}

	transaction, err := graphService.GetTransaction(transactionId)
	if err != nil {
		logger.Errorf("Failed to get transaction: %v", err)
		return nil, http.StatusInternalServerError, fmt.Errorf("Internal server error")
	}

	return &entity.GetTransactionResponseMVP{
		Data: transaction,
	}, http.StatusOK, nil
}
