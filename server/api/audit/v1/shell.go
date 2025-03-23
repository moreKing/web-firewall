package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model/entity"
)

type GetShellListReq struct {
	g.Meta    `path:"/shell"  tags:"审计日志" method:"get" summary:"shell日志"`
	Page      int    `json:"page" v:"min:0"`
	Limit     int    `json:"limit" v:"required|min:10|max:2000"`
	Username  string `json:"username"`
	Loginname string `json:"loginname"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
	Online    int    `json:"online"`
	ClientIp  string `json:"clientIp"`
}
type GetShellListRes struct {
	Data      *[]entity.LogShell `json:"data"`
	Total     int                `json:"total"`
	Page      int                `json:"page"`
	Limit     int                `json:"limit"`
	Timestamp int64              `json:"timestamp"`
}

type GetShellReplayTokenReq struct {
	g.Meta `path:"/shell-token"  tags:"审计日志" method:"get" summary:"shell回放token"`
}
type GetShellReplayTokenRes struct {
	Token string `json:"token"`
}

type GetShellReplayReq struct {
	g.Meta `path:"/shell-replay"  tags:"审计日志" method:"get" summary:"shell回放"`
	Token  string `json:"token" v:"required"`
	Id     string `json:"id" v:"required" dc:"回放id"`
}

type GetShellReplayRes struct {
}
