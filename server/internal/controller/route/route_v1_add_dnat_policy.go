package route

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/dao"
	"server/internal/model/do"
	"server/internal/model/entity"
	"server/internal/service"
	"time"

	"server/api/route/v1"
)

func (c *ControllerV1) AddDnatPolicy(ctx context.Context, req *v1.AddDnatPolicyReq) (res *v1.AddDnatPolicyRes, err error) {

	if len(req.Port) < 1 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "端口映射不能为空")
	}
	// 判断出网卡是否存在
	network, err := service.Network().GetNetwork()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	for _, n := range *network {
		if req.Iif == n.Name {
			goto CONTINUE
		}
	}

	g.Log().Error(ctx, "不存在的网卡名称", req.Iif)
	return nil, errors.New(fmt.Sprintf("不存在的网卡名称 %s", req.Iif))

CONTINUE:

	var id int64 = 0
	err = dao.DnatRules.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 获取全部策略进行排序
		var list []entity.DnatRules
		err := tx.Ctx(ctx).Model(&do.DnatRules{}).OrderAsc(dao.DnatRules.Columns().Position).Scan(&list)
		if err != nil {
			return err
		}
		// 添加策略
		g.Log().Debug(ctx, "port----->  ", req.Port)
		id, err = tx.Ctx(ctx).Model(&do.DnatRules{}).InsertAndGetId(&do.DnatRules{

			Protocol:  req.Protocol,
			Dip:       req.Dip,
			Iif:       req.Iif,
			Port:      req.Port,
			Dnat:      req.Dnat,
			Comment:   req.Comment,
			Position:  -1,
			CreatedAt: time.Now().Unix(),
		})
		if err != nil {
			return err
		}

		// 刷新策略
		if req.Position == 0 {
			if req.Add {
				list = append(list, entity.DnatRules{Id: id})
			} else {
				newSlice := make([]entity.DnatRules, len(list)+1)
				newSlice[0] = entity.DnatRules{Id: id}
				copy(newSlice[1:], list)
				list = newSlice
			}
		} else {
			for i, item := range list {
				if item.Id == int64(req.Position) {
					if req.Add && i == len(list)-1 {
						list = append(list, entity.DnatRules{Id: id})
					} else if !req.Add && i == 0 {
						newSlice := make([]entity.DnatRules, len(list)+1)
						newSlice[0] = entity.DnatRules{Id: id}
						copy(newSlice[1:], list)
						list = newSlice
					} else {
						newSlice := make([]entity.DnatRules, len(list)+1)
						position := i
						if req.Add {
							// 指定元素之后，否则就是之前
							position += 1
						}
						// 将前半部分复制自原有切片
						copy(newSlice[:position], list[:position])
						// 插入要添加的元素
						newSlice[position] = entity.DnatRules{Id: id}
						// 将后半部分复制自原有切片
						copy(newSlice[position+1:], list[position:])
						list = newSlice
					}

					break
				}
			}

		}

		for i, rules := range list {
			_, err := tx.Ctx(ctx).Model(&do.DnatRules{}).Where(dao.DnatRules.Columns().Id, rules.Id).Update(&do.DnatRules{
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

	return &v1.AddDnatPolicyRes{Id: id}, nil
}
