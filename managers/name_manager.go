package managers

import (
	"bytes"
	"strconv"

	"github.com/sirupsen/logrus"
)

type NameManager struct {
	nameTitle     []rune
	dayBirth      string
	charOfSat     map[string]string
	charOfShadow  map[string]string
	dayOfKalikini map[string]string
}

func NewNameManager(nameTitle string, daybirth string) *NameManager {
	return &NameManager{
		nameTitle: []rune(nameTitle),
		dayBirth:  daybirth,
		charOfSat: map[string]string{
			"ก": "1", "ด": "1", "ถ": "1", "ท": "1", "ภ": "1", "ฤ": "1", "ฦ": "1", "า": "1", "ำ": "1", "ุ": "1", "่": "1",
			"ข": "2", "ง": "2", "ช": "2", "บ": "2", "ป": "2", "เ": "2", "แ": "2", "ู": "2", "้": "2",
			"ฆ": "3", "ต": "3", "ฑ": "3", "ฒ": "3", "๋": "3",
			"ค": "4", "ธ": "4", "ญ": "4", "ร": "4", "ษ": "4", "ะ": "4", "โ": "4", "ั": "4", "ิ": "4",
			"ฉ": "5", "ฌ": "5", "ณ": "5", "น": "5", "ม": "5", "ห": "5", "ฎ": "5", "ฮ": "5", "ฬ": "5", "ึ": "5",
			"จ": "6", "ล": "6", "ว": "6", "อ": "6", "ใ": "6",
			"ซ": "7", "ศ": "7", "ส": "7", "๊": "7", "ี": "7", "ื": "7",
			"ผ": "8", "ฝ": "8", "พ": "8", "ฟ": "8", "ย": "8", "็": "8",
			"ฏ": "9", "ฐ": "9", "ไ": "9", "์": "9",
			"a": "1", "i": "1", "j": "1", "q": "1", "y": "1",
			"A": "1", "I": "1", "J": "1", "Q": "1", "Y": "1",
			"b": "2", "k": "2", "r": "2",
			"B": "2", "K": "2", "R": "2",
			"c": "3", "g": "3", "l": "3", "s": "3",
			"C": "3", "G": "3", "L": "3", "S": "3",
			"d": "4", "m": "4", "t": "4",
			"D": "4", "M": "4", "T": "4",
			"e": "5", "h": "5", "n": "5", "x": "5",
			"E": "5", "H": "5", "N": "5", "X": "5",
			"u": "6", "v": "6", "w": "6",
			"U": "6", "V": "6", "W": "6",
			"o": "7", "z": "7",
			"O": "7", "Z": "7",
			"f": "8", "p": "8",
			"F": "8", "P": "8",
		},
		charOfShadow: map[string]string{
			"อ": "6", "ะ": "6", "า": "6", "ิ": "6", "ี": "6", "ุ": "6", "ู": "6", "เ": "6", "โ": "6",
			"ก": "15", "ข": "15", "ค": "15", "ฆ": "15", "ง": "15",
			"จ": "8", "ฉ": "8", "ช": "8", "ซ": "8", "ฌ": "8", "ญ": "8",
			"ฎ": "17", "ฏ": "17", "ฐ": "17", "ฑ": "17", "ฒ": "17", "ณ": "17",
			"บ": "19", "ป": "19", "ผ": "19", "ฝ": "19", "พ": "19", "ฟ": "19", "ภ": "19", "ม": "19",
			"ศ": "21", "ษ": "21", "ส": "21", "ห": "21", "ฬ": "21", "ฮ": "21",
			"ด": "10", "ต": "10", "ถ": "10", "ท": "10", "ธ": "10", "น": "10",
			"ย": "12", "ร": "12", "ล": "12", "ว": "12",
		},
		dayOfKalikini: map[string]string{"sunday": "ศษสหฬฮ", "monday": "ะ้ื์แ๊ั็าิึี่ำุูเใไโอ", "tuesday": "กขคฆง", "wednesday1": "จฉชซฌญ", "wednesday2": "บปผฝพฟภม", "thursday": "ดตถทธน", "friday": "ยรลว", "saturday": "ฎฏฐฑฒณ"},
	}
}

func (manager *NameManager) ReangName() string {
	var b bytes.Buffer
	namex := manager.nameTitle
	for _, name := range string(namex) {
		for k, number := range manager.charOfSat {
			if string(name) == k {
				b.WriteString(number)
				break
			}
		}
	}

	return b.String()
}

//Sum of LekReang
func (manager *NameManager) LekSat() string {
	var leksatNum = 0
	for _, charX := range string(manager.ReangName()) {
		num, _ := strconv.Atoi(string(charX))
		leksatNum += num
	}
	return strconv.Itoa(leksatNum)
}

func (manager *NameManager) Shadow() string {
	shadow := 0
	namex := manager.nameTitle
	for _, name := range string(namex) {
		for k, number := range manager.charOfShadow {
			if string(name) == k {
				numShar, _ := strconv.Atoi(number)
				shadow += numShar
				break
			}
		}
	}

	return strconv.Itoa(shadow)

}

func (manager *NameManager) Kalakini() (kalakinis []string) {

	nameTitleRune := manager.nameTitle

	for day, kalakiniList := range manager.dayOfKalikini {

		if manager.dayBirth == day {

			for _, charx := range string(nameTitleRune) {

				for _, kalakiniCharx := range kalakiniList {

					if charx == kalakiniCharx {
						logrus.Println("KALAKINI", string(charx))
						kalakinis = append(kalakinis, string(charx))
						break
					}
				}
			}
			break
		}
	}

	return

}
