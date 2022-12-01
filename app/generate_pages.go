package app

import (
	"github.com/leobeosab/notion-blogger/internal/markdown"
	"github.com/leobeosab/notion-blogger/internal/notion"
	"github.com/leobeosab/notion-blogger/internal/utilities"
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
			continue
		}

		// Create Page Base
		title := markdown.RichTextArrToPlainString(info.Title)
		description := markdown.RichTextArrToPlainString(info.Description)
		mdPage := markdown.NewPage(c.PageContext(), title, page.ID, description, info.PublishDate)

		// Add Cover if it exists
		if info.CoverURL != nil {
			coverAlt := markdown.RichTextArrToPlainString(info.CoverAlt)
			mdPage.AddCover(*info.CoverURL, coverAlt)
		}

		// Set Header
		if err = mdPage.SetHeader(info.Categories, info.Tags); err != nil {
			log.Println("Could not set header for: " + page.URL)
			continue
		}

		// Import from Notion
		if err = mdPage.ImportNotionBlocks(info.Blocks); err != nil {
			log.Println("Could not import blocks for: " + page.URL)
			continue
		}

		// Download Assets to directory
		mdPage.DownloadAssets(config.OutputDirectory)

		// Write to file
		outputDirectory := markdown.RichTextArrToPlainString(info.OutputDirectory)
		if err = utilities.WriteStringToFile(mdPage.Build(), c.Config.OutputDirectory+c.Config.ContentDirectory+outputDirectory+"/", mdPage.ID+".md"); err != nil {
			log.Println("Could not write page to file: " + page.URL)
			continue
		}
	}

	return 0, nil
}
