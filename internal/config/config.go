package config

import (
	"fmt"
	"github.com/emillamm/envx"
	"os"
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

	var env envx.EnvX = os.Getenv

	port, _ := env.String("SERVER_PORT").Default("8080")
	//port := "8080"

	fmt.Printf("port: %d\n", port)

	url, _ := env.String("PROXY_URL").Default("https://jsonplaceholder.typicode.com/posts")
	//url := "https://jsonplaceholder.typicode.com/posts"

	shutdownTimeout, _ := env.Duration("SHUTDOWN_TIMEOUT_SECONDS").Default(5)
	//shutdownTimeout := 5

	clientTimeout, _ := env.Duration("CLIENT_TIMEOUT_SECONDS").Default(5)
	//clientTimeout := 5

	envType, _ := env.String("ENV_TYPE").Default("local")

	return &AppConfig{
		Server: ServerProxyConfig{
			Port:            port,
			ShutdownTimeout: shutdownTimeout,
			EnvType:         envType,
		},
		Client: ClientProxyConfig{
			URL:     url,
			Timeout: clientTimeout,
		},
	}, nil
}
