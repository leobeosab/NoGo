package main

import (
	"github.com/leobeosab/notion-blogger/app"
	"os"
)

func main() {
	notionSecret := os.Getenv("NOTION_SECRET")
	databaseId := os.Getenv("PAGE_DATABASE_ID")
	targetStatus := os.Getenv("MIGRATE_STATUS")

	app.RunNotionMigration(notionSecret, databaseId, targetStatus)
}
