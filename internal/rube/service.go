package rube

func NewService(conf *Config) *Service {
	return &Service{conf: conf}
}

type Service struct {
	conf *Config
}