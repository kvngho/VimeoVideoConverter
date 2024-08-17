package persistence

type VideoRepository interface {
	Update(url string, convertedURL string, videoType string) error
}
