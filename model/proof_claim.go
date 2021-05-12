package model

import "encoding/json"

const RealNameAuthentication = "RealNameAuthentication"
const IDCardAuthentication = "IDCardAuthentication"
const FingerprintAuthentication = "FingerprintAuthentication"
const EnterpriseAuthentication = "EnterpriseAuthentication"
const BusinessAuthentication = "BusinessAuthentication"
const VIPAuthentication =  "VIPAuthentication"

//由用户提交的原始数据
type CredentialSubject struct {
	ID string					`json:"id"`	//被签发方的ID
	ShortDescription string		`json:"shortDescription"`
	LongDescription	string		`json:"longDescription"`
	TypeCliam string			`json:"typeClaim"` //IDCardAuthentication
}

type Proof struct {
	Creator string 			`json:"creator"`		//发证方
	ChainAddr string		`json:"chainAddr"`		//发证方的区块链地址
	EncryptionType string	`json:"type"`			//加密类型，Keccak256
	SignatureValue string	`json:"signatureValue"`	//签名结果，对CredentialSubject，签名然后再base58编码
}

type ProofClaim struct {
	Id string					`json: "id"`		//ProofCliam ID 全局中唯一
	Issuer string				`json: "issuer"`	//发证方
	IssuanceDate string			`json: "issuance_date"`
	ExpirationDate string		`json: "expiration_date"`
	Result uint					`json: "result"`
	CredentialSubject CredentialSubject 	`json: "credential_subject"`
	Proof Proof					`json: "proof"`
}

func CredentialSubjectUnmarshal(CredentialSubjectStr string, credentialSubject *CredentialSubject) error {
	return json.Unmarshal([]byte(CredentialSubjectStr), &credentialSubject)
}

func CredentialSubjectMarshal(credentialSubjectEntry CredentialSubject) ([]byte, error)  {
	return json.Marshal(credentialSubjectEntry)
}

func ProofClaimUnmarshal(ProofClaimStr string, claim *ProofClaim) error {
	return json.Unmarshal([]byte(ProofClaimStr), &claim)
}

func ProofClaimMarshal(proofClaimEntry ProofClaim) ([]byte, error) {
	return json.Marshal(proofClaimEntry)
}
