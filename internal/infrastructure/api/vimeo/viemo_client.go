package vimeo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type Client struct {
	Token string
}

func NewVimeoClient(token string) *Client {
	return &Client{
		Token: token,
	}
}

func (c *Client) FetchVideoData(url string) (string, error) {
	videoID, err := FindVideoID(url)
	if err != nil {
		return "", err
	}

	requestURL := fmt.Sprintf("https://api.vimeo.com/videos/%s?fields=play", videoID)
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Accept", "application/vnd.vimeo.*+json;version=3.4")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var vimeoResp Response
	err = json.Unmarshal(body, &vimeoResp)
	if err != nil {
		return "", err
	}
	if len(vimeoResp.Play.Progressive) == 0 {
		return "", fmt.Errorf("존재하지 않는 비디오입니다")
	}
	return vimeoResp.Play.Progressive[0].Link, nil
}

func FindVideoID(url string) (string, error) {
	// 정규 표현식 패턴 정의
	re := regexp.MustCompile(`/(\d+)/?$`)

	// 패턴에 맞는 부분 찾기
	match := re.FindStringSubmatch(url)

	if len(match) > 1 {
		return match[1], nil
	}

	return "", nil
}
