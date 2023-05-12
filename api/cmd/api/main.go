package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/config"
	"github.com/project-nova/backend/api/internal/handler"
	"github.com/project-nova/backend/api/internal/repository"
	"github.com/project-nova/backend/pkg/constant"
	"github.com/project-nova/backend/pkg/database"
	"github.com/project-nova/backend/pkg/keymanagement"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/middleware"
	"github.com/project-nova/backend/pkg/s3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()

	flag.Parse()

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

	franchiseMap := map[string]*config.FranchiseConfig{}
	for _, franchise := range cfg.Protocol.FranchiseMap {
		franchiseMap[franchise.FranchiseInfo.Address] = franchise
	}

	db, err := database.NewGormDB(cfg.DatabaseURI)
	if err != nil {
		logger.Fatalf("Failed to connect to DB, error: %v", err)
	}

	ethClient, err := ethclient.Dial(cfg.ProviderURL)
	if err != nil {
		logger.Fatalf("Failed to connect to the blockchain provider, error: %v", err)
	}

	kmsClient := keymanagement.NewKmsClient(cfg.Region)

	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.Region),
	})
	if err != nil {
		logger.Fatalf("Failed to create aws session: %v", err)
	}
	s3Client := s3.NewS3Client(awsSession)

	walletMerkleProofRepository := repository.NewWalletMerkleProofDbImpl(db)
	nftTokenRepository := repository.NewNftTokenDbImpl(db)
	nftCollectionRepository := repository.NewNftCollectionDbImpl(db)
	nftAllowlistRepository := repository.NewNftAllowlistDbImpl(db)
	storyChapterRepository := repository.NewStoryChapterDbImpl(db)
	storyInfoRepository := repository.NewStoryInfoDbImpl(db)

	storyContentRepository, err := repository.NewStoryContentS3Impl(s3Client, cfg.S3ContentBucketName)
	if err != nil {
		logger.Fatalf("Failed to init story content s3 implementation: %v", err)
	}

	protocolStoryContentRepository := repository.NewProtocolStoryContentDbImpl(db)

	franchiseCollectionRepository := repository.NewFranchiseCollectionDbImpl(db)
	err = franchiseCollectionRepository.GetAndLoadFranchiseCollections()
	if err != nil {
		logger.Errorf("Failed to get and load franchise collections: %v", err)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello")
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Healthy")
	})

	r.GET(constant.MetricsPath, prometheusHandler())

	publicV1 := r.Group("/v1")
	publicV1.Use(cors.Default())
	publicV1.Use(middleware.Prometheus(cfg.AppID))
	{
		// Endpoint to get the merkle proof for the wallet address per allowlist
		publicV1.GET("/wallet/:walletAddress/proof", handler.NewGetWalletProofHandler(walletMerkleProofRepository))

		// Endpoint to return signing message
		publicV1.GET("/wallet/:walletAddress/sign-message", handler.NewGetWalletSignMessageHandler())

		// Endpoint to get all story chapters' information
		publicV1.GET("/story/:franchiseId/:storyNum", handler.NewGetStoryChaptersHandler(storyChapterRepository, storyInfoRepository))

		// Endpoint to get story chapter contents
		publicV1.GET("/story/:franchiseId/:storyNum/:chapterNum", handler.NewGetStoryChapterContentsHandler(storyContentRepository, storyChapterRepository, storyInfoRepository))

		// Endpoint to get the metadata of story nfts
		publicV1.POST("/nft/list", handler.NewGetNftsHandler(nftTokenRepository, franchiseCollectionRepository))

		// Endpoint to update nft backstory for the nft owner
		publicV1.POST("/nft/:id/backstory", handler.NewUpdateNftBackstoryHandler(nftTokenRepository))

		// Endpoint to get the metadata of nft collection
		publicV1.GET("/nft/collections", handler.NewGetNftCollectionsHandler(nftCollectionRepository, franchiseCollectionRepository))

	}

	adminV1 := r.Group("/admin/v1")
	adminV1.Use(middleware.AuthAdmin(kmsClient, []byte(cfg.AdminAuthMessage), cfg.AuthKeyId))
	{
		// Admin Endpoint to create wallet proof
		adminV1.POST("/wallet/:walletAddress/proof", handler.NewAdminCreateWalletProofHandler(walletMerkleProofRepository))

		// Admin Endpoint to get collection data
		adminV1.GET("/nft/collections", handler.NewAdminGetCollectionsHandler(nftCollectionRepository, ethClient))

		// Admin Endpoint to update collection address
		adminV1.POST("/nft/collection/:address", handler.NewAdminUpdateCollectionAddressHandler(nftCollectionRepository, franchiseCollectionRepository, nftAllowlistRepository))

		// Admin Endpoint to fetch and create nft metadata
		adminV1.POST("/nft/:id", handler.NewAdminCreateOrUpdateNftHandler(nftTokenRepository, ethClient))

		// Admin Endpoint to update nft owner address
		adminV1.POST("/nft/:id/owner", handler.NewAdminUpdateNftOwnerHandler(nftTokenRepository, ethClient))

		// Admin Endpoint to delete nft
		adminV1.DELETE("/nft/:id", handler.NewAdminDeleteNftHandler(nftTokenRepository))

		// Admin Endpoint to add a story chapter
		adminV1.POST("/story/:franchiseId/:storyNum/:chapterNum", handler.NewAdminCreateStoryChapterHandler(storyChapterRepository, storyInfoRepository))

		// Admin Endpoint to update chapter content to cache
		adminV1.POST("/story/:franchiseId/:storyNum/:chapterNum/cache", handler.NewAdminUpdateStoryChapterCacheHandler(storyContentRepository))
	}

	protocolV1 := r.Group("/protocol/v1")
	protocolV1.Use(cors.Default())
	{
		// Endpoint to list all franchise
		protocolV1.GET("/franchise", handler.NewGetFranchisesHandler(ethClient, franchiseMap))

		// Endpoint to get franchise collections
		protocolV1.GET("/franchise/:franchiseAddress", handler.NewGetFranchiseCollectionsHandler(ethClient, franchiseMap))

		// Endpoint to get characters per collection
		protocolV1.GET("/character/:franchiseAddress/:collectionAddress", handler.NewGetCharactersHandler(ethClient, franchiseMap))

		// Endpoint to get a single character per collection
		protocolV1.GET("/character/:franchiseAddress/:collectionAddress/:tokenId", handler.NewGetCharacterHandler(ethClient, franchiseMap))

		// Endpoint to get collectors of a character
		protocolV1.GET("/collector/:franchiseAddress/:collectionAddress/:tokenId", handler.NewGetCollectorsHandler(ethClient, franchiseMap))

		// Endpoint to get stories per collection
		protocolV1.GET("/story/:franchiseAddress/:collectionAddress", handler.NewGetStoriesHandler(ethClient, franchiseMap))

		// Endpoint to get a single story per collection
		protocolV1.GET("/story/:franchiseAddress/:collectionAddress/:tokenId", handler.NewGetStoryHandler(ethClient, franchiseMap))

		// Endpoint to get a story content
		protocolV1.GET("/story/content/:contentId", handler.NewGetStoryContentHandler(protocolStoryContentRepository))

		// Endpoint to post story content
		protocolV1.POST("/story/:franchiseAddress/:collectionAddress/content", handler.NewPostStoryContentHandler(protocolStoryContentRepository, cfg.Protocol.ContentUri))

		// Endpoint to get derivative stories of a story
		protocolV1.GET("/story/:franchiseAddress/:collectionAddress/:tokenId/derivatives", handler.NewGetDerivativesHandler())

		// Endpoint to get license information for an asset
		protocolV1.GET("/license/:franchiseAddress/:collectionAddress/:tokenId", handler.NewGetAssetLicensesHandler(ethClient, franchiseMap, cfg.PrimitiveTpeAbiPath))
	}

	port := fmt.Sprintf(":%d", cfg.Port)
	_ = r.Run(port)
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
