package config

import (
	"encoding/json"
	logPkg "log"
	"regexp"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/mehditeymorian/hermes/internal/db"
	"github.com/mehditeymorian/hermes/internal/emq"
	"github.com/mehditeymorian/hermes/internal/log"
	"github.com/tidwall/pretty"
	"go.uber.org/zap"
)

const PREFIX = "HERMES_"

type Config struct {
	Emq    emq.Config `koanf:"emq"`
	DB     db.Config  `koanf:"db"`
	Logger log.Config `koanf:"logger"`
}

func Load(path string) Config {
	var cfg Config

	k := koanf.New(".")

	// load default configuration
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		zap.L().Fatal("error loading default config", zap.Error(err))
	}

	// load configuration from file
	if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
		zap.L().Warn("error loading config.yaml", zap.Error(err))
	}

	// load environment variables
	cb := func(key string, value string) (string, interface{}) {
		finalKey := strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(key, PREFIX)), "__", ".")

		if strings.Contains(value, ",") {
			// remove all the whitespace from value
			// split the value using comma
			finalValue := strings.Split(removeWhitespace(value), ",")

			return finalKey, finalValue
		}

		return finalKey, value
	}
	if err := k.Load(env.ProviderWithValue(PREFIX, ".", cb), nil); err != nil {
		zap.L().Warn("error loading environment variables", zap.Error(err))
	}

	if err := k.Unmarshal("", &cfg); err != nil {
		zap.L().Fatal("error unmarshalling config", zap.Error(err))
	}

	indent, _ := json.MarshalIndent(cfg, "", "\t")
	indent = pretty.Color(indent, nil)
	cfgStrTemplate := `
	================ Loaded Configuration ================
	%s
	======================================================
	`
	logPkg.Printf(cfgStrTemplate, string(indent))

	return cfg
}

// removeWhitespace remove all the whitespaces from the input.
func removeWhitespace(in string) string {
	compile := regexp.MustCompile(`\s+`)

	return compile.ReplaceAllString(in, "")
}
