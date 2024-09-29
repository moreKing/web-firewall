package route

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/service"
	"time"

	"server/api/route/v1"
)

func (c *ControllerV1) GetLimitPolicy(ctx context.Context, req *v1.GetLimitPolicyReq) (res *v1.GetLimitPolicyRes, err error) {
	var list []entity.ForwardLimitRules
	err = dao.ForwardLimitRules.Ctx(ctx).OrderAsc(dao.ForwardLimitRules.Columns().Position).Scan(&list)
	if err != nil {
		return nil, err
	}

	network, err := service.Network().GetNetwork()
	if err != nil {
		return nil, err
	}

	return &v1.GetLimitPolicyRes{
		Data:      list,
		Network:   *network,
		Total:     len(list),
		Timestamp: time.Now().Unix(),
	}, nil
}
