package firewall

import (
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/entity"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

const (
	security = "security"
	filter   = "filter"
	nat      = "nat"
	mangle   = "mangle"
	raw      = "raw"
)

type Iptables struct {
}

func (i *Iptables) Flush(ctx context.Context) error {
	// 获取iptables规则
	str, err := exec.Command("iptables-save").CombinedOutput()
	if err != nil {
		return err
	}
	tmp := NewIptables(string(str))

	tmp.Init()

	//  增加input规则
	// 入站策略
	var inputList []entity.InputRules
	err = dao.InputRules.Ctx(ctx).OrderAsc(dao.InputRules.Columns().Position).Scan(&inputList)
	if err != nil {
		return err
	}

	for _, input := range inputList {
		//  判断协议类型
		if input.Protocol == "tcp" || input.Protocol == "udp" {
			if strings.TrimSpace(input.Ip) != "" {
				// 判断ip是否是多个
				for _, ip := range strings.Split(input.Ip, ",") {
					// 多个端口
					for _, port := range strings.Split(input.Port, ",") {
						tmpRule := fmt.Sprintf("-A %s ", ChainName[INPUT])
						if strings.Contains(ip, "-") {
							tmpRule += fmt.Sprintf(" -m iprange --src-range %s ", ip)
						} else {
							tmpRule += fmt.Sprintf(" -s %s ", ip)
						}

						port = strings.Replace(port, "-", ":", -1)
						tmpRule += fmt.Sprintf(" -p %s --dport %s ", input.Protocol, port)
						tmpRule += fmt.Sprintf(" -j %s ", strings.ToUpper(input.Policy))
						tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
					}
				}
			} else {
				for _, port := range strings.Split(input.Port, ",") {
					tmpRule := fmt.Sprintf("-A %s ", ChainName[INPUT])
					port = strings.Replace(port, "-", ":", -1)
					tmpRule += fmt.Sprintf(" -p %s --dport %s ", input.Protocol, port)
					tmpRule += fmt.Sprintf(" -j %s ", strings.ToUpper(input.Policy))
					tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
				}
			}

		}

		if input.Protocol == "ct" {
			if strings.TrimSpace(input.Ip) != "" {
				for _, ip := range strings.Split(input.Ip, ",") {
					tmpRule := fmt.Sprintf("-A %s ", ChainName[INPUT])
					if strings.Contains(ip, "-") {
						tmpRule += fmt.Sprintf(" -m iprange --src-range %s ", ip)
					} else {
						tmpRule += fmt.Sprintf(" -s %s ", ip)
					}

					tmpRule += fmt.Sprintf(" -m conntrack --ctstate %s ", strings.ToUpper(input.Ct))
					tmpRule += fmt.Sprintf(" -j %s ", strings.ToUpper(input.Policy))
					tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
				}
			} else {
				tmpRule := fmt.Sprintf("-A %s ", ChainName[INPUT])
				tmpRule += fmt.Sprintf(" -m conntrack --ctstate %s ", strings.ToUpper(input.Ct))
				tmpRule += fmt.Sprintf(" -j %s ", strings.ToUpper(input.Policy))
				tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
			}

		}

		if input.Protocol == "icmp" {
			if strings.TrimSpace(input.Ip) != "" {
				for _, ip := range strings.Split(input.Ip, ",") {
					for _, ic := range strings.Split(input.Icmp, ",") {
						tmpRule := fmt.Sprintf("-A %s ", ChainName[INPUT])
						if strings.Contains(ip, "-") {
							tmpRule += fmt.Sprintf(" -m iprange --src-range %s ", ip)
						} else {
							tmpRule += fmt.Sprintf(" -s %s", ip)
						}
						tmpRule += fmt.Sprintf(" -p icmp --icmp-type %s", ic)
						tmpRule += fmt.Sprintf(" -j %s", strings.ToUpper(input.Policy))
						tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
					}

				}
			} else {
				for _, ic := range strings.Split(input.Icmp, ",") {
					tmpRule := fmt.Sprintf("-A %s ", ChainName[INPUT])
					tmpRule += fmt.Sprintf(" -p icmp --icmp-type %s", ic)
					tmpRule += fmt.Sprintf(" -j %s", strings.ToUpper(input.Policy))
					tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
				}
			}
		}

	}

	//  增加入站限流
	var inputLimitList []entity.InputLimitRules
	err = dao.InputLimitRules.Ctx(ctx).OrderAsc(dao.InputLimitRules.Columns().Position).Scan(&inputLimitList)
	if err != nil {
		return err
	}

	for _, item := range inputLimitList {
		//  判断协议类型
		if item.Protocol == "tcp" || item.Protocol == "udp" {
			for _, ip := range strings.Split(item.Ip, ",") {
				// 多个端口
				for _, port := range strings.Split(item.Port, ",") {
					tmpRule := fmt.Sprintf("-A %s ", ChainName[LIMIT_INPUT])
					if strings.TrimSpace(ip) == "" {
					} else if strings.Contains(ip, "-") {
						tmpRule += fmt.Sprintf(" -m iprange --src-range %s ", ip)
					} else {
						tmpRule += fmt.Sprintf(" -s %s ", ip)
					}

					port = strings.Replace(port, "-", ":", -1)
					tmpRule += fmt.Sprintf(" -p %s --dport %s ", item.Protocol, port)

					tmpRule += fmt.Sprintf(" -m hashlimit --hashlimit-above %d%s  --hashlimit-name %s", item.Limit, item.Speed, gonanoid.Must())
					tmpRule += " -j DROP "
					tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
				}
			}

		} else {
			for _, ip := range strings.Split(item.Ip, ",") {
				tmpRule := fmt.Sprintf("-A %s ", ChainName[LIMIT_INPUT])
				if strings.TrimSpace(ip) == "" {

				} else if strings.Contains(ip, "-") {
					tmpRule += fmt.Sprintf(" -m iprange --src-range %s ", ip)
				} else {
					tmpRule += fmt.Sprintf(" -s %s ", ip)
				}
				tmpRule += fmt.Sprintf(" -m hashlimit --hashlimit-above %d%s  --hashlimit-name %s", item.Limit, item.Speed, gonanoid.Must())
				tmpRule += " -j DROP "
				tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
			}

		}
	}

	// todo 出站策略
	var outputList []entity.OutputRules
	err = dao.OutputRules.Ctx(ctx).OrderAsc(dao.OutputRules.Columns().Position).Scan(&outputList)
	if err != nil {
		return err
	}

	for _, item := range outputList {
		//  判断协议类型
		if item.Protocol == "tcp" || item.Protocol == "udp" {
			if strings.TrimSpace(item.Ip) != "" {
				// 判断ip是否是多个
				for _, ip := range strings.Split(item.Ip, ",") {
					// 多个端口
					for _, port := range strings.Split(item.Port, ",") {
						tmpRule := fmt.Sprintf("-A %s ", ChainName[OUTPUT])
						if strings.Contains(ip, "-") {
							tmpRule += fmt.Sprintf(" -m iprange --dst-range %s ", ip)
						} else {
							tmpRule += fmt.Sprintf(" -d %s ", ip)
						}

						port = strings.Replace(port, "-", ":", -1)
						tmpRule += fmt.Sprintf(" -p %s --dport %s ", item.Protocol, port)
						tmpRule += fmt.Sprintf(" -j %s ", strings.ToUpper(item.Policy))
						tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
					}
				}
			} else {
				for _, port := range strings.Split(item.Port, ",") {
					tmpRule := fmt.Sprintf("-A %s ", ChainName[OUTPUT])
					port = strings.Replace(port, "-", ":", -1)
					tmpRule += fmt.Sprintf(" -p %s --dport %s ", item.Protocol, port)
					tmpRule += fmt.Sprintf(" -j %s ", strings.ToUpper(item.Policy))
					tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
				}
			}

		}

		if item.Protocol == "ct" {
			if strings.TrimSpace(item.Ip) != "" {
				for _, ip := range strings.Split(item.Ip, ",") {
					tmpRule := fmt.Sprintf("-A %s ", ChainName[OUTPUT])
					if strings.Contains(ip, "-") {
						tmpRule += fmt.Sprintf(" -m iprange --dst-range %s ", ip)
					} else {
						tmpRule += fmt.Sprintf(" -d %s ", ip)
					}

					tmpRule += fmt.Sprintf(" -m conntrack --ctstate %s ", strings.ToUpper(item.Ct))
					tmpRule += fmt.Sprintf(" -j %s ", strings.ToUpper(item.Policy))
					tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
				}
			} else {
				tmpRule := fmt.Sprintf("-A %s ", ChainName[OUTPUT])
				tmpRule += fmt.Sprintf(" -m conntrack --ctstate %s ", strings.ToUpper(item.Ct))
				tmpRule += fmt.Sprintf(" -j %s ", strings.ToUpper(item.Policy))
				tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
			}

		}

		if item.Protocol == "icmp" {
			if strings.TrimSpace(item.Ip) != "" {
				for _, ip := range strings.Split(item.Ip, ",") {
					for _, ic := range strings.Split(item.Icmp, ",") {
						tmpRule := fmt.Sprintf("-A %s ", ChainName[OUTPUT])
						if strings.Contains(ip, "-") {
							tmpRule += fmt.Sprintf(" -m iprange --dst-range %s ", ip)
						} else {
							tmpRule += fmt.Sprintf(" -d %s", ip)
						}
						tmpRule += fmt.Sprintf(" -p icmp --icmp-type %s", ic)
						tmpRule += fmt.Sprintf(" -j %s", strings.ToUpper(item.Policy))
						tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
					}

				}
			} else {
				for _, ic := range strings.Split(item.Icmp, ",") {
					tmpRule := fmt.Sprintf("-A %s ", ChainName[OUTPUT])
					tmpRule += fmt.Sprintf(" -p icmp --icmp-type %s", ic)
					tmpRule += fmt.Sprintf(" -j %s", strings.ToUpper(item.Policy))
					tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
				}
			}
		}

	}

	var outputLimitList []entity.OutputLimitRules
	err = dao.OutputLimitRules.Ctx(ctx).OrderAsc(dao.OutputLimitRules.Columns().Position).Scan(&outputLimitList)
	if err != nil {
		return err
	}
	for _, item := range outputLimitList {
		//  判断协议类型
		if item.Protocol == "tcp" || item.Protocol == "udp" {
			for _, ip := range strings.Split(item.Ip, ",") {
				// 多个端口
				for _, port := range strings.Split(item.Port, ",") {
					tmpRule := fmt.Sprintf("-A %s ", ChainName[LIMIT_OUTPUT])
					if strings.TrimSpace(ip) == "" {
					} else if strings.Contains(ip, "-") {
						tmpRule += fmt.Sprintf(" -m iprange --dst-range %s ", ip)
					} else {
						tmpRule += fmt.Sprintf(" -d %s ", ip)
					}

					port = strings.Replace(port, "-", ":", -1)
					tmpRule += fmt.Sprintf(" -p %s --sport %s ", item.Protocol, port)

					tmpRule += fmt.Sprintf(" -m hashlimit --hashlimit-above %d%s  --hashlimit-name %s", item.Limit, item.Speed, gonanoid.Must())
					tmpRule += " -j DROP "
					tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
				}
			}

		} else {
			for _, ip := range strings.Split(item.Ip, ",") {
				tmpRule := fmt.Sprintf("-A %s ", ChainName[LIMIT_OUTPUT])
				if strings.TrimSpace(ip) == "" {

				} else if strings.Contains(ip, "-") {
					tmpRule += fmt.Sprintf(" -m iprange --dst-range %s ", ip)
				} else {
					tmpRule += fmt.Sprintf(" -s %s ", ip)
				}
				tmpRule += fmt.Sprintf(" -m hashlimit --hashlimit-above %d%s  --hashlimit-name %s", item.Limit, item.Speed, gonanoid.Must())
				tmpRule += " -j DROP "
				tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
			}

		}
	}

	// todo 转发策略
	var forwardList []entity.ForwardRules
	err = dao.ForwardRules.Ctx(ctx).OrderAsc(dao.ForwardRules.Columns().Position).Scan(&forwardList)
	if err != nil {
		return err
	}
	for _, input := range forwardList {
		//  判断协议类型
		if input.Protocol == "tcp" || input.Protocol == "udp" {
			for _, sip := range strings.Split(input.Sip, ",") {
				for _, dip := range strings.Split(input.Dip, ",") {
					// 多个端口
					for _, port := range strings.Split(input.Port, ",") {
						tmpRule := fmt.Sprintf("-A %s ", ChainName[FORWARD])
						if strings.TrimSpace(sip) == "" {
						} else if strings.Contains(sip, "-") {
							tmpRule += fmt.Sprintf(" -m iprange --src-range %s ", sip)
						} else {
							tmpRule += fmt.Sprintf(" -s %s ", sip)
						}

						if strings.TrimSpace(dip) == "" {
						} else if strings.Contains(dip, "-") {
							tmpRule += fmt.Sprintf(" -m iprange --dst-range %s ", dip)
						} else {
							tmpRule += fmt.Sprintf(" -d %s ", dip)
						}

						port = strings.Replace(port, "-", ":", -1)
						tmpRule += fmt.Sprintf(" -p %s --dport %s ", input.Protocol, port)
						tmpRule += fmt.Sprintf(" -j %s ", strings.ToUpper(input.Policy))
						tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
					}
				}

			}

		} else {
			for _, sip := range strings.Split(input.Sip, ",") {
				for _, dip := range strings.Split(input.Dip, ",") {
					tmpRule := fmt.Sprintf("-A %s ", ChainName[FORWARD])
					if strings.TrimSpace(sip) == "" {
					} else if strings.Contains(sip, "-") {
						tmpRule += fmt.Sprintf(" -m iprange --src-range %s ", sip)
					} else {
						tmpRule += fmt.Sprintf(" -s %s ", sip)
					}

					if strings.TrimSpace(dip) == "" {
					} else if strings.Contains(dip, "-") {
						tmpRule += fmt.Sprintf(" -m iprange --dst-range %s ", dip)
					} else {
						tmpRule += fmt.Sprintf(" -d %s ", dip)
					}
					tmpRule += fmt.Sprintf(" -j %s ", strings.ToUpper(input.Policy))
					tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
				}

			}
		}

	}

	// todo 转发流控

	var forwardLimitList []entity.ForwardLimitRules
	err = dao.ForwardLimitRules.Ctx(ctx).OrderAsc(dao.ForwardLimitRules.Columns().Position).Scan(&forwardLimitList)
	if err != nil {
		return err
	}

	for _, item := range forwardLimitList {
		//  判断协议类型
		if item.Protocol == "tcp" || item.Protocol == "udp" {
			for _, sip := range strings.Split(item.Sip, ",") {
				for _, dip := range strings.Split(item.Dip, ",") {
					// 多个端口
					for _, port := range strings.Split(item.Port, ",") {
						tmpRule := fmt.Sprintf("-A %s ", ChainName[LIMIT_FORWARD])
						if strings.TrimSpace(sip) == "" {
						} else if strings.Contains(sip, "-") {
							tmpRule += fmt.Sprintf(" -m iprange --src-range %s ", sip)
						} else {
							tmpRule += fmt.Sprintf(" -s %s ", sip)
						}

						if strings.TrimSpace(dip) == "" {
						} else if strings.Contains(dip, "-") {
							tmpRule += fmt.Sprintf(" -m iprange --dst-range %s ", dip)
						} else {
							tmpRule += fmt.Sprintf(" -d %s ", dip)
						}

						if item.PortType == "sport" {
							port = strings.Replace(port, "-", ":", -1)
							tmpRule += fmt.Sprintf(" -p %s --sport %s ", item.Protocol, port)
						}
						if item.PortType == "dport" {
							port = strings.Replace(port, "-", ":", -1)
							tmpRule += fmt.Sprintf(" -p %s --dport %s ", item.Protocol, port)
						}

						tmpRule += fmt.Sprintf(" -m hashlimit --hashlimit-above %d%s  --hashlimit-name %s", item.Limit, item.Speed, gonanoid.Must())
						tmpRule += " -j DROP "
						tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
					}
				}

			}

		} else {
			for _, sip := range strings.Split(item.Sip, ",") {
				for _, dip := range strings.Split(item.Dip, ",") {
					tmpRule := fmt.Sprintf("-A %s ", ChainName[LIMIT_FORWARD])
					if strings.TrimSpace(sip) == "" {
					} else if strings.Contains(sip, "-") {
						tmpRule += fmt.Sprintf(" -m iprange --src-range %s ", sip)
					} else {
						tmpRule += fmt.Sprintf(" -s %s ", sip)
					}

					if strings.TrimSpace(dip) == "" {
					} else if strings.Contains(dip, "-") {
						tmpRule += fmt.Sprintf(" -m iprange --dst-range %s ", dip)
					} else {
						tmpRule += fmt.Sprintf(" -d %s ", dip)
					}
					tmpRule += fmt.Sprintf(" -m hashlimit --hashlimit-above %d%s  --hashlimit-name %s", item.Limit, item.Speed, gonanoid.Must())
					tmpRule += " -j DROP "
					tmp.Filter.Rule = append(tmp.Filter.Rule, tmpRule)
				}
			}

		}
	}

	// todo 源地址转换
	var snatList []entity.SnatRules
	err = dao.SnatRules.Ctx(ctx).OrderAsc(dao.SnatRules.Columns().Position).Scan(&snatList)
	if err != nil {
		return err
	}

	for _, item := range snatList {
		for _, sip := range strings.Split(item.Sip, ",") {
			for _, dip := range strings.Split(item.Dip, ",") {
				tmpRule := fmt.Sprintf("-A %s", ChainName[SNAT])
				tmpRule += fmt.Sprintf(" -o %s", item.Oif)
				if strings.TrimSpace(sip) == "" {

				} else if strings.Contains(sip, "-") {
					tmpRule += fmt.Sprintf(" -m iprange --src-range  %s ", sip)
				} else {
					tmpRule += fmt.Sprintf(" -s %s ", sip)
				}

				if strings.TrimSpace(dip) == "" {

				} else if strings.Contains(dip, "-") {
					tmpRule += fmt.Sprintf(" -m iprange --dst-range  %s ", dip)
				} else {
					tmpRule += fmt.Sprintf(" -d %s ", dip)
				}

				if strings.TrimSpace(item.Snat) != "" {
					tmpRule += fmt.Sprintf(" -j SNAT --to-source %s ", item.Snat)
				} else {
					tmpRule += " -j MASQUERADE"
				}

				tmp.Nat.Rule = append(tmp.Nat.Rule, tmpRule)
			}

		}
	}

	// todo 目的地址转换
	var dnatList []model.DnatRulesets
	err = dao.DnatRules.Ctx(ctx).OrderAsc(dao.DnatRules.Columns().Position).Scan(&dnatList)
	if err != nil {
		return err
	}

	for _, item := range dnatList {
		if item.Protocol == "tcp" || item.Protocol == "udp" {
			for _, dip := range strings.Split(item.Dip, ",") {
				for _, port := range item.Port {
					tmpRule := fmt.Sprintf("-A %s", ChainName[DNAT])
					tmpRule += fmt.Sprintf(" -i %s", item.Iif)

					if strings.TrimSpace(dip) == "" {

					} else if strings.Contains(dip, "-") {
						tmpRule += fmt.Sprintf(" -m iprange --dst-range  %s ", dip)
					} else {
						tmpRule += fmt.Sprintf(" -d %s ", dip)
					}

					tmpRule += fmt.Sprintf(" -p %s --dport %d -j DNAT --to-destination %s:%d ", item.Protocol, port.Key, item.Dnat, port.Value)

					tmp.Nat.Rule = append(tmp.Nat.Rule, tmpRule)
				}
			}
		}

	}

	// 写入iptables
	g.Log().Debug(ctx, "iptables  \n", fmt.Sprintf("iptables-restore  <<EOF\n%s\nEOF", tmp.String()))
	err = exec.Command("sh", "-c", fmt.Sprintf("iptables-restore  <<EOF\n%s\nEOF", tmp.String())).Run()

	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

type iptables struct {
	Security struct {
		Table []string
		Rule  []string
	} //主要针对的是数据链路层的管理规则组合，只不过它是基于target目标来操作的。
	Filter struct {
		Table []string
		Rule  []string
	} //是netfilter中最重要的表，也是默认的表，主要负责数据包的过滤功能
	Nat struct {
		Table []string
		Rule  []string
	} //主要实现网络地址转换的表。可以自由转换数据报文中的ip和port
	Mangle struct {
		Table []string
		Rule  []string
	} // 主要实现数据包的拆分-修改-封装动作
	Raw struct {
		Table []string
		Rule  []string
	} // 通过关闭nat表的追踪功能，从而实现加速防火墙过滤的表
}

func NewIptables(RuleStr string) *iptables {
	isTable := false
	currentTable := ""
	ipt := iptables{
		Nat: struct {
			Table []string
			Rule  []string
		}{
			Table: []string{},
			Rule:  []string{},
		},
		Filter: struct {
			Table []string
			Rule  []string
		}{
			Table: []string{},
			Rule:  []string{},
		},
		Security: struct {
			Table []string
			Rule  []string
		}{
			Table: []string{},
			Rule:  []string{},
		},
		Mangle: struct {
			Table []string
			Rule  []string
		}{
			Table: []string{},
			Rule:  []string{},
		},
		Raw: struct {
			Table []string
			Rule  []string
		}{
			Table: []string{},
			Rule:  []string{},
		},
	}

	for _, line := range strings.Split(RuleStr, "\n") {
		line = strings.TrimSpace(line)
		if !isTable {
			switch line {
			case fmt.Sprintf("*%s", security):
				currentTable = security
				isTable = true

			case fmt.Sprintf("*%s", filter):
				currentTable = filter
				isTable = true

			case fmt.Sprintf("*%s", nat):
				currentTable = nat
				isTable = true

			case fmt.Sprintf("*%s", mangle):
				currentTable = mangle
				isTable = true

			case fmt.Sprintf("*%s", raw):
				currentTable = raw
				isTable = true
			}

			continue
		}

		if line == "COMMIT" {
			isTable = false
			currentTable = ""
			continue
		}

		switch currentTable {
		case security:
			if strings.HasPrefix(line, ":") {
				ipt.Security.Table = append(ipt.Security.Table, line)
			} else {
				ipt.Security.Rule = append(ipt.Security.Rule, line)
			}

		case filter:
			if strings.HasPrefix(line, ":") {
				ipt.Filter.Table = append(ipt.Filter.Table, line)
			} else {
				ipt.Filter.Rule = append(ipt.Filter.Rule, line)
			}

		case nat:
			if strings.HasPrefix(line, ":") {
				ipt.Nat.Table = append(ipt.Nat.Table, line)
			} else {
				ipt.Nat.Rule = append(ipt.Nat.Rule, line)
			}

		case mangle:

			if strings.HasPrefix(line, ":") {
				ipt.Mangle.Table = append(ipt.Mangle.Table, line)
			} else {
				ipt.Mangle.Rule = append(ipt.Mangle.Rule, line)
			}

		case raw:
			if strings.HasPrefix(line, ":") {
				ipt.Raw.Table = append(ipt.Raw.Table, line)
			} else {
				ipt.Raw.Rule = append(ipt.Raw.Rule, line)
			}
		}

	}
	return &ipt

}

func (i *iptables) Init() {

	isInput := false
	isOutput := false
	isForward := false
	isMInput := false
	isMInputLimit := false
	isMOutput := false
	isMOutputLimit := false
	isMForward := false
	isMForwardLimit := false

	// 判断是否存在基本的表，不存在则创建
	for _, v := range i.Filter.Table {
		if strings.HasPrefix(v, ":INPUT ") {
			isInput = true
			continue
		}
		if strings.HasPrefix(v, ":OUTPUT ") {
			isOutput = true
			continue
		}
		if strings.HasPrefix(v, ":FORWARD ") {
			isForward = true
			continue
		}

		// 自定义的链
		if strings.HasPrefix(v, fmt.Sprintf(":%s ", ChainName[INPUT])) {
			isMInput = true
			continue
		}
		if strings.HasPrefix(v, fmt.Sprintf(":%s ", ChainName[LIMIT_INPUT])) {
			isMInputLimit = true
			continue
		}
		if strings.HasPrefix(v, fmt.Sprintf(":%s ", ChainName[OUTPUT])) {
			isMOutput = true
			continue
		}
		if strings.HasPrefix(v, fmt.Sprintf(":%s ", ChainName[LIMIT_OUTPUT])) {
			isMOutputLimit = true
			continue
		}
		if strings.HasPrefix(v, fmt.Sprintf(":%s ", ChainName[FORWARD])) {
			isMForward = true
			continue
		}
		if strings.HasPrefix(v, fmt.Sprintf(":%s ", ChainName[LIMIT_FORWARD])) {
			isMForwardLimit = true
			continue
		}

	}

	if !isInput {
		i.Filter.Table = append(i.Filter.Table, ":INPUT ACCEPT [0:0]")
	}

	if !isOutput {
		i.Filter.Table = append(i.Filter.Table, ":OUTPUT ACCEPT [0:0]")
	}

	if !isForward {
		i.Filter.Table = append(i.Filter.Table, ":FORWARD ACCEPT [0:0]")
	}

	if !isMInputLimit {
		i.Filter.Table = append(i.Filter.Table, fmt.Sprintf(":%s - [0:0]", ChainName[LIMIT_INPUT]))
	}

	if !isMOutputLimit {
		i.Filter.Table = append(i.Filter.Table, fmt.Sprintf(":%s - [0:0]", ChainName[LIMIT_OUTPUT]))
	}

	if !isMForwardLimit {
		i.Filter.Table = append(i.Filter.Table, fmt.Sprintf(":%s - [0:0]", ChainName[LIMIT_FORWARD]))
	}

	if !isMInput {
		i.Filter.Table = append(i.Filter.Table, fmt.Sprintf(":%s - [0:0]", ChainName[INPUT]))
	}

	if !isMOutput {
		i.Filter.Table = append(i.Filter.Table, fmt.Sprintf(":%s - [0:0]", ChainName[OUTPUT]))
	}

	if !isMForward {
		i.Filter.Table = append(i.Filter.Table, fmt.Sprintf(":%s - [0:0]", ChainName[FORWARD]))
	}

	isPreRouting := false
	isPostRouting := false

	isMPreRouting := false
	isMPostRouting := false

	for _, v := range i.Nat.Table {
		if strings.HasPrefix(v, ":PREROUTING ") {
			isPreRouting = true
			continue
		}
		if strings.HasPrefix(v, ":POSTROUTING ") {
			isPostRouting = true
			continue
		}

		if strings.HasPrefix(v, fmt.Sprintf(":%s ", ChainName[DNAT])) {
			isMPreRouting = true
			continue
		}
		if strings.HasPrefix(v, fmt.Sprintf(":%s ", ChainName[SNAT])) {
			isMPostRouting = true
			continue
		}

	}

	if !isPreRouting {
		i.Nat.Table = append(i.Nat.Table, ":PREROUTING ACCEPT [0:0]")
	}
	if !isPostRouting {
		i.Nat.Table = append(i.Nat.Table, ":POSTROUTING ACCEPT [0:0]")
	}
	if !isMPreRouting {
		i.Nat.Table = append(i.Nat.Table, fmt.Sprintf(":%s - [0:0]", ChainName[DNAT]))
	}
	if !isMPostRouting {
		i.Nat.Table = append(i.Nat.Table, fmt.Sprintf(":%s - [0:0]", ChainName[SNAT]))
	}

	// 过滤规则
	newFilterRules := []string{}
	// 判断是否存在默认规则，不存在则创建

	loPattern := regexp.MustCompile(`^-A\s+INPUT\s+-i\s+lo\s+-j\s+ACCEPT$`)

	dropPattern := regexp.MustCompile(`^-A\s+INPUT\s+-j\s+DROP$`)

	for _, v := range i.Filter.Rule {
		v := strings.TrimSpace(v)

		if loPattern.MatchString(v) {
			continue
		}

		if dropPattern.MatchString(v) {
			continue
		}

		if strings.Contains(v, fmt.Sprintf(" %s ", ChainName[INPUT])) || strings.HasSuffix(v, fmt.Sprintf(" %s", ChainName[INPUT])) {
			continue
		}

		if strings.Contains(v, fmt.Sprintf(" %s ", ChainName[OUTPUT])) || strings.HasSuffix(v, fmt.Sprintf(" %s", ChainName[OUTPUT])) {
			continue
		}

		if strings.Contains(v, fmt.Sprintf(" %s ", ChainName[FORWARD])) || strings.HasSuffix(v, fmt.Sprintf(" %s", ChainName[FORWARD])) {
			continue
		}

		if strings.Contains(v, fmt.Sprintf(" %s ", ChainName[LIMIT_INPUT])) || strings.HasSuffix(v, fmt.Sprintf(" %s", ChainName[LIMIT_INPUT])) {
			continue
		}

		if strings.Contains(v, fmt.Sprintf(" %s ", ChainName[LIMIT_OUTPUT])) || strings.HasSuffix(v, fmt.Sprintf(" %s", ChainName[LIMIT_OUTPUT])) {
			continue
		}

		if strings.Contains(v, fmt.Sprintf(" %s ", ChainName[LIMIT_FORWARD])) || strings.HasSuffix(v, fmt.Sprintf(" %s", ChainName[LIMIT_FORWARD])) {
			continue
		}

		newFilterRules = append(newFilterRules, v)
	}

	i.Filter.Rule = append([]string{
		"-A INPUT -i lo -j ACCEPT",
		fmt.Sprintf("-A INPUT -j %s", ChainName[LIMIT_INPUT]),
		fmt.Sprintf("-A INPUT -j %s", ChainName[INPUT]),
		fmt.Sprintf("-A OUTPUT -j %s", ChainName[LIMIT_OUTPUT]),
		fmt.Sprintf("-A OUTPUT -j %s", ChainName[OUTPUT]),
		fmt.Sprintf("-A FORWARD -j %s", ChainName[LIMIT_FORWARD]),
		fmt.Sprintf("-A FORWARD -j %s", ChainName[FORWARD]),
	}, newFilterRules...)
	i.Filter.Rule = append(i.Filter.Rule, "-A INPUT -j DROP")

	newNatRules := []string{}

	for _, v := range i.Nat.Rule {
		v := strings.TrimSpace(v)
		if strings.Contains(v, fmt.Sprintf(" %s ", ChainName[SNAT])) || strings.HasSuffix(v, fmt.Sprintf(" %s", ChainName[SNAT])) {
			continue
		}

		if strings.Contains(v, fmt.Sprintf(" %s ", ChainName[DNAT])) || strings.HasSuffix(v, fmt.Sprintf(" %s", ChainName[DNAT])) {
			continue
		}

		newNatRules = append(newNatRules, v)
	}

	i.Nat.Rule = append([]string{
		fmt.Sprintf("-A PREROUTING -j %s", ChainName[DNAT]),
		fmt.Sprintf("-A POSTROUTING -j %s", ChainName[SNAT]),
	}, newNatRules...)

}

func (i *iptables) String() string {

	var res = fmt.Sprintf("*filter\n%s\n%s\nCOMMIT\n", strings.Join(i.Filter.Table, "\n"), strings.Join(i.Filter.Rule, "\n"))
	res += fmt.Sprintf("*nat\n%s\n%s\nCOMMIT\n", strings.Join(i.Nat.Table, "\n"), strings.Join(i.Nat.Rule, "\n"))

	if len(i.Mangle.Table) > 0 {
		res += fmt.Sprintf("*mangle\n%s\n%s\nCOMMIT\n", strings.Join(i.Mangle.Table, "\n"), strings.Join(i.Mangle.Rule, "\n"))
	}

	if len(i.Security.Table) > 0 {
		res += fmt.Sprintf("*mangle\n%s\n%s\nCOMMIT\n", strings.Join(i.Security.Table, "\n"), strings.Join(i.Security.Rule, "\n"))
	}

	if len(i.Raw.Table) > 0 {
		res += fmt.Sprintf("*mangle\n%s\n%s\nCOMMIT\n", strings.Join(i.Raw.Table, "\n"), strings.Join(i.Raw.Rule, "\n"))
	}

	return res
}
