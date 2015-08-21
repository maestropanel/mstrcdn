package main

import (
	"log"
	"code.google.com/p/gcfg"
	"fmt"
	"os"
	"flag"
)

var config Config
var configPath string

func init() {
	flag.StringVar(&configPath, "path", "/etc/maestropanel/agent/config/mstrcdn.conf", "Configuration file path")
	
	if(!nginx.fileExists(configPath)){
		fmt.Println("Config file not found: "+ configPath)
		os.Exit(1)
	}

	err := gcfg.ReadFileInto(&config, configPath)

	if err != nil {
		fmt.Println("Failed to parse config file: ", err)
		os.Exit(1)
	}
}

func main() {
	err := StartAgent(config.Api.SecretKey, config.Api.Port)
	if(err != nil){
		log.Fatalln(err.Error())		
	}
}

type Config struct {
	Api struct {
		Port       int
		SecretKey  string
		ConfigRoot string
		TemplatePath string
	}
}
