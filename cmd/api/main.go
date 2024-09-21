package main

import (
	"ai-content-creator-agent/internal/adapters/controllers"
	"ai-content-creator-agent/internal/adapters/repositories"
	"ai-content-creator-agent/internal/adapters/services"
	"ai-content-creator-agent/internal/domain/interfaces"
	"ai-content-creator-agent/internal/domain/usecases"
	"ai-content-creator-agent/internal/infrastructure/api"
	"ai-content-creator-agent/internal/infrastructure/config"

	// "ai-content-creator-agent/pkg/logger"
	// "ai-content-creator-agent/pkg/utils"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	configuration, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	// Initialize logger
	// logger.InitLogger()

	// Initialize services
	nlpService := services.NewNLPService(configuration.GeminiAPIKey)
	scheduleService := services.NewScheduleService()
	// sentimentAnalayzer := services.NewSentimentAnalysisService(configuration.GeminiAPIKey)

	// Initialize api services
	newsApi := api.NewNewsAPI(configuration.NewsAPIKey)

	// Initialize repositories
	companyRepo := repositories.NewCompanyRepository(context.Background(), configuration.ProjectID)
	contentRepo := repositories.NewContentRepository(context.Background(), configuration.ProjectID)

	// Initialize usecases
	researchUsecase := usecases.NewResearchUsecase(newsApi, nlpService)
	companyUsecase := usecases.NewCompanyUsecase(companyRepo, scheduleService)
	contentUsecase := usecases.NewContentUsecase(contentRepo, researchUsecase, nlpService)
	socialMediaUsecase := usecases.NewSocialMediaMgtUsecase(contentUsecase, companyUsecase)

	// Initialize Controllers
	companyController := controllers.NewCompanyController(companyUsecase)
	socialMediaController := controllers.NewSocialMediaMgtController(socialMediaUsecase)

	// Initialize social media management
	// Initialize router
	router := setupRouter(companyController, socialMediaController)

	// Start server
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func setupRouter(companyController interfaces.CompanyController, socialMediaController controllers.SocialMediaMgtController) *gin.Engine {
	router := gin.Default()

	// Initialize controllers
	router.POST("/post-content/:id", socialMediaController.PostOnFacebook)
	router.POST("/reply/:id", socialMediaController.ReplyToComments)
	router.POST("company/register", companyController.RegisterCompany)
	router.GET("company/:id", companyController.GetCompany)

	return router
}
