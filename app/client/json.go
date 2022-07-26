package client

import (
	"fmt"
	"log"

	"github.com/m7shapan/njson"
	"github.com/manticoresoftware/go-sdk/manticore"
)

type ManticoreJsonClient struct {
	Client manticore.Client
	Limit  int
}

func (c ManticoreJsonClient) FindAll(index string) (*ManticoreClientSearchResult, error) {

	res, err := c.Client.Json(
		"json/search",
		fmt.Sprintf(`{"index":"%s","query":{"match":{"title,description":"*"}},"limit":%d}`, index, c.Limit),
	)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return c.extractSearchResult(res), nil
}

func (c ManticoreJsonClient) FindByExpression(index string, expression string) (*ManticoreClientSearchResult, error) {

	res, err := c.Client.Json(
		"json/search",
		fmt.Sprintf(`{"index":"%s","query":{"match":{"title,description":"%s"}},"limit":%d}`, index, expression, c.Limit),
	)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return c.extractSearchResult(res), nil
}

func (c ManticoreJsonClient) extractSearchResult(jsonResult manticore.JsonAnswer) *ManticoreClientSearchResult {

	type resultHits struct {
		Hits []map[string]interface{} `njson:"hits.hits.#._source"`
	}

	hits := resultHits{}
	err := njson.Unmarshal([]byte(jsonResult.Answer), &hits)
	if err != nil {
		log.Printf("error: %v", err)
	}

	results := make(map[int]map[string]interface{})

	for index, hit := range hits.Hits {
		results[index] = hit
	}

	return &ManticoreClientSearchResult{results}
}
