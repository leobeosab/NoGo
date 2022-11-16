package markdown

import (
	"bytes"
	"github.com/dstotijn/go-notion"
	"github.com/leobeosab/notion-blogger/internal/utilities"
	"text/template"
)

type Image struct {
	Caption      string
	HostedURL    string
	DownloadPath string
}

func newImage(block notion.ImageBlock, pageId string) Image {

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

	downloadPath := utilities.GetAssetPath(fileName, pageId)

	s, err := RichTextArrToString(block.Caption)
	if err != nil {
		panic(err)
	}

	return Image{
		Caption:      s,
		HostedURL:    hostedUrl,
		DownloadPath: downloadPath,
	}
}

func (p *Page) AddImageToPage(block *notion.ImageBlock) error {
	md := newImage(*block, p.ID)

	template, err := template.ParseFiles("blocks/ImageTemplate.md")
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
	p.AddAsset(md.HostedURL, md.DownloadPath)
	return nil
}
