package utils

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/ripemd160"
)

//产生一个新的私钥
func GenerateKey() (*ecdsa.PrivateKey, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, errors.New("failed GenerateKey")
	}
	return key, nil
}

//使用私钥进行签名
func SignText(digestHash string, prv *ecdsa.PrivateKey) ([]byte, error) {
	msg := crypto.Keccak256([]byte(digestHash))
	sig, err := crypto.Sign(msg, prv)
	return sig, err
}

//通过公钥推导出地址
func GetAddressByPublicKey(key ecdsa.PublicKey) string {
	address := crypto.PubkeyToAddress(key)
	return  address.String()
}

func Base58Encode(message []byte) string {
	return base58.Encode(message)
}

func Base58Decode(msgStr string) []byte {
	return base58.Decode(msgStr)
}

func GetSHA256HashCode(message []byte) []byte {
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write(message)
	//计算哈希值
	bytes := hash.Sum(nil)
	//返回哈希值
	return bytes
}

func GetRipemd160HashCode(message []byte) []byte {
	hasher := ripemd160.New()
	hasher.Write(message)
	hashBytes := hasher.Sum(nil)
	return hashBytes
}

//使用原数据和签名恢复公钥地址
func VerifyToAddress(digestHash string, sig []byte) (string, error) {
	msg := crypto.Keccak256([]byte(digestHash))
	recoveredPub, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return "", err
	}
	recoveredAddr := crypto.PubkeyToAddress(*recoveredPub)
	return recoveredAddr.String(), nil
}
