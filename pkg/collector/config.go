package collector

import "time"

// TODO: should have output destination
type Config struct {
	Interval  time.Duration
	BatchSize int
}

func NewConfig() Config {
	return Config{
		Interval:  1 * time.Second,
		BatchSize: 10,
	}
}
