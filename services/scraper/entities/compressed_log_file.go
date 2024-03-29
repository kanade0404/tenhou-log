package entities

import (
	"encoding/json"
	"log"
	"strings"
)

type CompressedLogFile struct {
	File string `json:"file"`
	Size uint   `json:"size"`
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
