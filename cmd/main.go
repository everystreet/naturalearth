package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/mercatormaps/go-geojson"
	"github.com/mercatormaps/naturalearth"

	"github.com/BurntSushi/toml"
	"github.com/gosuri/uiprogress"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	AdminCountries110 = "countries-110"
	AdminCountries50  = "countries-50"
	AdminCountries10  = "countries-10"
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

	opts := []naturalearth.Option{
		naturalearth.RenameProperties(map[string]string{
			"NAME_EN": "name_en",
		}),
		naturalearth.AddProperties(geojson.Property{
			Name:  "source",
			Value: "naturalearth",
		}),
	}

	type updater func() (*naturalearth.UpdateProgress, error)
	updaters := make(map[string]updater)
	for name, uri := range config.DataSources {
		switch name {
		case AdminCountries110:
			fallthrough
		case AdminCountries50:
			fallthrough
		case AdminCountries10:
			props, ok := props[name]
			if !ok {
				fmt.Fprintf(os.Stderr, "Missing properties for data source '%s'\n", name)
				os.Exit(1)
			}

			suffix := "_" + strconv.Itoa(props.suffix)
			updaters[name] = func() (*naturalearth.UpdateProgress, error) {
				return naturalearth.Update(uri, suffix, store,
					append(opts, naturalearth.AddProperties(props.new...))...)
			}
		default:
			fmt.Fprintf(os.Stderr, "Unknown data source '%s': %s\n", name, uri)
			os.Exit(1)
		}

		fmt.Printf("%s: %s\n", name, uri)
	}

	uiprogress.Start()

	var numFeatures uint64
	wg := sync.WaitGroup{}
	for name, update := range updaters {
		wg.Add(1)
		go func(name string, update updater) {
			prog, err := update()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to process '%s': %s\n", name, err.Error())
				os.Exit(1)
			}

			bar := uiprogress.AddBar(int(prog.Total))
			bar.AppendCompleted()
			bar.PrependFunc(func(b *uiprogress.Bar) string {
				return name
			})

			for n := range prog.Progress {
				bar.Set(int(n))
			}

			bar.Set(int(prog.Total))
			atomic.AddUint64(&numFeatures, uint64(prog.Total))
			wg.Done()
		}(name, update)
	}
	wg.Wait()

	uiprogress.Stop()
	fmt.Printf("Processed %d features from %d source(s) successfully\n", numFeatures, len(updaters))
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

type properties struct {
	suffix int
	new    []geojson.Property
}

var props = map[string]properties{
	AdminCountries110: properties{
		suffix: 0,
		new: []geojson.Property{
			{
				Name:  "type",
				Value: "country",
			},
			{
				Name:  "max_zoom",
				Value: 2,
			},
		},
	},
	AdminCountries50: properties{
		suffix: 3,
		new: []geojson.Property{
			{
				Name:  "type",
				Value: "country",
			},
			{
				Name:  "max_zoom",
				Value: 3,
			},
		},
	},
	AdminCountries10: properties{
		suffix: 4,
		new: []geojson.Property{
			{
				Name:  "type",
				Value: "country",
			},
			{
				Name:  "max_zoom",
				Value: 6,
			},
		},
	},
}
