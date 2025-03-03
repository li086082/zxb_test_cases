package main

import (
	"zxb_test_cases/cases"
	"zxb_test_cases/config"
)

func main() {
	config.New("config", "config")
	cases.Execute()
}
