// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package audit

import (
	"context"

	"server/api/audit/v1"
)

type IAuditV1 interface {
	GetLoginList(ctx context.Context, req *v1.GetLoginListReq) (res *v1.GetLoginListRes, err error)
	CutLogin(ctx context.Context, req *v1.CutLoginReq) (res *v1.CutLoginRes, err error)
	GetSettingsList(ctx context.Context, req *v1.GetSettingsListReq) (res *v1.GetSettingsListRes, err error)
	GetShellList(ctx context.Context, req *v1.GetShellListReq) (res *v1.GetShellListRes, err error)
	GetShellReplayToken(ctx context.Context, req *v1.GetShellReplayTokenReq) (res *v1.GetShellReplayTokenRes, err error)
	GetShellReplay(ctx context.Context, req *v1.GetShellReplayReq) (res *v1.GetShellReplayRes, err error)
}
