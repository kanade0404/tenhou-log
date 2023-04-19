package usecases

import (
	"context"
	"fmt"
	"github.com/kanade0404/tenhou-log/pkg/driver/storage"
	"github.com/kanade0404/tenhou-log/pkg/http_handler"
	"github.com/kanade0404/tenhou-log/services/scraper/entities"
	"go.opentelemetry.io/otel"
	"log"
	"math/rand"
	"time"
)

type Result struct {
	Response string
	Error    error
}

var storageTracer = otel.Tracer("scraper/trace")

func StoreCompressedLogFiles(c context.Context, bucketName string, compressedLogFile []*entities.CompressedLogFile) ([]string, error) {
	ctx, span := storageTracer.Start(c, "StoreCompressedLogFiles")
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
			log.Println(err)
		}
	}()
	var newLogFileBucketNames []string
	checkStatus := func(files []*entities.CompressedLogFile) <-chan Result {
		resultChan := make(chan Result, 10)
		go func(files []*entities.CompressedLogFile) {
			defer close(resultChan)
			for _, file := range files {
				url := fmt.Sprintf("https://tenhou.net/sc/raw/dat/%s", file.File)
				g, err := http_handler.RequestGZip(url, "GET", nil)
				if err != nil {
					resultChan <- Result{
						Error: fmt.Errorf("failed request gzip file. file: %s: %v", url, err),
					}
					continue
				}
				if err := client.UploadObject(bucketName, bucketName, g, &storage.Option{ContentEncoding: "gzip"}); err != nil {
					resultChan <- Result{
						Error: fmt.Errorf("failed buffer copy: %v", err),
					}
					continue
				}
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				resultChan <- Result{Response: fmt.Sprintf("%s/%s", bucketName, file.File)}
			}
		}(compressedLogFile)
		return resultChan
	}
	for status := range checkStatus(compressedLogFile) {
		if status.Error != nil {
			log.Println(status.Error)
			continue
		} else {
			newLogFileBucketNames = append(newLogFileBucketNames, status.Response)
		}
	}
	return newLogFileBucketNames, nil
}
