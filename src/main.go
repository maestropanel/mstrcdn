package main


import (
	"code.google.com/p/gcfg"
	"fmt"
	"os"
)

var config Config

func init() {
	err := gcfg.ReadFileInto(&config, "config.ini")

	if err != nil {
		fmt.Println("Failed to parse gcfg data: ", err)
		os.Exit(1)
	}
}


func main(){
	StartAgent(config.Api.SecretKey, config.Api.Port)
}

type Config struct {
	Api struct {
		Port           int
		SecretKey      string
	}
}