package gm

import "testing"

/*
测试创建sm2密钥对
*/
func TestSaveFileKey(t *testing.T) {
	err := SaveFileKey()
	if err != nil {
		t.Fatal("SaveFileKey 执行错误:", err)
		return
	}
	t.Log("SaveFileKey 执行正确")
}
