package iotmaker_docker_builder_network_interface

import (
	"github.com/docker/docker/api/types/network"
)

type ContainerBuilderNetworkInterface interface {
	Init() (err error)
	GetConfiguration() (networkConfiguration *network.NetworkingConfig, err error)
	NetworkCreate(name, subnet, gateway string) (err error)
	GetNetworkName() (name string)
	Remove() (err error)
}
