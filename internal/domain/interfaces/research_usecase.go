package interfaces

type ResearchUsecase interface {
	ResearchNews(industry string) (string, error)
}