package xml

import (
	"fmt"
	"github.com/kanade0404/tenhou-log/xml/init"
	"strconv"
)

type Table int

const (
	Ippan    Table = 0x00
	Jyoukyuu Table = 0x80
	Tokujou  Table = 0x20
	Houou    Table = 0xA0
)

func (t Table) Name() string {
	switch t {
	case Ippan:
		return "一般"
	case Jyoukyuu:
		return "上級"
	case Tokujou:
		return "特上"
	case Houou:
		return "鳳凰"
	default:
		return "UNKNOWN"
	}
}

type GameInfo struct {
	Demo          bool
	Red           bool
	KuitanAtozuke bool
	East          bool
	Three         bool
	HighSpeed     bool
	Table
	Lobby string
}

// Marshal バイト列を構造体に格納する
func (g *GameInfo) Marshal() (XmlElement, error) {
	var result int
	if !g.Demo {
		result = result | 1
	}
	if !g.Red {
		result = result | 1<<1
	}
	if !g.KuitanAtozuke {
		result = result | 1<<2
	}
	if !g.East {
		result = result | 1<<3
	}
	if g.Three {
		result = result | 1<<4
	}
	if g.HighSpeed {
		result = result | 1<<6
	}
	result = result | int(g.Table)
	return newXmlElement(
		"GO",
		NewXmlAttrs(map[string]string{
			"type":  strconv.Itoa(result),
			"lobby": g.Lobby,
		}),
	), nil
}

// Unmarshal 構造体を元のタグの状態に戻す
func (g *GameInfo) Unmarshal(x XmlElement) error {
	return x.ForEachAttr(func(name, value string) error {
		switch name {
		case "lobby":
			g.Lobby = value
		case "type":
			gameType, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			g.Demo = (gameType & 0x01) == 0
			g.Red = (gameType&0x02)>>1 == 0
			g.KuitanAtozuke = (gameType&0x04)>>2 == 0
			g.East = (gameType&0x08)>>3 == 0
			g.Three = (gameType&0x10)>>4 == 1
			g.HighSpeed = (gameType&0x20)>>5 == 1
			g.Table = Table(gameType & 0xA0)
		default:
			return fmt.Errorf("unknown attr. name:%s, value:%s", name, value)
		}
		return nil
	})
}
