package main

import (
	"flag"
	"os"
	"fmt"	
	"github.com/datauniverse-lab/earth-asd/factory"
	"github.com/datauniverse-lab/earth-asd/process"
)

func main() {

	configSet := flag.String("config_set", os.Getenv("CONFIG_SET"), "config set")

	var fac factory.Factory

	if *configSet == "LIVE" {
		configHome := flag.String("config_home", os.Getenv("CONFIG_HOME"), "app env")
		configURL := flag.String("config_url", os.Getenv("CONFIG_URL"), "config url")

		flag.Parse()

		fmt.Println("CONFIG_HOME:", *configHome)
		fmt.Println("CONFIG_URL:", *configURL)

		fac = factory.Factory{JSONConfigPath: *configHome, JSONConfigURL: *configURL, ConfigSet: *configSet}

	} else {
		appEnv := flag.String("app-env", os.Getenv("CONFIG_HOME"), "app env")

		flag.Parse()

		if *appEnv == "" {
			*appEnv = `./`
		}
		fac = factory.Factory{JSONConfigPath: *appEnv, ConfigSet: *configSet}
	}

	fmt.Println("CONFIG_SET:", *configSet)
	fac.Initialize()

	if fac.Property.TeslaProcess {
		defer func() {
			fac.GrpcClient.Close()
		}()
	}
	
	var proc process.ASDProcess
	proc.Initialize(&fac)
	proc.Processing()
}
