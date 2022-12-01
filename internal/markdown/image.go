package markdown

import (
	"bytes"
	"github.com/dstotijn/go-notion"
	"github.com/leobeosab/notion-blogger/internal/utilities"
)

type Image struct {
	Caption   string
	HostedURL string
	AssetURL  string
	FileName  string
}

func (p *PageBuilder) newImage(block notion.ImageBlock) Image {

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

	s, err := p.RichTextArrToString(block.Caption)
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

func (p *PageBuilder) AddImageToPage(block *notion.ImageBlock) error {
	md := p.newImage(*block)

	template, err := p.FetchTemplate("ImageTemplate.md")
	if err != nil {
		return err
	}

	var mdBuffer bytes.Buffer
	err = template.Execute(&mdBuffer, md)
	if err != nil {
		return err
	}

	p.AddContent(mdBuffer.String())

	// Add asset to the page for downloading
	p.AddAsset(md.HostedURL, md.FileName)
	return nil
}
