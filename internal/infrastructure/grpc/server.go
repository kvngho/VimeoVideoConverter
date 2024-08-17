package grpc

import (
	"context"
	pb "github.com/kvngho/vimeovideoconverter/api/v1"
	"github.com/kvngho/vimeovideoconverter/internal/worker_queue"
	"github.com/rs/zerolog/log"
)

type VideoServer struct {
	pb.UnimplementedConverterServer
	workerPool *worker_queue.WorkerPool
}

func NewVideoServer(workerPool *worker_queue.WorkerPool) *VideoServer {
	return &VideoServer{workerPool: workerPool}
}

func (vs *VideoServer) ConvertVideo(ctx context.Context, req *pb.VideoRequest) (*pb.VideoResponse, error) {
	log.Info().Str("video_url", req.Url).Msg("Sending Event to Worker queue")
	vs.workerPool.SubmitJob(req.Url, req.Type)
	return &pb.VideoResponse{Success: true}, nil
}
