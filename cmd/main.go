package main

import (
	"github.com/leobeosab/notion-blogger/app"
	"log"
	"os"
)

func main() {
	notionSecret := os.Getenv("NOTION_SECRET")
	databaseId := os.Getenv("PAGE_DATABASE_ID")
	targetStatus := os.Getenv("MIGRATE_STATUS")

	config := &app.NotionMigrationsConfig{
		NotionSecret:    notionSecret,
		DatabaseId:      databaseId,
		ReadyStatus:     targetStatus,
		AssetDirectory:  os.Getenv("ASSET_PATH"),
		AssetURL:        os.Getenv("ASSET_URL"),
		BlocksDirectory: "",
	}

	_, err := app.RunNotionMigrations(config)
	if err != nil {
		log.Fatal(err)
	}
}
