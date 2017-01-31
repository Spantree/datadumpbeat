package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/spantree/datadumpbeat/beater"
)

func main() {
	err := beat.Run("datadumpbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
