package usecases

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/pkg/http_handler"
	"github.com/kanade0404/tenhou-log/services/ent"
	"github.com/kanade0404/tenhou-log/services/ent/compressedmjlog"
	"github.com/kanade0404/tenhou-log/services/scraper/entities"
	"github.com/kanade0404/tenhou-log/services/scraper/internal"
	"io"
	"log"
)

func ScrapingCompressedLog(count int) ([]*entities.CompressedLogFile, error) {
	body, err := http_handler.RequestHTTP("https://tenhou.net/sc/raw/list.cgi")
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(body)
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	logFiles := internal.UnmarshalFileListText(doc.Text())
	var fetchLogs []*entities.CompressedLogFile
	for i, str := range logFiles {
		if count != 0 && i == count {
			break
		}
		l, err := entities.Unmarshal(str)
		if err != nil {
			log.Println("error unmarshal entity: ", err)
			return nil, err
		}
		fetchLogs = append(fetchLogs, l)
	}
	return fetchLogs, nil
}
func FetchUpdatedLogFiles(ctx context.Context, db *ent.Client, logs []*entities.CompressedLogFile) ([]*ent.CompressedMJLog, error) {
	var logNames []string
	for i := range logs {
		logNames = append(logNames, logs[i].File)
	}

	currentLogFiles, err := db.CompressedMJLog.Query().Where(func(selector *sql.Selector) {
		selector.Where(sql.In(compressedmjlog.FieldName, logNames))
	}).All(ctx)
	if err != nil {
		return nil, err
	}
	resultLogMap := make(map[uuid.UUID]*ent.CompressedMJLog)
	for i := range currentLogFiles {
		for li := range logs {
			if currentLogFiles[i].Name == logs[li].File {
				if currentLogFiles[i].Size != logs[li].Size {
					resultLogMap[currentLogFiles[i].ID] = currentLogFiles[i]
				}
				break
			}
		}
	}
	var resultLogs []*ent.CompressedMJLog
	for k := range resultLogMap {
		resultLogs = append(resultLogs, resultLogMap[k])
	}
	return resultLogs, err
}
func StoreCompressedLogFile(ctx context.Context, db *ent.Client, fetchLogFiles []*entities.CompressedLogFile) ([]*entities.CompressedLogFile, error) {
	var (
		targetFiles []*entities.CompressedLogFile
		queries     []*ent.CompressedMJLogCreate
	)
	// DBから圧縮されたファイルレコード取ってくる
	var latestFileMap = make(map[string]*ent.CompressedMJLog)
	currentLogFiles, err := db.CompressedMJLog.Query().All(ctx)
	for _, file := range currentLogFiles {
		if _, ok := latestFileMap[file.Name]; ok {
			if file.InsertedAt.After(latestFileMap[file.Name].InsertedAt) {
				latestFileMap[file.Name] = file
			}
		} else {
			latestFileMap[file.Name] = file
		}
	}
	if err != nil {
		return nil, err
	}
	for _, fetchLog := range fetchLogFiles {
		for _, current := range latestFileMap {
			// ファイル名が一致してサイズが違うなら
			if fetchLog.File == current.Name {
				if fetchLog.Size != current.Size {
					targetFiles = append(targetFiles, fetchLog)
					queries = append(queries, db.CompressedMJLog.Create().SetSize(fetchLog.Size).SetName(fetchLog.File))
				}
				break
			}
		}
		// いずれにもヒットしなかった = 新規
		targetFiles = append(targetFiles, fetchLog)
		queries = append(queries, db.CompressedMJLog.Create().SetSize(fetchLog.Size).SetName(fetchLog.File))
	}
	if err := db.CompressedMJLog.CreateBulk(queries...).Exec(ctx); err != nil {
		return nil, err
	}
	return targetFiles, nil
}
