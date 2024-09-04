package nftables

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

/**
用来定义各个对象
*/

type Table struct {
	Name    string `json:"name"` // 仅支持英文
	Comment string `json:"comment"`
	Family  string `json:"family"` //ip ip6 inet arp bridge netdev
}

func (t *Table) ToString() string {
	res := t.Family + " " + t.Name
	if t.Comment != "" {
		res += fmt.Sprintf(` { comment "%s" ; }`, t.Comment)
	}

	return res
}

// AddTable 创建表/激活表，表存在不会报错
func (t *Table) AddTable() error {
	command := fmt.Sprintf("add table %s", t.ToString())
	//fmt.Println(command)
	cmd := exec.Command("nft", "-e", "-a", command)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		return err
	}
	return nil
}

// CreateTable 创建表，表存在会报错，在程序运行时建议使用此函数，方便获取创建时的handle，如果使用add则导致没有打印handle的意外情况
func (t *Table) CreateTable() (handle string, err error) {
	command := fmt.Sprintf("create table %s", t.ToString())
	//fmt.Println(command)
	cmd := exec.Command("nft", "-e", "-a", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(err.Error() + "\n" + string(output))
	}
	// 如果没有报错，则解析handle
	results := strings.Split(string(output), "\n")[0]
	re := regexp.MustCompile(`(\d+)\s*$`)
	handleByte := re.Find([]byte(results))
	if handleByte != nil {
		return string(handleByte), nil
	}
	return "", errors.New("no handle found")
}

// 删除表，程序启动先删除表，避免表与链属性冲突
func (t *Table) DeleteTable() error {
	command := fmt.Sprintf("delete table %s %s", t.Family, t.Name)
	cmd := exec.Command("nft", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(err.Error() + "\n" + string(output))
	}
	return nil
}

// Flush 清空表中所有的rule，不会删除chain，但是会删除rule
func (t *Table) Flush() error {
	cmd := exec.Command("nft", "flush", "table", t.Family, t.Name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(err.Error() + "\n" + string(output))
	}
	return nil
}

func DisableTable(handle int) error {
	return nil
}
