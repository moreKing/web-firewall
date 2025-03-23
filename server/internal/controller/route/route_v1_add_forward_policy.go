package route

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/dao"
	"server/internal/model/do"
	"server/internal/model/entity"
	"server/internal/service"
	"time"

	"server/api/route/v1"
)

func (c *ControllerV1) AddForwardPolicy(ctx context.Context, req *v1.AddForwardPolicyReq) (res *v1.AddForwardPolicyRes, err error) {

	var id int64 = 0
	err = dao.ForwardRules.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 获取全部策略进行排序
		var list []entity.ForwardRules
		err := tx.Ctx(ctx).Model(&do.ForwardRules{}).OrderAsc(dao.ForwardRules.Columns().Position).Scan(&list)
		if err != nil {
			return err
		}
		// 添加策略
		id, err = tx.Ctx(ctx).Model(&do.ForwardRules{}).InsertAndGetId(&do.ForwardRules{

			Protocol: req.Protocol,
			Sip:      req.Sip,
			Dip:      req.Dip,
			Port:     req.Port,
			Comment:  req.Comment,
			Policy:   req.Policy,

			Position:  -1,
			CreatedAt: time.Now().Unix(),
			DeletedAt: nil,
		})
		if err != nil {
			return err
		}

		// 刷新策略
		if req.Position == 0 {
			if req.Add {
				list = append(list, entity.ForwardRules{Id: id})
			} else {
				newSlice := make([]entity.ForwardRules, len(list)+1)
				newSlice[0] = entity.ForwardRules{Id: id}
				copy(newSlice[1:], list)
				list = newSlice
			}
		} else {
			for i, item := range list {
				if item.Id == int64(req.Position) {
					if req.Add && i == len(list)-1 {
						list = append(list, entity.ForwardRules{Id: id})
					} else if !req.Add && i == 0 {
						newSlice := make([]entity.ForwardRules, len(list)+1)
						newSlice[0] = entity.ForwardRules{Id: id}
						copy(newSlice[1:], list)
						list = newSlice
					} else {
						newSlice := make([]entity.ForwardRules, len(list)+1)
						position := i
						if req.Add {
							// 指定元素之后，否则就是之前
							position += 1
						}
						// 将前半部分复制自原有切片
						copy(newSlice[:position], list[:position])
						// 插入要添加的元素
						newSlice[position] = entity.ForwardRules{Id: id}
						// 将后半部分复制自原有切片
						copy(newSlice[position+1:], list[position:])
						list = newSlice
					}

					break
				}
			}

		}

		for i, rules := range list {
			_, err := tx.Ctx(ctx).Model(&do.ForwardRules{}).Where(dao.ForwardRules.Columns().Id, rules.Id).Update(&do.ForwardRules{
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

	return &v1.AddForwardPolicyRes{Id: id}, nil
}
