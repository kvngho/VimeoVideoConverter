package ent

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kvngho/vimeovideoconverter/internal/infrastructure/persistence/ent/productvideo"
	"github.com/kvngho/vimeovideoconverter/internal/infrastructure/persistence/ent/uservideo"
	"github.com/rs/zerolog/log"
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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		cfg.Username, cfg.Password, cfg.URL, cfg.Port, cfg.Name)
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
		_, err = videoObj.Update().SetURL(convertedURL).Save(context.Background())
		if err != nil {
			return err
		}

		return nil
	case "user":
		videoObj, err := ec.client.UserVideo.Query().Where(uservideo.VideoURL(url)).First(context.Background())
		if err != nil {
			return err
		}
		_, err = videoObj.Update().SetVideoURL(convertedURL).Save(context.Background())
		if err != nil {
			return err
		}
	default:
		return nil
	}
	return nil
}
