package route

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/service"
	"time"

	"server/api/route/v1"
)

func (c *ControllerV1) GetSnatPolicy(ctx context.Context, req *v1.GetSnatPolicyReq) (res *v1.GetSnatPolicyRes, err error) {

	var list []entity.SnatRules
	err = dao.SnatRules.Ctx(ctx).OrderAsc(dao.SnatRules.Columns().Position).Scan(&list)
	if err != nil {
		return nil, err
	}

	network, err := service.Network().GetNetwork()
	if err != nil {
		return nil, err
	}

	return &v1.GetSnatPolicyRes{
		Data:      list,
		Network:   *network,
		Total:     len(list),
		Timestamp: time.Now().Unix(),
	}, nil
}
