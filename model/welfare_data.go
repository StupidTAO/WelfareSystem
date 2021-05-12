package model

import "encoding/json"

type WelfareDonate struct {
	ClaimId	string
	Amount int32
	Priority float32
}

func WelfareDonateUnmarshal(WelfareDonateStr string, welfateDonate *WelfareDonate) error {
	return json.Unmarshal([]byte(WelfareDonateStr), welfateDonate)
}

func WelfareDonateMarshal(welfateDonate WelfareDonate) ([]byte, error)  {
	return json.Marshal(welfateDonate)
}

