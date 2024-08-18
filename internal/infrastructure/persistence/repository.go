package persistence

type VideoRepository interface {
	Update(url string, convertedURL string, videoType string) error
	GetAllReviews() ([]*ProductReview, error)
	GetAllUserVideos() ([]*UserVideo, error)
	GetAllProductVideos() ([]*ProductVideo, error)
	GetAllTalks() ([]*DeepingTalk, error)
}
