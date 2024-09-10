package nft

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/do"
	"server/internal/service"
	"server/utility/nftables"
	"strconv"
	"strings"
)

type sNft map[int]*[]model.Rulesets

func init() {
	service.RegisterNft(New())
}

func New() service.INft {
	// 1-7 分别获取
	var s sNft = make(map[int]*[]model.Rulesets)

	for i := 1; i < 10; i++ {
		var tmp []model.Rulesets
		// 注意使用position升序排列
		err := dao.Rulesets.Ctx(context.Background()).Where(dao.Rulesets.Columns().Chain, i).OrderAsc(dao.Rulesets.Columns().Position).Scan(&tmp)
		if err != nil {
			panic(err)
		}

		// 添加策略到系统
		for j, _ := range tmp {
			rule, err := nftables.AddRule(context.Background(), &nftables.Rule{
				Chain: nftables.ChainName[i],
				Add:   true,
				Expr:  tmp[j].Expr,
			})
			if err != nil {
				panic(err)
			}
			tmp[j].Handle = rule.Handle
		}
		g.Log().Debug(context.Background(), tmp)
		s[i] = &tmp
	}

	return &s

}

func (s *sNft) Add(ctx context.Context, rulesets *model.Rulesets, add bool) error {

	err := dao.Rulesets.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 添加到数据库中
		id, err := tx.Ctx(ctx).Model(&do.Rulesets{}).InsertAndGetId(&do.Rulesets{
			Comment:  rulesets.Comment,
			Chain:    rulesets.Chain,
			Position: -1,
			Expr:     rulesets.Expr,
		})
		if err != nil {
			return err
		}

		rulesets.Id = id

		// 确定是否在开头或者末尾
		if rulesets.Position > 0 {
			isAdd := false
			for _, r := range *((*s)[rulesets.Chain]) {
				if r.Id == int64(rulesets.Position) {
					isAdd = true
					position, err := strconv.Atoi(strings.TrimSpace(r.Handle))
					if err != nil {
						return err
					}
					tmpR, err := nftables.AddRule(ctx, &nftables.Rule{
						Chain:    nftables.ChainName[rulesets.Chain],
						Handle:   "",
						Add:      add,
						Position: position,
						Expr:     rulesets.Expr,
					})
					if err != nil {
						return err
					}

					rulesets.Handle = tmpR.Handle
					break
				}
			}
			if !isAdd {
				return errors.New("指定位置不存在规则")
			}
		} else {
			// position 不能重复，每次更新必须全量更新避免删除导致位置重复
			tmpR, err := nftables.AddRule(ctx, &nftables.Rule{
				Chain:    nftables.ChainName[rulesets.Chain],
				Handle:   "",
				Add:      add,
				Position: 0,
				Expr:     rulesets.Expr,
			})
			if err != nil {
				return err
			}

			rulesets.Handle = tmpR.Handle
		}

		// 查找对应位置进行数据插入
		err = s.add(rulesets, add)
		if err != nil {
			return err
		}

		// 刷新数据库position
		for i, rule := range *((*s)[rulesets.Chain]) {
			rule.Position = i
			_, err := tx.Ctx(ctx).Model(&do.Rulesets{}).Where(dao.Rulesets.Columns().Id, rule.Id).Update(&do.Rulesets{
				Position: i,
			})
			if err != nil {
				return err
			}
		}

		g.Log().Debug(ctx, "更新完链中的规则： ", *((*s)[rulesets.Chain]))

		return nil
	})

	return err
}

// func (s *sNft) Replace() {
//
// }
//
// func (s *sNft) ChangePosition() {
//
// }
//
// func (s *sNft) Delete() {
//
// }
func (s *sNft) GetChainList(chain int) *[]model.Rulesets {
	if chain < 1 || chain > 9 {
		g.Log().Error(context.Background(), "不存在的链id", chain)
		return nil
	}
	return (*s)[chain]
}

func (s *sNft) add(r *model.Rulesets, add bool) error {

	// 末尾新增
	if r.Position == 0 {
		if add {
			*((*s)[r.Chain]) = append(*((*s)[r.Chain]), *r)
			return nil
		}

		// 开头新增
		newSlice := make([]model.Rulesets, len(*((*s)[r.Chain]))+1)
		newSlice[0] = *r
		copy(newSlice[1:], *((*s)[r.Chain]))
		(*s)[r.Chain] = &newSlice
		return nil

	}

	// 创建新的切片，长度为原有切片长度加1
	newSlice := make([]model.Rulesets, len(*((*s)[r.Chain]))+1)
	postion := -1

	for i, tmp := range *((*s)[r.Chain]) {
		if int64(r.Position) == tmp.Id {
			postion = i
			break
		}
	}
	if postion == -1 {
		return errors.New(fmt.Sprintf("链 %s 没有规则id : %d ", nftables.ChainName[r.Chain], r.Position))
	}
	g.Log().Debug(context.Background(), fmt.Sprintf("r.Position: %d, position: %d ", r.Position, postion))

	if add {
		// 指定元素之后，否则就是之前
		postion += 1
	}
	// 将前半部分复制自原有切片
	copy(newSlice[:postion], (*((*s)[r.Chain]))[:postion])
	// 插入要添加的元素
	newSlice[postion] = *r
	// 将后半部分复制自原有切片
	copy(newSlice[postion+1:], (*((*s)[r.Chain]))[postion:])
	(*s)[r.Chain] = &newSlice

	return nil
}

func (s *sNft) Delete(ctx context.Context, id int64) error {
	var rule model.Rulesets
	err := dao.Rulesets.Ctx(ctx).Where(dao.Rulesets.Columns().Id, id).Scan(&rule)
	if err != nil {
		return err
	}

	err = dao.Rulesets.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 查找元素

		var position = -1
		for i, rulesets := range *(*s)[rule.Chain] {
			if rulesets.Id == id {
				//	 删除系统中的策略
				err := nftables.DeleteRule(&nftables.Rule{
					Chain:  nftables.ChainName[rulesets.Chain],
					Handle: rulesets.Handle,
				})
				if err != nil {
					return err
				}

				position = i
				break
			}
		}
		if position == -1 {
			return errors.New("无匹配的规则")
		}

		// 删除策略
		_, err := dao.Rulesets.Ctx(ctx).Where(dao.Rulesets.Columns().Id, id).Delete()
		if err != nil {
			return err
		}

		//	 删除数组中的元素
		*(*s)[rule.Chain] = append((*(*s)[rule.Chain])[:position], (*(*s)[rule.Chain])[position+1:]...)

		// 刷新数据库
		err = s.flushPositionDB(ctx, rule.Chain, tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *sNft) flushPositionDB(ctx context.Context, chain int, tx gdb.TX) error {

	// 刷新数据库position
	for i, rule := range *((*s)[chain]) {
		rule.Position = i
		_, err := tx.Ctx(ctx).Model(&do.Rulesets{}).Where(dao.Rulesets.Columns().Id, rule.Id).Update(&do.Rulesets{
			Position: i,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *sNft) flushDaoPositionDB(ctx context.Context, chain int) error {

	// 刷新数据库position
	for i, rule := range *((*s)[chain]) {
		rule.Position = i
		_, err := dao.Rulesets.Ctx(ctx).Where(dao.Rulesets.Columns().Id, rule.Id).Update(&do.Rulesets{
			Position: i,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *sNft) Replace(ctx context.Context, rule *model.Rulesets) error {
	// 获取数组中的策略
	isUpdate := false
	for i, rulesets := range *(*s)[rule.Chain] {
		if rulesets.Id == rule.Id {
			//	 由于nftables替换存在bug，各Linux表现不一致，这里选择删除重建
			//err := nftables.ReplaceRule(&nftables.Rule{
			//	Chain:  nftables.ChainName[rule.Chain],
			//	Handle: rulesets.Handle,
			//	Expr:   rule.Expr,
			//})
			//if err != nil {
			//	return err
			//}

			isUpdate = true
			// 新建
			position, err := strconv.Atoi(strings.TrimSpace(rulesets.Handle))
			if err != nil {
				return err
			}

			addRule, err := nftables.AddRule(ctx, &nftables.Rule{
				Chain:    nftables.ChainName[rulesets.Chain],
				Handle:   rulesets.Handle,
				Add:      false,
				Position: position,
				Expr:     rule.Expr,
			})
			if err != nil {
				return err
			}

			// 删除原规则
			_ = nftables.DeleteRule(&nftables.Rule{
				Chain:  nftables.ChainName[rule.Chain],
				Handle: rulesets.Handle,
			})

			//rule.Handle = rulesets.Handle
			rule.Handle = addRule.Handle
			(*(*s)[rule.Chain])[i] = *rule
			break
		}
	}

	if !isUpdate {
		return errors.New("未匹配的规则策略")
	}

	// 更新数据库
	_, err := dao.Rulesets.Ctx(ctx).Where(dao.Rulesets.Columns().Id, rule.Id).Update(&do.Rulesets{
		Comment: rule.Comment,
		Expr:    rule.Expr,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *sNft) UpdatePosition(ctx context.Context, rid, did int64, add bool) error {

	var rule model.Rulesets
	err := dao.Rulesets.Ctx(ctx).Where(dao.Rulesets.Columns().Id, rid).Scan(&rule)
	if err != nil {
		return err
	}

	//	原位置
	oldPosition := -1
	for i := 0; i < len(*(*s)[rule.Chain]); i++ {
		if (*(*s)[rule.Chain])[i].Id == rid {
			oldPosition = i
			rule = (*(*s)[rule.Chain])[i]
			break
		}
	}

	if oldPosition == -1 {
		return errors.New("没有符合条件的策略")
	}

	// 目标位置
	destPosition := -1
	if did == 0 {
		destPosition = 0
	} else {
		for i := 0; i < len(*(*s)[rule.Chain]); i++ {
			if (*(*s)[rule.Chain])[i].Id == did {
				destPosition = i
				break
			}
		}
		if destPosition == -1 {
			return errors.New("没有符合条件的目标位置")
		}
	}

	// 删除原策略
	_ = nftables.DeleteRule(&nftables.Rule{
		Chain:  nftables.ChainName[rule.Chain],
		Handle: rule.Handle,
	})

	// 插入新的策略
	position := 0
	if did != 0 {
		position, err = strconv.Atoi((*(*s)[rule.Chain])[destPosition].Handle)
		if err != nil {
			return err
		}
	}

	addRule, err := nftables.AddRule(ctx, &nftables.Rule{
		Chain:    nftables.ChainName[rule.Chain],
		Add:      add,
		Position: position,
		Expr:     rule.Expr,
	})

	if err != nil {
		return err
	}
	rule.Handle = addRule.Handle

	// 删除数组中原数据
	for i := 0; i < len(*(*s)[rule.Chain]); i++ {
		if (*(*s)[rule.Chain])[i].Id == rid {
			if i == 0 {
				*(*s)[rule.Chain] = (*(*s)[rule.Chain])[i+1:]
				break
			}
			if i == len(*(*s)[rule.Chain])-1 {
				*(*s)[rule.Chain] = (*(*s)[rule.Chain])[:i]
				break
			}

			*(*s)[rule.Chain] = append((*(*s)[rule.Chain])[:i], (*(*s)[rule.Chain])[i+1:]...)

			break
		}
	}

	// 插入到新的位置
	if did == 0 {
		if add {
			*(*s)[rule.Chain] = append(*(*s)[rule.Chain], rule)
		} else {
			*(*s)[rule.Chain] = append([]model.Rulesets{rule}, *(*s)[rule.Chain]...)
		}
	} else {
		for i := 0; i < len(*(*s)[rule.Chain]); i++ {
			if (*(*s)[rule.Chain])[i].Id == did {
				if add && i == len(*(*s)[rule.Chain])-1 {
					*(*s)[rule.Chain] = append(*(*s)[rule.Chain], rule)
					break
				}
				if !add && i == 0 {
					*(*s)[rule.Chain] = append([]model.Rulesets{rule}, *(*s)[rule.Chain]...)
					break
				}

				for _, tmp := range (*(*s)[rule.Chain])[:i] {
					g.Log().Debug(ctx, tmp)
				}

				if add {
					// 指定元素之后，否则就是之前
					i += 1
				}
				newSlice := make([]model.Rulesets, len(*(*s)[rule.Chain])+1)
				// 将前半部分复制自原有切片
				copy(newSlice[:i], (*(*s)[rule.Chain])[:i])
				// 插入要添加的元素
				newSlice[i] = rule
				// 将后半部分复制自原有切片
				copy(newSlice[i+1:], (*((*s)[rule.Chain]))[i:])
				(*s)[rule.Chain] = &newSlice

				break
			}
		}

	}

	err = s.flushDaoPositionDB(ctx, rule.Chain)
	if err != nil {
		return err
	}

	return nil
}
