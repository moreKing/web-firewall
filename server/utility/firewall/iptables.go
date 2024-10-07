package firewall

import "context"

type Iptables struct {
}

func (i *Iptables) Flush(ctx context.Context) error {

	return nil
}

func (i *Iptables) Add(table string, customTable string, dbName string) error {

	return nil

}
