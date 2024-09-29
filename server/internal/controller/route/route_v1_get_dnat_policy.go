package route

import (
	"context"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/service"
	"time"

	"server/api/route/v1"
)

func (c *ControllerV1) GetDnatPolicy(ctx context.Context, req *v1.GetDnatPolicyReq) (res *v1.GetDnatPolicyRes, err error) {

	var list []model.DnatRulesets
	err = dao.DnatRules.Ctx(ctx).OrderAsc(dao.DnatRules.Columns().Position).Scan(&list)
	if err != nil {
		return nil, err
	}

	network, err := service.Network().GetNetwork()
	if err != nil {
		return nil, err
	}

	return &v1.GetDnatPolicyRes{
		Data:      list,
		Network:   *network,
		Total:     len(list),
		Timestamp: time.Now().Unix(),
	}, nil
}
