package utils

import (
	"fmt"
	"testing"
)

func TestVerifyToAddress(t *testing.T) {
	prk, err := GenerateKey()
	if err != nil {
	t.Error(err.Error())
	return
	}

	text := "caohaitao"
	result, err := SignText(text, prk)
	if err != nil {
	t.Error(err.Error())
	return
	}

	addr, err := VerifyToAddress(text, result)
	if err != nil {
	t.Error(err.Error())
	return
	}
	addrRaw := GetAddressByPublicKey(prk.PublicKey)
	fmt.Println(addrRaw)
	fmt.Println(addr)
}

func TestGetRipemd160HashCode(t *testing.T) {
	strByte := "peter"
	hashCode := GetRipemd160HashCode([]byte(strByte))
	fmt.Printf("%x", hashCode)
	t.Log("GetSHA256HashCode pass")
}

func TestGetSHA256HashCode(t *testing.T) {
	strByte := "peter"
	hashCode := GetSHA256HashCode([]byte(strByte))
	fmt.Printf("%x", hashCode)
	t.Log("GetSHA256HashCode pass")
}


func TestBase58Encode(t *testing.T) {
	str := Base58Encode([]byte("peter"))
	fmt.Println(str)
	t.Log("Base58Encode pass")
}

func TestBase58Decode(t *testing.T)  {
	strByte := Base58Decode("DgUwp2V")
	fmt.Println(string(strByte))
	t.Log("Base58Decode pass")
}
