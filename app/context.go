package app

import (
	"context"
	"github.com/leobeosab/notion-blogger/internal/markdown"
)

type NotionMigrationsConfig struct {
	BlocksDirectory string
	AssetURL        string
	AssetDirectory  string
	DatabaseId      string
	ReadyStatus     string
	NotionSecret    string
}

type NotionMigrationsContext struct {
	C      context.Context
	Config *NotionMigrationsConfig
}

func NewNotionMigrationsContext(config *NotionMigrationsConfig) *NotionMigrationsContext {
	return &NotionMigrationsContext{
		C:      context.Background(),
		Config: config,
	}
}

func (n *NotionMigrationsContext) PageContext() *markdown.PageContext {
	return &markdown.PageContext{
		C: n.C,
		Config: &markdown.PageConfig{
			AssetDirectory:  n.Config.AssetDirectory,
			AssetURL:        n.Config.AssetURL,
			BlocksDirectory: n.Config.BlocksDirectory,
		},
	}
}
