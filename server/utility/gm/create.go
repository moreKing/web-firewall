package gm

import (
	"crypto/rand"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"os"
	"path/filepath"
)

// CreateSM2Key
/**
 *  @Description: 随机生成公私钥
 *  @return privateKey 私钥
 *  @return publicKey 公钥
 *  @return err
 */
func CreateSM2Key() (privateKey *sm2.PrivateKey, publicKey *sm2.PublicKey, err error) {
	// 生成sm2秘钥对
	privateKey, err = sm2.GenerateKey(rand.Reader)
	if err != nil {
		return
	}
	// 进行sm2公钥断言
	publicKey = privateKey.Public().(*sm2.PublicKey)
	return
}

// CreatePrivatePem
/**
 *  @Description: 创建Pem私钥文件
 *  @param privateKey 私钥
 *  @param pwd 密码
 *  @param path Pem私钥文件保存路径
 *  @return err
 */
func CreatePrivatePem(privateKey *sm2.PrivateKey, pwd []byte, path string) (err error) {
	// 将私钥反序列化并进行pem编码
	var privateKeyToPem []byte
	privateKeyToPem, err = x509.WritePrivateKeyToPem(privateKey, pwd)
	if err != nil {
		return err
	}
	// 将私钥写入磁盘
	if path == "" {
		path = "cert/sm2Private.Pem"
	}
	// 获取文件中的路径
	paths, _ := filepath.Split(path)
	err = os.MkdirAll(paths, os.ModePerm)
	if err != nil {
		return err
	}
	var file *os.File
	file, err = os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(privateKeyToPem)
	if err != nil {
		return err
	}
	return
}

// CreatePublicPem
/**
 *  @Description: 创建Pem公钥文件
 *  @param publicKey 公钥
 *  @param path Pem公钥文件保存路径
 *  @return err
 */
func CreatePublicPem(publicKey *sm2.PublicKey, path string) (err error) {
	// 将私钥反序列化并进行pem编码
	var publicKeyToPem []byte
	publicKeyToPem, err = x509.WritePublicKeyToPem(publicKey)
	if err != nil {
		return err
	}
	// 将私钥写入磁盘
	if path == "" {
		path = "cert/sm2Public.Pem"
	}
	// 获取文件中的路径
	paths, _ := filepath.Split(path)
	err = os.MkdirAll(paths, os.ModePerm)
	if err != nil {
		return err
	}
	var file *os.File
	file, err = os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(publicKeyToPem)
	if err != nil {
		return err
	}
	return
}

func SaveFileKey() error {
	privateKey, publicKey, err := CreateSM2Key()
	if err != nil {
		return err
	}
	err = CreatePrivatePem(privateKey, nil, "./privateKey.pem")
	if err != nil {
		return err
	}
	err = CreatePublicPem(publicKey, "./publicKey.pem")
	if err != nil {
		return err
	}
	return nil
}
