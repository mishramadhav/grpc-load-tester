package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	TargetServer TargetServer      `yaml:"targetServer"`
	Services     []Service         `yaml:"services"`
	LoadPattern  LoadPattern       `yaml:"loadPattern"`
	RateLimiting RateLimiting      `yaml:"rateLimiting"`
	Metadata     map[string]string `yaml:"metadata,omitempty"`
	TLS          *TLS              `yaml:"tls,omitempty"`
}

type TargetServer struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Service struct {
	Name    string   `yaml:"name"`
	Methods []Method `yaml:"methods"`
}

type Method struct {
	Name  string                 `yaml:"name"`
	Input map[string]interface{} `yaml:"input"`
}

type LoadPattern struct {
	Type            string        `yaml:"type"`
	ConcurrentUsers int           `yaml:"concurrentUsers"`
	Duration        time.Duration `yaml:"duration"`
	RampUp          RampUp        `yaml:"rampUp,omitempty"`
	Cooldown        Cooldown      `yaml:"cooldown,omitempty"`
}

type RampUp struct {
	Duration time.Duration `yaml:"duration"`
}

type Cooldown struct {
	Duration time.Duration `yaml:"duration"`
}

type RateLimiting struct {
	MaxRequestsPerSecond int `yaml:"maxRequestsPerSecond"`
}

type TLS struct {
	Enabled  bool   `yaml:"enabled"`
	CertFile string `yaml:"certFile,omitempty"`
	KeyFile  string `yaml:"keyFile,omitempty"`
}

func ParseConfigFile(filename string) (Config, error) {
	var config Config

	data, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
