package audit

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/dao"
	"server/internal/model/entity"
	"strings"
	"time"

	"server/api/audit/v1"
)

func (c *ControllerV1) GetShellList(ctx context.Context, req *v1.GetShellListReq) (res *v1.GetShellListRes, err error) {

	if req.StartTime > 0 && req.EndTime > 0 && req.StartTime >= req.EndTime {
		return nil, errors.New("开始时间不能大于等于结束时间")
	}

	nowTime := time.Now().Unix()
	if req.StartTime >= nowTime {
		return nil, errors.New("开始时间不能大于当前时间")
	}

	g.Log().Debug(ctx, "call v1.GetUserList", req)

	// 获取登录日志
	var parments []any
	var sqls []string

	if strings.TrimSpace(req.ClientIp) != "" {
		sqls = append(sqls, fmt.Sprintf("%s LIKE ?", dao.LogShell.Columns().ClientIp))
		parments = append(parments, fmt.Sprintf("%%%s%%", strings.TrimSpace(req.ClientIp)))
	}

	if req.Online == 1 {
		sqls = append(sqls, fmt.Sprintf("%s = ?", dao.LogShell.Columns().Online))
		parments = append(parments, true)
	}
	if req.Online == 2 {
		sqls = append(sqls, fmt.Sprintf("%s = ?", dao.LogShell.Columns().Online))
		parments = append(parments, false)
	}

	if strings.TrimSpace(req.Loginname) != "" {
		// 登录名
		sqls = append(sqls, fmt.Sprintf("%s LIKE ?", dao.LogShell.Columns().Loginname))
		parments = append(parments, fmt.Sprintf("%%%s%%", req.Loginname))
	}

	if strings.TrimSpace(req.Username) != "" {
		// 用户名
		sqls = append(sqls, fmt.Sprintf("%s LIKE ?", dao.LogShell.Columns().Username))
		parments = append(parments, fmt.Sprintf("%%%s%%", req.Username))
	}

	if req.StartTime > 0 {
		sqls = append(sqls, fmt.Sprintf("%s > ?", dao.LogShell.Columns().CreatedAt))
		parments = append(parments, req.StartTime)
	}

	if req.EndTime > 0 {

		if req.EndTime > nowTime {
			req.EndTime = nowTime
		}
		sqls = append(sqls, fmt.Sprintf("%s < ?", dao.LogShell.Columns().CreatedAt))
		parments = append(parments, req.EndTime)
	}

	logs := &[]entity.LogShell{}
	var total int

	err = dao.LogShell.Ctx(ctx).OrderDesc(dao.LogShell.Columns().CreatedAt).
		Where(strings.Join(sqls, " AND "), parments...).
		Limit(req.Page*req.Limit, req.Limit).ScanAndCount(logs, &total, true)

	if err != nil {
		return nil, err
	}

	return &v1.GetShellListRes{
		Data:      logs,
		Total:     total,
		Page:      req.Page,
		Limit:     req.Limit,
		Timestamp: time.Now().Unix(),
	}, nil
}
