package main

import (
	"os"
	"testing"
)

const (
	t_SRC = "testdata/src"
	t_DST = "testdata/dst"
)

var (
	t_config *AppConfig
)

func t_setup() {
	// tearing down before test rather than after allows for visual inspection if things go wrong
	t_teardown()

	if err := os.MkdirAll(t_DST, os.ModePerm); nil != err {
		panic(err)
	}
}

func t_teardown() {
	if _, err := os.Stat(t_DST); nil != err || !os.IsNotExist(err) {
		_ = os.RemoveAll(t_DST)
	}
}

func t_stringSlicesMatch(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, val := range slice1 {
		if val != slice2[i] {
			return false
		}
	}

	return true
}

func TestMain(m *testing.M) {
	t_setup()
	t_config, _ = config("testdata/conf.yml")

	code := m.Run()

	os.Exit(code)
}
