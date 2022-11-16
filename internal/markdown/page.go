package markdown

import (
	"github.com/leobeosab/notion-blogger/internal/utilities"
	"os"
	"strings"
)

type Page struct {
	sbContent      strings.Builder
	Title          string
	ID             string
	tags           []string
	categories     []string
	Assets         []PageAsset
	AssetDirectory string // Actual base directory to write to, usually static/...
	AssetURL       string // URL to base directory. For example if AssetDirectory was static/images/ AssetURL may just be images/ depending on your static site generator
}

type PageAsset struct {
	ContentURL string
	FileName   string
}

func NewPage(title string) *Page {
	id := strings.Replace(strings.ToLower(title), " ", "-", -1)

	return &Page{
		sbContent:      strings.Builder{},
		Title:          title,
		ID:             id,
		Assets:         make([]PageAsset, 0),
		AssetDirectory: strings.Replace(os.Getenv("ASSET_PATH"), "$PAGE_URI$", id, -1),
		AssetURL:       strings.Replace(os.Getenv("ASSET_URL"), "$PAGE_URI$", id, -1),
	}
}

func (p *Page) AddAsset(contentURL string, fileName string) {
	p.Assets = append(p.Assets, PageAsset{
		ContentURL: contentURL,
		FileName:   fileName,
	})
}

func (p *Page) AddBlock(block string) {
	p.sbContent.WriteString("\n" + block)
}

// DownloadAssets returns number of assets downloaded
func (p *Page) DownloadAssets(outputDirectory string) int {
	if os.Getenv("NO_OUTPUT_ASSETS") == "TRUE" {
		return 0
	}

	err := utilities.MakeDirectoryIfNotExists(outputDirectory + p.AssetDirectory)
	if err != nil {
		return 0
	}

	assetCount := 0
	for _, asset := range p.Assets {
		utilities.DownloadFile(asset.ContentURL, outputDirectory+p.AssetDirectory+asset.FileName)
		assetCount++
	}

	return assetCount
}

func (p *Page) Build() string {

	return p.sbContent.String()
}
