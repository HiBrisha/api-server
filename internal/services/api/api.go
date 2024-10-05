package api

import (
	logger "module/pkg/utils/logs"
	"net/http"
)

// APIServer struct chứa thông tin cấu hình server như cổng kết nối.
type APIServer struct {
	Port string
	Log  *logger.SysLog
	mux  *http.ServeMux
}

// Init initializes the API server
func (server *APIServer) Init() error {
	server.mux = http.NewServeMux()
	return nil
}

// Start starts the HTTP server
func (server *APIServer) Start() error {
	server.Log.Write("INFO", ":%s", server.Port)
	return http.ListenAndServe(":"+server.Port, server.mux)
}

func (apiServer *APIServer) Route(method, path string, handlerFunc http.HandlerFunc) {
	switch method {
	case http.MethodGet:
		apiServer.mux.HandleFunc(path, handlerFunc)
	case http.MethodPost:
		apiServer.mux.HandleFunc(path, handlerFunc)
	default:
		apiServer.Log.Write("WAR", "Unsupported method: %s\n", method)
	}
}
