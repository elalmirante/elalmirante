package providers

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/elalmirante/elalmirante-agent/conf"
	"github.com/elalmirante/elalmirante-agent/rpc"
	"github.com/elalmirante/elalmirante/models"
	"google.golang.org/grpc"
)

type Agent struct{}

func (a Agent) Deploy(s models.Server, ref string) (string, error) {

	parts := strings.Split(s.Key, "@")
	key := parts[0]
	host := parts[1]

	// if there is no port then add default port.
	if !strings.Contains(host, ":") {
		host = fmt.Sprintf("%s:%s", host, conf.DefaultPort)
	}

	// dial connection and call client
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return "", err
	}

	client := rpc.NewDeployServiceClient(conn)
	response, err := client.Deploy(context.Background(), &rpc.DeployRequest{
		Key: key,
		Ref: ref,
	})

	if err != nil {
		return "", err
	}

	// fill error from response if not blank
	if response.Error != "" {
		err = errors.New(response.Error)
	}

	return response.Output, err
}

func (a Agent) KeyFormat() string {
	return "<key>@<host>[:<port>]"
}

func (a Agent) ValidKey(key string) bool {

	parts := strings.Split(key, "@")

	if len(parts) != 2 {
		return false
	}

	host := parts[1]

	if strings.Contains(host, ":") {
		return len(strings.Split(host, ":")) == 2
	}

	return true
}
