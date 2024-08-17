package worker_queue

import (
	"github.com/kvngho/vimeovideoconverter/internal/application"
	"github.com/rs/zerolog/log"
)

type JobDescription struct {
	videoURL  string
	videoType string
}

type WorkerPool struct {
	MaxWorkers int
	JobQueue   chan JobDescription
	JobAction  *application.VideoUpdateService
}

func NewWorkerPool(maxWorkers int, videoSvc *application.VideoUpdateService) *WorkerPool {
	pool := &WorkerPool{
		MaxWorkers: maxWorkers,
		JobQueue:   make(chan JobDescription, 100),
		JobAction:  videoSvc,
	}
	return pool
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.MaxWorkers; i++ {
		index := i
		go func(workerID int) {
			for job := range wp.JobQueue {
				log.Info().Str("video_id", job.videoURL).Str("video_type", job.videoType).Msg("Job started")
				convertedUrl, err := wp.JobAction.UpdateInformation(job.videoURL, job.videoType)
				if err != nil {
					log.Error().Err(err).Str("video_id", job.videoURL).Msg("Job failed")
				} else {
					log.Info().Str("video_id", job.videoURL).Str("original_url", job.videoURL).Str("converted_url", convertedUrl).Msg("Job ended")
				}
			}
		}(index)
	}
}

func (wp *WorkerPool) SubmitJob(videoURL string, videoType string) {
	wp.JobQueue <- JobDescription{videoURL: videoURL, videoType: videoType}
}

func (wp *WorkerPool) Stop() {
	close(wp.JobQueue)
}
