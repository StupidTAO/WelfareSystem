package model

import (
	"fmt"
	"testing"
)

func TestProofClaimMarshal(t *testing.T) {
	proof := Proof {
		"did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1",
		"0x5a14c2136084A24853EfF907185f79AE77c36c7c",
		"Secp256k1",
		"02b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71",
	}

	credentialSubject := CredentialSubject {
		"did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1", "short test", "long test", RealNameAuthentication,
	}
	baseProof := ProofClaim{"12345678", "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN2",
		"2021-03-27", "2021-03-27",
		1, credentialSubject, proof,
	}
	strByte, err := ProofClaimMarshal(baseProof)
	if err != nil {
		t.Error("BaseDIDDocMarshal failed")
	}
	fmt.Println(string(strByte))
}

func TestProofClaimUnMarshal(t *testing.T) {
	proof := Proof {
		"did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1",
		"0x5a14c2136084A24853EfF907185f79AE77c36c7c",
		"Secp256k1",
		"02b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71",
	}

	credentialSubject := CredentialSubject {
		"did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1", "short test", "long test", RealNameAuthentication,
	}
	baseProof := ProofClaim{"12345678", "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN2",
		"2021-03-27", "2021-03-27",
		1, credentialSubject, proof,
	}
	strByte, err := ProofClaimMarshal(baseProof)
	if err != nil {
		t.Error("ProofClaimUnMarshal failed")
	}

	//从字符串到结构体
	proofEntry := new(ProofClaim)
	err = ProofClaimUnmarshal(string(strByte), proofEntry)
	if err != nil {
		t.Error("ProofClaimMarshal failed")
	}
}

func TestCredentialSubjectMarshal(t *testing.T) {
	csEntry := new(CredentialSubject)
	csEntry.ID = "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN2"
	csEntry.TypeCliam = IDCardAuthentication
	csEntry.ShortDescription = "342225199509082432"
	csEntry.LongDescription = "身份证号码"
	bs, err := CredentialSubjectMarshal(*csEntry)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(string(bs))
}

func TestCredentialSubjectUnmarshal(t *testing.T) {
	csEntry := new(CredentialSubject)
	csEntry.ID = "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN2"
	csEntry.TypeCliam = IDCardAuthentication
	csEntry.ShortDescription = "342225199509082432"
	csEntry.LongDescription = "身份证号码"
	bs, err := CredentialSubjectMarshal(*csEntry)
	if err != nil {
		t.Error(err.Error())
		return
	}

	csEntry1 := new(CredentialSubject)
	err = CredentialSubjectUnmarshal(string(bs), csEntry1)
	if err != nil {
		t.Error(err.Error())
		return
	}
}
