package web

import (
	"errors"
	"fmt"
	hub "github.com/StupidTAO/DIDHub/model"
	"github.com/ethereum/go-ethereum/common"
)

const INFO_PAGE = "<html lang=\"en\">\n\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=<device-width>, initial-scale=1.0\">\n    <meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\">\n    <title>DID信息详情</title>\n    <base target=\"_blank\">\n</head>\n\n<body background=\"\">\n    <table border=\"1\" align=\"center\" cellspacing=\"0\" cellpadding=\"10\" width=\"800\" height=\"70\">\n        <tr>\n            <th colspan=\"4\" bgcolor=\"BuryWood\">\n                基本信息\n            </th>\n        </tr>\n        <tr>\n            <th>did地址:</th>\n            <td>%s</td>\n            <th>公益积分:</th>\n            <td>%f</td>\n        </tr>\n        <tr>\n            <th>区块链地址:</th>\n            <td>%s</td>\n            <th>公益排名:</th>\n            <td>%d</td>\n        </tr>\n        <tr>\n            <th colspan=\"4\" bgcolor=\"BuryWood\">\n                投票信息\n            </th>\n        </tr>\n        <tr>\n            <th>票权:</th>\n            <td>%d</td>\n            <th>是否已投票:</th>\n            <td>%s</td>\n        </tr>\n        <tr>\n            <th>委托人地址:</th>\n            <td>%s</td>\n            <th>投票提案索引:</th>\n            <td>%d</td>\n        </tr>\n       \n    </table>\n</body>\n\n</html>\n"

func GetInfoPage(did string) (string, error){
	//基本信息
	//1.did对应的chain addr
	addrs, err := hub.FindDBDIDChainAddr(did)
	if err != nil {
		return "", err
	}
	if len(addrs) == 0 {
		return "", errors.New("the did is not exsit matched addr")
	}

	chainAddr := addrs[0].DidChainAddr
	//2.did对应的公益积分
	token, err := hub.GetWelfareToken(did)
	if err != nil {
		return "", errors.New("get welfare token occur error")
	}
	//3.公益积分排名
	rank := 3


	//投票信息
	address := common.HexToAddress(chainAddr)
	voteRight, vote, delegate, voteIndex, err := hub.ContractVoters(address)
	isVotes := ""
	if vote {
		isVotes = "是"
	} else {
		isVotes = "否"
	}
	infoPage := fmt.Sprintf(INFO_PAGE, did, token, chainAddr, rank, voteRight, isVotes, delegate, voteIndex)
	return infoPage, nil
}
