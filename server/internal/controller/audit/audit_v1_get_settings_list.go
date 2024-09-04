package audit

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/dao"
	"server/internal/global"
	"server/internal/model/entity"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"server/api/audit/v1"
)

func (c *ControllerV1) GetSettingsList(ctx context.Context, req *v1.GetSettingsListReq) (res *v1.GetSettingsListRes, err error) {

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

	// 判断接口是否存在
	if (req.Method == "" && req.Path != "") || (req.Method != "" && req.Path == "") {
		return nil, gerror.NewCode(gcode.CodeValidationFailed, "method path 必须同时存在或同时不存在")
	}
	if req.Method != "" && req.Path != "" {
		pathName := global.GetPathName(req.Path, strings.ToUpper(req.Method))
		if pathName == "" {
			return nil, gerror.NewCode(gcode.CodeValidationFailed, fmt.Sprintf("不存在请求方式：%s path:%s 的接口 请检查", req.Method, req.Path))
		}

		sqls = append(sqls, fmt.Sprintf("%s = ?", dao.LogSettings.Columns().Name))
		parments = append(parments, pathName)
	}

	if req.State == 1 {
		sqls = append(sqls, fmt.Sprintf("%s = ?", dao.LogSettings.Columns().Success))
		parments = append(parments, true)
	}
	if req.State == 2 {
		sqls = append(sqls, fmt.Sprintf("%s = ?", dao.LogSettings.Columns().Success))
		parments = append(parments, false)
	}

	if strings.TrimSpace(req.Loginname) != "" {
		// 登录名
		sqls = append(sqls, fmt.Sprintf("%s LIKE ?", dao.LogSettings.Columns().Loginname))
		parments = append(parments, fmt.Sprintf("%%%s%%", req.Loginname))
	}

	if strings.TrimSpace(req.Username) != "" {
		// 用户名
		sqls = append(sqls, fmt.Sprintf("%s LIKE ?", dao.LogSettings.Columns().Username))
		parments = append(parments, fmt.Sprintf("%%%s%%", req.Username))
	}

	if req.StartTime > 0 {
		sqls = append(sqls, fmt.Sprintf("%s > ?", dao.LogSettings.Columns().CreatedAt))
		parments = append(parments, req.StartTime)
	}

	if req.EndTime > 0 {

		if req.EndTime > nowTime {
			req.EndTime = nowTime
		}
		sqls = append(sqls, fmt.Sprintf("%s < ?", dao.LogSettings.Columns().CreatedAt))
		parments = append(parments, req.EndTime)
	}

	logs := &[]entity.LogSettings{}
	var total int

	err = dao.LogSettings.Ctx(ctx).OrderDesc(dao.LogSettings.Columns().CreatedAt).
		Where(strings.Join(sqls, " AND "), parments...).
		Limit(req.Page*req.Limit, req.Limit).ScanAndCount(logs, &total, true)

	return &v1.GetSettingsListRes{
		Data:      logs,
		Total:     total,
		Page:      req.Page,
		Limit:     req.Limit,
		Timestamp: time.Now().Unix(),
	}, err

}
