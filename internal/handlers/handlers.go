package handlers

import (
	"encoding/json"
	"io/ioutil"
	logger "module/pkg/utils/logs"
	"net/http"
)

type Info struct {
	Log *logger.SysLog
}
type UserData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (info *Info) InsertUser(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	// Ensure the body is closed after reading

	var data UserData
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	info.Log.Write("INFO", "Received Name: %s, Password: %s\n", data.Name, data.Password)
}
