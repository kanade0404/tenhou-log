package usecases

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/kanade0404/tenhou-log/pkg/http_handler"
	"github.com/kanade0404/tenhou-log/services/ent"
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
	log.Println("text: ", doc.Text())

	logFiles := internal.UnmarshalFileListText(doc.Text())
	log.Printf("logs: \n%v", logFiles)
	var fetchLogs []*entities.CompressedLogFile
	for i, str := range logFiles {
		if count != 0 && i == count {
			break
		}
		log.Println("str: ", str)
		l, err := entities.Unmarshal(str)
		if err != nil {
			log.Println("error unmarshal entity: ", err)
			return nil, err
		}
		fetchLogs = append(fetchLogs, l)
	}
	return fetchLogs, nil
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
	if err := db.Debug().CompressedMJLog.CreateBulk(queries...).Exec(ctx); err != nil {
		return nil, err
	}
	return targetFiles, nil
}
