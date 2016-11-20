package main

import (
	"fmt"
	"log"

	"github.com/bodgix/gowikiwordcount/wiki"
)

func main() {
	page, err := wiki.GetPage("21721040")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("###", page.Title)
	fmt.Println(page.Words)
}
