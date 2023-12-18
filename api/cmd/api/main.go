package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/project-nova/backend/api/internal/config"
	"github.com/project-nova/backend/api/internal/handler"
	"github.com/project-nova/backend/api/internal/service/thegraph"
	xhttp "github.com/project-nova/backend/pkg/http"
	"github.com/project-nova/backend/pkg/logger"
)

func main() {
	r := gin.Default()

	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	Logger, err := logger.InitLogger(logger.Levels.Info)
	if err != nil {
		logger.Fatalf("Failed to init logger, error: %v", err)
	}
	defer func() {
		_ = Logger.Sync()
	}()

	cfg, err := config.GetConfig()
	if err != nil {
		logger.Fatalf("Failed to init config, error: %v", err)
	}

	httpClient := xhttp.NewClient(&xhttp.ClientConfig{})

	theGraphClientAlpha := graphql.NewClient("https://api.thegraph.com/subgraphs/name/edisonz0718/storyprotocol-v0-alpha")
	theGraphServiceAlpha := thegraph.NewTheGraphServiceAlphaImpl(theGraphClientAlpha)

	// initialize handlers
	protocolHandler := handler.NewProtocolHandler(theGraphServiceAlpha, httpClient)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello")
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Healthy")
	})

	// Alpha endpoints
	protocol := r.Group("/protocol")
	protocol.Use(cors.Default())
	{
		// Endpoint to list ip orgs
		protocol.POST("/iporg", protocolHandler.ListIpOrgsHandler)

		// Endpoint to get an ip org
		protocol.GET("/iporg/:ipOrgId", protocolHandler.GetIpOrgHandler)

		// Endpoint to get a single ip asset
		protocol.GET("/ipasset/:ipAssetId", protocolHandler.GetIpAssetHandler)

		// Endpoint to list ip assets
		protocol.POST("/ipasset", protocolHandler.ListIpAssetsHandler)

		// Endpoint to get a single license
		protocol.GET("/license/:licenseId", protocolHandler.GetLicenseHandler)

		// Endpoint to list licenses from an ip asset
		protocol.POST("/license", protocolHandler.ListLicensesHandler)

		// Endpoint to get transaction
		protocol.GET("/transaction/:transactionId", protocolHandler.GetTransactionHandler)

		// Endpoint to list transactions
		protocol.POST("/transaction", protocolHandler.ListTransactionsHandler)

		// Endpoint to get a relatioinship
		protocol.GET("/relationship/:relationshipId", protocolHandler.GetRelationshipHandler)

		// Endpoint to list relatioinships
		protocol.POST("/relationship", protocolHandler.ListRelationshipsHandler)

		// Endpoint to get a relastionship type
		protocol.GET("/relationship-type", protocolHandler.GetRelationshipTypeHandler)

		// Endpoint to list relastionship types
		protocol.POST("/relationship-type", protocolHandler.ListRelationshipTypesHandler)

		// Endpoint to list modules
		protocol.POST("/module", protocolHandler.ListModulesHandler)

		// Endpoint to get a module
		protocol.GET("/module/:moduleId", protocolHandler.GetModuleHandler)

		// Endpoint to list hooks
		protocol.POST("/hook", protocolHandler.ListHooksHandler)

		// Endpoint to get a hook
		protocol.GET("/hook/:hookId", protocolHandler.GetHookHandler)

		// Endpoint to list license params
		protocol.POST("/license-params", protocolHandler.ListLicenseParamsHandler)
	}

	port := fmt.Sprintf(":%d", cfg.Port)
	_ = r.Run(port)
}
