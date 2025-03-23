package gm

import "testing"

/*
测试sm2加解密功能
*/

func TestSM3(t *testing.T) {
	slat := "0XQaLdUJJg4nKHoj7M8Ax"
	password := SM3("admin", slat)
	t.Log(password)
}

func TestSM2(t *testing.T) {
	data := "admin"
	encrypt, err := Sm2Encrypt(data)
	if err != nil {

		t.Error(err)
		return
	}
	t.Log(encrypt)

	decode, err := Sm2Decode(encrypt)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(decode)
}
