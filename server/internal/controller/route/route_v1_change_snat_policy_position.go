package route

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/dao"
	"server/internal/model/do"
	"server/internal/model/entity"
	"server/internal/service"

	"server/api/route/v1"
)

func (c *ControllerV1) ChangeSnatPolicyPosition(ctx context.Context, req *v1.ChangeSnatPolicyPositionReq) (res *v1.ChangeSnatPolicyPositionRes, err error) {

	// 获取所有的规则
	var list []entity.SnatRules
	err = dao.SnatRules.Ctx(ctx).OrderAsc(dao.SnatRules.Columns().Position).Scan(&list)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 找到对应的规则

	var tmp entity.SnatRules
	for i, rules := range list {
		if rules.Id == req.ID {
			tmp = rules
			if i == 0 {
				list = list[i+1:]
			} else if len(list)-1 == i {
				list = list[:i]
			} else {
				list = append(list[:i], list[i+1:]...)
			}

			goto Position
		}
	}

	return nil, errors.New("不存在的id")

Position:
	var des = 0
	if req.Position == 0 {
		goto OK
	}

	for j, dest := range list {
		if dest.Id == req.Position {
			des = j
			goto OK
		}
	}

	return nil, errors.New("不存在的position")
OK:
	// 重新排序

	if req.Position == 0 {
		if req.Add {
			list = append(list, tmp)
		} else {
			newSlice := make([]entity.SnatRules, len(list)+1)
			newSlice[0] = tmp
			copy(newSlice[1:], list)
			list = newSlice
		}
	} else {
		if req.Add && des == len(list)-1 {
			list = append(list, tmp)
		} else if !req.Add && des == 0 {
			newSlice := make([]entity.SnatRules, len(list)+1)
			newSlice[0] = tmp
			copy(newSlice[1:], list)
			list = newSlice
		} else {
			newSlice := make([]entity.SnatRules, len(list)+1)
			position := des
			if req.Add {
				// 指定元素之后，否则就是之前
				position += 1
			}
			// 将前半部分复制自原有切片
			copy(newSlice[:position], list[:position])
			// 插入要添加的元素
			newSlice[position] = tmp
			// 将后半部分复制自原有切片
			copy(newSlice[position+1:], list[position:])
			list = newSlice
		}

	}

	// 将排序的结果更新到数据库
	err = dao.SnatRules.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		for i, rules := range list {
			_, err := tx.Ctx(ctx).Model(&do.SnatRules{}).Where(dao.SnatRules.Columns().Id, rules.Id).Update(&do.SnatRules{
				Position: i,
			})
			if err != nil {
				return err
			}
		}

		// 更新防火墙
		return service.Policy().Flush(ctx)

	})

	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return nil, nil
}
