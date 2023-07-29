package main

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/xml"
	"fmt"
	xml2 "github.com/kanade0404/tenhou-log/pkg/parser/xml"
	"io"
	"log"
)

type MJLog struct {
}

func next(r *xml.Decoder) (xml2.XmlElement, error) {
	for {
		token, err := r.Token()
		if err != nil {
			return xml2.XmlElement{}, err
		}
		switch t := token.(type) {
		case xml.StartElement:
			var attr []xml2.XmlAttr
			for _, a := range t.Attr {
				attr = append(attr, xml2.XmlAttr{
					Name:  a.Name.Local,
					Value: a.Value,
				})
			}
			return xml2.XmlElement{
				Name: t.Name.Local,
				Attr: attr,
			}, nil
		}
	}
}
func main() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	reader, err := client.Bucket("tmp-tenhou-log-dev-f2f840ed754beba3").Object("scb2023031600.html.gz").ReadCompressed(true).NewReader(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer reader.Close()
	data, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))
	//records := strings.Split(string(data), "\n")
	//for i := range records {
	//	if records[i] != "" {
	//		data := strings.Split(records[i], "|")
	//		if len(data) > 0 {
	//			fmt.Printf("code: %s, time: %s, room: %s, players: %s\n", data[0], data[1], data[2], data[3])
	//		}
	//	}
	//}
}
func unmarshal(decoder *xml.Decoder) (*MJLog, error) {
	m := &MJLog{}
	var startGame bool
	for {
		token, err := next(decoder)
		if err != nil {
			if err == io.EOF {
				return m, nil
			}
			return nil, err
		}
		fmt.Println(token.Name)
		if startGame {
			// ツモや打牌の解析
			fmt.Println(token.Name)
		} else {
			switch token.Name {
			case "mjloggm":
				// mjlogの解析
				fmt.Println(token.AttrByName("ver"))
			case "SHUFFLE":
				// shuffleの解析
				fmt.Println(token.AttrByName("seed"))
			case "GO":
				// goの解析
				fmt.Println(token.AttrByName("type"))
				fmt.Println(token.AttrByName("lobby"))
			case "UN":
				// 対局開始前のUN
				fmt.Println(token.AttrByName("n0"))
				fmt.Println(token.AttrByName("n1"))
				fmt.Println(token.AttrByName("n2"))
				fmt.Println(token.AttrByName("n3"))
				fmt.Println(token.AttrByName("dan"))
				fmt.Println(token.AttrByName("rate"))
				fmt.Println(token.AttrByName("sx"))
			case "TAIKYOKU":
				// oyaは常に0なので、保持しない
				startGame = true
			default:
				return nil, fmt.Errorf("unexpected token %v \n %v", token.Name, token)
			}
		}
	}
}
