package markdown

import (
	"github.com/dstotijn/go-notion"
	"github.com/leobeosab/notion-blogger/internal/utilities"
	"os"
	"reflect"
	"strings"
)

type Page struct {
	sbContent      strings.Builder
	Title          string
	ID             string
	NotionID       string
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

func NewPage(title string, notionID string) *Page {
	id := strings.Replace(strings.ToLower(title), " ", "-", -1)

	return &Page{
		sbContent:      strings.Builder{},
		Title:          title,
		ID:             id,
		NotionID:       notionID,
		Assets:         make([]PageAsset, 0),
		AssetDirectory: strings.Replace(os.Getenv("ASSET_PATH"), "$PAGE_URI$", id, -1),
		AssetURL:       strings.Replace(os.Getenv("ASSET_URL"), "$PAGE_URI$", id, -1),
	}
}

func (p *Page) ImportNotionBlocks(blocks []notion.Block) error {
	for _, block := range blocks {
		blockType := strings.Replace(reflect.TypeOf(block).String(), "*notion.", "", -1)
		switch blockType {
		case "Heading1Block":
			err := p.AddHeading1ToPage(block.(*notion.Heading1Block))
			if err != nil {
				return err
			}
			break

		case "ParagraphBlock":
			err := p.AddParagraphToPage(block.(*notion.ParagraphBlock))
			if err != nil {
				return err
			}
			break

		case "CodeBlock":
			err := p.AddCodeToPage(block.(*notion.CodeBlock))
			if err != nil {
				return err
			}
			break

		case "ImageBlock":
			err := p.AddImageToPage(block.(*notion.ImageBlock))
			if err != nil {
				return err
			}
			break
		}
	}

	return nil
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
