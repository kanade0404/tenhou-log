package storage

import (
	"bytes"
	"context"
	"io"
	"sync"

	"github.com/kanade0404/tenhou-log/services/scraper/entities"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type Storage struct {
	client *storage.Client
	ctx    context.Context
}

type Option struct {
	ContentEncoding string
}

func NewStorage(ctx context.Context) (*Storage, func() error, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, nil, err
	}
	return &Storage{
			client: client,
			ctx:    ctx,
		}, func() error {
			if err := client.Close(); err != nil {
				return err
			}
			return nil
		}, nil
}

func (s *Storage) UploadObject(bucketName string, objectName string, b []byte, option *Option) error {
	writer := s.client.Bucket(bucketName).Object(objectName).NewWriter(s.ctx)
	if option != nil {
		if option.ContentEncoding != "" {
			writer.ContentEncoding = option.ContentEncoding
		}
	}
	if _, err := io.Copy(writer, bytes.NewBuffer(b)); err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetObject(bucketName string, objectName string) ([]byte, error) {
	reader, err := s.client.Bucket(bucketName).Object(objectName).NewReader(s.ctx)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	if err := reader.Close(); err != nil {
		return nil, err
	}
	return b, nil
}

type FailureListFile struct {
}
type FailureLog struct {
	Log   *entities.CompressedLogFile
	Error error
}

/*
ListObjects
bucketNameにあるfileNamesにmatchするファイルを探し、
existLogsには存在するファイル、notExistLogsには存在しないファイル、failureLogsにはエラー終了したファイルを格納する。
*/
func (s *Storage) ListObjects(bucketName string, fileNames []string) (existLogs []*entities.CompressedLogFile, notExistLogs []*entities.CompressedLogFile, failureLogs []*FailureLog, err error) {
	// 並列処理でfileNameにmatchするファイルはexistLogsに、matchしないファイルはnotExistLogsに、エラー終了したファイルはfailureLogsに格納する
	// 並列処理の数はfileNamesの数とする
	// 並列処理の終了を待つために、fileNamesの数分のchannelを用意する
	type result struct {
		file struct {
			name string
			size int64
		}
		found bool
		err   error
	}
	wg := sync.WaitGroup{}
	ch := make(chan result, len(fileNames))
	for _, fileName := range fileNames {
		wg.Add(1)
		go func(fileName string) {
			defer wg.Done()
			query := &storage.Query{
				Prefix: fileName,
			}
			it := s.client.Bucket(bucketName).Objects(s.ctx, query)
			for {
				attrs, err := it.Next()
				if err == iterator.Done {
					ch <- result{
						file: struct {
							name string
							size int64
						}{name: fileName},
						found: false,
					}
					return
				}
				if err != nil {
					ch <- result{
						file: struct {
							name string
							size int64
						}{
							name: fileName,
						},
						err: err,
					}
					return
				}
				ch <- result{
					file: struct {
						name string
						size int64
					}{name: fileName, size: attrs.Size},
					found: true,
				}
				return
			}
		}(fileName)
	}
	wg.Wait()
	close(ch)
	for r := range ch {
		if r.err != nil {
			failureLogs = append(failureLogs, &FailureLog{
				Log: &entities.CompressedLogFile{
					File: r.file.name,
				},
				Error: err,
			})
		} else if r.found {
			existLogs = append(existLogs, &entities.CompressedLogFile{
				File: r.file.name,
				Size: uint(r.file.size),
			})
		} else if !r.found {
			notExistLogs = append(notExistLogs, &entities.CompressedLogFile{
				File: r.file.name,
			})
		}
	}
	return existLogs, notExistLogs, failureLogs, nil

}
