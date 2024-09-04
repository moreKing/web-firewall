package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model/entity"
)

type GetLoginListReq struct {
	g.Meta    `path:"/login"  tags:"审计日志" method:"get" summary:"登录日志"`
	Page      int    `json:"page" v:"min:0"`
	Limit     int    `json:"limit" v:"required|min:10|max:2000"`
	Username  string `json:"username"`
	Loginname string `json:"loginname"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
	Online    int    `json:"online"`
	State     int    `json:"state" v:"in:0,1,2#传入的参数 0.全部 1.成功 2.失败 中的一个" dc:" 0.全部 1.成功 2.失败"` //0.全部 1.启用 2.有效期 3.禁用
}
type GetLoginListRes struct {
	Data      *[]entity.LogLogins `json:"data"`
	Total     int                 `json:"total"`
	Page      int                 `json:"page"`
	Limit     int                 `json:"limit"`
	Timestamp int64               `json:"timestamp"`
}

type CutLoginReq struct {
	g.Meta `path:"/cut-login/{uuid}"  tags:"审计日志" method:"post" summary:"切断登录"`
	UUID   string `json:"uuid" path:"uuid"`
}
type CutLoginRes struct {
}
