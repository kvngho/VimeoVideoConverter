package server

import (
	"github.com/gorilla/mux"
	"github.com/kvngho/vimeovideoconverter/internal/controller"
	"github.com/kvngho/vimeovideoconverter/internal/worker_queue"
	"net/http"
)

type WebServer struct {
	router *mux.Router
	queue  *worker_queue.WorkerPool
}

func NewWebServer(queue *worker_queue.WorkerPool) *WebServer {
	return &WebServer{
		router: mux.NewRouter(),
		queue:  queue,
	}
}

func (s *WebServer) Start() error {
	s.routes()
	return http.ListenAndServe(":50052", s.router)
}

func (s *WebServer) routes() {
	userController := controller.NewVideoController(s.queue)
	s.router.HandleFunc("/video", userController.UpdateVideo).Methods("POST")
}
