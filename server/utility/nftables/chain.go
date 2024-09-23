package nftables

import (
	"fmt"
	"os/exec"
	"strings"
)

// 关于链的一些操作

type Chain struct {
	Name    string `json:"name"`
	Table   string `json:"table"`
	Family  string `json:"family"`
	Comment string `json:"comment"` // 备注
	//	 基本链还是常规链 在这里就有区分了，以下只有基本链才有属性
	Type     string `json:"type"`     // filter nat route
	Hook     string `json:"hook"`     // prerouting postrouting input output forward
	Priority int    `json:"priority"` // 优先级
	Policy   string `json:"policy"`   // accept drop
}

func (c *Chain) Add() error {

	if c.Type == "" {
		command := fmt.Sprintf(`add chain %s %s %s`,
			c.Family, c.Table, c.Name)

		if strings.TrimSpace(c.Comment) != "" {
			command += fmt.Sprintf(` { comment "%s" ; }`, c.Comment)
		}

		cmd := exec.Command("nft", command)
		output, err := cmd.CombinedOutput()
		fmt.Println(string(output))
		if err != nil {
			return err
		}

		return nil

	}

	command := fmt.Sprintf(`add chain %s %s %s { type %s hook %s priority %d ; policy %s; }`,
		c.Family, c.Table, c.Name, c.Type, c.Hook, c.Priority, c.Policy)

	if strings.TrimSpace(c.Comment) != "" {
		command += fmt.Sprintf(` { comment "%s" ; }`, c.Comment)
	}

	cmd := exec.Command("nft", command)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		return err
	}
	return nil
}
