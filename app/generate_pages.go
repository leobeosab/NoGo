package app

import (
	"fmt"
	"github.com/leobeosab/notion-blogger/internal/markdown"
	"github.com/leobeosab/notion-blogger/internal/notion"
	"log"
)

func RunNotionMigrations(config *NotionMigrationsConfig) (int, error) {
	c := NewNotionMigrationsContext(config)

	conn := notion.NewConnection(config.NotionSecret)

	pages, err := conn.FetchDatabasePagesBasedOnStatus(config.ReadyStatus, config.DatabaseId)
	if err != nil {
		panic(err)
	}

	for _, page := range *pages {
		info, err := conn.FetchPageInfo(&page)
		if err != nil {
			log.Println("Could not fetch blocks for page: " + page.URL)
		}

		// Create Page Base
		mdPage := markdown.NewPage(c.PageContext(), markdown.RichTextArrToPlainString(info.Title), page.ID)

		// Add Cover if it exists
		if info.CoverURL != nil {
			mdPage.AddCover(*info.CoverURL)
		}

		// Import from Notion
		if err = mdPage.ImportNotionBlocks(info.Blocks); err != nil {
			log.Println("Could not import blocks for: " + page.URL)
		}

		// Build to String
		fmt.Println(mdPage.Title + ":")
		fmt.Println(mdPage.Build())

		// Download Assets to directory
		mdPage.DownloadAssets(config.OutputDirectory)

		// Write to file
	}

	return 0, nil
}
