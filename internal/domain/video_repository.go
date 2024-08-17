package domain

type VideoRepository interface {
	FetchVideoData(URL string) (string, error)
}
