package markdown

import (
	"bytes"
	"github.com/dstotijn/go-notion"
	"github.com/leobeosab/notion-blogger/internal/utilities"
	"os"
	"text/template"
)

type Image struct {
	Caption   string
	HostedURL string
	AssetURL  string
	FileName  string
}

func (p *Page) newImage(block notion.ImageBlock) Image {

	var hostedUrl string
	if block.File != nil {
		hostedUrl = block.File.URL
	} else {
		hostedUrl = block.External.URL
	}

	fileName, err := utilities.GetFileNameFromURL(hostedUrl)
	if err != nil {
		// TODO: handle this gracefully
		panic(err)
	}

	s, err := RichTextArrToString(block.Caption)
	if err != nil {
		panic(err)
	}

	return Image{
		Caption:   s,
		HostedURL: hostedUrl,
		FileName:  fileName,
		AssetURL:  p.AssetURL,
	}
}

func (p *Page) AddImageToPage(block *notion.ImageBlock) error {
	md := p.newImage(*block)

	template, err := template.ParseFiles(os.Getenv("BLOCKS_DIRECTORY") + "/ImageTemplate.md")
	if err != nil {
		return err
	}

	var mdBuffer bytes.Buffer
	err = template.Execute(&mdBuffer, md)
	if err != nil {
		return err
	}

	p.AddBlock(mdBuffer.String())

	// Add asset to the page for downloading
	p.AddAsset(md.HostedURL, md.FileName)
	return nil
}
