package main

import (
	"encoding/json"

	"github.com/ananyap/goodname/domains"
	"github.com/ananyap/goodname/managers"
)

func main() {
	nameTitle := "พีระพงศ์"
	dayBirth := "monday"
	nameManager := managers.NewNameManager(nameTitle, dayBirth)
	nameManager.Kalakini()
	myName := domains.NameDomain{
		NameTitle: nameTitle,
		DayBirth:  dayBirth,
		LekReang:  nameManager.ReangName(),
		LekSat:    nameManager.LekSat(),
		Shadow:    nameManager.Shadow(),
		Kalakini:  nameManager.Kalakini(),
	}

	resp, _ := json.Marshal(myName)
	print(string(resp))
}
