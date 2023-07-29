package usecases

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/kanade0404/tenhou-log/pkg/driver/storage"
	"github.com/kanade0404/tenhou-log/pkg/driver/tracer"
	"github.com/kanade0404/tenhou-log/pkg/http_handler"
	"github.com/kanade0404/tenhou-log/pkg/logger"
	"github.com/kanade0404/tenhou-log/services/scraper/entities"
)

var storageTracer = tracer.NewTracer("services/scraper/usecases/storage")

const baseUrl = "https://tenhou.net/sc/raw/dat"

type FailureFileName struct {
	Name  string
	Error error
}

// StoreCompressedLogFiles stores compressed log files to GCS
/**
  - @param c: context
  - @param bucketName: GCS bucket name
  - @param existLogs: already exist logs
  - @param notExistLogs: not exist logs
  - @return successLogs: success logs
  - @return err: error
*/
func StoreCompressedLogFiles(c context.Context, bucketName string, existLogs, notExistLogs []*entities.CompressedLogFile) ([]string, []FailureFileName, error) {
	logger.Info("StoreCompressedLogFiles start.")
	defer logger.Info("StoreCompressedLogFiles end.")
	type result struct {
		Response string
		Error    error
	}
	ctx, span := storageTracer.Start(c, "StoreCompressedLogFiles")
	defer span.End()
	ctx, cancel := context.WithTimeout(ctx, time.Minute*10)
	defer cancel()
	client, storageClose, err := storage.NewStorage(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed establish gcs client connection: %v", err)
	}
	defer func() {
		err := storageClose()
		if err != nil {
			logger.Warn(err.Error())
		}
	}()
	targetLogs := collect2UploadLogs(ctx, existLogs, notExistLogs)
	targetLogs = append(targetLogs, notExistLogs...)
	checkStatus := func(files []*entities.CompressedLogFile) <-chan result {
		resultChan := make(chan result, 10)
		go func(files []*entities.CompressedLogFile) {
			defer close(resultChan)
			for _, file := range files {
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				if err := storeGzipLog(ctx, client, bucketName, file.File); err != nil {
					resultChan <- result{
						Error: fmt.Errorf("failed buffer copy: %v", err),
					}
					continue
				} else {
					resultChan <- result{Response: fmt.Sprintf("%s/%s", bucketName, file.File)}
				}
			}
		}(notExistLogs)
		return resultChan
	}
	var (
		uploadedFileNames []string
		failureFileNames  []FailureFileName
	)
	for status := range checkStatus(notExistLogs) {
		if status.Error != nil {
			failureFileNames = append(failureFileNames, struct {
				Name  string
				Error error
			}{Name: status.Response, Error: status.Error})
		} else {
			uploadedFileNames = append(uploadedFileNames, status.Response)
		}
	}
	return uploadedFileNames, failureFileNames, nil
}

// storeGzipLog stores gzip log file to GCS.
/**
 * @param ctx context
 * @param client GCS client
 * @param bucketName GCS bucket name
 * @param fileName file name
 * @return error
 */
func storeGzipLog(ctx context.Context, client *storage.Storage, bucketName, fileName string) error {
	logger.Info("storeGzipLog start.")
	defer logger.Info("storeGzipLog end.")
	_, span := storageTracer.Start(ctx, "storeGzipLog")
	defer span.End()
	url := fmt.Sprintf("%s/%s", baseUrl, fileName)
	g, err := http_handler.RequestGZip(url, "GET", nil)
	if err != nil {
		return fmt.Errorf("failed request gzip file. file: %s: %v", url, err)
	}
	if err := client.UploadObject(bucketName, fileName, g, &storage.Option{ContentEncoding: "gzip"}); err != nil {
		return fmt.Errorf("failed buffer copy: %v", err)
	}
	return nil
}

type FetchUploadedLogsResponse struct {
	ExistLogs    []*entities.CompressedLogFile
	NotExistLogs []*entities.CompressedLogFile
	FailureLogs  []*storage.FailureLog
	Error        error
}

// FetchUploadedLogs fetches uploaded logs from GCS.
/**
 * @param bucketName GCS bucket name
 * @param compressedLogFiles compressed log files
 * @return existLogs, notExistLogs, failureLogs, error
 */
func FetchUploadedLogs(c context.Context, bucketName string, compressedLogFiles []*CompressedLog) (*FetchUploadedLogsResponse, error) {
	logger.Info("FetchUploadedLogs start fetch uploaded logs.")
	defer logger.Info("FetchUploadedLogs end fetch uploaded logs.")
	ctx, span := storageTracer.Start(c, "FetchUploadedLogs")
	defer span.End()
	ctx, cancel := context.WithTimeout(ctx, time.Minute*10)
	defer cancel()
	client, storageClose, err := storage.NewStorage(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed establish gcs client connection: %v", err)
	}
	defer func() {
		err := storageClose()
		if err != nil {
			logger.Warn(err.Error())
		}
	}()
	targetLogFileNames := make([]string, len(compressedLogFiles))
	for i := range compressedLogFiles {
		targetLogFileNames = append(targetLogFileNames, compressedLogFiles[i].Name)
	}
	var (
		existLogs    []*entities.CompressedLogFile
		notExistLogs []*entities.CompressedLogFile
		failureLogs  []*storage.FailureLog
	)
	if existLogs, notExistLogs, failureLogs, err = client.ListObjects(bucketName, targetLogFileNames); err != nil {
		logger.ErrorF("failed list objects.bucketName: %s, targetLogFileNames: %v, err: %v", bucketName, targetLogFileNames, err)
		return nil, err
	} else {
		return &FetchUploadedLogsResponse{
			ExistLogs:    existLogs,
			NotExistLogs: notExistLogs,
			FailureLogs:  failureLogs,
		}, nil
	}
}

/**
 * collect2UploadLogs collects logs that need to be uploaded.
 * @param fetchLogs logs that fetched from GCS
 * @param uploadedLogs logs that uploaded to GCS
 * @return logs that need to be uploaded
 */
func collect2UploadLogs(ctx context.Context, fetchLogs, uploadedLogs []*entities.CompressedLogFile) []*entities.CompressedLogFile {
	logger.Info("collect2UploadLogs start.")
	defer logger.Info("collect2UploadLogs end.")
	ctx, span := storageTracer.Start(ctx, "collect2UploadLogs")
	defer span.End()
	var logs []*entities.CompressedLogFile
	for _, fetchLog := range fetchLogs {
		for _, uploadedLog := range uploadedLogs {
			if fetchLog.File == uploadedLog.File && fetchLog.Size != uploadedLog.Size {
				logger.InfoF("log file size is different. file: %s, fetchLogSize: %d, uploadedLogSize: %d", fetchLog.File, fetchLog.Size, uploadedLog.Size)
				logs = append(logs, fetchLog)
				break
			}
		}
	}
	return logs
}
