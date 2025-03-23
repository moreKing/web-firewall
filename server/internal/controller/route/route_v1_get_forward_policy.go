package route

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/service"
	"time"

	"server/api/route/v1"
)

func (c *ControllerV1) GetForwardPolicy(ctx context.Context, req *v1.GetForwardPolicyReq) (res *v1.GetForwardPolicyRes, err error) {
	var list []entity.ForwardRules
	err = dao.ForwardRules.Ctx(ctx).OrderAsc(dao.ForwardRules.Columns().Position).Scan(&list)
	if err != nil {
		return nil, err
	}

	network, err := service.Network().GetNetwork()
	if err != nil {
		return nil, err
	}

	return &v1.GetForwardPolicyRes{
		Data:      list,
		Network:   *network,
		Total:     len(list),
		Timestamp: time.Now().Unix(),
	}, nil
}
