package nft

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/shirou/gopsutil/v3/net"
	"server/internal/model"
	"server/internal/service"
	"time"

	"server/api/nft/v1"
)

func (c *ControllerV1) GetPolicy(ctx context.Context, req *v1.GetPolicyReq) (res *v1.GetPolicyRes, err error) {

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
		g.Log().Debug(ctx, stat)
	}

	return &v1.GetPolicyRes{
		Data:      *(service.Nft().GetChainList(req.Chain)),
		Network:   network,
		Total:     len(*(service.Nft().GetChainList(req.Chain))),
		Timestamp: time.Now().Unix(),
	}, nil
}
