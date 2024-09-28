package firewall

import "context"

type Iptables struct {
}

func (i *Iptables) Flush(ctx context.Context) error {

	return nil
}
