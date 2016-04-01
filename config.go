package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
	"regexp"
	"strings"
)

type Config struct {
	Dir   string `default:"/var/log/errors"`
	Debug bool   `default:false`
	Host  string `default:"0.0.0.0"`
	Port  int    `default:"8080"`
}

type App struct {
	Name string
	Key  string
}

var (
	config Config
	apps   []App
)

const (
	appName    string = "breaker"
	keylength  int    = 32
	appVersion string = "0.0.1"
)

func readConfig() error {
	var app App
	err := envconfig.Process(appName, &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	format := "Host: %s\nPort: %d\nDebug: %v\nOutput dir: %s\n"
	_, err = fmt.Printf(format, config.Host, config.Port, config.Debug, config.Dir)
	if err != nil {
		log.Fatal(err.Error())
	}

	re, _ := regexp.Compile(strings.ToUpper(appName) + "_" + "([a-zA-Z0]+)_KEY")

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if sub := re.FindStringSubmatch(pair[0]); sub != nil {
			app.Name = sub[1]
			app.Key = pair[1]
			apps = append(apps, app)
		}
	}

	if len(apps) == 0 {
		log.Fatal("Error: no apps defined, exiting")
	}

	return nil
}
