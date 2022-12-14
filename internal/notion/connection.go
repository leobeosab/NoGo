package notion

import (
	"context"
	"fmt"
	"github.com/dstotijn/go-notion"
)

type PageInfoResponse struct {
	Title           []notion.RichText
	Blocks          []notion.Block
	OutputDirectory []notion.RichText
	Description     []notion.RichText
	Categories      []string
	Tags            []string
	PublishDate     string
	CoverURL        *string
	CoverAlt        []notion.RichText
}

type Connection struct {
	client *notion.Client
}

func NewConnection(apiKey string) *Connection {
	return &Connection{
		notion.NewClient(apiKey),
	}
}

func (conn Connection) FetchDatabase(databaseId string) {
	fmt.Println("Looking for " + databaseId)
	db, err := conn.client.FindDatabaseByID(context.Background(), databaseId)
	if err != nil {
		panic(err)
	}

	fmt.Println(db)
}

func (conn Connection) FetchDatabasePagesBasedOnStatus(status string, databaseId string) (*[]notion.Page, error) {

	query := notion.DatabaseQuery{
		Filter: &notion.DatabaseQueryFilter{
			Property: "Status",
			DatabaseQueryPropertyFilter: notion.DatabaseQueryPropertyFilter{
				Status: &notion.StatusDatabaseQueryFilter{Equals: status},
			},
		},
		//TODO: Handle recursion for making sure we get all of the pages
		// Over 1000 pages on a start up is a lot but not an excuse to be a lazy engineer :)
		PageSize: 1000,
	}

	pages, err := conn.client.QueryDatabase(context.Background(), databaseId, &query)
	if err != nil {
		return nil, err
	}

	return &pages.Results, nil
}

// FetchPageInfo gets all of the blocks, title and custom properties we clue in on
//TODO: Handle recursion for making sure we  get all of the blocks
func (conn Connection) FetchPageInfo(page *notion.Page) (*PageInfoResponse, error) {
	// Fetch Blocks
	res, err := conn.client.FindBlockChildrenByID(context.Background(), page.ID, nil)
	if err != nil {
		return nil, err
	}

	// Get Title
	title := page.Properties.(notion.DatabasePageProperties)["Name"].Title

	// Get Output Directory
	outputDir := page.Properties.(notion.DatabasePageProperties)["Output Directory"].RichText

	// Get Description
	description := page.Properties.(notion.DatabasePageProperties)["Description"].RichText

	// Get Cover Alt
	coverAlt := page.Properties.(notion.DatabasePageProperties)["Covert Alt"].RichText

	// Get Date
	var publishDate string
	pd := page.Properties.(notion.DatabasePageProperties)["Publish Date"].Date
	if pd != nil {
		publishDate = pd.Start.String()
	} else {
		publishDate = ""
	}

	// Get Categories
	c := page.Properties.(notion.DatabasePageProperties)["Categories"].MultiSelect
	categories := make([]string, 0)
	for _, v := range c {
		categories = append(categories, v.Name)
	}

	// Get Tags
	t := page.Properties.(notion.DatabasePageProperties)["Tags"].MultiSelect
	tags := make([]string, 0)
	for _, v := range t {
		tags = append(tags, v.Name)
	}

	// Cover URL
	var coverURL *string
	if page.Cover != nil {
		if page.Cover.File != nil {
			coverURL = &page.Cover.File.URL
		} else if page.Cover.External != nil {
			coverURL = &page.Cover.External.URL
		}
	}

	return &PageInfoResponse{
		Blocks:          res.Results,
		Title:           title,
		OutputDirectory: outputDir,
		CoverURL:        coverURL,
		PublishDate:     publishDate,
		Categories:      categories,
		Tags:            tags,
		Description:     description,
		CoverAlt:        coverAlt,
	}, nil
}
