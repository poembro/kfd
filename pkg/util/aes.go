package util

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/base64"
    "encoding/json"
    "encoding/pem"
    "errors"
    //"kfd/pkg/logger"

    jsoniter "github.com/json-iterator/go"
)

type TokenInfo struct {
    AppId    int64 `json:"app_id"`    // appId
    UserId   int64 `json:"user_id"`   // 用户id
    Account string `json:"account"` // 设备id
    Expire   int64 `json:"expire"`    // 过期时间
}

// GetToken 获取token
func GetToken(appId, userId int64, account string, expire int64, publicKey string) (string, error) {
    info := TokenInfo{
        AppId:    appId,
        UserId:   userId,
        Account: account,
        Expire:   expire,
    }
    bytes, err := json.Marshal(info)
    if err != nil {
        //logger.Sugar.Error(err)
        return "", err
    }

    token, err := RsaEncrypt(bytes, []byte(publicKey))
    if err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(token), nil
}

// DecryptToken 对加密的token进行解码
func DecryptToken(token string, privateKey string) (*TokenInfo, error) {
    bytes, err := base64.StdEncoding.DecodeString(token)
    if err != nil {
        //logger.Sugar.Error(err)
        return nil, err
    }
    result, err := RsaDecrypt(bytes, Str2bytes(privateKey))
    if err != nil {
        //logger.Sugar.Error(err)
        return nil, err
    }

    var info TokenInfo
    err = jsoniter.Unmarshal(result, &info)
    if err != nil {
        //logger.Sugar.Error(err)
        return nil, err
    }
    return &info, nil
}

// 加密
func RsaEncrypt(origData []byte, publicKey []byte) ([]byte, error) {
    //解密pem格式的公钥
    block, _ := pem.Decode(publicKey)
    if block == nil {
        return nil, errors.New("public key error")
    }
    // 解析公钥
    pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return nil, err
    }
    // 类型断言
    pub := pubInterface.(*rsa.PublicKey)
    //加密
    return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte, privateKey []byte) ([]byte, error) {
    //解密
    block, _ := pem.Decode(privateKey)
    if block == nil {
        return nil, errors.New("private key error!")
    }
    //解析PKCS1格式的私钥
    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        return nil, err
    }
    // 解密
    return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}


// PrivateKey 私钥
var PrivateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCvkih3KRdHG6TgZv8cQ7wCICbzYuFOBkjjZatTjNIJleCK51FA
lYvGuh00on4bubIfa47bMadPkRiQ9udAv1BJ3C3cmEKMJc9Y3EpEA/ZLDEPuO2QU
T6k8wV5Au6N8bxciLNHTrzOH6qQQ8V4xQ++sPS2dqDkRqLN59SxFjJJlNwIDAQAB
AoGAHSOX3bcPKvkWkzvk6U8AnCWz8T8e/7EhVcRg+/vqPDcIzmT34k0vpqrq//pc
DcPzIS3bxttl1lnRhvXDicZ2z5T81SG9fQWzGOewLzh71y3IHxXmM/J9ZPhx9w9V
yzVw47lUl5jUm4+NtjpUGwiSxlIu5slhGP2dezIT8Xyd/HkCQQDYm4/dh2qbP2Nj
hW0MUN0aJ+LcEwyjxzYn/T1iOrHU3veAsWuTTLzJ9KV+5LCv5o/zgnIAsbgXwKAf
n3EQt9ANAkEAz4AUgx/k0HAHq4ijFBmEODdHUVpaknTC498hTasaUCBGHCxvazBt
l5sQ+WXpcDHmlG/cWvgaBqoDAkPa/XR1UwJAeHH4C3zzQKR8xag5vPFyIMsxEKLf
EmsBqDbe3TI6FF5vTfZaFSxEhiAtrmPIA+e2//b3IX+xGDQaVbs5CczMOQJBALy8
2GuIqBz7uc5Jw2P17bEgfss9ryKF9/tDKsy1tIJrSLo5pMLP0u9SSbc//nxht0UP
FlTcVPf4zjHSPK/LrScCQBWzzDOZAQFiseCH2TKjZJKcQTphTj6dg6mx12lwg4u1
0QTi+9QfOy9m2L1M2uZZ8s2iEP7jXvqltQ9cV9CEtAs=
-----END RSA PRIVATE KEY-----
`

// 公钥: 根据私钥生成
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var PublicKey = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCvkih3KRdHG6TgZv8cQ7wCICbz
YuFOBkjjZatTjNIJleCK51FAlYvGuh00on4bubIfa47bMadPkRiQ9udAv1BJ3C3c
mEKMJc9Y3EpEA/ZLDEPuO2QUT6k8wV5Au6N8bxciLNHTrzOH6qQQ8V4xQ++sPS2d
qDkRqLN59SxFjJJlNwIDAQAB
-----END PUBLIC KEY-----
`


/*
#利用git的openssl模块 生成rsa 

##双击bash.exe 弹出窗口
### D:\soft\git\Git\bin\bash.exe

##切换到指定目录
cd C:\Users\Administrator\rsa-test-golang

##生成私钥
openssl genrsa -out rsa_private_key.pem 1024

##利用私钥生成公钥
openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
*/