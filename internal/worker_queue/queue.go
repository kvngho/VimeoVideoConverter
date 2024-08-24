package worker_queue

import (
	"github.com/kvngho/vimeovideoconverter/internal/application"
	"github.com/rs/zerolog/log"
	"time"
)

type JobDescription struct {
	videoURL  string
	videoType string
}

type JobFailDescription struct {
	JobDescription
	FailedCount int
}
type WorkerPool struct {
	MaxWorkers     int
	JobQueue       chan JobDescription
	FailedJobQueue chan JobFailDescription
	JobAction      *application.VideoUpdateService
}

func NewWorkerPool(maxWorkers int, videoSvc *application.VideoUpdateService) *WorkerPool {
	pool := &WorkerPool{
		MaxWorkers:     maxWorkers,
		JobQueue:       make(chan JobDescription, 100),
		FailedJobQueue: make(chan JobFailDescription, 100),
		JobAction:      videoSvc,
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
					log.Info().Err(err).Str("video_id", job.videoURL).Msg("Job failed - Submit the job to failed queue")
					wp.ReSubmitJob(job, 1)
				} else {
					log.Info().Str("video_id", job.videoURL).Str("original_url", job.videoURL).Str("converted_url", convertedUrl).Msg("Job ended")
				}
			}
		}(index)
	}
	// Handler Failed Job
	go func() {
		log.Info().Msg("Start handling failed job")
		for failedJob := range wp.FailedJobQueue {
			time.Sleep(30 * time.Second) // 모든 작업에 대해서 5초 대기, 추후 수정 필요
			convertedUrl, err := wp.JobAction.UpdateInformation(failedJob.videoURL, failedJob.videoType)
			if err != nil {
				log.Info().Str("video_id", failedJob.videoURL).Int("retry", failedJob.FailedCount).Msg("Job failed - Requeue")
				if failedJob.FailedCount < 3 {
					wp.ReSubmitJob(failedJob.JobDescription, failedJob.FailedCount+1)
					continue
				}
				log.Error().Str("video_url", failedJob.videoURL).Msg("Error. This video url is not yet valid")
			} else {
				log.Info().Str("video_id", failedJob.videoURL).Str("original_url", failedJob.videoURL).Str("converted_url", convertedUrl).Msg("Job ended")
			}
		}
	}()
}

func (wp *WorkerPool) SubmitJob(videoURL string, videoType string) {
	wp.JobQueue <- JobDescription{videoURL: videoURL, videoType: videoType}
}

func (wp *WorkerPool) ReSubmitJob(jobDesc JobDescription, failedCount int) {
	wp.FailedJobQueue <- JobFailDescription{jobDesc, failedCount}
}

func (wp *WorkerPool) Stop() {
	close(wp.JobQueue)
}
