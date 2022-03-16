package koanfconf

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

func Load(configFileName, envPrefix string, defaultConfig interface{}, config interface{}) *koanf.Koanf {

	k := koanf.New(".")

	// load default configuration from file
	if err := k.Load(structs.Provider(defaultConfig, "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	// load configuration from file
	if err := k.Load(file.Provider(configFileName), yaml.Parser()); err != nil {
		log.Printf("error loading %s: %v", configFileName, err)
	}

	// load environment variables
	if err := k.Load(env.Provider(envPrefix, ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(
			strings.TrimPrefix(s, envPrefix)), "__", ".")
	}), nil); err != nil {
		log.Printf("error loading environment variables: %s", err)
	}

	if err := k.Unmarshal("", config); err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}
	indent, _ := json.MarshalIndent(config, "", "\t")
	cfgStrTemplate := `
	================ Loaded Configuration ================
	%s
	======================================================
	`
	log.Printf(cfgStrTemplate, string(indent))
	return k
}
