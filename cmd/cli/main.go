package main

import (
	"github.com/leobeosab/notion-blogger/app"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	a := &cli.App{
		Name:  "NoGo",
		Usage: "Turn a Notion database into markdown blog entries",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "database-id",
				Usage: "Database Id containing the pages to be converted",
			},
			&cli.StringFlag{
				Name:  "ready-status",
				Value: "ready",
				Usage: "Status name for pages to convert. IE status:ready on the page property",
			},
			&cli.StringFlag{
				Name:  "asset-path",
				Value: "static/assets/img/posts/$PAGE_URI$/",
				Usage: "Path to downlaod static assets (ie images) to",
			},
			&cli.StringFlag{
				Name:  "asset-url",
				Value: "assets/img/posts/$PAGE_URI$/",
				Usage: "Path from root of site for asseets",
			},
		},

		Action: func(c *cli.Context) error {
			config := &app.NotionMigrationsConfig{
				NotionSecret:   os.Getenv("NOTION_SECRET"),
				DatabaseId:     c.String("database-id"),
				ReadyStatus:    c.String("ready-status"),
				AssetDirectory: c.String("asset-path"),
				AssetURL:       c.String("asset-url"),
			}

			_, err := app.RunNotionMigrations(config)
			return err
		},
	}

	if err := a.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
