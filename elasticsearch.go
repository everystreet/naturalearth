package naturalearth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/everystreet/go-geojson/v2"
)

type ElasticSearch struct {
	cli           *elasticsearch.Client
	clientVersion string
	serverVersion string
}

func (e *ElasticSearch) Connect(hosts ...string) error {
	e.clientVersion = elasticsearch.Version

	cli, err := elasticsearch.NewClient(elasticsearch.Config{Addresses: hosts})
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	e.cli = cli

	resp, err := cli.Info()
	if err != nil {
		return fmt.Errorf("failed to get server info: %w", err)
	} else if resp.IsError() {
		return fmt.Errorf("[%s]: %s", resp.Status(), resp.String())
	}

	var info map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return fmt.Errorf("failed to decode info: %w", err)
	}
	e.serverVersion = info["version"].(map[string]interface{})["number"].(string)

	return nil
}

func (e *ElasticSearch) Insert(feat *geojson.Feature, key string) error {
	body, err := json.Marshal(feat)
	if err != nil {
		return fmt.Errorf("failed to marshal feature: %w", err)
	}

	req := esapi.IndexRequest{
		Index:      "features",
		DocumentID: key,
		Body:       bytes.NewBuffer(body),
	}

	resp, err := req.Do(context.Background(), e.cli)
	if err != nil {
		return fmt.Errorf("failed to process request: %w", err)
	}
	defer resp.Body.Close()
	if resp.IsError() {
		return fmt.Errorf("[%s]: %s", resp.Status(), resp.String())
	}
	return nil
}
