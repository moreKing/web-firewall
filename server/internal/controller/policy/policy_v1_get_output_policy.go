package policy

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/service"
	"time"

	"server/api/policy/v1"
)

func (c *ControllerV1) GetOutputPolicy(ctx context.Context, req *v1.GetOutputPolicyReq) (res *v1.GetOutputPolicyRes, err error) {

	var list []entity.OutputRules
	err = dao.OutputRules.Ctx(ctx).OrderAsc(dao.OutputRules.Columns().Position).Scan(&list)
	if err != nil {
		return nil, err
	}

	network, err := service.Network().GetNetwork()
	if err != nil {
		return nil, err
	}

	return &v1.GetOutputPolicyRes{
		Data:      list,
		Network:   *network,
		Total:     len(list),
		Timestamp: time.Now().Unix(),
	}, nil
}
