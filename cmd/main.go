package main

import (
	"module/internal/handlers"
	api "module/internal/services/api"
	logger "module/pkg/utils/logs"
	"net/http"
)

var log logger.SysLog = logger.SysLog{
	StrDir: "../docs/logs/",
}

func main() {
	log.Init()
	//Test API Server
	apiServer := &api.APIServer{
		Port: "3000",
		Log:  &log,
	}

	handler := handlers.Info{
		Log: &log,
	}
	apiServer.Route(http.MethodPost, "/login/signup", handler.InsertUser)

	// Initialize the server
	err := apiServer.Init()
	if err != nil {
		log.Write("ERR", "Unable to initialize server: %v\n", err)
		return
	}

	// Start the server
	if err := apiServer.Start(); err != nil {
		log.Write("ERR", "Unable to start server: %v\n", err)
	}

}
