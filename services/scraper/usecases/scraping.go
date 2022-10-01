package usecases

import (
	"context"
	"database/sql"
	"external/http_handler"
	"github.com/PuerkitoBio/goquery"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
	"scraper/entities"
	"scraper/internal"
	"scraper/models"
)

func ScrapingCompressedLog() ([]*entities.CompressedLogFile, error) {
	body, err := http_handler.RequestHTTP("https://tenhou.net/sc/raw/list.cgi")
	defer body.Close()
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}
	log.Println("text: ", doc.Text())

	logFiles := internal.UnmarshalFileListText(doc.Text())
	log.Printf("logs: \n%v", logFiles)
	var fetchLogs []*entities.CompressedLogFile
	for _, str := range logFiles {
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
func StoreCompressedLogFile(ctx context.Context, db *sql.DB, fetchLogFiles []*entities.CompressedLogFile) (models.CompressedLogFileSlice, error) {
	// DBから圧縮されたファイルレコード取ってくる
	currentLogFiles, err := models.CompressedLogFiles().All(ctx, db)
	var newLogFiles models.CompressedLogFileSlice
	var diffLogFileNames models.CompressedLogFileSlice
	for _, fetchLog := range fetchLogFiles {
		for _, current := range currentLogFiles {
			if fetchLog.File == current.Name {
				if fetchLog.Size != current.Size {
					diffLogFileNames = append(diffLogFileNames, fetchLog.Transformer())
				}
				break
			}
		}
		// いずれにもヒットしなかった = 新規
		diffLogFileNames = append(diffLogFileNames, fetchLog.Transformer())
		newLogFiles = append(newLogFiles, fetchLog.Transformer())
	}
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	for _, file := range newLogFiles {
		if err := file.Insert(ctx, tx, boil.Infer()); err != nil {
			err := tx.Rollback()
			if err != nil {
				return nil, err
			}
			return nil, err
		}
	}
	for _, diffLogFile := range diffLogFileNames {
		if _, err := diffLogFile.Update(ctx, tx, boil.Infer()); err != nil {
			err := tx.Rollback()
			if err != nil {
				return nil, err
			}
			return nil, err
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return diffLogFileNames, nil
}
