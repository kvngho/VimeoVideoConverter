package vimeo

import "time"

type Response struct {
	Play struct {
		Progressive []struct {
			Type               string    `json:"type"`
			Codec              string    `json:"codec"`
			Width              int       `json:"width"`
			Height             int       `json:"height"`
			LinkExpirationTime time.Time `json:"link_expiration_time"`
			Link               string    `json:"link"`
			CreatedTime        time.Time `json:"created_time"`
			Fps                float64   `json:"fps"`
			Size               int       `json:"size"`
			Md5                string    `json:"md5"`
			Rendition          string    `json:"rendition"`
		} `json:"progressive"`
		Hls struct {
			LinkExpirationTime time.Time `json:"link_expiration_time"`
			Link               string    `json:"link"`
		} `json:"hls"`
		Dash struct {
			LinkExpirationTime time.Time `json:"link_expiration_time"`
			Link               string    `json:"link"`
		} `json:"dash"`
		Status string `json:"status"`
	} `json:"play"`
}
