package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/spantree/datadumperbeat/beater"
)

func main() {
	err := beat.Run("datadumperbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
