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

func (svc *VideoUpdateService) UpdateInformationALL() (int, error) {
	var count int
	reviews, err := svc.Database.GetAllReviews()
	if err != nil {
		return 0, err
	}
	for _, review := range reviews {
		_, err = svc.UpdateInformation(review.ReviewVideo, "review")
		count++
	}
	talks, err := svc.Database.GetAllTalks()
	if err != nil {
		return 0, err
	}
	for _, talk := range talks {
		_, err = svc.UpdateInformation(talk.TalkVideo, "talk")
		count++
	}
	productVideos, err := svc.Database.GetAllProductVideos()
	if err != nil {
		return 0, err
	}
	for _, productVideo := range productVideos {
		_, err = svc.UpdateInformation(productVideo.URL, "admin")
		count++
	}
	userVideos, err := svc.Database.GetAllUserVideos()
	if err != nil {
		return 0, err
	}
	for _, userVideo := range userVideos {
		_, err = svc.UpdateInformation(userVideo.VideoURL, "user")
		count++
	}
	return count, nil

}
