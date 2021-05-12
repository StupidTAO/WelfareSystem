package web

import (
	"fmt"
	"testing"
)

func TestGetInfoPage(t *testing.T) {
	did := "did:welfare:yHB56rdByDPpghbcLi13W3SaXYU"
	pageInfo, err := GetInfoPage(did)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(pageInfo)
}
