package managers

import (
	"github.com/ananyap/goodname/domains"
)

type MiracleManager struct {
	lekSat         string
	numbersMiracle []domains.Number
}

func NewMiracleManager(lekSat string, numbersMiracle []domains.Number) *MiracleManager {
	return &MiracleManager{lekSat: lekSat, numbersMiracle: numbersMiracle}
}

func (manager *MiracleManager) NumMiracle() (numberMiracle domains.Number) {
	for _, NumberOBJ := range manager.numbersMiracle {
		if manager.lekSat == NumberOBJ.PairNumber {
			numberMiracle = NumberOBJ
			break
		}
	}

	return
}
