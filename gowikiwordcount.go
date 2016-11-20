package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bodgix/gowikiwordcount/tools"
	"github.com/bodgix/gowikiwordcount/wiki"
)

func main() {
	pageID := os.Args[1]
	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	page, err := wiki.GetPage(pageID)
	if err != nil {
		log.Fatal(err)
	}

	wordsWithCount := tools.TopN(page.Words, n)
	fmt.Println("URL:", page.URL)
	fmt.Println("Title:", page.Title)
	fmt.Println()
	for _, word := range wordsWithCount {
		fmt.Println("-", word.Count, word.Word)
	}
}
