package main

import (
	"fmt"
	not "github.com/dstotijn/go-notion"
	"github.com/leobeosab/notion-blogger/internal/markdown"
	"github.com/leobeosab/notion-blogger/internal/notion"
	"os"
	"reflect"
	"strings"
)

func main() {
	notionSecret := os.Getenv("NOTION_SECRET")
	conn := notion.NewConnection(notionSecret)

	conn.FetchDatabase("9e2d76ca14cd43c9b6b80b517b7209d5")

	pages, err := conn.FetchDatabasePagesBasedOnStatus("ready", "9e2d76ca14cd43c9b6b80b517b7209d5")
	if err != nil {
		panic(err)
	}

	fmt.Println(pages)

	blocks, err := conn.FetchPageBlocks((*pages)[0].ID)

	page := markdown.NewPage("Oh no")

	// Extract this to a functionforloop
	for _, block := range blocks {
		blockType := strings.Replace(reflect.TypeOf(block).String(), "*notion.", "", -1)
		fmt.Println(blockType)
		switch blockType {
		case "Heading1Block":
			err := page.AddHeading1ToPage(block.(*not.Heading1Block))
			if err != nil {
				panic(err)
			}
			break

		case "ParagraphBlock":
			err := page.AddParagraphToPage(block.(*not.ParagraphBlock))
			if err != nil {
				panic(err)
			}
			break

		case "CodeBlock":
			err := page.AddCodeToPage(block.(*not.CodeBlock))
			if err != nil {
				panic(err)
			}
			break
		}
	}

}
