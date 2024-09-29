package network

import (
	"net"
	"server/internal/model"
	"server/internal/service"
)

type sNetwork struct {
}

func init() {
	service.RegisterNetwork(&sNetwork{})
}

func (s *sNetwork) GetNetwork() (*[]model.Network, error) {
	var network []model.Network

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, stat := range interfaces {
		network = append(network, model.Network{
			Index: stat.Index,
			Name:  stat.Name,
		})

	}

	return &network, nil
}

func (s *sNetwork) IsNetwork(interfaceName string) (ok bool, err error) {

	network, err := s.GetNetwork()
	if err != nil {
		return false, err
	}

	for _, n := range *network {
		if n.Name == interfaceName {
			return true, nil
		}
	}

	return false, nil
}
