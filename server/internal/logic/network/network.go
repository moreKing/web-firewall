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
		if stat.Name == "lo" {
			continue
		}
		addrs, err := stat.Addrs()
		if err != nil {
			return nil, err
		}

		ips := []string{}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ips = append(ips, ipnet.IP.To4().String())
				}
			}

		}

		network = append(network, model.Network{
			Index: stat.Index,
			Name:  stat.Name,
			Ip:    ips,
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
