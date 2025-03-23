package policy

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/service"
	"time"

	"server/api/policy/v1"
)

func (c *ControllerV1) GetInputLimitPolicy(ctx context.Context, req *v1.GetInputLimitPolicyReq) (res *v1.GetInputLimitPolicyRes, err error) {
	var list []entity.InputLimitRules
	err = dao.InputLimitRules.Ctx(ctx).OrderAsc(dao.InputLimitRules.Columns().Position).Scan(&list)
	if err != nil {
		return nil, err
	}

	network, err := service.Network().GetNetwork()
	if err != nil {
		return nil, err
	}

	return &v1.GetInputLimitPolicyRes{
		Data:      list,
		Network:   *network,
		Total:     len(list),
		Timestamp: time.Now().Unix(),
	}, nil
}
