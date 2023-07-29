package log_downloader

import (
	"encoding/json"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/kanade0404/tenhou-log/pkg/http_handler"
)

type request struct {
	Name string `json:"name"`
}

func init() {
	functions.HTTP("Download", Download)
}

func Download(w http.ResponseWriter, r *http.Request) {
	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http_handler.ErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
	http_handler.Response(w, http.StatusOK, req)
}
