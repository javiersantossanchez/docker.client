package docker

import (
	"net/http"
	"time"

	"github.com/tv42/httpunix"
)

/*
class docker connector
*/
type DockerConnector struct {
}

func (docker *DockerConnector) socketPath() string {
	return "/var/run/docker.sock"
}

func (docker *DockerConnector) BaseUrl() string {
	return "http+unix://myservice//v1.23"
}

/*
* Get client to connect
 */
func (docker *DockerConnector) GetConnector() http.Client {

	u := &httpunix.Transport{
		DialTimeout:           100 * time.Millisecond,
		RequestTimeout:        1 * time.Second,
		ResponseHeaderTimeout: 1 * time.Second,
	}
	u.RegisterLocation("myservice", docker.socketPath())

	// If you want to use http: with the same client:
	t := &http.Transport{}
	t.RegisterProtocol(httpunix.Scheme, u)
	return http.Client{
		Transport: t,
	}

}
