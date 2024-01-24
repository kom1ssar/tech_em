package env

import (
	"errors"
	"github.com/kom1ssar/tech_em/internal/config"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	hostEnvName            = "HTTP_HOST"
	portEnvName            = "HTTP_PORT"
	httpTimeoutEnvName     = "HTTP_TIMEOUT"
	httpIdleTimeoutEnvName = "HTTP_IDLE_TIMEOUT"
)

var _ config.HTTPServerConfig = (*httpConfig)(nil)

type httpConfig struct {
	host        string
	port        int
	timeout     time.Duration
	idleTimeout time.Duration
}

func NewHTTPServerConfig() (config.HTTPServerConfig, error) {
	host := os.Getenv(hostEnvName)
	if len(host) == 0 {
		return nil, errors.New("http_host env not found")
	}

	portString := os.Getenv(portEnvName)
	if len(portString) == 0 {
		return nil, errors.New("http_port env not found")
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		return nil, errors.New("http_port env parse to int err")
	}

	//timeout := os.Getenv(httpTimeoutEnvName)
	//if len(timeout) == 0 {
	//	return nil, errors.New("http_timeout env not found")
	//}
	//
	//idleTimeout := os.Getenv(httpIdleTimeoutEnvName)
	//if len(idleTimeout) == 0 {
	//	return nil, errors.New("http_idle_timeout env not found")
	//}

	return &httpConfig{
		host:        host,
		port:        port,
		timeout:     4 * time.Second, //todo
		idleTimeout: 30 * time.Second,
	}, nil
}

func (h *httpConfig) Address() string {
	return net.JoinHostPort(h.host, strconv.Itoa(h.port))
}

func (h *httpConfig) GetTimeout() time.Duration {
	return h.timeout
}

func (h *httpConfig) GetIdleTimeout() time.Duration {
	return h.idleTimeout
}
