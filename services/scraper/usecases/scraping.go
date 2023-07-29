package usecases

import (
	"context"
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"

	"github.com/kanade0404/tenhou-log/pkg/driver/tracer"
	"github.com/kanade0404/tenhou-log/pkg/http_handler"
	"github.com/kanade0404/tenhou-log/pkg/logger"
	"github.com/kanade0404/tenhou-log/services/scraper/entities"
	"github.com/kanade0404/tenhou-log/services/scraper/internal"
)

var scrapingTracer = tracer.NewTracer("services/scraper/usecases/scraping")

type CompressedLog struct {
	Name string
	Body []byte
	Size uint
}
type ScrapingCompressedLogListResponse struct {
	SuccessLogs []*CompressedLog
	FailureLogs []FailureFileName
}

// ScrapingCompressedLogList は圧縮されたログのリストを取得する
/*
  - @param c: context
  - @param count: nilの場合は全てのログを取得する
  - @return ScrapingCompressedLogListResponse: 成功したログと失敗したログのリスト
  - @return err: エラー
*/
func ScrapingCompressedLogList(c context.Context, count *int) (*ScrapingCompressedLogListResponse, error) {
	logger.Info("ScrapingCompressedLogList start")
	defer logger.Info("ScrapingCompressedLogList end")
	ctx, span := scrapingTracer.Start(c, "ScrapingCompressedLogList")
	defer span.End()
	body, err := http_handler.RequestHTTP("https://tenhou.net/sc/raw/list.cgi")
	if err != nil {
		return nil, fmt.Errorf("ScrapingCompressedLogList failed. %w", err)
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			go logger.Error(err.Error())
		}
	}(body)
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, fmt.Errorf("ScrapingCompressedLogList failed. %w", err)
	}
	logFiles := internal.UnmarshalFileListText(ctx, doc.Text())
	type result struct {
		Body  []byte
		Log   *entities.CompressedLogFile
		Error error
	}
	checkFetchStatus := func(logFiles []string) <-chan result {
		resultChan := make(chan result)
		go func(files []string) {
			logger.InfoF("ScrapingCompressedLogList start fetch. total: %d", len(logFiles))
			for i := range files {
				if count != nil && i == *count {
					break
				}
				logger.InfoF("ScrapingCompressedLogList start fetch. %d/%d", i+1, len(files))
				l, err := entities.Unmarshal(ctx, files[i])
				if err != nil {
					err = fmt.Errorf("ScrapingCompressedLogList error unmarshal entity: %v", err)
					resultChan <- result{
						Log:   l,
						Error: err,
					}
					logger.ErrorF("ScrapingCompressedLogList end fetch. %d/%d\n%s", i+1, len(files), err.Error())
					continue
				}
				url := fmt.Sprintf("https://tenhou.net/sc/raw/dat/%s", l.File)
				g, err := http_handler.RequestGZip(url, "GET", nil)
				logger.InfoF("ScrapingCompressedLogList fetch target: %s. %d/%d", url, i+1, len(files))
				if err != nil {
					err = fmt.Errorf("ScrapingCompressedLogList error get gzip logs. URL: %s", url)
					resultChan <- result{
						Log:   l,
						Error: err,
					}
					logger.ErrorF("ScrapingCompressedLogList end fetch. %d/%d\n%s", i+1, len(files), err.Error())
					continue
				}
				resultChan <- result{
					Body: g,
					Log:  l,
				}
				logger.InfoF("ScrapingCompressedLogList end fetch. %d/%d", i+1, len(files))
			}
			logger.Info("ScrapingCompressedLogList end fetch")
		}(logFiles)
		return resultChan
	}
	var (
		successLogs []*CompressedLog
		failureLogs []FailureFileName
	)
	for status := range checkFetchStatus(logFiles) {
		if status.Error != nil {
			failureLogs = append(failureLogs, FailureFileName{
				Name:  status.Log.File,
				Error: status.Error,
			})
		} else {
			successLogs = append(successLogs, &CompressedLog{
				Name: status.Log.File,
				Body: status.Body,
				Size: status.Log.Size,
			})
		}
	}
	return &ScrapingCompressedLogListResponse{
		SuccessLogs: successLogs,
		FailureLogs: failureLogs,
	}, nil
}
