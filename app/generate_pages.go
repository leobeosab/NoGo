package app

import (
	"fmt"
	"github.com/leobeosab/notion-blogger/internal/markdown"
	"github.com/leobeosab/notion-blogger/internal/notion"
)

func RunNotionMigrations(config *NotionMigrationsConfig) (int, error) {
	c := NewNotionMigrationsContext(config)

	conn := notion.NewConnection(config.NotionSecret)

	pages, err := conn.FetchDatabasePagesBasedOnStatus(config.ReadyStatus, config.DatabaseId)
	if err != nil {
		panic(err)
	}

	pageID := (*pages)[0].ID
	blocks, err := conn.FetchPageBlocks(pageID)

	page := markdown.NewPage(c.PageContext(), "oh yeah", pageID)

	if err = page.ImportNotionBlocks(blocks); err != nil {
		panic(err)
	}

	fmt.Println(page.Build())

	page.DownloadAssets("./ignore/")
	return 0, nil
}
