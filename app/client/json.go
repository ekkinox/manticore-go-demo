package client

import (
	"github.com/manticoresoftware/go-sdk/manticore"
)

type ManticoreJsonClient struct {
	Client manticore.Client
	Limit  int
}

func (c ManticoreJsonClient) FindAll(index string) (*ManticoreClientSearchResult, error) {
	//todo
	return &ManticoreClientSearchResult{}, nil
}

func (c ManticoreJsonClient) FindByExpression(index string, expression string) (*ManticoreClientSearchResult, error) {
	//todo
	return &ManticoreClientSearchResult{}, nil
}
