package init

import "github.com/kanade0404/tenhou-log/xml/hai"

type InitUser struct {
	Name  string
	Point uint
	Hais  []hai.Hai
}

type Init struct {
	Seed
	LeaderUserName string
	InitUsers      []InitUser
}
