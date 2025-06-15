package config

import (
	"fmt"
	"os"
)

type ClientProxyConfig struct {
	URL string
}

type ServerProxyConfig struct {
	Port string
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

	fmt.Printf("URL: %d\n", port)

	return &AppConfig{
		Server: ServerProxyConfig{
			Port: port,
		},
		Client: ClientProxyConfig{
			URL: url,
		},
	}, nil
}
