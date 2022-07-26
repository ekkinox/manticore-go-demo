package client

import (
	"fmt"
	"log"

	"github.com/manticoresoftware/go-sdk/manticore"
)

type ManticoreSqlClient struct {
	Client manticore.Client
	Limit  int
}

func (c ManticoreSqlClient) FindAll(index string) (*ManticoreClientSearchResult, error) {

	res, err := c.Client.Sphinxql(fmt.Sprintf(`select * from %s limit %d`, index, c.Limit))
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return c.extractSearchResult(res[0]), nil
}

func (c ManticoreSqlClient) FindByExpression(index string, expression string) (*ManticoreClientSearchResult, error) {

	res, err := c.Client.Sphinxql(fmt.Sprintf(`select * from %s where match('%s') limit %d`, index, expression, c.Limit))
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return c.extractSearchResult(res[0]), nil
}

func (c ManticoreSqlClient) extractSearchResult(sqlResult manticore.Sqlresult) *ManticoreClientSearchResult {

	results := make(map[int]map[string]interface{})

	for ri, row := range sqlResult.Rows {
		result := make(map[string]interface{})
		for ci, col := range row {
			result[sqlResult.Schema[ci].Name] = col
		}

		results[ri] = result
	}

	return &ManticoreClientSearchResult{results}
}
