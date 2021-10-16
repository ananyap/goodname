package main

import (
	"database/sql"
	"encoding/json"

	"github.com/ananyap/goodname/domains"
	"github.com/ananyap/goodname/managers"
	_ "github.com/go-sql-driver/mysql"
)

func DBConnect() (*sql.DB, error) {
	//db, err := sql.Open("mysql", "ananya:In^telliP24.pa@tcp(localhost:3306)/ananyadb")
	db, err := sql.Open("mysql", "root:IntelliP24@tcp(localhost:3306)/ananyadb")
	return db, err
}

func main() {
	db, err := DBConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	nameTitle := "อณัญญา"
	dayBirth := "sunday"
	nameManager := managers.NewNameManager(nameTitle, dayBirth)
	kalakiniManager := managers.NewKalakiniManager(nameTitle, dayBirth)
	myName := domains.NameDomain{
		NameTitle: nameTitle,
		DayBirth:  dayBirth,
		LekReang:  nameManager.ReangName(),
		LekSat:    nameManager.LekSat(),
		Shadow:    nameManager.Shadow(),
		Kalakini:  kalakiniManager.Kalakini(),
	}

	dbManager := managers.NewDbManager(db)
	dbManager.SetRealNames()
	dbManager.SetNumbersMiracle()

	realList := dbManager.GetRealNames()    //[]domains.RealName
	numList := dbManager.GetNumberMiracle() //[]domains.Number

	miracleManager := managers.NewMiracleManager(myName.LekSat, numList)
	myName.NumberMiracle = miracleManager.NumMiracle()

	matchingManager := managers.NewMatchingManager(myName, numList, realList)
	listNameMatched := matchingManager.MatchingLekSatAndSha()

	//resp, _ := json.Marshal(myName)
	//print(string(resp))

	resMathing, _ := json.Marshal(listNameMatched)
	print(string(resMathing))

}
