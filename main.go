package main

// Gets required env from command line args and/or file

// TODO: support more complex env validation, like "isString" or "length" or 0<val<10
// TODO: Check if env var names are invalid (not uppercase A-Z?)
// TODO: Log which env is checked for

import (
	"flag"
	"fmt"
	"gopkg.in/v1/yaml"
	"io/ioutil"
	"os"
)

// RequiredEnv is the structure of a file containing required env vars
type RequiredEnv struct {
	Env []string `yaml:"env"`
}

func main() {
	fileFlag := flag.String("file", "", "path to json/yaml file containing required env")
	flag.Parse()
	envFile := string(*fileFlag)

	// Command line arguments are passed in as required ENV vars
	envList := flag.Args()

	// If a file is given, parse it to find required env vars
	if envFile != "" {
		requiredEnv, err := getRequiredEnvFromFile(envFile)
		if err != nil {
			fmt.Println("error: unable to read file", envFile)
			os.Exit(2)
		}
		envList = append(envList, requiredEnv.Env...)
	}

	if len(envList) == 0 {
		fmt.Println("error: no required environment vars were specified")
		os.Exit(3)
	}

	if !areEnvVarsSet(envList) {
		fmt.Println("error: some required env vars are not set")
		os.Exit(1)
	}

	os.Exit(0)
}

func getRequiredEnvFromFile(f string) (required RequiredEnv, err error) {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return
	}
	if err = yaml.Unmarshal(data, &required); err != nil {
		return
	}
	return
}

func areEnvVarsSet(requiredEnv []string) bool {
	allSet := true
	for _, val := range requiredEnv {
		//fmt.Println("Checking if", val, "is set")
		if os.Getenv(val) == "" {
			fmt.Println("Env var", val, "is unset (or empty string)")
			allSet = false
		}
	}
	return allSet
}
