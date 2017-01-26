// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period time.Duration `config:"period"`
	Locale string `config:"locale"`
	EventsToPublish int `config:"events_per_period"`
}

var DefaultConfig = Config{
	Period: 1 * time.Second,
	Locale: "en",
	EventsToPublish: 10,
}
