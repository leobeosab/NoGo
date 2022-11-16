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
	os.Setenv("ASSET_PATH", "assets/img/posts/$PAGE_URI$/")

	//	mdDirectory := os.Getenv("MD_DIRECTORY")
	//	noOutput := os.Getenv("NO_OUTPUT_ASSETS")

	notionSecret := os.Getenv("NOTION_SECRET")
	conn := notion.NewConnection(notionSecret)

	conn.FetchDatabase("9e2d76ca14cd43c9b6b80b517b7209d5")

	pages, err := conn.FetchDatabasePagesBasedOnStatus("ready", "9e2d76ca14cd43c9b6b80b517b7209d5")
	if err != nil {
		panic(err)
	}

	blocks, err := conn.FetchPageBlocks((*pages)[0].ID)

	page := markdown.NewPage("Oh no")

	// Extract this to a functionforloop
	for _, block := range blocks {
		blockType := strings.Replace(reflect.TypeOf(block).String(), "*notion.", "", -1)
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

		case "ImageBlock":
			err := page.AddImageToPage(block.(*not.ImageBlock))
			if err != nil {
				panic(err)
			}
			break
		}
	}

	fmt.Println(page.Build())

	page.DownloadAssets("./ignore/")
}
