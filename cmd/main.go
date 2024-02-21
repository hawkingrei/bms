package main

import (
	"flag"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/hawkingrei/bms/config"
	"github.com/hawkingrei/bms/service"
	"github.com/pingcap/log"
	"go.uber.org/zap"
)

func flagSet() *flag.FlagSet {
	flagSet := flag.NewFlagSet("redp", flag.ExitOnError)
	flagSet.Bool("version", false, "print version string")
	flagSet.String("config", "", "path to config file")
	return flagSet
}

func loadmeta(configFile string) (meta *config.Configure, err error) {
	if configFile != "" {
		_, err = toml.DecodeFile(configFile, &meta)
		if err != nil {
			return
		}
	}
	return
}

func main() {
	flagSet := flagSet()
	flagSet.Parse(os.Args[1:])
	configFile := flagSet.Lookup("config").Value.String()
	config, err := loadmeta(configFile)
	if err != nil {
		log.Error("load config failed", zap.Error(err))
		os.Exit(0)
	}
	service, err := service.NewService(config)
	if err != nil {
		log.Error("load config failed", zap.Error(err))
		os.Exit(0)
	}
	service.Run()
}
