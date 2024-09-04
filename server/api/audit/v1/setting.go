package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model/entity"
)

type GetSettingsListReq struct {
	g.Meta `path:"/settings"  tags:"审计日志" method:"get" summary:"操作日志"`
	Page   int `json:"page" v:"min:0"`
	Limit  int `json:"limit" v:"required|min:10|max:2000"`

	Username  string `json:"username"`
	Loginname string `json:"loginname"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
	State     int    `json:"state" v:"in:0,1,2#传入的参数 0.全部 1.成功 2.失败 中的一个" dc:" 0.全部 1.成功 2.失败"` //0.全部 1.启用 2.有效期 3.禁用

	Method string `json:"method"`
	Path   string `json:"path"`
}
type GetSettingsListRes struct {
	Data      *[]entity.LogSettings `json:"data"`
	Total     int                   `json:"total"`
	Page      int                   `json:"page"`
	Limit     int                   `json:"limit"`
	Timestamp int64                 `json:"timestamp"`
}
