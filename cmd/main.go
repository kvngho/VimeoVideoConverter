package main

import (
	pb "github.com/kvngho/vimeovideoconverter/api/v1"
	"github.com/kvngho/vimeovideoconverter/internal/application"
	config2 "github.com/kvngho/vimeovideoconverter/internal/config"
	"github.com/kvngho/vimeovideoconverter/internal/infrastructure/api/vimeo"
	gprc2 "github.com/kvngho/vimeovideoconverter/internal/infrastructure/grpc"
	"github.com/kvngho/vimeovideoconverter/internal/infrastructure/persistence/ent"
	"github.com/kvngho/vimeovideoconverter/internal/worker_queue"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	cfg, err := config2.LoadConfig()
	log.Debug().Interface("config", cfg).Msg("Load config")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config file")
		return
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen")
	}

	grpcServer := grpc.NewServer()

	vimeoClient := vimeo.NewVimeoClient(cfg.VimeoToken)
	entClient := ent.NewEntClient(ent.Config{
		Name:     cfg.DatabaseName,
		Username: cfg.DatabaseUser,
		Password: cfg.DatabasePassword,
		Port:     cfg.DatabasePort,
		URL:      cfg.DatabaseURL,
	})
	defer entClient.Close()
	videoUpdateService := application.NewVideoUpdateService(entClient, vimeoClient)
	queue := worker_queue.NewWorkerPool(cfg.WorkerNumber, videoUpdateService)
	queue.Start()
	defer queue.Stop()
	pb.RegisterConverterServer(grpcServer, gprc2.NewVideoServer(queue))
	log.Info().Msg("Starting grpc server...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("Failed to serve")
	}
}
