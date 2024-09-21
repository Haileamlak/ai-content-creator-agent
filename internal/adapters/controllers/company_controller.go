package controllers

import (
	"ai-content-creator-agent/internal/domain/entities"
	"ai-content-creator-agent/internal/domain/interfaces"
	"ai-content-creator-agent/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CompanyController struct {
	companyUsecase   interfaces.CompanyUsecase
}

func NewCompanyController(companyUsecase interfaces.CompanyUsecase) *CompanyController {
	return &CompanyController{
		companyUsecase:   companyUsecase,
	}
}

func (cc *CompanyController) RegisterCompany(c *gin.Context) {
	var companyRequest entities.Company

	err := c.ShouldBindJSON(&companyRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request" + err.Error()})
		return
	}

	companyRequest.CreatedAt = time.Now()
	companyRequest.UpdatedAt = time.Now()

	companyID, err := utils.GenerateRandomID(16)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate company ID"})
		return
	}

	companyRequest.ID = companyID

	err = cc.companyUsecase.RegisterCompany(&companyRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company: "  + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"company": companyRequest})
}

func (cc *CompanyController) GetCompany(c *gin.Context) {
	companyID := c.Param("id")

	company, err := cc.companyUsecase.GetCompany(companyID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "company not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"company": company})
}

func (cc *CompanyController) UpdateCompany(c *gin.Context) {
	companyID := c.Param("id")

	var companyRequest struct {
		CompanyName         string   `json:"companyName" binding:"required"`
		Industry            string   `json:"industry" binding:"required"`
		MissionStatement    string   `json:"missionStatement" binding:"required"`
		BrandVoice          string   `json:"brandVoice"`
		LogoURL             string   `json:"logoURL"`
		TargetAudience      string   `json:"targetAudience"`
		ContentGoals        []string `json:"contentGoals"`
		KeyMessages         []string `json:"keyMessages"`
		FacebookPageID      string   `json:"facebookPageID"`
		FacebookAccessToken string   `json:"facebookAccessToken"`
		PostingFrequency    string   `json:"postingFrequency"`
		PrivacyPolicy       string   `json:"privacyPolicy"`
	}

	err := c.ShouldBindJSON(&companyRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	company, err := entities.NewCompany(companyID, companyRequest.CompanyName, companyRequest.Industry, companyRequest.MissionStatement, companyRequest.BrandVoice, companyRequest.LogoURL, companyRequest.TargetAudience, companyRequest.ContentGoals, companyRequest.KeyMessages, companyRequest.FacebookPageID, companyRequest.FacebookAccessToken, companyRequest.PostingFrequency)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = cc.companyUsecase.UpdateCompany(company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update company"})
		return
	}

	c.JSON(http.StatusOK, company)
}

func (cc *CompanyController) DeleteCompany(c *gin.Context) {
	companyID := c.Param("id")

	err := cc.companyUsecase.DeleteCompany(companyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete company: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}