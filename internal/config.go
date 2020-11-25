package internal

import "os"

type Config struct {
	Port string
	Local bool
	WebEnabled bool
	WorkerEnabled bool
	GeneratorEnabled bool
}

func LoadConfig() (*Config, error) {
	var conf Config
	conf.Port = os.Getenv("PORT")
	conf.Local = os.Getenv("LOCAL") == "TRUE"
	conf.WebEnabled = os.Getenv("WEB_ENABLED") == "TRUE"
	conf.WorkerEnabled = os.Getenv("WORKER_ENABLED") == "TRUE"
	conf.GeneratorEnabled = os.Getenv("GENERATOR_ENABLED") == "TRUE"
	return &conf, nil
}