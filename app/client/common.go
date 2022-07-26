package client

import (
	"log"

	"github.com/manticoresoftware/go-sdk/manticore"
	"github.com/spf13/viper"
)

type ManticoreClientSearchResult struct {
	Results map[int]map[string]interface{}
}

type ManticoreClient interface {
	FindAll(index string) (*ManticoreClientSearchResult, error)
	FindByExpression(index string, expression string) (*ManticoreClientSearchResult, error)
}

func InitManticoreClient() (ManticoreClient, error) {

	host := viper.GetString("MANTICORE_HOST")
	port := viper.GetInt("MANTICORE_PORT")
	mode := viper.GetString("MANTICORE_MODE")
	limit := viper.GetInt("MANTICORE_LIMIT")

	client := manticore.NewClient()
	client.SetServer(host, uint16(port))

	opened, err := client.Open()
	if !opened {
		log.Printf("error: %v", err)
		return nil, err
	}

	if mode == "json" {
		return ManticoreJsonClient{client, limit}, nil
	} else {
		return ManticoreSqlClient{client, limit}, nil
	}
}
