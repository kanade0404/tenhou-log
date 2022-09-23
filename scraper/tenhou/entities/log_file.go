package entities

import (
	"encoding/json"
	"log"
	"strings"
)

type Log struct {
	File string `json:"file"`
	Size uint   `json:"size"`
}

func Unmarshal(logText string) (*Log, error) {
	var l Log
	t := strings.ReplaceAll(strings.ReplaceAll(strings.Replace(logText, "},", "}", 1), "\n", ""), "'", `"`)
	log.Println(t)
	err := json.Unmarshal([]byte(t), &l)
	if err != nil {
		return nil, err
	}
	return &l, nil
}
