package usecases

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"external/http_handler"
	"fmt"
	"io"
	"scraper/models"
	"time"
)

func StoreCompressedLogFiles(c context.Context, bucketName string, compressedLogFile models.CompressedLogFileSlice) ([]string, error) {
	ctx, cancel := context.WithTimeout(c, time.Second*50)
	defer cancel()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	var newLogFileBucketNames []string
	for _, l := range compressedLogFile {
		g, err := http_handler.RequestGZip(fmt.Sprintf("https://tenhou.net/sc/raw/dat/%s", l.Name), "GET", nil)
		if err != nil {
			return nil, err
		}
		wc := client.Bucket(bucketName).Object(l.Name).NewWriter(ctx)
		wc.ChunkSize = 0
		buf := bytes.NewBuffer(g)
		if _, err := io.Copy(wc, buf); err != nil {
			return nil, err
		}
		if err := wc.Close(); err != nil {
			return nil, err
		}
		newLogFileBucketNames = append(newLogFileBucketNames, fmt.Sprintf("%s/%s", bucketName, l.Name))
	}
	return newLogFileBucketNames, nil
}
