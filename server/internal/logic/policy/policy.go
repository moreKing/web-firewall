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
