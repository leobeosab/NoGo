package app

import (
	"fmt"
	"github.com/leobeosab/notion-blogger/internal/markdown"
	"github.com/leobeosab/notion-blogger/internal/notion"
)

func RunNotionMigration(notionSecret string, pageDatabaseID string, migrateStatus string) (int, error) {
	conn := notion.NewConnection(notionSecret)

	pages, err := conn.FetchDatabasePagesBasedOnStatus(migrateStatus, pageDatabaseID)
	if err != nil {
		panic(err)
	}

	pageID := (*pages)[0].ID
	blocks, err := conn.FetchPageBlocks(pageID)

	page := markdown.NewPage("Oh no", pageID)

	if err = page.ImportNotionBlocks(blocks); err != nil {
		panic(err)
	}

	fmt.Println(page.Build())

	page.DownloadAssets("./ignore/")
	return 0, nil
}
