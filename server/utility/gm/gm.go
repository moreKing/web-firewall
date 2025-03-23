package gm

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"strings"
	"unsafe"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"github.com/tjfoc/gmsm/sm4"
	"github.com/tjfoc/gmsm/x509"
)

var (
	sm4Key = []byte("more!@#$1234qwer")

	priKey *sm2.PrivateKey
	pubkey *sm2.PublicKey
)

func init() {
	publicKey := `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEw+8enwmcH9decrv6jFEeCqJPOIc5
ZFB0G0Q+miahGKpowWZCD2II+zYnkqMPVE7C000asKepgIKXNXc4Srv2sw==
-----END PUBLIC KEY-----`

	privateKey := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgl+63hqBrq6lnuMcS
3Ovbig6RvEJxXSUTtEupQf4yOsCgCgYIKoEcz1UBgi2hRANCAATD7x6fCZwf115y
u/qMUR4Kok84hzlkUHQbRD6aJqEYqmjBZkIPYgj7NieSow9UTsLTTRqwp6mAgpc1
dzhKu/az
-----END PRIVATE KEY-----`

	var err error
	ctx := gctx.New()
	priKey, err = x509.ReadPrivateKeyFromPem([]byte(privateKey), nil)
	if err != nil {
		//log.Fatal(err)
		g.Log().Fatal(ctx, err)
	}
	pubkey, err = x509.ReadPublicKeyFromPem([]byte(publicKey))
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

}

// sm4加密
func Sm4Ecb(data string) []byte {
	ecbMsg, err := sm4.Sm4Ecb(sm4Key, []byte(data), true)
	if err != nil {
		// global.UMI_LOG.Error(err)
		return nil
	}
	return ecbMsg
}

// sm4解密
func Sm4Dec(data []byte) string {
	if len(data) == 0 {
		return ""
	}
	ecbMsg, err := sm4.Sm4Ecb(sm4Key, data, false)
	if err != nil {
		// global.UMI_LOG.Error(err)
		return ""
	}
	return string(ecbMsg)
}

// SM3加密，slat为随机生成的盐值，无解密
func SM3(data string, slat string) string {
	// datas, _ := hex.DecodeString(data)
	h := sm3.New()
	h.Write([]byte(data)) //解密
	sum := h.Sum([]byte(slat))
	cipherStr := hex.EncodeToString(sum)
	return cipherStr
}

// sm4加密
func Sm4EcbCmdLine(data []byte) string {
	ecbMsg, err := sm4.Sm4Ecb([]byte("moujun!@#%123asd"), data, true)
	if err != nil {
		// global.UMI_LOG.Error(err)
		return ""
	}
	return string(ecbMsg)
}

// Sm2Decode  解密
/**
 * @Description: 解密(私钥解密)
 * @param privateKey 私钥
 * @param cipherStr 加密后的字符串
 * @return data 解密后的数据
 * @return err
 */
func Sm2Decode(cipherStr string) (data string, err error) {

	if strings.TrimSpace(cipherStr) == "" {

		return "", nil
	}
	// 16进制字符串转[]byte
	bytes, _ := hex.DecodeString(cipherStr)
	// sm2解密
	var dataByte []byte
	dataByte, err = priKey.DecryptAsn1(bytes)
	if err != nil {
		return data, err
	}
	// byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&dataByte))
	return *str, err
}

// Sm2Encrypt  加密
/**
 * @Description: SM2加密(公钥加密)
 * @param publicKey 公钥
 * @param data 需要加密的数据
 * @return cipherStr 加密后的字符串
 */
func Sm2Encrypt(data string) (cipherStr string, err error) {
	if data == "" {
		return "", nil
	}
	// 将字符串转为[]byte
	dataByte := []byte(data)
	// sm2加密
	cipherTxt, err := pubkey.EncryptAsn1(dataByte, rand.Reader)
	if err != nil {
		return
	}
	// 转为16进制字符串输出
	//cipherStr = fmt.Sprintf("%x", cipherTxt)
	cipherStr = hex.EncodeToString(cipherTxt)
	return
}
