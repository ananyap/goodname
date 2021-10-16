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

	nameTitle := "พีระพงศ์"
	dayBirth := "monday"
	nameManager := managers.NewNameManager(nameTitle, dayBirth)
	myName := domains.NameDomain{
		NameTitle: nameTitle,
		DayBirth:  dayBirth,
		LekReang:  nameManager.ReangName(),
		LekSat:    nameManager.LekSat(),
		Shadow:    nameManager.Shadow(),
		Kalakini:  nameManager.Kalakini(),
	}

	dbManager := managers.NewDbManager(db)
	dbManager.SetRealNames()
	dbManager.SetNumbersMiracle()

	realList := dbManager.GetRealNames() //[]domains.RealName
	numList := dbManager.GetNumberMiracle() //[]domains.Number

	miracleManager := managers.NewMiracleManager(myName.LekSat, numList)
	myName.NumberMiracle = miracleManager.NumGoodOrBad()

	resp, _ := json.Marshal(myName)
	print(string(resp))

}
