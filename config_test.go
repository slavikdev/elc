package main

import (
	"flag"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	resetFlags()
	config := NewConfig()
	assert.Equal(t, defPort, config.Port)
	assert.Equal(t, defHost, config.Host)

	testPort := 1234
	flag.Set("p", strconv.Itoa(testPort))
	testHost := "example.com"
	flag.Set("h", testHost)
	assert.Equal(t, testPort, config.Port)
	assert.Equal(t, testHost, config.Host)
}

func resetFlags() {
	flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
}
