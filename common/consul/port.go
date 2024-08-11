package discovery

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var ErrNotFound = errors.New("no service addresses found")

type RegistryApi interface {
	Register(instanceID string, serviceName string, hostPort string) error

	Deregister(instanceID string, serviceName string) error

	ServiceAddresses(serviceID string) ([]string, error)

	ReportHealthyState(instanceID string) error
}

func GenerateInstanceID(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}
