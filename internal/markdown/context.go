package markdown

import "context"

type PageConfig struct {
	AssetDirectory  string
	AssetURL        string
	BlocksDirectory string
}

type PageContext struct {
	C      context.Context
	Config *PageConfig
}
