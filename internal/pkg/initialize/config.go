package initialize

import (
	stdlog "log"
	"path/filepath"

	"github.com/nite-coder/blackbear/pkg/config"
	"github.com/nite-coder/blackbear/pkg/config/provider/env"
	"github.com/nite-coder/blackbear/pkg/config/provider/file"
)

func Config(confPath string) error {
	// env
	envProvider := env.New()
	config.AddProvider(envProvider)

	// file
	fileProvider := file.New()
	if confPath != "" {
		path, configName := filepath.Split(confPath)
		if path == "" {
			path = "./"
		}
		fileProvider.AddPath(path)
		fileProvider.SetConfigName(configName)
	} else {
		fileProvider.SetConfigName("app.yaml")
	}

	err := fileProvider.Load()
	if err != nil {
		stdlog.Println("initialize: no file provider :", err)
	} else {
		config.AddProvider(fileProvider)
	}

	return nil
}
