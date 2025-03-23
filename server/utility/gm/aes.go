package gm

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

//高级加密标准（Adevanced Encryption Standard ,AES）

// 16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法
// key不能泄露

var PwdKey []byte
var blockSize = 16

func init() {
	ctx := gctx.New()
	aesKey, err := g.Cfg().Get(ctx, "encrypt.aesKey", "y+FZNFOjFBS45ur67sLeRCbHv@we^%h8")
	if err != nil {
		return
	}
	PwdKey = aesKey.Bytes()
}

// PKCS7Padding PKCS7 填充模式
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding 填充的反向操作，删除填充字符串
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	} else {
		//获取填充字符串长度
		unpadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文
		return origData[:(length - unpadding)], nil
	}
}

// AesEcrypt 实现加密
func AesEcrypt(origData []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(PwdKey)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	//blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	origData = PKCS7Padding(origData, blockSize)
	//采用AES加密方法中CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block, PwdKey[:blockSize])
	crypted := make([]byte, len(origData))
	//执行加密
	blocMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// AesDeCrypt 实现解密
func AesDeCrypt(cypted []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(PwdKey)
	if err != nil {
		return nil, err
	}
	//获取块大小
	//blockSize := block.BlockSize()
	//fmt.Println(key[:blockSize])
	//创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, PwdKey[:blockSize])
	origData := make([]byte, len(cypted))
	//这个函数也可以用来解密
	blockMode.CryptBlocks(origData, cypted)
	//去除填充字符串
	origData, err = PKCS7UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, err
}

// DecodeAES AES解密 CBC模式 iv=key[:16]
func DecodeAES(pwd string) ([]byte, error) {
	pwdByte, err := hex.DecodeString(pwd)
	if err != nil {
		return nil, err
	}
	//执行AES解密
	return AesDeCrypt(pwdByte)

}

// EncodeAES AES 加密 CBC模式 iv=key[:16]
func EncodeAES(content []byte) (string, error) {
	result, err := AesEcrypt(content)
	if err != nil {
		return "", err
	}
	str := hex.EncodeToString(result)
	return str, err
}
