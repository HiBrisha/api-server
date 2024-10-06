package main

import (
	"module/internal/handlers"
	api "module/internal/services/api"
	logger "module/pkg/utils/logs"
	"net/http"
)

func main() {
	log := logger.SysLog{
		StrDir: "../api-server/docs/logs/",
	}
	log.Init()
	//Test API Server
	server := &api.APIServer{
		Port: ":8080",
		Log:  &log,
	}

	handlers := handlers.Info{
		Log: &log,
	}

	// Thiết lập các route
	server.Route(http.MethodPost, "/login/signup", handlers.InsertUser)

	server.Route(http.MethodGet, "/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Khởi động server
	if err := server.Init(); err != nil {
		log.Write("ERR", "Failed to start server: %s", err)
	}

}
