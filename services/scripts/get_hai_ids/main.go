package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var num, haiType string
	fmt.Print("牌タイプを指定してください（萬子:1,筒子:2,索子:3,字牌:4）:")
	_, err := fmt.Scan(&haiType)
	if err != nil {
		log.Fatalln(err)
	}
	if haiType == "4" {
		fmt.Print("どの字牌ですか？（東:1,南:2,西:3,北:4,白:5,發:6,中:7）:")
	} else {
		fmt.Print("数牌の数字を指定してください:")
	}
	_, err = fmt.Scan(&num)
	if err != nil {
		log.Fatalln(err)
	}
	n, err := strconv.Atoi(num)
	if err != nil {
		log.Fatalln(err)
	}
	var result []string
	if haiType == "1" {
		for i := 0; i < 4; i++ {
			result = append(result, strconv.Itoa(n-1+i))
		}
	} else if haiType == "2" {
		for i := 0; i < 4; i++ {
			result = append(result, strconv.Itoa(n*4+i+35))
		}
	} else if haiType == "3" {

	} else if haiType == "4" {

	}
	fmt.Println(strings.Join(result, ","))
}
