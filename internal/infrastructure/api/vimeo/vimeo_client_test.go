package vimeo_test

import (
	"github.com/kvngho/vimeovideoconverter/internal/infrastructure/api/vimeo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindVideoID(t *testing.T) {
	url := "http://vimeo.com/1234"
	url1 := "http://vimeo.com/12345/"

	id1, err := vimeo.FindVideoID(url)
	assert.NoError(t, err)
	assert.Equal(t, id1, "1234")

	id2, err := vimeo.FindVideoID(url1)
	assert.NoError(t, err)
	assert.Equal(t, id2, "12345")
}

func TestClient_FetchVideoData(t *testing.T) {
	url := ""
	token := ""
	vimeoClient := vimeo.NewVimeoClient(token)
	_, err := vimeoClient.FetchVideoData(url)
	assert.NoError(t, err)
}
