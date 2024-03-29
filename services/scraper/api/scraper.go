package api

import (
	"context"
	"encoding/json"
	"github.com/kanade0404/tenhou-log/pkg/http_handler"
	"github.com/kanade0404/tenhou-log/services/ent"
	"github.com/kanade0404/tenhou-log/services/scraper/usecases"
	"net/http"
	"time"
)

type request struct {
	Count  int       `json:"count"`
	Before time.Time `json:"before"`
	After  time.Time `json:"after"`
}

func (a *Api) Scraper(w http.ResponseWriter, r *http.Request) {
	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http_handler.ErrorResponse(w, http.StatusBadRequest, err.Error())
	}
	res, err := ScrapingAndStore(a.Context, a.client, a.config.CompressedLogBucketName(), req)
	if err != nil {
		http_handler.ErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
	http_handler.Response(w, http.StatusOK, res)
}

func ScrapingAndStore(ctx context.Context, db *ent.Client, bucketName string, req request) ([]string, error) {
	logFiles, err := usecases.ScrapingCompressedLog(req.Count)
	if err != nil {
		return nil, err
	}
	compressedLogFile, err := usecases.StoreCompressedLogFile(ctx, db, logFiles)
	if err != nil {
		return nil, err
	}
	return usecases.StoreCompressedLogFiles(ctx, bucketName, compressedLogFile)
}
