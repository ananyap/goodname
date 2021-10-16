package managers

import (
	"sort"
	"strconv"

	"github.com/ananyap/goodname/domains"
)

type MatchingManager struct {
	myName   domains.NameDomain
	numList  []domains.Number
	realList []domains.RealName
}

func NewMatchingManager(myName domains.NameDomain, numList []domains.Number, realList []domains.RealName) *MatchingManager {
	return &MatchingManager{
		myName: myName, numList: numList, realList: realList,
	}
}

func (manager *MatchingManager) MatchingLekSatAndSha() (namesMatched domains.NamesMatchDomain) {

	var listRealNameD []domains.RealName

	myLeksat, _ := strconv.Atoi(manager.myName.LekSat)
	mySha, _ := strconv.Atoi(manager.myName.Shadow)

	realNameList := manager.realList
	numMiracle := manager.numList

	for _, nameObj := range realNameList {

		xPlus := myLeksat + nameObj.LeksatThai
		shaPlus := mySha + nameObj.Shadow

		for _, numMira := range numMiracle {

			if strconv.Itoa(xPlus) == numMira.PairNumber {
				if numMira.PairType == "D10" || numMira.PairType == "D8" || numMira.PairType == "D5" {

					for _, numMiraObjCheckeShaD := range numMiracle {

						if strconv.Itoa(shaPlus) == numMiraObjCheckeShaD.PairNumber {

							if numMira.PairType == "D10" || numMira.PairType == "D8" || numMira.PairType == "D5" {

								kManager := NewKalakiniManager(manager.myName.NameTitle, manager.myName.DayBirth)

								listRealNameD = append(listRealNameD, domains.RealName{
									ThaiName:           nameObj.ThaiName,
									ReangThai:          nameObj.ReangThai,
									LeksatThai:         nameObj.LeksatThai,
									Shadow:             nameObj.Shadow,
									PairPoint:          numMira.PairPoint,
									MatchedLekSatPlusD: xPlus,
									MatchedShaNumPlusD: shaPlus,
									SurNameKalakini:    kManager.Kalakini(),
								})
							}
							break
						}
					}

				}

				break
			}

		}
	}

	namesMatched.MyName = manager.myName
	namesMatched.RealNameList = listRealNameD

	sort.SliceStable(namesMatched.RealNameList, func(i, j int) bool {
		return namesMatched.RealNameList[i].PairPoint > namesMatched.RealNameList[j].PairPoint
	})

	return namesMatched
}
