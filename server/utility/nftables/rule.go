package nftables

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"os/exec"
	"regexp"
	"strings"
)

const (
	_ = iota
	ADD
	REPLACE
	DELETE
)

type Rule struct {
	Chain    string       `json:"chain"`
	Handle   string       `json:"handle"`
	Add      bool         `json:"add"`      // 使用add 还是insert
	Position int          `json:"position"` // 添加规则时插入的位置
	Expr     []Expression `json:"expr"`
}

// Expression 一个简单的表达式,此处为了本项目服务，如需要更灵活的请自行定义结构体
type Expression struct {
	Type     string `json:"type"`     // match or policy
	Protocol string `json:"protocol"` // meta tcp udp ip...
	Field    string `json:"field"`    // saddr sport dport ...
	Value    string `json:"value"`    //  192.168.1.1、{ 192.168.1.1,192.168.1.1-192.168.1.5}
	Policy   string `json:"policy"`   // accept 、drop 、 reject
}

func AddRule(ctx context.Context, rule *Rule) (r *Rule, err error) {
	var command string
	if rule.Add {
		command = fmt.Sprintf("add %s", rule.ToCommand(ADD))
	} else {
		command = fmt.Sprintf("insert %s", rule.ToCommand(ADD))
	}
	//fmt.Println(command)
	g.Log().Debug(ctx, command)
	cmd := exec.Command("nft", "-e", "-a", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.New(err.Error() + "\n" + string(output))
	}
	// 如果没有报错，则解析handle
	results := strings.Split(string(output), "\n")[0]
	re := regexp.MustCompile(`(\d+)\s*$`)
	handleByte := re.Find([]byte(results))
	if handleByte != nil {
		rule.Handle = string(handleByte)
		return rule, nil
	}
	return nil, errors.New("no handle found")
}

func DeleteRule(rule *Rule) (err error) {
	command := fmt.Sprintf("delete %s", rule.ToCommand(DELETE))
	//fmt.Println(command)
	cmd := exec.Command("nft", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(err.Error() + "\n" + string(output))
	}
	return nil
}

func ReplaceRule(rule *Rule) (err error) {
	command := fmt.Sprintf("replace %s", rule.ToCommand(REPLACE))
	//fmt.Println(command)
	cmd := exec.Command("nft", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(err.Error() + "\n" + string(output))
	}
	return nil
}

// ToCommand 将rule转成命令模式
func (r *Rule) ToCommand(mod int) string {

	switch mod {
	case ADD:
		res := fmt.Sprintf("rule %s %s %s ", tableName.Family, tableName.Name, r.Chain)

		if r.Position > 0 {
			res += fmt.Sprintf(" handle  %d ", r.Position)
		}

		for _, expression := range r.Expr {
			if expression.Type == "match" {
				//  为空就跳过
				if (strings.TrimSpace(expression.Protocol) == "ip" || strings.TrimSpace(expression.Protocol) == "tcp" || strings.TrimSpace(expression.Protocol) == "udp") && strings.TrimSpace(expression.Value) == "" {
					continue
				}

				res += fmt.Sprintf(" %s %s ", expression.Protocol, expression.Field)
				if strings.TrimSpace(expression.Field) == "vmap" || strings.Contains(expression.Value, ",") {
					res += fmt.Sprintf("{ %s } ", expression.Value)
				} else {
					res += fmt.Sprintf(" %s ", expression.Value)
				}

				continue
			}
			res += fmt.Sprintf(" %s ", expression.Policy)
		}
		return res

	case REPLACE:
		res := fmt.Sprintf("rule %s %s %s handle %s ", tableName.Family, tableName.Name, r.Chain, r.Handle)
		for _, expression := range r.Expr {
			if expression.Type == "match" {

				//  为空就跳过
				if (strings.TrimSpace(expression.Protocol) == "ip" || strings.TrimSpace(expression.Protocol) == "tcp" || strings.TrimSpace(expression.Protocol) == "udp") && strings.TrimSpace(expression.Value) == "" {
					continue
				}

				res += fmt.Sprintf(" %s %s ", expression.Protocol, expression.Field)
				if strings.TrimSpace(expression.Field) == "vmap" || strings.Contains(expression.Value, ",") {
					res += fmt.Sprintf("{ %s } ", expression.Value)
				} else {
					res += fmt.Sprintf(" %s ", expression.Value)
				}
				continue
			}
			res += fmt.Sprintf(" %s ", expression.Policy)
		}
		return res

	case DELETE:

		return fmt.Sprintf("rule %s %s %s handle %s ", tableName.Family, tableName.Name, r.Chain, r.Handle)

	}

	return ""

}
