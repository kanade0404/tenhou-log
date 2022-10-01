package api

import (
	"context"
	"database/sql"
	"external/http_handler"
	"net/http"
	"scraper/configs"
	"scraper/usecases"
)

type Api struct {
	context.Context
	Db *sql.DB
	configs.Env
}

func (a *Api) Scraper(w http.ResponseWriter, _ *http.Request) {
	res, err := ScrapingAndStore(a.Context, a.Db, a.Env.LogFile.CompressedLogBucketName)
	if err != nil {
		http_handler.ErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
	http_handler.Response(w, http.StatusOK, res)
}
func ScrapingAndStore(ctx context.Context, db *sql.DB, bucketName string) ([]string, error) {
	logFiles, err := usecases.ScrapingCompressedLog()
	if err != nil {
		return nil, err
	}
	compressedLogFile, err := usecases.StoreComporessedLogFile(ctx, db, logFiles)
	if err != nil {
		return nil, err
	}
	return usecases.StoreCompressedLogFiles(ctx, bucketName, compressedLogFile)
}
