package policy

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/service"
	"time"

	"server/api/policy/v1"
)

func (c *ControllerV1) GetInputPolicy(ctx context.Context, req *v1.GetInputPolicyReq) (res *v1.GetInputPolicyRes, err error) {

	var list []entity.InputRules
	err = dao.InputRules.Ctx(ctx).OrderAsc(dao.InputRules.Columns().Position).Scan(&list)
	if err != nil {
		return nil, err
	}

	network, err := service.Network().GetNetwork()
	if err != nil {
		return nil, err
	}

	return &v1.GetInputPolicyRes{
		Data:      list,
		Network:   *network,
		Total:     len(list),
		Timestamp: time.Now().Unix(),
	}, nil
}
