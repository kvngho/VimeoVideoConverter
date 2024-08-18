package ent

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kvngho/vimeovideoconverter/internal/infrastructure/persistence"
	"github.com/kvngho/vimeovideoconverter/internal/infrastructure/persistence/ent/deepingtalk"
	"github.com/kvngho/vimeovideoconverter/internal/infrastructure/persistence/ent/productreview"
	"github.com/kvngho/vimeovideoconverter/internal/infrastructure/persistence/ent/productvideo"
	"github.com/kvngho/vimeovideoconverter/internal/infrastructure/persistence/ent/uservideo"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

type Config struct {
	Username string
	Password string
	URL      string
	Port     string
	Name     string
}
type EntClient struct {
	client *Client
}

func NewEntClient(cfg Config) *EntClient {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&tls=custom",
		cfg.Username, cfg.Password, cfg.URL, cfg.Port, cfg.Name)
	rootCertPool := x509.NewCertPool()
	pem, err := os.ReadFile("cert.pem")
	if err != nil {
		log.Error().Err(err).Msgf("failed to read cert.pem")
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Error().Msgf("failed to append certs")
	}
	err = mysql.RegisterTLSConfig("custom", &tls.Config{
		RootCAs: rootCertPool,
	})
	if err != nil {
		log.Error().Err(err).Msgf("failed to register custom TLS config")
	}

	drv, err := sql.Open(dialect.MySQL, dsn)
	if err != nil {
		log.Error().Err(err).Msg("failed opening connection to mysql")
	}

	client := NewClient(Driver(drv))
	return &EntClient{client: client}
}

func (ec *EntClient) Close() error {
	err := ec.client.Close()
	if err != nil {
		return err
	}
	return nil
}

func (ec *EntClient) Update(url string, convertedURL string, videoType string) error {
	switch strings.ToLower(videoType) {
	case "admin":
		videoObj, err := ec.client.ProductVideo.Query().Where(productvideo.URL(url)).First(context.Background())
		if err != nil {
			return err
		}
		_, err = videoObj.Update().SetPlayableVideo(convertedURL).Save(context.Background())
		if err != nil {
			return err
		}
	case "user":
		videoObj, err := ec.client.UserVideo.Query().Where(uservideo.VideoURL(url)).First(context.Background())
		if err != nil {
			return err
		}
		_, err = videoObj.Update().SetPlayableVideo(convertedURL).Save(context.Background())
		if err != nil {
			return err
		}
	case "talk":
		videoObj, err := ec.client.DeepingTalk.Query().Where(deepingtalk.TalkVideo(url)).First(context.Background())
		if err != nil {
			return err
		}
		_, err = videoObj.Update().SetPlayableVideo(convertedURL).Save(context.Background())
		if err != nil {
			return err
		}
	case "review":
		videoObj, err := ec.client.ProductReview.Query().Where(productreview.ReviewVideo(url)).First(context.Background())
		if err != nil {
			return err
		}
		_, err = videoObj.Update().SetPlayableVideo(convertedURL).Save(context.Background())
		if err != nil {
			return err
		}
	default:
		return nil
	}
	return nil
}

func (ec *EntClient) GetAllReviews() ([]*persistence.ProductReview, error) {
	var resultReviews []*persistence.ProductReview
	reviews, err := ec.client.ProductReview.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	for _, review := range reviews {
		resultReviews = append(resultReviews, &persistence.ProductReview{
			PlayableVideo: review.PlayableVideo,
			ReviewVideo:   review.ReviewVideo,
		})
	}
	return resultReviews, nil
}

func (ec *EntClient) GetAllUserVideos() ([]*persistence.UserVideo, error) {
	var userVideos []*persistence.UserVideo
	videos, err := ec.client.UserVideo.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	for _, video := range videos {
		userVideos = append(userVideos, &persistence.UserVideo{
			PlayableVideo: video.PlayableVideo,
			VideoURL:      video.VideoURL,
		})
	}
	return userVideos, nil
}

func (ec *EntClient) GetAllProductVideos() ([]*persistence.ProductVideo, error) {
	var productVideos []*persistence.ProductVideo
	videos, err := ec.client.ProductVideo.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	for _, video := range videos {
		productVideos = append(productVideos, &persistence.ProductVideo{
			PlayableVideo: video.PlayableVideo,
			URL:           video.URL,
		})
	}
	return productVideos, nil
}

func (ec *EntClient) GetAllTalks() ([]*persistence.DeepingTalk, error) {
	var talkVideos []*persistence.DeepingTalk
	videos, err := ec.client.DeepingTalk.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	for _, video := range videos {
		talkVideos = append(talkVideos, &persistence.DeepingTalk{
			PlayableVideo: video.PlayableVideo,
			TalkVideo:     video.TalkVideo,
		})
	}
	return talkVideos, nil
}
