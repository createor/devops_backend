package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

// 公钥
var publicKey = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvrfNL+kIizOIkER8Xi/8
qOpM7KBglv3yVxG8R8y/tjvGYpV1Hq5fSNnr6TiVnOuT06aIAZL8opz9TtEAr1fr
JEs95RzU62dOHQSqost79uK2vrLzNAledMzpi5E/+Ft6k2QhEOqloxk3bjVja2/j
Je97grF37Y3VeBZOeQOpXQzZC98WuKkeMybArKo/6SlzefQtdjPQ3zDGWUR91Z9Q
DQXdFI+A+Edbi/3XLijKOZ+r0FSu2tio5Tu01Nvj4eIicaLDg0Q2kwneOIc2ub89
9IMtesxWnLBsuCxPkt7UmtNrzlhHLktqiG5JRLE201DwDMYoQmQ4cdLPV3/BIpho
uwIDAQAB
-----END PUBLIC KEY-----
`

// 私钥
var privateKey = `
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC+t80v6QiLM4iQ
RHxeL/yo6kzsoGCW/fJXEbxHzL+2O8ZilXUerl9I2evpOJWc65PTpogBkvyinP1O
0QCvV+skSz3lHNTrZ04dBKqiy3v24ra+svM0CV50zOmLkT/4W3qTZCEQ6qWjGTdu
NWNrb+Ml73uCsXftjdV4Fk55A6ldDNkL3xa4qR4zJsCsqj/pKXN59C12M9DfMMZZ
RH3Vn1ANBd0Uj4D4R1uL/dcuKMo5n6vQVK7a2KjlO7TU2+Ph4iJxosODRDaTCd44
hza5vz30gy16zFacsGy4LE+S3tSa02vOWEcuS2qIbklEsTbTUPAMxihCZDhx0s9X
f8EimGi7AgMBAAECggEAW517qbot6oCU54ienbg7jQEQdtML0zymP4E7itomden9
ALp/CoAFMb/Nfbk61ais2I005Fyxk4QKguQPiiuXv1WNpBPXjEWR1oq5VX6eTBjY
ZH8eKS8e+si8n9jke++l0EvXPoMZkmG4qO5oleGnoj+Ke1u5Gpp5ozhD9gn2P8Xo
rJjZTr9oWcXq3UhfRLvKzSrFkYgnQM/T7cSqtvqcT4VUzVjpfuGuttPzL392ipCC
TVpelUe0IiiSp7uuCQZMs39Q3KBa4OAwhovYdhFZKO40u47n5GJJT5qrAiTb2hcQ
bK1/sdU45inuTmYqApuUvBD4R++mb7jei/sCWgZsEQKBgQD2PGPh7vE+u0G0m1QO
1YeNW4H7edz28dDoE+AI/9a53IuoBB43Cxz6ViqlE4zuWX9gfTkNY9xBmGKsuvv6
gN7D3SHObJ4MLD4LYppqO2VpUlERLQAJn+b2L3t/CqEfFhsD7fdwRiqGZsaIR9cr
RoRnk3R2nv/M8E1t8efx1pXjlwKBgQDGR9VSW5wzPz9iZk1bNK/fcXUNNkVxdeEn
uB/3xgN/fBq9Q3Z0IZrgQTQwkDxYLOV+2QfKQYKQXbUAfMTRpSdMFyV5PoRG6tQ0
Ovcy+BU9CpxtHjLL2L8CGyyOh354KmtUefxDsnfxQssGabKrHmxsLb8DQZXn8PZo
bpLtlkL4fQKBgQCKWBGyPcJTAXiAFYkbsIKxPAmClcw8/k3mJkyIId2tnSjl5DJp
sJe+Wq0pBBv5SlVTi+eDC2kTfZ9q9r9d1gvStaopxULjCfRuBx9Eskxe6T3czZCo
16s3BCR5kypFQfE5uvh7nyCDVLkUlnBgwwTfAKy9fMWxig2myPQNHwglzwKBgFjz
tu7IrG2NLUlercuB+niadLGlrEe3Y3gnMSg+DCmwKmrIDicRQGLkvZ4fxwKjuZ1L
jiQdeY58i4wZbU7D8bpAFA6tjjgmd2arIWUbSKPm08BcMNukdCRkvnt+q60LErWG
ODbCpO52UZCh8Ia2ElwBtdSnIrI4NsMo//9YTtkdAoGAQpVwgginKuOXb91l6GKH
IVEIYH64YTWOglRy7MQKmRS/11l9hgJ2TW4YnYHWLLFdEflfJIggUJFMhuF8Mj9b
KJVORy9qtmGqhOtkgHJ216QW4JSBMVBTQIOGyTDTpdFSUGVUQh5YkpPxSyrdKV7Z
fRQeYYmDTgs/2eveEfgmRyQ=
-----END PRIVATE KEY-----
`

// GetPublicKey 获取公钥
// 返回:
//
//	string: 公钥
func GetPublicKey() string {
	return publicKey
}

// DecryptData rsa解密函数
//
// 参数:
//
//	secret: string, 密文
//
// 返回:
//
//	string: 解密后的明文
//	error: 错误
func DecryptData(secret string) (string, error) {
	block, _ := pem.Decode([]byte(privateKey)) // 读取私钥
	if block == nil || block.Type != "PRIVATE KEY" {
		return "", fmt.Errorf("错误的私钥")
	}
	pkcs8Key, err := x509.ParsePKCS8PrivateKey(block.Bytes) // 格式转换
	if err != nil {
		return "", err
	}
	rsaPrivateKey, ok := pkcs8Key.(*rsa.PrivateKey) // 类型断言
	if !ok {
		return "", fmt.Errorf("类型转换失败")
	}
	cipherText, err := base64.StdEncoding.DecodeString(secret) // base64解码
	if err != nil {
		return "", fmt.Errorf("base64解码失败")
	}
	plainText, err := rsa.DecryptPKCS1v15(nil, rsaPrivateKey, cipherText) // rsa解密
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}
