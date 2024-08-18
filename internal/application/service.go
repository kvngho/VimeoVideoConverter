package application

import (
	"github.com/kvngho/vimeovideoconverter/internal/domain"
	"github.com/kvngho/vimeovideoconverter/internal/infrastructure/persistence"
	"github.com/rs/zerolog/log"
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
		if review.ReviewVideo == "" {
			continue
		}
		_, err = svc.UpdateInformation(review.ReviewVideo, "review")
		count++
	}
	log.Info().Int("count", count).Msg("review completed")

	talks, err := svc.Database.GetAllTalks()
	if err != nil {
		return 0, err
	}
	for _, talk := range talks {
		if talk.TalkVideo == "" {
			continue
		}
		_, err = svc.UpdateInformation(talk.TalkVideo, "talk")
		count++
	}
	log.Info().Int("count", count).Msg("talk completed")

	productVideos, err := svc.Database.GetAllProductVideos()
	if err != nil {
		return 0, err
	}
	for _, productVideo := range productVideos {
		if productVideo.URL == "" {
			continue
		}
		_, err = svc.UpdateInformation(productVideo.URL, "admin")
		count++
	}
	log.Info().Int("count", count).Msg("admin completed")

	userVideos, err := svc.Database.GetAllUserVideos()
	if err != nil {
		return 0, err
	}
	for _, userVideo := range userVideos {
		if userVideo.VideoURL == "" {
			continue
		}
		_, err = svc.UpdateInformation(userVideo.VideoURL, "user")
		count++
	}
	log.Info().Int("count", count).Msg("user completed")
	return count, nil

}
