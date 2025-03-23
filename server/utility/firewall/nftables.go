package firewall

import (
	"context"
	"fmt"
	"os/exec"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/entity"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
)

/**
代码的逻辑是，当用户策略发生变更需要下发策略时，为了防止用户自行操作导致防火墙策略与系统重的不一致，因此生成全量策略全量完整下发，避免下发策略失败
*/

type Nftables struct {
	basicStr string
}

const (
	_ = iota
	INPUT
	OUTPUT
	DNAT
	SNAT
	LIMIT_INPUT
	LIMIT_OUTPUT
	FORWARD
	LIMIT_FORWARD
	IP_B_W
	INPUT_BASIC
	OUTPUT_BASIC
	FORWARD_BASIC
)

var ChainName = map[int]string{
	INPUT:         "m_input",
	OUTPUT:        "m_output",
	DNAT:          "m_dnat",
	SNAT:          "m_snat",
	LIMIT_INPUT:   "m_limit_input",
	LIMIT_OUTPUT:  "m_limit_output",
	FORWARD:       "m_forward",
	IP_B_W:        "m_ip_b_w", // ip黑白名单
	INPUT_BASIC:   "m_input_basic",
	OUTPUT_BASIC:  "m_output_basic",
	FORWARD_BASIC: "m_forward_basic",
	LIMIT_FORWARD: "m_limit_forward",
}

var speedMap = map[string]string{
	"kb/s": "kbytes/second",
	"mb/s": "mbytes/second",
	"kb/m": "kbytes/minute",
	"mb/m": "mbytes/minute",
}

var basicStr = ""

func NewNftables() *Nftables {
	inputPriority := g.Cfg().MustGet(context.Background(), "firewall.chainPriority.input", 100).Int()
	// 出站策略 output链
	outputPriority := g.Cfg().MustGet(context.Background(), "firewall.chainPriority.output", 100).Int()

	//todo 转发策略（作为网关时）forward 链
	forwardPriority := g.Cfg().MustGet(context.Background(), "firewall.chainPriority.forward", 100).Int()

	//DNAT prerouting链
	preroutingPriority := g.Cfg().MustGet(context.Background(), "firewall.chainPriority.prerouting", 100).Int()
	//SNAT postrouting
	postroutingPriority := g.Cfg().MustGet(context.Background(), "firewall.chainPriority.postrouting", 100).Int()

	basicStr = fmt.Sprintf(`
add table inet web-firewall
delete table inet web-firewall
add table inet web-firewall

add chain inet web-firewall %s {type filter hook input priority %d; policy drop;}
add chain inet web-firewall %s
add chain inet web-firewall %s

add rule inet web-firewall %s iif lo accept
add rule inet web-firewall %s jump %s
add rule inet web-firewall %s jump %s

add chain inet web-firewall %s {type filter hook output priority %d; policy accept;}
add chain inet web-firewall %s
add chain inet web-firewall %s

add rule inet web-firewall %s jump %s
add rule inet web-firewall %s jump %s

add chain inet web-firewall %s {type filter hook forward priority %d; policy accept;}
add chain inet web-firewall %s
add chain inet web-firewall %s

add rule inet web-firewall %s jump %s
add rule inet web-firewall %s jump %s

add chain inet web-firewall %s {type nat hook prerouting priority %d; policy accept;}
add chain inet web-firewall %s {type nat hook postrouting priority %d; policy accept;}


`, ChainName[INPUT_BASIC], inputPriority, ChainName[INPUT], ChainName[LIMIT_INPUT], ChainName[INPUT_BASIC], ChainName[INPUT_BASIC], ChainName[LIMIT_INPUT], ChainName[INPUT_BASIC], ChainName[INPUT],
		ChainName[OUTPUT_BASIC], outputPriority, ChainName[OUTPUT], ChainName[LIMIT_OUTPUT], ChainName[OUTPUT_BASIC], ChainName[LIMIT_OUTPUT], ChainName[OUTPUT_BASIC], ChainName[OUTPUT],
		ChainName[FORWARD_BASIC], forwardPriority, ChainName[FORWARD], ChainName[LIMIT_FORWARD], ChainName[FORWARD_BASIC], ChainName[LIMIT_FORWARD], ChainName[FORWARD_BASIC], ChainName[FORWARD],
		ChainName[DNAT], preroutingPriority,
		ChainName[SNAT], postroutingPriority)

	return &Nftables{basicStr}

}

// Flush 从数据库中读取数据，生成nft备份文件，并将备份文件重新导入
func (n *Nftables) Flush(ctx context.Context) error {

	var ruleList []string

	// 入站策略
	var inputList []entity.InputRules
	err := dao.InputRules.Ctx(ctx).OrderAsc(dao.InputRules.Columns().Position).Scan(&inputList)
	if err != nil {
		return err
	}

	for _, input := range inputList {
		line := ""
		if strings.TrimSpace(input.Ip) != "" {
			if strings.Contains(input.Ip, ",") {
				line += fmt.Sprintf("ip saddr { %s } ", input.Ip)
			} else {
				line += fmt.Sprintf("ip saddr %s ", input.Ip)
			}
		}
		if input.Protocol == "tcp" || input.Protocol == "udp" {
			if strings.Contains(input.Port, ",") {
				line += fmt.Sprintf("%s dport { %s } ", input.Protocol, input.Port)
			} else {
				line += fmt.Sprintf("%s dport %s ", input.Protocol, input.Port)
			}
		}

		if input.Protocol == "icmp" {
			tmp := strings.Split(input.Icmp, ",")
			for i := range tmp {
				tmp[i] = fmt.Sprintf("%s:%s ", tmp[i], input.Policy)
			}

			line += fmt.Sprintf("icmp type vmap { %s } ", strings.Join(tmp, ", "))
		}

		if input.Protocol == "ct" {
			tmp := strings.Split(input.Ct, ",")
			for i := range tmp {
				tmp[i] = fmt.Sprintf("%s:%s ", tmp[i], input.Policy)
			}
			line += fmt.Sprintf("ct state vmap { %s } ", strings.Join(tmp, ", "))
		}

		line += fmt.Sprintf(" %s", input.Policy)

		ruleList = append(ruleList, fmt.Sprintf("add rule inet web-firewall %s %s", ChainName[INPUT], line))
	}

	// todo 入站流控

	var inputLimitList []entity.InputLimitRules
	err = dao.InputLimitRules.Ctx(ctx).OrderAsc(dao.InputLimitRules.Columns().Position).Scan(&inputLimitList)
	if err != nil {
		return err
	}

	g.Log().Debug(ctx, "output:== ", inputLimitList)

	for _, limit := range inputLimitList {
		line := ""
		if strings.TrimSpace(limit.Ip) != "" {
			if strings.Contains(limit.Ip, ",") {
				line += fmt.Sprintf("ip saddr { %s } ", limit.Ip)
			} else {
				line += fmt.Sprintf("ip saddr %s ", limit.Ip)
			}
		}
		if limit.Protocol == "tcp" || limit.Protocol == "udp" {
			if strings.Contains(limit.Port, ",") {
				line += fmt.Sprintf("%s dport { %s } ", limit.Protocol, limit.Port)
			} else {
				line += fmt.Sprintf("%s dport %s ", limit.Protocol, limit.Port)
			}
		}

		line += fmt.Sprintf(" limit rate over %d %s drop", limit.Limit, speedMap[limit.Speed])

		g.Log().Debug(ctx, line)

		ruleList = append(ruleList, fmt.Sprintf("add rule inet web-firewall %s %s", ChainName[LIMIT_INPUT], line))
	}

	// todo 出站策略
	var outputList []entity.OutputRules
	err = dao.OutputRules.Ctx(ctx).OrderAsc(dao.OutputRules.Columns().Position).Scan(&outputList)
	if err != nil {
		return err
	}

	//g.Log().Debug(ctx, "output:== ", outputList)

	for _, output := range outputList {
		line := ""
		if strings.TrimSpace(output.Ip) != "" {
			if strings.Contains(output.Ip, ",") {
				line += fmt.Sprintf("ip daddr { %s } ", output.Ip)
			} else {
				line += fmt.Sprintf("ip daddr %s ", output.Ip)
			}
		}
		if output.Protocol == "tcp" || output.Protocol == "udp" {
			if strings.Contains(output.Port, ",") {
				line += fmt.Sprintf("%s dport { %s } ", output.Protocol, output.Port)
			} else {
				line += fmt.Sprintf("%s dport %s ", output.Protocol, output.Port)
			}
		}

		if output.Protocol == "icmp" {
			tmp := strings.Split(output.Icmp, ",")
			for i := range tmp {
				tmp[i] = fmt.Sprintf("%s:%s ", tmp[i], output.Policy)
			}

			line += fmt.Sprintf("icmp type vmap { %s } ", strings.Join(tmp, ", "))
		}

		if output.Protocol == "ct" {
			tmp := strings.Split(output.Ct, ",")
			for i := range tmp {
				tmp[i] = fmt.Sprintf("%s:%s ", tmp[i], output.Policy)
			}
			line += fmt.Sprintf("ct state vmap { %s } ", strings.Join(tmp, ", "))
		}

		line += fmt.Sprintf(" %s", output.Policy)
		g.Log().Debug(ctx, line)

		ruleList = append(ruleList, fmt.Sprintf("add rule inet web-firewall %s %s", ChainName[OUTPUT], line))
	}
	// todo 出站流控

	var outputLimitList []entity.OutputLimitRules
	err = dao.OutputLimitRules.Ctx(ctx).OrderAsc(dao.OutputLimitRules.Columns().Position).Scan(&outputLimitList)
	if err != nil {
		return err
	}

	for _, limit := range outputLimitList {
		line := ""
		if strings.TrimSpace(limit.Ip) != "" {
			if strings.Contains(limit.Ip, ",") {
				line += fmt.Sprintf("ip daddr { %s } ", limit.Ip)
			} else {
				line += fmt.Sprintf("ip daddr %s ", limit.Ip)
			}
		}
		if limit.Protocol == "tcp" || limit.Protocol == "udp" {
			if strings.Contains(limit.Port, ",") {
				line += fmt.Sprintf("%s sport { %s } ", limit.Protocol, limit.Port)
			} else {
				line += fmt.Sprintf("%s sport %s ", limit.Protocol, limit.Port)
			}
		}

		line += fmt.Sprintf(" limit rate over %d %s drop", limit.Limit, speedMap[limit.Speed])

		g.Log().Debug(ctx, line)

		ruleList = append(ruleList, fmt.Sprintf("add rule inet web-firewall %s %s", ChainName[LIMIT_OUTPUT], line))
	}

	// todo 源地址转换

	var snatList []entity.SnatRules
	err = dao.SnatRules.Ctx(ctx).OrderAsc(dao.SnatRules.Columns().Position).Scan(&snatList)
	if err != nil {
		return err
	}

	for _, item := range snatList {
		line := ""
		if strings.TrimSpace(item.Oif) != "" {
			line += fmt.Sprintf("oifname %s ", item.Oif)
		}

		if strings.TrimSpace(item.Sip) != "" {
			if strings.Contains(item.Sip, ",") {
				line += fmt.Sprintf("ip saddr { %s } ", item.Sip)
			} else {
				line += fmt.Sprintf("ip saddr %s ", item.Sip)
			}
		}

		if strings.TrimSpace(item.Dip) != "" {
			if strings.Contains(item.Dip, ",") {
				line += fmt.Sprintf("ip daddr { %s } ", item.Dip)
			} else {
				line += fmt.Sprintf("ip daddr %s ", item.Dip)
			}
		}

		if strings.TrimSpace(item.Snat) != "" {
			line += fmt.Sprintf(" snat ip to %s ", item.Snat)
		} else {
			line += " masquerade"
		}

		g.Log().Debug(ctx, line)

		ruleList = append(ruleList, fmt.Sprintf("add rule inet web-firewall %s %s", ChainName[SNAT], line))
	}

	// todo 目的地址转换
	var dnatList []model.DnatRulesets
	err = dao.DnatRules.Ctx(ctx).OrderAsc(dao.DnatRules.Columns().Position).Scan(&dnatList)
	if err != nil {
		return err
	}

	for _, item := range dnatList {
		line := ""
		if strings.TrimSpace(item.Iif) != "" {
			line += fmt.Sprintf("iifname %s ", item.Iif)
		}

		if strings.TrimSpace(item.Dip) != "" {
			if strings.Contains(item.Dip, ",") {
				line += fmt.Sprintf("ip daddr { %s } ", item.Dip)
			} else {
				line += fmt.Sprintf("ip daddr %s ", item.Dip)
			}
		}
		// 判断port是否为空
		if len(item.Port) == 0 {
			// add rule inet web-firewall m_dnat ip daddr 192.168.1.200 dnat ip to 1.2.2.2
			ruleList = append(ruleList, fmt.Sprintf("add rule inet web-firewall %s %s", ChainName[DNAT], line+fmt.Sprintf("dnat ip to %s", item.Dnat)))
			continue
		}

		// 兼容1.3.0版本
		for _, port := range item.Port {
			// 判断port是否有protocol，有就用 没有就使用全局的旧版本数据
			tmpProtocol := strings.TrimSpace(port.Protocol)
			if strings.TrimSpace(port.Protocol) == "" {
				tmpProtocol = strings.TrimSpace(item.Protocol)
			}
			if tmpProtocol == "tcp" || tmpProtocol == "udp" {
				ruleList = append(ruleList, fmt.Sprintf("add rule inet web-firewall %s %s", ChainName[DNAT], line+fmt.Sprintf("%s dport %v  dnat ip to %s:%v", tmpProtocol, port.Key, item.Dnat, port.Value)))
			} else {
				ruleList = append(ruleList, fmt.Sprintf("add rule inet web-firewall %s %s", ChainName[DNAT], line+fmt.Sprintf("tcp dport %v  dnat ip to %s:%v", port.Key, item.Dnat, port.Value)))
				ruleList = append(ruleList, fmt.Sprintf("add rule inet web-firewall %s %s", ChainName[DNAT], line+fmt.Sprintf("udp dport %v  dnat ip to %s:%v", port.Key, item.Dnat, port.Value)))

			}

		}
	}
	// todo 转发策略
	var forwardList []entity.ForwardRules
	err = dao.ForwardRules.Ctx(ctx).OrderAsc(dao.ForwardRules.Columns().Position).Scan(&forwardList)
	if err != nil {
		return err
	}

	for _, item := range forwardList {
		line := ""
		if strings.TrimSpace(item.Sip) != "" {
			if strings.Contains(item.Sip, ",") {
				line += fmt.Sprintf("ip saddr { %s } ", item.Sip)
			} else {
				line += fmt.Sprintf("ip saddr %s ", item.Sip)
			}
		}

		if strings.TrimSpace(item.Dip) != "" {
			if strings.Contains(item.Dip, ",") {
				line += fmt.Sprintf("ip daddr { %s } ", item.Dip)
			} else {
				line += fmt.Sprintf("ip daddr %s ", item.Dip)
			}
		}

		if item.Protocol == "tcp" || item.Protocol == "udp" {
			if strings.Contains(item.Port, ",") {
				line += fmt.Sprintf("%s dport { %s } ", item.Protocol, item.Port)
			} else {
				line += fmt.Sprintf("%s dport %s ", item.Protocol, item.Port)
			}
		}
		line += fmt.Sprintf(" %s", item.Policy)

		g.Log().Debug(ctx, line)

		ruleList = append(ruleList, fmt.Sprintf("add rule inet web-firewall %s %s", ChainName[FORWARD], line))
	}

	// todo 转发流控

	var forwardLimitList []entity.ForwardLimitRules
	err = dao.ForwardLimitRules.Ctx(ctx).OrderAsc(dao.ForwardLimitRules.Columns().Position).Scan(&forwardLimitList)
	if err != nil {
		return err
	}

	for _, item := range forwardLimitList {
		line := ""

		if strings.TrimSpace(item.Sip) != "" {
			if strings.Contains(item.Sip, ",") {
				line += fmt.Sprintf("ip saddr { %s } ", item.Sip)
			} else {
				line += fmt.Sprintf("ip saddr %s ", item.Sip)
			}
		}

		if strings.TrimSpace(item.Dip) != "" {
			if strings.Contains(item.Dip, ",") {
				line += fmt.Sprintf("ip daddr { %s } ", item.Dip)
			} else {
				line += fmt.Sprintf("ip daddr %s ", item.Dip)
			}
		}

		if item.Protocol == "tcp" || item.Protocol == "udp" {
			if strings.Contains(item.Port, ",") {
				line += fmt.Sprintf("%s %s { %s } ", item.Protocol, item.PortType, item.Port)
			} else {
				line += fmt.Sprintf("%s %s %s ", item.Protocol, item.PortType, item.Port)
			}
		}

		g.Log().Debug(ctx, line)

		line += fmt.Sprintf(" limit rate over %d %s drop", item.Limit, speedMap[item.Speed])

		ruleList = append(ruleList, fmt.Sprintf("add rule inet web-firewall %s %s", ChainName[LIMIT_FORWARD], line))
	}

	g.Log().Debug(ctx, basicStr, strings.Join(ruleList, "\n"))
	err = exec.Command("sh", "-c", fmt.Sprintf("nft -f - <<EOF\n%s\n%s\nEOF", basicStr, strings.Join(ruleList, "\n"))).Run()
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	return nil
}
