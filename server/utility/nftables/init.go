package nftables

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var tableName = &Table{
	Name: "web-firewall",
	//Comment: "web防火墙系统",
	Family: "inet",
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
	LIMIT_FORWARD: "m_limit_forwarded",
}

func init() {

	_ = tableName.DeleteTable()
	//	 创建表
	err := tableName.AddTable()
	if err != nil {
		panic(err)
	}

	//	创建链
	//入站策略 input链
	inputPriority := g.Cfg().MustGet(context.Background(), "firewall.chainPriority.input", 100).Int()
	inputBasic := Chain{
		Name:   ChainName[INPUT_BASIC],
		Table:  tableName.Name,
		Family: tableName.Family,
		//Comment:  "入站规则策略",
		Type:     "filter",
		Hook:     "input",
		Priority: inputPriority,
		Policy:   "drop",
	}
	err = inputBasic.Add()
	if err != nil {
		panic(err)
	}

	// 连接数/流量 控制 input链/output链/forward 链 优先级高 目前不做forward流控，仅做上传/下载 即 出站/入站 策略流控

	limitIn := Chain{
		Name:   ChainName[LIMIT_INPUT],
		Table:  tableName.Name,
		Family: tableName.Family,
		//Comment: "入站流量控制(限制客户端上传速度/包数量)",
	}
	err = limitIn.Add()
	if err != nil {
		panic(err)
	}

	// 添加入站策略到基础链
	_, err = AddRule(context.Background(), &Rule{
		Chain:    ChainName[INPUT_BASIC],
		Add:      true,
		Position: 0,
		Expr: []Expression{
			{
				Type:     "match",
				Protocol: "jump",
				Field:    ChainName[LIMIT_INPUT],
			},
		},
	})
	if err != nil {
		panic(err)
	}

	// 入站默认常规链 优先级最高

	// 添加一条策略，运行本地lo所有访问
	_, err = AddRule(context.Background(), &Rule{
		Chain:    ChainName[INPUT_BASIC],
		Add:      false,
		Position: 0,
		Expr: []Expression{
			{
				Type:     "match",
				Protocol: "meta",
				Field:    "iif",
				Value:    "1",
			}, {
				Type:   "policy",
				Policy: "accept",
			},
		},
	})
	if err != nil {
		panic(err)
	}

	// 入站默认策略
	input := Chain{
		Name:   ChainName[INPUT],
		Table:  tableName.Name,
		Family: tableName.Family,
		//Comment: "入站规则策略",
	}
	err = input.Add()
	if err != nil {
		panic(err)
	}

	// 添加入站策略到基础链
	_, err = AddRule(context.Background(), &Rule{
		Chain:    ChainName[INPUT_BASIC],
		Add:      true,
		Position: 0,
		Expr: []Expression{
			{
				Type:     "match",
				Protocol: "jump",
				Field:    ChainName[INPUT],
			},
		},
	})
	if err != nil {
		panic(err)
	}

	// 出站策略 output链
	outputPriority := g.Cfg().MustGet(context.Background(), "firewall.chainPriority.output", 100).Int()
	outputBasic := Chain{
		Name:   ChainName[OUTPUT_BASIC],
		Table:  tableName.Name,
		Family: tableName.Family,
		//Comment:  "出站规则策略",
		Type:     "filter",
		Hook:     "output",
		Priority: outputPriority,
		Policy:   "accept",
	}
	err = outputBasic.Add()
	if err != nil {
		panic(err)
	}

	limitOUT := Chain{
		Name:   ChainName[LIMIT_OUTPUT],
		Table:  tableName.Name,
		Family: tableName.Family,
		//Comment: "出站流量控制(限制客户端下载速度/包数量)",
	}
	err = limitOUT.Add()
	if err != nil {
		panic(err)
	}

	// 添加入站策略到基础链
	_, err = AddRule(context.Background(), &Rule{
		Chain:    ChainName[OUTPUT_BASIC],
		Add:      true,
		Position: 0,
		Expr: []Expression{
			{
				Type:     "match",
				Protocol: "jump",
				Field:    ChainName[LIMIT_OUTPUT],
			},
		},
	})
	if err != nil {
		panic(err)
	}

	output := Chain{
		Name:   ChainName[OUTPUT],
		Table:  tableName.Name,
		Family: tableName.Family,
		//Comment: "出站规则策略",
	}
	err = output.Add()
	if err != nil {
		panic(err)
	}

	_, err = AddRule(context.Background(), &Rule{
		Chain:    ChainName[OUTPUT_BASIC],
		Add:      true,
		Position: 0,
		Expr: []Expression{
			{
				Type:     "match",
				Protocol: "jump",
				Field:    ChainName[OUTPUT],
			},
		},
	})
	if err != nil {
		panic(err)
	}

	//DNAT prerouting链
	preroutingPriority := g.Cfg().MustGet(context.Background(), "firewall.chainPriority.prerouting", 100).Int()
	dnat := Chain{
		Name:   ChainName[DNAT],
		Table:  tableName.Name,
		Family: tableName.Family,
		//Comment:  "目的地址转换",
		Type:     "nat",
		Hook:     "prerouting",
		Priority: preroutingPriority,
		Policy:   "accept",
	}
	err = dnat.Add()
	if err != nil {
		panic(err)
	}

	//SNAT postrouting
	postroutingPriority := g.Cfg().MustGet(context.Background(), "firewall.chainPriority.postrouting", 100).Int()
	snat := Chain{
		Name:   ChainName[SNAT],
		Table:  tableName.Name,
		Family: tableName.Family,
		//Comment:  "源地址转换",
		Type:     "nat",
		Hook:     "postrouting",
		Priority: postroutingPriority,
		Policy:   "accept",
	}
	err = snat.Add()
	if err != nil {
		panic(err)
	}

	//todo 转发策略（作为网关时）forward 链 暂时不实现
	forwardPriority := g.Cfg().MustGet(context.Background(), "firewall.chainPriority.forward", 100).Int()
	forwardBasic := Chain{
		Name:   ChainName[FORWARD_BASIC],
		Table:  tableName.Name,
		Family: tableName.Family,
		//Comment:  "转发规则策略",
		Type:     "filter",
		Hook:     "forward",
		Priority: forwardPriority,
		Policy:   "accept",
	}
	err = forwardBasic.Add()
	if err != nil {
		panic(err)
	}

	limitForward := Chain{
		Name:   ChainName[LIMIT_FORWARD],
		Table:  tableName.Name,
		Family: tableName.Family,
		//Comment: "转发流量控制(限制客户端下载速度/包数量)",
	}
	err = limitForward.Add()
	if err != nil {
		panic(err)
	}

	// 添加入站策略到基础链
	_, err = AddRule(context.Background(), &Rule{
		Chain:    ChainName[FORWARD_BASIC],
		Add:      true,
		Position: 0,
		Expr: []Expression{
			{
				Type:     "match",
				Protocol: "jump",
				Field:    ChainName[LIMIT_FORWARD],
			},
		},
	})
	if err != nil {
		panic(err)
	}

	forward := Chain{
		Name:   ChainName[FORWARD],
		Table:  tableName.Name,
		Family: tableName.Family,
		//Comment: "转发规则策略",
	}
	err = forward.Add()
	if err != nil {
		panic(err)
	}

	_, err = AddRule(context.Background(), &Rule{
		Chain:    ChainName[FORWARD_BASIC],
		Add:      true,
		Position: 0,
		Expr: []Expression{
			{
				Type:     "match",
				Protocol: "jump",
				Field:    ChainName[FORWARD],
			},
		},
	})
	if err != nil {
		panic(err)
	}
	// ip黑名单 prerouting链 优先级高

	ip_b_w := Chain{
		Name:   ChainName[IP_B_W],
		Table:  tableName.Name,
		Family: tableName.Family,
		//Comment:  "ip地址黑白名单",
		Type:     "filter",
		Hook:     "prerouting",
		Priority: 0,
		Policy:   "accept",
	}
	err = ip_b_w.Add()
	if err != nil {
		panic(err)
	}

}
