package application

import (
	"github.com/kvngho/vimeovideoconverter/internal/domain"
	"github.com/kvngho/vimeovideoconverter/internal/infrastructure/persistence"
)

type VideoUpdateService struct {
	Database persistence.VideoRepository
	Video    domain.VideoRepository
}

func NewVideoUpdateService(database persistence.VideoRepository, video domain.VideoRepository) *VideoUpdateService {
	return &VideoUpdateService{
		Database: database,
		Video:    video,
	}
}

func (svc *VideoUpdateService) UpdateInformation(url string, videoType string) (string, error) {
	convertedURL, err := svc.Video.FetchVideoData(url)
	if err != nil {
		return "", err
	}

	err = svc.Database.Update(url, convertedURL, videoType)
	if err != nil {
		return "", err
	}
	return convertedURL, nil
}
