package interfaces


type SocialMediaMgtController interface {
	PostContent(companyID string) error
	ReplyToComments(companyID string) error
}