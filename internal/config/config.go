package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type ClientProxyConfig struct {
	URL     string
	Timeout time.Duration
}

type ServerProxyConfig struct {
	Port            string
	ShutdownTimeout time.Duration
	EnvType         string
}

type AppConfig struct {
	Server ServerProxyConfig
	Client ClientProxyConfig
}

func New() (cfg *AppConfig, err error) {

	cfg, err = readEnvConfig()

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func readEnvConfig() (*AppConfig, error) {
	port := os.Getenv("SERVER_PORT")
	//port := "8080"

	fmt.Printf("port: %d\n", port)

	url := os.Getenv("PROXY_URL")
	//url := "https://jsonplaceholder.typicode.com/posts"

	shutdownTimeout := os.Getenv("SHUTDOWN_TIMEOUT_SECONDS")
	shutdownTimeoutInt, _ := strconv.Atoi(shutdownTimeout)
	//shutdownTimeout := 5

	clientTimeout := os.Getenv("CLIENT_TIMEOUT_SECONDS")
	clientTimeoutInt, _ := strconv.Atoi(clientTimeout)
	//clientTimeout := 5

	envType := os.Getenv("ENV_TYPE")

	return &AppConfig{
		Server: ServerProxyConfig{
			Port:            port,
			ShutdownTimeout: time.Duration(shutdownTimeoutInt),
			EnvType:         envType,
		},
		Client: ClientProxyConfig{
			URL:     url,
			Timeout: time.Duration(clientTimeoutInt),
		},
	}, nil
}
