package api

import (
	logger "module/pkg/utils/logs"
	"net/http"
)

// APIServer struct chứa thông tin cấu hình server như cổng kết nối.
type APIServer struct {
	Port string
	Log  *logger.SysLog
}

// Init khởi tạo server
func (server *APIServer) Init() error {
	server.Log.Write("INFO", "Starting server on port: %s", server.Port)

	// Bắt đầu server
	if err := http.ListenAndServe(server.Port, nil); err != nil {
		server.Log.Write("ERR", "Error starting server: %s", err)
		return err
	}
	return nil
}

// Route thêm route cho server
func (apiServer *APIServer) Route(method, path string, handlerFunc http.HandlerFunc) {
	switch method {
	case http.MethodGet:
		http.HandleFunc(path, handlerFunc)
	case http.MethodPost:
		http.HandleFunc(path, handlerFunc)
	default:
		apiServer.Log.Write("WAR", "Unsupported method: %s", method)
	}
}
