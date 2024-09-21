package controllers

import (
	"ai-content-creator-agent/internal/domain/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SocialMediaMgtController interface {
	PostOnFacebook(c *gin.Context)
	ReplyToComments(c *gin.Context)
}

type socialMediaMgtController struct {
	socialMediaMgtUsecase interfaces.SocialMediaMgtUsecase
}

// NewSocialMediaMgtController creates a new SocialMediaMgtController.
func NewSocialMediaMgtController(socialMediaMgtUsecase interfaces.SocialMediaMgtUsecase) SocialMediaMgtController {
	return &socialMediaMgtController{
		socialMediaMgtUsecase: socialMediaMgtUsecase,
	}
}

// PostOnFacebook handles the request to post content on Facebook.
func (smc *socialMediaMgtController) PostOnFacebook(c *gin.Context) {
	var CompanyID struct {
		ID string `json:"id"`
	}

	CompanyID.ID = c.Param("id")
	// if err := c.ShouldBindJSON(&CompanyID); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	err := smc.socialMediaMgtUsecase.PostContent(CompanyID.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Content posted successfully"})
}

// ReplyToComments handles the request to reply to comments on Facebook.
func (smc *socialMediaMgtController) ReplyToComments(c *gin.Context) {
	var CompanyID struct {
		ID string `json:"id"`
	}

	if err := c.ShouldBindJSON(&CompanyID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := smc.socialMediaMgtUsecase.ReplyToComments(CompanyID.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Replied to comments successfully"})
}

