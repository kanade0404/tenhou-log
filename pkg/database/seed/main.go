package main

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/kanade0404/tenhou-log/pkg/config"
	"github.com/kanade0404/tenhou-log/pkg/database"
	"github.com/kanade0404/tenhou-log/services/ent"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

var chineseNumeral = map[int]string{
	2:  "二",
	3:  "三",
	4:  "四",
	5:  "五",
	6:  "六",
	7:  "七",
	8:  "八",
	9:  "九",
	10: "十",
}

func TenhouDans() []string {
	dans := []string{"新人"}
	for i := 9; i > 0; i-- {
		dans = append(dans, fmt.Sprintf("%d級", i))
	}
	dans = append(dans, "初段")
	for i := 2; i < 11; i++ {
		dans = append(dans, fmt.Sprintf("%s段", chineseNumeral[i]))
	}
	return dans
}

func main() {
	log.Println("Run start")
	ctx := context.Background()
	env, err := config.NewEnv(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := database.CreateClient(ctx, env.Dialect(), env.ConnectionString())
	if err != nil {
		log.Fatalf("failed database connection: %v", err)
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalf("close database client: %v", err)
		}
	}(client)
	tx, err := client.Tx(ctx)
	if err != nil {
		log.Fatalf("starting a transaction: %v", err)
	}
	dans := TenhouDans()
	for _, dan := range dans {
		err := tx.Dan.Create().SetName(dan).OnConflict(sql.ConflictColumns("name")).DoNothing().Exec(ctx)
		if err != nil && !strings.Contains(err.Error(), "sql: no rows in result set") {
			err := database.Rollback(tx, err)
			if err != nil {
				log.Fatalf("rollback: %v", err)
			}
		}
	}
	if err := tx.Commit(); err != nil {
		log.Fatalf("failed commit: %v", err)
	}
	log.Println("Run end")
}
