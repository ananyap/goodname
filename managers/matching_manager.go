package managers

import "github.com/ananyap/goodname/domains"

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
