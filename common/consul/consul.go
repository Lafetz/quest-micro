package discovery

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	consul "github.com/hashicorp/consul/api"
)

type ConsulRegistry struct {
	client *consul.Client
}

func NewConsulRegistry(addr string) (*ConsulRegistry, error) {
	config := consul.DefaultConfig()
	config.Address = addr
	client, err := consul.NewClient(config)

	if err != nil {
		return nil, err
	}
	return &ConsulRegistry{client: client}, nil
}

func (r *ConsulRegistry) Register(instanceID string, serviceName string, hostPort string) error {
	parts := strings.Split(hostPort, ":")
	if len(parts) != 2 {
		return errors.New("hostPort must be in a form of <host>:<port>, example: localhost:8081")
	}
	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}
	return r.client.Agent().ServiceRegister(&consul.AgentServiceRegistration{
		Address: parts[0],
		ID:      instanceID,
		Name:    serviceName,
		Port:    port,
		Check:   &consul.AgentServiceCheck{GRPC: hostPort, Interval: "5s", Timeout: "5s", DeregisterCriticalServiceAfter: "15s", CheckID: instanceID},
	})
}

func (r *ConsulRegistry) Deregister(instanceID string, _ string) error {
	return r.client.Agent().ServiceDeregister(instanceID)
}

func (r *ConsulRegistry) ServiceAddresses(serviceName string) ([]string, error) {
	entries, _, err := r.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	} else if len(entries) == 0 {
		return nil, ErrNotFound
	}
	var res []string

	for _, e := range entries {
		res = append(res, fmt.Sprintf("%s:%d", e.Service.Address, e.Service.Port))
	}
	return res, nil
}

func (r *ConsulRegistry) ReportHealthyState(instanceID string) error {
	return r.client.Agent().PassTTL(instanceID, "")
}
