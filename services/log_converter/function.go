package log_converter

import (
	"encoding/json"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/kanade0404/tenhou-log/pkg/http_handler"
	"net/http"
)

type request struct {
	Name string `json:"name"`
}

func init() {
	functions.HTTP("Convert", Convert)
}
func Convert(w http.ResponseWriter, r *http.Request) {
	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http_handler.ErrorResponse(w, http.StatusBadRequest, err.Error())
	}
	http_handler.Response(w, http.StatusOK, req)
}
