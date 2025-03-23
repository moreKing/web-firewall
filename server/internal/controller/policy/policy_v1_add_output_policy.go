package policy

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/dao"
	"server/internal/model/do"
	"server/internal/model/entity"
	"server/internal/service"
	"time"

	"server/api/policy/v1"
)

func (c *ControllerV1) AddOutputPolicy(ctx context.Context, req *v1.AddOutputPolicyReq) (res *v1.AddOutputPolicyRes, err error) {

	var id int64 = 0
	err = dao.OutputRules.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		// 获取全部策略进行排序
		var list []entity.OutputRules
		err := tx.Ctx(ctx).Model(&do.OutputRules{}).OrderAsc(dao.OutputRules.Columns().Position).Scan(&list)
		if err != nil {
			return err
		}
		// 添加策略
		id, err = tx.Ctx(ctx).Model(&do.OutputRules{}).InsertAndGetId(&do.OutputRules{
			Protocol:  req.Protocol,
			Port:      req.Port,
			Ip:        req.Ip,
			Ct:        req.Ct,
			Icmp:      req.Icmp,
			Comment:   req.Comment,
			Position:  -1,
			Policy:    req.Policy,
			CreatedAt: time.Now().Unix(),
		})
		if err != nil {
			return err
		}

		// 刷新策略
		if req.Position == 0 {
			if req.Add {
				list = append(list, entity.OutputRules{Id: id})
			} else {
				newSlice := make([]entity.OutputRules, len(list)+1)
				newSlice[0] = entity.OutputRules{Id: id}
				copy(newSlice[1:], list)
				list = newSlice
			}
		} else {

			for i, item := range list {
				if item.Id == int64(req.Position) {
					if req.Add && i == len(list)-1 {
						list = append(list, entity.OutputRules{Id: id})
					} else if !req.Add && i == 0 {
						newSlice := make([]entity.OutputRules, len(list)+1)
						newSlice[0] = entity.OutputRules{Id: id}
						copy(newSlice[1:], list)
						list = newSlice
					} else {
						newSlice := make([]entity.OutputRules, len(list)+1)
						position := i
						if req.Add {
							// 指定元素之后，否则就是之前
							position += 1
						}
						// 将前半部分复制自原有切片
						copy(newSlice[:position], list[:position])
						// 插入要添加的元素
						newSlice[position] = entity.OutputRules{Id: id}
						// 将后半部分复制自原有切片
						copy(newSlice[position+1:], list[position:])
						list = newSlice
					}

					break
				}
			}

		}

		for i, rules := range list {
			_, err := tx.Ctx(ctx).Model(&do.OutputRules{}).Where(dao.OutputRules.Columns().Id, rules.Id).Update(&do.OutputRules{
				Position: i,
			})
			if err != nil {
				return err
			}
		}

		return service.Policy().Flush(ctx)
	})

	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return &v1.AddOutputPolicyRes{Id: id}, nil
}
