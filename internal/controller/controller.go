package controller

import (
	"encoding/json"
	"github.com/kvngho/vimeovideoconverter/internal/worker_queue"
	"github.com/rs/zerolog/log"
	"net/http"
)

type VideoController struct {
	queue *worker_queue.WorkerPool
}

type UserRequest struct {
	VideoURL  string `json:"video_url"`
	VideoType string `json:"video_type"`
}

func NewVideoController(queue *worker_queue.WorkerPool) *VideoController {
	return &VideoController{
		queue: queue,
	}
}

func (vc *VideoController) UpdateVideo(w http.ResponseWriter, r *http.Request) {
	var userRequest UserRequest

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Info().Str("video_type", userRequest.VideoType).Str("video_url", userRequest.VideoURL).Msg("get requests")
	vc.queue.SubmitJob(userRequest.VideoURL, userRequest.VideoType)
	res := struct {
		Status string `json:"status"`
	}{
		Status: "success",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return

}
