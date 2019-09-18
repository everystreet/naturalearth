package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"sync/atomic"

	"github.com/mercatormaps/naturalearth"
	"github.com/mercatormaps/naturalearth/data"

	"github.com/BurntSushi/toml"
	"github.com/gosuri/uiprogress"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	Admin110BoundaryLines = "boundaries-110"
	Admin50BoundaryLines  = "boundaries-50"
	Admin50StatesLines    = "states-50"
)

func main() {
	updateConf := confUpdate()

	switch kingpin.Parse() {
	case updateConf.cmd.FullCommand():
		update(&updateConf)
	}
}

func confUpdate() updateConf {
	cmd := kingpin.Command("update", "Update Natural Earth data.")
	return updateConf{
		cmd:    cmd,
		config: cmd.Flag("config", "Path to config file (.toml)").Short('c').Required().String(),
	}
}

func update(conf *updateConf) {
	config, err := readConfig(*conf.config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read config file '%s': %s\n", *conf.config, err.Error())
		os.Exit(1)
	}

	store := &naturalearth.ElasticSearch{}
	if err := store.Connect(config.Store.Hosts...); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to data stores %v: %s\n", config.Store.Hosts, err.Error())
		os.Exit(1)
	}

	uiprogress.Start()

	concurrency := 2
	tokens := make(chan struct{}, concurrency)
	defer close(tokens)
	for i := 0; i < concurrency; i++ {
		tokens <- struct{}{}
	}

	var numFeatures, storedFeatures uint64
	var wg = sync.WaitGroup{}
	for name, uri := range config.DataSources {
		token := <-tokens
		wg.Add(1)

		go func(name, uri string) {
			defer func() {
				tokens <- token
				wg.Done()
			}()

			source, ok := data.Source(name)
			if !ok {
				fmt.Fprintf(os.Stderr, "Unknown data source '%s'\n", name)
				os.Exit(1)
			}

			source.Label = label(source.Name)
			if err := source.Load(uri, store); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to process '%s': %v\n", name, err)
				os.Exit(1)
			}

			atomic.AddUint64(&numFeatures, uint64(source.NumFeatures()))
			atomic.AddUint64(&storedFeatures, uint64(source.StoredFeatures()))
		}(name, uri)
	}
	wg.Wait()

	uiprogress.Stop()
	fmt.Printf("\nStored %d features (out of %d in total) from %d source(s) successfully\n",
		storedFeatures, numFeatures, len(config.DataSources))
}

type updateConf struct {
	cmd    *kingpin.CmdClause
	config *string
}

func readConfig(path string) (*config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf := config{}
	if err := toml.Unmarshal(buf, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

type config struct {
	DataSources map[string]string `toml:"data-sources"`
	Store       struct {
		Hosts []string `toml:"hosts"`
	} `toml:"data-store"`
}

func label(name string) string {
	return fmt.Sprintf("%-*s", data.MaxNameLen(), name)
}
