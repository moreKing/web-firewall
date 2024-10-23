package policy

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"regexp"
	"server/internal/dao"
	"server/internal/model/do"
	"server/internal/service"
	"server/utility/firewall"
)

type sPolicy struct {
	p firewall.Firewall
}

func init() {
	service.RegisterPolicy(New())
}

func New() service.IPolicy {
	// 检测配置文件中web的端口
	addr := g.Cfg().MustGet(context.Background(), "server.address", ":8000").String()
	re := regexp.MustCompile(`(\d+)\s*$`)
	portByte := re.Find([]byte(addr))

	g.Log().Info(context.Background(), string(portByte))

	_, err2 := dao.InputRules.Ctx(context.Background()).Where(dao.InputRules.Columns().Id, 4).Update(do.InputRules{
		Port: string(portByte),
	})
	if err2 != nil {
		g.Log().Error(context.Background(), err2)
	}

	s := &sPolicy{p: firewall.New()}
	err := s.Flush(context.Background())
	if err != nil {
		panic(err)
	}
	return s
}

func (s *sPolicy) Flush(ctx context.Context) error {
	return s.p.Flush(ctx)
}
