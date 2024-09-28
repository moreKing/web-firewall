package policy

import (
	"context"
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
	return &sPolicy{p: firewall.New()}
}

func (s *sPolicy) Flush(ctx context.Context) error {
	return s.p.Flush(ctx)
}
