package notion

import (
	"context"
	"fmt"
	"github.com/dstotijn/go-notion"
)

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
			Property: "status",
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

//TODO: Handle recursion for making sure we  get all of the blocks
func (conn Connection) FetchPageBlocks(pageId string) ([]notion.Block, error) {

	res, err := conn.client.FindBlockChildrenByID(context.Background(), pageId, nil)
	if err != nil {
		return nil, err
	}

	return res.Results, nil
}
