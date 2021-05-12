package router

import (
	"fmt"
	hub "github.com/StupidTAO/DIDHub/model"
	"github.com/WelfareSystem/log"
	"github.com/WelfareSystem/model"
	"github.com/WelfareSystem/utils"
	"github.com/WelfareSystem/web"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome!")
}

func Transfer(w http.ResponseWriter, r *http.Request) {
	//读取数据
	r.ParseForm()
	welfareDonate := r.Form["rawDonate"][0]
	signDonate := r.Form["sigDonate"][0]
	addr := r.Form["addr"][0]
	signDonateBytes := utils.Base58Decode(signDonate)
	//验证签名
	b, _ := VerifyClaim(welfareDonate, signDonateBytes, addr)
	if b {
		log.Info("signature verify pass")
	}

	var wdEntry, err = GetWelfareDonate(welfareDonate)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	//根据claimID获取claim内容
	//claimId := "48pVXCQk37Rzf9mMWbEboaaoQKJ3"
	claims, err := hub.FindHubDIDClaim(wdEntry.ClaimId)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	if len(claims) == 0 {
		fmt.Fprintln(w, "the claimId is not exsit")
		return
	}

	claim := claims[0]
	var proofClaim model.ProofClaim
	err = model.ProofClaimUnmarshal(claim.DidClaim, &proofClaim)
	if err != nil {
		fmt.Fprintln(w, "proof claim is not complete data")
		return
	}

	//获取发证方区块链地址，并对签名部分做验签
	addr = proofClaim.Proof.ChainAddr
	signText := utils.Base58Decode(proofClaim.Proof.SignatureValue)
	originText, err := model.CredentialSubjectMarshal(proofClaim.CredentialSubject)
	if err != nil {
		fmt.Fprintln(w, "proof claim CredentialSubject data is have problem")
		return
	}
	b, err = VerifyClaim(string(originText), []byte(signText), addr)
	if b {
		log.Info("age is older than 10 years, allow donations")
	} else {
		log.Info("proof claim signature verify failed, err: %s", err.Error())
		return
	}

	//获取捐款额度，并将交易上链
	from := claim.Did
	to := "did:welfare:123456789abcdefghijklmnopqrs"
	tx := createTx(from, to, uint32(wdEntry.Amount), wdEntry.Priority)
	hub.InsertHubTransaction(tx)

	fmt.Fprintln(w, "donate success, tx id is: ", tx.TxId)
}

//验证客户端发来的声明数据
func VerifyClaim(rawDonate string, signDonate []byte, addr string) (bool, error) {
	addrDerive, err := utils.VerifyToAddress(rawDonate, signDonate)
	if err != nil {
		return false, err
	}
	if addrDerive != addr {
		return false, err
	}

	return true, nil
}

//根据参数创建交易
func createTx(from string, to string, amount uint32, projectPriority float32) hub.Transaction {
	tx := new(hub.Transaction)

	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(10000)
	randStr := strconv.Itoa(randInt)
	id := utils.Base58Encode(utils.GetSHA256HashCode([]byte(randStr)))
	tx.TxId = id
	tx.FromAddr = from
	tx.ToAddr = to
	tx.Amount = amount
	tx.ProjectPriority = projectPriority
	tx.Contribution = 0
	tx.CreateTime = time.Now().Add((24*14+8) * time.Hour)
	tx.UpdateTime = time.Now().Add((24*14+8) * time.Hour)
	tx.HasCaculate = 0
	tx.Payload = "donate"
	return *tx
}

func GetWelfareDonate(wf string) (model.WelfareDonate, error) {
	wdEntry := new(model.WelfareDonate)
	err := model.WelfareDonateUnmarshal(wf, wdEntry)
	if err != nil {
		return model.WelfareDonate{}, err
	}
	return *wdEntry, nil
}

func InfoPage(w http.ResponseWriter, r *http.Request) {
	//读取参数
	r.ParseForm()
	did := r.Form["did"][0]
	//调用函数获取页面
	fmt.Println("info page did is: ", did)
	info, err := web.GetInfoPage(did)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	w.Write([]byte(info))
}
