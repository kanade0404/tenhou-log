package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/kanade0404/tenhou-log/pkg/config"
	_ "github.com/lib/pq"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/*
atlasのversioned migrationで作成されたmigration fileを実行する。
*/
func main() {
	if len(os.Args) != 2 {
		log.Fatalln("migration file name prefix is required.")
	}
	ctx := context.Background()
	f, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	files, err := ioutil.ReadDir(fmt.Sprintf("%s/%s/%s", filepath.Dir(f), filepath.Base(f), "services/database/migrations"))
	if err != nil {
		log.Fatalln(err)
	}
	fileName := findMatchedNameSQLFile(files, os.Args[1])
	if fileName == "" {
		log.Fatalf("matched SQL file is not exist. name: %s", os.Args[1])
	}
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/%s/%s", filepath.Dir(f), filepath.Base(f), "services/database/migrations", fileName))
	if err != nil {
		log.Fatalln(err)
	}
	env, err := config.NewEnv(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	db, err := sql.Open(env.Dialect(), env.ConnectionString())
	if err != nil {
		log.Fatalln(err)
	}
	b := bytes.NewBuffer(content)
	if _, err := db.Exec(b.String()); err != nil {
		log.Fatalln(err)
	}
}
func findMatchedNameSQLFile(files []fs.FileInfo, name string) string {
	for i := range files {
		if strings.Index(files[i].Name(), fmt.Sprintf("_%s.sql", os.Args[1])) != -1 {
			return files[i].Name()
		}
	}
	return ""
}
