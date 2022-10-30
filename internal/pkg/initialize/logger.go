package initialize

import (
	"github.com/nite-coder/blackbear/pkg/config"
	"github.com/nite-coder/blackbear/pkg/log"
	"github.com/nite-coder/blackbear/pkg/log/handler/console"
	"github.com/nite-coder/blackbear/pkg/log/handler/json"
)

type LogSetting struct {
	ID               string
	Type             string
	MinLevel         string `mapstructure:"min_level"`
	ConnectionString string `mapstructure:"connection_string"`
}

func Logger() error {
	logSettings := []LogSetting{}
	err := config.Scan("logging", &logSettings)
	if err != nil {
		return err
	}

	logger := log.New()
	appID, _ := config.String("app.id")
	if len(appID) > 0 {
		logger = logger.Str("app_id", appID).Logger()
	}

	for _, logSetting := range logSettings {
		switch logSetting.Type {
		case "console":
			opts := console.ConsoleOptions{DisableColor: false}
			clog := console.New(opts)
			levels := log.GetLevelsFromMinLevel(logSetting.MinLevel)
			logger.AddHandler(clog, levels...)
		case "json":
			jsonHandler := json.New()
			levels := log.GetLevelsFromMinLevel(logSetting.MinLevel)
			logger.AddHandler(jsonHandler, levels...)
		}
	}

	log.SetLogger(logger)
	return nil
}
