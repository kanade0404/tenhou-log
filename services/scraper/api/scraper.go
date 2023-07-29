package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"

	"github.com/kanade0404/tenhou-log/pkg/http_handler"
	"github.com/kanade0404/tenhou-log/pkg/logger"
	"github.com/kanade0404/tenhou-log/services/scraper/usecases"
)

type request struct {
	Count  *int       `json:"count,omitempty"`
	Before *time.Time `json:"before,omitempty"`
	After  *time.Time `json:"after,omitempty"`
}

var scraperTracer = otel.GetTracerProvider().Tracer("github.com/kanade0404/tenhou-log/services/scraper/api/scraper")

// Scraper は圧縮されたログを取得し、GCSに保存する
/**
 * @param w: http.ResponseWriter
 * @param r: http.Request
 * @return err: エラー
 */
func (a *Api) Scraper(w http.ResponseWriter, r *http.Request) {
	logger.Info("Scraper start.")
	defer logger.Info("Scraper end.")
	ctx, span := scraperTracer.Start(a.Context, "Scraper")
	defer span.End()
	a.Context = ctx
	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error(err.Error())
		http_handler.ErrorResponse(w, http.StatusBadRequest, err.Error())
	}
	logger.InfoF("Request Body: {Count: %d, Before: %s, After: %s}", req.Count, req.Before, req.After)
	successes, failures, err := ScrapingAndStore(a.Context, a.config.CompressedLogBucketName(), req)
	if err != nil {
		logger.Error(err.Error())
		http_handler.ErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
	type resBody struct {
		SuccessLogs []string `json:"success_logs"`
		FailureLogs []string `json:"failure_logs"`
	}
	body := resBody{
		SuccessLogs: successes,
	}
	for i := range failures {
		body.FailureLogs = append(body.FailureLogs, failures[i].Name)
	}
	logger.InfoF("Response Body: {SuccessLogs: %v, FailureLogs: %v}", body.SuccessLogs, body.FailureLogs)
	http_handler.Response(w, http.StatusOK, body)
}

// ScrapingAndStore は圧縮されたログを取得し、GCSに保存する
/**
 * @param ctx: context.Context
 * @param bucketName: string
 * @param req: request
 * @return successLogs: []string
 * @return failureLogs: []usecases.FailureFileName
 * @return err: エラー
 */
func ScrapingAndStore(ctx context.Context, bucketName string, req request) ([]string, []usecases.FailureFileName, error) {
	logger.Info("ScrapingAndStore start.")
	defer logger.Info("ScrapingAndStore end.")
	ctx, span := scraperTracer.Start(ctx, "ScrapingAndStore")
	defer span.End()
	res, err := usecases.ScrapingCompressedLogList(ctx, req.Count)
	if err != nil {
		return nil, nil, err
	}
	fetchUploadedLogsRes, err := usecases.FetchUploadedLogs(ctx, bucketName, res.SuccessLogs)
	if err != nil {
		return nil, nil, err
	}
	success, failure, err := usecases.StoreCompressedLogFiles(ctx, bucketName, fetchUploadedLogsRes.ExistLogs, fetchUploadedLogsRes.NotExistLogs)
	if err != nil {
		return nil, nil, err
	}
	var failureLogs []usecases.FailureFileName
	failureLogs = append(failureLogs, failure...)
	for i := range fetchUploadedLogsRes.FailureLogs {
		failureLogs = append(failureLogs, usecases.FailureFileName{
			Name:  fetchUploadedLogsRes.FailureLogs[i].Log.File,
			Error: fetchUploadedLogsRes.FailureLogs[i].Error,
		})
	}
	return success, failureLogs, nil
}
