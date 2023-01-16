package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var i string
	_, err := fmt.Scan(&i)
	if err != nil {
		log.Fatalln(err)
	}
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Fatalln(err)
	}
	if id >= 0 && id <= 35 {
		fmt.Printf("萬子の%d\n", (id+1)/4)
	} else if id >= 36 && id <= 71 {
		// circles
		fmt.Printf("筒子の%d\n", (id-35)/4)
	} else if id >= 72 && id <= 106 {
		// bamboos
		fmt.Printf("索子の%d\n", (id-71)/4)
	} else if id >= 107 && id <= 135 {
		// honors
		switch (id - 106) / 4 {
		case 0:
			fmt.Println("東")
		case 1:
			fmt.Println("南")
		case 2:
			fmt.Println("西")
		case 4:
			fmt.Println("北")
		case 5:
			fmt.Println("白")
		case 6:
			fmt.Println("發")
		case 7:
			fmt.Println("中")
		}
	} else {
		log.Fatalf("input must be between 0 and 135. input is %d", id)
	}
}
