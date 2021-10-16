package managers

import (
	"database/sql"

	"github.com/ananyap/goodname/domains"
)

type DbManager struct {
	db             *sql.DB
	realNames      []domains.RealName
	numbersMiracle []domains.Number
}

func NewDbManager(myDb *sql.DB) *DbManager {
	return &DbManager{db: myDb}
}

func (manager *DbManager) SetNumbersMiracle() {

	results, err := manager.db.Query("SELECT pairnumber, pairtype, pairpoint FROM numbers order by pairnumberid ")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var myNumber domains.Number
		err = results.Scan(&myNumber.PairNumber, &myNumber.PairType, &myNumber.PairPoint)
		if err != nil {
			panic(err.Error())
		}

		manager.numbersMiracle = append(manager.numbersMiracle, myNumber)

	}
}

func (manager *DbManager) GetNumberMiracle() []domains.Number {
	return manager.numbersMiracle
}

func (manager *DbManager) SetRealNames() {

	results, err := manager.db.Query("SELECT thainame, reangthai, leksat_thai, shadow FROM realname order by nameid")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var myName domains.RealName
		err = results.Scan(&myName.ThaiName, &myName.ReangThai, &myName.LeksatThai, &myName.Shadow)
		if err != nil {
			panic(err.Error())
		}

		manager.realNames = append(manager.realNames, myName)

	}
}

func (manager *DbManager) GetRealNames() []domains.RealName {
	return manager.realNames
}
