package entities

import (
	"errors"
	"time"
)

// Company represents a Company entity with all necessary information for content creation.
type Company struct {
    ID                  string    `json:"id" firestore:"id"`                                       // Unique identifier for the Company.
    CompanyName         string    `json:"company_name," binding:"required" firestore:"company_name"`                   // The official name of the company.
    Industry            string    `json:"industry," binding:"required" firestore:"industry"`                           // The industry in which the company operates.
    MissionStatement    string    `json:"mission_statement," binding:"required" firestore:"mission_statement"`         // A brief description of the company's mission and core values.
    BrandVoice          string    `json:"brand_voice," binding:"required" firestore:"brand_voice"`                     // The tone and style of communication.
    LogoURL             string    `json:"logo_url," binding:"required" firestore:"logo_url"`                           // URL to the company's logo.
    TargetAudience      string    `json:"target_audience," binding:"required" firestore:"target_audience"`             // The target audience information.
    ContentGoals        []string  `json:"content_goals," binding:"required" firestore:"content_goals"`                 // What the Company wants to achieve with their content.
    KeyMessages         []string  `json:"key_messages," binding:"required" firestore:"key_messages"`                   // The main messages or themes the content should convey.
    FacebookPageID      string    `json:"facebook_page_id," binding:"required" firestore:"facebook_page_id"`           // The Facebook page ID of the company.
    FacebookAccessToken string    `json:"facebook_access_token," binding:"required" firestore:"facebook_access_token"` // The Facebook access token of the company.
    PostingFrequency    string    `json:"posting_frequency," binding:"required" firestore:"posting_frequency"`         // Frequency of content posting.
    CreatedAt           time.Time `json:"created_at," firestore:"created_at"`                       // When the Company record was created.
    UpdatedAt           time.Time `json:"updated_at," firestore:"updated_at"`                       // When the Company record was last updated.
}

// NewCompany initializes a new Company with the given information.
func NewCompany(id, companyName, industry, missionStatement, brandVoice, logoURL, targetAudience string, contentGoals, keyMessages []string, facebookPageID, facebookAcessToken, postingFrequency string) (*Company, error) {
	// Validation for mandatory fields
	if companyName == "" || industry == "" || missionStatement == "" {
		return nil, errors.New("mandatory fields cannot be empty")
	}

	// Create the Company object
	Company := &Company{
		ID:                  id,
		CompanyName:         companyName,
		Industry:            industry,
		MissionStatement:    missionStatement,
		BrandVoice:          brandVoice,
		LogoURL:             logoURL,
		TargetAudience:      targetAudience,
		ContentGoals:        contentGoals,
		KeyMessages:         keyMessages,
		FacebookPageID:      facebookPageID,
		FacebookAccessToken: facebookAcessToken,
		PostingFrequency:    postingFrequency,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	return Company, nil
}

// // UpdateCompany updates the existing Company information.
// func (c *Company) UpdateCompany(companyName, industry, missionStateme, brandVoice, logoURL string, colorScheme [ string, targetAudience string, contentGoals, keyMessages, contentPillars, callToActions []string, socialMediaProfiles map[string]string, privacyPolicy string) error {
// 	if companyName == "" || industry == "" || missionStatement == "" {
// 		return errors.New("mandatory fields cannot be empty")
// 	}

// 	// Update Company details
// 	c.CompanyName = companyName
// 	c.Industry = industry
// 	c.MissionStatement = missionStatement
// 	c.BrandVoice = brandVoice

// 	c.LogoURL = logoURL
// 	c.ColorScheme = colorScheme

// 	c.TargetAudience = targetAudience
// 	c.ContentGoals = contentGoals
// 	c.KeyMessages = keyMessages
// 	c.ContentPillars = contentPillars
// 	c.CallToActions = callToActions
// 	c.SocialMediaProfiles = socialMediaProfiles
// 	c.PrivacyPolicy = privacyPolicy
// 	c.UpdatedAt = time.Now()

// 	return nil
// }

// // ValidateSocialMediaProfiles ensures that social media profiles contain valid URLs.
// func (c *Company) ValidateSocialMediaProfiles() error {
// 	for platform, url := range c.SocialMediaProfiles {
// 		if !strings.HasPrefix(url, "http") {
// 			return errors.New(platform + " profile URL is invalid")
// 		}
// 	}
// 	return nil
// }
