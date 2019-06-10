package naturalearth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mercatormaps/go-geojson"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/pkg/errors"
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
		return errors.Wrap(err, "failed to connect")
	}
	e.cli = cli

	resp, err := cli.Info()
	if err != nil {
		return errors.Wrap(err, "failed to get server info")
	} else if resp.IsError() {
		return fmt.Errorf("[%s]: %s", resp.Status(), resp.String())
	}

	var info map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return errors.Wrap(err, "failed to decode info")
	}
	e.serverVersion = info["version"].(map[string]interface{})["number"].(string)

	return nil
}

func (e *ElasticSearch) Insert(feat *geojson.Feature, idSuffix string) (string, error) {
	var name string
	if err := feat.Properties.GetType("name_en", &name); err != nil {
		return "", errors.Wrap(err, "missing or invalid 'name_en' property")
	}
	key := strings.NewReplacer(" ", "_").Replace(strings.ToLower(name)) + idSuffix

	body, err := json.Marshal(feat)
	if err != nil {
		return key, errors.Wrap(err, "failed to marshal feature")
	}

	req := esapi.IndexRequest{
		Index:      "features",
		DocumentID: key,
		Body:       bytes.NewBuffer(body),
	}

	resp, err := req.Do(context.Background(), e.cli)
	if err != nil {
		return key, errors.Wrap(err, "failed to process request")
	}
	defer resp.Body.Close()
	if resp.IsError() {
		return key, fmt.Errorf("[%s]: %s", resp.Status(), resp.String())
	}

	return key, nil
}
