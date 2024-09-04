package gm

import (
	"crypto"
	"crypto/rand"
	"encoding/hex"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

// sm2 签名跟验证签名
var (
	signPublicKey  *sm2.PublicKey
	signPrivateKey *sm2.PrivateKey
)

func init() {
	signPublicKeyStr := `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAE4khh/gwbQAimfO3j9sXR7WMYuH6m
owQDXVDQquD8wCBRbCtV2WaPlajRg0mABQ5CBhwoBF7elqvppM3MsSSWkQ==
-----END PUBLIC KEY-----`

	signPrivateKeyStr := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgZUdy/PeFcDs7hL4e
RM1Fur5Qt7NJqA1loT+Tw04sTVSgCgYIKoEcz1UBgi2hRANCAATiSGH+DBtACKZ8
7eP2xdHtYxi4fqajBANdUNCq4PzAIFFsK1XZZo+VqNGDSYAFDkIGHCgEXt6Wq+mk
zcyxJJaR
-----END PRIVATE KEY-----`

	var err error

	signPrivateKey, err = x509.ReadPrivateKeyFromPem([]byte(signPrivateKeyStr), nil)
	if err != nil {
		panic(err)
	}
	signPublicKey, err = x509.ReadPublicKeyFromPem([]byte(signPublicKeyStr))
	if err != nil {
		panic(err)
	}

}

// Sign
/**
 *  @Description: 签名
 *  @param privateKey 私钥
 *  @param msg 需要签名的内容
 *  @param signer
 *  @return sign 签名字符串
 *  @return err
 */
func SM2Sign(msg []byte) (sign string, err error) {
	var signByte []byte
	// sm2签名
	signByte, err = signPrivateKey.Sign(rand.Reader, msg, crypto.SHA256)
	if err != nil {
		return "", err
	}
	// 转为16进制字符串输出
	sign = hex.EncodeToString(signByte)
	return sign, nil
}

// Verify
/**
 *  @Description: 验签
 *  @param publicKey 公钥
 *  @param msg 需要验签的内容
 *  @param sign 签名字符串
 *  @return verify
 */
func SM2SignVerify(msg, sign string) (verify bool) {
	// 16进制字符串转[]byte
	msgBytes := []byte(msg)
	signBytes, _ := hex.DecodeString(sign)
	// sm2验签
	verify = signPublicKey.Verify(msgBytes, signBytes)
	return
}
