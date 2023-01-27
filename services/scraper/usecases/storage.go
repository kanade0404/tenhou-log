package usecases

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/kanade0404/tenhou-log/pkg/http_handler"
	"github.com/kanade0404/tenhou-log/services/scraper/entities"
	"io"
	"log"
	"math/rand"
	"time"
)

type Result struct {
	Response string
	Error    error
}

func StoreCompressedLogFiles(c context.Context, bucketName string, compressedLogFile []*entities.CompressedLogFile) ([]string, error) {
	ctx, cancel := context.WithTimeout(c, time.Second*180)
	defer cancel()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed establish gcs client connection: %v", err)
	}
	defer client.Close()
	var newLogFileBucketNames []string
	checkStatus := func(files []*entities.CompressedLogFile) <-chan Result {
		resultChan := make(chan Result, 10)
		go func(files []*entities.CompressedLogFile) {
			defer close(resultChan)
			for _, file := range files {
				log.Printf("start WriteBucket. file:%s", file.File)
				f := fmt.Sprintf("https://tenhou.net/sc/raw/dat/%s", file.File)
				g, err := http_handler.RequestGZip(f, "GET", nil)
				if err != nil {
					resultChan <- Result{
						Error: fmt.Errorf("failed request gzip file. file: %s: %v", f, err),
					}
					log.Printf("end WriteBucket. file:%s", file.File)
					continue
				}
				log.Printf("get file. file:%s", file.File)
				wc := client.Bucket(bucketName).Object(file.File).NewWriter(ctx)
				buf := bytes.NewBuffer(g)
				if _, err := io.Copy(wc, buf); err != nil {
					resultChan <- Result{
						Error: fmt.Errorf("failed buffer copy: %v", err),
					}
					log.Printf("end WriteBucket. file:%s", file.File)
					continue
				}
				log.Printf("copy file. file:%s", file.File)
				if err := wc.Close(); err != nil {
					resultChan <- Result{
						Error: fmt.Errorf("failed writer close: %v", err),
					}
					log.Printf("end WriteBucket. file:%s", file.File)
					continue
				}
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				log.Printf("close writer. file:%s", file.File)
				resultChan <- Result{Response: fmt.Sprintf("%s/%s", bucketName, file.File)}
				log.Printf("end WriteBucket. file:%s", file.File)
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
