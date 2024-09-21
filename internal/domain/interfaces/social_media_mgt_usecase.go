package interfaces

type SocialMediaMgtUsecase interface {
	PostContent(companyID string) error
	ReplyToComments(companyID string) error
}