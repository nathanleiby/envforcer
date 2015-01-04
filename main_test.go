package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var currentDir = "./"

// CheckError fails if there is an error
func CheckError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func Test_GetRequiredEnvFromFile(t *testing.T) {
	for _, file := range []string{
		"example-env.yml",
		"example-env.json",
	} {
		requiredEnv, err := getRequiredEnvFromFile(file)
		CheckError(t, err)
		assert.Equal(t, len(requiredEnv.Env), 2, "incorrect required env count in file %s", file)
	}
}

func Test_GetRequiredEnvFromFileError(t *testing.T) {
	file := "does-not-exist"
	_, err := getRequiredEnvFromFile(file)
	if err == nil {
		t.Fatalf("Expected %s to not exist", file)
	}

}
