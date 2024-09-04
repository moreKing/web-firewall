package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"server/internal/dao"
	"server/internal/global"
	"server/internal/model"
	"server/internal/model/do"
	"server/internal/service"
	"slices"
	"strings"
)

func (s *sMiddleware) ConfigLog(r *ghttp.Request) {
	//reqBody := r.GetBodyString()
	r.Middleware.Next()
	//g.Log().Debug(r.Context(), "ConfigLog \n", r.GetHandlerResponse())
	if s.logWhitelist(r) {
		return
	}

	g.Log().Debug(r.Context(), "ConfigLog \n", r.GetBodyString())
	m := service.BizCtx().Get(r.Context())
	go func() {
		defer func() {
			if err := recover(); err != nil {
				g.Log().Error(context.Background(), err)
			}
		}()
		var resByte = []byte("")
		if r.GetHandlerResponse() != nil {
			resByte, _ = json.Marshal(r.GetHandlerResponse())
		}
		user := &model.UserInfo{}
		err := dao.Users.Ctx(context.Background()).Where(dao.Users.Columns().Id, m.User.ID).FieldsEx(dao.Users.Columns().Slat, dao.Users.Columns().Password, dao.Users.Columns().TotpToken).Scan(user)
		if err != nil {
			g.Log().Error(context.Background(), err)
			return
		}
		errlog := ""
		if r.GetError() != nil {
			errlog = r.GetError().Error()
		}

		log := do.LogSettings{
			Name:          global.GetPathName(r.Request.URL.Path, r.Method),
			Loginname:     user.Loginname,
			Username:      user.Username,
			ClientIp:      r.GetClientIp(),
			UserId:        user.Id,
			Success:       r.GetError() == nil,
			RequestMethod: strings.ToUpper(r.Method),
			RequestPath:   r.Request.URL.Path,
			RequestBody:   s.reMasking(r.GetBodyString()),
			ResponseCode:  r.Response.Status,
			ResponseBody:  s.reMasking(string(resByte)),
			ResponseError: errlog,
		}

		_, err = dao.LogSettings.Ctx(context.Background()).Insert(&log)
		if err != nil {
			g.Log().Error(context.Background(), err)
			return
		}

	}()

}

// 密码脱敏
func (s *sMiddleware) reMasking(body string) (req string) {
	req = body
	passwords := regMask.FindAllString(body, -1)
	if passwords != nil && len(passwords) > 0 {
		for _, password := range passwords {
			req = strings.ReplaceAll(req, password, string(regMaskContent.ReplaceAll([]byte(password), []byte(`:"***"`))))
		}
	}
	return
}

// 配置日志白名单,true代表匹配成功
func (s *sMiddleware) logWhitelist(r *ghttp.Request) bool {
	if strings.ToUpper(r.Method) == "GET" {
		return true
	}
	//POST:/api/v1/accounts/search
	var list []string

	return slices.Contains(list, fmt.Sprintf("%s:%s", strings.ToUpper(r.Method), r.Request.URL.Path))
}
