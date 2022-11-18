package main

import (
	"fmt"
	"github.com/leobeosab/notion-blogger/internal/markdown"
	"github.com/leobeosab/notion-blogger/internal/notion"
	"os"
)

func main() {
	notionSecret := os.Getenv("NOTION_SECRET")
	conn := notion.NewConnection(notionSecret)

	pages, err := conn.FetchDatabasePagesBasedOnStatus("ready", "9e2d76ca14cd43c9b6b80b517b7209d5")
	if err != nil {
		panic(err)
	}

	blocks, err := conn.FetchPageBlocks((*pages)[0].ID)

	page := markdown.NewPage("Oh no")

	if err = page.ImportNotionBlocks(blocks); err != nil {
		panic(err)
	}

	fmt.Println(page.Build())

	page.DownloadAssets("./ignore/")
}
