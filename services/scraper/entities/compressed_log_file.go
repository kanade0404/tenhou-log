package entities

import (
	"encoding/json"
	"log"
	"scraper/models"
	"strings"
)

type CompressedLogFile struct {
	File string `json:"file"`
	Size int    `json:"size"`
}

func Unmarshal(logText string) (*CompressedLogFile, error) {
	var l CompressedLogFile
	t := strings.ReplaceAll(strings.ReplaceAll(strings.Replace(logText, "},", "}", 1), "\n", ""), "'", `"`)
	log.Println(t)
	err := json.Unmarshal([]byte(t), &l)
	if err != nil {
		return nil, err
	}
	return &l, nil
}
func (c *CompressedLogFile) Transformer() *models.CompressedLogFile {
	return &models.CompressedLogFile{
		Name: c.File,
		Size: c.Size,
	}
}
